package query

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kisielk/sqlstruct"
)

func isField(key string) bool {
	if key == "sort_by" {
		return false
	}
	if key == "limit" {
		return false
	}
	if key == "page" {
		return false
	}
	if key == "conjunction" {
		return false
	}
	return true
}

func translateOp(op string) string {
	if op == "LIKE" {
		return "~~"
	}
	if op == "NOT_LIKE" {
		return "!~~"
	}
	if op == "REGEX_I" {
		return "~*"
	}
	if op == "REGEX" {
		return "~"
	}
	if op == "NOT_REGEX_I" {
		return "!~*"
	}
	if op == "NOT_REGEX" {
		return "!~"
	}
	if op == "NE" {
		return "<>"
	}
	if op == "LTE" {
		return "<="
	}
	if op == "LT" {
		return "<"
	}
	if op == "GTE" {
		return ">="
	}
	if op == "GT" {
		return ">"
	}
	return "="
}

type QueryConfig struct {
	Pattern string         `json:"pattern"`
	Args    map[string]any `json:"args"`
	Limit   string         `json:"limit"`
	SortBy  string         `json:"sort_by"`
	Page    string         `json:"page"`
}

// Query value structure based on elements:
// [1] value only assumes string:
//
//	name=Billy+Goat ==>  { pattern: 'name = ?', param: "Billy Goat" }
//
// [2] value and operator assumes string type is string:
//
//	name="LIKE;Billy" ==> { pattern: 'name LIKE ?', param: "Billy" }
func NewQueryConfig(queryArgs url.Values, model any) (q *QueryConfig, err error) {
	types, err := reflectStruct(model)
	if err != nil {
		return
	}
	conjunction := queryArgs.Get("conjunction")
	if conjunction != " OR " {
		conjunction = " AND "
	}
	pfx := ""

	q = &QueryConfig{
		Args: make(map[string]any),
	}
	for key, values := range queryArgs {
		if !isField(key) {
			continue
		}
		for _, value := range values {
			parts := strings.SplitN(value, "|", 2)
			op := "="
			val := ""
			typ := types[key]
			if len(parts) == 1 {
				val = parts[0]
			}
			if len(parts) == 2 {
				val = parts[1]
				op = translateOp(parts[0])
			}
			q.Pattern += fmt.Sprintf("%s%s %s ?", pfx, key, op)
			pfx = conjunction
			var v any
			v, err = convertString(val, typ)
			if err != nil {
				return
			}
			q.Args[key] = v
		}
	}

	q.Limit = queryArgs.Get("limit")
	q.SortBy = queryArgs.Get("sort_by")
	q.Page = queryArgs.Get("page")
	return
}

func convertString(val string, typ string) (converted any, err error) {
	switch typ {
	case "int64":
		converted, err = strconv.ParseInt(val, 10, 64)
	case "string":
		converted = val
	case "uuid":
		converted, err = uuid.Parse(val)
	case "time":
		var i int64
		i, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
		converted = time.UnixMilli(i)
	}
	return
}

func reflectStruct(model any) (types map[string]string, err error) {
	typ := reflect.TypeOf(model)
	if typ.Kind() != reflect.Struct {
		err = fmt.Errorf("%s is not a struct", typ)
		return
	}
	types = make(map[string]string)
	for i := 0; i < typ.NumField(); i++ {
		fld := typ.Field(i)
		snc := sqlstruct.ToSnakeCase(fld.Name)
		qt := fld.Tag.Get("qt")
		if qt != "" {
			parts := strings.SplitN(qt, ";", 2)
			// Formats:
			// type SomeStruct struct {
			//    SomeId    uuid.UUID  `qt:"field_name;uuid"`  //- uses field_name instead of some_id for field name, uuid for type
			//    SomeTime  time.Time  `qt:"time"`             //- uses snake cased some_time for field name time.Time for type
			// }
			if len(parts) > 1 {
				types[parts[1]] = parts[0]
			} else {
				types[snc] = parts[0]
			}
			continue
		}
		types[snc] = "string"
	}
	return
}

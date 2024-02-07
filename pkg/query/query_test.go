package query

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/google/uuid"
)

type TestStruct struct {
	Id      uuid.UUID `qt:"uuid"`
	Created time.Time `qt:"time"`
	Count   int64     `qt:"test_struct_count;int64"`
	Name    string
}

func TestNewQueryConfig(t *testing.T) {
	someUuid, err := uuid.NewRandom()
	if err != nil {
		t.Fatalf("cant generate a uuid: %v", err)
		return
	}
	now := time.Now()
	start := now.Add(time.Hour * -2)

	queryString := fmt.Sprintf("conjunction=AND&id=NE|%s&created=LT|%d&created=GT|%d&test_struct_count=30&name=REGEX|%s&sort_by=created DESC",
		someUuid.String(), now.UnixMilli(), start.UnixMilli(), url.QueryEscape("[A-F][a-z]+.*"))
	values, err := url.ParseQuery(queryString)
	if err != nil {
		t.Fatalf("cant parse query string %s: %v", queryString, err)
		return
	}

	result, err := NewQueryConfig(values, TestStruct{})
	if err != nil {
		t.Fatalf("cant create new query config: %v", err)
		return
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("json marshal error: %v", err)
		return
	}
	fmt.Printf("Result: %s", string(bytes))
}

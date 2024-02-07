package logger

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

func preamble() (result string) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "?"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	return fmt.Sprintf("%s:%d %s", filepath.Base(file), line, fnName)

}

func Infof(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	log.Printf("INFO: %s: %s", preamble(), msg)
}

func Warnf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	log.Printf("WARN: %s: %s", preamble(), msg)
}

func Errorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	log.Printf("ERROR: %s: %s", preamble(), msg)
}
func Printf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	log.Printf("DEBUG: %s: %s", preamble(), msg)
}
func Println(msg string) {
	log.Printf("DEBUG: %s: %s", preamble(), msg)
}

package utilitys

import (
	"log"
	"runtime"
)

func LogError(err error) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.Printf("Error: %v, di file %s, baris ke %d\n", err, file, line)
	} else {
		log.Printf("Error: %v\n", err)
	}
}

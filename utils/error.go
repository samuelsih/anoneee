package utils

import (
	"fmt"
	"log"
	"runtime"
)

func CheckError(any interface{}, a ...interface{}) error {
	if any != nil {
		err := (error)(nil)

		if _, ok := any.(string); ok {
			err = fmt.Errorf(any.(string), a...)
		} else if _, ok := any.(error); ok {
			err = fmt.Errorf(any.(error).Error(), a...)
		} else {
			err = fmt.Errorf("%v", err)
		}

		_, fn, line, _ := runtime.Caller(1)
		log.Printf("ERROR: [%s:%d] %v \n", fn, line, err)

		return err
	}

	return nil
}

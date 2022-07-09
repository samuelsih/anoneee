package utils

import "errors"

func CustomErrReturn(err string) error {
	return errors.New(err)
}
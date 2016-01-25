package main

import (
	"errors"
	"fmt"
)

func FindFormat(code int) (string, error) {
	if val, ok := StandardFormats[code]; ok {
		return val, nil
	} else {
		return "", errors.New(fmt.Sprintf("code (%d) is not recognized as standard format\n\r", code))
	}
}

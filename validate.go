package main

import (
	"errors"
	"fmt"
	"time"
)

func Validate(date string) (int, error) { // Change error to bool????
	for code, format := range StandardFormats {
		_, err := time.Parse(format, date)
		if err == nil {
			return code, nil
		}
	}

	return 99, errors.New(fmt.Sprintf("datetime (%s) is not recognized as standard format\n\r", date))
}

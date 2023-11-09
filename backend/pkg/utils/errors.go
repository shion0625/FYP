package utils

import (
	"errors"
	"fmt"
)

// To append the message to the error.
func AppendMessageToError(err error, message string) error {
	if err == nil {
		return errors.New(message)
	}

	return fmt.Errorf("%w \n%s", err, message)
}

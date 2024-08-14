package util

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"math/rand"
	"time"
)

func GenerateAccountNumber() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return fmt.Sprintf("%016d", r.Int63n(1e16))
}

func FormatValidationErrors(err error) map[string]string {
	errorArr := make(map[string]string)
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, err := range err.(validator.ValidationErrors) {
			errorArr[err.Field()] = err.Tag()
		}
	}

	return errorArr
}

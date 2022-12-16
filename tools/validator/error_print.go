package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func PrintError(err error) {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		fmt.Println("InvalidValidationError", err)
		return
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, validationError := range validationErrors {
			//fmt.Print(i, err)
			fmt.Print(validationError.Namespace())
			fmt.Print("-")
			fmt.Print(validationError.Field())
			fmt.Print("-")
			fmt.Print(validationError.StructNamespace())
			fmt.Print("-")
			fmt.Print(validationError.StructField())
			fmt.Print("-")
			fmt.Print(validationError.Tag())
			fmt.Print("-")
			fmt.Print(validationError.ActualTag())
			fmt.Print("-")
			fmt.Print(validationError.Kind())
			fmt.Print("-")
			fmt.Print(validationError.Type())
			fmt.Print("-")
			fmt.Print(validationError.Value())
			fmt.Print("-")
			fmt.Print(validationError.Param())
			fmt.Println()
		}
	}
}

package custom

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"testing"
)

type MyStruct struct {
	String string `validate:"custom-val"`
}

func TestCustomType(t *testing.T) {
	type DbBackedUser struct {
		Name sql.NullString `validate:"required"`
		Age  sql.NullInt64  `validate:"required"`
	}

	validate := validator.New()
	validate.RegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})

	//x := DbBackedUser{Name: sql.NullString{String: "", Valid: true}, Age: sql.NullInt64{Int64: 0, Valid: false}}
	x := DbBackedUser{Name: sql.NullString{String: "123", Valid: true}, Age: sql.NullInt64{Int64: 123, Valid: true}}

	err := validate.Struct(x)

	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
}

func ValidateValuer(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {
		val, err := valuer.Value()
		if err == nil {
			return val
		}
	}
	return nil
}

func TestCustomValidate(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("custom-val", CustomValidate)

	str := MyStruct{String: "yao"}
	err := validate.Struct(str)
	fmt.Println(err)
}

func CustomValidate(fl validator.FieldLevel) bool {
	return fl.Field().String() == "yqj"
}

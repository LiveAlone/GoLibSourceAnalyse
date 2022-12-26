package sample

import (
	"fmt"
	validator2 "github.com/LiveAlone/GoLibSourceAnalyse/tools/validator"
	"github.com/go-playground/validator/v10"
	"testing"
)

var validate *validator.Validate

func TestSample(t *testing.T) {
	validate = validator.New()
	validateLocalStruct()
	//validateVar()
}

func validateLocalStruct() {
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	err := validate.Struct(address)
	if err != nil {
		fmt.Println(err)
	}
	//validator2.PrintError(err)

	//user := &User{
	//	FirstName:      "Badger",
	//	LastName:       "Smith",
	//	Age:            135,
	//	Email:          "Badger.Smith@gmail.com",
	//	FavouriteColor: "#000-",
	//	Addresses:      []*Address{address},
	//}
	//err := validate.Struct(user)
	//validator2.PrintError(err)
}

func validateVar() {
	myEmail := "joeybloggs.gmail.com"
	//myEmail := "yaoqijunmail@foxmail.com"
	err := validate.Var(myEmail, "required,email")
	validator2.PrintError(err)
}

// User contains user information
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`
	Addresses      []*Address `validate:"required,dive,required"`
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

package rules

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

type Data struct {
	Name    string
	Email   string
	Details *Details
}

type Details struct {
	FamilyMembers *FamilyMembers
	Salary        string
}

type FamilyMembers struct {
	FatherName string
	MotherName string
}

type Data2 struct {
	Name string
	Age  uint32
}

var validate = validator.New()

func TestRuleConfig(t *testing.T) {
	//validateStruct()

	validateStructNested()
}

func validateStructNested() {
	data := Data{
		Name:  "11sdfddd111",
		Email: "zytel3301@mail.com",
		Details: &Details{
			Salary: "1000",
		},
	}

	rules1 := map[string]string{
		"Name":    "min=4,max=6",
		"Email":   "required,email",
		"Details": "required",
	}

	rules2 := map[string]string{
		"Salary":        "number",
		"FamilyMembers": "required",
	}

	rules3 := map[string]string{
		"FatherName": "required,min=4,max=32",
		"MotherName": "required,min=4,max=32",
	}

	validate.RegisterStructValidationMapRules(rules1, Data{})
	validate.RegisterStructValidationMapRules(rules2, Details{})
	validate.RegisterStructValidationMapRules(rules3, FamilyMembers{})
	err := validate.Struct(data)

	fmt.Println(err)
}

func validateStruct() {
	data := Data2{
		Name: "leo",
		Age:  1000,
	}

	rules := map[string]string{
		"Name": "min=4,max=6",
		"Age":  "min=4,max=6",
	}

	validate.RegisterStructValidationMapRules(rules, Data2{})

	err := validate.Struct(data)
	fmt.Println(err)
	fmt.Println()
}

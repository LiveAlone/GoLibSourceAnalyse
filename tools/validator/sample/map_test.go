package sample

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	data := map[string]interface{}{
		"name":  "Arshiya Kiani",
		"email": "zytel3301@gmail.com",
		"details": map[string]interface{}{
			"family_members": map[string]interface{}{
				"father_name": "Micheal",
				"mother_name": "Hannah",
			},
			"salary": "1000",
			"phones": []map[string]interface{}{
				{
					"number": "11-111-1111",
					"remark": "home",
				},
				{
					"number": "22-222-2222",
					"remark": "work",
				},
			},
		},
	}

	// Rules must be set as the structure as the data itself. If you want to dive into the
	// map, just declare its rules as a map
	rules := map[string]interface{}{
		"name":  "min=4,max=32",
		"email": "required,email",
		"details": map[string]interface{}{
			"family_members": map[string]interface{}{
				"father_name": "required,min=4,max=32",
				"mother_name": "required,min=4,max=32",
			},
			"salary": "number",
			"phones": map[string]interface{}{
				"number": "required,min=4,max=32",
				"remark": "required,min=1,max=32",
			},
		},
	}

	fmt.Println(validate.ValidateMap(data, rules))
}

func TestMapValidate(t *testing.T) {
	user := map[string]interface{}{"name": "Arshiya Kiani", "email": "zytel3301@gmail.com"}
	rules := map[string]interface{}{"name": "required,min=8,max=32", "email": "omitempty,required,email"}

	errs := validate.ValidateMap(user, rules)

	if len(errs) > 0 {
		fmt.Println(errs)
	}
}

package dive

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

type Test struct {
	Array []string          `validate:"required,gt=0,dive,required"`
	Map   map[string]string `validate:"required,gt=0,dive,keys,keymax,endkeys,required,max=1000"`
}

func TestDive(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("keymax", "max=10")

	var test Test
	test.Array = []string{"123"}
	test.Map = map[string]string{"test > than 10": ""}
	err := validate.Struct(test)
	fmt.Println(err)
}

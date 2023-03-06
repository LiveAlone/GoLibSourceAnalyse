package jsoniter

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func TestSerial(t *testing.T) {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := jsoniter.Marshal(group)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(b))
}

func TestOneLine(t *testing.T) {
	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	fmt.Println(jsoniter.Get(val, "Colors", 0).ToString())
}

func TestMapSort(t *testing.T) {
	m := map[string]interface{}{
		"3": 3,
		"1": 1,
		"2": 2,
	}
	//b, err := json.Marshal(m)
	//fmt.Println(string(b), err)
	//b, err := jsoniter.Marshal(m)
	//fmt.Println(string(b), err)

	// 兼容map key 顺序
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(m)
	fmt.Println(string(b), err)
}

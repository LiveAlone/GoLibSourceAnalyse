package model

import "fmt"

var num = 100

func SelectOne(id string) string {
	return "YQJ"
}

func SelectList() string {
	return "None"
}

type UserDao struct{}

func (d *UserDao) Update(id, name, phoneNumber string) int64 {
	fmt.Println("userDao is update")
	return 100
}

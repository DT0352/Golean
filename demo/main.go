package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	type User struct {
		UserId   int    `json:"user_id" bson:"user_id" form:"msg"`
		UserName string `json:"user_name" bson:"user_name " form:"msg"`
	}
	u := &User{
		UserId:   10,
		UserName: "哈哈",
	}
	j, _ := json.Marshal(u)
	fmt.Println(string(j))
	t := reflect.TypeOf(u)
	field := t.Elem().Field(0)
	fmt.Println(field.Tag.Get("json"))
}

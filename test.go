package main

import (
	"encoding/json"
	"fmt"
)

func Foo() {
	var arr1 []string
	fmt.Println(arr1)
	var arr2 []string
	fmt.Println(arr2)
}

func Bar() {
	type Person struct {
		Name    string    `json:"name"`
		Age     int       `json:"age"`
		Hobbies *[]string `json:"hobbies"`
	}

	// 定义一个空切片
	var hobbies *[]string
	//var hobbies = make([]string, 0)
	// 构造一个 Person 对象
	person := Person{
		Name:    "Alice",
		Age:     25,
		Hobbies: hobbies,
	}

	// 将 Person 对象转换为 JSON 数据
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("JSON marshaling failed:", err)
		return
	}

	// 打印 JSON 数据
	fmt.Println(string(jsonBytes))
}

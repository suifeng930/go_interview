package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	name string `json:"name"`  // 小写属性不能被转码 包可见
}

func main() {

	js := `{ "name":"11l"}`

	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err!=nil {
		fmt.Println("err: ",err)
		return
	}
	fmt.Println("people: ",p)  // people:  {}
}

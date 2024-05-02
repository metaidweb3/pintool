package main

import (
	"fmt"
	"self-tool/tool"
	"testing"
)

func TestJson(t *testing.T) {
	data := make(map[string]interface{})
	data["pk"] = "1"
	data["address"] = "2"
	data["net"] = "3"
	tool.WriteJSON("./setting.json", data)
	d, err := tool.ReadJSON("./setting.json")
	fmt.Println(err, d)
}

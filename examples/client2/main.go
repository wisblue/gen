//go:generate gen client github.com/wzshiming/gen/examples/route2

package main

import (
	"fmt"
	"time"
)

func main() {
	Client.SetURLByStr("http://127.0.0.1:8080")
	item := &Item{
		Name:    "hello",
		Message: "Time " + time.Now().String(),
	}
	itemID, err := CreateItem(1, item)
	if err != nil {
		fmt.Println(err)
		return
	}

	ret, err := GetItem(1, itemID)
	if err != nil {
		fmt.Println(err)
		return
	}
	if ret.Name != item.Name || ret.Message != item.Message {
		fmt.Println("error")
		return
	}
	fmt.Println("ok")
}

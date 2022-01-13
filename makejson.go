package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string

	userMap := make(map[string]string)

	fmt.Print("Enter a name: ")
	fmt.Scan(&name)

	fmt.Print("Enter an address: ")
	fmt.Scan(&address)

	userMap["name:"+name] = "address:" + address

	jsonUserMap, _ := json.Marshal(userMap)

	fmt.Print(string(jsonUserMap))
}

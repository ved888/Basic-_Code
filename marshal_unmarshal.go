package main

import (
	"encoding/json"
	"fmt"
)

type ourMap struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func main() {
	//var myMap = make(map[string]string)
	//var myMap = map[string]string{}
	//myMap["name"] = "amar"
	//myMap["phone"] = "1236547894"
	var myMap = map[string]string{"name": "Amar", "phone": "987564213"}
	fmt.Println(myMap)

	//jsonData, _ := json.Marshal(myMap)
	jsonData, err := json.Marshal(myMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Json string from map ", string(jsonData))
	fmt.Println("Now unmarshaling the json string")
	var structData ourMap
	json.Unmarshal([]byte(jsonData), &structData)

	fmt.Println(" unmarshal data ", structData)

}

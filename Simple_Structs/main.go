package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isadult"`
}

func Marshalling() {

}

func main() {
	person := Person{Name: "john", Age: 20, IsAdult: true}
	fmt.Print(person)
	//person into json (Marshalling) form:=slice
	marshall, err := json.Marshal(person)
	if err != nil {
		fmt.Print("something Went to Wrong")
	}
	fmt.Print(marshall)
	fmt.Print(string(marshall))
	var dummy Person
	err = json.Unmarshal(marshall,&dummy)
	fmt.Print(dummy)
	Marshalling()
}

//json to Person (unMarshalling)

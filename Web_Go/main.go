package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetData() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	defer res.Body.Close()
	if err != nil {
		fmt.Print("Something Went to Wrong")
		return
	}
	fmt.Println(res.Body)
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print("Something Went to Wrong")
		return
	}
	fmt.Print(string(data))

}

func main() {
	fmt.Print("hello")
	GetData()
}

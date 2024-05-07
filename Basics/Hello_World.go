package main

import (
	"fmt"
	"os"
)

var m int = 12

func File() {
	file, err := os.Create("Sample.txt")
	if err != nil {
		fmt.Print("Created", err)
		return
	}
	defer file.Close()

	ofile, err := os.Open("Sample.txt")
	if err != nil {
		fmt.Print("opened", err)
		return

	}
	defer ofile.Close()

}
func Writing_Content() {
	file, err := os.Open("Sample.txt")
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Print("Somthing Went to Wrong")
	}
	content := "Hey its me"
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Print("Somthing Went to Wrong")
	}
	fileSize := fileInfo.Size()

	// Create a byte slice to read the file content
	contentt := make([]byte, fileSize)

	// Read the content of the file into the byte slice
	_, err = file.Read(contentt)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(contentt)
}

func main() {
	fmt.Println("Jay Bajrang")
	var i int = 42
	fmt.Print(i, " ")

	//a := 25
	fmt.Print(m, " ")
	var m int
	m = int(i)
	fmt.Println()
	fmt.Printf("%T,%v ", m, m)
	fmt.Printf("%v ", m)

	a := 21
	b := 34
	fmt.Println(a + b)
	File()
	Writing_Content()

}

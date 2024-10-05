package main

import "fmt"

func main() {
	var printValue string = "Hello world"
	printme(printValue)
}

func printme(printValue string) {
	fmt.Println(printValue)
}

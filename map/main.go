package main

import "fmt"


func main() {
	colors:= map[string]string{
		"red":   "\033[91m",
        "green": "\033[92m",
        "yellow": "\033[93m",
        "blue": "\033[94m",
	}
	fmt.Println(colors)
}
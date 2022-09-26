package main

// Package main indicates an executable type package, opposed to "helper". If main isn't defined go build won't generate an executable.

import "fmt"

// standard go libraries https://pkg.go.dev/std

func main() {
	fmt.Println("Hello World")
}

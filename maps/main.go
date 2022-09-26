package main

import "fmt"

func main() {
	//1
	// map_def := map[string]string{
	// 	"red":   "blue",
	// 	"green": "yellow",
	// }
	//2
	//var map_def map[string]string
	//3
	// map_def := make(map[string]string)
	// map_def["red"] = "blue"

	// fmt.Println(map_def)

	//Iteration
	map_def := map[string]string{
		"red":   "blue",
		"green": "yellow",
		"blue":  "black",
	}
	iterateMap(map_def)

}

func iterateMap(maps map[string]string) {
	for a, c := range maps {
		fmt.Println("key = ", a, "Value = ", c)
	}

}
// Maps are passed as reference, structs by value. Meaning you can change the original value of map without pointer
// Like slices

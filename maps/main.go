package main

import "fmt"

func main() {

	//maps always need to matches types in keys and values
	//1 way to declare a map, with initial values
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"white": "#ffffff",
	}

	//2 way to declare map
	//var colors map[string]string

	//3 way to declare map
	//colors := make(map[string]string)

	colors["black"] = "#000000" // add value to map

	delete(colors, "black") // delete value from map

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}

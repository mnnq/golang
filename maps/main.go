package main

import "fmt"

func main() {
	/*
			//map[string-->keys type]string -->>values type
			colors := map[string]string{
				"red":   "#ff0000",
				"green": "#f33000",
			}
			var colors map[string]string
			colors := make(map[string]string)
			//way to add values
		colors["white"] = "#ffjjj333"
		colors["hite"] = "#ffjjj333"
		//delete items from map
		delete(colors, "white")
	*/

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#f33000",
		"white": "#f33000",
	}

	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for %v is %v \n", color, hex)
	}
}

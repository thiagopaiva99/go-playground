package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	fmt.Println(colors)

	newColors := make(map[string]string)
	newColors["white"] = "#ffffff"
	fmt.Println(newColors)
	delete(newColors, "white")
}

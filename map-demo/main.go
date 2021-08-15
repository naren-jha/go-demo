package main

import "fmt"

func main() {
	// method1
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
	}

	fmt.Println(colors) // map[green:#00FF00 red:#FF0000]

	// method 2
	var colors2 map[string]string // empty map
	fmt.Println(colors2)

	// method 3
	colors3 := make(map[string]string)
	colors3["white"] = "#FFFFFF"
	fmt.Println(colors3)

	delete(colors, "red")
	fmt.Println(colors) // map[green:#00FF00]
}

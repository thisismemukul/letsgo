package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// 	"blue":  "#0000ff",
	// }
	colors := make(map[string]string)
	colors["white"] = "#ffffff"
	colors["black"] = "#000000"
	printMap(colors)
	delete(colors, "black")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(
			"Color:", color, ", Hex:", hex,
		)
	}

}

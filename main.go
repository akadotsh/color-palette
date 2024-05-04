package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(palette())
}

func generateRandomColor() string {
	red := rand.Intn(256)
	blue := rand.Intn(256)
	green := rand.Intn(256)
	rgbString := fmt.Sprintf("rgb(%d,%d,%d)", red, blue, green)

	return rgbString
}

func palette() []string {
	palettes := make([]string, 0)

	for i := 0; i < 5; i++ {
		palettes = append(palettes, generateRandomColor())
	}

	return palettes
}

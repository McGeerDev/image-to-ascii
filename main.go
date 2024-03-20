package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func init() {
	// Registering the jpeg format so the image interface
	// Knows how to decode the jpeg image
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func main() {
	file, err := os.Open("image.jpeg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	// The second return value is the format of the image
	imgCfg, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	width := imgCfg.Width
	height := imgCfg.Height

	fmt.Println("Width:", width)
	fmt.Println("Height:", height)

	// Without this line there is a memory out of bounds error
	file.Seek(0, 0)

	img, _, err := image.Decode(file)

	destinationFile, err := os.Create("ascii.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer destinationFile.Close()

	for y := range height {
		for x := range width {
			r, g, b, a := img.At(x, y).RGBA()
			brightness := caclulateBrightness(r, g, b, a)
			char := getAsciiChar(brightness)
			destinationFile.WriteString(char)
		}
		destinationFile.WriteString("\n")
	}
	fmt.Println("Done")
}

func caclulateBrightness(r, g, b, a uint32) uint32 {
	r = (r >> 8) & 0xFF
	g = (g >> 8) & 0xFF
	b = (b >> 8) & 0xFF
	a = (a >> 8) & 0xFF
	return (r + g + b + a) / 4
}

func getAsciiChar(brightness uint32) string {
	asciiPixels := []string{"_", ".", ":", "-", "=", "+", "*", "#", "%", "@"}
	amountOfPixels := len(asciiPixels)

	pixel := float32(brightness) / 255 * float32(amountOfPixels-1)
	var pix int = int(pixel)

	return asciiPixels[pix]
}

package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func init() {
	// damn important or else At(), Bounds() functions will
	// caused memory pointer error!!
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func main() {

	file, err := os.Open("self.jpeg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	imgCfg, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	width := imgCfg.Width
	height := imgCfg.Height

	fmt.Println("Width:", width)
	fmt.Println("Height:", height)

	file.Seek(0, 0)

	img, _, err := image.Decode(file)

	fmt.Println(img.At(10, 10).RGBA())
	for y := range height {
		for x := range width {
			fmt.Println(img.At(x, y).RGBA())
			fmt.Printf("[%d, %d] ", x, y)
		}
	}
}

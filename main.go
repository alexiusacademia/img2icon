package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	ico "github.com/Kodeworks/golang-image-ico"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage img2icon <input_image> <output_image>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Failed to open input file: %v\n", inputFile)
		return
	}

	defer file.Close()

	var img image.Image

	switch strings.ToLower(strings.Split(inputFile, ".")[1]) {
	case "jpg", "jpeg":
		img, err = jpeg.Decode(file)
	case "png":
		img, err = png.Decode(file)
	default:
		fmt.Println("Unsupported image format. Please use JPG or PNG file formats.")
		return
	}

	if err != nil {
		fmt.Printf("Failed to decode input image: %v\n", inputFile)
	}

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Failed to create output file: %v\n", outputFile)
		return
	}

	defer outFile.Close()

	if err := ico.Encode(outFile, img); err != nil {
		fmt.Printf("Failed to encode image to icon file: %v\n", err)
	}
}

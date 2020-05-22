package main

import (
	"fmt"

	"github.com/fatih/color"
)

func logOutput(title, image, largeImage, videoLink string) {
	color.Green("Title:")
	fmt.Println(title)

	color.Green("Original image:")
	fmt.Println(image)

	color.Green("Large image:")
	fmt.Println(largeImage)

	color.Green("Video link:")
	fmt.Println(videoLink)
}

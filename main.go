package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math/rand"
	"os"
)

func main() {
	max := 256

	arr := initializeArray(max)
	arrShuffled := shuffle(arr)
	img := generateImage(arrShuffled)

	f, err := os.Create("test.gif")

	if err != nil {
		fmt.Printf("Error Occured! %q", err)
	}

	createNewGif(f, []*image.Paletted{img})

}

func generateImage(arr []int) *image.Paletted {
	palette := []color.Color{color.White, color.Black}

	rect := image.Rect(0, 0, len(arr), len(arr))
	img := image.NewPaletted(rect, palette)

	for k, v := range arr {
		fmt.Printf("K: %v, V: %v \n", v, v)
		img.SetColorIndex(k, len(arr)-v, uint8(1))
	}
	return img
}

func createNewGif(out io.Writer, imgs []*image.Paletted) {
	// Create a GIF for every
	images := []*image.Paletted{}
	delays := []int{}
	for _, v := range imgs {
		images = append(images, v)
		delays = append(delays, 10)
	}

	anim := gif.GIF{Delay: delays, Image: images}

	gif.EncodeAll(out, &anim)

}

func generateRandomImage(width, height int) *image.Paletted {
	palette := []color.Color{color.White, color.Black}

	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, palette)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := rand.Intn(2)
			img.SetColorIndex(x, y, uint8(c))
		}
	}

	return img
}

func initializeArray(length int) []int {
	arr := []int{}
	// This is for random numbers
	// for i := 0; i < length; i++ {
	// 	arr = append(arr, rand.Intn(maxNum))
	// }
	for i := 0; i < length; i++ {
		arr = append(arr, i)
	}
	return arr
}

func shuffle(arr []int) []int {
	for i := 0; i < 1000; i++ {
		temp1 := rand.Intn(len(arr))
		temp2 := rand.Intn(len(arr))

		temp := arr[temp1]

		arr[temp1] = arr[temp2]
		arr[temp2] = temp
	}
	return arr
}

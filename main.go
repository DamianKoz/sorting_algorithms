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
	max := 80

	arr := initializeArray(max)
	arrShuffled := shuffle(arr)
	frames := BubbleSort(arrShuffled)
	images := generateImages(frames)
	f, err := os.Create("test.gif")

	if err != nil {
		fmt.Printf("Error Occured! %q", err)
	}

	createNewGif(f, images)

}

// The Part about the visualizer
type FramesCollection struct {
	Arr [][]int
}

func (fc *FramesCollection) AddFrame(arr []int) {
	temp := make([]int, len(arr))
	copy(temp, arr)
	fc.Arr = append(fc.Arr, temp)
}

func generateImages(arr FramesCollection) (imgs []*image.Paletted) {
	for _, v := range arr.Arr {
		img := generateImage(v)
		imgs = append(imgs, img)

	}
	fmt.Print("Successfully created image array from single images.")
	return imgs
}

func generateImage(arr []int) *image.Paletted {
	multiplier := 256 / len(arr)
	palette := []color.Color{color.White, color.Black}

	rect := image.Rect(0, 0, 256, 256)
	img := image.NewPaletted(rect, palette)

	for k, v := range arr {
		img.SetColorIndex(k*multiplier, (len(arr)-v)*multiplier, uint8(1))
	}
	// for i := 0; i < len(arr); i++ {
	// 	img.SetColorIndex(i, len(arr)-arr[i], uint8(1))
	// }
	return img
}

func createNewGif(out io.Writer, imgs []*image.Paletted) {
	// Create a GIF for every
	images := []*image.Paletted{}
	delays := []int{}
	for _, v := range imgs {
		images = append(images, v)
		delays = append(delays, 0)
	}

	anim := gif.GIF{Delay: delays, Image: images}

	gif.EncodeAll(out, &anim)
	fmt.Print("Successfully created gif from single images.")

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

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
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

// Algorithms start here

func BubbleSort(arr []int) FramesCollection {
	swaped := false
	framesResult := FramesCollection{}
	framesResult.AddFrame(arr)

	for !swaped {
		swaped = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
				swaped = false
				framesResult.AddFrame(arr)

			}
			if isSorted(arr) {
				return framesResult
			}
		}
	}
	fmt.Print(framesResult)
	return framesResult
}

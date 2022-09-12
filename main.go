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

	current_algorithm := "insertionsort"
	numbers := 25

	createNewVisualisation(numbers, current_algorithm)
}

type fn func([]int) FramesCollection

func createNewVisualisation(max int, alg string) {
	algs := map[string]fn{
		"bubblesort":    BubbleSort,
		"insertionsort": InsertionSort,
	}

	if max == 0 {
		max = 15
	}

	arr := initializeArray(max)
	frames := algs[alg](shuffle(arr))
	images := generateImages(frames)

	file_path := "algorithm_gifs/" + alg + ".gif"
	f, err := os.Create(file_path)

	if err != nil {
		fmt.Printf("\nError Occured! %q \n", err)
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
	imageSize := 1000
	pixelSize := imageSize / len(arr) / 2
	multiplier := imageSize / len(arr)
	palette := []color.Color{color.White, color.Black}

	rect := image.Rect(0, 0, imageSize, imageSize)
	img := image.NewPaletted(rect, palette)

	for k, v := range arr {
		img.SetColorIndex(k*multiplier, (len(arr)-v)*multiplier, uint8(1))
		for i := -pixelSize; i < pixelSize; i++ {
			for j := -pixelSize; j < pixelSize; j++ {
				img.SetColorIndex(k*multiplier+i+(pixelSize), (len(arr)-v)*multiplier+j-(pixelSize), uint8(1))
			}

		}
	}
	return img
}

func createNewGif(out io.Writer, imgs []*image.Paletted) {
	images := []*image.Paletted{}
	delays := []int{}
	for _, v := range imgs {
		images = append(images, v)
		delays = append(delays, 1)
	}

	anim := gif.GIF{Delay: delays, Image: images}

	gif.EncodeAll(out, &anim)
	fmt.Print("\nSuccessfully created gif from single images.")

}

func initializeArray(length int) []int {
	arr := []int{}
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

func InsertionSort(arr []int) FramesCollection {
	framesResult := FramesCollection{}
	framesResult.AddFrame(arr)
	for i := 0; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j -= 1
			framesResult.AddFrame(arr)
		}
		framesResult.AddFrame(arr)
	}
	return framesResult
}

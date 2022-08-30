package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math/rand"
	"os"
)

func main() {
	width := 200
	height := 200

	f, _ := os.Create("test.gif")

	createNewGif(f, width, height)

}

func createNewGif(out io.Writer, width, height int) {
	images := []*image.Paletted{}
	for i := 0; i < 50; i++ {
		img := generateRandomImage(width, height)
		images := append(images, img)
	}

	anim := gif.GIF{Delay: []int{0}, Image: images}

	gif.EncodeAll(out, &anim)

}

func generateRandomImage(width, height int) *image.Paletted {
	palette := []color.Color{color.White, color.Black}

	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, palette)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := rand.Intn(2)
			img.Set(x, y, palette[c])
		}
	}

	img.Set(width/3, height/2, color.Black)
	return img
}

func initializeArray(maxNum, length int) []int {
	arr := []int{}
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(maxNum))
	}
	return arr
}

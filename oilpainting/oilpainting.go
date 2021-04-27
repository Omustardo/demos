// Based on http://supercomputingblog.com/graphics/oil-painting-algorithm/
package main

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

const (
	Radius = 6
	Bins   = 32
)

func OilPaint(img *image.NRGBA) *image.NRGBA {
	out := image.NewNRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))

	// Each pixel needs to be put in an intensity 'bin'.
	// Intensity is defined as (r+g+b)/3
	// Bin size of 20 is decent, so intensity bins start with 0-19, 20-39, 40-59, etc.
	// The bin is also affected by neighboring pixels, in an arbitrary radius (let's say 5).

	for x := 0; x < out.Bounds().Dx(); x++ {
		for y := 0; y < out.Bounds().Dy(); y++ {
			out.Set(x, y, calculatePixel(x, y, img))
		}
	}
	return out
}

func calculatePixel(x, y int, img *image.NRGBA) color.Color {
	// For nearby pixels, calculate intensity to determine which bin the target falls in.
	bins := make([]int, Bins)
	// For each bin, keep track of its colors so we can average it out later.
	sumCol := make([][3]int, Bins)

	for i := max(0, x-Radius); i < min(img.Bounds().Dx(), x+Radius); i++ {
		for j := max(0, y-Radius); j < min(img.Bounds().Dy(), y+Radius); j++ {
			r, g, b, _ := img.At(i, j).RGBA()

			// Normalize intensity (put it in a range from 0 to 1.0)
			// 0x101 comes from golang upscaling color bytes from 8 to 16 bits. https://blog.golang.org/go-image-package
			// 0xff is because they are originally 8 bit values
			normIntensity := (float32(r) + float32(g) + float32(b)) / 3 / 0x101 / 0xff
			// Convert the [0-1] range to a [0-NumBins) range.
			bin := min(int(normIntensity*float32(Bins)), Bins-1)
			bins[bin]++
			sumCol[bin][0] += int(r / 0x101)
			sumCol[bin][1] += int(g / 0x101)
			sumCol[bin][2] += int(b / 0x101)
		}
	}

	// Find bin with the highest pixel count.
	var maxIndex int
	for b := 1; b < len(bins); b++ {
		if bins[b] > bins[maxIndex] {
			maxIndex = b
		}
	}

	count := bins[maxIndex]
	r := uint8(sumCol[maxIndex][0] / count)
	g := uint8(sumCol[maxIndex][1] / count)
	b := uint8(sumCol[maxIndex][2] / count)
	return &color.NRGBA{R: r, G: g, B: b, A: 255}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	// Read in original image.
	in, err := os.Open(`C:\workspace\Go\src\github.com\omustardo\demos\oilpainting\headshot.jpg`)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	img, _, err := image.Decode(in)
	if err != nil {
		log.Fatal("Error decoding file: ", err)
	}
	in.Close()

	// To NRGBA
	m := image.NewNRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
	draw.Draw(m, m.Bounds(), img, img.Bounds().Min, draw.Src)

	// Modify the copy in place.
	for i := 0; i < 1; i++ {
		m = OilPaint(m)
	}

	// Write to file
	toImg, _ := os.Create(`C:\workspace\Go\src\github.com\omustardo\demos\oilpainting\out.png`)
	defer toImg.Close()
	png.Encode(toImg, m)
}

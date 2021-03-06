package util

import (
	"image"
	"image/png"
	"os"
	"time"

	"github.com/goxjs/gl"
)

// Why is this not in the standard time library?
func GetTimeMillis() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

// SaveScreenshot reads pixel data from OpenGL buffers, so it must be run in the same main thread as the rest
// of OpenGL.
// TODO: write to file in a goroutine and return a (chan err), or just ignore slow errors. Handling errors that can be
// caught immediately is fine. Blocking while writing to file adds way too much delay.
func SaveScreenshot(width, height int, path string) error {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	gl.ReadPixels(img.Pix, 0, 0, width, height, gl.RGBA, gl.UNSIGNED_BYTE)

	// Need to flip the image vertically since the pixels are provided with (0,0) in the top left corner.
	for row := 0; row < height/2; row++ {
		for col := 0; col < width; col++ {
			temp := img.At(col, row)
			img.Set(col, row, img.At(col, height-1-row))
			img.Set(col, height-1-row, temp)
		}
	}

	out, err := os.Create(path) // TODO: WebGL isn't happy with this (no syscalls allowed). Maybe store screenshots in memory or a temp file, and give an option to export?
	if err != nil {
		return err
	}
	return png.Encode(out, img)
}

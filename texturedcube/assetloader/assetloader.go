package assetloader

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png" // for PNG decoding

	"github.com/goxjs/gl"
)

// LoadTexture from local assets.
func LoadTexture(path string) (*gl.Texture, error) {
	// based on https://developer.mozilla.org/en-US/docs/Web/API/WebGL_API/Tutorial/Using_textures_in_WebGL and https://golang.org/pkg/image/

	fileData, err := loadFile(path)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewBuffer(fileData))
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// Need to flip the image vertically since OpenGL considers 0,0 to be the top left corner.
	flipImageVertically(img)

	// Image checking from https://github.com/go-gl-legacy/glh/blob/master/texture.go
	switch trueim := img.(type) {
	case *image.RGBA:
		return LoadTextureData(width, height, trueim.Pix), nil
	case *image.NRGBA: // What is NRGBA? It seems to act exactly like RGBA.
		return LoadTextureData(width, height, trueim.Pix), nil
	default:
		// copy := image.NewRGBA(trueim.Bounds())
		// draw.Draw(copy, trueim.Bounds(), trueim, image.Pt(0, 0), draw.Src)
		return nil, fmt.Errorf("unsupported texture format %T", img)
	}
}

func LoadTextureData(width, height int, data []uint8) *gl.Texture {
	// gl.Enable(gl.TEXTURE_2D) // some sources says this is needed, but it doesn't seem to be. In fact, it gives an "invalid capability" message in webgl.
	texture := gl.CreateTexture()
	gl.BindTexture(gl.TEXTURE_2D, texture)
	// NOTE: gl.FLOAT isn't enabled for texture data types unless gl.getExtension('OES_texture_float'); is set, so just use gl.UNSIGNED_BYTE
	//   See http://stackoverflow.com/questions/23124597/storing-floats-in-a-texture-in-opengl-es  http://stackoverflow.com/questions/22666556/webgl-texture-creation-trouble
	gl.TexImage2D(gl.TEXTURE_2D, 0, width, height, gl.RGBA, gl.UNSIGNED_BYTE, data) // TODO: Does layering RGBA images work? Or do we need to sort by Z value and draw in that order.
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.GenerateMipmap(gl.TEXTURE_2D)
	// gl.BindTexture(gl.TEXTURE_2D, gl.Texture{Value: 0}) // in js demo, they bind to null to prevent using the wrong texture by mistake. No way to do that with structs in Go, so just hope for the best?
	return &texture
}

func flipImageVertically(img image.Image) error {
	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	// Image checking from https://github.com/go-gl-legacy/glh/blob/master/texture.go
	switch trueim := img.(type) {
	case *image.RGBA:
		for row := 0; row < height/2; row++ {
			for col := 0; col < width; col++ {
				temp := img.At(col, row)
				trueim.Set(col, row, img.At(col, height-1-row))
				trueim.Set(col, height-1-row, temp)
			}
		}
	case *image.NRGBA: // What is NRGBA? It seems to act exactly like RGBA.
		for row := 0; row < height/2; row++ {
			for col := 0; col < width; col++ {
				temp := img.At(col, row)
				trueim.Set(col, row, img.At(col, height-1-row))
				trueim.Set(col, height-1-row, temp)
			}
		}
	default:
		return fmt.Errorf("unknown image type: %T", img)
	}
	return nil
}

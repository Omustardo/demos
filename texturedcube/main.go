// Reimplementation of Mozilla's textured cube demo, in Golang / Gopherjs
// https://developer.mozilla.org/en-US/docs/Web/API/WebGL_API/Tutorial/Using_textures_in_WebGL
package main

import (
	"flag"
	"log"
	"os"
	"time"

	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"github.com/omustardo/demos/texturedcube/assetloader"
	"github.com/omustardo/demos/texturedcube/camera"
	"github.com/omustardo/demos/texturedcube/cube"
	"github.com/omustardo/demos/texturedcube/shader"
	"github.com/omustardo/demos/texturedcube/view"
)

var (
	windowWidth  = flag.Int("window_width", 1000, "initial window width")
	windowHeight = flag.Int("window_height", 1000, "initial window height")

	frameRate = flag.Duration("framerate", time.Second/60, `Cap on framerate. Provide with units, like "16.66ms"`)
)

func init() {
	// log print with .go file and line number.
	log.SetFlags(log.Lshortfile)
	log.SetOutput(os.Stdout)
}

func main() {
	// Initialize gl constants and the glfw window. Note that this must be done before all other gl usage.
	if err := view.Initialize(*windowWidth, *windowHeight, "Spinning Textured Cube"); err != nil {
		log.Fatal(err)
	}
	defer view.Terminate()

	// Initialize Shaders
	if err := shader.Initialize(); err != nil {
		log.Fatal(err)
	}
	if err := gl.GetError(); err != 0 {
		log.Fatalf("gl error: %v", err)
	}

	// Load standard meshes.
	cube.Initialize()

	// Make a camera positioned at the origin, looking down negative Z.
	cam := camera.Camera{
		Pos:  mgl32.Vec3{0, 0, 0},
		Near: 0.1,
		Far:  1000,
	}

	tex, err := assetloader.LoadTexture("assets/sample_texture.png")
	if err != nil {
		log.Fatalf("error loading texture: %v", err)
	}
	texturedCube := cube.Cube{
		Center:   mgl32.Vec3{0, 0, -100},
		Dim:      mgl32.Vec3{32, 32, 32},
		Rotation: mgl32.Vec3{},
	}
	rotationPerSecond := float32(math.Pi / 4)

	ticker := time.NewTicker(*frameRate)
	for !view.Window.ShouldClose() {
		glfw.PollEvents() // Reads window events, like keyboard and mouse input.

		// Update the cube's X and Z rotation.
		texturedCube.Rotation[0] += rotationPerSecond * float32((*frameRate).Seconds())
		texturedCube.Rotation[2] += rotationPerSecond * float32((*frameRate).Seconds())

		// Set up Model-View-Projection Matrix and send it to the shader program. This could be moved outside of the loop as long as we never move the camera or resize the window.
		mvMatrix := cam.ModelView()
		w, h := view.Window.GetSize()
		pMatrix := cam.Projection(float32(w), float32(h))
		shader.Texture.SetMVPMatrix(pMatrix, mvMatrix)

		// Clear screen, then Draw everything
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		texturedCube.Draw(*tex)

		// Swaps the buffer that was drawn on to be visible. The visible buffer becomes the one that gets drawn on until it's swapped again.
		view.Window.SwapBuffers()
		<-ticker.C // wait up to the framerate cap.
	}
}

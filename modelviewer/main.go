// Demo of loading and displaying a DAE model.
package main

import (
	"flag"
	"log"
	"math"
	"os"
	"time"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"github.com/omustardo/demos/modelviewer/assetloader"
	"github.com/omustardo/demos/modelviewer/camera"
	"github.com/omustardo/demos/modelviewer/entity"
	"github.com/omustardo/demos/modelviewer/model"
	"github.com/omustardo/demos/modelviewer/shader"
	"github.com/omustardo/demos/modelviewer/view"
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
	if err := view.Initialize(*windowWidth, *windowHeight, "Model Viewer Demo"); err != nil {
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

	// Load model.
	mesh, err := assetloader.LoadDAE("assets/vehicle0.dae")
	if err != nil {
		log.Fatal(err)
	}
	model := &model.Model{
		Mesh: mesh,
		Entity: entity.Entity{
			Position: mgl32.Vec3{0, 0, 0},
			Rotation: mgl32.Vec3{},
			Scale:    mgl32.Vec3{1, 1, 1},
		},
	}

	cam := &camera.Camera{
		Target:       model,
		TargetOffset: mgl32.Vec3{0, 0, 50},
		Up:           mgl32.Vec3{0, 1, 0},
		Near:         0.1,
		Far:          10000,
		FOV:          math.Pi / 2.0,
	}

	rotationPerSecond := float32(math.Pi / 4)

	ticker := time.NewTicker(*frameRate)
	for !view.Window.ShouldClose() {
		glfw.PollEvents() // Reads window events, like keyboard and mouse input.

		// Update the X and Z rotation.
		model.Rotation[0] += rotationPerSecond * float32((*frameRate).Seconds())
		model.Rotation[2] += rotationPerSecond * float32((*frameRate).Seconds())

		// Set up Model-View-Projection Matrix and send it to the shader program.
		mvMatrix := cam.ModelView()
		w, h := view.Window.GetSize()
		pMatrix := cam.ProjectionPerspective(float32(w), float32(h))
		shader.Model.SetMVPMatrix(pMatrix, mvMatrix)

		// Clear screen, then Draw everything
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		model.Render()

		// Swaps the buffer that was drawn on to be visible. The visible buffer becomes the one that gets drawn on until it's swapped again.
		view.Window.SwapBuffers()
		<-ticker.C // wait up to the framerate cap.
	}
}

package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/omustardo/demos/sdl-input-and-display/keyboard"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	gametick                = time.Second / 3
	framerate               = 60
	winTitle                = "Go-SDL2 Render"
	winWidth, winHeight int = 640, 360
)

func init() {
	// OpenGl needs to run on one thread, evidently.
	// https://github.com/go-gl/gl/issues/13
	runtime.LockOSThread()
}

func main() {
	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_OPENGL|sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalf("Failed to create window: %s\n", err)
		return
	}
	defer window.Destroy()

	context, err := sdl.GL_CreateContext(window)
	if err != nil {
		log.Fatalln(err)
	}
	defer sdl.GL_DeleteContext(context)
	// 0 == immediate updates, 1 == VSYNC, -1 == late swap tearing
	if err = sdl.GL_SetSwapInterval(1); err != nil {
		log.Fatalln("Error setting swap interval (vsync):", err)
	}
	if err = sdl.GL_SetAttribute(sdl.GL_DOUBLEBUFFER, 1); err != nil {
		log.Fatalln("Error turning on GL double buffering:", err)
	}

	if err := gl.Init(); err != nil {
		log.Fatalln("Error initializing opengl:", err)
	}
	projectionMode()
	setViewport(winWidth, winHeight)

	modelViewMode()
	setModelViewOptions()

	keyboardHandler := keyboard.NewHandler()

	// Game State
	rect := sdl.Rect{X: 0, Y: 0, W: 100, H: 100}

	running := true
	ticker := time.NewTicker(time.Second / framerate)
	fmt.Println("Framerate Capped at:", time.Duration(time.Second/framerate), " per frame")
	fmt.Println("Game tick rate:", gametick)
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				log.Println("Got a QuitEvent.")
				running = false
				break
			}
		}
		// TODO: Decouple input handling and framerate.

		// Read input
		keyboardHandler.Update() // Note: This only works because sdl.PollEvent is called above until all events are processed.
		//fmt.Println(keyboardHandler.String() + "\n---")
		if keyboardHandler.LeftPressed() && !keyboardHandler.WasLeftPressed() {
			rect.X -= 100
		}
		if keyboardHandler.RightPressed() && !keyboardHandler.WasRightPressed() {
			rect.X += 100
		}
		if keyboardHandler.DownPressed() && !keyboardHandler.WasDownPressed() {
			rect.Y -= 100
		}
		if keyboardHandler.UpPressed() && !keyboardHandler.WasUpPressed() {
			rect.Y += 100
		}

		// TODO: I think these only need to be done once at the very start? Does it hurt to do them every time?

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.Enable(gl.BLEND)
		gl.Enable(gl.POINT_SMOOTH)
		gl.Enable(gl.LINE_SMOOTH)
		gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
		gl.LoadIdentity()

		RectFilled(float32(rect.X), float32(rect.Y), float32(rect.X+rect.W), float32(rect.Y-rect.H), 1, 0.5, 0.2, 1)

		sdl.GL_SwapWindow(window)

		<-ticker.C // wait based on framerate
	}
}

func projectionMode() {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
}

func setViewport(width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Ortho(-float64(width)/2, float64(width)/2, -float64(height)/2, float64(height)/2, -1.0, 1.0)
}

func modelViewMode() {
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

func setModelViewOptions() {
	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.FRONT) // ??
	// gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA) // ??
	gl.Enable(gl.BLEND)
	gl.Enable(gl.ALPHA_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.Disable(gl.DEPTH_TEST)
}

func RectFilled(x1, y1, x2, y2, r, g, b, a float32) {
	gl.Begin(gl.TRIANGLES)
	gl.Color4f(r, g, b, a)

	gl.Vertex3f(x1, y1, 0)
	gl.Vertex3f(x2, y1, 0)
	gl.Vertex3f(x2, y2, 0)

	gl.Vertex3f(x1, y1, 0)
	gl.Vertex3f(x2, y2, 0)
	gl.Vertex3f(x1, y2, 0)
	gl.End()
}

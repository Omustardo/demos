package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"path/filepath"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"github.com/omustardo/demos/opengl/camera"
	"github.com/omustardo/demos/opengl/fps"
	"github.com/omustardo/demos/opengl/keyboard"
	"github.com/omustardo/demos/opengl/mouse"
	"github.com/omustardo/demos/opengl/shader"
	"github.com/omustardo/demos/opengl/shape"
	"github.com/omustardo/demos/opengl/util"
)

var (
	windowWidth  = flag.Int("window_width", 1000, "initial window width")
	windowHeight = flag.Int("window_height", 1000, "initial window height")
	WindowSize   [2]int // Up to date window size.
)

const (
	gametick  = time.Second / 3
	framerate = time.Second / 60

	// Screenshots are saved in the target folder. Their name is the millisecond timestamp when they are taken.
	screenshotPath = `C:\Users\Omar\Desktop\screenshots\`
)

func init() {
	log.SetFlags(log.Lshortfile) // log print with .go file and line number.
	log.SetOutput(os.Stdout)

	// To log to multiple locations:
	//var b bytes.Buffer
	//bufWriter := bufio.NewWriter(&b)
	//log.SetOutput(io.MultiWriter(os.Stdout, os.Stderr, bufWriter))
}

func main() {
	// TODO: Loading screen.
	WindowSize[0], WindowSize[1] = *windowWidth, *windowHeight
	windowWidth, windowHeight = nil, nil // Clear the flags. They're only for initialization and shouldn't be used elsewhere.

	err := glfw.Init(gl.ContextWatcher)
	if err != nil {
		log.Fatal(err)
	}
	defer glfw.Terminate()
	glfw.WindowHint(glfw.Samples, 16) // Anti-aliasing.

	// Window hints to require OpenGL 3.2 or above, and to disable deprecated functions. https://open.gl/context#GLFW
	// These hints are not supported since we're using goxjs/glfw rather than the regular glfw, but should be used in a
	// standard desktop glfw project. TODO: Add support for these in goxjs/glfw/hint_glfw.go or consider using a conditional build rule.
	//glfw.WindowHint(glfw.ContextVersionMajor, 3)
	//glfw.WindowHint(glfw.ContextVersionMinor, 2)
	//glfw.WindowHint(glfw.OpenGLProfile, glfw.OPENGL_CORE_PROFILE)
	//glfw.WindowHint(glfw.OpenGLForwardCompatible, gl.TRUE)

	// Note CreateWindow ignores input size for WebGL/HTML canvas - it expands to fill browser window.
	// It still matters for desktop.
	window, err := glfw.CreateWindow(WindowSize[0], WindowSize[1], "Graphics Demo", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	window.MakeContextCurrent()
	fmt.Printf("OpenGL: %s %s %s; %v samples.\n", gl.GetString(gl.VENDOR), gl.GetString(gl.RENDERER), gl.GetString(gl.VERSION), gl.GetInteger(gl.SAMPLES))
	fmt.Printf("GLSL: %s.\n", gl.GetString(gl.SHADING_LANGUAGE_VERSION))

	glfw.SwapInterval(1) // Vsync.

	gl.ClearColor(0, 0, 0, 1) // Background Color
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.CULL_FACE) // NOTE: If triangles appear to be missing, this is probably the cause. The order that vertices are listed matters.
	//gl.Enable(gl.DEPTH_TEST) // TODO: Enable once everything uses 3D meshes. For now just depend on draw order.
	//gl.DepthFunc(gl.LESS) // Accept fragment if it's closer to the camera than the former one

	shape.LoadModels()

	// Set up a callback for when the window is resized. Call it once for good measure.
	framebufferSizeCallback := func(w *glfw.Window, framebufferSize0, framebufferSize1 int) {
		gl.Viewport(0, 0, framebufferSize0, framebufferSize1)
		WindowSize[0], WindowSize[1] = w.GetSize()
	}
	{
		framebufferSizeX, framebufferSizeY := window.GetFramebufferSize()
		framebufferSizeCallback(window, framebufferSizeX, framebufferSizeY)
	}
	window.SetFramebufferSizeCallback(framebufferSizeCallback)
	// TODO: Support fullscreen

	// Init shaders.
	if err := shader.Initialize(); err != nil {
		log.Fatal(err)
	}

	if err := gl.GetError(); err != 0 {
		log.Fatalf("gl error: %v", err)
	}

	mouseHandler, mouseButtonCallback, cursorPositionCallback := mouse.NewHandler()
	window.SetMouseButtonCallback(mouseButtonCallback)
	window.SetCursorPosCallback(cursorPositionCallback)
	keyboardHandler, keyboardCallback := keyboard.NewHandler()
	window.SetKeyCallback(keyboardCallback)
	// TODO: window.SetScrollCallback()
	// TODO: gestures / touchpad support

	fpsCounter := fps.NewFPSCounter()

	player := &shape.Rect{
		X: 0, Y: 0,
		Width:  100,
		Height: 100,
		R:      0.8, G: 0.1, B: 0.3, A: 1,
		Angle: 0,
	}
	cam := camera.NewTargetCamera(player)

	miscCircles := []*shape.Circle{
		{
			P:      mgl32.Vec3{100, 200, 0},
			Radius: 20,
			R:      0.2, G: 0.7, B: 0.5, A: 1,
		},
		{
			P:      mgl32.Vec3{-200, -100, 0},
			Radius: 15,
			R:      0.4, G: 0.9, B: 0.1, A: 1,
		},
		{
			P:      mgl32.Vec3{0, 50, 0},
			Radius: 35,
			R:      1, G: 0.5, B: 0.2, A: 1,
		},
	}

	// Generate parallax rectangles.
	parallaxObjects := shape.GenParallaxRects(cam, 500, 8, 5, 0.1, 0.2)                                // Near
	parallaxObjects = append(parallaxObjects, shape.GenParallaxRects(cam, 300, 5, 3.5, 0.35, 0.5)...)  // Med
	parallaxObjects = append(parallaxObjects, shape.GenParallaxRects(cam, 200, 2, 0.5, 0.75, 0.85)...) // Far
	parallaxObjects = append(parallaxObjects, shape.GenParallaxRects(cam, 100, 1, 0.1, 0.9, 0.95)...)  // Distant
	// Put the parallax info in buffers on the GPU. TODO: Consider using a single interleaved buffer. Stride and offset are annoying though.
	parallaxPositionBuffer, parallaxTranslationBuffer, parallaxTranslationRatioBuffer, parallaxAngleBuffer, parallaxScaleBuffer, parallaxColorBuffer := shape.GetParallaxBuffers(parallaxObjects)

	ticker := time.NewTicker(framerate)
	gameTicker := time.NewTicker(gametick)
	debugLogTicker := time.NewTicker(time.Second)
	for !window.ShouldClose() {
		fpsCounter.Update()

		// Handle Input
		keyboardHandler.Update()
		mouseHandler.Update()
		ApplyInputs(keyboardHandler, mouseHandler, player)

		// Run game logic
		select {
		case _, ok := <-gameTicker.C: // do stuff with game logic on ticks to minimize expensive calculations.
			if ok {
			}
		default:
		}
		cam.Update()

		// Set up Model-View-Projection Matrix and send it to the shader program.
		mvMatrix := cam.ModelView()
		pMatrix := mgl32.Ortho(-float32(WindowSize[0])/2, float32(WindowSize[0])/2,
			-float32(WindowSize[1])/2, float32(WindowSize[1])/2,
			cam.Near(), cam.Far())
		shader.Basic.SetMVPMatrix(pMatrix, mvMatrix)
		shader.Parallax.SetMVPMatrix(pMatrix, mvMatrix)

		// Clear screen, then Draw everything
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT) // TODO: Some cool graphical effects result from not clearing the screen.
		shape.DrawXYZAxes()
		for _, c := range miscCircles {
			c.Draw()
		}

		// Draw parallax objects
		// Old inefficient way of drawing the rectangles one by one:
		//for _, r := range parallaxObjects {
		//	r.DrawFilled()
		//}
		// New batched method:
		shape.DrawParallaxBuffers(6*len(parallaxObjects) /* vertices in total */, cam.Position().Vec2(),
			parallaxPositionBuffer, parallaxTranslationBuffer, parallaxTranslationRatioBuffer,
			parallaxAngleBuffer, parallaxScaleBuffer, parallaxColorBuffer)

		player.Draw()

		window.SwapBuffers() // Swaps the buffer that was drawn on to be visible. The visible buffer becomes the one that gets drawn on until it's swapped again.
		glfw.PollEvents()    // Reads window events, like input (?)

		// Debug logging - limited to once every X seconds to avoid spam.
		select {
		case _, ok := <-debugLogTicker.C:
			if ok {
				fmt.Println("location:", cam.Position())
				if mouseHandler.LeftPressed() {
					fmt.Println("detected mouse press at", mouseHandler.Position)
				}
				fmt.Println(fpsCounter.GetFPS(), "fps")
			}
		default:
		}
		<-ticker.C // wait up to 1/60th of a second. This caps framerate to 60 FPS.
	}
}

func ApplyInputs(keyboardHandler *keyboard.Handler, mouseHandler *mouse.Handler, player shape.Shape) {
	var move mgl32.Vec2
	if keyboardHandler.IsKeyDown(glfw.KeyA) || keyboardHandler.LeftPressed() {
		move[0] = -1
	}
	if keyboardHandler.IsKeyDown(glfw.KeyD) || keyboardHandler.RightPressed() {
		move[0] = 1
	}
	if keyboardHandler.IsKeyDown(glfw.KeyW) || keyboardHandler.UpPressed() {
		move[1] = 1
	}
	if keyboardHandler.IsKeyDown(glfw.KeyS) || keyboardHandler.DownPressed() {
		move[1] = -1
	}
	// Without this check, Normalize() could result in [NaN, NaN]
	if move.Len() != 0 {
		move = move.Normalize().Mul(10)
	}
	player.ModifyCenter(move[0], move[1])

	if keyboardHandler.IsKeyDown(glfw.KeySpace) && !keyboardHandler.WasKeyDown(glfw.KeySpace) {
		util.SaveScreenshot(WindowSize[0], WindowSize[1], filepath.Join(screenshotPath, fmt.Sprintf("%d.png", util.GetTimeMillis())))
	}

	if mouseHandler.LeftPressed() {
	}
	if mouseHandler.RightPressed() {
	}
}

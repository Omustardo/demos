package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"github.com/omustardo/demos/opengl/camera"
	"github.com/omustardo/demos/opengl/fps"
	"github.com/omustardo/demos/opengl/keyboard"
	"github.com/omustardo/demos/opengl/mouse"
	"github.com/omustardo/demos/opengl/shader"
	"github.com/omustardo/demos/opengl/shape"
)

var (
	windowWidth  = flag.Int("window_width", 1000, "initial window width")
	windowHeight = flag.Int("window_height", 1000, "initial window height")
	WindowSize   [2]int // Up to date window size.
)

const (
	gametick  = time.Second / 3
	framerate = time.Second / 60
)

func main() {
	err := glfw.Init(gl.ContextWatcher)
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()
	glfw.WindowHint(glfw.Samples, 16) // Anti-aliasing.

	// Note CreateWindow ignores input size for WebGL/HTML canvas - it expands to fill browser window.
	// It still matters for desktop.
	window, err := glfw.CreateWindow(*windowWidth, *windowHeight, "Graphics Demo", nil, nil)
	if err != nil {
		panic(err)
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
	//gl.DepthFunc(gl.LESS) // Accept fragment if it closer to the camera than the former one

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

	// Init shaders.
	if err := shader.SetupProgram(); err != nil {
		panic(err)
	}

	if err := gl.GetError(); err != 0 {
		fmt.Printf("gl error: %v", err)
		return
	}

	mouseHandler, mouseButtonCallback, cursorPositionCallback := mouse.NewHandler()
	window.SetMouseButtonCallback(mouseButtonCallback)
	window.SetCursorPosCallback(cursorPositionCallback)
	keyboardHandler, keyboardCallback := keyboard.NewHandler()
	window.SetKeyCallback(keyboardCallback)
	// TODO: window.SetScrollCallback()

	fpsCounter := fps.NewFPSCounter()

	player := &shape.Rect{
		X: 0, Y: 0,
		Width:  100,
		Height: 100,
		R:      0.8, G: 0.1, B: 0.3, A: 1,
		Angle: float32(math.Pi / 2),
	}
	cam := camera.CameraI(camera.NewTargetCamera(player)) // TODO: Consider having NewCamera functions return CameraI's

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

	genParallaxRects := func(count int, minWidth, maxWidth, minSpeedRatio, maxSpeedRatio float32) []shape.Shape {
		shapes := make([]shape.Shape, count)
		for i := 0; i < count; i++ {
			shapes[i] = &shape.ParallaxRect{
				Rect: shape.Rect{
					X: rand.Float32()*2000 - 1000, Y: rand.Float32()*2000 - 1000, // Note not even distribution - they are drawn from bottom left corner so everything is Up and Right shifted slightly
					R: rand.Float32(), G: rand.Float32(), B: rand.Float32(), A: 1,
					Width:  rand.Float32()*(maxWidth-minWidth) + minWidth,
					Height: rand.Float32()*(maxWidth-minWidth) + minWidth,
					Angle:  rand.Float32() * 360,
				},
				Camera:        cam, // TODO: Changing the camera should be allowed, but right now it breaks this.
				LocationRatio: rand.Float32()*(maxSpeedRatio-minSpeedRatio) + minSpeedRatio,
			}
		}
		return shapes
	}

	parallaxObjects := genParallaxRects(500, 8, 5, 0.1, 0.2)                                // Near
	parallaxObjects = append(parallaxObjects, genParallaxRects(300, 5, 3.5, 0.35, 0.5)...)  // Med
	parallaxObjects = append(parallaxObjects, genParallaxRects(100, 2, 0.5, 0.75, 0.85)...) // Far
	parallaxObjects = append(parallaxObjects, genParallaxRects(50, 1, 0.1, 0.9, 0.95)...)   // Distant

	ticker := time.NewTicker(framerate)
	gameTicker := time.NewTicker(gametick)
	debugLogTicker := time.NewTicker(time.Second * 2)
	for !window.ShouldClose() {
		fpsCounter.Update()

		// Handle Input
		keyboardHandler.Update()
		mouseHandler.Update()
		ApplyInputs(keyboardHandler, mouseHandler, player, cam)

		// Run game logic
		select {
		case _, ok := <-gameTicker.C: // do stuff with game logic on ticks to minimize expensive calculations.
			if ok {
				fmt.Println(fpsCounter.GetFPS(), "fps")
			}
		default:
		}
		cam.Update()

		// player.Angle = float32(time.Now().Nanosecond() / 1000 % 180) // For testing rotation

		// Set up Model-View-Projection Matrix and send it to the shader program.
		mvMatrix := cam.ModelView()
		pMatrix := mgl32.Ortho(-float32(WindowSize[0])/2, float32(WindowSize[0])/2,
			-float32(WindowSize[1])/2, float32(WindowSize[1])/2,
			cam.Near(), cam.Far())
		shader.SetMVPMatrix(pMatrix, mvMatrix)

		// Clear screen, then Draw everything
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		shape.DrawXYZAxes()
		for _, c := range miscCircles {
			c.Draw()
		}
		for _, obj := range parallaxObjects {
			obj.DrawFilled()
		}
		player.Draw()

		window.SwapBuffers()
		glfw.PollEvents()

		// Debug logging - limited to once every X seconds to avoid spam.
		select {
		case _, ok := <-debugLogTicker.C:
			if ok {
				fmt.Println("location:", cam.Position())
			}
		default:
		}
		<-ticker.C // wait up to 1/60th of a second. This caps framerate to 60 FPS.
	}
}

func ApplyInputs(keyboardHandler *keyboard.Handler, mouseHandler *mouse.Handler, player shape.Shape, cam camera.CameraI) {
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

	if mouseHandler.LeftPressed() {
	}
	if mouseHandler.RightPressed() {
	}
}

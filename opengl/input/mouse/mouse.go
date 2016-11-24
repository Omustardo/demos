// mouse handles mouse interaction with a glfw window.
package mouse

// TODO: How well does this handle unusual events? Try unplugging mouse. Using multiple mice.
// TODO: Make fields "read-only" by making them private and providing getters.

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/goxjs/glfw"
)

// Handler is the singleton mouse handler. It should be initialized with mouse.Initialize(), and then
// all mouse related input should be obtained though it.
var Handler *handler

func Initialize(window *glfw.Window) {
	if window == nil {
		panic("window is nil")
	}
	var (
		mouseButtonCallback    glfw.MouseButtonCallback
		cursorPositionCallback glfw.CursorPosCallback
		scrollCallback         glfw.ScrollCallback
	)
	Handler, mouseButtonCallback, cursorPositionCallback, scrollCallback = newHandler()
	window.SetMouseButtonCallback(mouseButtonCallback)
	window.SetCursorPosCallback(cursorPositionCallback)
	window.SetScrollCallback(scrollCallback)
}

type handler struct {
	// State maps from buttons to whether they are pressed.
	State, PreviousState map[glfw.MouseButton]bool

	// position is the screen coordinate where the mouse pointer is.
	// (0,0) is the top left of the drawable region (i.e. not including the title bar in a desktop environment).
	// Down and right are positive. Up and left are negative.
	position, previousPosition mgl32.Vec2

	// Scroll holds how much scrolling has occurred since the start of the program.
	// PreviousScroll is how much scrolling occurred since the start of the program, ignoring the most recent frame.
	// To determine changes, subtract the two.
	// The Y value is the standard forward/back, while the left/right scrolling available on some mice is in the X value.
	// While glfw says the value is a float, I've only seen it as integers. One "tick" is +1 or -1 depending on direction.
	// Positive for Forward/Left. Negative for Back/Right. 0 by default.
	Scroll, PreviousScroll mgl32.Vec2
}

func newHandler() (*handler, glfw.MouseButtonCallback, glfw.CursorPosCallback, glfw.ScrollCallback) {
	h := &handler{
		State:         make(map[glfw.MouseButton]bool),
		PreviousState: make(map[glfw.MouseButton]bool),
	}
	return h, h.mouseButtonCallback, h.cursorPosCallback, h.scrollCallback
}

// mouseButtonCallback is a function for glfw to call when a button event occurs.
func (h *handler) mouseButtonCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	// Note that this will overwrite any unhandled actions.
	// For example, if you press the left mouse button and then release it without calling
	// handler.Update() in between, it will appear as if no action was taken.
	// I think this is fine, since you shouldn't be able to click and release within 1/60th of a second.
	// If there are noticeable missing keypresses, then this is almost certainly the problem.
	h.setState(button, action)
}

// cursorPosCallback is a function for glfw to call when a button event occurs.
func (h *handler) cursorPosCallback(window *glfw.Window, xpos, ypos float64) {
	// log.Println("got cursor pos event:", xpos, ypos)
	h.position[0] = float32(xpos)
	h.position[1] = float32(ypos)
}

// scrollCallback is a function for glfw to call when a scroll wheel event occurs.
func (h *handler) scrollCallback(window *glfw.Window, xoff, yoff float64) {
	// log.Println("got scroll event:", xoff, yoff)
	h.Scroll[0] += float32(xoff)
	h.Scroll[1] += float32(yoff)
}

func (h *handler) setState(button glfw.MouseButton, action glfw.Action) {
	switch action {
	case glfw.Press:
		h.State[button] = true
	case glfw.Release:
		h.State[button] = false
	}
}

// Update is expected to be called once per frame, or more.
func (h *handler) Update() {
	h.PreviousState = h.State
	h.State = make(map[glfw.MouseButton]bool) // TODO: making a new map every frame isn't good for garbage collection.
	for b, pressed := range h.PreviousState {
		if pressed {
			h.State[b] = true
		}
	}
	h.previousPosition = h.position
	h.PreviousScroll = h.Scroll
}

func (h *handler) LeftPressed() bool {
	return h.State[glfw.MouseButtonLeft]
}
func (h *handler) RightPressed() bool {
	return h.State[glfw.MouseButtonRight]
}
func (h *handler) WasLeftPressed() bool {
	return h.PreviousState[glfw.MouseButtonLeft]
}
func (h *handler) WasRightPressed() bool {
	return h.PreviousState[glfw.MouseButtonRight]
}

// Position returns the screen coordinate where the mouse pointer is.
// (0,0) is the top left of the drawable region (i.e. not including the title bar in a desktop environment).
// Down and right are positive. Up and left are negative.
func (h *handler) Position() mgl32.Vec2 {
	return h.position
}
func (h *handler) PreviousPosition() mgl32.Vec2 {
	return h.previousPosition
}

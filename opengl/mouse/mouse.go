// mouse handles mouse interaction with a glfw window.
package mouse

// TODO: How well does this handle unusual events? Try unplugging mouse. Using multiple mice.
// TODO: Make fields "read-only" by making them private and providing getters.

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/goxjs/glfw"
)

type Handler struct {
	// State maps from buttons to whether they are pressed.
	State, PreviousState map[glfw.MouseButton]bool

	// Position is the space coordinate where the mouse pointer is.
	// (0,0) is the top left of the drawable region (i.e. not including the title bar in a desktop environment).
	// Down and right are positive. Up and left are negative.
	Position, PreviousPosition mgl32.Vec2

	// Scroll holds how much scrolling has occurred since the start of the program.
	// PreviousScroll is how much scrolling occurred since the start of the program, ignoring the most recent frame.
	// To determine changes, subtract the two.
	// The Y value is the standard forward/back, while the left/right scrolling available on some mice is in the X value.
	// While glfw says the value is a float, I've only seen it as integers. One "tick" is +1 or -1 depending on direction.
	// Positive for Forward/Left. Negative for Back/Right. 0 by default.
	Scroll, PreviousScroll mgl32.Vec2
}

func NewHandler() (*Handler, glfw.MouseButtonCallback, glfw.CursorPosCallback, glfw.ScrollCallback) {
	h := &Handler{
		State:         make(map[glfw.MouseButton]bool),
		PreviousState: make(map[glfw.MouseButton]bool),
	}
	return h, h.MouseButtonCallback, h.CursorPosCallback, h.ScrollCallback
}

// MouseButtonCallback is a function for glfw to call when a button event occurs.
func (h *Handler) MouseButtonCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	// Note that this will overwrite any unhandled actions.
	// For example, if you press the left mouse button and then release it without calling
	// handler.Update() in between, it will appear as if no action was taken.
	// I think this is fine, since you shouldn't be able to click and release within 1/60th of a second.
	// If there are noticeable missing keypresses, then this is almost certainly the problem.
	h.setState(button, action)
}

// CursorPosCallback is a function for glfw to call when a button event occurs.
func (h *Handler) CursorPosCallback(window *glfw.Window, xpos, ypos float64) {
	// log.Println("got cursor pos event:", xpos, ypos)
	h.Position[0] = float32(xpos)
	h.Position[1] = float32(ypos)
}

func (h *Handler) ScrollCallback(window *glfw.Window, xoff, yoff float64) {
	// log.Println("got scroll event:", xoff, yoff)
	h.Scroll[0] += float32(xoff)
	h.Scroll[1] += float32(yoff)
}

func (h *Handler) setState(button glfw.MouseButton, action glfw.Action) {
	switch action {
	case glfw.Press:
		h.State[button] = true
	case glfw.Release:
		h.State[button] = false
	}
}

// Update is expected to be called roughly once per frame. A likely choice is
// whenever a physics step occurs.
func (h *Handler) Update() {
	h.PreviousState = h.State
	h.State = make(map[glfw.MouseButton]bool) // TODO: making a new map every frame isn't good for garbage collection.
	for b, pressed := range h.PreviousState {
		if pressed {
			h.State[b] = true
		}
	}
	h.PreviousPosition = h.Position
	h.PreviousScroll = h.Scroll
}

func (h *Handler) LeftPressed() bool {
	return h.State[glfw.MouseButtonLeft]
}
func (h *Handler) RightPressed() bool {
	return h.State[glfw.MouseButtonRight]
}
func (h *Handler) WasLeftPressed() bool {
	return h.PreviousState[glfw.MouseButtonLeft]
}
func (h *Handler) WasRightPressed() bool {
	return h.PreviousState[glfw.MouseButtonRight]
}

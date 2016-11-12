package camera

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera interface {
	ModelView() mgl32.Mat4
	Near() float32
	Far() float32
	Update()
	Position() mgl32.Vec3
}

// TODO: DirectionalCamera that follows player orientation (up on the screen is always the direction the player faces).
// TODO: TrailingCamera. (in progress) Based on player position, but has its own max speed and delay, so stays behind player a bit.
// TODO: ZoomCamera. Allows changing of the mgl32.LookAt so you can zoom in/out with limits.

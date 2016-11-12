package camera

import "github.com/go-gl/mathgl/mgl32"

// Compile time check that FreeCamera implements CameraI
var _ Camera = (*FreeCamera)(nil)

// FreeCamera is a camera that is not attached to any player. It can be scrolled around the level by modifying the Pos.
type FreeCamera struct {
	Pos [3]float32
}

func (c *FreeCamera) ModelView() mgl32.Mat4 {
	return mgl32.LookAt(
		c.Pos[0], c.Pos[1], c.Pos[2], // Camera Position
		c.Pos[0], c.Pos[1], -1, // Target Position. Looking down on Z.
		0, 1, 0) // Up vector
}

func (c *FreeCamera) Update() {
	// Enforce bounds. This camera is always looking down onto the XY plane.
	if c.Pos[2] <= 0 {
		c.Pos[2] = 0
	}
}

func (c *FreeCamera) Near() float32 {
	return 0.1
}

func (c *FreeCamera) Far() float32 {
	return 100
}

func (c *FreeCamera) Position() mgl32.Vec3 {
	return c.Pos
}

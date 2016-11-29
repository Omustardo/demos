package camera

import "github.com/go-gl/mathgl/mgl32"

// Camera is a perspective camera that looks in the negative Z, with positive Y being up.
type Camera struct {
	Pos       mgl32.Vec3
	Near, Far float32
}

func (c *Camera) ModelView() mgl32.Mat4 {
	return mgl32.LookAt(
		c.Pos[0], c.Pos[1], c.Pos[2], // Camera Position
		c.Pos[0], c.Pos[1], -1, // Target Position. Looking down on Z.
		0, 1, 0) // Up vector
}

func (c *Camera) Projection(width, height float32) mgl32.Mat4 {
	return mgl32.Perspective(45, float32(width)/float32(height), c.Near, c.Far)
}

func (c *Camera) Position() mgl32.Vec3 {
	return c.Pos
}

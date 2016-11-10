package camera

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/omustardo/demos/opengl/entity"
)

var _ CameraI = (*TargetCamera)(nil)

// TargetCamera is a camera that is always locked to an entity.
type TargetCamera struct {
	target entity.Entity
	Pos    mgl32.Vec3 // This should always reflect the target's position, but always with a Z value > 0
}

func NewTargetCamera(target entity.Entity) *TargetCamera {
	p := &TargetCamera{
		target: target,
	}
	p.Update()
	p.Pos[2] = 1 // Not necessary with the bounds enforcing in Update(), but nice to be safe.
	return p
}

func (c *TargetCamera) ModelView() mgl32.Mat4 {
	//mat := mgl32.Translate3D(float32(draw.WindowSize[0]), float32(draw.WindowSize[1]), 0)
	//mat = mat.Mul4(mgl32.HomogRotate3DZ(c.Pos[2]))
	//mat = mat.Mul4(mgl32.Translate3D(-c.Pos[0], -c.Pos[1], -c.Pos[2])) // TODO: use entity position
	return mgl32.LookAt(
		c.Pos[0], c.Pos[1], c.Pos[2], // Camera Position
		c.Pos[0], c.Pos[1], 0, // Target Position.  Looking down on Z.
		0, 1, 0) // Up vector // TODO: For a camera that has Screen Up always as the direction the entity is facing, I think we just need to modify this line.
}

func (c *TargetCamera) Update() {
	targetCenter := c.target.Position()
	c.Pos[0], c.Pos[1] = targetCenter.X(), targetCenter.Y()
	// Enforce bounds. This camera is always looking down onto the XY plane.
	if c.Pos[2] <= c.Near() {
		c.Pos[2] = c.Near() + 0.05
	}
}

func (c *TargetCamera) Near() float32 {
	return 0.1
}

func (c *TargetCamera) Far() float32 {
	return 100
}

func (c *TargetCamera) Position() mgl32.Vec3 {
	return c.Pos
}

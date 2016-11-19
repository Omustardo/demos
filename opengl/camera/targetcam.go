package camera

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/omustardo/demos/opengl/entity"
)

var _ Camera = (*TargetCamera)(nil)

// TargetCamera is an orthographic camera that is always locked to an entity.
type TargetCamera struct {
	target entity.Entity
	Pos    mgl32.Vec3 // This should always reflect the target's position, but always with a Z value > 0
}

func NewTargetCamera(target entity.Entity) Camera {
	p := &TargetCamera{
		target: target,
	}
	p.Update()
	p.Pos[2] = 1 // Not necessary with the bounds enforcing in Update(), but nice to be safe.
	return p
}

func (c *TargetCamera) ModelView() mgl32.Mat4 {
	targetPos := c.target.Position()
	return mgl32.LookAt(
		targetPos[0], targetPos[1], 1, // Camera Position. Always above target.
		targetPos[0], targetPos[1], 0, // Target Position.
		0, 1, 0) // Up vector // TODO: For a camera that has Screen Up always as the direction the entity is facing, I think we just need to modify this line.
}

func (c *TargetCamera) Projection(width, height float32) mgl32.Mat4 {
	return mgl32.Ortho(-width/2, width/2,
		-height/2, height/2,
		c.Near(), c.Far())
}

func (c *TargetCamera) Update() {}

func (c *TargetCamera) Near() float32 {
	return 0.1
}

func (c *TargetCamera) Far() float32 {
	return 100
}

func (c *TargetCamera) Position() mgl32.Vec3 {
	pos := c.target.Position()
	pos[2] = 1
	return pos
}

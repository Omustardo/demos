package camera

import (
	"log"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/omustardo/demos/modelviewer/entity"
)

// Camera is a camera that is always positioned at an offset from the target entity. Zoomer can modify the length
// of the offset. The camera always looks toward a target with the provided Up vector being up.
type Camera struct {
	Target entity.Target
	// TargetOffset determines where the camera is positioned in relation to the target.
	TargetOffset mgl32.Vec3
	// Up is a vector pointing in the same direction as the top of the screen.
	Up mgl32.Vec3
	// Near and Far are the range to render entities in front of the camera. Be sure to make them small and large enough
	// to compensate for Zoomer. For example, if your object is 100 units away and your Zoomer can zoom in to 300%,
	// and zoom out to 25%, then the you must set Near<33.3 and Far>400 in order to always keep the target in view.
	Near, Far float32
	// Field of view in radians. This only matters if using a perspective projection and can be ignored if using an orthographic projection.
	FOV float32
}

func (c *Camera) ModelView() mgl32.Mat4 {
	return mgl32.LookAtV(c.Position(), c.Target.Center(), c.Up)
}

func (c *Camera) ProjectionOrthographic(width, height float32) mgl32.Mat4 {
	return mgl32.Ortho(-width/2, width/2,
		-height/2, height/2,
		c.Near, c.Far)
}

func (c *Camera) ProjectionPerspective(width, height float32) mgl32.Mat4 {
	return mgl32.Perspective(c.FOV, float32(width)/float32(height), c.Near, c.Far)
}

func (c *Camera) Update() {
	if c.Near >= c.Far {
		log.Println("camera's near is closer than far - nothing will render")
	}
}

func (c *Camera) Position() mgl32.Vec3 {
	return c.Target.Center().Add(c.TargetOffset)
}

// ScreenToWorldCoord2D returns the world coordinates of a point on the screen.
// This depends on the camera always looking directly down onto the XY-plane. e.g. camera position has positive Z.
// The screen space coordinate is expected in the coordinate system where the top left corner is (0,0), Y increases down, and X increases right.
func (c *Camera) ScreenToWorldCoord2D(screenPoint mgl32.Vec2, windowWidth, windowHeight int) mgl32.Vec2 {
	pos := c.Position()
	return mgl32.Vec2{
		pos.X() + (screenPoint.X() - float32(windowWidth)/2),
		pos.Y() - (screenPoint.Y() - float32(windowHeight)/2),
	}
}

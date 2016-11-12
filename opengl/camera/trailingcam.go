package camera

/* TODO: Implement this. Something like: https://docs.unity3d.com/ScriptReference/Vector3.SmoothDamp.html

import (
	"log"
	"github.com/omustardo/demos/opengl/entity"
)


var _ Camera = (*TrailingCamera)(nil)

// Trailing Camera follows a target with has a short delay.
type TrailingCamera struct {
	TargetCamera

	MaxTrailingDistance float32
}

func NewTrailingCamera(target entity.Entity) Camera {
	c := &TrailingCamera{
		TargetCamera: TargetCamera{
			target: target,
		},
		MaxTrailingDistance: 200,
	}
	c.Pos = c.target.Position()
	c.Update()
	c.Pos[2] = 1 // Not necessary with the bounds enforcing in Update(), but nice to be safe.
	return c
}

func (c *TrailingCamera) Update() {
	defer func() {
		// Enforce bounds. This camera is always looking down onto the XY plane.
		if c.Pos[2] <= c.Near() {
			c.Pos[2] = c.Near() + 0.05
		}
	}()

	targetCenter := c.TargetCamera.target.Position().Vec2()
	if targetCenter[0] == c.Pos[0] && targetCenter[1] == c.Pos[1] {
		return
	}

	// Vector pointing from target to camera
	v := c.Position().Vec2().Sub(targetCenter)
	trailingDistance := v.Len()
	if trailingDistance == 0 {
		log.Println("trailing==0")
		return
	}
	v = v.Normalize()

	// Reduce distance between camera and target:
	//   by at least 1% per frame
	//   if under 50% of MaxTrailingDistance do 2%
	//   under 20% do 5%
	//   under 5% just snap to 0
	// If trailing distance is still > MaxTrailingDistance, snap to max dist.
	switch {
	case trailingDistance >= c.MaxTrailingDistance:
		trailingDistance = c.MaxTrailingDistance
	//case trailingDistance >= 0.05 * c.MaxTrailingDistance:
	//
	//case trailingDistance >= 0.5 * c.MaxTrailingDistance:
	//	trailingDistance *= 0.99
	//case trailingDistance >= 0.2 * c.MaxTrailingDistance:
	//	trailingDistance *= 0.98
	//case trailingDistance >= 0.05 * c.MaxTrailingDistance:
	//	trailingDistance *= 0.95
	default:
		log.Println("Snapping")
		trailingDistance *= 1-(trailingDistance/c.MaxTrailingDistance)
	}

	c.Pos[0], c.Pos[1] = targetCenter[0] + v[0] * trailingDistance, targetCenter[1] + v[1] * trailingDistance
}

*/

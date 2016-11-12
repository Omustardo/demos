package shape

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/goxjs/gl"
	"github.com/omustardo/demos/opengl/camera"
	"github.com/omustardo/demos/opengl/parallax"
	"github.com/omustardo/demos/opengl/shader"
)

type Shape interface {
	// Draw draws an outline of a Shape using line segments.
	Draw()
	// DrawFilled draws a filled in Shape using triangles.
	DrawFilled()
	// SetCenter sets the Shape to the specified position.
	SetCenter(x, y float32)
	// ModifyCenter moves the Shape by the specified amount.
	ModifyCenter(x, y float32)
	// Center is a point about which all actions, like rotation, are defined.
	Center() mgl32.Vec2
	// TODO: Get rid of this in favor of Center? and make Center a vec3 for future proofing?
	// Consider the ability to choose a center point for rotating around an arbitrary point.
	Position() mgl32.Vec3
}

// Loads models into buffers on the GPU. Must be called after glfw.Init() @@@ Is this true? What is it dependent on? The gl.CreateBuffer() definitely doesn't work if I put it in an init() method.
func LoadModels() {
	loadRectangles()
	loadCircles()
}

var _ parallax.Parallax = (*ParallaxRect)(nil)
var _ Shape = (*ParallaxRect)(nil)

type ParallaxRect struct {
	Rect
	Camera camera.Camera
	// Essentially, how this object moves in comparison to the camera. 1 is the same speed. 0.2 is 20% of camera speed.
	// The larger the number, the further away the object appears to be. For example, a ratio of 0.95 means the object
	// barely move when the camera moves - just like something that's very far away.
	// Negative numbers will make it move in the opposite direction, which isn't recommended.
	LocationRatio float32
}

func (r *ParallaxRect) GetParallaxPosition() mgl32.Vec2 {
	cPos := r.Camera.Position()
	return mgl32.Vec2{cPos.X()*r.LocationRatio + r.X, cPos.Y()*r.LocationRatio + r.Y}
}

func (r *ParallaxRect) Draw() {
	setDefaults()
	// TODO
}

func (r *ParallaxRect) DrawFilled() {
	setDefaults()
	// Save original position
	xTemp, yTemp := r.X, r.Y

	// Modify to place at correct parallax position.
	pos := r.GetParallaxPosition()
	r.X, r.Y = pos.X(), pos.Y()
	// Draw and then set original coordinates back.
	r.Rect.DrawFilled()
	r.X, r.Y = xTemp, yTemp
}

func setDefaults() {
	setColor(1, 0.1, 1, 1) // Default to a bright purple.
	shader.SetTranslationMatrix(0, 0, 0)
	shader.SetRotationMatrix(0, 0, 0)
	shader.SetScaleMatrix(1, 1, 1)
}

func setColor(r, g, b, a float32) {
	// TODO: Is bounding needed? Test what OpenGL does if given larger/smaller numbers?
	bound := func(x float32) float32 {
		if x > 1 {
			return 1
		}
		if x < 0 {
			return 0
		}
		return x
	}
	gl.Uniform4f(shader.ColorUniform, bound(r), bound(g), bound(b), bound(a))
}

func vec2ToFloat32(vecs []mgl32.Vec2) []float32 {
	a := make([]float32, len(vecs)*2)
	for i := 0; i < len(vecs); i++ {
		a[i*2] = vecs[i].X()
		a[i*2+1] = vecs[i].Y()
	}
	return a
}

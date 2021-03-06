package shape

import (
	"encoding/binary"
	"math/rand"

	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/goxjs/gl"
	"github.com/omustardo/demos/opengl/bytecoder"
	"github.com/omustardo/demos/opengl/entity"
	"github.com/omustardo/demos/opengl/shader"
	"github.com/omustardo/demos/opengl/util"
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

// Loads models into buffers on the GPU. Must be called after glfw.Init()
func LoadModels() {
	loadRectangles()
	loadCircles()
}

var _ Shape = (*ParallaxRect)(nil)

type ParallaxRect struct {
	Rect
	Target entity.Entity
	// Essentially, how this object moves in comparison to the camera (or some other target).
	// 1 is the same speed. 0.2 is 20% of camera speed.
	// The larger the number, the further away the object appears to be. For example, a ratio of 0.95 means the object
	// barely move when the camera moves - just like something that's very far away.
	// Negative numbers will make it move in the opposite direction, which isn't recommended.
	TranslationRatio float32
}

func GenParallaxRects(target entity.Entity, count int, minWidth, maxWidth, minSpeedRatio, maxSpeedRatio float32) []ParallaxRect {
	shapes := make([]ParallaxRect, count)
	for i := 0; i < count; i++ {
		shapes[i] = ParallaxRect{
			Rect: Rect{
				X: rand.Float32()*2000 - 1000, Y: rand.Float32()*2000 - 1000, // Note not even distribution - they are drawn from bottom left corner so everything is Up and Right shifted slightly
				R: rand.Float32(), G: rand.Float32(), B: rand.Float32(), A: 1,
				Width:  rand.Float32()*(maxWidth-minWidth) + minWidth,
				Height: rand.Float32()*(maxWidth-minWidth) + minWidth,
				Angle:  rand.Float32() * 360,
			},
			Target:           target,
			TranslationRatio: rand.Float32()*(maxSpeedRatio-minSpeedRatio) + minSpeedRatio,
		}
	}
	return shapes
}

func GetParallaxBuffers(arr []ParallaxRect) (parallaxPositionBuffer, parallaxTranslationBuffer, parallaxTranslationRatioBuffer, parallaxAngleBuffer, parallaxScaleBuffer, parallaxColorBuffer gl.Buffer) {
	lower, upper := float32(-0.5), float32(0.5)
	vertices := []float32{
		// Triangle 1
		lower, lower, 0,
		upper, upper, 0,
		lower, upper, 0,
		// Triangle 2
		lower, lower, 0,
		upper, lower, 0,
		upper, upper, 0,
	}
	var posData, transData, transRatioData, angleData, scaleData, colorData []float32
	for _, rect := range arr {
		posData = append(posData, vertices...)
		for i := 0; i < 6; i++ {
			tx, ty, _ := rect.Position().Elem()
			transData = append(transData, tx, ty)
			transRatioData = append(transRatioData, rect.TranslationRatio)
			angleData = append(angleData, rect.Angle)
			scaleData = append(scaleData, rect.Width, rect.Height)
			colorData = append(colorData, rect.R, rect.G, rect.B, rect.A)
		}
	}
	parallaxPositionBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, parallaxPositionBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, bytecoder.Float32(binary.LittleEndian, posData...), gl.STATIC_DRAW)

	parallaxTranslationBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, parallaxTranslationBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, bytecoder.Float32(binary.LittleEndian, transData...), gl.STATIC_DRAW)

	parallaxTranslationRatioBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, parallaxTranslationRatioBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, bytecoder.Float32(binary.LittleEndian, transRatioData...), gl.STATIC_DRAW)

	parallaxAngleBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, parallaxAngleBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, bytecoder.Float32(binary.LittleEndian, angleData...), gl.STATIC_DRAW)

	parallaxScaleBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, parallaxScaleBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, bytecoder.Float32(binary.LittleEndian, scaleData...), gl.STATIC_DRAW)

	parallaxColorBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, parallaxColorBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, bytecoder.Float32(binary.LittleEndian, colorData...), gl.STATIC_DRAW)
	return
}

func DrawParallaxBuffers(numObjects int, camPos mgl32.Vec2, parallaxPositionBuffer, parallaxTranslationBuffer, parallaxTranslationRatioBuffer, parallaxAngleBuffer, parallaxScaleBuffer, parallaxColorBuffer gl.Buffer) {
	gl.UseProgram(shader.Parallax.Program)
	gl.BindBuffer(gl.ARRAY_BUFFER, parallaxPositionBuffer)
	gl.VertexAttribPointer(shader.Parallax.PositionAttrib, 3 /* floats per vertex */, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(shader.Parallax.PositionAttrib)

	gl.BindBuffer(gl.ARRAY_BUFFER, parallaxTranslationBuffer)
	gl.VertexAttribPointer(shader.Parallax.TranslationAttrib, 2, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(shader.Parallax.TranslationAttrib)

	gl.BindBuffer(gl.ARRAY_BUFFER, parallaxTranslationRatioBuffer)
	gl.VertexAttribPointer(shader.Parallax.TranslationRatioAttrib, 1, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(shader.Parallax.TranslationRatioAttrib)

	gl.BindBuffer(gl.ARRAY_BUFFER, parallaxAngleBuffer)
	gl.VertexAttribPointer(shader.Parallax.AngleAttrib, 1, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(shader.Parallax.AngleAttrib)

	gl.BindBuffer(gl.ARRAY_BUFFER, parallaxScaleBuffer)
	gl.VertexAttribPointer(shader.Parallax.ScaleAttrib, 2, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(shader.Parallax.ScaleAttrib)

	gl.BindBuffer(gl.ARRAY_BUFFER, parallaxColorBuffer)
	gl.VertexAttribPointer(shader.Parallax.ColorAttrib, 4, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(shader.Parallax.ColorAttrib)

	shader.Parallax.SetCameraPosition(camPos)
	gl.DrawArrays(gl.TRIANGLES, 0, numObjects)
}

func (r *ParallaxRect) GetParallaxPosition() mgl32.Vec2 {
	cPos := r.Target.Position()
	return mgl32.Vec2{cPos.X()*r.TranslationRatio + r.X, cPos.Y()*r.TranslationRatio + r.Y}
}

func (r *ParallaxRect) Draw() {
	shader.Basic.SetDefaults()
	shader.Parallax.SetCameraPosition(r.Target.Position().Vec2())
}

func (r *ParallaxRect) DrawFilled() {
	shader.Basic.SetDefaults()
	// Save original position
	xTemp, yTemp := r.X, r.Y

	// Modify to place at correct parallax position.
	pos := r.GetParallaxPosition()
	r.X, r.Y = pos.X(), pos.Y()
	// Draw and then set original coordinates back.
	r.Rect.DrawFilled()
	r.X, r.Y = xTemp, yTemp
}

type OrbitingRect struct {
	Rect
	revolutionSpeed int64 // milliseconds to go entirely around the orbit. i.e. one year for the earth.
	orbit           Circle
	orbitTarget     entity.Entity // Makes the center of the orbit an object that can move. If nil, just uses the orbit's static center.
	rotateSpeed     int64         // milliseconds to do one full rotation. 0 to not rotate. i.e. one day for the earth.
}

func NewOrbitingRect(rect Rect, orbitCenter mgl32.Vec2, orbitRadius float32, orbitTarget entity.Entity, revolutionSpeed, rotateSpeed int64) *OrbitingRect {
	r := &OrbitingRect{
		Rect: rect,
		orbit: Circle{
			P:      orbitCenter.Vec3(0),
			Radius: orbitRadius,
			R:      0.6, G: 0.6, B: 0.6, A: 1.0,
		},
		orbitTarget:     orbitTarget,
		revolutionSpeed: revolutionSpeed,
		rotateSpeed:     rotateSpeed,
	}
	r.Update()
	return r
}

// TODO: Allow orbits and rotations to go counterclockwise.

func (r *OrbitingRect) Update() {
	if r.orbitTarget != nil {
		r.orbit.P = r.orbitTarget.Position()
	}
	now := util.GetTimeMillis()
	percentRevolution := float32(now%r.revolutionSpeed) / float32(r.revolutionSpeed)
	rads := percentRevolution * 2 * math.Pi
	offset := mgl32.Vec2{float32(math.Cos(float64(rads))), float32(math.Sin(float64(rads)))}.Mul(r.orbit.Radius)
	x, y := r.orbit.Center().Add(offset).Elem()
	r.SetCenter(x, y)

	percentRotation := float32(now%r.rotateSpeed) / float32(r.rotateSpeed)
	r.Angle = percentRotation * 2 * math.Pi
}

func (r *OrbitingRect) DrawOrbit() {
	r.orbit.Draw()
}

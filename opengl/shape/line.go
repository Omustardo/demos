package shape

import (
	"encoding/binary"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/goxjs/gl"
	"github.com/omustardo/demos/opengl/shader"
	"golang.org/x/mobile/exp/f32"
)

var _ Shape = (*Line)(nil)

type Line struct {
	P1, P2     mgl32.Vec3
	R, G, B, A float32
}

func (l *Line) SetCenter(x, y float32) {
	center := l.Center()
	l.P1[0] += x - center[0]
	l.P2[0] += x - center[0]
	l.P1[1] += y - center[1]
	l.P2[1] += y - center[1]
}
func (l *Line) ModifyCenter(x, y float32) {
	l.P1[0] += x
	l.P2[0] += x
	l.P1[1] += y
	l.P2[1] += y
}

func (l *Line) Center() mgl32.Vec2 {
	x := (l.P1[0] + l.P2[0]) / 2
	y := (l.P1[1] + l.P2[1]) / 2
	return mgl32.Vec2{x, y}
}

func (l *Line) Position() mgl32.Vec3 { // Shape implements entity.Entity - rethink this, right now it's the same as Center().
	return l.Center().Vec3(0)
}

func (l *Line) Draw() {
	setDefaults()
	vbuffer := gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, vbuffer)
	vertices := f32.Bytes(binary.LittleEndian,
		l.P1[0], l.P1[1], l.P1[2],
		l.P2[0], l.P2[1], l.P2[2],
	)
	gl.BufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(shader.VertexPositionAttrib) // https://www.opengl.org/sdk/docs/man2/xhtml/glEnableVertexAttribArray.xml
	itemSize := 3                                           // we use vertices made up of 3 floats
	itemCount := 2                                          // 2 points
	gl.VertexAttribPointer(shader.VertexPositionAttrib, itemSize, gl.FLOAT, false, 0, 0)

	// pMatrix := mgl32.Ortho2D(0, float32(WindowSize[0]), float32(WindowSize[1]), 0)
	// mvMatrix := mgl32.Translate3D(0, 0, 0) // Rectangle coordinates are being provided as world coords... TODO: have a basic shape and just translate it.
	// rotMatrix := mgl32.HomogRotate2D(angle) TODO: combine this with Projection and transform matrices in vertex shader

	setColor(l.R, l.G, l.B, l.A) // set color
	gl.DrawArrays(gl.LINES, 0, itemCount)

	cleanup(vbuffer)
}

// DrawFilled for a line is equivalent to Draw, but still required for the Shape interface.
func (l *Line) DrawFilled() {
	l.Draw()
}

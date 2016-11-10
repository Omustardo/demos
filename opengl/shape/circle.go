package shape

import (
	"encoding/binary"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/goxjs/gl"
	"github.com/omustardo/demos/opengl/shader"
	"golang.org/x/mobile/exp/f32"
)

const numCircleSegments = 60

var _ Shape = (*Circle)(nil)

var (

	// Vertices are the float32 coordinates of vertices that make up a circle, converted to a byte array.
	// This is the format required by OpenGL vertex buffers. These two buffers are used for all circles by modifying
	// the Scale, Rotation, and Translation matrices in the vertex shader.

	// triangles for drawing a filled circle.
	circleTriangleSegmentVertices []byte
	// line segments for drawing an empty circle.
	circleLineSegmentVertices []byte
)

func init() {
	// Generates triangles to make a full circle entered at [0,0]. Not just the edges.
	tmp := mgl32.Circle(1.0, 1.0, numCircleSegments)

	// The values are good as is for making triangles.
	circleTriangleSegmentVertices = f32.Bytes(binary.LittleEndian, vec2ToFloat32(tmp)...)

	// To get the line segment vertices, just ignore the first of every trio since that's the center.
	lineSegments := make([]mgl32.Vec2, numCircleSegments*2)
	for i := 0; i < numCircleSegments; i++ {
		lineSegments[i*2], lineSegments[i*2+1] = tmp[i*3+1], tmp[i*3+2]
	}
	circleLineSegmentVertices = f32.Bytes(binary.LittleEndian, vec2ToFloat32(lineSegments)...)
}

// TODO: Consider only having a single circle and modifying it in the shader via Rotate, Scale, Translate.
type Circle struct {
	P          mgl32.Vec3
	Radius     float32
	R, G, B, A float32
}

func (c *Circle) SetCenter(x, y float32) {
	c.P[0], c.P[1] = x, y
}
func (c *Circle) ModifyCenter(x, y float32) {
	c.P[0] += x
	c.P[1] += y
}
func (c *Circle) Center() mgl32.Vec2 {
	return c.P.Vec2()
}
func (c *Circle) Position() mgl32.Vec3 {
	return c.P
}

func (c *Circle) Draw() {
	setDefaults()
	setColor(c.R, c.G, c.B, c.A)
	shader.SetTranslationMatrix(c.P.X(), c.P.Y(), 0)
	shader.SetScaleMatrix(c.Radius, c.Radius, 0)

	vbuffer := gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, vbuffer)
	gl.BufferData(gl.ARRAY_BUFFER, circleLineSegmentVertices, gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(shader.VertexPositionAttrib) // https://www.opengl.org/sdk/docs/man2/xhtml/glEnableVertexAttribArray.xml
	itemSize := 2                                           // we use vertices made up of 2 floats
	gl.VertexAttribPointer(shader.VertexPositionAttrib, itemSize, gl.FLOAT, false, 0, 0)

	itemCount := numCircleSegments * 2 // 2 vertices per segment
	gl.DrawArrays(gl.LINE_LOOP, 0, itemCount)

	cleanup(vbuffer)
}

func (c *Circle) DrawFilled() {
	setDefaults()
	setColor(c.R, c.G, c.B, c.A)
	shader.SetTranslationMatrix(c.P.X(), c.P.Y(), 0)
	shader.SetScaleMatrix(c.Radius, c.Radius, 0)

	vbuffer := gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, vbuffer)
	gl.BufferData(gl.ARRAY_BUFFER, circleTriangleSegmentVertices, gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(shader.VertexPositionAttrib) // https://www.opengl.org/sdk/docs/man2/xhtml/glEnableVertexAttribArray.xml
	itemSize := 2                                           // we use vertices made up of 2 floats
	gl.VertexAttribPointer(shader.VertexPositionAttrib, itemSize, gl.FLOAT, false, 0, 0)

	// pMatrix := mgl32.Ortho2D(0, float32(WindowSize[0]), float32(WindowSize[1]), 0)
	// mvMatrix := mgl32.Translate3D(0, 0, 0) // Rectangle coordinates are being provided as world coords... TODO: have a basic shape and just translate it.
	// rotMatrix := mgl32.HomogRotate2D(angle) TODO: combine this with Projection and transform matrices in vertex shader

	itemCount := numCircleSegments // One triangle per segment
	gl.DrawArrays(gl.TRIANGLES, 0, itemCount)

	cleanup(vbuffer)
}

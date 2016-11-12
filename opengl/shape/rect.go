package shape

import (
	"encoding/binary"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/goxjs/gl"
	"github.com/omustardo/demos/opengl/shader"
	"golang.org/x/mobile/exp/f32"
)

var _ Shape = (*Rect)(nil)

// rectVertices are the float32 coordinates of two triangles (composing a 1x1 square), converted to a byte array.
// This is the format required by OpenGL vertex buffers. This one buffer is used for all rectangles by modifying
// the Scale, Rotation, and Translation matrices in the vertex shader.
// NOTE: Be careful of using len(rectVertices). It's NOT an array of floats - it's an array of bytes.
var (
	rectTriangleBuffer  gl.Buffer
	rectLineStripBuffer gl.Buffer
)

func loadRectangles() {
	lower, upper := float32(-0.5), float32(0.5)
	rectTriangleVertices := f32.Bytes(binary.LittleEndian,
		// Triangle 1
		lower, lower, 0,
		upper, upper, 0,
		lower, upper, 0,
		// Triangle 2
		lower, lower, 0,
		upper, lower, 0,
		upper, upper, 0,
	)
	rectTriangleBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, rectTriangleBuffer)                   // Bind the target buffer so we can store values in it. https://www.opengl.org/sdk/docs/man4/html/glBindBuffer.xhtml
	gl.BufferData(gl.ARRAY_BUFFER, rectTriangleVertices, gl.STATIC_DRAW) // store values in buffer

	rectLineStripVertices := f32.Bytes(binary.LittleEndian,
		lower, lower, 0,
		lower, upper, 0,
		upper, upper, 0,
		upper, lower, 0,
		lower, lower, 0,
	)
	rectLineStripBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, rectLineStripBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, rectLineStripVertices, gl.STATIC_DRAW)
}

type Rect struct {
	// X, Y are the center coordinate of the rectangle.
	X, Y          float32
	Width, Height float32
	// Angle is rotation around the center.
	Angle      float32
	R, G, B, A float32
}

func (r *Rect) Draw() {
	setDefaults()
	setColor(r.R, r.G, r.B, r.A)
	shader.SetRotationMatrix2D(r.Angle) // TODO: Make sure we rotate around the rectangle's center, rather than bottom left.
	shader.SetScaleMatrix(r.Width, r.Height, 0)
	shader.SetTranslationMatrix(r.X, r.Y, 0)

	gl.BindBuffer(gl.ARRAY_BUFFER, rectLineStripBuffer)
	itemSize := 3                                           // we use vertices made up of 3 floats
	itemCount := 5                                          // 4 segments, which requires 5 points
	gl.EnableVertexAttribArray(shader.VertexPositionAttrib) // https://www.opengl.org/sdk/docs/man2/xhtml/glEnableVertexAttribArray.xml
	gl.VertexAttribPointer(shader.VertexPositionAttrib, itemSize, gl.FLOAT, false, 0, 0)
	gl.DrawArrays(gl.LINE_STRIP, 0, itemCount)

	gl.DisableVertexAttribArray(shader.VertexPositionAttrib)
}

func (r *Rect) DrawFilled() {
	setDefaults()
	setColor(r.R, r.G, r.B, r.A)
	shader.SetRotationMatrix2D(r.Angle) // TODO: Make sure we rotate around the rectangle's center, rather than bottom left.
	shader.SetScaleMatrix(r.Width, r.Height, 0)
	shader.SetTranslationMatrix(r.X, r.Y, 0)

	gl.BindBuffer(gl.ARRAY_BUFFER, rectTriangleBuffer)
	itemSize := 3                                           // because the points consist of 3 floats
	itemCount := 6                                          // number of vertices in total
	gl.EnableVertexAttribArray(shader.VertexPositionAttrib) // https://www.opengl.org/sdk/docs/man2/xhtml/glEnableVertexAttribArray.xml
	gl.VertexAttribPointer(shader.VertexPositionAttrib, itemSize, gl.FLOAT, false, 0, 0)
	gl.DrawArrays(gl.TRIANGLES, 0, itemCount)

	gl.DisableVertexAttribArray(shader.VertexPositionAttrib)
}

func (r *Rect) SetCenter(x, y float32) {
	r.X = x
	r.Y = y
}

func (r *Rect) ModifyCenter(x, y float32) {
	r.X += x
	r.Y += y
}

func (r *Rect) Position() mgl32.Vec3 {
	return r.Center().Vec3(0)
}

func (r *Rect) Center() mgl32.Vec2 {
	return mgl32.Vec2{r.X, r.Y}
	/* Finding center coordinate based on bottom left x,y coords.
	rad := float64(mgl32.DegToRad(r.Angle))
	cos := float32(math.Cos(rad))
	sin := float32(math.Sin(rad))

	return mgl32.Vec2{
		r.X + r.Width/2*cos - r.Height/2*sin,
		r.Y + r.Width/2*sin + r.Height/2*cos,
	}
	*/
}

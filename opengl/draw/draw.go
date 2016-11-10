package draw

import (
//	"encoding/binary"
//	"github.com/goxjs/gl"
// "github.com/go-gl/mathgl/mgl32"
//	"golang.org/x/mobile/exp/f32"
)

var (
/*
	VertexPositionUniform gl.Uniform
	// ProjectionMatrixUniform  gl.Uniform
	// TranslationMatrixUniform gl.Uniform
	// RotationMatrixUniform    gl.Uniform
	// ScaleMatrixUniform       gl.Uniform
	MVMatrixUniform gl.Uniform
	PMatrixUniform  gl.Uniform
	ColorUniform    gl.Uniform

	VertexPositionAttrib gl.Attrib
*/
)

/*
func RectFilled(x1, y1, x2, y2, r, g, b, a float32) {
	pMatrix := mgl32.Ortho2D(0, float32(WindowSize[0]), float32(WindowSize[1]), 0)
	mvMatrix := mgl32.Translate3D(0, 0, 0) // Rectangle coordinates are being provided as world coords. TODO: have a basic rectangle shape and just translate it.
	// rotMatrix := mgl32.HomogRotate2D(angle) TODO: combine this with Projection and transform matrices in vertex shader

	// gl.UniformMatrix4fv(ProjectionMatrixUniform, pMatrix[:])   // view
	// gl.UniformMatrix4fv(TranslationMatrixUniform, mvMatrix[:]) // position
	// TODO: rotation

	gl.Uniform4f(ColorUniform, r, g, b, a) // set color

	// NOTE: Be careful of using len(vertices). It's NOT an array of floats - it's an array of bytes.
	vertices := f32.Bytes(binary.LittleEndian,
		// Triangle 1
		x1, y1, 0,
		x1, y2, 0,
		x2, y2, 0,
		// Triangle 2
		x1, y1, 0,
		x2, y1, 0,
		x2, y2, 0,
	)

	vbuffer := gl.CreateBuffer()                             // Generate buffer and returns a reference to it. https://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenBuffers.xml
	gl.BindBuffer(gl.ARRAY_BUFFER, vbuffer)                  // Bind the target buffer so we can store values in it. https://www.opengl.org/sdk/docs/man4/html/glBindBuffer.xhtml
	gl.BufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW) // store values in buffer

	itemSize := 3                                    // because the points consist of 3 floats
	gl.EnableVertexAttribArray(VertexPositionAttrib) // https://www.opengl.org/sdk/docs/man2/xhtml/glEnableVertexAttribArray.xml
	gl.VertexAttribPointer(VertexPositionAttrib, itemSize, gl.FLOAT, false, 0, 0)

	itemCount := 6 // itemSize is number of points
	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, itemCount)

	gl.DisableVertexAttribArray(VertexPositionAttrib)
	gl.DeleteBuffer(vbuffer) // TODO: Unsure if this is needed.
	// gl.BindBuffer(gl.ARRAY_BUFFER, gl.Buffer{Value: 0}) //Unbind buffer
}

func TriangleFilled(x1, y1, x2, y2, x3, y3, r, g, b, a float32) {
	// Create VBO. More info: http://www.songho.ca/opengl/gl_vbo.html
	triangleVertexPositionBuffer := gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, triangleVertexPositionBuffer)
	vertices := f32.Bytes(binary.LittleEndian,
		x1, y1, 0,
		x2, y2, 0,
		x3, y3, 0,
	)
	gl.BufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(VertexPositionAttrib) // https://www.opengl.org/sdk/docs/man2/xhtml/glEnableVertexAttribArray.xml
	itemSize := 3                                    // we use vertices made up of 3 floats
	gl.VertexAttribPointer(VertexPositionAttrib, itemSize, gl.FLOAT, false, 0, 0)

	pMatrix := mgl32.Ortho2D(0, float32(WindowSize[0]), float32(WindowSize[1]), 0)
	mvMatrix := mgl32.Translate3D(0, 0, 0) // Rectangle coordinates are being provided as world coords... TODO: have a basic shape and just translate it.
	// rotMatrix := mgl32.HomogRotate2D(angle) TODO: combine this with Projection and transform matrices in vertex shader

	gl.Uniform4f(ColorUniform, r, g, b, a)                     // set color
	gl.UniformMatrix4fv(ProjectionMatrixUniform, pMatrix[:])   // set Projection
	gl.UniformMatrix4fv(TranslationMatrixUniform, mvMatrix[:]) // set world transform (translate to position)
	itemCount := 3                                             // 3 points
	gl.DrawArrays(gl.TRIANGLES, 0, itemCount)

	// gl.BindBuffer(gl.ARRAY_BUFFER, gl.Buffer{Value: 0}) // Unbind buffer
	gl.DeleteBuffer(triangleVertexPositionBuffer)
	gl.DisableVertexAttribArray(VertexPositionAttrib)
}
*/
/*
func Line(x1, y1, x2, y2, r, g, b, a float32) {
	vbuffer := gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, vbuffer)
	vertices := f32.Bytes(binary.LittleEndian,
		x1, y1, 0,
		x2, y2, 0,
	)
	gl.BufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(VertexPositionAttrib) // https://www.opengl.org/sdk/docs/man2/xhtml/glEnableVertexAttribArray.xml
	itemSize := 3                                    // we use vertices made up of 3 floats
	gl.VertexAttribPointer(VertexPositionAttrib, itemSize, gl.FLOAT, false, 0, 0)

	// pMatrix := mgl32.Ortho2D(0, float32(WindowSize[0]), float32(WindowSize[1]), 0)
	// mvMatrix := mgl32.Translate3D(0, 0, 0) // Rectangle coordinates are being provided as world coords... TODO: have a basic shape and just translate it.
	// rotMatrix := mgl32.HomogRotate2D(angle) TODO: combine this with Projection and transform matrices in vertex shader

	gl.Uniform4f(ColorUniform, r, g, b, a) // set color
	// gl.UniformMatrix4fv(ProjectionMatrixUniform, pMatrix[:])   // set Projection
	// gl.UniformMatrix4fv(TranslationMatrixUniform, mvMatrix[:]) // set world transform (translate to position)
	itemCount := 2 // 2 points
	gl.DrawArrays(gl.LINES, 0, itemCount)

	// gl.BindBuffer(gl.ARRAY_BUFFER, gl.Buffer{Value: 0}) // Unbind buffer
	gl.DeleteBuffer(vbuffer)
	gl.DisableVertexAttribArray(VertexPositionAttrib)
}
*/

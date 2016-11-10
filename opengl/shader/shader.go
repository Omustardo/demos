package shader

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/goxjs/gl"
	"github.com/goxjs/gl/glutil"
)

// Online live shader editor: http://shdr.bkcore.com/

const (
	// VertexSource is a vertex shader. The goal is to get a vertex in world space to a point in OpenGL's -1 to +1 screen coordinates.
	VertexSource = `//#version 120 // OpenGL 2.1.
//#version 100 // WebGL.
attribute vec3 aVertexPosition;

// TODO: This calculation is the same for each vertex. Pass the combined matrix into a uTranslationRotationScaleMatrix
// so it isn't recalculated every time.
uniform mat4 uTranslationMatrix;
uniform mat4 uRotationMatrix;
uniform mat4 uScaleMatrix;

uniform mat4 uMVMatrix; // Model-View (transforms the input vertex to the camera's view of the world)
uniform mat4 uPMatrix;  // Projection (transforms camera's view into screen space)

void main() {
	vec4 worldPosition = uTranslationMatrix * uRotationMatrix * uScaleMatrix * vec4(aVertexPosition, 1.0);
	gl_Position = uPMatrix * uMVMatrix * worldPosition;

	// Original: gl_Position = uPMatrix * uMVMatrix * vec4(aVertexPosition, 1.0);
}
`
	FragmentSource = `//#version 120 // OpenGL 2.1.
//#version 100 // WebGL.
#ifdef GL_ES
precision highp float; // set floating point precision. TODO: Use mediump if performance is an issue.
#endif
uniform vec4 uColor;
void main() {
	gl_FragColor = uColor;
}
`
)

var (
	TranslationMatrixUniform gl.Uniform
	RotationMatrixUniform    gl.Uniform
	ScaleMatrixUniform       gl.Uniform

	MVMatrixUniform gl.Uniform
	PMatrixUniform  gl.Uniform
	ColorUniform    gl.Uniform

	// TODO: Learn more about shaders. What's attrib vs uniform?
	VertexPositionAttrib gl.Attrib
)

func SetupProgram() error {
	program, err := glutil.CreateProgram(VertexSource, FragmentSource)
	if err != nil {
		panic(err)
	}
	gl.ValidateProgram(program)
	if gl.GetProgrami(program, gl.VALIDATE_STATUS) != gl.TRUE {
		return fmt.Errorf("gl validate status: %s", gl.GetProgramInfoLog(program))
	}
	gl.UseProgram(program)

	// Get gl "names" of variables in the shader program.
	// https://www.opengl.org/sdk/docs/man/html/glUniform.xhtml
	TranslationMatrixUniform = gl.GetUniformLocation(program, "uTranslationMatrix")
	RotationMatrixUniform = gl.GetUniformLocation(program, "uRotationMatrix")
	ScaleMatrixUniform = gl.GetUniformLocation(program, "uScaleMatrix")
	MVMatrixUniform = gl.GetUniformLocation(program, "uMVMatrix")
	PMatrixUniform = gl.GetUniformLocation(program, "uPMatrix")
	ColorUniform = gl.GetUniformLocation(program, "uColor")
	VertexPositionAttrib = gl.GetAttribLocation(program, "aVertexPosition")
	return nil
}

func SetMVPMatrix(pMatrix, mvMatrix mgl32.Mat4) {
	gl.UniformMatrix4fv(PMatrixUniform, pMatrix[:])
	gl.UniformMatrix4fv(MVMatrixUniform, mvMatrix[:])
}

func SetTranslationMatrix(x, y, z float32) {
	translateMatrix := mgl32.Translate3D(x, y, z)
	gl.UniformMatrix4fv(TranslationMatrixUniform, translateMatrix[:])
}

func SetRotationMatrix2D(z float32) {
	rotationMatrix := mgl32.Rotate3DZ(z).Mat4() // TODO: Use quaternions.
	gl.UniformMatrix4fv(RotationMatrixUniform, rotationMatrix[:])
}
func SetRotationMatrix(x, y, z float32) {
	rotationMatrix := mgl32.Rotate3DX(x).Mul3(mgl32.Rotate3DY(y)).Mul3(mgl32.Rotate3DZ(z)).Mat4() // TODO: Use quaternions.
	gl.UniformMatrix4fv(RotationMatrixUniform, rotationMatrix[:])
}

func SetScaleMatrix(x, y, z float32) {
	scaleMatrix := mgl32.Scale3D(x, y, z)
	gl.UniformMatrix4fv(ScaleMatrixUniform, scaleMatrix[:])
}

package main

import (
	"encoding/binary"

	"fmt"

	"github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"golang.org/x/mobile/exp/f32"
)

const (
	// Shaders
	VertexSource = `
		#version 150

		in vec2 position;
		in vec3 color;

		out vec3 Color;

		void main()
		{
				Color = color;
				gl_Position = vec4(position, 0.0, 1.0);
		}
`
	FragmentSource = `
		#version 150

		in vec3 Color;

		out vec4 outColor;

		void main()
		{
				outColor = vec4(Color, 1.0);
		}
	`
)

func main() {
	err := glfw.Init(gl.ContextWatcher)
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()
	glfw.WindowHint(glfw.Samples, 16) // Anti-aliasing.
	glfw.WindowHint(glfw.Resizable, gl.FALSE)

	// Window hints to require OpenGL 3.2 or above, and to disable deprecated functions. glfw section of https://open.gl/context
	// These hints are not supported since we're using goxjs/glfw rather than the regular glfw, but should be used in a
	// standard desktop glfw project. TODO: Add support for these in goxjs/glfw/hint_glfw.go
	//glfw.WindowHint(glfw.ContextVersionMajor, 3)
	//glfw.WindowHint(glfw.ContextVersionMinor, 2)
	//glfw.WindowHint(glfw.OpenGLProfile, glfw.OPENGL_CORE_PROFILE)
	//glfw.WindowHint(glfw.OpenGLForwardCompatible, gl.TRUE)

	window, err := glfw.CreateWindow(800, 600, "OpenGL", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	fmt.Printf("OpenGL: %s %s %s; %v samples.\n", gl.GetString(gl.VENDOR), gl.GetString(gl.RENDERER), gl.GetString(gl.VERSION), gl.GetInteger(gl.SAMPLES))
	fmt.Printf("GLSL: %s.\n", gl.GetString(gl.SHADING_LANGUAGE_VERSION))

	// Vertex Input
	vertices := []float32{
		-0.5, 0.5, 1.0, 0.0, 0.0, // Top-left
		0.5, 0.5, 0.0, 1.0, 0.0, // Top-right
		0.5, -0.5, 0.0, 0.0, 1.0, // Bottom-right
		-0.5, -0.5, 1.0, 1.0, 1.0, // Bottom-left
	}

	verticesBytes := f32.Bytes(binary.LittleEndian, vertices...)
	vbo := gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)                           // Bind the target buffer so we can store values in it. https://www.opengl.org/sdk/docs/man4/html/glBindBuffer.xhtml
	gl.BufferData(gl.ARRAY_BUFFER, verticesBytes, gl.STATIC_DRAW) // store values in buffer

	// Element Array
	elements := Uint32ToBytes(binary.LittleEndian, 0, 1, 2, 2, 3, 0)
	ebo := gl.CreateBuffer()
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, elements, gl.STATIC_DRAW)

	// Compiling Shaders. Code is from https://github.com/goxjs/gl/blob/master/glutil/glutil.go
	vertexShader, err := loadShader(gl.VERTEX_SHADER, VertexSource)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := loadShader(gl.FRAGMENT_SHADER, FragmentSource)
	if err != nil {
		panic(err)
	}

	// Combining shaders into a program.
	program := gl.CreateProgram()
	if !program.Valid() {
		panic("no programs available")
	}
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	// Flag shaders for deletion when program is unlinked.
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	if gl.GetProgrami(program, gl.LINK_STATUS) == 0 {
		panic(gl.GetProgramInfoLog(program))
	}
	gl.ValidateProgram(program)
	if gl.GetProgrami(program, gl.VALIDATE_STATUS) != gl.TRUE {
		panic(gl.GetProgramInfoLog(program))
	}
	gl.UseProgram(program)

	// Making the link between vertex data and attributes
	posAttrib := gl.GetAttribLocation(program, "position")
	gl.VertexAttribPointer(posAttrib, 2, gl.FLOAT, false, 5*4, 0)
	gl.EnableVertexAttribArray(posAttrib)

	colAttrib := gl.GetAttribLocation(program, "color")
	gl.EnableVertexAttribArray(colAttrib)
	gl.VertexAttribPointer(colAttrib, 3, gl.FLOAT, false, 5*4, 2*4)

	for !window.ShouldClose() {
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, 0)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func loadShader(shaderType gl.Enum, src string) (gl.Shader, error) {
	shader := gl.CreateShader(shaderType)
	if !shader.Valid() {
		return gl.Shader{}, fmt.Errorf("glutil: could not create shader (type %v)", shaderType)
	}
	gl.ShaderSource(shader, src)
	gl.CompileShader(shader)
	if gl.GetShaderi(shader, gl.COMPILE_STATUS) == 0 {
		defer gl.DeleteShader(shader)
		return gl.Shader{}, fmt.Errorf("shader compile: %s", gl.GetShaderInfoLog(shader))
	}
	return shader, nil
}

// Bytes returns the byte representation of uint32 values in the given byte
// order. byteOrder must be either binary.BigEndian or binary.LittleEndian.
func Uint32ToBytes(byteOrder binary.ByteOrder, values ...uint32) []byte {
	le := false
	switch byteOrder {
	case binary.BigEndian:
	case binary.LittleEndian:
		le = true
	default:
		panic(fmt.Sprintf("invalid byte order %v", byteOrder))
	}

	b := make([]byte, 4*len(values))
	for i, v := range values {
		if le {
			b[4*i+0] = byte(v >> 0)
			b[4*i+1] = byte(v >> 8)
			b[4*i+2] = byte(v >> 16)
			b[4*i+3] = byte(v >> 24)
		} else {
			b[4*i+0] = byte(v >> 24)
			b[4*i+1] = byte(v >> 16)
			b[4*i+2] = byte(v >> 8)
			b[4*i+3] = byte(v >> 0)
		}
	}
	return b
}

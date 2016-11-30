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

		void main()
		{
			gl_Position = vec4(position, 0.0, 1.0);
		}
`
	FragmentSource = `
		#version 150

		out vec4 outColor;

		void main()
		{
				outColor = vec4(1.0, 1.0, 1.0, 1.0);
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
		0.0, 0.5, // Vertex 1 (X, Y)
		0.5, -0.5, // Vertex 2 (X, Y)
		-0.5, -0.5, // Vertex 3 (X, Y)
	}
	verticesBytes := f32.Bytes(binary.LittleEndian, vertices...)
	vbo := gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)                           // Bind the target buffer so we can store values in it. https://www.opengl.org/sdk/docs/man4/html/glBindBuffer.xhtml
	gl.BufferData(gl.ARRAY_BUFFER, verticesBytes, gl.STATIC_DRAW) // store values in buffer

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
	gl.VertexAttribPointer(posAttrib, 2, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(posAttrib)

	for !window.ShouldClose() {
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
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

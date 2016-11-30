package main

import (
	"github.com/goxjs/gl"
	"github.com/goxjs/glfw"
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

	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

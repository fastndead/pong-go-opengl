package glfwUtils

import "github.com/go-gl/glfw/v3.2/glfw"

var window *glfw.Window

func CreateWindow(name string, height int, width int) (*glfw.Window, func()) {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, int(glfw.OpenGLCoreProfile))
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	createdWindow, err := glfw.CreateWindow(width, height, name, nil, nil)
	if err != nil {
		panic(err)
	}

	createdWindow.MakeContextCurrent()

	window = createdWindow
	terminate := func() {
		glfw.Terminate()
	}
	return createdWindow, terminate
}

func AddKeyCallback(cb glfw.KeyCallback) {
	window.SetKeyCallback(cb)
}

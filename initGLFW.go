package main

import "github.com/go-gl/glfw/v3.2/glfw"

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, int(glfw.OpenGLCoreProfile))
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Test", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	window.SetKeyCallback(handleKeyPress)
	return window
}

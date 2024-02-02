package main

import (
	"go-graphics/glfwUtils"
	"go-graphics/openGlUtils"
	"runtime"
)

func main() {
	runtime.LockOSThread()

	window, terminate := glfwUtils.CreateWindow("Pong", 500, 500)
	defer terminate()
	glfwUtils.AddKeyCallback(keyPressCallback)

	program := openGlUtils.InitProgram()
	reset()

	for !window.ShouldClose() {
		handlePressedKeys()
		draw(window, program)
	}
}

package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

var (
	pressedKeys map[glfw.Key]bool = make(map[glfw.Key]bool)
)

func handleKeyPress(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		reset()
	}

	if action == glfw.Press {
		pressedKeys[key] = true
	}

	if action == glfw.Release {
		pressedKeys[key] = false
	}

	if key == glfw.KeyLeft && action == glfw.Press {
		directionAngle += 10
	}

	if key == glfw.KeyRight && action == glfw.Press {
		directionAngle -= 10
	}

	if key == glfw.KeyLeftControl && action == glfw.Press {
		currentSpeed -= 0.001
	}

	if key == glfw.KeyC && mods == glfw.ModControl {
		w.SetShouldClose(true)
	}
	if key == glfw.KeyR && action == glfw.Press {
		circleTrace = []float32{}
	}
	if key == glfw.KeyS && action == glfw.Press {
		showCircleTrace = !showCircleTrace
	}
}

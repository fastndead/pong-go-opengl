package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

var (
	pressedKeys map[glfw.Key]bool = make(map[glfw.Key]bool)
)

func keyPressCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		reset()
	}

	if action == glfw.Press {
		pressedKeys[key] = true
	}

	if action == glfw.Release {
		pressedKeys[key] = false
	}

	if key == glfw.KeyC && mods == glfw.ModControl {
		w.SetShouldClose(true)
	}
}

func handlePressedKeys() {
}

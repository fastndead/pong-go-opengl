package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/go-gl/glfw/v3.2/glfw"
)

var (
	startGameTimer time.Timer
)

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()

	program := initOpenGL()
	reset()
	for !window.ShouldClose() {
		select {
		case <-startGameTimer.C:
			fmt.Println("timer fired")
			startGame()
		default:
		}
		draw(window, program)
	}
}

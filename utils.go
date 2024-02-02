package main

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var (
	circles []*geometry
	lines   []*geometry
)

var globalAcceleration = []float32{0, -0.0001}

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	///////////////////////////////
	// The part where drawing executes
	drawables := getDrawables()

	for _, d := range drawables {
		d.draw()
	}
	///////////////////////////////

	glfw.PollEvents()
	window.SwapBuffers()
}

func getDrawables() []*drawable {
	detectCollisions(circles, lines)
	var drawables []*drawable

	for _, circle := range circles {
		updateCircle(circle)
		drawables = append(drawables, circle.drawable)
	}

	for _, line := range lines {
		updateLine(line)
		drawables = append(drawables, line.drawable)
	}

	return drawables
}

func reset() {
	fmt.Println("Reset")

	initialVelocityA := []float32{0.007, 0}
	initialVelocityB := []float32{-0.007, 0}
	circles = []*geometry{
		makeCircle([]float32{-0.5, 0.2}, initialVelocityA),
		makeCircle([]float32{0.3, 0.1}, initialVelocityB),
		makeCircle([]float32{0.5, 0.2}, initialVelocityB),
		makeCircle([]float32{0.8, 0.8}, initialVelocityB),
		makeCircle([]float32{0.6, 0.7}, initialVelocityB),
		makeCircle([]float32{0.4, 0.5}, initialVelocityB),
	}
	lines = []*geometry{
		makeLine([]float32{0.9, 0.9}, []float32{-0.9, 0.9}),   // top
		makeLine([]float32{-0.9, -0.9}, []float32{0.9, -0.9}), // bottom
		makeLine([]float32{0.9, 0.9}, []float32{0.9, -0.9}),   // right
		makeLine([]float32{-0.9, -0.9}, []float32{-0.9, 0.9}), // left
	}

}

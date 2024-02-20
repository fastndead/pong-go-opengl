package main

import (
	"fmt"

	"math/rand"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var (
	circles []*geometry
	lines   []*geometry
)

var globalAcceleration = []float32{0, -0.00009}

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
	addAcceleration(circles)
	detectCollisions(circles, lines)
	for _, circle := range circles {
		circle.position = calculatePosition(circle.position, circle.velocity)
	}

	for _, line := range lines {

		x1 := line.position[0]
		y1 := line.position[1]
		x2 := line.position[2]
		y2 := line.position[3]

		positionStart := calculatePosition([]float32{x1, y1}, line.velocity)
		positionEnd := calculatePosition([]float32{x2, y2}, line.velocity)

		line.position[0] = positionStart[0]
		line.position[1] = positionStart[1]
		line.position[2] = positionEnd[0]
		line.position[3] = positionEnd[1]
	}

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

	// initialVelocityA := []float32{0.007, 0}
	// initialVelocityB := []float32{-0.007, 0}

	circles = []*geometry{
		// makeCircle([]float32{-0.5, 0.2}, initialVelocityA),
		// makeCircle([]float32{0.3, 0.1}, initialVelocityB),
		// makeCircle([]float32{0.5, 0.2}, initialVelocityB),
		// makeCircle([]float32{0.8, 0.8}, initialVelocityB),
		// makeCircle([]float32{0.6, 0.7}, initialVelocityB),
		// makeCircle([]float32{0.4, 0.5}, initialVelocityB),
	}

	particleCount := 500
	for i := 0; i < particleCount; i++ {
		x, y := generateRandomPointInTrapezoid([]float32{-0.9, -0.9}, []float32{0.9, 0.9}, []float32{-0.03, 0}, []float32{0.03, 0})
		// vx := rand.Float32()*0.0016 - 0.008
		// vy := rand.Float32()*0.0016 - 0.008

		// x := float32(0)
		// y := float32(0)
		vx := float32(0)
		vy := float32(0)
		velocity := []float32{vx, vy}
		newCircle := makeCircle([]float32{x, y}, velocity)
		circles = append(circles, newCircle)
	}

	lines = []*geometry{
		makeLine([]float32{0.9, 0.9}, []float32{-0.9, 0.9}, 0.7),   // top
		makeLine([]float32{-0.9, -0.9}, []float32{0.9, -0.9}, 0.7), // bottom
		makeLine([]float32{0.9, 0.9}, []float32{0.9, -0.9}, 0.7),   // right
		makeLine([]float32{-0.9, -0.9}, []float32{-0.9, 0.9}, 0.7), // left
		makeLine([]float32{-1, 1}, []float32{-0.03, 0}, 0.8),       //
		makeLine([]float32{1, 1}, []float32{0.03, 0}, 0.8),         //
	}

}
func generateRandomPointInTrapezoid(topLeft, topRight, bottomLeft, bottomRight []float32) (float32, float32) {
	y := randomFloat32(0, 0.9-radius-0.01)

	xRightBound := ((topLeft[0]-bottomLeft[0])*(y-bottomLeft[1])/(topLeft[1]-bottomLeft[1]) - bottomLeft[0]) - radius
	x := randomFloat32(-xRightBound, xRightBound)

	return x, y

}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

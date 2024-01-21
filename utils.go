package main

import (
	"fmt"
	"math"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)

	return vao
}

func normalizeAngle(degrees float32) float32 {
	// Normalize the angle to be within -360 and 360 degrees
	for degrees < -360 || degrees > 360 {
		if degrees < -360 {
			degrees += 360
		} else if degrees > 360 {
			degrees -= 360
		}
	}

	// Convert negative angles to positive angles
	if degrees < 0 {
		degrees += 360
	}

	return degrees
}

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

func degreesToRadians(degrees float32) float32 {
	return degrees * (math.Pi / 180)
}

func handlePressedKeys() {
	if pressedKeys[glfw.KeyW] {
		leftPadY += 0.01
	}

	if pressedKeys[glfw.KeyS] {
		leftPadY -= 0.01
	}

	if pressedKeys[glfw.KeyUp] {
		rightPadY += 0.01
	}

	if pressedKeys[glfw.KeyDown] {
		rightPadY -= 0.01
	}
}

func calculateGeometry() {
	handlePressedKeys()
	// circle
	directionAngle = normalizeAngle(directionAngle)
	updateCircle(circle, float32(radius))
	xPosition += currentSpeed * float32(math.Cos(float64(degreesToRadians(float32(directionAngle)))))
	yPosition += currentSpeed * float32(math.Sin(float64(degreesToRadians(float32(directionAngle)))))

	// circle trace
	// updateCircleTrace(xPosition, yPosition)
	// trace := drawable{glType: gl.LINES, content: makeVao(circleTrace), length: int32(len(circleTrace))}

	// collisions
	detectCollisions()

	// pin pads
	updatePinPads()

	detectCircleOutOfBound()
}

func getDrawables() []*drawable {
	calculateGeometry()

	// assembling drawables
	drawables := []*drawable{circle, pinPads}
	/* if showCircleTrace {
		drawables = append(drawables, &trace)
	} */

	return drawables
}

func detectCircleOutOfBound() {
	if xPosition > float32(1)+radius || xPosition < float32(-1)-radius {
		reset()
	}
}

func detectCollisions() {
	// top
	if yPosition >= (float32(1) - radius) {
		yPosition = float32(1) - radius

		directionAngle = 360 - directionAngle
	}

	// bottom
	if yPosition <= (float32(-1) + radius) {
		yPosition = float32(-1) + radius
		directionAngle = 360 - directionAngle
	}

	// right
	/* if xPosition >= (float32(1) - radius) {
		xPosition = float32(1) - radius
		if directionAngle > 90 {
			directionAngle = 540 - directionAngle
		} else {
			directionAngle = 180 - directionAngle
		}
	} */

	// left
	/* if xPosition <= (float32(-1) + radius) {
		xPosition = float32(-1) + radius

		if directionAngle > 180 {
			directionAngle = 540 - directionAngle
		} else {
			directionAngle = 180 - directionAngle
		}
	} */

	// left pin pad
	if isCircleRectangleOverlap(xPosition, yPosition, radius, leftPadX, leftPadY, pinpadWidth, pinpadHeight) {
		xPosition = float32(leftPadX+pinpadWidth/2) + radius + 0.001

		if directionAngle > 180 {
			directionAngle = 540 - directionAngle
		} else {
			directionAngle = 180 - directionAngle
		}
	}

	if isCircleRectangleOverlap(xPosition, yPosition, radius, rightPadX, rightPadY, pinpadWidth, pinpadHeight) {
		xPosition = float32(rightPadX-pinpadWidth/2) - radius - 0.001

		if directionAngle > 90 {
			directionAngle = 540 - directionAngle
		} else {
			directionAngle = 180 - directionAngle
		}
	}

}

func isCircleRectangleOverlap(circleX, circleY, radius, rectX, rectY, width, height float32) bool {
	circleDistanseX := float32(math.Abs(float64(circleX) - float64(rectX)))
	circleDistanseY := float32(math.Abs(float64(circleY) - float64(rectY)))
	if circleDistanseX > width/2+radius {
		return false
	}

	if circleDistanseY > height/2+radius {
		return false
	}

	if circleDistanseX > width/2+radius {
		return false
	}

	if circleDistanseX <= width/2 {
		return true
	}
	if circleDistanseY <= height/2 {
		return true
	}
	return false
}

func startGame() {
	currentSpeed = 0.001
}

func reset() {
	fmt.Println("Reset")
	yPosition = float32(0)
	xPosition = float32(0)
	radius = float32(.1)
	startSpeed = float32(.01)
	acceleration = float32(.0000)
	currentSpeed = startSpeed
	directionAngle = float32(145)
	circle = initCircle()
	startGameTimer := time.NewTicker(time.Second * 5)
	defer startGameTimer.Stop()

	// pin pads
	leftPadX = -1
	leftPadY = 0
	rightPadX = 1
	rightPadY = 0
	pinPads = initPinPads()

	// circle trace
	// circleTrace = []float32{xPosition, yPosition, xPosition, yPosition}
	showCircleTrace = false
}

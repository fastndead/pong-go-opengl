package main

import (
	"math"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// circle properties
var (
	yPosition      float32
	xPosition      float32
	radius         float32
	startSpeed     float32
	acceleration   float32
	currentSpeed   float32
	directionAngle float32
	circle         *drawable
)

func updateCircle(d *drawable, radius float32) *drawable {
	var vertices []float32
	numSegments := segments

	var prevX, prevY float32
	for i := 0; i <= numSegments; i++ {
		angle := 2.0 * 3.14159 * float32(i) / float32(numSegments)
		y := float32(-radius)*float32(math.Cos(float64(angle))) + float32(yPosition)
		x := float32(radius)*float32(math.Sin(float64(angle))) + float32(xPosition)

		if i == 0 {
			prevX = x
			prevY = y
			continue
		}

		vertices = append(vertices, x, y)
		vertices = append(vertices, prevX, prevY)
		vertices = append(vertices, xPosition, yPosition)
		prevX = x
		prevY = y
	}

	d.content = makeVao(vertices)

	return d
}

func initCircle() *drawable {
	return &drawable{length: segments * 3, glType: gl.TRIANGLES}
}

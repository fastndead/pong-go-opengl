package main

import (
	"go-graphics/openGlUtils"
	"math"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	segments = 100
	radius   = 0.09
)

func updateCircle(circleGeometry *geometry) *drawable {
	d := circleGeometry.drawable
	position := calculatePosition(circleGeometry.position, circleGeometry.velocity, globalAcceleration)
	xPosition := position[0]
	yPosition := position[1]

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

	d.content = openGlUtils.MakeVao(vertices)

	return d
}

func makeCircle(position []float32, velocity []float32) *geometry {
	newCircleDrawable := &drawable{length: segments * 3, glType: gl.TRIANGLES}
	newCircleGeometry := &geometry{drawable: newCircleDrawable, position: position, velocity: velocity}
	updateCircle(newCircleGeometry)

	return newCircleGeometry
}

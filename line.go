package main

import (
	"go-graphics/openGlUtils"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func updateLine(lineGeometry *geometry) *drawable {
	d := lineGeometry.drawable

	x1 := lineGeometry.position[0]
	y1 := lineGeometry.position[1]
	x2 := lineGeometry.position[2]
	y2 := lineGeometry.position[3]

	positionStart := calculatePosition([]float32{x1, y1}, lineGeometry.velocity, []float32{0, 0})
	positionEnd := calculatePosition([]float32{x2, y2}, lineGeometry.velocity, []float32{0, 0})

	lineGeometry.position[0] = positionStart[0]
	lineGeometry.position[1] = positionStart[1]
	lineGeometry.position[2] = positionEnd[0]
	lineGeometry.position[3] = positionEnd[1]

	d.content = openGlUtils.MakeVao(
		[]float32{
			lineGeometry.position[0],
			lineGeometry.position[1],
			lineGeometry.position[2],
			lineGeometry.position[3],
		})

	return d
}

func makeLine(start, end []float32) *geometry {
	newLineDrawable := &drawable{length: 4, glType: gl.LINES}
	newLineGeometry := &geometry{drawable: newLineDrawable, position: append(start, end...), velocity: []float32{0, 0}}
	return newLineGeometry
}

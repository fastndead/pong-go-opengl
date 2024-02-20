package main

import (
	"go-graphics/openGlUtils"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func updateLine(lineGeometry *geometry) *drawable {
	d := lineGeometry.drawable

	d.content = openGlUtils.MakeVao(
		[]float32{
			lineGeometry.position[0],
			lineGeometry.position[1],
			lineGeometry.position[2],
			lineGeometry.position[3],
		})

	return d
}

func makeLine(start, end []float32, elascitity float32) *geometry {
	newLineDrawable := &drawable{length: 4, glType: gl.LINES}
	newLineGeometry := &geometry{drawable: newLineDrawable, position: append(start, end...), velocity: []float32{0, 0}, resting: true, elasticity: elascitity}
	return newLineGeometry
}

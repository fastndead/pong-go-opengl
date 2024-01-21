package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

var (
	leftPadX  float32
	leftPadY  float32
	rightPadX float32
	rightPadY float32

	pinPads *drawable
)

const pinpadWidth float32 = 0.01
const pinpadHeight float32 = 0.4

func initPinPads() *drawable {
	return &drawable{
		length: 12,
		glType: gl.TRIANGLES,
	}
}

func updatePinPads() *drawable {
	vertices := []float32{
		leftPadX + pinpadWidth,
		leftPadY + pinpadHeight/2,
		leftPadX + pinpadWidth,
		leftPadY - pinpadHeight/2,
		leftPadX - pinpadWidth,
		leftPadY + pinpadHeight/2,

		leftPadX - pinpadWidth,
		leftPadY - pinpadHeight/2,
		leftPadX + pinpadWidth,
		leftPadY - pinpadHeight/2,
		leftPadX - pinpadWidth,
		leftPadY + pinpadHeight/2,

		rightPadX + pinpadWidth,
		rightPadY + pinpadHeight/2,
		rightPadX + pinpadWidth,
		rightPadY - pinpadHeight/2,
		rightPadX - pinpadWidth,
		rightPadY + pinpadHeight/2,

		rightPadX - pinpadWidth,
		rightPadY - pinpadHeight/2,
		rightPadX + pinpadWidth,
		rightPadY - pinpadHeight/2,
		rightPadX - pinpadWidth,
		rightPadY + pinpadHeight/2,
	}

	pinPads.content = makeVao(vertices)
	return pinPads
}

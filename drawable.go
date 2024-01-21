package main

import "github.com/go-gl/gl/v4.1-core/gl"

type drawable struct {
	content uint32
	length  int32
	glType  uint32
}

func (d *drawable) draw() {
	gl.BindVertexArray(d.content)
	gl.DrawArrays(d.glType, 0, d.length)
}

package main

func calculatePosition(position []float32, velocity []float32, acceleration []float32) []float32 {
	position[0] += velocity[0] + acceleration[0]
	position[1] += velocity[1] + acceleration[1]
	return position
}

type geometry struct {
	drawable *drawable
	position []float32
	velocity []float32
}

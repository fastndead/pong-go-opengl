package main

import (
	"fmt"
	"math"
)

const restTreshold float64 = 0.0005

func elasticCollisionCircles(v1, v2, pos1, pos2 []float32, dampingFactor1, dampingFactor2 float32) ([]float32, []float32) {
	m1 := float64(radius)
	m2 := float64(radius)

	x1 := float64(pos1[0])
	y1 := float64(pos1[1])
	x2 := float64(pos2[0])
	y2 := float64(pos2[1])

	u1x := float64(v1[0])
	u1y := float64(v1[1])
	u2x := float64(v2[0])
	u2y := float64(v2[1])

	angle := math.Atan2(y2-y1, x2-x1)

	v1x := u1x*math.Cos(angle) + u1y*math.Sin(angle)
	v1y := u1y*math.Cos(angle) - u1x*math.Sin(angle)
	v2x := u2x*math.Cos(angle) + u2y*math.Sin(angle)
	v2y := u2y*math.Cos(angle) - u2x*math.Sin(angle)

	v1xPrime := ((m1-m2)*v1x + 2*m2*v2x) / (m1 + m2)
	v2xPrime := ((m2-m1)*v2x + 2*m1*v1x) / (m1 + m2)
	v1yPrime := v1y
	v2yPrime := v2y

	u1xPrime := v1xPrime*math.Cos(-angle) + v1yPrime*math.Sin(-angle)
	u1yPrime := v1yPrime*math.Cos(-angle) - v1xPrime*math.Sin(-angle)
	u2xPrime := v2xPrime*math.Cos(-angle) + v2yPrime*math.Sin(-angle)
	u2yPrime := v2yPrime*math.Cos(-angle) - v2xPrime*math.Sin(-angle)

	c1resultVelocity := []float32{float32(u1xPrime) * dampingFactor1, float32(u1yPrime) * dampingFactor1}
	c2resultVelocity := []float32{float32(u2xPrime) * dampingFactor2, float32(u2yPrime) * dampingFactor2}

	return c1resultVelocity, c2resultVelocity

}

func elasticCollisionCircleLine(vc, pc, pl []float32, dampingFactor float32) []float32 {
	pl1 := []float32{pl[0], pl[1]}
	pl2 := []float32{pl[2], pl[3]}
	dirLine := []float32{pl2[0] - pl1[0], pl2[1] - pl1[1]}

	vcx := vc[0] * dampingFactor
	vcy := vc[1] * dampingFactor

	lineLength := float32(math.Sqrt(float64(dirLine[0]*dirLine[0] + dirLine[1]*dirLine[1])))

	dirLine[0] /= lineLength
	dirLine[1] /= lineLength

	velocityNorm := []float32{-dirLine[1], dirLine[0]}
	dotProduct := 2 * (vcx*velocityNorm[0] + vcy*velocityNorm[1])

	resultV := []float32{vcx - dotProduct*velocityNorm[0], vcy - dotProduct*velocityNorm[1]}
	return resultV
}

func addAcceleration(gArr []*geometry) {
	for _, g := range gArr {
		g.velocity[0] = g.velocity[0] + globalAcceleration[0]
		g.velocity[1] = g.velocity[1] + globalAcceleration[1]

	}
}

func boundingBoxConstraint(circle *geometry) {
	const boundingBoxValue = float32(.9)
	if circle.position[0]+radius > boundingBoxValue {
		circle.position[0] = boundingBoxValue - radius

	}

	if circle.position[1]+radius > boundingBoxValue {
		circle.position[1] = boundingBoxValue - radius
	}

	if circle.position[0]-radius < -boundingBoxValue {
		circle.position[0] = -boundingBoxValue + radius
	}

	if circle.position[1]-radius < -boundingBoxValue {
		circle.position[1] = -boundingBoxValue + radius
	}
}

// todo: add sweep and prune to this, for now performance is dogshit
func detectCollisions(cGeometry []*geometry, lGeometry []*geometry) []*geometry {

	// collision detection between circle and a circle
	for i := 0; i < len(cGeometry); i++ {
		for j := i + 1; j < len(cGeometry); j++ {
			g1 := cGeometry[i]
			g2 := cGeometry[j]
			dist := math.Sqrt(math.Pow(float64(g2.position[0]-g1.position[0]), 2) + math.Pow(float64(g2.position[1]-g1.position[1]), 2))
			if dist <= float64(radius*2) {
				g1.velocity, g2.velocity = elasticCollisionCircles(g1.velocity, g2.velocity, g1.position, g2.position, g1.elasticity, g2.elasticity)

				clippingDot := []float32{(g1.position[0] + g2.position[0]) / 2, (g1.position[1] + g2.position[1]) / 2}
				g1.position = clipPosition(g1.position, clippingDot)
				g2.position = clipPosition(g2.position, clippingDot)

				// if g1.velocity[0] == 0 && g1.velocity[1] == 0 {
				// 	g1.resting = true
				// } else {
				// 	g1.resting = false
				// }
				// if g2.velocity[0] == 0 && g2.velocity[1] == 0 {
				// 	g2.resting = true
				// } else {
				// 	g2.resting = false
				// }

				if g1.resting && (g1.velocity[0] > 0 || g1.velocity[1] > 0) {
					fmt.Println("here's the error g1")
				}
				if g2.resting && (g2.velocity[0] > 0 || g2.velocity[1] > 0) {
					fmt.Println("here's the error g2")
				}
			}
		}
	}
	// collision detection between circle and a line
	for _, circle := range cGeometry {
		for _, line := range lGeometry {
			x1 := line.position[0]
			y1 := line.position[1]
			x2 := line.position[2]
			y2 := line.position[3]

			cx := circle.position[0]
			cy := circle.position[1]

			length := dist(x1, y1, x2, y2)
			dot := (((cx - x1) * (x2 - x1)) + ((cy - y1) * (y2 - y1))) / float32(math.Pow(float64(length), 2))

			closestX := x1 + (dot * (x2 - x1))
			closestY := y1 + (dot * (y2 - y1))

			onSegment := linePoint(x1, y1, x2, y2, closestX, closestY)
			if !onSegment {
				continue
			}

			distClosestToCircle := dist(closestX, closestY, cx, cy)

			if distClosestToCircle < radius {
				circle.velocity = elasticCollisionCircleLine(circle.velocity, circle.position, line.position, line.elasticity)
				// circle.position = clipPosition(circle.position, []float32{closestX, closestY})
			}
		}
		boundingBoxConstraint(circle)

		if len(circle.previousVelocities) > 100 {
			circle.previousVelocities = []float32{}
		}
		circle.previousVelocities = append(circle.previousVelocities, circle.velocity...)

		// if len(circle.previousVelocities) == 100 && !circle.resting {
		// 	var xVelocitiesSum, yVelocitiesSum float32
		// 	for i := 0; i < len(circle.previousVelocities); i += 2 {
		// 		xVelocitiesSum += circle.previousVelocities[i]
		// 		yVelocitiesSum += circle.previousVelocities[i+1]
		// 	}
		// 	fmt.Println(xVelocitiesSum)
		// 	fmt.Println(yVelocitiesSum)
		//
		// 	restingThresholdPosition := float64(.02)
		// 	if math.Abs(float64(xVelocitiesSum)) < restingThresholdPosition && math.Abs(float64(yVelocitiesSum)) < restingThresholdPosition {
		// 		fmt.Println("found rester")
		// 		circle.resting = true
		// 		circle.velocity = []float32{0, 0}
		// 	}
		// }

	}

	return cGeometry
}

func clipPosition(cPos, clipTarget []float32) []float32 {
	directionVector := []float32{clipTarget[0] - cPos[0], clipTarget[1] - cPos[1]}
	directionVectorMagnitude := math.Sqrt(float64(directionVector[0]*directionVector[0] + directionVector[1]*directionVector[1]))

	unitVector := []float32{directionVector[0] / float32(directionVectorMagnitude), directionVector[1] / float32(directionVectorMagnitude)}

	resultPos := []float32{clipTarget[0] - unitVector[0]*radius, clipTarget[1] - unitVector[1]*radius}

	return resultPos
}

func linePoint(x1, y1, x2, y2, px, py float32) bool {
	d1 := dist(px, py, x1, y1)
	d2 := dist(px, py, x2, y2)

	d := dist(x1, y1, x2, y2)

	var buffer float64 = 0.0001
	return math.Abs(float64(d1+d2-d)) < buffer
}

func dist(x1, y1, x2, y2 float32) float32 {
	distX := x1 - x2
	distY := y1 - y2
	return float32(math.Sqrt(float64(distX*distX + distY*distY)))
}

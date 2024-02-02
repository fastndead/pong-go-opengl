package main

import (
	"math"
)

func elasticCollisionCircles(v1, v2, pos1, pos2 []float32) ([]float32, []float32) {
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

	return []float32{float32(u1xPrime), float32(u1yPrime)}, []float32{float32(u2xPrime), float32(u2yPrime)}
}

func elasticCollisionCircleLine(vc, pc, pl []float32) []float32 {
	pl1 := []float32{pl[0], pl[1]}
	pl2 := []float32{pl[2], pl[3]}
	// Calculate the direction vector of the line
	dirLine := []float32{pl2[0] - pl1[0], pl2[1] - pl1[1]}

	// Calculate the vector from one end of the line segment to the circle's center
	dirToPoint := []float32{pc[0] - pl1[0], pc[1] - pl1[1]}

	// Calculate the length of the line segment
	lineLength := float32(math.Sqrt(float64(dirLine[0]*dirLine[0] + dirLine[1]*dirLine[1])))

	// Normalize the direction vector of the line
	dirLine[0] /= lineLength
	dirLine[1] /= lineLength

	// Project the vector from one end of the line segment to the circle's center onto the direction vector of the line
	distanceAlongLine := dirToPoint[0]*dirLine[0] + dirToPoint[1]*dirLine[1]

	// Calculate the closest point on the line to the circle's center
	closestPoint := []float32{pl1[0] + distanceAlongLine*dirLine[0], pl1[1] + distanceAlongLine*dirLine[1]}

	// Calculate the vector from the circle's center to the closest point on the line
	dirToClosest := []float32{closestPoint[0] - pc[0], closestPoint[1] - pc[1]}

	// Calculate the distance from the circle's center to the closest point on the line
	distanceToClosest := float32(math.Sqrt(float64(dirToClosest[0]*dirToClosest[0] + dirToClosest[1]*dirToClosest[1])))

	// Check if the circle is overlapping with the line
	if distanceToClosest <= radius {
		// Calculate the reflection of the circle's velocity vector about the line (elastic collision)
		velocityNorm := []float32{-dirLine[1], dirLine[0]}
		dotProduct := 2 * (vc[0]*velocityNorm[0] + vc[1]*velocityNorm[1])
		reflection := []float32{vc[0] - dotProduct*velocityNorm[0], vc[1] - dotProduct*velocityNorm[1]}
		return reflection
	}
	// If the circle and the line do not overlap, return the original velocity
	return vc
}

func detectCollisions(cGeometry []*geometry, lGeometry []*geometry) []*geometry {
	for i := 0; i < len(cGeometry); i++ {
		for j := i + 1; j < len(cGeometry); j++ {
			g1 := cGeometry[i]
			g2 := cGeometry[j]
			dist := math.Sqrt(math.Pow(float64(g2.position[0]-g1.position[0]), 2) + math.Pow(float64(g2.position[1]-g1.position[1]), 2))
			if dist <= float64(radius*2) {
				v1, v2 := elasticCollisionCircles(g1.velocity, g2.velocity, g1.position, g2.position)

				g1.velocity = v1
				g2.velocity = v2
			}
		}
	}

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
			// todo radius here is not tied to an instance of a circle
			if distClosestToCircle <= radius {

				circle.velocity = elasticCollisionCircleLine(circle.velocity, circle.position, line.position)
			}
		}
	}
	return cGeometry
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

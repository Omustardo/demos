package shape

import "github.com/go-gl/mathgl/mgl32"

const axisLength = 1e9

var xAxis, yAxis, zAxis *Line

func init() {
	xAxis = &Line{
		P1: mgl32.Vec3{-axisLength, 0, 0},
		P2: mgl32.Vec3{axisLength, 0, 0},
		R:  1, G: 0, B: 0, A: 1,
	}
	yAxis = &Line{
		P1: mgl32.Vec3{0, -axisLength, 0},
		P2: mgl32.Vec3{0, axisLength, 0},
		R:  0, G: 1, B: 0, A: 1,
	}
	zAxis = &Line{
		P1: mgl32.Vec3{0, 0, -axisLength},
		P2: mgl32.Vec3{0, 0, axisLength},
		R:  0, G: 0, B: 1, A: 1,
	}
}

// DrawXYZAxes is a utility function that simply draws the three basic X,Y,Z axes
func DrawXYZAxes() {
	xAxis.Draw()
	yAxis.Draw()
	zAxis.Draw()
}

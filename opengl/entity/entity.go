package entity

import "github.com/go-gl/mathgl/mgl32"
type Entity interface{
	Position() mgl32.Vec3
}

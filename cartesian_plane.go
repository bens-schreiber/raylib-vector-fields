package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CartesianPlane struct {
	Origin rl.Vector2
}

func CreateCartesianPlane() CartesianPlane {
	return CartesianPlane{
		Origin: rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2},
	}
}

func (cp *CartesianPlane) Draw() {
	rl.DrawLineV(rl.NewVector2(0, cp.Origin.Y), rl.NewVector2(float32(rl.GetScreenWidth()), cp.Origin.Y), rl.Gray)
	rl.DrawLineV(rl.NewVector2(cp.Origin.X, 0), rl.NewVector2(cp.Origin.X, float32(rl.GetScreenHeight())), rl.Gray)
}

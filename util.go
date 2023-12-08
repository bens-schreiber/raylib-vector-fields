package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func RandomVector2InScreen() rl.Vector2 {
	screenWidth := int32(rl.GetScreenWidth())
	screenHeight := int32(rl.GetScreenHeight())
	return rl.Vector2{X: float32(rl.GetRandomValue(0, screenWidth)), Y: float32(rl.GetRandomValue(0, screenHeight))}
}

package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1000, 1000, "")

	rl.SetTargetFPS(60)

	rl.DisableCursor()

	plane := CreateCartesianPlane()
	origin := plane.Origin

	bodies := CreateBodies(10000)

	for !rl.WindowShouldClose() {

		for i := range bodies {
			UpdateVectorField(&bodies[i], origin)
			bodies[i].DecrementTTL()
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for _, body := range bodies {
			body.Draw()
		}

		plane.Draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

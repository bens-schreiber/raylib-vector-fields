package main

import (
	math "math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/*
UpdateVectorField updates the position of a body using polar coordinates.

Note that the origin is in the center of the screen. Also note that up is negative y, and right is positive x,
because raylib's coordinate system is originated at (0,0) in the top left corner of the screen.

To adjust for this, we use the magnitude from the distance to our true origin (the center of the screen)

F(x, y) = (

	sqrt((x - originX)^2 + (y - originY)^2) * cos(atan2(y - originY, x - originX) + 0.0001 * sqrt((x - originX)^2 + (y - originY)^2)),
	sqrt((x - originX)^2 + (y - originY)^2) * sin(atan2(y - originY, x - originX) + 0.0001 * sqrt((x - originX)^2 + (y - originY)^2))

)

or

F(x, y) = (

	magnitudeFromOrigin * cos(angleFromOrigin + 0.0001 * magnitudeFromOrigin),
	magnitudeFromOrigin * sin(angleFromOrigin + 0.0001 * magnitudeFromOrigin)

)
*/
func UpdateVectorField(b *Body, origin rl.Vector2) {
	b.PreviousPosition = b.Position

	distance := math.Sqrt(math.Pow(float64(b.Position.X-origin.X), 2) + math.Pow(float64(b.Position.Y-origin.Y), 2))

	angle := math.Atan2(float64(b.Position.Y-origin.Y), float64(b.Position.X-origin.X)) + 0.0001*distance

	b.Position.X = origin.X + float32(distance*math.Cos(angle))
	b.Position.Y = origin.Y + float32(distance*math.Sin(angle))
}

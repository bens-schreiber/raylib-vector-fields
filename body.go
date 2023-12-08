package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Body struct {
	PreviousPosition rl.Vector2
	Position         rl.Vector2
	Color            rl.Color
	TTL              int32
}

func (b *Body) Draw() {
	rl.DrawCircleV(b.Position, 1, b.Color)
	rl.DrawLineV(b.PreviousPosition, b.Position, b.Color)
}

func (b *Body) DecrementTTL() {
	b.TTL--
	if b.TTL > 0 {
		return
	}

	// Reset body position and TTL on timeout
	b.Position = RandomVector2InScreen()
	b.TTL = rl.GetRandomValue(0, 1000)
	b.PreviousPosition = b.Position
}

func CreateBody(position rl.Vector2) Body {
	return Body{
		PreviousPosition: position,
		Position:         position,
		Color:            rl.NewColor(uint8(rl.GetRandomValue(0, 255)), uint8(rl.GetRandomValue(0, 255)), uint8(rl.GetRandomValue(0, 255)), 255),
		TTL:              rl.GetRandomValue(0, 1000),
	}
}

func CreateBodies(count int) []Body {
	bodies := make([]Body, count)
	for i := range bodies {
		bodies[i] = CreateBody(RandomVector2InScreen())
	}
	return bodies
}

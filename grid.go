package main

import (
	"math/rand"
	"time"
)

type entity struct {
	x     int
	y     int
	alive int
}

type Grid struct {
	entities []entity
}

func makeGrid(columns int, rows int) *Grid {
	g := &Grid{}
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			rand.Seed(time.Now().UnixNano())
			g.entities = append(g.entities, entity{x: i, y: j, alive: rand.Intn(2)})
		}
	}

	return g
}

package main

import "fmt"

type coordinate struct {
	x int
	y int
}

type Grid struct {
	coordinates []coordinate
}

func makeGrid(columns int, rows int) *Grid {
	g := &Grid{}
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			g.coordinates = append(g.coordinates, coordinate{x: i, y: j})
		}
	}
	fmt.Println(columns * rows)
	fmt.Println(len(g.coordinates))

	return g
}

package main

func getLivingNeighbours(e *entity, i int, g *Grid) []entity {
	livingNeighbours := make([]entity, 0)

	if e.x != 0 {
		left := g.entities[i-screenHeight]
		if left.alive == 1 {
			livingNeighbours = append(livingNeighbours, left)
		}
	}
	if e.x < screenWidth-1 {
		right := g.entities[i+screenHeight]
		if right.alive == 1 {
			livingNeighbours = append(livingNeighbours, right)
		}
	}
	if e.y > 0 {
		up := g.entities[i-1]
		if up.alive == 1 {
			livingNeighbours = append(livingNeighbours, up)
		}
	}
	if e.y < screenHeight-1 {
		down := g.entities[i+1]
		if down.alive == 1 {
			livingNeighbours = append(livingNeighbours, down)
		}
	}
	if e.x < screenWidth-1 && e.y > 0 {
		upRight := g.entities[i+(screenHeight-1)]
		if upRight.alive == 1 {
			livingNeighbours = append(livingNeighbours, upRight)
		}
	}
	if e.x > 0 && e.y > 0 {
		upLeft := g.entities[i-(screenHeight-1)]
		if upLeft.alive == 1 {
			livingNeighbours = append(livingNeighbours, upLeft)
		}
	}
	if e.x < screenWidth-1 && e.y < screenHeight-1 {
		downRight := g.entities[i+screenHeight]
		if downRight.alive == 1 {
			livingNeighbours = append(livingNeighbours, downRight)
		}
	}
	if e.x != 0 && e.y < screenHeight-1 {
		downLeft := g.entities[i-(screenHeight-1)]
		if downLeft.alive == 1 {
			livingNeighbours = append(livingNeighbours, downLeft)
		}
	}

	return livingNeighbours
}

func shouldLive(e *entity, i int, g *Grid) bool {
	living := getLivingNeighbours(e, i, g)

	if len(living) == 3 && e.alive == 0 {
		return true
	}
	if len(living) == 2 {
		if e.alive == 1 {
			return true
		}
	}

	return false
}

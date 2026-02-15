package game

import world "proc-gen-world/World"

type Game struct {
	world.World
	
}

func NewGame() (Game, error) {
	return Game{}, nil
}

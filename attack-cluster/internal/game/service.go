package game

import "sync"

type Game struct {
	attacks int
	mu      sync.Mutex
}

func NewGame() *Game {
	return &Game{attacks: 0}
}

func (g *Game) ApplyAttack() int {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.attacks += 1

	return g.attacks
}

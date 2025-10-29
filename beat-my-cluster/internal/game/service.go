package game

import "sync"

type Game struct {
	ClusterHealth int
	mu            sync.Mutex
}

func NewGame() *Game {
	return &Game{ClusterHealth: 100}
}

func (g *Game) ApplyAttack(intensity int) int {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.ClusterHealth -= intensity
	if g.ClusterHealth < 0 {
		g.ClusterHealth = 0
	}
	return g.ClusterHealth
}

func (g *Game) Heal(amount int) int {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.ClusterHealth += amount
	if g.ClusterHealth > 100 {
		g.ClusterHealth = 100
	}
	return g.ClusterHealth
}

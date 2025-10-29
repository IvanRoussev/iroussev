package api

import (
	"net/http"

	"github.com/IvanRoussev/iroussev/beat-my-cluster/internal/db"
	"github.com/gin-gonic/gin"
)

type AttackRequest struct {
	PlayerName string `json:"player" binding:"required"`
	Intensity  int    `json:"intensity" binding:"required"`
}

// POST /attack
func (s *Server) HandleAttack(c *gin.Context) {
	var req AttackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find or create player
	var player db.Player
	if err := s.DB.FirstOrCreate(&player, db.Player{Name: req.PlayerName}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find/create player"})
		return
	}

	// Record attack
	attack := db.Attack{
		PlayerID:  player.ID,
		Intensity: req.Intensity,
	}
	if err := s.DB.Create(&attack).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record attack"})
		return
	}

	// Update cluster health
	health := s.Game.ApplyAttack(req.Intensity)

	c.JSON(http.StatusOK, gin.H{
		"message":        "Attack recorded",
		"cluster_health": health,
	})
}

// GET /leaderboard
func (s *Server) HandleLeaderboard(c *gin.Context) {
	type PlayerScore struct {
		Name        string
		TotalAttack int
	}

	var leaderboard []PlayerScore

	s.DB.Model(&db.Attack{}).
		Select("players.name, SUM(attacks.intensity) as total_attack").
		Joins("left join players on players.id = attacks.player_id").
		Group("players.name").
		Order("total_attack desc").
		Scan(&leaderboard)

	c.JSON(http.StatusOK, gin.H{"leaderboard": leaderboard})
}

// GET /healthz
func (s *Server) HandleHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"cluster_health": s.Game.ClusterHealth})
}

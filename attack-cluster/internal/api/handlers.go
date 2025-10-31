package api

import (
	"fmt"
	"net/http"

	"github.com/IvanRoussev/iroussev/attack-cluster/internal/db"
	"github.com/gin-gonic/gin"
)

type AttackRequest struct {
	PlayerName string `json:"player" binding:"required"`
}

func (s *Server) HandleAttack(c *gin.Context) {
	var req AttackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Print(req.PlayerName)
	var player db.Player

	result := s.DB.Where("name = ?", req.PlayerName).FirstOrCreate(&player, db.Player{Name: req.PlayerName})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	player.NumAttacks++

	fmt.Println(&player)

	if err := s.DB.Save(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"number_attacks": player.NumAttacks,
	})

}

func (s *Server) HandleLeaderboard(c *gin.Context) {

	var players []db.Player

	err := s.DB.Order("num_attacks DESC").Find(&players).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"leaderboard": players})
}

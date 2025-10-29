package api

import (
	"github.com/IvanRoussev/iroussev/beat-my-cluster/internal/game"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Router *gin.Engine
	DB     *gorm.DB
	Game   *game.Game
}

func (s *Server) SetupRouter() {
	r := gin.Default()

	r.POST("/attack", s.HandleAttack)
	r.GET("/leaderboard", s.HandleLeaderboard)
	r.GET("/health", s.HandleHealth)

	s.Router = r
}

func (s *Server) Start(address string) error {
	return s.Router.Run(address)
}

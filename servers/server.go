package servers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Server contains the server resources
type Server struct {
	Router gin.Engine
	DB     *gorm.DB
}

func New(db *gorm.DB) *Server {
	return nil
}

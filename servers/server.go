package servers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var (
// userService service.UserService
)

// Server contains the server resources
type Server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

// New new server initialization
func New(db *gorm.DB) *Server {
	r := gin.Default()
	server := &Server{
		DB:     db,
		Router: r,
	}
	server.userRoutes()
	return server
}

// Start begin the server execution process
func (s *Server) Start(port string) error {
	return s.Router.Run(port)
}

func created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func badRequestFromError(c *gin.Context, err error) {
	log.Error().Err(err)
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"message": err.Error(),
		"data":    nil,
	})
}

package servers

import (
	"context"
	"fmt"

	"github.com/LordRahl90/auth-manager-ot/domain/users/entities"
	"github.com/LordRahl90/auth-manager-ot/domain/users/interfaces"
	"github.com/LordRahl90/auth-manager-ot/domain/users/service"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
)

var userService interfaces.UserService

func (s *Server) userRoutes() {
	userService = service.DefaultUserService(s.DB)
	user := s.Router.Group("users")
	{
		user.POST("", createUser)
		user.POST("/authenticate", authenticate)
		user.GET("/me", getUserDetails)
	}
}

func authenticate(c *gin.Context) {

}

func createUser(c *gin.Context) {
	ctx := context.Background()
	ctx, span := otel.Tracer("users").Start(ctx, "UserHandle_Create")
	defer span.End()

	var view entities.User
	if err := c.ShouldBindJSON(&view); err != nil {
		badRequestFromError(c, err)
		return
	}

	err := userService.Create(ctx, &view)
	if err != nil {
		badRequestFromError(c, err)
		return
	}

	created(c, "user created successfully", view)
}

func getUserDetails(ctx *gin.Context) {
	fmt.Printf("hello user details")
}

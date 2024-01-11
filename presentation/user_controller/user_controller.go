// presentation/user_controller.go - Camada de Apresentação (Presentation Layer)
package presentation

import (
	"github.com/RafaelCava/kitkit-back-go/domain/user_usecase"
	"github.com/gin-gonic/gin"
)

// UserController lida com as solicitações relacionadas ao usuário.
type UserController struct {
	userService user_usecase.UserService
}

// NewUserController cria uma nova instância de UserController.
func NewUserController(userService user_usecase.UserService) *UserController {
	return &UserController{userService}
}

// RegisterRoutes registra rotas relacionadas ao usuário.
func (controller *UserController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/:id", controller.getUserByID)
}

func (controller *UserController) getUserByID(c *gin.Context) {
	userID := c.Param("id")
	// Converter userID para uint ou lidar com erros, dependendo da sua lógica
	// ...

	// Obter usuário usando o serviço de usuário
	user, err := controller.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Falha ao obter usuário"})
		return
	}

	// Responder com o usuário
	c.JSON(200, user)
}

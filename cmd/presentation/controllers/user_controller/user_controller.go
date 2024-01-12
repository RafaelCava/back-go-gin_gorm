// presentation/controller/user_controller/user_controller.go - Camada de Apresentação (Presentation Layer)
package user_controller

import (
	"github.com/RafaelCava/kitkit-back-go/cmd/domain/models/user_models"
	"github.com/RafaelCava/kitkit-back-go/cmd/domain/usecases/user_usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
	RegisterRoutes(router *gin.RouterGroup)
	getUserByID(c *gin.Context)
	createUser(c *gin.Context)
}

// UserControllerImpl lida com as solicitações relacionadas ao usuário.
type UserControllerImpl struct {
	userService user_usecase.UserService
}

// NewUserControllerImpl cria uma nova instância de UserController.
func NewUserControllerImpl(userService user_usecase.UserService) *UserControllerImpl {
	return &UserControllerImpl{userService}
}

// RegisterRoutes registra rotas relacionadas ao usuário.
func (controller *UserControllerImpl) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/:id", controller.getUserByID)
	router.POST("", controller.createUser)
}

func (controller *UserControllerImpl) getUserByID(c *gin.Context) {
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

type CreateUserRequest struct {
	Name     *string `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

func (controller *UserControllerImpl) createUser(c *gin.Context) {
	var request CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Falha ao criar usuário"})
		return
	}
	user := &user_models.User{
		ID:       uuid.New().String(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	userID, err := controller.userService.CreateUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Falha ao criar usuário"})
		return
	}

	// Responder com o ID do usuário
	c.JSON(200, gin.H{"id": userID})
}

// presentation/controller/user_controller/user_controller.go - Camada de Apresentação (Presentation Layer)
package user_controller

import (
	"net/http"

	"github.com/RafaelCava/kitkit-back-go/cmd/domain/models/user_models"
	"github.com/RafaelCava/kitkit-back-go/cmd/domain/usecases/user_usecase"
	"github.com/RafaelCava/kitkit-back-go/cmd/presentation/utils/http_util"
	_ "github.com/RafaelCava/kitkit-back-go/docs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
	RegisterRoutes(router *gin.RouterGroup)
	getUserByID(c *gin.Context)
	createUser(c *gin.Context)
	getAllUsers(c *gin.Context)
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
	router.GET("", controller.getAllUsers)
	router.POST("", controller.createUser)
}

// GetUserById godoc
//
//		@Summary		Retorna um usuário
//		@Description	Retorna um usuário pelo ID
//		@Tags			Users
//		@Accept			json
//		@Produce		json
//		@Param			id	path		string	true	"User ID"
//	  @Success		200	{object}	user_models.User
//	  @Failure		400	{object}	http_util.HTTPError
//		@Router			/users/{id} [get]
func (controller *UserControllerImpl) getUserByID(c *gin.Context) {
	userID := c.Param("id")
	// Converter userID para uint ou lidar com erros, dependendo da sua lógica
	// ...

	// Obter usuário usando o serviço de usuário
	user, err := controller.userService.GetUserByID(userID)
	if err != nil {
		http_util.NewError(c, http.StatusBadRequest, err)
		return
	}

	// Responder com o usuário
	c.JSON(http.StatusOK, user)
}

// GetAllUsers godoc
//
//	@Summary		Retorna todos os usuários
//	@Description	Retorna todos os usuários
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	user_models.User
//	@Failure		400	{object}	http_util.HTTPError
//	@Router			/users [get]
func (controller *UserControllerImpl) getAllUsers(c *gin.Context) {
	users, err := controller.userService.FindAll()
	if err != nil {
		http_util.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

type CreateUserRequest struct {
	Name     *string `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

// CreateUser godoc
//
//	@Summary		Cria um usuário
//	@Description	Recurso de criação
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		user_controller.CreateUserRequest	true	"Add User"
//	@Failure		400	{object}	http_util.HTTPError
//	@Router			/users [post]
func (controller *UserControllerImpl) createUser(c *gin.Context) {
	var request CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_util.NewError(c, http.StatusBadRequest, http.ErrBodyNotAllowed)
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
		http_util.NewError(c, http.StatusBadRequest, http.ErrAbortHandler)
		return
	}

	// Responder com o ID do usuário
	c.JSON(http.StatusOK, gin.H{"id": userID})
}

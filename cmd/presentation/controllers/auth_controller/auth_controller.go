// presentation/controller/auth/auth_controller.go - Camada de Apresentação (Presentation Layer)
package auth_controller

import (
	"errors"
	"net/http"

	"github.com/RafaelCava/kitkit-back-go/cmd/domain/usecases/auth_usecase"
	"github.com/RafaelCava/kitkit-back-go/cmd/presentation/utils/http_util"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	RegisterRoutes(router *gin.RouterGroup)
	login(c *gin.Context)
}

type AuthControllerImpl struct {
	authService auth_usecase.AuthService
}

func NewAuthControllerImpl(authService auth_usecase.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{authService}
}

func (controller *AuthControllerImpl) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/login", controller.login)
}

// Login godoc
//
//		@Summary		  Autentica um usuário
//		@Description  Autentica um usuário pela senha
//		@Tags			    Auth
//		@Accept			  json
//		@Produce		  json
//	  @Param			  auth	body		auth_usecase.AuthLoginRequest	true	"Authenticate"
//	  @Success		  200	{string}	string	"Autorizado"
//	  @Failure		  400	{object}	http_util.HTTPError
//		@Router			  /auth/login [post]
func (controller *AuthControllerImpl) login(c *gin.Context) {
	authRequest := auth_usecase.AuthLoginRequest{}
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		http_util.NewError(c, http.StatusBadRequest, errors.New("invalid request"))
		return
	}

	isAuthenticated, err := controller.authService.Validate(&authRequest)
	if err != nil {
		http_util.NewError(c, http.StatusBadRequest, err)
		return
	}
	if !isAuthenticated {
		http_util.NewError(c, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	} else {
		token, err := controller.authService.GenerateToken(map[string]string{"email": authRequest.Email})
		if err != nil {
			http_util.NewError(c, http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

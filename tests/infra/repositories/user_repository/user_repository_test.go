// infra/repositories/user_repository/user_repository_test.go
package user_repository_test

import (
	"testing"

	"github.com/RafaelCava/kitkit-back-go/cmd/domain/models/user_models"
	"github.com/RafaelCava/kitkit-back-go/cmd/infra/repositories/user_repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Erro ao configurar o banco de dados de teste")
	}
	db.AutoMigrate(&user_models.User{})
	return db
}

func TestGormUserRepository_FindByID(t *testing.T) {
	// Configurando o banco de dados de teste
	db := setupTestDB()
	// Criando instância de GormUserRepository
	repo := user_repository.NewGormUserRepository(db)

	// Inserindo um usuário de teste no banco de dados
	testUserID := "1"
	testUserName := "Test User"
	testUserEmail := "test@example.com"
	testUserPassword := "password"
	testUser := &user_models.User{
		ID:       testUserID,
		Name:     &testUserName,
		Email:    testUserEmail,
		Password: testUserPassword,
	}
	db.Create(testUser)

	// Testando a função FindByID
	resultUser, err := repo.FindByID(testUserID)

	// Verificando se não ocorreu erro
	assert.Nil(t, err, "Erro inesperado")

	// Verificando se o usuário retornado é o esperado
	assert.NotNil(t, resultUser, "Usuário não encontrado")
	assert.Equal(t, testUserID, resultUser.ID, "IDs de usuário não correspondem")
	assert.Equal(t, testUserName, *resultUser.Name, "Nomes de usuário não correspondem")
	assert.Equal(t, testUserEmail, resultUser.Email, "E-mails de usuário não correspondem")
	assert.Equal(t, testUserPassword, resultUser.Password, "Senhas de usuário não correspondem")
}

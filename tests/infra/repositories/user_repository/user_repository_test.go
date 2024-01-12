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

var db *gorm.DB

func setupTestDB() {
	data, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Erro ao configurar o banco de dados de teste")
	}
	data.AutoMigrate(&user_models.User{})
	db = data
}

func TestMain(m *testing.M) {
	setupTestDB()
	m.Run()
}

func clearUserTable() {
	db.Exec("DELETE FROM users")
}

func makeSut() *user_repository.GormUserRepository {
	return user_repository.NewGormUserRepository(db)
}

func TestGormUserRepository_FindByID(t *testing.T) {
	clearUserTable()
	repo := makeSut()

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

func TestGormUserRepository_ReturnError(t *testing.T) {
	clearUserTable()
	repo := makeSut()

	testUserID := "1"

	// Testando a função FindByID
	resultUser, err := repo.FindByID(testUserID)

	// Verificando se ocorreu erro
	assert.NotNil(t, err, "Erro esperado")
	assert.Nil(t, resultUser, "Usuário encontrado")
}

func TestGormUserRepository_Create(t *testing.T) {
	clearUserTable()
	repo := makeSut()
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

	userID, err := repo.Create(testUser)

	assert.Nil(t, err, "Erro inesperado")
	assert.Equal(t, testUserID, userID, "IDs de usuário não correspondem")
}

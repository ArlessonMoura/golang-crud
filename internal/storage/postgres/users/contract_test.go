package users

import (
	"testing"

	"meu-treino-golang/users-crud/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryImplementsPort(t *testing.T) {
	var _ service.IUserRepository = (*Repository)(nil)
}

func TestRepositoryInstantiation(t *testing.T) {
	repo := NewRepository(nil)
	assert.NotNil(t, repo)
}

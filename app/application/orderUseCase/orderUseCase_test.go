package orderUseCase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockStructure struct {
	Status string `json:status`
}

func Test_ValidateInDatabase(t *testing.T) {
	t.Parallel()
	t.Run("Validate saveInDatabase", func(t *testing.T) {
		assert.Equal(t, SaveInDatabase(100, "Matias"), true)
	})
}

func Test_OrderUseCase(t *testing.T) {
	t.Parallel()
	t.Run("OrderUseCase", func(t *testing.T) {

		responseEstado, responseError := OrderUseCase()
		assert.Equal(t, responseEstado, true)
		assert.Equal(t, responseError, nil)
	})
}

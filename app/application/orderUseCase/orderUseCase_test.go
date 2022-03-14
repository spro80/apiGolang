package orderUseCase

import (
	"fmt"
	"testing"
)

type mockStructure struct {
	Status string `json:status`
}

/*
func Test_ValidateInDatabase(t *testing.T) {
	t.Parallel()

	var resStatus, resError = SaveInDatabase(100, "Matias")

	t.Run("Validate saveInDatabase", func(t *testing.T) {
		assert.Equal(t, resStatus, true)
		assert.Equal(t, resError, nil)
	})
}*/

func Test_OrderUseCase(t *testing.T) {
	t.Parallel()
	t.Run("OrderUseCase", func(t *testing.T) {

		response := NewOrderUseCase()
		res, err := response.Process()
		fmt.Println(response)
		fmt.Println(res)
		fmt.Println(err)

		//assert.Equal(t, nil, nil)
		/*responseEstado, responseError := OrderUseCase()
		assert.Equal(t, responseEstado, true)
		assert.Equal(t, responseError, nil)
		*/
	})
}

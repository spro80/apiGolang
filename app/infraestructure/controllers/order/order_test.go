package order

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockStructure struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Test_ConvertEntity(t *testing.T) {

	t.Parallel()

	t.Run("When an entity of a specific type is delivered and another of another type returns with the same structure ok", func(t *testing.T) {
		/*e1 := mockStructure{
			Id:   1,
			Name: "name",
		}*/
		assert.Equal(t, CalculateSum(195, 4), 199)
	})

}

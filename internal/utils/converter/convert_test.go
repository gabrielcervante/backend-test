package converter

import (
	"testing"

	"github.com/japhy-tech/backend-test/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestConverter(t *testing.T) {
	breedsDB := mocks.NewBreedsDB(t)

	converterStr := converter{BreedsDB: breedsDB}

	t.Run("Test Read File", func(t *testing.T) {
		assert.NotNil(t, converterStr.readFile("../../../breeds.csv"))
	})
}

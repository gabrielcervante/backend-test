package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreeds(t *testing.T) {
	t.Run("Test Breeds Sucess", func(t *testing.T) {
		breeds := Breeds{
			ID:                       1,
			Species:                  "species",
			PetSize:                  "pet size",
			Name:                     "name",
			Weight:                   1,
			AverageMaleAdultWeight:   1,
			AverageFemaleAdultWeight: 1,
		}

		assert.Nil(t, breeds.Validate())
	})

	t.Run("Test Breeds Error", func(t *testing.T) {
		breeds := Breeds{
			ID:                       0,
			Species:                  "species",
			PetSize:                  "pet size",
			Name:                     "name",
			Weight:                   1,
			AverageMaleAdultWeight:   1,
			AverageFemaleAdultWeight: 1,
		}

		assert.NotNil(t, breeds)
	})
}

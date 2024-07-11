package breeds

import (
	"context"
	"testing"

	"github.com/japhy-tech/backend-test/internal/entity"
	"github.com/japhy-tech/backend-test/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBreedsCRUD(t *testing.T) {
	breedsEntity := entity.Breeds{
		ID:                       1,
		Species:                  "species",
		PetSize:                  "pet size",
		Name:                     "name",
		Weight:                   1,
		AverageMaleAdultWeight:   1,
		AverageFemaleAdultWeight: 1,
	}

	breedsRepo := mocks.NewBreedsDB(t)
	breedsCRUD := NewBreeds(breedsRepo)

	t.Run("Test Create", func(t *testing.T) {

		breedsRepo.On("Insert", mock.Anything, breedsEntity).Return(nil).Once()

		err := breedsCRUD.Create(context.Background(), breedsEntity)

		assert.Nil(t, err)
	})

	t.Run("Test Read By ID", func(t *testing.T) {

		breedsRepo.On("SelectById", mock.Anything, 1).Return([]entity.Breeds{breedsEntity}, nil).Once()

		resp, err := breedsCRUD.Read(context.Background(), "id", "1")

		assert.Equal(t, []entity.Breeds{breedsEntity}, resp)
		assert.Nil(t, err)
	})

	t.Run("Test Read By Name", func(t *testing.T) {

		breedsRepo.On("SelectByName", mock.Anything, "raj").Return([]entity.Breeds{breedsEntity}, nil).Once()

		resp, err := breedsCRUD.Read(context.Background(), "name", "raj")

		assert.Equal(t, []entity.Breeds{breedsEntity}, resp)
		assert.Nil(t, err)
	})

	t.Run("Test Read By Pet Size", func(t *testing.T) {

		breedsRepo.On("SelectByPetSize", mock.Anything, "1").Return([]entity.Breeds{breedsEntity}, nil).Once()

		resp, err := breedsCRUD.Read(context.Background(), "pet_size", "1")

		assert.Equal(t, []entity.Breeds{breedsEntity}, resp)
		assert.Nil(t, err)
	})

	t.Run("Test Read By Species", func(t *testing.T) {

		breedsRepo.On("SelectBySpecies", mock.Anything, "cat").Return([]entity.Breeds{breedsEntity}, nil).Once()

		resp, err := breedsCRUD.Read(context.Background(), "species", "cat")

		assert.Equal(t, []entity.Breeds{breedsEntity}, resp)
		assert.Nil(t, err)
	})

	t.Run("Test Read By Weight", func(t *testing.T) {

		breedsRepo.On("SelectByWeight", mock.Anything, 1).Return([]entity.Breeds{breedsEntity}, nil).Once()

		resp, err := breedsCRUD.Read(context.Background(), "weight", "1")

		assert.Equal(t, []entity.Breeds{breedsEntity}, resp)
		assert.Nil(t, err)
	})

	t.Run("Test Update", func(t *testing.T) {

		breedsRepo.On("Update", mock.Anything, breedsEntity).Return(nil).Once()

		err := breedsCRUD.Update(context.Background(), breedsEntity)

		assert.Nil(t, err)
	})

	t.Run("Test Delete", func(t *testing.T) {

		breedsRepo.On("Delete", mock.Anything, 1).Return(nil).Once()

		err := breedsCRUD.Delete(context.Background(), 1)

		assert.Nil(t, err)
	})
}

package breeds

import (
	"context"
	"errors"
	"strconv"

	"github.com/japhy-tech/backend-test/internal/entity"
	"github.com/japhy-tech/backend-test/internal/ports"
)

var readFilters = map[string]bool{
	"id":       true,
	"name":     false,
	"species":  false,
	"pet_size": false,
	"weight":   true,
}

type breeds struct {
	db ports.BreedsDB
}

func (b breeds) Create(ctx context.Context, breeds entity.Breeds) error {
	if err := breeds.Validate(); err != nil {
		return err
	}

	return b.db.Insert(ctx, breeds)
}

func (b breeds) Read(ctx context.Context, filterName string, input string) ([]entity.Breeds, error) {
	if input == "" {
		return nil, errors.New("sorry, your query is empty")
	}

	switch filterName {
	case "id":
		convertedID, err := strconv.Atoi(input)
		if err != nil {
			return nil, err
		}

		return b.db.SelectById(ctx, convertedID)

	case "name":
		return b.db.SelectByName(ctx, input)
	case "species":
		return b.db.SelectBySpecies(ctx, input)
	case "pet_size":
		return b.db.SelectByPetSize(ctx, input)

	case "weight":
		convertedWeight, err := strconv.Atoi(input)
		if err != nil {
			return nil, err
		}

		return b.db.SelectByWeight(ctx, convertedWeight)
	}

	return nil, errors.New("you didn't provided any filter")
}

func (b breeds) Update(ctx context.Context, breeds entity.Breeds) error {
	if err := breeds.Validate(); err != nil {
		return err
	}

	return b.db.Update(ctx, breeds)
}

func (b breeds) Delete(ctx context.Context, id int) error {
	if id == 0 {
		return errors.New("sorry, you didn't provided the id")
	}

	return b.db.Delete(ctx, id)
}

func NewBreeds(repo ports.BreedsDB) ports.BreedsCRUD {
	return breeds{repo}
}

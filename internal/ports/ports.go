package ports

import (
	"context"

	"github.com/japhy-tech/backend-test/internal/entity"
)

type BreedsCRUD interface {
	Create(ctx context.Context, breeds entity.Breeds) error
	Read(ctx context.Context, filtername, input string) ([]entity.Breeds, error)
	Update(ctx context.Context, breeds entity.Breeds) error
	Delete(ctx context.Context, id int) error
}

type BreedsDB interface {
	Insert(context.Context, entity.Breeds) error
	SelectById(ctx context.Context, id int) (breeds []entity.Breeds, err error)
	SelectBySpecies(ctx context.Context, species string) (breeds []entity.Breeds, err error)
	SelectByPetSize(ctx context.Context, petSize string) (breeds []entity.Breeds, err error)
	SelectByWeight(ctx context.Context, weight int) (breeds []entity.Breeds, err error)
	SelectByName(ctx context.Context, name string) (breeds []entity.Breeds, err error)
	Update(ctx context.Context, breeds entity.Breeds) (err error)
	Delete(ctx context.Context, id int) (err error)
}

type Converter interface {
	ConvertCSVToSQL(context.Context, string)
}

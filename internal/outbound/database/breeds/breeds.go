package database

import (
	"context"
	"database/sql"

	"github.com/japhy-tech/backend-test/internal/entity"
	"github.com/japhy-tech/backend-test/internal/outbound/database/repo"
	"github.com/japhy-tech/backend-test/internal/ports"
)

type breeds struct {
	*sql.DB
}

func (b breeds) Insert(ctx context.Context, breeds entity.Breeds) (err error) {
	_, err = b.ExecContext(ctx, insert, breedsInputInsert(breeds)...)
	return
}

func (b breeds) SelectById(ctx context.Context, id int) (breeds []entity.Breeds, err error) {
	rows, err := b.QueryContext(ctx, SelectById, id)
	return repo.Scan(rows, err, breedsOutput)
}

func (b breeds) SelectBySpecies(ctx context.Context, species string) (breeds []entity.Breeds, err error) {
	rows, err := b.QueryContext(ctx, SelectBySpecies, species)
	return repo.Scan(rows, err, breedsOutput)
}

func (b breeds) SelectByPetSize(ctx context.Context, petSize string) (breeds []entity.Breeds, err error) {
	rows, err := b.QueryContext(ctx, SelectByPetSize, petSize)
	return repo.Scan(rows, err, breedsOutput)
}

func (b breeds) SelectByWeight(ctx context.Context, weight int) (breeds []entity.Breeds, err error) {
	rows, err := b.QueryContext(ctx, SelectByWeight, weight)
	return repo.Scan(rows, err, breedsOutput)
}

func (b breeds) SelectByName(ctx context.Context, name string) (breeds []entity.Breeds, err error) {
	rows, err := b.QueryContext(ctx, SelectByName, name)
	return repo.Scan(rows, err, breedsOutput)
}

func (b breeds) Update(ctx context.Context, breeds entity.Breeds) (err error) {
	_, err = b.ExecContext(ctx, update, breedsInputUpdate(breeds)...)
	return
}

func (b breeds) Delete(ctx context.Context, id int) (err error) {
	_, err = b.ExecContext(ctx, delete, id)
	return
}

func NewBreeds(db *sql.DB) ports.BreedsDB {
	return breeds{db}
}

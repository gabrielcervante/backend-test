package database

import "github.com/japhy-tech/backend-test/internal/entity"

func breedsInputInsert(breeds entity.Breeds) []any {
	return []any{
		breeds.ID,
		breeds.Species,
		breeds.PetSize,
		breeds.Name,
		breeds.Weight,
		breeds.AverageMaleAdultWeight,
		breeds.AverageFemaleAdultWeight,
	}
}

func breedsInputUpdate(breeds entity.Breeds) []any {
	return []any{
		breeds.Species,
		breeds.PetSize,
		breeds.Name,
		breeds.Weight,
		breeds.AverageMaleAdultWeight,
		breeds.AverageFemaleAdultWeight,
		breeds.ID,
	}
}

func breedsOutput(breeds *entity.Breeds) []any {
	return []any{
		&breeds.ID,
		&breeds.Species,
		&breeds.PetSize,
		&breeds.Name,
		&breeds.Weight,
		&breeds.AverageMaleAdultWeight,
		&breeds.AverageFemaleAdultWeight,
	}
}

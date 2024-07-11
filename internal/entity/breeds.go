package entity

import (
	"errors"

	"github.com/japhy-tech/backend-test/internal/utils/validator"
)

type Breeds struct {
	ID                       int    `csv:"id" json:"id"`
	Species                  string `csv:"species" json:"species"`
	PetSize                  string `csv:"pet_size" json:"pet_size"`
	Name                     string `csv:"name" json:"name"`
	Weight                   int    `csv:"weight" json:"weight"`
	AverageMaleAdultWeight   int    `csv:"average_male_adult_weight" json:"average_male_adult_weight"`
	AverageFemaleAdultWeight int    `csv:"average_female_adult_weight" json:"average_female_adult_weight"`
}

func (b Breeds) Validate() error {
	if isEmpty := validator.IsEmpty(b.ID); isEmpty {
		return errors.New("sorry, you didn't provide the id")
	}

	if isEmpty := validator.IsEmpty(b.Species); isEmpty {
		return errors.New("sorry, you didn't provide the specie")
	}

	if isEmpty := validator.IsEmpty(b.PetSize); isEmpty {
		return errors.New("sorry, you didn't provide the pet size")
	}

	if isEmpty := validator.IsEmpty(b.Name); isEmpty {
		return errors.New("sorry, you didn't provide the name")
	}

	if isEmpty := validator.IsEmpty(b.Weight); isEmpty {
		return errors.New("sorry, you didn't provide the weight")
	}

	return nil
}

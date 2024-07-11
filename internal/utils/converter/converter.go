package converter

import (
	"context"
	"os"
	"strings"

	"github.com/japhy-tech/backend-test/internal/entity"
	"github.com/japhy-tech/backend-test/internal/ports"

	"github.com/gocarina/gocsv"
)

type converter struct {
	ports.BreedsDB
}

func (c converter) ConvertCSVToSQL(ctx context.Context, filePath string) {
	breeds := c.readFile(filePath)

	c.save(ctx, breeds)
}

func (c converter) readFile(filePath string) []entity.Breeds {
	var breeds []entity.Breeds

	breedsFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer breedsFile.Close()

	if err := gocsv.UnmarshalFile(breedsFile, &breeds); err != nil {
		panic(err)
	}

	return breeds
}

func (c converter) save(ctx context.Context, breeds []entity.Breeds) {
	for i := 0; i < len(breeds); i++ {
		if err := c.Insert(ctx, breeds[i]); err != nil && !strings.Contains(err.Error(), "Duplicate entry") {
			panic(err)
		}
	}
}

func NewConverter(db ports.BreedsDB) ports.Converter {
	return converter{db}
}

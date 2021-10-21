package repository

import (
	"bufio"
	"os"
	"path"

	"github.com/mteixidorc/trips/internal/shared"
	"github.com/mteixidorc/trips/internal/shared/domain/errors"
	"github.com/mteixidorc/trips/internal/shared/domain/value"
	"github.com/mteixidorc/trips/internal/trips/domain"
)

const defaultFileCitiesDataPath = "/data/cities.txt"

// FileCityRepository
// Implements CityRepository interface and reads cities records from a file
type FileCityRepository struct {
	fileData map[value.ValueObjectID]string
}

func NewFileCityRepository() (*FileCityRepository, error) {
	cities := map[value.ValueObjectID]string{}

	fileCitiesPath := path.Join(shared.Root, defaultFileCitiesDataPath)
	f, err := os.Open(fileCitiesPath)

	if err != nil {
		// TODO build a better error message
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var cityId int64 = 1
	for scanner.Scan() {
		cities[value.NewValueObjectID(cityId)] = scanner.Text()
		cityId++
	}

	return &FileCityRepository{
		fileData: cities,
	}, nil
}

func (repository FileCityRepository) Get(id value.ValueObjectID) (*domain.City, error) {
	city, exists := repository.fileData[id]
	if !exists {
		return nil, errors.NewRecordNotExistsError(id)
	}

	return domain.NewCity(id, city), nil
}

package storage

import (
	"errors"
	"fmt"
	"net/http"

	"go-jsonapi-example/internal/model"

	"github.com/manyminds/api2go"
)

type CarStorage struct {
	cars    map[string]*model.Car
	idCount int
}

func NewCarStorage() *CarStorage {
	return &CarStorage{make(map[string]*model.Car), 1}
}

func (s CarStorage) GetAll() map[string]*model.Car {
	return s.cars
}

func (s CarStorage) GetOne(id string) (model.Car, error) {
	user, ok := s.cars[id]
	if ok {
		return *user, nil
	}
	errMessage := fmt.Sprintf("Car for id %s not found", id)
	return model.Car{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (s *CarStorage) Insert(c model.Car) string {
	id := fmt.Sprintf("%d", s.idCount)
	c.ID = id
	s.cars[id] = &c
	s.idCount++
	return id
}

func (s *CarStorage) Delete(id string) error {
	_, exists := s.cars[id]
	if !exists {
		return fmt.Errorf("Car with id %s does not exist", id)
	}
	delete(s.cars, id)

	return nil
}

func (s *CarStorage) Update(c model.Car) error {
	_, exists := s.cars[c.ID]
	if !exists {
		return fmt.Errorf("Car with id %s does not exist", c.ID)
	}
	s.cars[c.ID] = &c

	return nil
}

package storage

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"go-jsonapi-example/internal/model"

	"github.com/manyminds/api2go"
)

type CarStorage struct {
	mutex   sync.RWMutex
	cars    map[uint64]*model.Car
	idCount uint64
}

type Cars []model.Car

func (c Cars) Len() int           { return len(c) }
func (c Cars) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Cars) Less(i, j int) bool { return c[i].ID < c[j].ID }

func NewCarStorage() *CarStorage {
	return &CarStorage{cars: make(map[uint64]*model.Car), idCount: 1}
}

func (s CarStorage) GetAll() Cars {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	result := make(Cars, 0, len(s.cars))
	for _, car := range s.cars {
		result = append(result, *car)
	}
	return result
}

func (s CarStorage) GetOne(id uint64) (model.Car, error) {
	s.mutex.RLock()
	car, ok := s.cars[id]
	if ok {
		defer s.mutex.RUnlock()
		return *car, nil
	}
	s.mutex.RUnlock()
	errMessage := fmt.Sprintf("Car for id %s not found", id)
	return model.Car{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (s *CarStorage) Insert(c model.Car) uint64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	c.ID = s.idCount
	s.cars[s.idCount] = &c
	s.idCount++
	return c.ID
}

func (s *CarStorage) Delete(id uint64) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, exists := s.cars[id]
	if !exists {
		return fmt.Errorf("Car with id %s does not exist", id)
	}
	delete(s.cars, id)

	return nil
}

func (s *CarStorage) Update(c model.Car) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, exists := s.cars[c.ID]
	if !exists {
		return fmt.Errorf("Car with id %s does not exist", c.ID)
	}
	s.cars[c.ID] = &c

	return nil
}

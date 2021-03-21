package resource

import (
	"errors"
	"net/http"
	"sort"
	"strconv"

	"go-jsonapi-example/internal/model"
	"go-jsonapi-example/internal/storage"

	"github.com/manyminds/api2go"
)

// CarResource for api2go routes
type CarResource struct {
	CarStorage *storage.CarStorage
}

// FindAll to satisfy api2go data source interface
func (s CarResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: s.CarStorage.GetAll()}, nil
}

// PaginatedFindAll can be used to load cars in chunks
func (s CarResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      storage.Cars
		number, size, offset, limit string
	)
	cars := s.CarStorage.GetAll()
	sort.Sort(cars)

	numberQuery, ok := r.QueryParams["page[number]"]
	if ok {
		number = numberQuery[0]
	}
	sizeQuery, ok := r.QueryParams["page[size]"]
	if ok {
		size = sizeQuery[0]
	}
	offsetQuery, ok := r.QueryParams["page[offset]"]
	if ok {
		offset = offsetQuery[0]
	}
	limitQuery, ok := r.QueryParams["page[limit]"]
	if ok {
		limit = limitQuery[0]
	}

	if size != "" {
		sizeI, err := strconv.ParseUint(size, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		numberI, err := strconv.ParseUint(number, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		start := sizeI * (numberI - 1)
		for i := start; i < start+sizeI; i++ {
			if i >= uint64(len(cars)) {
				break
			}
			result = append(result, cars[i])
		}
	} else {
		limitI, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		offsetI, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		for i := offsetI; i < offsetI+limitI; i++ {
			if i >= uint64(len(cars)) {
				break
			}
			result = append(result, cars[i])
		}
	}

	return uint(len(cars)), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
func (s CarResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	intID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	car, err := s.CarStorage.GetOne(intID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	return &Response{Res: car}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s CarResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	car, ok := obj.(model.Car)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	if ok, httpErr := car.Verify(); !ok {
		return &Response{}, httpErr
	}

	id := s.CarStorage.Insert(car)
	car.ID = id

	return &Response{Res: car, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s CarResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	intID, err := strconv.ParseUint(id, 10, 64)
	if err == nil {
		err = s.CarStorage.Delete(intID)
	}
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the car
func (s CarResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	car, ok := obj.(model.Car)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	if ok, httpErr := car.Verify(); !ok {
		return &Response{}, httpErr
	}

	err := s.CarStorage.Update(car)
	return &Response{Code: http.StatusOK}, err
}

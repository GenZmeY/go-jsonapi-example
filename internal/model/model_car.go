package model

import (
	"errors"

	"fmt"
	"net/http"
	"strconv"

	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go/jsonapi"
)

type Car struct {
	ID      uint64 `json:"-"`
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Price   uint64 `json:"price"`
	Status  string `json:"status"`  // OnTheWay, InStock, Sold, Discontinued
	Mileage int64  `json:"mileage"` // I suppose it should be made unsigned, but that's what the task says ¯\_(ツ)_/¯
}

func (c Car) Verify() (bool, api2go.HTTPError) {
	var verifyErrors []api2go.Error
	var httpError api2go.HTTPError

	if c.Brand == "" {
		newErr := newVerifyError(
			"Invalid Attribute",
			"attribute cannot be empty",
			"/data/attributes/brand")
		verifyErrors = append(verifyErrors, newErr)
	}
	if c.Model == "" {
		newErr := newVerifyError(
			"Invalid Attribute",
			"attribute cannot be empty",
			"/data/attributes/model")
		verifyErrors = append(verifyErrors, newErr)
	}
	if c.Status != "OnTheWay" &&
		c.Status != "InStock" &&
		c.Status != "Sold" &&
		c.Status != "Discontinued" {
		newErr := newVerifyError(
			"Invalid Attribute",
			"attribute must be one of: OnTheWay, InStock, Sold, Discontinued",
			"/data/attributes/status")
		verifyErrors = append(verifyErrors, newErr)
	}

	ok := len(verifyErrors) == 0

	if !ok {
		httpError = api2go.NewHTTPError(
			errors.New("Invalid content"),
			"Invalid content",
			http.StatusBadRequest)
		httpError.Errors = verifyErrors
	}

	return ok, httpError
}

func newVerifyError(title string, detail string, pointer string) api2go.Error {
	var newError api2go.Error
	newError.Title = title
	newError.Detail = detail
	newError.Source = &api2go.ErrorSource{Pointer: pointer}
	return newError
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Car) GetID() string {
	return fmt.Sprintf("%d", c.ID)
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Car) SetID(id string) error {
	if id == "" {
		c.ID = 0
		return nil
	}

	intID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	c.ID = intID
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (c Car) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (c Car) GetReferencedIDs() []jsonapi.ReferenceID {
	return []jsonapi.ReferenceID{}
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (c Car) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	return []jsonapi.MarshalIdentifier{}
}

// SetToManyReferenceIDs sets the sweets reference IDs and satisfies the jsonapi.UnmarshalToManyRelations interface
func (c *Car) SetToManyReferenceIDs(name string, IDs []string) error {
	return errors.New("There is no to-many relationship with the name " + name)
}

// AddToManyIDs adds some new sweets that a users loves so much
func (c *Car) AddToManyIDs(name string, IDs []string) error {
	return errors.New("There is no to-many relationship with the name " + name)
}

// DeleteToManyIDs removes some sweets from a users because they made him very sick
func (c *Car) DeleteToManyIDs(name string, IDs []string) error {
	return errors.New("There is no to-many relationship with the name " + name)
}

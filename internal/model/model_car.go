package model

import (
	"errors"

	"github.com/manyminds/api2go/jsonapi"
)

type Car struct {
	ID     string `json:"-"`
	Brand  string `json:"brand"`
	Model  string `json:"model"`
	Price  uint   `json:"price"`
	Status string `json:"status"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u Car) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *Car) SetID(id string) error {
	u.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u Car) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u Car) GetReferencedIDs() []jsonapi.ReferenceID {
	return []jsonapi.ReferenceID{}
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u Car) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	return []jsonapi.MarshalIdentifier{}
}

// SetToManyReferenceIDs sets the sweets reference IDs and satisfies the jsonapi.UnmarshalToManyRelations interface
func (u *Car) SetToManyReferenceIDs(name string, IDs []string) error {
	return errors.New("There is no to-many relationship with the name " + name)
}

// AddToManyIDs adds some new sweets that a users loves so much
func (u *Car) AddToManyIDs(name string, IDs []string) error {
	return errors.New("There is no to-many relationship with the name " + name)
}

// DeleteToManyIDs removes some sweets from a users because they made him very sick
func (u *Car) DeleteToManyIDs(name string, IDs []string) error {
	return errors.New("There is no to-many relationship with the name " + name)
}

// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook_test

import (
	"github.com/tmc/appledocs/generated/addressbook"
)

// Suppress unused import errors
var _ = addressbook.NewABPerson

// ExampleABPerson_ImageData demonstrates using ImageData on a ABPerson instance.
// Returns data that contains a picture of this person.
//
// Note: This example is not executed because ImageData crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleABPerson_ImageData() {
	obj := addressbook.NewABPerson()
	_ = obj.ImageData()
	}

// ExampleABPerson_LinkedPeople demonstrates using LinkedPeople on a ABPerson instance.
// Returns the array of all person records that are linked to the person this record represents.
func ExampleABPerson_LinkedPeople() {
	obj := addressbook.NewABPerson()
	_ = obj.LinkedPeople()
	// Output:
	}

// ExampleABPerson_ParentGroups demonstrates using ParentGroups on a ABPerson instance.
// Returns an array of the address book groups that this person belongs to.
func ExampleABPerson_ParentGroups() {
	obj := addressbook.NewABPerson()
	_ = obj.ParentGroups()
	// Output:
	}

// ExampleABPerson_VCardRepresentation demonstrates using VCardRepresentation on a ABPerson instance.
// Returns the vCard representation of the person record as a data object in vCard format.
func ExampleABPerson_VCardRepresentation() {
	obj := addressbook.NewABPerson()
	_ = obj.VCardRepresentation()
	// Output:
	}


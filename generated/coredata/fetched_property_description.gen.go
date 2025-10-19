// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FetchedPropertyDescription] class.
var fetchedPropertyDescriptionClass = _FetchedPropertyDescriptionClass{objc.GetClass("NSFetchedPropertyDescription")}

type _FetchedPropertyDescriptionClass struct {
	class objc.Class
}

// A description object used to define which properties are fetched from Core Data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedPropertyDescription

type FetchedPropertyDescription struct {
	PropertyDescription
}

// FetchedPropertyDescriptionFrom constructs a [FetchedPropertyDescription] from an unsafe.Pointer.
//
// A description object used to define which properties are fetched from Core Data.
func FetchedPropertyDescriptionFrom(ptr unsafe.Pointer) FetchedPropertyDescription {
	return FetchedPropertyDescription{
		PropertyDescription: PropertyDescriptionFrom(ptr),
	}
}




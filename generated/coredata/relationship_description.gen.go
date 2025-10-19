// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RelationshipDescription] class.
var relationshipDescriptionClass = _RelationshipDescriptionClass{objc.GetClass("NSRelationshipDescription")}

type _RelationshipDescriptionClass struct {
	class objc.Class
}

// A description of a relationship between two entities. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSRelationshipDescription

type RelationshipDescription struct {
	PropertyDescription
}

// RelationshipDescriptionFrom constructs a [RelationshipDescription] from an unsafe.Pointer.
//
// A description of a relationship between two entities.
func RelationshipDescriptionFrom(ptr unsafe.Pointer) RelationshipDescription {
	return RelationshipDescription{
		PropertyDescription: PropertyDescriptionFrom(ptr),
	}
}




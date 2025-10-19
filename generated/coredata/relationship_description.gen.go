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

// An interface definition for the [RelationshipDescription] class.
type IRelationshipDescription interface {
	IPropertyDescription
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
// Alloc allocates a new instance without initialization.
func (rc _RelationshipDescriptionClass) Alloc() RelationshipDescription {
	rv := objc.Send[RelationshipDescription](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (rc _RelationshipDescriptionClass) New() RelationshipDescription {
	rv := objc.Send[RelationshipDescription](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RelationshipDescription) Init() RelationshipDescription {
	rv := objc.Send[RelationshipDescription](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RelationshipDescription) Autorelease() RelationshipDescription {
	rv := objc.Send[RelationshipDescription](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRelationshipDescription creates a new RelationshipDescription instance.
func NewRelationshipDescription() RelationshipDescription {
	return relationshipDescriptionClass.New()
}





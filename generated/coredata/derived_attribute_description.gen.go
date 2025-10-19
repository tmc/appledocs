// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DerivedAttributeDescription] class.
var derivedAttributeDescriptionClass = _DerivedAttributeDescriptionClass{objc.GetClass("NSDerivedAttributeDescription")}

type _DerivedAttributeDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [DerivedAttributeDescription] class.
type IDerivedAttributeDescription interface {
	IAttributeDescription
}

// A description of an attribute that derives its value by performing a calculation on a related attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSDerivedAttributeDescription

type DerivedAttributeDescription struct {
	AttributeDescription
}

// DerivedAttributeDescriptionFrom constructs a [DerivedAttributeDescription] from an unsafe.Pointer.
//
// A description of an attribute that derives its value by performing a calculation on a related attribute.
func DerivedAttributeDescriptionFrom(ptr unsafe.Pointer) DerivedAttributeDescription {
	return DerivedAttributeDescription{
		AttributeDescription: AttributeDescriptionFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (dc _DerivedAttributeDescriptionClass) Alloc() DerivedAttributeDescription {
	rv := objc.Send[DerivedAttributeDescription](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (dc _DerivedAttributeDescriptionClass) New() DerivedAttributeDescription {
	rv := objc.Send[DerivedAttributeDescription](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DerivedAttributeDescription) Init() DerivedAttributeDescription {
	rv := objc.Send[DerivedAttributeDescription](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DerivedAttributeDescription) Autorelease() DerivedAttributeDescription {
	rv := objc.Send[DerivedAttributeDescription](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDerivedAttributeDescription creates a new DerivedAttributeDescription instance.
func NewDerivedAttributeDescription() DerivedAttributeDescription {
	return derivedAttributeDescriptionClass.New()
}





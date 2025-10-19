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




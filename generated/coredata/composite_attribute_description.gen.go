// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CompositeAttributeDescription] class.
var compositeAttributeDescriptionClass = _CompositeAttributeDescriptionClass{objc.GetClass("NSCompositeAttributeDescription")}

type _CompositeAttributeDescriptionClass struct {
	class objc.Class
}

// A description of an attribute that derives its value by composing other attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCompositeAttributeDescription

type CompositeAttributeDescription struct {
	AttributeDescription
}

// CompositeAttributeDescriptionFrom constructs a [CompositeAttributeDescription] from an unsafe.Pointer.
//
// A description of an attribute that derives its value by composing other attributes.
func CompositeAttributeDescriptionFrom(ptr unsafe.Pointer) CompositeAttributeDescription {
	return CompositeAttributeDescription{
		AttributeDescription: AttributeDescriptionFrom(ptr),
	}
}




// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AttributeDescription] class.
var attributeDescriptionClass = _AttributeDescriptionClass{objc.GetClass("NSAttributeDescription")}

type _AttributeDescriptionClass struct {
	class objc.Class
}

// A description of a single attribute belonging to an entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAttributeDescription

type AttributeDescription struct {
	PropertyDescription
}

// AttributeDescriptionFrom constructs a [AttributeDescription] from an unsafe.Pointer.
//
// A description of a single attribute belonging to an entity.
func AttributeDescriptionFrom(ptr unsafe.Pointer) AttributeDescription {
	return AttributeDescription{
		PropertyDescription: PropertyDescriptionFrom(ptr),
	}
}




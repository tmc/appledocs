// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PropertyDescription] class.
var propertyDescriptionClass = _PropertyDescriptionClass{objc.GetClass("NSPropertyDescription")}

type _PropertyDescriptionClass struct {
	class objc.Class
}

// A description of a single property belonging to an entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription

type PropertyDescription struct {
	objectivec.Object
}

// PropertyDescriptionFrom constructs a [PropertyDescription] from an unsafe.Pointer.
//
// A description of a single property belonging to an entity.
func PropertyDescriptionFrom(ptr unsafe.Pointer) PropertyDescription {
	return PropertyDescription{objectivec.Object{objc.ID(ptr)}}
}




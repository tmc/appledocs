// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EntityDescription] class.
var entityDescriptionClass = _EntityDescriptionClass{objc.GetClass("NSEntityDescription")}

type _EntityDescriptionClass struct {
	class objc.Class
}

// A description of a Core Data entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityDescription

type EntityDescription struct {
	objectivec.Object
}

// EntityDescriptionFrom constructs a [EntityDescription] from an unsafe.Pointer.
//
// A description of a Core Data entity.
func EntityDescriptionFrom(ptr unsafe.Pointer) EntityDescription {
	return EntityDescription{objectivec.Object{objc.ID(ptr)}}
}




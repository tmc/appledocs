// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EntityMapping] class.
var entityMappingClass = _EntityMappingClass{objc.GetClass("NSEntityMapping")}

type _EntityMappingClass struct {
	class objc.Class
}

// A mapping instance that specifies how to map an entity from a source to a destination managed object model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSEntityMapping

type EntityMapping struct {
	objectivec.Object
}

// EntityMappingFrom constructs a [EntityMapping] from an unsafe.Pointer.
//
// A mapping instance that specifies how to map an entity from a source to a destination managed object model.
func EntityMappingFrom(ptr unsafe.Pointer) EntityMapping {
	return EntityMapping{objectivec.Object{objc.ID(ptr)}}
}




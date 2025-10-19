// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PropertyMapping] class.
var propertyMappingClass = _PropertyMappingClass{objc.GetClass("NSPropertyMapping")}

type _PropertyMappingClass struct {
	class objc.Class
}

// A mapping instance that specifies in a model how to map from a property in a source entity to a property in a destination entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyMapping

type PropertyMapping struct {
	objectivec.Object
}

// PropertyMappingFrom constructs a [PropertyMapping] from an unsafe.Pointer.
//
// A mapping instance that specifies in a model how to map from a property in a source entity to a property in a destination entity.
func PropertyMappingFrom(ptr unsafe.Pointer) PropertyMapping {
	return PropertyMapping{objectivec.Object{objc.ID(ptr)}}
}




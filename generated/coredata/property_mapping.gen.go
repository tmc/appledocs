// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PropertyMapping] class.
var (
	propertyMappingClass     _PropertyMappingClass
	propertyMappingClassOnce sync.Once
)

func getPropertyMappingClass() _PropertyMappingClass {
	propertyMappingClassOnce.Do(func() {
		propertyMappingClass = _PropertyMappingClass{objc.GetClass("NSPropertyMapping")}
	})
	return propertyMappingClass
}

type _PropertyMappingClass struct {
	class objc.Class
}

// An interface definition for the [PropertyMapping] class.
type IPropertyMapping interface {
	objectivec.IObject
}

// A mapping instance that specifies in a model how to map from a property in a source entity to a property in a destination entity.
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

// Alloc allocates a new instance without initialization.
func (pc _PropertyMappingClass) Alloc() PropertyMapping {
	rv := objc.Send[PropertyMapping](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PropertyMappingClass) New() PropertyMapping {
	rv := objc.Send[PropertyMapping](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PropertyMapping) Init() PropertyMapping {
	rv := objc.Send[PropertyMapping](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PropertyMapping) Autorelease() PropertyMapping {
	rv := objc.Send[PropertyMapping](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPropertyMapping creates a new PropertyMapping instance.
func NewPropertyMapping() PropertyMapping {
	return getPropertyMappingClass().New()
}


// The name of the property in the destination entity for the property mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyMapping/name
func (p_ PropertyMapping) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("name"))
	return rv
}

// SetName sets the value of the name property.
// The name of the property in the destination entity for the property mapping.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyMapping/name
func (p_ PropertyMapping) SetName(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setName:"), value)
}



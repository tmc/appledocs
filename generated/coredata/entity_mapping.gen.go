// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EntityMapping] class.
var (
	entityMappingClass     _EntityMappingClass
	entityMappingClassOnce sync.Once
)

func getEntityMappingClass() _EntityMappingClass {
	entityMappingClassOnce.Do(func() {
		entityMappingClass = _EntityMappingClass{objc.GetClass("NSEntityMapping")}
	})
	return entityMappingClass
}

type _EntityMappingClass struct {
	class objc.Class
}

// An interface definition for the [EntityMapping] class.
type IEntityMapping interface {
	objectivec.IObject
}

// A mapping instance that specifies how to map an entity from a source to a destination managed object model.
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

// Alloc allocates a new instance without initialization.
func (ec _EntityMappingClass) Alloc() EntityMapping {
	rv := objc.Send[EntityMapping](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EntityMappingClass) New() EntityMapping {
	rv := objc.Send[EntityMapping](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EntityMapping) Init() EntityMapping {
	rv := objc.Send[EntityMapping](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EntityMapping) Autorelease() EntityMapping {
	rv := objc.Send[EntityMapping](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEntityMapping creates a new EntityMapping instance.
func NewEntityMapping() EntityMapping {
	return getEntityMappingClass().New()
}





// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EntityDescription] class.
var (
	entityDescriptionClass     _EntityDescriptionClass
	entityDescriptionClassOnce sync.Once
)

func getEntityDescriptionClass() _EntityDescriptionClass {
	entityDescriptionClassOnce.Do(func() {
		entityDescriptionClass = _EntityDescriptionClass{objc.GetClass("NSEntityDescription")}
	})
	return entityDescriptionClass
}

type _EntityDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [EntityDescription] class.
type IEntityDescription interface {
	objectivec.IObject
}

// A description of a Core Data entity.
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

// Alloc allocates a new instance without initialization.
func (ec _EntityDescriptionClass) Alloc() EntityDescription {
	rv := objc.Send[EntityDescription](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EntityDescriptionClass) New() EntityDescription {
	rv := objc.Send[EntityDescription](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EntityDescription) Init() EntityDescription {
	rv := objc.Send[EntityDescription](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EntityDescription) Autorelease() EntityDescription {
	rv := objc.Send[EntityDescription](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEntityDescription creates a new EntityDescription instance.
func NewEntityDescription() EntityDescription {
	return getEntityDescriptionClass().New()
}





// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PropertyDescription] class.
var (
	propertyDescriptionClass     _PropertyDescriptionClass
	propertyDescriptionClassOnce sync.Once
)

func getPropertyDescriptionClass() _PropertyDescriptionClass {
	propertyDescriptionClassOnce.Do(func() {
		propertyDescriptionClass = _PropertyDescriptionClass{objc.GetClass("NSPropertyDescription")}
	})
	return propertyDescriptionClass
}

type _PropertyDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [PropertyDescription] class.
type IPropertyDescription interface {
	objectivec.IObject
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

// Alloc allocates a new instance without initialization.
func (pc _PropertyDescriptionClass) Alloc() PropertyDescription {
	rv := objc.Send[PropertyDescription](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PropertyDescriptionClass) New() PropertyDescription {
	rv := objc.Send[PropertyDescription](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PropertyDescription) Init() PropertyDescription {
	rv := objc.Send[PropertyDescription](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PropertyDescription) Autorelease() PropertyDescription {
	rv := objc.Send[PropertyDescription](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPropertyDescription creates a new PropertyDescription instance.
func NewPropertyDescription() PropertyDescription {
	return getPropertyDescriptionClass().New()
}





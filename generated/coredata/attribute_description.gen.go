// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AttributeDescription] class.
var (
	attributeDescriptionClass     _AttributeDescriptionClass
	attributeDescriptionClassOnce sync.Once
)

func getAttributeDescriptionClass() _AttributeDescriptionClass {
	attributeDescriptionClassOnce.Do(func() {
		attributeDescriptionClass = _AttributeDescriptionClass{objc.GetClass("NSAttributeDescription")}
	})
	return attributeDescriptionClass
}

type _AttributeDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [AttributeDescription] class.
type IAttributeDescription interface {
	IPropertyDescription
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

// Alloc allocates a new instance without initialization.
func (ac _AttributeDescriptionClass) Alloc() AttributeDescription {
	rv := objc.Send[AttributeDescription](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AttributeDescriptionClass) New() AttributeDescription {
	rv := objc.Send[AttributeDescription](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AttributeDescription) Init() AttributeDescription {
	rv := objc.Send[AttributeDescription](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AttributeDescription) Autorelease() AttributeDescription {
	rv := objc.Send[AttributeDescription](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAttributeDescription creates a new AttributeDescription instance.
func NewAttributeDescription() AttributeDescription {
	return getAttributeDescriptionClass().New()
}





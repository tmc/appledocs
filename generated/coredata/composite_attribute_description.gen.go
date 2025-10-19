// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CompositeAttributeDescription] class.
var compositeAttributeDescriptionClass = _CompositeAttributeDescriptionClass{objc.GetClass("NSCompositeAttributeDescription")}

type _CompositeAttributeDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [CompositeAttributeDescription] class.
type ICompositeAttributeDescription interface {
	IAttributeDescription
}

// A description of an attribute that derives its value by composing other attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCompositeAttributeDescription

type CompositeAttributeDescription struct {
	AttributeDescription
}

// CompositeAttributeDescriptionFrom constructs a [CompositeAttributeDescription] from an unsafe.Pointer.
//
// A description of an attribute that derives its value by composing other attributes.
func CompositeAttributeDescriptionFrom(ptr unsafe.Pointer) CompositeAttributeDescription {
	return CompositeAttributeDescription{
		AttributeDescription: AttributeDescriptionFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (cc _CompositeAttributeDescriptionClass) Alloc() CompositeAttributeDescription {
	rv := objc.Send[CompositeAttributeDescription](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (cc _CompositeAttributeDescriptionClass) New() CompositeAttributeDescription {
	rv := objc.Send[CompositeAttributeDescription](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CompositeAttributeDescription) Init() CompositeAttributeDescription {
	rv := objc.Send[CompositeAttributeDescription](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CompositeAttributeDescription) Autorelease() CompositeAttributeDescription {
	rv := objc.Send[CompositeAttributeDescription](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCompositeAttributeDescription creates a new CompositeAttributeDescription instance.
func NewCompositeAttributeDescription() CompositeAttributeDescription {
	return compositeAttributeDescriptionClass.New()
}





// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CompositeAttributeDescription] class.
var (
	CompositeAttributeDescriptionClass     _CompositeAttributeDescriptionClass
	CompositeAttributeDescriptionClassOnce sync.Once
)

func getCompositeAttributeDescriptionClass() _CompositeAttributeDescriptionClass {
	CompositeAttributeDescriptionClassOnce.Do(func() {
		CompositeAttributeDescriptionClass = _CompositeAttributeDescriptionClass{objc.GetClass("NSCompositeAttributeDescription")}
	})
	return CompositeAttributeDescriptionClass
}

type _CompositeAttributeDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [CompositeAttributeDescription] class.
type ICompositeAttributeDescription interface {
	IAttributeDescription
}

// A description of an attribute that derives its value by composing other attributes.
//
// Composite attributes enable you to define and store complex data types, and then query, index, and apply constraints to those types. Model classes use dictionaries to represent those composites in-memory, where each dictionary contains keys corresponding to the names of the underlying attributes. You may use composite attributes anywhere you use standard attributes, including lightweight migrations and CloudKit, through . You can even nest composites inside other composites to create complex object hierarchies without additional model classes. In most scenarios, prefer to use Xcode’s model editor to add composite attributes to your entities and then regenerate your model classes. However, if you need to create composites dynamically at runtime, create an instance of this class and populate its property with the necessary attribute descriptions. You can access a composite’s underlying attributes using namespaced key paths and property-like setters and getters, as the following example demonstrates:
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getCompositeAttributeDescriptionClass().New()
}


// The composed attribute descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCompositeAttributeDescription/elements
func (c_ CompositeAttributeDescription) Elements() []AttributeDescription {
	rv := objc.Send[[]AttributeDescription](c_.ID, objc.Sel("elements"))
	return rv
}


// SetElements sets the value of the elements property.
// The composed attribute descriptions.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCompositeAttributeDescription/elements
func (c_ CompositeAttributeDescription) SetElements(value []AttributeDescription) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setElements:"), nsArray)
}




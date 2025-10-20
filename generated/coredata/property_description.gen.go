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

// A description of a single property belonging to an entity.
//
// A property describes a single value within an object managed by the Core Data Framework. There are different types of property, each represented by a subclass which encapsulates the specific property behavior—see , , and . Note that a property name cannot be the same as any no-parameter method name of or . For example, you cannot give a property the name “description”. There are hundreds of methods on which may conflict with property names—and this list can grow without warning from frameworks or other libraries. You should avoid very general words (like “font”, and “color”) and words or phrases which overlap with Cocoa paradigms (such as “isEditing” and “objectSpecifier”). Properties—relationships as well as attributes—may be transient. A managed object context knows about transient properties and tracks changes made to them. Transient properties are ignored by the persistent store, and not just during saves: you cannot fetch using a predicate based on transients (although you can use transient properties to filter in memory yourself).
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


// The name of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/name
func (p_ PropertyDescription) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("name"))
	return rv
}

// SetName sets the value of the name property.
// The name of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPropertyDescription/name
func (p_ PropertyDescription) SetName(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setName:"), value)
}



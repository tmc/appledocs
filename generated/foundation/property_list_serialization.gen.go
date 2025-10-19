// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PropertyListSerialization] class.
var (
	propertyListSerializationClass     _PropertyListSerializationClass
	propertyListSerializationClassOnce sync.Once
)

func getPropertyListSerializationClass() _PropertyListSerializationClass {
	propertyListSerializationClassOnce.Do(func() {
		propertyListSerializationClass = _PropertyListSerializationClass{objc.GetClass("NSPropertyListSerialization")}
	})
	return propertyListSerializationClass
}

type _PropertyListSerializationClass struct {
	class objc.Class
}

// An interface definition for the [PropertyListSerialization] class.
type IPropertyListSerialization interface {
	objectivec.IObject
}

// An object that converts between a property list and one of several serialized representations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization
type PropertyListSerialization struct {
	objectivec.Object
}

// PropertyListSerializationFrom constructs a [PropertyListSerialization] from an unsafe.Pointer.
//
// An object that converts between a property list and one of several serialized representations.
func PropertyListSerializationFrom(ptr unsafe.Pointer) PropertyListSerialization {
	return PropertyListSerialization{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PropertyListSerializationClass) Alloc() PropertyListSerialization {
	rv := objc.Send[PropertyListSerialization](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PropertyListSerializationClass) New() PropertyListSerialization {
	rv := objc.Send[PropertyListSerialization](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PropertyListSerialization) Init() PropertyListSerialization {
	rv := objc.Send[PropertyListSerialization](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PropertyListSerialization) Autorelease() PropertyListSerialization {
	rv := objc.Send[PropertyListSerialization](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPropertyListSerialization creates a new PropertyListSerialization instance.
func NewPropertyListSerialization() PropertyListSerialization {
	return getPropertyListSerializationClass().New()
}





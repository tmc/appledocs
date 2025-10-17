// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PropertyListSerialization] class.
var propertyListSerializationClass = _PropertyListSerializationClass{objc.GetClass("NSPropertyListSerialization")}

type _PropertyListSerializationClass struct {
	class objc.Class
}

// An interface definition for the [PropertyListSerialization] class.
type IPropertyListSerialization interface {
	objectivec.IObject
}

// An object that converts between a property list and one of several serialized representations. [Full Topic]
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




// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AppleEventDescriptor] class.
var appleEventDescriptorClass = _AppleEventDescriptorClass{objc.GetClass("NSAppleEventDescriptor")}

type _AppleEventDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AppleEventDescriptor] class.
type IAppleEventDescriptor interface {
	objectivec.IObject
}

// A wrapper for the Apple event descriptor data type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor

type AppleEventDescriptor struct {
	objectivec.Object
}

// AppleEventDescriptorFrom constructs a [AppleEventDescriptor] from an unsafe.Pointer.
//
// A wrapper for the Apple event descriptor data type.
func AppleEventDescriptorFrom(ptr unsafe.Pointer) AppleEventDescriptor {
	return AppleEventDescriptor{objectivec.Object{objc.ID(ptr)}}
}




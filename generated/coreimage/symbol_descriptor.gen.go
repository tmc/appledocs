// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [symbolDescriptor] class.
var symbolDescriptorClass = _symbolDescriptorClass{objc.GetClass("symbolDescriptor")}

type _symbolDescriptorClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/symbolDescriptor-c.ivar

type symbolDescriptor struct {
	objectivec.Object
}

// symbolDescriptorFrom constructs a [symbolDescriptor] from an unsafe.Pointer.
func symbolDescriptorFrom(ptr unsafe.Pointer) symbolDescriptor {
	return symbolDescriptor{objectivec.Object{objc.ID(ptr)}}
}




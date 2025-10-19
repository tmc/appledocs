// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BarcodeDescriptor] class.
var barcodeDescriptorClass = _BarcodeDescriptorClass{objc.GetClass("CIBarcodeDescriptor")}

type _BarcodeDescriptorClass struct {
	class objc.Class
}

// An abstract base class that represents a machine-readable code’s attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBarcodeDescriptor

type BarcodeDescriptor struct {
	objectivec.Object
}

// BarcodeDescriptorFrom constructs a [BarcodeDescriptor] from an unsafe.Pointer.
//
// An abstract base class that represents a machine-readable code’s attributes.
func BarcodeDescriptorFrom(ptr unsafe.Pointer) BarcodeDescriptor {
	return BarcodeDescriptor{objectivec.Object{objc.ID(ptr)}}
}




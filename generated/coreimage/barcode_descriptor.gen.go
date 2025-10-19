// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BarcodeDescriptor] class.
var (
	barcodeDescriptorClass     _BarcodeDescriptorClass
	barcodeDescriptorClassOnce sync.Once
)

func getBarcodeDescriptorClass() _BarcodeDescriptorClass {
	barcodeDescriptorClassOnce.Do(func() {
		barcodeDescriptorClass = _BarcodeDescriptorClass{objc.GetClass("CIBarcodeDescriptor")}
	})
	return barcodeDescriptorClass
}

type _BarcodeDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [BarcodeDescriptor] class.
type IBarcodeDescriptor interface {
	objectivec.IObject
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

// Alloc allocates a new instance without initialization.
func (bc _BarcodeDescriptorClass) Alloc() BarcodeDescriptor {
	rv := objc.Send[BarcodeDescriptor](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BarcodeDescriptorClass) New() BarcodeDescriptor {
	rv := objc.Send[BarcodeDescriptor](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BarcodeDescriptor) Init() BarcodeDescriptor {
	rv := objc.Send[BarcodeDescriptor](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BarcodeDescriptor) Autorelease() BarcodeDescriptor {
	rv := objc.Send[BarcodeDescriptor](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBarcodeDescriptor creates a new BarcodeDescriptor instance.
func NewBarcodeDescriptor() BarcodeDescriptor {
	return getBarcodeDescriptorClass().New()
}





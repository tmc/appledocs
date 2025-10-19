// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTLTensorDescriptor] class.
var mTLTensorDescriptorClass = _MTLTensorDescriptorClass{objc.GetClass("MTLTensorDescriptor")}

type _MTLTensorDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTLTensorDescriptor] class.
type IMTLTensorDescriptor interface {
	objectivec.IObject
}

// A configuration type for creating new tensor instances. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor

type MTLTensorDescriptor struct {
	objectivec.Object
}

// MTLTensorDescriptorFrom constructs a [MTLTensorDescriptor] from an unsafe.Pointer.
//
// A configuration type for creating new tensor instances.
func MTLTensorDescriptorFrom(ptr unsafe.Pointer) MTLTensorDescriptor {
	return MTLTensorDescriptor{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (mc _MTLTensorDescriptorClass) Alloc() MTLTensorDescriptor {
	rv := objc.Send[MTLTensorDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MTLTensorDescriptorClass) New() MTLTensorDescriptor {
	rv := objc.Send[MTLTensorDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTLTensorDescriptor) Init() MTLTensorDescriptor {
	rv := objc.Send[MTLTensorDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTLTensorDescriptor) Autorelease() MTLTensorDescriptor {
	rv := objc.Send[MTLTensorDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLTensorDescriptor creates a new MTLTensorDescriptor instance.
func NewMTLTensorDescriptor() MTLTensorDescriptor {
	return mTLTensorDescriptorClass.New()
}





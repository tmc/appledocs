// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTLRenderPassDescriptor] class.
var mTLRenderPassDescriptorClass = _MTLRenderPassDescriptorClass{objc.GetClass("MTLRenderPassDescriptor")}

type _MTLRenderPassDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTLRenderPassDescriptor] class.
type IMTLRenderPassDescriptor interface {
	objectivec.IObject
}

// A group of render targets that hold the results of a render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassDescriptor

type MTLRenderPassDescriptor struct {
	objectivec.Object
}

// MTLRenderPassDescriptorFrom constructs a [MTLRenderPassDescriptor] from an unsafe.Pointer.
//
// A group of render targets that hold the results of a render pass.
func MTLRenderPassDescriptorFrom(ptr unsafe.Pointer) MTLRenderPassDescriptor {
	return MTLRenderPassDescriptor{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (mc _MTLRenderPassDescriptorClass) Alloc() MTLRenderPassDescriptor {
	rv := objc.Send[MTLRenderPassDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MTLRenderPassDescriptorClass) New() MTLRenderPassDescriptor {
	rv := objc.Send[MTLRenderPassDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTLRenderPassDescriptor) Init() MTLRenderPassDescriptor {
	rv := objc.Send[MTLRenderPassDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTLRenderPassDescriptor) Autorelease() MTLRenderPassDescriptor {
	rv := objc.Send[MTLRenderPassDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLRenderPassDescriptor creates a new MTLRenderPassDescriptor instance.
func NewMTLRenderPassDescriptor() MTLRenderPassDescriptor {
	return mTLRenderPassDescriptorClass.New()
}





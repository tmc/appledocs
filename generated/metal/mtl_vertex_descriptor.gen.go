// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTLVertexDescriptor] class.
var mTLVertexDescriptorClass = _MTLVertexDescriptorClass{objc.GetClass("MTLVertexDescriptor")}

type _MTLVertexDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTLVertexDescriptor] class.
type IMTLVertexDescriptor interface {
	objectivec.IObject
}

// An object that describes how to organize and map data to a vertex function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexDescriptor

type MTLVertexDescriptor struct {
	objectivec.Object
}

// MTLVertexDescriptorFrom constructs a [MTLVertexDescriptor] from an unsafe.Pointer.
//
// An object that describes how to organize and map data to a vertex function.
func MTLVertexDescriptorFrom(ptr unsafe.Pointer) MTLVertexDescriptor {
	return MTLVertexDescriptor{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (mc _MTLVertexDescriptorClass) Alloc() MTLVertexDescriptor {
	rv := objc.Send[MTLVertexDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MTLVertexDescriptorClass) New() MTLVertexDescriptor {
	rv := objc.Send[MTLVertexDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTLVertexDescriptor) Init() MTLVertexDescriptor {
	rv := objc.Send[MTLVertexDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTLVertexDescriptor) Autorelease() MTLVertexDescriptor {
	rv := objc.Send[MTLVertexDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLVertexDescriptor creates a new MTLVertexDescriptor instance.
func NewMTLVertexDescriptor() MTLVertexDescriptor {
	return mTLVertexDescriptorClass.New()
}





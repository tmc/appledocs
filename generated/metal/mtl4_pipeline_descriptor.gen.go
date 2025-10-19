// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTL4PipelineDescriptor] class.
var mTL4PipelineDescriptorClass = _MTL4PipelineDescriptorClass{objc.GetClass("MTL4PipelineDescriptor")}

type _MTL4PipelineDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4PipelineDescriptor] class.
type IMTL4PipelineDescriptor interface {
	objectivec.IObject
}

// A parent class referenced by other Metal classes. [Full Topic]

type MTL4PipelineDescriptor struct {
	objectivec.Object
}

// MTL4PipelineDescriptorFrom constructs a [MTL4PipelineDescriptor] from an unsafe.Pointer.
//
// A parent class referenced by other Metal classes.
func MTL4PipelineDescriptorFrom(ptr unsafe.Pointer) MTL4PipelineDescriptor {
	return MTL4PipelineDescriptor{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (mc _MTL4PipelineDescriptorClass) Alloc() MTL4PipelineDescriptor {
	rv := objc.Send[MTL4PipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MTL4PipelineDescriptorClass) New() MTL4PipelineDescriptor {
	rv := objc.Send[MTL4PipelineDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4PipelineDescriptor) Init() MTL4PipelineDescriptor {
	rv := objc.Send[MTL4PipelineDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4PipelineDescriptor) Autorelease() MTL4PipelineDescriptor {
	rv := objc.Send[MTL4PipelineDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4PipelineDescriptor creates a new MTL4PipelineDescriptor instance.
func NewMTL4PipelineDescriptor() MTL4PipelineDescriptor {
	return mTL4PipelineDescriptorClass.New()
}





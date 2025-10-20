// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTL4ComputePipelineDescriptor] class.
var (
	MTL4ComputePipelineDescriptorClass     _MTL4ComputePipelineDescriptorClass
	MTL4ComputePipelineDescriptorClassOnce sync.Once
)

func getMTL4ComputePipelineDescriptorClass() _MTL4ComputePipelineDescriptorClass {
	MTL4ComputePipelineDescriptorClassOnce.Do(func() {
		MTL4ComputePipelineDescriptorClass = _MTL4ComputePipelineDescriptorClass{objc.GetClass("MTL4ComputePipelineDescriptor")}
	})
	return MTL4ComputePipelineDescriptorClass
}

type _MTL4ComputePipelineDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4ComputePipelineDescriptor] class.
type IMTL4ComputePipelineDescriptor interface {
	objectivec.IObject
}

// Describes a compute pipeline state.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor
type MTL4ComputePipelineDescriptor struct {
	objectivec.Object
}

// MTL4ComputePipelineDescriptorFrom constructs a [MTL4ComputePipelineDescriptor] from an unsafe.Pointer.
//
// Describes a compute pipeline state.
func MTL4ComputePipelineDescriptorFrom(ptr unsafe.Pointer) MTL4ComputePipelineDescriptor {
	return MTL4ComputePipelineDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4ComputePipelineDescriptorClass) Alloc() MTL4ComputePipelineDescriptor {
	rv := objc.Send[MTL4ComputePipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4ComputePipelineDescriptorClass) New() MTL4ComputePipelineDescriptor {
	rv := objc.Send[MTL4ComputePipelineDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4ComputePipelineDescriptor) Init() MTL4ComputePipelineDescriptor {
	rv := objc.Send[MTL4ComputePipelineDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4ComputePipelineDescriptor) Autorelease() MTL4ComputePipelineDescriptor {
	rv := objc.Send[MTL4ComputePipelineDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4ComputePipelineDescriptor creates a new MTL4ComputePipelineDescriptor instance.
func NewMTL4ComputePipelineDescriptor() MTL4ComputePipelineDescriptor {
	return getMTL4ComputePipelineDescriptorClass().New()
}


// The maximum total number of threads that Metal can execute in a single threadgroup for the compute function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/maxTotalThreadsPerThreadgroup
func (m_ MTL4ComputePipelineDescriptor) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}


// SetMaxTotalThreadsPerThreadgroup sets the value of the maxTotalThreadsPerThreadgroup property.
// The maximum total number of threads that Metal can execute in a single threadgroup for the compute function.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/maxTotalThreadsPerThreadgroup
func (m_ MTL4ComputePipelineDescriptor) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}



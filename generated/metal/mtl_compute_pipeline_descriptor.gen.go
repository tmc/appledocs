// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ComputePipelineDescriptor] class.
var (
	ComputePipelineDescriptorClass     _ComputePipelineDescriptorClass
	ComputePipelineDescriptorClassOnce sync.Once
)

func getComputePipelineDescriptorClass() _ComputePipelineDescriptorClass {
	ComputePipelineDescriptorClassOnce.Do(func() {
		ComputePipelineDescriptorClass = _ComputePipelineDescriptorClass{objc.GetClass("MTLComputePipelineDescriptor")}
	})
	return ComputePipelineDescriptorClass
}

type _ComputePipelineDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [ComputePipelineDescriptor] class.
type IComputePipelineDescriptor interface {
	objectivec.IObject
}

// An instance describing the desired GPU state for a kernel call in a compute pass.
//
// A pipeline descriptor provides information necessary for creating an instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor
type ComputePipelineDescriptor struct {
	objectivec.Object
}

// ComputePipelineDescriptorFrom constructs a [ComputePipelineDescriptor] from an unsafe.Pointer.
//
// An instance describing the desired GPU state for a kernel call in a compute pass.
func ComputePipelineDescriptorFrom(ptr unsafe.Pointer) ComputePipelineDescriptor {
	return ComputePipelineDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ComputePipelineDescriptorClass) Alloc() ComputePipelineDescriptor {
	rv := objc.Send[ComputePipelineDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ComputePipelineDescriptorClass) New() ComputePipelineDescriptor {
	rv := objc.Send[ComputePipelineDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComputePipelineDescriptor) Init() ComputePipelineDescriptor {
	rv := objc.Send[ComputePipelineDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComputePipelineDescriptor) Autorelease() ComputePipelineDescriptor {
	rv := objc.Send[ComputePipelineDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComputePipelineDescriptor creates a new ComputePipelineDescriptor instance.
func NewComputePipelineDescriptor() ComputePipelineDescriptor {
	return getComputePipelineDescriptorClass().New()
}


// A string that identifies the instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/label
func (c_ ComputePipelineDescriptor) Label() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// A string that identifies the instance.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/label
func (c_ ComputePipelineDescriptor) SetLabel(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLabel:"), value)
}



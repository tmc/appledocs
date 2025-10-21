// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ComputePassDescriptor] class.
var (
	ComputePassDescriptorClass     _ComputePassDescriptorClass
	ComputePassDescriptorClassOnce sync.Once
)

func getComputePassDescriptorClass() _ComputePassDescriptorClass {
	ComputePassDescriptorClassOnce.Do(func() {
		ComputePassDescriptorClass = _ComputePassDescriptorClass{objc.GetClass("MTLComputePassDescriptor")}
	})
	return ComputePassDescriptorClass
}

type _ComputePassDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [ComputePassDescriptor] class.
type IComputePassDescriptor interface {
	objectivec.IObject
}

// A description of how to dispatch execution of pass commands and GPU performance sampling.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassDescriptor
type ComputePassDescriptor struct {
	objectivec.Object
}

// ComputePassDescriptorFrom constructs a [ComputePassDescriptor] from an unsafe.Pointer.
//
// A description of how to dispatch execution of pass commands and GPU performance sampling.
func ComputePassDescriptorFrom(ptr unsafe.Pointer) ComputePassDescriptor {
	return ComputePassDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ComputePassDescriptorClass) Alloc() ComputePassDescriptor {
	rv := objc.Send[ComputePassDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ComputePassDescriptorClass) New() ComputePassDescriptor {
	rv := objc.Send[ComputePassDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComputePassDescriptor) Init() ComputePassDescriptor {
	rv := objc.Send[ComputePassDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComputePassDescriptor) Autorelease() ComputePassDescriptor {
	rv := objc.Send[ComputePassDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComputePassDescriptor creates a new ComputePassDescriptor instance.
func NewComputePassDescriptor() ComputePassDescriptor {
	return getComputePassDescriptorClass().New()
}


// Creates a default compute pass descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassDescriptor/computePassDescriptor
func (cc _ComputePassDescriptorClass) ComputePassDescriptor() ComputePassDescriptor {
	rv := objc.Send[ComputePassDescriptor](objc.ID(cc.class), objc.Sel("computePassDescriptor"))
	return rv
}

// The strategy for dispatching any compute commands encoded in the compute pass.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepassdescriptor/dispatchtype
func (c_ ComputePassDescriptor) DispatchType() DispatchType {
	rv := objc.Send[DispatchType](c_.ID, objc.Sel("dispatchType"))
	return rv
}


// SetDispatchType sets the value of the dispatchType property.
// The strategy for dispatching any compute commands encoded in the compute pass.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepassdescriptor/dispatchtype
func (c_ ComputePassDescriptor) SetDispatchType(value DispatchType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDispatchType:"), value)
}

// The sample buffers that the compute pass can access.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepassdescriptor/samplebufferattachments
func (c_ ComputePassDescriptor) SampleBufferAttachments() MTLComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLComputePassSampleBufferAttachmentDescriptorArray](c_.ID, objc.Sel("sampleBufferAttachments"))
	return rv
}


// SetSampleBufferAttachments sets the value of the sampleBufferAttachments property.
// The sample buffers that the compute pass can access.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepassdescriptor/samplebufferattachments
func (c_ ComputePassDescriptor) SetSampleBufferAttachments(value IMTLComputePassSampleBufferAttachmentDescriptorArray) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSampleBufferAttachments:"), value)
}




// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AccelerationStructurePassSampleBufferAttachmentDescriptor] class.
var (
	AccelerationStructurePassSampleBufferAttachmentDescriptorClass     _AccelerationStructurePassSampleBufferAttachmentDescriptorClass
	AccelerationStructurePassSampleBufferAttachmentDescriptorClassOnce sync.Once
)

func getAccelerationStructurePassSampleBufferAttachmentDescriptorClass() _AccelerationStructurePassSampleBufferAttachmentDescriptorClass {
	AccelerationStructurePassSampleBufferAttachmentDescriptorClassOnce.Do(func() {
		AccelerationStructurePassSampleBufferAttachmentDescriptorClass = _AccelerationStructurePassSampleBufferAttachmentDescriptorClass{objc.GetClass("MTLAccelerationStructurePassSampleBufferAttachmentDescriptor")}
	})
	return AccelerationStructurePassSampleBufferAttachmentDescriptorClass
}

type _AccelerationStructurePassSampleBufferAttachmentDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AccelerationStructurePassSampleBufferAttachmentDescriptor] class.
type IAccelerationStructurePassSampleBufferAttachmentDescriptor interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptor
type AccelerationStructurePassSampleBufferAttachmentDescriptor struct {
	objectivec.Object
}

// AccelerationStructurePassSampleBufferAttachmentDescriptorFrom constructs a [AccelerationStructurePassSampleBufferAttachmentDescriptor] from an unsafe.Pointer.
func AccelerationStructurePassSampleBufferAttachmentDescriptorFrom(ptr unsafe.Pointer) AccelerationStructurePassSampleBufferAttachmentDescriptor {
	return AccelerationStructurePassSampleBufferAttachmentDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructurePassSampleBufferAttachmentDescriptorClass) Alloc() AccelerationStructurePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[AccelerationStructurePassSampleBufferAttachmentDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccelerationStructurePassSampleBufferAttachmentDescriptorClass) New() AccelerationStructurePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[AccelerationStructurePassSampleBufferAttachmentDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptor) Init() AccelerationStructurePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[AccelerationStructurePassSampleBufferAttachmentDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptor) Autorelease() AccelerationStructurePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[AccelerationStructurePassSampleBufferAttachmentDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructurePassSampleBufferAttachmentDescriptor creates a new AccelerationStructurePassSampleBufferAttachmentDescriptor instance.
func NewAccelerationStructurePassSampleBufferAttachmentDescriptor() AccelerationStructurePassSampleBufferAttachmentDescriptor {
	return getAccelerationStructurePassSampleBufferAttachmentDescriptorClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurepasssamplebufferattachmentdescriptor/startofencodersampleindex
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() int {
	rv := objc.Send[int](a_.ID, objc.Sel("startOfEncoderSampleIndex"))
	return rv
}


// SetStartOfEncoderSampleIndex sets the value of the startOfEncoderSampleIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurepasssamplebufferattachmentdescriptor/startofencodersampleindex
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStartOfEncoderSampleIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurepasssamplebufferattachmentdescriptor/endofencodersampleindex
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() int {
	rv := objc.Send[int](a_.ID, objc.Sel("endOfEncoderSampleIndex"))
	return rv
}


// SetEndOfEncoderSampleIndex sets the value of the endOfEncoderSampleIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurepasssamplebufferattachmentdescriptor/endofencodersampleindex
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEndOfEncoderSampleIndex:"), value)
}

// A specialized memory buffer that the GPU uses to store its counter data during the acceleration structure pass.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurepasssamplebufferattachmentdescriptor/samplebuffer
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptor) SampleBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("sampleBuffer"))
	return rv
}


// SetSampleBuffer sets the value of the sampleBuffer property.
// A specialized memory buffer that the GPU uses to store its counter data during the acceleration structure pass.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurepasssamplebufferattachmentdescriptor/samplebuffer
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptor) SetSampleBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleBuffer:"), value)
}




// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ComputePassSampleBufferAttachmentDescriptor] class.
var (
	ComputePassSampleBufferAttachmentDescriptorClass     _ComputePassSampleBufferAttachmentDescriptorClass
	ComputePassSampleBufferAttachmentDescriptorClassOnce sync.Once
)

func getComputePassSampleBufferAttachmentDescriptorClass() _ComputePassSampleBufferAttachmentDescriptorClass {
	ComputePassSampleBufferAttachmentDescriptorClassOnce.Do(func() {
		ComputePassSampleBufferAttachmentDescriptorClass = _ComputePassSampleBufferAttachmentDescriptorClass{objc.GetClass("MTLComputePassSampleBufferAttachmentDescriptor")}
	})
	return ComputePassSampleBufferAttachmentDescriptorClass
}

type _ComputePassSampleBufferAttachmentDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [ComputePassSampleBufferAttachmentDescriptor] class.
type IComputePassSampleBufferAttachmentDescriptor interface {
	objectivec.IObject
	StartOfEncoderSampleIndex() uint
	SetStartOfEncoderSampleIndex(value uint)
	EndOfEncoderSampleIndex() int
	SetEndOfEncoderSampleIndex(value int)
	SampleBuffer() unsafe.Pointer
	SetSampleBuffer(value unsafe.Pointer)
}

// A configuration that instructs the GPU where to store counter data from the beginning and end of a compute pass.
//
// For more context about configuring sample buffer attachments for compute passes, see . That article is one of a series in about sampling Metal hardware counters for performance measurement.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptor
type ComputePassSampleBufferAttachmentDescriptor struct {
	objectivec.Object
}

// ComputePassSampleBufferAttachmentDescriptorFrom constructs a [ComputePassSampleBufferAttachmentDescriptor] from an unsafe.Pointer.
//
// A configuration that instructs the GPU where to store counter data from the beginning and end of a compute pass.
func ComputePassSampleBufferAttachmentDescriptorFrom(ptr unsafe.Pointer) ComputePassSampleBufferAttachmentDescriptor {
	return ComputePassSampleBufferAttachmentDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ComputePassSampleBufferAttachmentDescriptorClass) Alloc() ComputePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[ComputePassSampleBufferAttachmentDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ComputePassSampleBufferAttachmentDescriptorClass) New() ComputePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[ComputePassSampleBufferAttachmentDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComputePassSampleBufferAttachmentDescriptor) Init() ComputePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[ComputePassSampleBufferAttachmentDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComputePassSampleBufferAttachmentDescriptor) Autorelease() ComputePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[ComputePassSampleBufferAttachmentDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComputePassSampleBufferAttachmentDescriptor creates a new ComputePassSampleBufferAttachmentDescriptor instance.
func NewComputePassSampleBufferAttachmentDescriptor() ComputePassSampleBufferAttachmentDescriptor {
	return getComputePassSampleBufferAttachmentDescriptorClass().New()
}


// An index within a counter sample buffer that tells the GPU where to store counter data from the start of a compute pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptor/startOfEncoderSampleIndex
func (c_ ComputePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("startOfEncoderSampleIndex"))
	return rv
}


// SetStartOfEncoderSampleIndex sets the value of the startOfEncoderSampleIndex property.
// An index within a counter sample buffer that tells the GPU where to store counter data from the start of a compute pass.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptor/startOfEncoderSampleIndex
func (c_ ComputePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStartOfEncoderSampleIndex:"), value)
}

// An index within a counter sample buffer that tells the GPU where to store counter data from the end of a compute pass.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepasssamplebufferattachmentdescriptor/endofencodersampleindex
func (c_ ComputePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() int {
	rv := objc.Send[int](c_.ID, objc.Sel("endOfEncoderSampleIndex"))
	return rv
}


// SetEndOfEncoderSampleIndex sets the value of the endOfEncoderSampleIndex property.
// An index within a counter sample buffer that tells the GPU where to store counter data from the end of a compute pass.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepasssamplebufferattachmentdescriptor/endofencodersampleindex
func (c_ ComputePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEndOfEncoderSampleIndex:"), value)
}

// A specialized memory buffer that the GPU uses to store its counter data during a compute pass.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepasssamplebufferattachmentdescriptor/samplebuffer
func (c_ ComputePassSampleBufferAttachmentDescriptor) SampleBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sampleBuffer"))
	return rv
}


// SetSampleBuffer sets the value of the sampleBuffer property.
// A specialized memory buffer that the GPU uses to store its counter data during a compute pass.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepasssamplebufferattachmentdescriptor/samplebuffer
func (c_ ComputePassSampleBufferAttachmentDescriptor) SetSampleBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSampleBuffer:"), value)
}




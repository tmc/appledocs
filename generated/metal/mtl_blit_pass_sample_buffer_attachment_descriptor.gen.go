// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BlitPassSampleBufferAttachmentDescriptor] class.
var (
	BlitPassSampleBufferAttachmentDescriptorClass     _BlitPassSampleBufferAttachmentDescriptorClass
	BlitPassSampleBufferAttachmentDescriptorClassOnce sync.Once
)

func getBlitPassSampleBufferAttachmentDescriptorClass() _BlitPassSampleBufferAttachmentDescriptorClass {
	BlitPassSampleBufferAttachmentDescriptorClassOnce.Do(func() {
		BlitPassSampleBufferAttachmentDescriptorClass = _BlitPassSampleBufferAttachmentDescriptorClass{objc.GetClass("MTLBlitPassSampleBufferAttachmentDescriptor")}
	})
	return BlitPassSampleBufferAttachmentDescriptorClass
}

type _BlitPassSampleBufferAttachmentDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [BlitPassSampleBufferAttachmentDescriptor] class.
type IBlitPassSampleBufferAttachmentDescriptor interface {
	objectivec.IObject
	StartOfEncoderSampleIndex() uint
	SetStartOfEncoderSampleIndex(value uint)
	EndOfEncoderSampleIndex() int
	SetEndOfEncoderSampleIndex(value int)
	SampleBuffer() unsafe.Pointer
	SetSampleBuffer(value unsafe.Pointer)
}

// A configuration that instructs the GPU where to store counter data from the beginning and end of a blit pass.
//
// See for more context about configuring instances of this type. That article is one of a series of articles in .
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptor
type BlitPassSampleBufferAttachmentDescriptor struct {
	objectivec.Object
}

// BlitPassSampleBufferAttachmentDescriptorFrom constructs a [BlitPassSampleBufferAttachmentDescriptor] from an unsafe.Pointer.
//
// A configuration that instructs the GPU where to store counter data from the beginning and end of a blit pass.
func BlitPassSampleBufferAttachmentDescriptorFrom(ptr unsafe.Pointer) BlitPassSampleBufferAttachmentDescriptor {
	return BlitPassSampleBufferAttachmentDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BlitPassSampleBufferAttachmentDescriptorClass) Alloc() BlitPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[BlitPassSampleBufferAttachmentDescriptor](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BlitPassSampleBufferAttachmentDescriptorClass) New() BlitPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[BlitPassSampleBufferAttachmentDescriptor](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BlitPassSampleBufferAttachmentDescriptor) Init() BlitPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[BlitPassSampleBufferAttachmentDescriptor](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BlitPassSampleBufferAttachmentDescriptor) Autorelease() BlitPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[BlitPassSampleBufferAttachmentDescriptor](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBlitPassSampleBufferAttachmentDescriptor creates a new BlitPassSampleBufferAttachmentDescriptor instance.
func NewBlitPassSampleBufferAttachmentDescriptor() BlitPassSampleBufferAttachmentDescriptor {
	return getBlitPassSampleBufferAttachmentDescriptorClass().New()
}


// An index within a counter sample buffer that tells the GPU where to store counter data from the start of a blit pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptor/startOfEncoderSampleIndex
func (b_ BlitPassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](b_.ID, objc.Sel("startOfEncoderSampleIndex"))
	return rv
}


// SetStartOfEncoderSampleIndex sets the value of the startOfEncoderSampleIndex property.
// An index within a counter sample buffer that tells the GPU where to store counter data from the start of a blit pass.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptor/startOfEncoderSampleIndex
func (b_ BlitPassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex(value uint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setStartOfEncoderSampleIndex:"), value)
}

// An index within a counter sample buffer that tells the GPU where to store counter data from the end of a blit pass.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpasssamplebufferattachmentdescriptor/endofencodersampleindex
func (b_ BlitPassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() int {
	rv := objc.Send[int](b_.ID, objc.Sel("endOfEncoderSampleIndex"))
	return rv
}


// SetEndOfEncoderSampleIndex sets the value of the endOfEncoderSampleIndex property.
// An index within a counter sample buffer that tells the GPU where to store counter data from the end of a blit pass.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpasssamplebufferattachmentdescriptor/endofencodersampleindex
func (b_ BlitPassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setEndOfEncoderSampleIndex:"), value)
}

// A specialized memory buffer that the GPU uses to store its counter data during the blit pass.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpasssamplebufferattachmentdescriptor/samplebuffer
func (b_ BlitPassSampleBufferAttachmentDescriptor) SampleBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("sampleBuffer"))
	return rv
}


// SetSampleBuffer sets the value of the sampleBuffer property.
// A specialized memory buffer that the GPU uses to store its counter data during the blit pass.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitpasssamplebufferattachmentdescriptor/samplebuffer
func (b_ BlitPassSampleBufferAttachmentDescriptor) SetSampleBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSampleBuffer:"), value)
}




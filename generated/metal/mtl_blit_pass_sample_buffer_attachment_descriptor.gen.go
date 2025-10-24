// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLBlitPassSampleBufferAttachmentDescriptor */


/* debug [class_header]: Header for MTLBlitPassSampleBufferAttachmentDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BlitPassSampleBufferAttachmentDescriptor */
// An interface definition for the [BlitPassSampleBufferAttachmentDescriptor] class.
type IBlitPassSampleBufferAttachmentDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BlitPassSampleBufferAttachmentDescriptor */
	// properties:
	EndOfEncoderSampleIndex() uint
	SetEndOfEncoderSampleIndex(value uint)
	SampleBuffer() unsafe.Pointer
	SetSampleBuffer(value unsafe.Pointer)
	StartOfEncoderSampleIndex() uint
	SetStartOfEncoderSampleIndex(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BlitPassSampleBufferAttachmentDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BlitPassSampleBufferAttachmentDescriptor */
// Alloc allocates a new instance without initialization.
func (bc _BlitPassSampleBufferAttachmentDescriptorClass) Alloc() BlitPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[BlitPassSampleBufferAttachmentDescriptor](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BlitPassSampleBufferAttachmentDescriptor */
// A configuration that instructs the GPU where to store counter data from the beginning and end of a blit pass.
//
// See for more context about configuring instances of this type. That article is one of a series of articles in .


// A configuration that instructs the GPU where to store counter data from the beginning and end of a blit pass.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BlitPassSampleBufferAttachmentDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BlitPassSampleBufferAttachmentDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BlitPassSampleBufferAttachmentDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BlitPassSampleBufferAttachmentDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BlitPassSampleBufferAttachmentDescriptor */

// An index within a counter sample buffer that tells the GPU where to store counter data from the end of a blit pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptor/endOfEncoderSampleIndex
func (b_ BlitPassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](b_.ID, objc.Sel("endOfEncoderSampleIndex"))
	return rv
}/* debug [instance_properties/getter]: endOfEncoderSampleIndex */


// An index within a counter sample buffer that tells the GPU where to store counter data from the end of a blit pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptor/endOfEncoderSampleIndex
func (b_ BlitPassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex(value uint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setEndOfEncoderSampleIndex:"), value)
}/* debug [instance_properties/setter]: endOfEncoderSampleIndex */


// A specialized memory buffer that the GPU uses to store its counter data during the blit pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptor/sampleBuffer
func (b_ BlitPassSampleBufferAttachmentDescriptor) SampleBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("sampleBuffer"))
	return rv
}/* debug [instance_properties/getter]: sampleBuffer */


// A specialized memory buffer that the GPU uses to store its counter data during the blit pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptor/sampleBuffer
func (b_ BlitPassSampleBufferAttachmentDescriptor) SetSampleBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSampleBuffer:"), value)
}/* debug [instance_properties/setter]: sampleBuffer */


// An index within a counter sample buffer that tells the GPU where to store counter data from the start of a blit pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptor/startOfEncoderSampleIndex
func (b_ BlitPassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](b_.ID, objc.Sel("startOfEncoderSampleIndex"))
	return rv
}/* debug [instance_properties/getter]: startOfEncoderSampleIndex */


// An index within a counter sample buffer that tells the GPU where to store counter data from the start of a blit pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptor/startOfEncoderSampleIndex
func (b_ BlitPassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex(value uint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setStartOfEncoderSampleIndex:"), value)
}/* debug [instance_properties/setter]: startOfEncoderSampleIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLBlitPassSampleBufferAttachmentDescriptor */




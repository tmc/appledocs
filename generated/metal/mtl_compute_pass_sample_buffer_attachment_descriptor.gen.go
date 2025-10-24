// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLComputePassSampleBufferAttachmentDescriptor */


/* debug [class_header]: Header for MTLComputePassSampleBufferAttachmentDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ComputePassSampleBufferAttachmentDescriptor */
// An interface definition for the [ComputePassSampleBufferAttachmentDescriptor] class.
type IComputePassSampleBufferAttachmentDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ComputePassSampleBufferAttachmentDescriptor */
	// properties:
	EndOfEncoderSampleIndex() uint
	SetEndOfEncoderSampleIndex(value uint)
	SampleBuffer() unsafe.Pointer
	SetSampleBuffer(value unsafe.Pointer)
	StartOfEncoderSampleIndex() uint
	SetStartOfEncoderSampleIndex(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ComputePassSampleBufferAttachmentDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ComputePassSampleBufferAttachmentDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _ComputePassSampleBufferAttachmentDescriptorClass) Alloc() ComputePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[ComputePassSampleBufferAttachmentDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ComputePassSampleBufferAttachmentDescriptor */
// A configuration that instructs the GPU where to store counter data from the beginning and end of a compute pass.
//
// For more context about configuring sample buffer attachments for compute passes, see . That article is one of a series in about sampling Metal hardware counters for performance measurement.


// A configuration that instructs the GPU where to store counter data from the beginning and end of a compute pass.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ComputePassSampleBufferAttachmentDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ComputePassSampleBufferAttachmentDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ComputePassSampleBufferAttachmentDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ComputePassSampleBufferAttachmentDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ComputePassSampleBufferAttachmentDescriptor */

// An index within a counter sample buffer that tells the GPU where to store counter data from the end of a compute pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptor/endOfEncoderSampleIndex
func (c_ ComputePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("endOfEncoderSampleIndex"))
	return rv
}/* debug [instance_properties/getter]: endOfEncoderSampleIndex */


// An index within a counter sample buffer that tells the GPU where to store counter data from the end of a compute pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptor/endOfEncoderSampleIndex
func (c_ ComputePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEndOfEncoderSampleIndex:"), value)
}/* debug [instance_properties/setter]: endOfEncoderSampleIndex */


// A specialized memory buffer that the GPU uses to store its counter data during a compute pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptor/sampleBuffer
func (c_ ComputePassSampleBufferAttachmentDescriptor) SampleBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sampleBuffer"))
	return rv
}/* debug [instance_properties/getter]: sampleBuffer */


// A specialized memory buffer that the GPU uses to store its counter data during a compute pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptor/sampleBuffer
func (c_ ComputePassSampleBufferAttachmentDescriptor) SetSampleBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSampleBuffer:"), value)
}/* debug [instance_properties/setter]: sampleBuffer */


// An index within a counter sample buffer that tells the GPU where to store counter data from the start of a compute pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptor/startOfEncoderSampleIndex
func (c_ ComputePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("startOfEncoderSampleIndex"))
	return rv
}/* debug [instance_properties/getter]: startOfEncoderSampleIndex */


// An index within a counter sample buffer that tells the GPU where to store counter data from the start of a compute pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptor/startOfEncoderSampleIndex
func (c_ ComputePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStartOfEncoderSampleIndex:"), value)
}/* debug [instance_properties/setter]: startOfEncoderSampleIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLComputePassSampleBufferAttachmentDescriptor */




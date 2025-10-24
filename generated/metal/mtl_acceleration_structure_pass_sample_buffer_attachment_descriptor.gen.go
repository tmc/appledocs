// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLAccelerationStructurePassSampleBufferAttachmentDescriptor */


/* debug [class_header]: Header for MTLAccelerationStructurePassSampleBufferAttachmentDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccelerationStructurePassSampleBufferAttachmentDescriptor */
// An interface definition for the [AccelerationStructurePassSampleBufferAttachmentDescriptor] class.
type IAccelerationStructurePassSampleBufferAttachmentDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AccelerationStructurePassSampleBufferAttachmentDescriptor */
	// properties:
	EndOfEncoderSampleIndex() uint
	SetEndOfEncoderSampleIndex(value uint)
	SampleBuffer() unsafe.Pointer
	SetSampleBuffer(value unsafe.Pointer)
	StartOfEncoderSampleIndex() uint
	SetStartOfEncoderSampleIndex(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccelerationStructurePassSampleBufferAttachmentDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccelerationStructurePassSampleBufferAttachmentDescriptor */
// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructurePassSampleBufferAttachmentDescriptorClass) Alloc() AccelerationStructurePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[AccelerationStructurePassSampleBufferAttachmentDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccelerationStructurePassSampleBufferAttachmentDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptor
type AccelerationStructurePassSampleBufferAttachmentDescriptor struct {
	objectivec.Object
}

// AccelerationStructurePassSampleBufferAttachmentDescriptorFrom constructs a [AccelerationStructurePassSampleBufferAttachmentDescriptor] from an unsafe.Pointer.
func AccelerationStructurePassSampleBufferAttachmentDescriptorFrom(ptr unsafe.Pointer) AccelerationStructurePassSampleBufferAttachmentDescriptor {
	return AccelerationStructurePassSampleBufferAttachmentDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccelerationStructurePassSampleBufferAttachmentDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccelerationStructurePassSampleBufferAttachmentDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccelerationStructurePassSampleBufferAttachmentDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccelerationStructurePassSampleBufferAttachmentDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccelerationStructurePassSampleBufferAttachmentDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptor/endOfEncoderSampleIndex
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("endOfEncoderSampleIndex"))
	return rv
}/* debug [instance_properties/getter]: endOfEncoderSampleIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptor/endOfEncoderSampleIndex
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEndOfEncoderSampleIndex:"), value)
}/* debug [instance_properties/setter]: endOfEncoderSampleIndex */


// A specialized memory buffer that the GPU uses to store its counter data during the acceleration structure pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptor/sampleBuffer
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptor) SampleBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("sampleBuffer"))
	return rv
}/* debug [instance_properties/getter]: sampleBuffer */


// A specialized memory buffer that the GPU uses to store its counter data during the acceleration structure pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptor/sampleBuffer
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptor) SetSampleBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleBuffer:"), value)
}/* debug [instance_properties/setter]: sampleBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptor/startOfEncoderSampleIndex
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("startOfEncoderSampleIndex"))
	return rv
}/* debug [instance_properties/getter]: startOfEncoderSampleIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptor/startOfEncoderSampleIndex
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStartOfEncoderSampleIndex:"), value)
}/* debug [instance_properties/setter]: startOfEncoderSampleIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLAccelerationStructurePassSampleBufferAttachmentDescriptor */




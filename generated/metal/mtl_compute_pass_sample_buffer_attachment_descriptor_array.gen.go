// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLComputePassSampleBufferAttachmentDescriptorArray */


/* debug [class_header]: Header for MTLComputePassSampleBufferAttachmentDescriptorArray */
// The class instance for the [ComputePassSampleBufferAttachmentDescriptorArray] class.
var (
	ComputePassSampleBufferAttachmentDescriptorArrayClass     _ComputePassSampleBufferAttachmentDescriptorArrayClass
	ComputePassSampleBufferAttachmentDescriptorArrayClassOnce sync.Once
)

func getComputePassSampleBufferAttachmentDescriptorArrayClass() _ComputePassSampleBufferAttachmentDescriptorArrayClass {
	ComputePassSampleBufferAttachmentDescriptorArrayClassOnce.Do(func() {
		ComputePassSampleBufferAttachmentDescriptorArrayClass = _ComputePassSampleBufferAttachmentDescriptorArrayClass{objc.GetClass("MTLComputePassSampleBufferAttachmentDescriptorArray")}
	})
	return ComputePassSampleBufferAttachmentDescriptorArrayClass
}

type _ComputePassSampleBufferAttachmentDescriptorArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ComputePassSampleBufferAttachmentDescriptorArray */
// An interface definition for the [ComputePassSampleBufferAttachmentDescriptorArray] class.
type IComputePassSampleBufferAttachmentDescriptorArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ComputePassSampleBufferAttachmentDescriptorArray */
	// properties:
	CounterSets() CounterSet /* not a class type */
	SetCounterSets(value CounterSet /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ComputePassSampleBufferAttachmentDescriptorArray */
	// methods:
	SetObjectAtIndexedSubscript(attachment IMTLComputePassSampleBufferAttachmentDescriptor, attachmentIndex uint)
	ObjectAtIndexedSubscript(attachmentIndex uint) IComputePassSampleBufferAttachmentDescriptor
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ComputePassSampleBufferAttachmentDescriptorArray */
// Alloc allocates a new instance without initialization.
func (cc _ComputePassSampleBufferAttachmentDescriptorArrayClass) Alloc() ComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ComputePassSampleBufferAttachmentDescriptorArray](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ComputePassSampleBufferAttachmentDescriptorArrayClass) New() ComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ComputePassSampleBufferAttachmentDescriptorArray](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComputePassSampleBufferAttachmentDescriptorArray) Init() ComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ComputePassSampleBufferAttachmentDescriptorArray](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComputePassSampleBufferAttachmentDescriptorArray) Autorelease() ComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ComputePassSampleBufferAttachmentDescriptorArray](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComputePassSampleBufferAttachmentDescriptorArray creates a new ComputePassSampleBufferAttachmentDescriptorArray instance.
func NewComputePassSampleBufferAttachmentDescriptorArray() ComputePassSampleBufferAttachmentDescriptorArray {
	return getComputePassSampleBufferAttachmentDescriptorArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ComputePassSampleBufferAttachmentDescriptorArray */
// A container that stores an array of sample buffer attachments for a compute pass.
//
// The number of elements in the array is at least the number of elements in an instance’s property.


// A container that stores an array of sample buffer attachments for a compute pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptorArray
type ComputePassSampleBufferAttachmentDescriptorArray struct {
	objectivec.Object
}

// ComputePassSampleBufferAttachmentDescriptorArrayFrom constructs a [ComputePassSampleBufferAttachmentDescriptorArray] from an unsafe.Pointer.
//
// A container that stores an array of sample buffer attachments for a compute pass.
func ComputePassSampleBufferAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) ComputePassSampleBufferAttachmentDescriptorArray {
	return ComputePassSampleBufferAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ComputePassSampleBufferAttachmentDescriptorArray *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ComputePassSampleBufferAttachmentDescriptorArray */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ComputePassSampleBufferAttachmentDescriptorArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ComputePassSampleBufferAttachmentDescriptorArray */

// Sets the descriptor object for the specified sample buffer attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (c_ ComputePassSampleBufferAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLComputePassSampleBufferAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}/* debug [instance_methods/method]: SetObjectAtIndexedSubscript */


// Returns the descriptor object for the specified sample buffer attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptorArray/subscript(_:)
func (c_ ComputePassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IComputePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[ComputePassSampleBufferAttachmentDescriptor](c_.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ComputePassSampleBufferAttachmentDescriptorArray */

// The counter sets supported by the device object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/countersets
func (c_ ComputePassSampleBufferAttachmentDescriptorArray) CounterSets() CounterSet /* not a class type */ {
	rv := objc.Send[CounterSet](c_.ID, objc.Sel("counterSets"))
	return rv
}/* debug [instance_properties/getter]: counterSets */


// The counter sets supported by the device object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/countersets
func (c_ ComputePassSampleBufferAttachmentDescriptorArray) SetCounterSets(value CounterSet /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCounterSets:"), value)
}/* debug [instance_properties/setter]: counterSets */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLComputePassSampleBufferAttachmentDescriptorArray */




// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray */


/* debug [class_header]: Header for MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray */
// The class instance for the [AccelerationStructurePassSampleBufferAttachmentDescriptorArray] class.
var (
	AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass     _AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass
	AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClassOnce sync.Once
)

func getAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass() _AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass {
	AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClassOnce.Do(func() {
		AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass = _AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass{objc.GetClass("MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray")}
	})
	return AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass
}

type _AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccelerationStructurePassSampleBufferAttachmentDescriptorArray */
// An interface definition for the [AccelerationStructurePassSampleBufferAttachmentDescriptorArray] class.
type IAccelerationStructurePassSampleBufferAttachmentDescriptorArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AccelerationStructurePassSampleBufferAttachmentDescriptorArray */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccelerationStructurePassSampleBufferAttachmentDescriptorArray */
	// methods:
	SetObjectAtIndexedSubscript(attachment IMTLAccelerationStructurePassSampleBufferAttachmentDescriptor, attachmentIndex uint)
	ObjectAtIndexedSubscript(attachmentIndex uint) IAccelerationStructurePassSampleBufferAttachmentDescriptor
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccelerationStructurePassSampleBufferAttachmentDescriptorArray */
// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass) Alloc() AccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[AccelerationStructurePassSampleBufferAttachmentDescriptorArray](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass) New() AccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[AccelerationStructurePassSampleBufferAttachmentDescriptorArray](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptorArray) Init() AccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[AccelerationStructurePassSampleBufferAttachmentDescriptorArray](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptorArray) Autorelease() AccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[AccelerationStructurePassSampleBufferAttachmentDescriptorArray](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructurePassSampleBufferAttachmentDescriptorArray creates a new AccelerationStructurePassSampleBufferAttachmentDescriptorArray instance.
func NewAccelerationStructurePassSampleBufferAttachmentDescriptorArray() AccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	return getAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccelerationStructurePassSampleBufferAttachmentDescriptorArray */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray
type AccelerationStructurePassSampleBufferAttachmentDescriptorArray struct {
	objectivec.Object
}

// AccelerationStructurePassSampleBufferAttachmentDescriptorArrayFrom constructs a [AccelerationStructurePassSampleBufferAttachmentDescriptorArray] from an unsafe.Pointer.
func AccelerationStructurePassSampleBufferAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) AccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	return AccelerationStructurePassSampleBufferAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccelerationStructurePassSampleBufferAttachmentDescriptorArray *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccelerationStructurePassSampleBufferAttachmentDescriptorArray */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccelerationStructurePassSampleBufferAttachmentDescriptorArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccelerationStructurePassSampleBufferAttachmentDescriptorArray */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLAccelerationStructurePassSampleBufferAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}/* debug [instance_methods/method]: SetObjectAtIndexedSubscript */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray/subscript(_:)
func (a_ AccelerationStructurePassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IAccelerationStructurePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[AccelerationStructurePassSampleBufferAttachmentDescriptor](a_.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccelerationStructurePassSampleBufferAttachmentDescriptorArray */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray */




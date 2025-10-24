// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLAccelerationStructurePassDescriptor */


/* debug [class_header]: Header for MTLAccelerationStructurePassDescriptor */
// The class instance for the [AccelerationStructurePassDescriptor] class.
var (
	AccelerationStructurePassDescriptorClass     _AccelerationStructurePassDescriptorClass
	AccelerationStructurePassDescriptorClassOnce sync.Once
)

func getAccelerationStructurePassDescriptorClass() _AccelerationStructurePassDescriptorClass {
	AccelerationStructurePassDescriptorClassOnce.Do(func() {
		AccelerationStructurePassDescriptorClass = _AccelerationStructurePassDescriptorClass{objc.GetClass("MTLAccelerationStructurePassDescriptor")}
	})
	return AccelerationStructurePassDescriptorClass
}

type _AccelerationStructurePassDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccelerationStructurePassDescriptor */
// An interface definition for the [AccelerationStructurePassDescriptor] class.
type IAccelerationStructurePassDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AccelerationStructurePassDescriptor */
	// properties:
	SampleBufferAttachments() IMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccelerationStructurePassDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccelerationStructurePassDescriptor */
// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructurePassDescriptorClass) Alloc() AccelerationStructurePassDescriptor {
	rv := objc.Send[AccelerationStructurePassDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccelerationStructurePassDescriptorClass) New() AccelerationStructurePassDescriptor {
	rv := objc.Send[AccelerationStructurePassDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructurePassDescriptor) Init() AccelerationStructurePassDescriptor {
	rv := objc.Send[AccelerationStructurePassDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructurePassDescriptor) Autorelease() AccelerationStructurePassDescriptor {
	rv := objc.Send[AccelerationStructurePassDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructurePassDescriptor creates a new AccelerationStructurePassDescriptor instance.
func NewAccelerationStructurePassDescriptor() AccelerationStructurePassDescriptor {
	return getAccelerationStructurePassDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccelerationStructurePassDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassDescriptor
type AccelerationStructurePassDescriptor struct {
	objectivec.Object
}

// AccelerationStructurePassDescriptorFrom constructs a [AccelerationStructurePassDescriptor] from an unsafe.Pointer.
func AccelerationStructurePassDescriptorFrom(ptr unsafe.Pointer) AccelerationStructurePassDescriptor {
	return AccelerationStructurePassDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccelerationStructurePassDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccelerationStructurePassDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassDescriptor/accelerationStructurePassDescriptor
func (ac _AccelerationStructurePassDescriptorClass) AccelerationStructurePassDescriptor() IAccelerationStructurePassDescriptor {
	rv := objc.Send[AccelerationStructurePassDescriptor](objc.ID(ac.class), objc.Sel("accelerationStructurePassDescriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AccelerationStructurePassDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccelerationStructurePassDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccelerationStructurePassDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccelerationStructurePassDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassDescriptor/sampleBufferAttachments
func (a_ AccelerationStructurePassDescriptor) SampleBufferAttachments() IMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[AccelerationStructurePassSampleBufferAttachmentDescriptorArray](a_.ID, objc.Sel("sampleBufferAttachments"))
	return rv
}/* debug [instance_properties/getter]: sampleBufferAttachments */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLAccelerationStructurePassDescriptor */




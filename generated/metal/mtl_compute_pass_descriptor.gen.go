// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLComputePassDescriptor */


/* debug [class_header]: Header for MTLComputePassDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ComputePassDescriptor */
// An interface definition for the [ComputePassDescriptor] class.
type IComputePassDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ComputePassDescriptor */
	// properties:
	DispatchType() DispatchType
	SetDispatchType(value DispatchType)
	SampleBufferAttachments() IMTLComputePassSampleBufferAttachmentDescriptorArray
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ComputePassDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ComputePassDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _ComputePassDescriptorClass) Alloc() ComputePassDescriptor {
	rv := objc.Send[ComputePassDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ComputePassDescriptor */
// A description of how to dispatch execution of pass commands and GPU performance sampling.


// A description of how to dispatch execution of pass commands and GPU performance sampling.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ComputePassDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ComputePassDescriptor */

// Creates a default compute pass descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassDescriptor/computePassDescriptor
func (cc _ComputePassDescriptorClass) ComputePassDescriptor() IComputePassDescriptor {
	rv := objc.Send[ComputePassDescriptor](objc.ID(cc.class), objc.Sel("computePassDescriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ComputePassDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ComputePassDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ComputePassDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ComputePassDescriptor */

// The strategy for dispatching any compute commands encoded in the compute pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassDescriptor/dispatchType
func (c_ ComputePassDescriptor) DispatchType() DispatchType {
	rv := objc.Send[DispatchType](c_.ID, objc.Sel("dispatchType"))
	return rv
}/* debug [instance_properties/getter]: dispatchType */


// The strategy for dispatching any compute commands encoded in the compute pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassDescriptor/dispatchType
func (c_ ComputePassDescriptor) SetDispatchType(value DispatchType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDispatchType:"), value)
}/* debug [instance_properties/setter]: dispatchType */


// The sample buffers that the compute pass can access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassDescriptor/sampleBufferAttachments
func (c_ ComputePassDescriptor) SampleBufferAttachments() IMTLComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ComputePassSampleBufferAttachmentDescriptorArray](c_.ID, objc.Sel("sampleBufferAttachments"))
	return rv
}/* debug [instance_properties/getter]: sampleBufferAttachments */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLComputePassDescriptor */




// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLBlitPassSampleBufferAttachmentDescriptorArray */


/* debug [class_header]: Header for MTLBlitPassSampleBufferAttachmentDescriptorArray */
// The class instance for the [BlitPassSampleBufferAttachmentDescriptorArray] class.
var (
	BlitPassSampleBufferAttachmentDescriptorArrayClass     _BlitPassSampleBufferAttachmentDescriptorArrayClass
	BlitPassSampleBufferAttachmentDescriptorArrayClassOnce sync.Once
)

func getBlitPassSampleBufferAttachmentDescriptorArrayClass() _BlitPassSampleBufferAttachmentDescriptorArrayClass {
	BlitPassSampleBufferAttachmentDescriptorArrayClassOnce.Do(func() {
		BlitPassSampleBufferAttachmentDescriptorArrayClass = _BlitPassSampleBufferAttachmentDescriptorArrayClass{objc.GetClass("MTLBlitPassSampleBufferAttachmentDescriptorArray")}
	})
	return BlitPassSampleBufferAttachmentDescriptorArrayClass
}

type _BlitPassSampleBufferAttachmentDescriptorArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BlitPassSampleBufferAttachmentDescriptorArray */
// An interface definition for the [BlitPassSampleBufferAttachmentDescriptorArray] class.
type IBlitPassSampleBufferAttachmentDescriptorArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BlitPassSampleBufferAttachmentDescriptorArray */
	// properties:
	CounterSets() CounterSet /* not a class type */
	SetCounterSets(value CounterSet /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BlitPassSampleBufferAttachmentDescriptorArray */
	// methods:
	SetObjectAtIndexedSubscript(attachment IMTLBlitPassSampleBufferAttachmentDescriptor, attachmentIndex uint)
	ObjectAtIndexedSubscript(attachmentIndex uint) IBlitPassSampleBufferAttachmentDescriptor
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BlitPassSampleBufferAttachmentDescriptorArray */
// Alloc allocates a new instance without initialization.
func (bc _BlitPassSampleBufferAttachmentDescriptorArrayClass) Alloc() BlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[BlitPassSampleBufferAttachmentDescriptorArray](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BlitPassSampleBufferAttachmentDescriptorArrayClass) New() BlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[BlitPassSampleBufferAttachmentDescriptorArray](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BlitPassSampleBufferAttachmentDescriptorArray) Init() BlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[BlitPassSampleBufferAttachmentDescriptorArray](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BlitPassSampleBufferAttachmentDescriptorArray) Autorelease() BlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[BlitPassSampleBufferAttachmentDescriptorArray](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBlitPassSampleBufferAttachmentDescriptorArray creates a new BlitPassSampleBufferAttachmentDescriptorArray instance.
func NewBlitPassSampleBufferAttachmentDescriptorArray() BlitPassSampleBufferAttachmentDescriptorArray {
	return getBlitPassSampleBufferAttachmentDescriptorArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BlitPassSampleBufferAttachmentDescriptorArray */
// A container that stores an array of sample buffer attachments for a blit pass.
//
// The number of elements in the array is at least the number of elements in an instance’s property.


// A container that stores an array of sample buffer attachments for a blit pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptorArray
type BlitPassSampleBufferAttachmentDescriptorArray struct {
	objectivec.Object
}

// BlitPassSampleBufferAttachmentDescriptorArrayFrom constructs a [BlitPassSampleBufferAttachmentDescriptorArray] from an unsafe.Pointer.
//
// A container that stores an array of sample buffer attachments for a blit pass.
func BlitPassSampleBufferAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) BlitPassSampleBufferAttachmentDescriptorArray {
	return BlitPassSampleBufferAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BlitPassSampleBufferAttachmentDescriptorArray *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BlitPassSampleBufferAttachmentDescriptorArray */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BlitPassSampleBufferAttachmentDescriptorArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BlitPassSampleBufferAttachmentDescriptorArray */

// Copies the properties of a blit pass sample buffer attachment descriptor instance to the properties of one of the array’s instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (b_ BlitPassSampleBufferAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLBlitPassSampleBufferAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}/* debug [instance_methods/method]: SetObjectAtIndexedSubscript */


// Accesses one of the array’s blit pass sample buffer attachment descriptor instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptorArray/subscript(_:)
func (b_ BlitPassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IBlitPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[BlitPassSampleBufferAttachmentDescriptor](b_.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BlitPassSampleBufferAttachmentDescriptorArray */

// The counter sets supported by the device object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/countersets
func (b_ BlitPassSampleBufferAttachmentDescriptorArray) CounterSets() CounterSet /* not a class type */ {
	rv := objc.Send[CounterSet](b_.ID, objc.Sel("counterSets"))
	return rv
}/* debug [instance_properties/getter]: counterSets */


// The counter sets supported by the device object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevice/countersets
func (b_ BlitPassSampleBufferAttachmentDescriptorArray) SetCounterSets(value CounterSet /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCounterSets:"), value)
}/* debug [instance_properties/setter]: counterSets */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLBlitPassSampleBufferAttachmentDescriptorArray */




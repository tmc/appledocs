// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLRenderPassColorAttachmentDescriptorArray */


/* debug [class_header]: Header for MTLRenderPassColorAttachmentDescriptorArray */
// The class instance for the [RenderPassColorAttachmentDescriptorArray] class.
var (
	RenderPassColorAttachmentDescriptorArrayClass     _RenderPassColorAttachmentDescriptorArrayClass
	RenderPassColorAttachmentDescriptorArrayClassOnce sync.Once
)

func getRenderPassColorAttachmentDescriptorArrayClass() _RenderPassColorAttachmentDescriptorArrayClass {
	RenderPassColorAttachmentDescriptorArrayClassOnce.Do(func() {
		RenderPassColorAttachmentDescriptorArrayClass = _RenderPassColorAttachmentDescriptorArrayClass{objc.GetClass("MTLRenderPassColorAttachmentDescriptorArray")}
	})
	return RenderPassColorAttachmentDescriptorArrayClass
}

type _RenderPassColorAttachmentDescriptorArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RenderPassColorAttachmentDescriptorArray */
// An interface definition for the [RenderPassColorAttachmentDescriptorArray] class.
type IRenderPassColorAttachmentDescriptorArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RenderPassColorAttachmentDescriptorArray */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RenderPassColorAttachmentDescriptorArray */
	// methods:
	SetObjectAtIndexedSubscript(attachment IMTLRenderPassColorAttachmentDescriptor, attachmentIndex uint)
	ObjectAtIndexedSubscript(attachmentIndex uint) IRenderPassColorAttachmentDescriptor
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RenderPassColorAttachmentDescriptorArray */
// Alloc allocates a new instance without initialization.
func (rc _RenderPassColorAttachmentDescriptorArrayClass) Alloc() RenderPassColorAttachmentDescriptorArray {
	rv := objc.Send[RenderPassColorAttachmentDescriptorArray](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RenderPassColorAttachmentDescriptorArrayClass) New() RenderPassColorAttachmentDescriptorArray {
	rv := objc.Send[RenderPassColorAttachmentDescriptorArray](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RenderPassColorAttachmentDescriptorArray) Init() RenderPassColorAttachmentDescriptorArray {
	rv := objc.Send[RenderPassColorAttachmentDescriptorArray](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RenderPassColorAttachmentDescriptorArray) Autorelease() RenderPassColorAttachmentDescriptorArray {
	rv := objc.Send[RenderPassColorAttachmentDescriptorArray](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRenderPassColorAttachmentDescriptorArray creates a new RenderPassColorAttachmentDescriptorArray instance.
func NewRenderPassColorAttachmentDescriptorArray() RenderPassColorAttachmentDescriptorArray {
	return getRenderPassColorAttachmentDescriptorArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RenderPassColorAttachmentDescriptorArray */
// An array of render pass color attachment descriptor objects.


// An array of render pass color attachment descriptor objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassColorAttachmentDescriptorArray
type RenderPassColorAttachmentDescriptorArray struct {
	objectivec.Object
}

// RenderPassColorAttachmentDescriptorArrayFrom constructs a [RenderPassColorAttachmentDescriptorArray] from an unsafe.Pointer.
//
// An array of render pass color attachment descriptor objects.
func RenderPassColorAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) RenderPassColorAttachmentDescriptorArray {
	return RenderPassColorAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RenderPassColorAttachmentDescriptorArray *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RenderPassColorAttachmentDescriptorArray */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RenderPassColorAttachmentDescriptorArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RenderPassColorAttachmentDescriptorArray */

// Sets the descriptor for the specified color attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassColorAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (r_ RenderPassColorAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLRenderPassColorAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}/* debug [instance_methods/method]: SetObjectAtIndexedSubscript */


// Returns the descriptor object for the specified color attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPassColorAttachmentDescriptorArray/subscript(_:)
func (r_ RenderPassColorAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IRenderPassColorAttachmentDescriptor {
	rv := objc.Send[RenderPassColorAttachmentDescriptor](r_.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RenderPassColorAttachmentDescriptorArray */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLRenderPassColorAttachmentDescriptorArray */




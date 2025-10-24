// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4RenderPipelineColorAttachmentDescriptorArray */


/* debug [class_header]: Header for MTL4RenderPipelineColorAttachmentDescriptorArray */
// The class instance for the [MTL4RenderPipelineColorAttachmentDescriptorArray] class.
var (
	MTL4RenderPipelineColorAttachmentDescriptorArrayClass     _MTL4RenderPipelineColorAttachmentDescriptorArrayClass
	MTL4RenderPipelineColorAttachmentDescriptorArrayClassOnce sync.Once
)

func getMTL4RenderPipelineColorAttachmentDescriptorArrayClass() _MTL4RenderPipelineColorAttachmentDescriptorArrayClass {
	MTL4RenderPipelineColorAttachmentDescriptorArrayClassOnce.Do(func() {
		MTL4RenderPipelineColorAttachmentDescriptorArrayClass = _MTL4RenderPipelineColorAttachmentDescriptorArrayClass{objc.GetClass("MTL4RenderPipelineColorAttachmentDescriptorArray")}
	})
	return MTL4RenderPipelineColorAttachmentDescriptorArrayClass
}

type _MTL4RenderPipelineColorAttachmentDescriptorArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4RenderPipelineColorAttachmentDescriptorArray */
// An interface definition for the [MTL4RenderPipelineColorAttachmentDescriptorArray] class.
type IMTL4RenderPipelineColorAttachmentDescriptorArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4RenderPipelineColorAttachmentDescriptorArray */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4RenderPipelineColorAttachmentDescriptorArray */
	// methods:
	Reset()
	SetObjectAtIndexedSubscript(attachment IMTL4RenderPipelineColorAttachmentDescriptor, attachmentIndex uint)
	ObjectAtIndexedSubscript(attachmentIndex uint) IMTL4RenderPipelineColorAttachmentDescriptor
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4RenderPipelineColorAttachmentDescriptorArray */
// Alloc allocates a new instance without initialization.
func (mc _MTL4RenderPipelineColorAttachmentDescriptorArrayClass) Alloc() MTL4RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4RenderPipelineColorAttachmentDescriptorArrayClass) New() MTL4RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptorArray](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4RenderPipelineColorAttachmentDescriptorArray) Init() MTL4RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptorArray](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4RenderPipelineColorAttachmentDescriptorArray) Autorelease() MTL4RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptorArray](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4RenderPipelineColorAttachmentDescriptorArray creates a new MTL4RenderPipelineColorAttachmentDescriptorArray instance.
func NewMTL4RenderPipelineColorAttachmentDescriptorArray() MTL4RenderPipelineColorAttachmentDescriptorArray {
	return getMTL4RenderPipelineColorAttachmentDescriptorArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4RenderPipelineColorAttachmentDescriptorArray */
// An array of color attachment descriptions for a render pipeline.


// An array of color attachment descriptions for a render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptorArray
type MTL4RenderPipelineColorAttachmentDescriptorArray struct {
	objectivec.Object
}

// MTL4RenderPipelineColorAttachmentDescriptorArrayFrom constructs a [MTL4RenderPipelineColorAttachmentDescriptorArray] from an unsafe.Pointer.
//
// An array of color attachment descriptions for a render pipeline.
func MTL4RenderPipelineColorAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) MTL4RenderPipelineColorAttachmentDescriptorArray {
	return MTL4RenderPipelineColorAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4RenderPipelineColorAttachmentDescriptorArray *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4RenderPipelineColorAttachmentDescriptorArray */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4RenderPipelineColorAttachmentDescriptorArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4RenderPipelineColorAttachmentDescriptorArray */

// Resets the elements of the descriptor array
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptorArray/reset()
func (m_ MTL4RenderPipelineColorAttachmentDescriptorArray) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */


// Sets an attachment at an index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (m_ MTL4RenderPipelineColorAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTL4RenderPipelineColorAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}/* debug [instance_methods/method]: SetObjectAtIndexedSubscript */


// Accesses a color attachment at a specific index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptorArray/subscript(_:)
func (m_ MTL4RenderPipelineColorAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IMTL4RenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptor](m_.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4RenderPipelineColorAttachmentDescriptorArray */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4RenderPipelineColorAttachmentDescriptorArray */




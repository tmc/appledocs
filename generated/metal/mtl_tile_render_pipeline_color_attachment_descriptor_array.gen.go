// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLTileRenderPipelineColorAttachmentDescriptorArray */


/* debug [class_header]: Header for MTLTileRenderPipelineColorAttachmentDescriptorArray */
// The class instance for the [TileRenderPipelineColorAttachmentDescriptorArray] class.
var (
	TileRenderPipelineColorAttachmentDescriptorArrayClass     _TileRenderPipelineColorAttachmentDescriptorArrayClass
	TileRenderPipelineColorAttachmentDescriptorArrayClassOnce sync.Once
)

func getTileRenderPipelineColorAttachmentDescriptorArrayClass() _TileRenderPipelineColorAttachmentDescriptorArrayClass {
	TileRenderPipelineColorAttachmentDescriptorArrayClassOnce.Do(func() {
		TileRenderPipelineColorAttachmentDescriptorArrayClass = _TileRenderPipelineColorAttachmentDescriptorArrayClass{objc.GetClass("MTLTileRenderPipelineColorAttachmentDescriptorArray")}
	})
	return TileRenderPipelineColorAttachmentDescriptorArrayClass
}

type _TileRenderPipelineColorAttachmentDescriptorArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TileRenderPipelineColorAttachmentDescriptorArray */
// An interface definition for the [TileRenderPipelineColorAttachmentDescriptorArray] class.
type ITileRenderPipelineColorAttachmentDescriptorArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TileRenderPipelineColorAttachmentDescriptorArray */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TileRenderPipelineColorAttachmentDescriptorArray */
	// methods:
	SetObjectAtIndexedSubscript(attachment IMTLTileRenderPipelineColorAttachmentDescriptor, attachmentIndex uint)
	ObjectAtIndexedSubscript(attachmentIndex uint) ITileRenderPipelineColorAttachmentDescriptor
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TileRenderPipelineColorAttachmentDescriptorArray */
// Alloc allocates a new instance without initialization.
func (tc _TileRenderPipelineColorAttachmentDescriptorArrayClass) Alloc() TileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptorArray](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TileRenderPipelineColorAttachmentDescriptorArrayClass) New() TileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptorArray](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TileRenderPipelineColorAttachmentDescriptorArray) Init() TileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptorArray](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TileRenderPipelineColorAttachmentDescriptorArray) Autorelease() TileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptorArray](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTileRenderPipelineColorAttachmentDescriptorArray creates a new TileRenderPipelineColorAttachmentDescriptorArray instance.
func NewTileRenderPipelineColorAttachmentDescriptorArray() TileRenderPipelineColorAttachmentDescriptorArray {
	return getTileRenderPipelineColorAttachmentDescriptorArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TileRenderPipelineColorAttachmentDescriptorArray */
// An array of color attachment descriptors for the tile render pipeline.


// An array of color attachment descriptors for the tile render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptorArray
type TileRenderPipelineColorAttachmentDescriptorArray struct {
	objectivec.Object
}

// TileRenderPipelineColorAttachmentDescriptorArrayFrom constructs a [TileRenderPipelineColorAttachmentDescriptorArray] from an unsafe.Pointer.
//
// An array of color attachment descriptors for the tile render pipeline.
func TileRenderPipelineColorAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) TileRenderPipelineColorAttachmentDescriptorArray {
	return TileRenderPipelineColorAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TileRenderPipelineColorAttachmentDescriptorArray *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TileRenderPipelineColorAttachmentDescriptorArray */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TileRenderPipelineColorAttachmentDescriptorArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TileRenderPipelineColorAttachmentDescriptorArray */

// Sets the render pipeline state for a specified color attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (t_ TileRenderPipelineColorAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLTileRenderPipelineColorAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}/* debug [instance_methods/method]: SetObjectAtIndexedSubscript */


// Returns the render pipeline state for the specified color attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptorArray/subscript(_:)
func (t_ TileRenderPipelineColorAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) ITileRenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptor](t_.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TileRenderPipelineColorAttachmentDescriptorArray */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLTileRenderPipelineColorAttachmentDescriptorArray */




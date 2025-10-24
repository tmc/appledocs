// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLTileRenderPipelineColorAttachmentDescriptor */


/* debug [class_header]: Header for MTLTileRenderPipelineColorAttachmentDescriptor */
// The class instance for the [TileRenderPipelineColorAttachmentDescriptor] class.
var (
	TileRenderPipelineColorAttachmentDescriptorClass     _TileRenderPipelineColorAttachmentDescriptorClass
	TileRenderPipelineColorAttachmentDescriptorClassOnce sync.Once
)

func getTileRenderPipelineColorAttachmentDescriptorClass() _TileRenderPipelineColorAttachmentDescriptorClass {
	TileRenderPipelineColorAttachmentDescriptorClassOnce.Do(func() {
		TileRenderPipelineColorAttachmentDescriptorClass = _TileRenderPipelineColorAttachmentDescriptorClass{objc.GetClass("MTLTileRenderPipelineColorAttachmentDescriptor")}
	})
	return TileRenderPipelineColorAttachmentDescriptorClass
}

type _TileRenderPipelineColorAttachmentDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TileRenderPipelineColorAttachmentDescriptor */
// An interface definition for the [TileRenderPipelineColorAttachmentDescriptor] class.
type ITileRenderPipelineColorAttachmentDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TileRenderPipelineColorAttachmentDescriptor */
	// properties:
	PixelFormat() PixelFormat
	SetPixelFormat(value PixelFormat)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TileRenderPipelineColorAttachmentDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TileRenderPipelineColorAttachmentDescriptor */
// Alloc allocates a new instance without initialization.
func (tc _TileRenderPipelineColorAttachmentDescriptorClass) Alloc() TileRenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptor](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TileRenderPipelineColorAttachmentDescriptorClass) New() TileRenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptor](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TileRenderPipelineColorAttachmentDescriptor) Init() TileRenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptor](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TileRenderPipelineColorAttachmentDescriptor) Autorelease() TileRenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptor](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTileRenderPipelineColorAttachmentDescriptor creates a new TileRenderPipelineColorAttachmentDescriptor instance.
func NewTileRenderPipelineColorAttachmentDescriptor() TileRenderPipelineColorAttachmentDescriptor {
	return getTileRenderPipelineColorAttachmentDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TileRenderPipelineColorAttachmentDescriptor */
// A description of a tile-shading render pipeline’s color render target.


// A description of a tile-shading render pipeline’s color render target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptor
type TileRenderPipelineColorAttachmentDescriptor struct {
	objectivec.Object
}

// TileRenderPipelineColorAttachmentDescriptorFrom constructs a [TileRenderPipelineColorAttachmentDescriptor] from an unsafe.Pointer.
//
// A description of a tile-shading render pipeline’s color render target.
func TileRenderPipelineColorAttachmentDescriptorFrom(ptr unsafe.Pointer) TileRenderPipelineColorAttachmentDescriptor {
	return TileRenderPipelineColorAttachmentDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TileRenderPipelineColorAttachmentDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TileRenderPipelineColorAttachmentDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TileRenderPipelineColorAttachmentDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TileRenderPipelineColorAttachmentDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TileRenderPipelineColorAttachmentDescriptor */

// The pixel format associated with the tile shading render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptor/pixelFormat
func (t_ TileRenderPipelineColorAttachmentDescriptor) PixelFormat() PixelFormat {
	rv := objc.Send[PixelFormat](t_.ID, objc.Sel("pixelFormat"))
	return rv
}/* debug [instance_properties/getter]: pixelFormat */


// The pixel format associated with the tile shading render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptor/pixelFormat
func (t_ TileRenderPipelineColorAttachmentDescriptor) SetPixelFormat(value PixelFormat) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPixelFormat:"), value)
}/* debug [instance_properties/setter]: pixelFormat */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLTileRenderPipelineColorAttachmentDescriptor */




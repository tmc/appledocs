// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [TileRenderPipelineColorAttachmentDescriptor] class.
type ITileRenderPipelineColorAttachmentDescriptor interface {
	objectivec.IObject
}

// A description of a tile-shading render pipeline’s color render target.
//
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

// Alloc allocates a new instance without initialization.
func (tc _TileRenderPipelineColorAttachmentDescriptorClass) Alloc() TileRenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptor](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The pixel format associated with the tile shading render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptor/pixelFormat
func (t_ TileRenderPipelineColorAttachmentDescriptor) PixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("pixelFormat"))
	return rv
}


// SetPixelFormat sets the value of the pixelFormat property.
// The pixel format associated with the tile shading render pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptor/pixelFormat
func (t_ TileRenderPipelineColorAttachmentDescriptor) SetPixelFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPixelFormat:"), value)
}



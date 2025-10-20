// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [TiledLayer] class.
var (
	tiledLayerClass     _TiledLayerClass
	tiledLayerClassOnce sync.Once
)

func getTiledLayerClass() _TiledLayerClass {
	tiledLayerClassOnce.Do(func() {
		tiledLayerClass = _TiledLayerClass{objc.GetClass("CATiledLayer")}
	})
	return tiledLayerClass
}

type _TiledLayerClass struct {
	class objc.Class
}

// An interface definition for the [TiledLayer] class.
type ITiledLayer interface {
	ILayer
}

// A layer that provides a way to asynchronously provide tiles of the layer’s content, potentially cached at multiple levels of detail.
//
// As more data is required by the renderer, the layer’s method is called on one or more background threads to supply the drawing operations to fill in one tile of data. The clip bounds and current transformation matrix (CTM) of the drawing context can be used to determine the bounds and resolution of the tile being requested. Regions of the layer may be invalidated using the method however the update will be asynchronous. While the next display update will most likely not contain the updated content, a future update will.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATiledLayer
type TiledLayer struct {
	Layer
}

// TiledLayerFrom constructs a [TiledLayer] from an unsafe.Pointer.
//
// A layer that provides a way to asynchronously provide tiles of the layer’s content, potentially cached at multiple levels of detail.
func TiledLayerFrom(ptr unsafe.Pointer) TiledLayer {
	return TiledLayer{
		Layer: LayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TiledLayerClass) Alloc() TiledLayer {
	rv := objc.Send[TiledLayer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TiledLayerClass) New() TiledLayer {
	rv := objc.Send[TiledLayer](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TiledLayer) Init() TiledLayer {
	rv := objc.Send[TiledLayer](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TiledLayer) Autorelease() TiledLayer {
	rv := objc.Send[TiledLayer](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTiledLayer creates a new TiledLayer instance.
func NewTiledLayer() TiledLayer {
	return getTiledLayerClass().New()
}


// The maximum size of each tile used to create the layer’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATiledLayer/tileSize
func (t_ TiledLayer) TileSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("tileSize"))
	return rv
}


// SetTileSize sets the value of the tileSize property.
// The maximum size of each tile used to create the layer’s content.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATiledLayer/tileSize
func (t_ TiledLayer) SetTileSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTileSize:"), value)
}



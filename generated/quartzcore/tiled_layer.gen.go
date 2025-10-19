// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TiledLayer] class.
var tiledLayerClass = _TiledLayerClass{objc.GetClass("CATiledLayer")}

type _TiledLayerClass struct {
	class objc.Class
}

// An interface definition for the [TiledLayer] class.
type ITiledLayer interface {
	ILayer
}

// A layer that provides a way to asynchronously provide tiles of the layer’s content, potentially cached at multiple levels of detail. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return tiledLayerClass.New()
}





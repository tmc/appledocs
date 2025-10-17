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




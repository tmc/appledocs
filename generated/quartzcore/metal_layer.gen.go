// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MetalLayer] class.
var metalLayerClass = _MetalLayerClass{objc.GetClass("CAMetalLayer")}

type _MetalLayerClass struct {
	class objc.Class
}

// A Core Animation layer that Metal can render into, typically displayed onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer

type MetalLayer struct {
	Layer
}

// MetalLayerFrom constructs a [MetalLayer] from an unsafe.Pointer.
//
// A Core Animation layer that Metal can render into, typically displayed onscreen.
func MetalLayerFrom(ptr unsafe.Pointer) MetalLayer {
	return MetalLayer{
		Layer: LayerFrom(ptr),
	}
}




// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [EmitterLayer] class.
var emitterLayerClass = _EmitterLayerClass{objc.GetClass("CAEmitterLayer")}

type _EmitterLayerClass struct {
	class objc.Class
}

// A layer that emits, animates, and renders a particle system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer

type EmitterLayer struct {
	Layer
}

// EmitterLayerFrom constructs a [EmitterLayer] from an unsafe.Pointer.
//
// A layer that emits, animates, and renders a particle system.
func EmitterLayerFrom(ptr unsafe.Pointer) EmitterLayer {
	return EmitterLayer{
		Layer: LayerFrom(ptr),
	}
}




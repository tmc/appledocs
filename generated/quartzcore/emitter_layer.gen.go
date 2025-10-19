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

// An interface definition for the [EmitterLayer] class.
type IEmitterLayer interface {
	ILayer
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
// Alloc allocates a new instance without initialization.
func (ec _EmitterLayerClass) Alloc() EmitterLayer {
	rv := objc.Send[EmitterLayer](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ec _EmitterLayerClass) New() EmitterLayer {
	rv := objc.Send[EmitterLayer](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EmitterLayer) Init() EmitterLayer {
	rv := objc.Send[EmitterLayer](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EmitterLayer) Autorelease() EmitterLayer {
	rv := objc.Send[EmitterLayer](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEmitterLayer creates a new EmitterLayer instance.
func NewEmitterLayer() EmitterLayer {
	return emitterLayerClass.New()
}





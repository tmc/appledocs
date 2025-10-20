// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [EmitterLayer] class.
var (
	emitterLayerClass     _EmitterLayerClass
	emitterLayerClassOnce sync.Once
)

func getEmitterLayerClass() _EmitterLayerClass {
	emitterLayerClassOnce.Do(func() {
		emitterLayerClass = _EmitterLayerClass{objc.GetClass("CAEmitterLayer")}
	})
	return emitterLayerClass
}

type _EmitterLayerClass struct {
	class objc.Class
}

// An interface definition for the [EmitterLayer] class.
type IEmitterLayer interface {
	ILayer
}

// A layer that emits, animates, and renders a particle system.
//
// The particles, defined by instances of , are drawn above the layer’s background color and border. The following code shows how to set up a simple point (the default is ) particle emitter. It uses an image named as the cell contents and, by setting the emitter cell’s to doc://com.apple.documentation/documentation/corefoundation/cgfloat/1845230-pi , the particles are emitted in all directions.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getEmitterLayerClass().New()
}


// Specifies the emitter shape.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterShape
func (e_ EmitterLayer) EmitterShape() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("emitterShape"))
	return rv
}

// SetEmitterShape sets the value of the emitterShape property.
// Specifies the emitter shape.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterShape
func (e_ EmitterLayer) SetEmitterShape(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmitterShape:"), value)
}



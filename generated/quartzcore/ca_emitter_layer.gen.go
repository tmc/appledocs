// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [EmitterLayer] class.
var (
	EmitterLayerClass     _EmitterLayerClass
	EmitterLayerClassOnce sync.Once
)

func getEmitterLayerClass() _EmitterLayerClass {
	EmitterLayerClassOnce.Do(func() {
		EmitterLayerClass = _EmitterLayerClass{objc.GetClass("CAEmitterLayer")}
	})
	return EmitterLayerClass
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


// Defines a multiplier that is applied to the cell-defined birth rate. Animatable
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/birthRate
func (e_ EmitterLayer) BirthRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("birthRate"))
	return rv
}


// SetBirthRate sets the value of the birthRate property.
// Defines a multiplier that is applied to the cell-defined birth rate. Animatable

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/birthRate
func (e_ EmitterLayer) SetBirthRate(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setBirthRate:"), value)
}

// The array emitter cells attached to the layer.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterCells
func (e_ EmitterLayer) EmitterCells() []EmitterCell {
	rv := objc.Send[[]EmitterCell](e_.ID, objc.Sel("emitterCells"))
	return rv
}


// SetEmitterCells sets the value of the emitterCells property.
// The array emitter cells attached to the layer.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterCells
func (e_ EmitterLayer) SetEmitterCells(value []EmitterCell) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmitterCells:"), nsArray)
}

// Determines the depth of the emitter shape.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterDepth
func (e_ EmitterLayer) EmitterDepth() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("emitterDepth"))
	return rv
}


// SetEmitterDepth sets the value of the emitterDepth property.
// Determines the depth of the emitter shape.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterDepth
func (e_ EmitterLayer) SetEmitterDepth(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmitterDepth:"), value)
}

// Specifies the emitter mode.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterMode
func (e_ EmitterLayer) EmitterMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("emitterMode"))
	return rv
}


// SetEmitterMode sets the value of the emitterMode property.
// Specifies the emitter mode.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterMode
func (e_ EmitterLayer) SetEmitterMode(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmitterMode:"), value)
}

// The position of the center of the particle emitter. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterPosition
func (e_ EmitterLayer) EmitterPosition() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](e_.ID, objc.Sel("emitterPosition"))
	return rv
}


// SetEmitterPosition sets the value of the emitterPosition property.
// The position of the center of the particle emitter. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterPosition
func (e_ EmitterLayer) SetEmitterPosition(value coregraphics.CGPoint) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmitterPosition:"), value)
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

// Determines the size of the particle emitter shape. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterSize
func (e_ EmitterLayer) EmitterSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](e_.ID, objc.Sel("emitterSize"))
	return rv
}


// SetEmitterSize sets the value of the emitterSize property.
// Determines the size of the particle emitter shape. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterSize
func (e_ EmitterLayer) SetEmitterSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmitterSize:"), value)
}

// Specifies the center of the particle emitter shape along the z-axis. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterZPosition
func (e_ EmitterLayer) EmitterZPosition() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("emitterZPosition"))
	return rv
}


// SetEmitterZPosition sets the value of the emitterZPosition property.
// Specifies the center of the particle emitter shape along the z-axis. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterZPosition
func (e_ EmitterLayer) SetEmitterZPosition(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmitterZPosition:"), value)
}

// Defines a multiplier applied to the cell-defined lifetime range when particles are created. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/lifetime
func (e_ EmitterLayer) Lifetime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("lifetime"))
	return rv
}


// SetLifetime sets the value of the lifetime property.
// Defines a multiplier applied to the cell-defined lifetime range when particles are created. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/lifetime
func (e_ EmitterLayer) SetLifetime(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLifetime:"), value)
}

// Defines whether the layer flattens the particles into its plane.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/preservesDepth
func (e_ EmitterLayer) PreservesDepth() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("preservesDepth"))
	return rv
}


// SetPreservesDepth sets the value of the preservesDepth property.
// Defines whether the layer flattens the particles into its plane.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/preservesDepth
func (e_ EmitterLayer) SetPreservesDepth(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPreservesDepth:"), value)
}

// Defines how particle cells are rendered into the layer.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/renderMode
func (e_ EmitterLayer) RenderMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("renderMode"))
	return rv
}


// SetRenderMode sets the value of the renderMode property.
// Defines how particle cells are rendered into the layer.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/renderMode
func (e_ EmitterLayer) SetRenderMode(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRenderMode:"), value)
}

// Defines a multiplier applied to the cell-defined particle scale.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/scale
func (e_ EmitterLayer) Scale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("scale"))
	return rv
}


// SetScale sets the value of the scale property.
// Defines a multiplier applied to the cell-defined particle scale.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/scale
func (e_ EmitterLayer) SetScale(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setScale:"), value)
}

// Specifies the seed used to initialize the random number generator.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/seed
func (e_ EmitterLayer) Seed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("seed"))
	return rv
}


// SetSeed sets the value of the seed property.
// Specifies the seed used to initialize the random number generator.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/seed
func (e_ EmitterLayer) SetSeed(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSeed:"), value)
}

// Defines a multiplier applied to the cell-defined particle spin. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/spin
func (e_ EmitterLayer) Spin() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("spin"))
	return rv
}


// SetSpin sets the value of the spin property.
// Defines a multiplier applied to the cell-defined particle spin. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/spin
func (e_ EmitterLayer) SetSpin(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSpin:"), value)
}

// Defines a multiplier applied to the cell-defined particle velocity. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/velocity
func (e_ EmitterLayer) Velocity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("velocity"))
	return rv
}


// SetVelocity sets the value of the velocity property.
// Defines a multiplier applied to the cell-defined particle velocity. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/velocity
func (e_ EmitterLayer) SetVelocity(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setVelocity:"), value)
}

// The angle, in radians, defining a cone around the emission angle. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/emissionrange
func (e_ EmitterLayer) EmissionRange() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("emissionRange"))
	return rv
}


// SetEmissionRange sets the value of the emissionRange property.
// The angle, in radians, defining a cone around the emission angle. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/emissionrange
func (e_ EmitterLayer) SetEmissionRange(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmissionRange:"), value)
}




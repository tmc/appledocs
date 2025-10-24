// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	// properties:
	EmitterShape() EmitterLayerEmitterShape /* not a class type */
	SetEmitterShape(value EmitterLayerEmitterShape /* not a class type */)
	EmissionRange() float64
	SetEmissionRange(value float64)
	BirthRate() float32
	SetBirthRate(value float32)
	EmitterCells() IEmitterCell
	SetEmitterCells(value IEmitterCell)
	EmitterDepth() float64
	SetEmitterDepth(value float64)
	EmitterMode() EmitterLayerEmitterMode /* not a class type */
	SetEmitterMode(value EmitterLayerEmitterMode /* not a class type */)
	EmitterPosition() objc.IObject /* cross-framework: Point */
	SetEmitterPosition(value objc.IObject /* cross-framework: Point */)
	EmitterSize() objc.IObject /* cross-framework: Size */
	SetEmitterSize(value objc.IObject /* cross-framework: Size */)
	EmitterZPosition() float64
	SetEmitterZPosition(value float64)
	Lifetime() float32
	SetLifetime(value float32)
	PreservesDepth() bool
	SetPreservesDepth(value bool)
	RenderMode() EmitterLayerRenderMode /* not a class type */
	SetRenderMode(value EmitterLayerRenderMode /* not a class type */)
	Scale() float32
	SetScale(value float32)
	Seed() unsafe.Pointer
	SetSeed(value unsafe.Pointer)
	Spin() float32
	SetSpin(value float32)
	Velocity() float32
	SetVelocity(value float32)
	// methods:
}

// A layer that emits, animates, and renders a particle system.
//
// The particles, defined by instances of , are drawn above the layer’s background color and border. The following code shows how to set up a simple point (the default is ) particle emitter. It uses an image named as the cell contents and, by setting the emitter cell’s to doc://com.apple.documentation/documentation/corefoundation/cgfloat/1845230-pi , the particles are emitted in all directions.


// A layer that emits, animates, and renders a particle system.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterShape
func (e_ EmitterLayer) EmitterShape() EmitterLayerEmitterShape /* not a class type */ {
	rv := objc.Send[EmitterLayerEmitterShape](e_.ID, objc.Sel("emitterShape"))
	return rv
}


// Specifies the emitter shape.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterShape
func (e_ EmitterLayer) SetEmitterShape(value EmitterLayerEmitterShape /* not a class type */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmitterShape:"), value)
}


// The angle, in radians, defining a cone around the emission angle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/emissionrange
func (e_ EmitterLayer) EmissionRange() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("emissionRange"))
	return rv
}


// The angle, in radians, defining a cone around the emission angle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/emissionrange
func (e_ EmitterLayer) SetEmissionRange(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmissionRange:"), value)
}


// Defines a multiplier that is applied to the cell-defined birth rate. Animatable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/birthrate
func (e_ EmitterLayer) BirthRate() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("birthRate"))
	return rv
}


// Defines a multiplier that is applied to the cell-defined birth rate. Animatable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/birthrate
func (e_ EmitterLayer) SetBirthRate(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setBirthRate:"), value)
}


// The array emitter cells attached to the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/emittercells
func (e_ EmitterLayer) EmitterCells() IEmitterCell {
	rv := objc.Send[EmitterCell](e_.ID, objc.Sel("emitterCells"))
	return rv
}


// The array emitter cells attached to the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/emittercells
func (e_ EmitterLayer) SetEmitterCells(value IEmitterCell) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmitterCells:"), value)
}


// Determines the depth of the emitter shape.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/emitterdepth
func (e_ EmitterLayer) EmitterDepth() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("emitterDepth"))
	return rv
}


// Determines the depth of the emitter shape.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/emitterdepth
func (e_ EmitterLayer) SetEmitterDepth(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmitterDepth:"), value)
}


// Specifies the emitter mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/emittermode
func (e_ EmitterLayer) EmitterMode() EmitterLayerEmitterMode /* not a class type */ {
	rv := objc.Send[EmitterLayerEmitterMode](e_.ID, objc.Sel("emitterMode"))
	return rv
}


// Specifies the emitter mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/emittermode
func (e_ EmitterLayer) SetEmitterMode(value EmitterLayerEmitterMode /* not a class type */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmitterMode:"), value)
}


// The position of the center of the particle emitter. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/emitterposition
func (e_ EmitterLayer) EmitterPosition() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[corefoundation.Point](e_.ID, objc.Sel("emitterPosition"))
	return rv
}


// The position of the center of the particle emitter. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/emitterposition
func (e_ EmitterLayer) SetEmitterPosition(value objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmitterPosition:"), value)
}


// Determines the size of the particle emitter shape. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/emittersize
func (e_ EmitterLayer) EmitterSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](e_.ID, objc.Sel("emitterSize"))
	return rv
}


// Determines the size of the particle emitter shape. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/emittersize
func (e_ EmitterLayer) SetEmitterSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmitterSize:"), value)
}


// Specifies the center of the particle emitter shape along the z-axis. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/emitterzposition
func (e_ EmitterLayer) EmitterZPosition() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("emitterZPosition"))
	return rv
}


// Specifies the center of the particle emitter shape along the z-axis. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/emitterzposition
func (e_ EmitterLayer) SetEmitterZPosition(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmitterZPosition:"), value)
}


// Defines a multiplier applied to the cell-defined lifetime range when particles are created. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/lifetime
func (e_ EmitterLayer) Lifetime() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("lifetime"))
	return rv
}


// Defines a multiplier applied to the cell-defined lifetime range when particles are created. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/lifetime
func (e_ EmitterLayer) SetLifetime(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLifetime:"), value)
}


// Defines whether the layer flattens the particles into its plane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/preservesdepth
func (e_ EmitterLayer) PreservesDepth() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("preservesDepth"))
	return rv
}


// Defines whether the layer flattens the particles into its plane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/preservesdepth
func (e_ EmitterLayer) SetPreservesDepth(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPreservesDepth:"), value)
}


// Defines how particle cells are rendered into the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/rendermode
func (e_ EmitterLayer) RenderMode() EmitterLayerRenderMode /* not a class type */ {
	rv := objc.Send[EmitterLayerRenderMode](e_.ID, objc.Sel("renderMode"))
	return rv
}


// Defines how particle cells are rendered into the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/rendermode
func (e_ EmitterLayer) SetRenderMode(value EmitterLayerRenderMode /* not a class type */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRenderMode:"), value)
}


// Defines a multiplier applied to the cell-defined particle scale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/scale
func (e_ EmitterLayer) Scale() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("scale"))
	return rv
}


// Defines a multiplier applied to the cell-defined particle scale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/scale
func (e_ EmitterLayer) SetScale(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setScale:"), value)
}


// Specifies the seed used to initialize the random number generator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/seed
func (e_ EmitterLayer) Seed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("seed"))
	return rv
}


// Specifies the seed used to initialize the random number generator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/seed
func (e_ EmitterLayer) SetSeed(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSeed:"), value)
}


// Defines a multiplier applied to the cell-defined particle spin. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/spin
func (e_ EmitterLayer) Spin() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("spin"))
	return rv
}


// Defines a multiplier applied to the cell-defined particle spin. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/spin
func (e_ EmitterLayer) SetSpin(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSpin:"), value)
}


// Defines a multiplier applied to the cell-defined particle velocity. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/velocity
func (e_ EmitterLayer) Velocity() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("velocity"))
	return rv
}


// Defines a multiplier applied to the cell-defined particle velocity. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemitterlayer/velocity
func (e_ EmitterLayer) SetVelocity(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setVelocity:"), value)
}




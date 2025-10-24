// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EmitterCell] class.
var (
	EmitterCellClass     _EmitterCellClass
	EmitterCellClassOnce sync.Once
)

func getEmitterCellClass() _EmitterCellClass {
	EmitterCellClassOnce.Do(func() {
		EmitterCellClass = _EmitterCellClass{objc.GetClass("CAEmitterCell")}
	})
	return EmitterCellClass
}

type _EmitterCellClass struct {
	class objc.Class
}

// An interface definition for the [EmitterCell] class.
type IEmitterCell interface {
	objectivec.IObject
	// properties:
	EmissionLatitude() float64
	SetEmissionLatitude(value float64)
	AlphaRange() float32
	SetAlphaRange(value float32)
	AlphaSpeed() float32
	SetAlphaSpeed(value float32)
	BirthRate() float32
	SetBirthRate(value float32)
	BlueRange() float32
	SetBlueRange(value float32)
	BlueSpeed() float32
	SetBlueSpeed(value float32)
	Color() objectivec.IObject
	SetColor(value objectivec.IObject)
	Contents() unsafe.Pointer
	SetContents(value unsafe.Pointer)
	ContentsRect() objc.IObject /* cross-framework: Rect */
	SetContentsRect(value objc.IObject /* cross-framework: Rect */)
	ContentsScale() float64
	SetContentsScale(value float64)
	EmissionLongitude() float64
	SetEmissionLongitude(value float64)
	EmissionRange() float64
	SetEmissionRange(value float64)
	EmitterCells() IEmitterCell
	SetEmitterCells(value IEmitterCell)
	GreenRange() float32
	SetGreenRange(value float32)
	GreenSpeed() float32
	SetGreenSpeed(value float32)
	IsEnabled() bool
	SetIsEnabled(value bool)
	Lifetime() float32
	SetLifetime(value float32)
	LifetimeRange() float32
	SetLifetimeRange(value float32)
	MagnificationFilter() objc.IObject /* cross-framework: NSString */
	SetMagnificationFilter(value objc.IObject /* cross-framework: NSString */)
	MinificationFilter() objc.IObject /* cross-framework: NSString */
	SetMinificationFilter(value objc.IObject /* cross-framework: NSString */)
	MinificationFilterBias() float32
	SetMinificationFilterBias(value float32)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	RedRange() float32
	SetRedRange(value float32)
	RedSpeed() float32
	SetRedSpeed(value float32)
	Scale() float64
	SetScale(value float64)
	ScaleRange() float64
	SetScaleRange(value float64)
	ScaleSpeed() float64
	SetScaleSpeed(value float64)
	Spin() float64
	SetSpin(value float64)
	SpinRange() float64
	SetSpinRange(value float64)
	Style() unsafe.Pointer
	SetStyle(value unsafe.Pointer)
	Velocity() float64
	SetVelocity(value float64)
	VelocityRange() float64
	SetVelocityRange(value float64)
	XAcceleration() float64
	SetXAcceleration(value float64)
	YAcceleration() float64
	SetYAcceleration(value float64)
	ZAcceleration() float64
	SetZAcceleration(value float64)
	// methods:
}

// The definition of a particle emitted by a particle layer.
//
// The class represents one source of particles being emitted by a object. An emitter cell defines the direction and properties of the emitted particles. Emitter cells can have an array of sub-cells, which lets the particles themselves emit particles.


// The definition of a particle emitted by a particle layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell
type EmitterCell struct {
	objectivec.Object
}

// EmitterCellFrom constructs a [EmitterCell] from an unsafe.Pointer.
//
// The definition of a particle emitted by a particle layer.
func EmitterCellFrom(ptr unsafe.Pointer) EmitterCell {
	return EmitterCell{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EmitterCellClass) Alloc() EmitterCell {
	rv := objc.Send[EmitterCell](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EmitterCellClass) New() EmitterCell {
	rv := objc.Send[EmitterCell](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EmitterCell) Init() EmitterCell {
	rv := objc.Send[EmitterCell](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EmitterCell) Autorelease() EmitterCell {
	rv := objc.Send[EmitterCell](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEmitterCell creates a new EmitterCell instance.
func NewEmitterCell() EmitterCell {
	return getEmitterCellClass().New()
}



// The latitudinal orientation of the emission angle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionLatitude
func (e_ EmitterCell) EmissionLatitude() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("emissionLatitude"))
	return rv
}


// The latitudinal orientation of the emission angle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionLatitude
func (e_ EmitterCell) SetEmissionLatitude(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmissionLatitude:"), value)
}


// The amount by which the alpha component of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/alpharange
func (e_ EmitterCell) AlphaRange() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("alphaRange"))
	return rv
}


// The amount by which the alpha component of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/alpharange
func (e_ EmitterCell) SetAlphaRange(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAlphaRange:"), value)
}


// The speed, in seconds, at which the alpha component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/alphaspeed
func (e_ EmitterCell) AlphaSpeed() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("alphaSpeed"))
	return rv
}


// The speed, in seconds, at which the alpha component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/alphaspeed
func (e_ EmitterCell) SetAlphaSpeed(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAlphaSpeed:"), value)
}


// The number of emitted objects created every second. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/birthrate
func (e_ EmitterCell) BirthRate() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("birthRate"))
	return rv
}


// The number of emitted objects created every second. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/birthrate
func (e_ EmitterCell) SetBirthRate(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setBirthRate:"), value)
}


// The amount by which the blue color component of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/bluerange
func (e_ EmitterCell) BlueRange() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("blueRange"))
	return rv
}


// The amount by which the blue color component of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/bluerange
func (e_ EmitterCell) SetBlueRange(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setBlueRange:"), value)
}


// The speed, in seconds, at which the blue color component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/bluespeed
func (e_ EmitterCell) BlueSpeed() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("blueSpeed"))
	return rv
}


// The speed, in seconds, at which the blue color component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/bluespeed
func (e_ EmitterCell) SetBlueSpeed(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setBlueSpeed:"), value)
}


// The color of each emitted object. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/color
func (e_ EmitterCell) Color() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("color"))
	return rv
}


// The color of each emitted object. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/color
func (e_ EmitterCell) SetColor(value objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setColor:"), value)
}


// An object that provides the contents of the layer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/contents
func (e_ EmitterCell) Contents() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("contents"))
	return rv
}


// An object that provides the contents of the layer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/contents
func (e_ EmitterCell) SetContents(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setContents:"), value)
}


// A rectangle (in the unit coordinate space) that specifies the portion of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/contentsrect
func (e_ EmitterCell) ContentsRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](e_.ID, objc.Sel("contentsRect"))
	return rv
}


// A rectangle (in the unit coordinate space) that specifies the portion of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/contentsrect
func (e_ EmitterCell) SetContentsRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setContentsRect:"), value)
}


// The scale factor of the cell contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/contentsscale
func (e_ EmitterCell) ContentsScale() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("contentsScale"))
	return rv
}


// The scale factor of the cell contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/contentsscale
func (e_ EmitterCell) SetContentsScale(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setContentsScale:"), value)
}


// The longitudinal orientation of the emission angle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/emissionlongitude
func (e_ EmitterCell) EmissionLongitude() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("emissionLongitude"))
	return rv
}


// The longitudinal orientation of the emission angle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/emissionlongitude
func (e_ EmitterCell) SetEmissionLongitude(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmissionLongitude:"), value)
}


// The angle, in radians, defining a cone around the emission angle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/emissionrange
func (e_ EmitterCell) EmissionRange() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("emissionRange"))
	return rv
}


// The angle, in radians, defining a cone around the emission angle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/emissionrange
func (e_ EmitterCell) SetEmissionRange(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmissionRange:"), value)
}


// An optional array containing the sub-cells of this cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/emittercells
func (e_ EmitterCell) EmitterCells() IEmitterCell {
	rv := objc.Send[EmitterCell](e_.ID, objc.Sel("emitterCells"))
	return rv
}


// An optional array containing the sub-cells of this cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/emittercells
func (e_ EmitterCell) SetEmitterCells(value IEmitterCell) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmitterCells:"), value)
}


// The amount by which the green color component of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/greenrange
func (e_ EmitterCell) GreenRange() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("greenRange"))
	return rv
}


// The amount by which the green color component of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/greenrange
func (e_ EmitterCell) SetGreenRange(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setGreenRange:"), value)
}


// The speed, in seconds, at which the green color component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/greenspeed
func (e_ EmitterCell) GreenSpeed() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("greenSpeed"))
	return rv
}


// The speed, in seconds, at which the green color component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/greenspeed
func (e_ EmitterCell) SetGreenSpeed(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setGreenSpeed:"), value)
}


// A Boolean value indicating whether or not cells from this emitter are rendered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/isenabled
func (e_ EmitterCell) IsEnabled() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value indicating whether or not cells from this emitter are rendered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/isenabled
func (e_ EmitterCell) SetIsEnabled(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsEnabled:"), value)
}


// The lifetime of the cell, in seconds. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/lifetime
func (e_ EmitterCell) Lifetime() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("lifetime"))
	return rv
}


// The lifetime of the cell, in seconds. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/lifetime
func (e_ EmitterCell) SetLifetime(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLifetime:"), value)
}


// The mean value by which the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/lifetimerange
func (e_ EmitterCell) LifetimeRange() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("lifetimeRange"))
	return rv
}


// The mean value by which the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/lifetimerange
func (e_ EmitterCell) SetLifetimeRange(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLifetimeRange:"), value)
}


// The filter used when increasing the size of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/magnificationfilter
func (e_ EmitterCell) MagnificationFilter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("magnificationFilter"))
	return rv
}


// The filter used when increasing the size of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/magnificationfilter
func (e_ EmitterCell) SetMagnificationFilter(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMagnificationFilter:"), value)
}


// The filter used when reducing the size of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/minificationfilter
func (e_ EmitterCell) MinificationFilter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("minificationFilter"))
	return rv
}


// The filter used when reducing the size of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/minificationfilter
func (e_ EmitterCell) SetMinificationFilter(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMinificationFilter:"), value)
}


// The bias factor used by the minification filter to determine the levels of detail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/minificationfilterbias
func (e_ EmitterCell) MinificationFilterBias() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("minificationFilterBias"))
	return rv
}


// The bias factor used by the minification filter to determine the levels of detail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/minificationfilterbias
func (e_ EmitterCell) SetMinificationFilterBias(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMinificationFilterBias:"), value)
}


// The name of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/name
func (e_ EmitterCell) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("name"))
	return rv
}


// The name of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/name
func (e_ EmitterCell) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setName:"), value)
}


// The amount by which the red color component of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/redrange
func (e_ EmitterCell) RedRange() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("redRange"))
	return rv
}


// The amount by which the red color component of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/redrange
func (e_ EmitterCell) SetRedRange(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRedRange:"), value)
}


// The speed, in seconds, at which the red color component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/redspeed
func (e_ EmitterCell) RedSpeed() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("redSpeed"))
	return rv
}


// The speed, in seconds, at which the red color component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/redspeed
func (e_ EmitterCell) SetRedSpeed(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRedSpeed:"), value)
}


// Specifies the scale factor applied to the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/scale
func (e_ EmitterCell) Scale() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("scale"))
	return rv
}


// Specifies the scale factor applied to the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/scale
func (e_ EmitterCell) SetScale(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setScale:"), value)
}


// Specifies the range over which the scale value can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/scalerange
func (e_ EmitterCell) ScaleRange() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("scaleRange"))
	return rv
}


// Specifies the range over which the scale value can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/scalerange
func (e_ EmitterCell) SetScaleRange(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setScaleRange:"), value)
}


// The speed at which the scale changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/scalespeed
func (e_ EmitterCell) ScaleSpeed() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("scaleSpeed"))
	return rv
}


// The speed at which the scale changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/scalespeed
func (e_ EmitterCell) SetScaleSpeed(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setScaleSpeed:"), value)
}


// The rotational velocity, measured in radians per second, to apply to the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/spin
func (e_ EmitterCell) Spin() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("spin"))
	return rv
}


// The rotational velocity, measured in radians per second, to apply to the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/spin
func (e_ EmitterCell) SetSpin(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSpin:"), value)
}


// The amount by which the spin of the cell can vary over its lifetime. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/spinrange
func (e_ EmitterCell) SpinRange() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("spinRange"))
	return rv
}


// The amount by which the spin of the cell can vary over its lifetime. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/spinrange
func (e_ EmitterCell) SetSpinRange(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSpinRange:"), value)
}


// An optional dictionary containing additional style values that are not explicitly defined by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/style
func (e_ EmitterCell) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("style"))
	return rv
}


// An optional dictionary containing additional style values that are not explicitly defined by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/style
func (e_ EmitterCell) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setStyle:"), value)
}


// The initial velocity of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/velocity
func (e_ EmitterCell) Velocity() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("velocity"))
	return rv
}


// The initial velocity of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/velocity
func (e_ EmitterCell) SetVelocity(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setVelocity:"), value)
}


// The amount by which the velocity of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/velocityrange
func (e_ EmitterCell) VelocityRange() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("velocityRange"))
	return rv
}


// The amount by which the velocity of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/velocityrange
func (e_ EmitterCell) SetVelocityRange(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setVelocityRange:"), value)
}


// The x component of an acceleration vector applied to cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/xacceleration
func (e_ EmitterCell) XAcceleration() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("xAcceleration"))
	return rv
}


// The x component of an acceleration vector applied to cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/xacceleration
func (e_ EmitterCell) SetXAcceleration(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setXAcceleration:"), value)
}


// The y component of an acceleration vector applied to cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/yacceleration
func (e_ EmitterCell) YAcceleration() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("yAcceleration"))
	return rv
}


// The y component of an acceleration vector applied to cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/yacceleration
func (e_ EmitterCell) SetYAcceleration(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setYAcceleration:"), value)
}


// The z component of an acceleration vector applied to cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/zacceleration
func (e_ EmitterCell) ZAcceleration() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("zAcceleration"))
	return rv
}


// The z component of an acceleration vector applied to cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/zacceleration
func (e_ EmitterCell) SetZAcceleration(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setZAcceleration:"), value)
}




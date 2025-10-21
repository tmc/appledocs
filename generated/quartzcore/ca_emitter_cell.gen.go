// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	ShouldArchiveValueForKey(key string) bool
}

// The definition of a particle emitted by a particle layer.
//
// The class represents one source of particles being emitted by a object. An emitter cell defines the direction and properties of the emitted particles. Emitter cells can have an array of sub-cells, which lets the particles themselves emit particles.
//
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


// Returns the default value of the property with the specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/defaultValue(forKey:)
func (ec _EmitterCellClass) DefaultValueForKey(key string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("defaultValueForKey:"), objc.String(key))
	return rv
}

// Creates and returns an instance of .
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emitterCell
func (ec _EmitterCellClass) EmitterCell() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("emitterCell"))
	return rv
}

// Returns a Boolean value indicating whether the value for a given key should be archived.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/shouldArchiveValue(forKey:)
func (e_ EmitterCell) ShouldArchiveValueForKey(key string) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("shouldArchiveValueForKey:"), objc.String(key))
	return rv
}

// The amount by which the alpha component of the cell can vary. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/alphaRange
func (e_ EmitterCell) AlphaRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("alphaRange"))
	return rv
}


// SetAlphaRange sets the value of the alphaRange property.
// The amount by which the alpha component of the cell can vary. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/alphaRange
func (e_ EmitterCell) SetAlphaRange(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAlphaRange:"), value)
}

// The speed, in seconds, at which the alpha component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/alphaSpeed
func (e_ EmitterCell) AlphaSpeed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("alphaSpeed"))
	return rv
}


// SetAlphaSpeed sets the value of the alphaSpeed property.
// The speed, in seconds, at which the alpha component changes over the lifetime of the cell. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/alphaSpeed
func (e_ EmitterCell) SetAlphaSpeed(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAlphaSpeed:"), value)
}

// The number of emitted objects created every second. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/birthRate
func (e_ EmitterCell) BirthRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("birthRate"))
	return rv
}


// SetBirthRate sets the value of the birthRate property.
// The number of emitted objects created every second. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/birthRate
func (e_ EmitterCell) SetBirthRate(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setBirthRate:"), value)
}

// The amount by which the blue color component of the cell can vary. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/blueRange
func (e_ EmitterCell) BlueRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("blueRange"))
	return rv
}


// SetBlueRange sets the value of the blueRange property.
// The amount by which the blue color component of the cell can vary. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/blueRange
func (e_ EmitterCell) SetBlueRange(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setBlueRange:"), value)
}

// The speed, in seconds, at which the blue color component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/blueSpeed
func (e_ EmitterCell) BlueSpeed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("blueSpeed"))
	return rv
}


// SetBlueSpeed sets the value of the blueSpeed property.
// The speed, in seconds, at which the blue color component changes over the lifetime of the cell. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/blueSpeed
func (e_ EmitterCell) SetBlueSpeed(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setBlueSpeed:"), value)
}

// The color of each emitted object. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/color
func (e_ EmitterCell) Color() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](e_.ID, objc.Sel("color"))
	return rv
}


// SetColor sets the value of the color property.
// The color of each emitted object. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/color
func (e_ EmitterCell) SetColor(value coregraphics.CGColorRef) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setColor:"), value)
}

// An object that provides the contents of the layer. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/contents
func (e_ EmitterCell) Contents() objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("contents"))
	return rv
}


// SetContents sets the value of the contents property.
// An object that provides the contents of the layer. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/contents
func (e_ EmitterCell) SetContents(value objc.ID) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setContents:"), value)
}

// A rectangle (in the unit coordinate space) that specifies the portion of that the receiver should draw. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/contentsRect
func (e_ EmitterCell) ContentsRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](e_.ID, objc.Sel("contentsRect"))
	return rv
}


// SetContentsRect sets the value of the contentsRect property.
// A rectangle (in the unit coordinate space) that specifies the portion of that the receiver should draw. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/contentsRect
func (e_ EmitterCell) SetContentsRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setContentsRect:"), value)
}

// The scale factor of the cell contents.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/contentsScale
func (e_ EmitterCell) ContentsScale() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("contentsScale"))
	return rv
}


// SetContentsScale sets the value of the contentsScale property.
// The scale factor of the cell contents.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/contentsScale
func (e_ EmitterCell) SetContentsScale(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setContentsScale:"), value)
}

// The latitudinal orientation of the emission angle. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionLatitude
func (e_ EmitterCell) EmissionLatitude() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("emissionLatitude"))
	return rv
}


// SetEmissionLatitude sets the value of the emissionLatitude property.
// The latitudinal orientation of the emission angle. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionLatitude
func (e_ EmitterCell) SetEmissionLatitude(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmissionLatitude:"), value)
}

// The longitudinal orientation of the emission angle. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionLongitude
func (e_ EmitterCell) EmissionLongitude() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("emissionLongitude"))
	return rv
}


// SetEmissionLongitude sets the value of the emissionLongitude property.
// The longitudinal orientation of the emission angle. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionLongitude
func (e_ EmitterCell) SetEmissionLongitude(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmissionLongitude:"), value)
}

// The angle, in radians, defining a cone around the emission angle. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionRange
func (e_ EmitterCell) EmissionRange() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("emissionRange"))
	return rv
}


// SetEmissionRange sets the value of the emissionRange property.
// The angle, in radians, defining a cone around the emission angle. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionRange
func (e_ EmitterCell) SetEmissionRange(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmissionRange:"), value)
}

// An optional array containing the sub-cells of this cell.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emitterCells
func (e_ EmitterCell) EmitterCells() []EmitterCell {
	rv := objc.Send[[]EmitterCell](e_.ID, objc.Sel("emitterCells"))
	return rv
}


// SetEmitterCells sets the value of the emitterCells property.
// An optional array containing the sub-cells of this cell.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emitterCells
func (e_ EmitterCell) SetEmitterCells(value []EmitterCell) {
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

// The amount by which the green color component of the cell can vary. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/greenRange
func (e_ EmitterCell) GreenRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("greenRange"))
	return rv
}


// SetGreenRange sets the value of the greenRange property.
// The amount by which the green color component of the cell can vary. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/greenRange
func (e_ EmitterCell) SetGreenRange(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setGreenRange:"), value)
}

// The speed, in seconds, at which the green color component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/greenSpeed
func (e_ EmitterCell) GreenSpeed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("greenSpeed"))
	return rv
}


// SetGreenSpeed sets the value of the greenSpeed property.
// The speed, in seconds, at which the green color component changes over the lifetime of the cell. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/greenSpeed
func (e_ EmitterCell) SetGreenSpeed(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setGreenSpeed:"), value)
}

// A Boolean value indicating whether or not cells from this emitter are rendered.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/isEnabled
func (e_ EmitterCell) Enabled() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("enabled"))
	return rv
}


// SetEnabled sets the value of the enabled property.
// A Boolean value indicating whether or not cells from this emitter are rendered.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/isEnabled
func (e_ EmitterCell) SetEnabled(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEnabled:"), value)
}

// The lifetime of the cell, in seconds. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/lifetime
func (e_ EmitterCell) Lifetime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("lifetime"))
	return rv
}


// SetLifetime sets the value of the lifetime property.
// The lifetime of the cell, in seconds. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/lifetime
func (e_ EmitterCell) SetLifetime(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLifetime:"), value)
}

// The mean value by which the of the cell can vary. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/lifetimeRange
func (e_ EmitterCell) LifetimeRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("lifetimeRange"))
	return rv
}


// SetLifetimeRange sets the value of the lifetimeRange property.
// The mean value by which the of the cell can vary. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/lifetimeRange
func (e_ EmitterCell) SetLifetimeRange(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLifetimeRange:"), value)
}

// The filter used when increasing the size of the content.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/magnificationFilter
func (e_ EmitterCell) MagnificationFilter() string {
	rv := objc.Send[string](e_.ID, objc.Sel("magnificationFilter"))
	return rv
}


// SetMagnificationFilter sets the value of the magnificationFilter property.
// The filter used when increasing the size of the content.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/magnificationFilter
func (e_ EmitterCell) SetMagnificationFilter(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMagnificationFilter:"), objc.String(value))
}

// The filter used when reducing the size of the content.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/minificationFilter
func (e_ EmitterCell) MinificationFilter() string {
	rv := objc.Send[string](e_.ID, objc.Sel("minificationFilter"))
	return rv
}


// SetMinificationFilter sets the value of the minificationFilter property.
// The filter used when reducing the size of the content.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/minificationFilter
func (e_ EmitterCell) SetMinificationFilter(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMinificationFilter:"), objc.String(value))
}

// The bias factor used by the minification filter to determine the levels of detail.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/minificationFilterBias
func (e_ EmitterCell) MinificationFilterBias() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("minificationFilterBias"))
	return rv
}


// SetMinificationFilterBias sets the value of the minificationFilterBias property.
// The bias factor used by the minification filter to determine the levels of detail.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/minificationFilterBias
func (e_ EmitterCell) SetMinificationFilterBias(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMinificationFilterBias:"), value)
}

// The name of the cell.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/name
func (e_ EmitterCell) Name() string {
	rv := objc.Send[string](e_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the cell.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/name
func (e_ EmitterCell) SetName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setName:"), objc.String(value))
}

// The amount by which the red color component of the cell can vary. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/redRange
func (e_ EmitterCell) RedRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("redRange"))
	return rv
}


// SetRedRange sets the value of the redRange property.
// The amount by which the red color component of the cell can vary. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/redRange
func (e_ EmitterCell) SetRedRange(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRedRange:"), value)
}

// The speed, in seconds, at which the red color component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/redSpeed
func (e_ EmitterCell) RedSpeed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("redSpeed"))
	return rv
}


// SetRedSpeed sets the value of the redSpeed property.
// The speed, in seconds, at which the red color component changes over the lifetime of the cell. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/redSpeed
func (e_ EmitterCell) SetRedSpeed(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRedSpeed:"), value)
}

// Specifies the scale factor applied to the cell. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/scale
func (e_ EmitterCell) Scale() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("scale"))
	return rv
}


// SetScale sets the value of the scale property.
// Specifies the scale factor applied to the cell. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/scale
func (e_ EmitterCell) SetScale(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setScale:"), value)
}

// Specifies the range over which the scale value can vary. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/scaleRange
func (e_ EmitterCell) ScaleRange() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("scaleRange"))
	return rv
}


// SetScaleRange sets the value of the scaleRange property.
// Specifies the range over which the scale value can vary. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/scaleRange
func (e_ EmitterCell) SetScaleRange(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setScaleRange:"), value)
}

// The speed at which the scale changes over the lifetime of the cell. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/scaleSpeed
func (e_ EmitterCell) ScaleSpeed() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("scaleSpeed"))
	return rv
}


// SetScaleSpeed sets the value of the scaleSpeed property.
// The speed at which the scale changes over the lifetime of the cell. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/scaleSpeed
func (e_ EmitterCell) SetScaleSpeed(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setScaleSpeed:"), value)
}

// The rotational velocity, measured in radians per second, to apply to the cell. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/spin
func (e_ EmitterCell) Spin() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("spin"))
	return rv
}


// SetSpin sets the value of the spin property.
// The rotational velocity, measured in radians per second, to apply to the cell. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/spin
func (e_ EmitterCell) SetSpin(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSpin:"), value)
}

// The amount by which the spin of the cell can vary over its lifetime. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/spinRange
func (e_ EmitterCell) SpinRange() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("spinRange"))
	return rv
}


// SetSpinRange sets the value of the spinRange property.
// The amount by which the spin of the cell can vary over its lifetime. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/spinRange
func (e_ EmitterCell) SetSpinRange(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSpinRange:"), value)
}

// An optional dictionary containing additional style values that are not explicitly defined by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/style
func (e_ EmitterCell) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("style"))
	return rv
}


// SetStyle sets the value of the style property.
// An optional dictionary containing additional style values that are not explicitly defined by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/style
func (e_ EmitterCell) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setStyle:"), value)
}

// The initial velocity of the cell. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/velocity
func (e_ EmitterCell) Velocity() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("velocity"))
	return rv
}


// SetVelocity sets the value of the velocity property.
// The initial velocity of the cell. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/velocity
func (e_ EmitterCell) SetVelocity(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setVelocity:"), value)
}

// The amount by which the velocity of the cell can vary. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/velocityRange
func (e_ EmitterCell) VelocityRange() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("velocityRange"))
	return rv
}


// SetVelocityRange sets the value of the velocityRange property.
// The amount by which the velocity of the cell can vary. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/velocityRange
func (e_ EmitterCell) SetVelocityRange(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setVelocityRange:"), value)
}

// The x component of an acceleration vector applied to cell.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/xAcceleration
func (e_ EmitterCell) XAcceleration() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("xAcceleration"))
	return rv
}


// SetXAcceleration sets the value of the xAcceleration property.
// The x component of an acceleration vector applied to cell.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/xAcceleration
func (e_ EmitterCell) SetXAcceleration(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setXAcceleration:"), value)
}

// The y component of an acceleration vector applied to cell.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/yAcceleration
func (e_ EmitterCell) YAcceleration() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("yAcceleration"))
	return rv
}


// SetYAcceleration sets the value of the yAcceleration property.
// The y component of an acceleration vector applied to cell.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/yAcceleration
func (e_ EmitterCell) SetYAcceleration(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setYAcceleration:"), value)
}

// The z component of an acceleration vector applied to cell.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/zAcceleration
func (e_ EmitterCell) ZAcceleration() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("zAcceleration"))
	return rv
}


// SetZAcceleration sets the value of the zAcceleration property.
// The z component of an acceleration vector applied to cell.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/zAcceleration
func (e_ EmitterCell) SetZAcceleration(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setZAcceleration:"), value)
}

// A Boolean value indicating whether or not cells from this emitter are rendered.
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/isenabled
func (e_ EmitterCell) IsEnabled() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isEnabled"))
	return rv
}


// SetIsEnabled sets the value of the isEnabled property.
// A Boolean value indicating whether or not cells from this emitter are rendered.

//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/isenabled
func (e_ EmitterCell) SetIsEnabled(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsEnabled:"), value)
}




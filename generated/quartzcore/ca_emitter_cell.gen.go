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

/* debug [class.gen.go]: Generating class CAEmitterCell */


/* debug [class_header]: Header for CAEmitterCell */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EmitterCell */
// An interface definition for the [EmitterCell] class.
type IEmitterCell interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for EmitterCell */
	// properties:
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
	Color() ColorRef /* not a class type */
	SetColor(value ColorRef /* not a class type */)
	Contents() objc.ID
	SetContents(value objc.ID)
	ContentsRect() corefoundation.CGRect
	SetContentsRect(value corefoundation.CGRect)
	ContentsScale() float64
	SetContentsScale(value float64)
	EmissionLatitude() float64
	SetEmissionLatitude(value float64)
	EmissionLongitude() float64
	SetEmissionLongitude(value float64)
	EmissionRange() float64
	SetEmissionRange(value float64)
	EmitterCells() []EmitterCell
	SetEmitterCells(value []EmitterCell)
	GreenRange() float32
	SetGreenRange(value float32)
	GreenSpeed() float32
	SetGreenSpeed(value float32)
	Enabled() bool
	SetEnabled(value bool)
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
	Style() objc.IObject /* cross-framework: NSDictionary */
	SetStyle(value objc.IObject /* cross-framework: NSDictionary */)
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
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EmitterCell */
	// methods:
	ShouldArchiveValueForKey(key objc.IObject /* cross-framework: NSString */) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EmitterCell */
// Alloc allocates a new instance without initialization.
func (ec _EmitterCellClass) Alloc() EmitterCell {
	rv := objc.Send[EmitterCell](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EmitterCell */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EmitterCell *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EmitterCell */

// Returns the default value of the property with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/defaultValue(forKey:)
func (ec _EmitterCellClass) DefaultValueForKey(key objc.IObject /* cross-framework: NSString */) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("defaultValueForKey:"), key)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultValueForKey) */


// Creates and returns an instance of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emitterCell
func (ec _EmitterCellClass) EmitterCell() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ec.class), objc.Sel("emitterCell"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EmitterCell) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EmitterCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EmitterCell */

// Returns a Boolean value indicating whether the value for a given key should be archived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/shouldArchiveValue(forKey:)
func (e_ EmitterCell) ShouldArchiveValueForKey(key objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("shouldArchiveValueForKey:"), key)
	return rv
}/* debug [instance_methods/method]: ShouldArchiveValueForKey */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EmitterCell */

// The amount by which the alpha component of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/alphaRange
func (e_ EmitterCell) AlphaRange() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("alphaRange"))
	return rv
}/* debug [instance_properties/getter]: alphaRange */


// The amount by which the alpha component of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/alphaRange
func (e_ EmitterCell) SetAlphaRange(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAlphaRange:"), value)
}/* debug [instance_properties/setter]: alphaRange */


// The speed, in seconds, at which the alpha component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/alphaSpeed
func (e_ EmitterCell) AlphaSpeed() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("alphaSpeed"))
	return rv
}/* debug [instance_properties/getter]: alphaSpeed */


// The speed, in seconds, at which the alpha component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/alphaSpeed
func (e_ EmitterCell) SetAlphaSpeed(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAlphaSpeed:"), value)
}/* debug [instance_properties/setter]: alphaSpeed */


// The number of emitted objects created every second. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/birthRate
func (e_ EmitterCell) BirthRate() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("birthRate"))
	return rv
}/* debug [instance_properties/getter]: birthRate */


// The number of emitted objects created every second. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/birthRate
func (e_ EmitterCell) SetBirthRate(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setBirthRate:"), value)
}/* debug [instance_properties/setter]: birthRate */


// The amount by which the blue color component of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/blueRange
func (e_ EmitterCell) BlueRange() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("blueRange"))
	return rv
}/* debug [instance_properties/getter]: blueRange */


// The amount by which the blue color component of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/blueRange
func (e_ EmitterCell) SetBlueRange(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setBlueRange:"), value)
}/* debug [instance_properties/setter]: blueRange */


// The speed, in seconds, at which the blue color component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/blueSpeed
func (e_ EmitterCell) BlueSpeed() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("blueSpeed"))
	return rv
}/* debug [instance_properties/getter]: blueSpeed */


// The speed, in seconds, at which the blue color component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/blueSpeed
func (e_ EmitterCell) SetBlueSpeed(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setBlueSpeed:"), value)
}/* debug [instance_properties/setter]: blueSpeed */


// The color of each emitted object. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/color
func (e_ EmitterCell) Color() ColorRef /* not a class type */ {
	rv := objc.Send[ColorRef](e_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// The color of each emitted object. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/color
func (e_ EmitterCell) SetColor(value ColorRef /* not a class type */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// An object that provides the contents of the layer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/contents
func (e_ EmitterCell) Contents() objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("contents"))
	return rv
}/* debug [instance_properties/getter]: contents */


// An object that provides the contents of the layer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/contents
func (e_ EmitterCell) SetContents(value objc.ID) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setContents:"), value)
}/* debug [instance_properties/setter]: contents */


// A rectangle (in the unit coordinate space) that specifies the portion of that the receiver should draw. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/contentsRect
func (e_ EmitterCell) ContentsRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](e_.ID, objc.Sel("contentsRect"))
	return rv
}/* debug [instance_properties/getter]: contentsRect */


// A rectangle (in the unit coordinate space) that specifies the portion of that the receiver should draw. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/contentsRect
func (e_ EmitterCell) SetContentsRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setContentsRect:"), value)
}/* debug [instance_properties/setter]: contentsRect */


// The scale factor of the cell contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/contentsScale
func (e_ EmitterCell) ContentsScale() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("contentsScale"))
	return rv
}/* debug [instance_properties/getter]: contentsScale */


// The scale factor of the cell contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/contentsScale
func (e_ EmitterCell) SetContentsScale(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setContentsScale:"), value)
}/* debug [instance_properties/setter]: contentsScale */


// The latitudinal orientation of the emission angle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionLatitude
func (e_ EmitterCell) EmissionLatitude() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("emissionLatitude"))
	return rv
}/* debug [instance_properties/getter]: emissionLatitude */


// The latitudinal orientation of the emission angle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionLatitude
func (e_ EmitterCell) SetEmissionLatitude(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmissionLatitude:"), value)
}/* debug [instance_properties/setter]: emissionLatitude */


// The longitudinal orientation of the emission angle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionLongitude
func (e_ EmitterCell) EmissionLongitude() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("emissionLongitude"))
	return rv
}/* debug [instance_properties/getter]: emissionLongitude */


// The longitudinal orientation of the emission angle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionLongitude
func (e_ EmitterCell) SetEmissionLongitude(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmissionLongitude:"), value)
}/* debug [instance_properties/setter]: emissionLongitude */


// The angle, in radians, defining a cone around the emission angle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionRange
func (e_ EmitterCell) EmissionRange() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("emissionRange"))
	return rv
}/* debug [instance_properties/getter]: emissionRange */


// The angle, in radians, defining a cone around the emission angle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionRange
func (e_ EmitterCell) SetEmissionRange(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmissionRange:"), value)
}/* debug [instance_properties/setter]: emissionRange */


// An optional array containing the sub-cells of this cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emitterCells
func (e_ EmitterCell) EmitterCells() []EmitterCell {
	rv := objc.Send[[]EmitterCell](e_.ID, objc.Sel("emitterCells"))
	return rv
}/* debug [instance_properties/getter]: emitterCells */


// An optional array containing the sub-cells of this cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emitterCells
func (e_ EmitterCell) SetEmitterCells(value []EmitterCell) {
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
}/* debug [instance_properties/setter]: emitterCells */


// The amount by which the green color component of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/greenRange
func (e_ EmitterCell) GreenRange() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("greenRange"))
	return rv
}/* debug [instance_properties/getter]: greenRange */


// The amount by which the green color component of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/greenRange
func (e_ EmitterCell) SetGreenRange(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setGreenRange:"), value)
}/* debug [instance_properties/setter]: greenRange */


// The speed, in seconds, at which the green color component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/greenSpeed
func (e_ EmitterCell) GreenSpeed() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("greenSpeed"))
	return rv
}/* debug [instance_properties/getter]: greenSpeed */


// The speed, in seconds, at which the green color component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/greenSpeed
func (e_ EmitterCell) SetGreenSpeed(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setGreenSpeed:"), value)
}/* debug [instance_properties/setter]: greenSpeed */


// A Boolean value indicating whether or not cells from this emitter are rendered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/isEnabled
func (e_ EmitterCell) Enabled() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value indicating whether or not cells from this emitter are rendered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/isEnabled
func (e_ EmitterCell) SetEnabled(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// The lifetime of the cell, in seconds. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/lifetime
func (e_ EmitterCell) Lifetime() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("lifetime"))
	return rv
}/* debug [instance_properties/getter]: lifetime */


// The lifetime of the cell, in seconds. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/lifetime
func (e_ EmitterCell) SetLifetime(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLifetime:"), value)
}/* debug [instance_properties/setter]: lifetime */


// The mean value by which the of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/lifetimeRange
func (e_ EmitterCell) LifetimeRange() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("lifetimeRange"))
	return rv
}/* debug [instance_properties/getter]: lifetimeRange */


// The mean value by which the of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/lifetimeRange
func (e_ EmitterCell) SetLifetimeRange(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLifetimeRange:"), value)
}/* debug [instance_properties/setter]: lifetimeRange */


// The filter used when increasing the size of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/magnificationFilter
func (e_ EmitterCell) MagnificationFilter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("magnificationFilter"))
	return rv
}/* debug [instance_properties/getter]: magnificationFilter */


// The filter used when increasing the size of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/magnificationFilter
func (e_ EmitterCell) SetMagnificationFilter(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMagnificationFilter:"), value)
}/* debug [instance_properties/setter]: magnificationFilter */


// The filter used when reducing the size of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/minificationFilter
func (e_ EmitterCell) MinificationFilter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("minificationFilter"))
	return rv
}/* debug [instance_properties/getter]: minificationFilter */


// The filter used when reducing the size of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/minificationFilter
func (e_ EmitterCell) SetMinificationFilter(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMinificationFilter:"), value)
}/* debug [instance_properties/setter]: minificationFilter */


// The bias factor used by the minification filter to determine the levels of detail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/minificationFilterBias
func (e_ EmitterCell) MinificationFilterBias() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("minificationFilterBias"))
	return rv
}/* debug [instance_properties/getter]: minificationFilterBias */


// The bias factor used by the minification filter to determine the levels of detail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/minificationFilterBias
func (e_ EmitterCell) SetMinificationFilterBias(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMinificationFilterBias:"), value)
}/* debug [instance_properties/setter]: minificationFilterBias */


// The name of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/name
func (e_ EmitterCell) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/name
func (e_ EmitterCell) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// The amount by which the red color component of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/redRange
func (e_ EmitterCell) RedRange() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("redRange"))
	return rv
}/* debug [instance_properties/getter]: redRange */


// The amount by which the red color component of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/redRange
func (e_ EmitterCell) SetRedRange(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRedRange:"), value)
}/* debug [instance_properties/setter]: redRange */


// The speed, in seconds, at which the red color component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/redSpeed
func (e_ EmitterCell) RedSpeed() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("redSpeed"))
	return rv
}/* debug [instance_properties/getter]: redSpeed */


// The speed, in seconds, at which the red color component changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/redSpeed
func (e_ EmitterCell) SetRedSpeed(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRedSpeed:"), value)
}/* debug [instance_properties/setter]: redSpeed */


// Specifies the scale factor applied to the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/scale
func (e_ EmitterCell) Scale() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("scale"))
	return rv
}/* debug [instance_properties/getter]: scale */


// Specifies the scale factor applied to the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/scale
func (e_ EmitterCell) SetScale(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setScale:"), value)
}/* debug [instance_properties/setter]: scale */


// Specifies the range over which the scale value can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/scaleRange
func (e_ EmitterCell) ScaleRange() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("scaleRange"))
	return rv
}/* debug [instance_properties/getter]: scaleRange */


// Specifies the range over which the scale value can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/scaleRange
func (e_ EmitterCell) SetScaleRange(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setScaleRange:"), value)
}/* debug [instance_properties/setter]: scaleRange */


// The speed at which the scale changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/scaleSpeed
func (e_ EmitterCell) ScaleSpeed() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("scaleSpeed"))
	return rv
}/* debug [instance_properties/getter]: scaleSpeed */


// The speed at which the scale changes over the lifetime of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/scaleSpeed
func (e_ EmitterCell) SetScaleSpeed(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setScaleSpeed:"), value)
}/* debug [instance_properties/setter]: scaleSpeed */


// The rotational velocity, measured in radians per second, to apply to the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/spin
func (e_ EmitterCell) Spin() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("spin"))
	return rv
}/* debug [instance_properties/getter]: spin */


// The rotational velocity, measured in radians per second, to apply to the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/spin
func (e_ EmitterCell) SetSpin(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSpin:"), value)
}/* debug [instance_properties/setter]: spin */


// The amount by which the spin of the cell can vary over its lifetime. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/spinRange
func (e_ EmitterCell) SpinRange() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("spinRange"))
	return rv
}/* debug [instance_properties/getter]: spinRange */


// The amount by which the spin of the cell can vary over its lifetime. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/spinRange
func (e_ EmitterCell) SetSpinRange(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSpinRange:"), value)
}/* debug [instance_properties/setter]: spinRange */


// An optional dictionary containing additional style values that are not explicitly defined by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/style
func (e_ EmitterCell) Style() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](e_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */


// An optional dictionary containing additional style values that are not explicitly defined by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/style
func (e_ EmitterCell) SetStyle(value objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setStyle:"), value)
}/* debug [instance_properties/setter]: style */


// The initial velocity of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/velocity
func (e_ EmitterCell) Velocity() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("velocity"))
	return rv
}/* debug [instance_properties/getter]: velocity */


// The initial velocity of the cell. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/velocity
func (e_ EmitterCell) SetVelocity(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setVelocity:"), value)
}/* debug [instance_properties/setter]: velocity */


// The amount by which the velocity of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/velocityRange
func (e_ EmitterCell) VelocityRange() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("velocityRange"))
	return rv
}/* debug [instance_properties/getter]: velocityRange */


// The amount by which the velocity of the cell can vary. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/velocityRange
func (e_ EmitterCell) SetVelocityRange(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setVelocityRange:"), value)
}/* debug [instance_properties/setter]: velocityRange */


// The x component of an acceleration vector applied to cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/xAcceleration
func (e_ EmitterCell) XAcceleration() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("xAcceleration"))
	return rv
}/* debug [instance_properties/getter]: xAcceleration */


// The x component of an acceleration vector applied to cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/xAcceleration
func (e_ EmitterCell) SetXAcceleration(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setXAcceleration:"), value)
}/* debug [instance_properties/setter]: xAcceleration */


// The y component of an acceleration vector applied to cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/yAcceleration
func (e_ EmitterCell) YAcceleration() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("yAcceleration"))
	return rv
}/* debug [instance_properties/getter]: yAcceleration */


// The y component of an acceleration vector applied to cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/yAcceleration
func (e_ EmitterCell) SetYAcceleration(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setYAcceleration:"), value)
}/* debug [instance_properties/setter]: yAcceleration */


// The z component of an acceleration vector applied to cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/zAcceleration
func (e_ EmitterCell) ZAcceleration() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("zAcceleration"))
	return rv
}/* debug [instance_properties/getter]: zAcceleration */


// The z component of an acceleration vector applied to cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/zAcceleration
func (e_ EmitterCell) SetZAcceleration(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setZAcceleration:"), value)
}/* debug [instance_properties/setter]: zAcceleration */


// A Boolean value indicating whether or not cells from this emitter are rendered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/isenabled
func (e_ EmitterCell) IsEnabled() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value indicating whether or not cells from this emitter are rendered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/isenabled
func (e_ EmitterCell) SetIsEnabled(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAEmitterCell */




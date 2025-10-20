// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EmitterCell] class.
var (
	emitterCellClass     _EmitterCellClass
	emitterCellClassOnce sync.Once
)

func getEmitterCellClass() _EmitterCellClass {
	emitterCellClassOnce.Do(func() {
		emitterCellClass = _EmitterCellClass{objc.GetClass("CAEmitterCell")}
	})
	return emitterCellClass
}

type _EmitterCellClass struct {
	class objc.Class
}

// An interface definition for the [EmitterCell] class.
type IEmitterCell interface {
	objectivec.IObject
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



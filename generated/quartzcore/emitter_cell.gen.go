// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EmitterCell] class.
var emitterCellClass = _EmitterCellClass{objc.GetClass("CAEmitterCell")}

type _EmitterCellClass struct {
	class objc.Class
}

// An interface definition for the [EmitterCell] class.
type IEmitterCell interface {
	objectivec.IObject
}

// The definition of a particle emitted by a particle layer. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return emitterCellClass.New()
}





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




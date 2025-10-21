// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRMediaPlaybackClusterPlaybackPosition] class.
var (
	MTRMediaPlaybackClusterPlaybackPositionClass     _MTRMediaPlaybackClusterPlaybackPositionClass
	MTRMediaPlaybackClusterPlaybackPositionClassOnce sync.Once
)

func getMTRMediaPlaybackClusterPlaybackPositionClass() _MTRMediaPlaybackClusterPlaybackPositionClass {
	MTRMediaPlaybackClusterPlaybackPositionClassOnce.Do(func() {
		MTRMediaPlaybackClusterPlaybackPositionClass = _MTRMediaPlaybackClusterPlaybackPositionClass{objc.GetClass("MTRMediaPlaybackClusterPlaybackPosition")}
	})
	return MTRMediaPlaybackClusterPlaybackPositionClass
}

type _MTRMediaPlaybackClusterPlaybackPositionClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterPlaybackPosition] class.
type IMTRMediaPlaybackClusterPlaybackPosition interface {
	IMTRMediaPlaybackClusterPlaybackPositionStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterPlaybackPosition
type MTRMediaPlaybackClusterPlaybackPosition struct {
	MTRMediaPlaybackClusterPlaybackPositionStruct
}

// MTRMediaPlaybackClusterPlaybackPositionFrom constructs a [MTRMediaPlaybackClusterPlaybackPosition] from an unsafe.Pointer.
func MTRMediaPlaybackClusterPlaybackPositionFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterPlaybackPosition {
	return MTRMediaPlaybackClusterPlaybackPosition{
		MTRMediaPlaybackClusterPlaybackPositionStruct: MTRMediaPlaybackClusterPlaybackPositionStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterPlaybackPositionClass) Alloc() MTRMediaPlaybackClusterPlaybackPosition {
	rv := objc.Send[MTRMediaPlaybackClusterPlaybackPosition](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterPlaybackPositionClass) New() MTRMediaPlaybackClusterPlaybackPosition {
	rv := objc.Send[MTRMediaPlaybackClusterPlaybackPosition](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterPlaybackPosition) Init() MTRMediaPlaybackClusterPlaybackPosition {
	rv := objc.Send[MTRMediaPlaybackClusterPlaybackPosition](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterPlaybackPosition) Autorelease() MTRMediaPlaybackClusterPlaybackPosition {
	rv := objc.Send[MTRMediaPlaybackClusterPlaybackPosition](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterPlaybackPosition creates a new MTRMediaPlaybackClusterPlaybackPosition instance.
func NewMTRMediaPlaybackClusterPlaybackPosition() MTRMediaPlaybackClusterPlaybackPosition {
	return getMTRMediaPlaybackClusterPlaybackPositionClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterplaybackposition/updatedat
func (m_ MTRMediaPlaybackClusterPlaybackPosition) UpdatedAt() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("updatedAt"))
	return rv
}


// SetUpdatedAt sets the value of the updatedAt property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterplaybackposition/updatedat
func (m_ MTRMediaPlaybackClusterPlaybackPosition) SetUpdatedAt(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdatedAt:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterplaybackposition/position
func (m_ MTRMediaPlaybackClusterPlaybackPosition) Position() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("position"))
	return rv
}


// SetPosition sets the value of the position property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterplaybackposition/position
func (m_ MTRMediaPlaybackClusterPlaybackPosition) SetPosition(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPosition:"), value)
}




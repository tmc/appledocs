// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	Position() objc.IObject /* cross-framework: NSNumber */
	SetPosition(value objc.IObject /* cross-framework: NSNumber */)
	UpdatedAt() objc.IObject /* cross-framework: NSNumber */
	SetUpdatedAt(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterplaybackposition/position
func (m_ MTRMediaPlaybackClusterPlaybackPosition) Position() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("position"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterplaybackposition/position
func (m_ MTRMediaPlaybackClusterPlaybackPosition) SetPosition(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPosition:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterplaybackposition/updatedat
func (m_ MTRMediaPlaybackClusterPlaybackPosition) UpdatedAt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("updatedAt"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterplaybackposition/updatedat
func (m_ MTRMediaPlaybackClusterPlaybackPosition) SetUpdatedAt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdatedAt:"), value)
}

// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaPlaybackClusterPlaybackPositionStruct] class.
var (
	MTRMediaPlaybackClusterPlaybackPositionStructClass     _MTRMediaPlaybackClusterPlaybackPositionStructClass
	MTRMediaPlaybackClusterPlaybackPositionStructClassOnce sync.Once
)

func getMTRMediaPlaybackClusterPlaybackPositionStructClass() _MTRMediaPlaybackClusterPlaybackPositionStructClass {
	MTRMediaPlaybackClusterPlaybackPositionStructClassOnce.Do(func() {
		MTRMediaPlaybackClusterPlaybackPositionStructClass = _MTRMediaPlaybackClusterPlaybackPositionStructClass{objc.GetClass("MTRMediaPlaybackClusterPlaybackPositionStruct")}
	})
	return MTRMediaPlaybackClusterPlaybackPositionStructClass
}

type _MTRMediaPlaybackClusterPlaybackPositionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterPlaybackPositionStruct] class.
type IMTRMediaPlaybackClusterPlaybackPositionStruct interface {
	objectivec.IObject
	// properties:
	Position() objc.IObject /* cross-framework: NSNumber */
	SetPosition(value objc.IObject /* cross-framework: NSNumber */)
	UpdatedAt() objc.IObject /* cross-framework: NSNumber */
	SetUpdatedAt(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterPlaybackPositionStruct
type MTRMediaPlaybackClusterPlaybackPositionStruct struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterPlaybackPositionStructFrom constructs a [MTRMediaPlaybackClusterPlaybackPositionStruct] from an unsafe.Pointer.
func MTRMediaPlaybackClusterPlaybackPositionStructFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterPlaybackPositionStruct {
	return MTRMediaPlaybackClusterPlaybackPositionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterPlaybackPositionStructClass) Alloc() MTRMediaPlaybackClusterPlaybackPositionStruct {
	rv := objc.Send[MTRMediaPlaybackClusterPlaybackPositionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterPlaybackPositionStructClass) New() MTRMediaPlaybackClusterPlaybackPositionStruct {
	rv := objc.Send[MTRMediaPlaybackClusterPlaybackPositionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterPlaybackPositionStruct) Init() MTRMediaPlaybackClusterPlaybackPositionStruct {
	rv := objc.Send[MTRMediaPlaybackClusterPlaybackPositionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterPlaybackPositionStruct) Autorelease() MTRMediaPlaybackClusterPlaybackPositionStruct {
	rv := objc.Send[MTRMediaPlaybackClusterPlaybackPositionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterPlaybackPositionStruct creates a new MTRMediaPlaybackClusterPlaybackPositionStruct instance.
func NewMTRMediaPlaybackClusterPlaybackPositionStruct() MTRMediaPlaybackClusterPlaybackPositionStruct {
	return getMTRMediaPlaybackClusterPlaybackPositionStructClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterplaybackpositionstruct/position
func (m_ MTRMediaPlaybackClusterPlaybackPositionStruct) Position() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("position"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterplaybackpositionstruct/position
func (m_ MTRMediaPlaybackClusterPlaybackPositionStruct) SetPosition(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPosition:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterplaybackpositionstruct/updatedat
func (m_ MTRMediaPlaybackClusterPlaybackPositionStruct) UpdatedAt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("updatedAt"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterplaybackpositionstruct/updatedat
func (m_ MTRMediaPlaybackClusterPlaybackPositionStruct) SetUpdatedAt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdatedAt:"), value)
}

// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitDuration] class.
var unitDurationClass = _UnitDurationClass{objc.GetClass("NSUnitDuration")}

type _UnitDurationClass struct {
	class objc.Class
}

// An interface definition for the [UnitDuration] class.
type IUnitDuration interface {
	IDimension
}

// A unit of measure for a duration of time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration

type UnitDuration struct {
	Dimension
}

// UnitDurationFrom constructs a [UnitDuration] from an unsafe.Pointer.
//
// A unit of measure for a duration of time.
func UnitDurationFrom(ptr unsafe.Pointer) UnitDuration {
	return UnitDuration{
		Dimension: DimensionFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (uc _UnitDurationClass) Alloc() UnitDuration {
	rv := objc.Send[UnitDuration](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (uc _UnitDurationClass) New() UnitDuration {
	rv := objc.Send[UnitDuration](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitDuration) Init() UnitDuration {
	rv := objc.Send[UnitDuration](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitDuration) Autorelease() UnitDuration {
	rv := objc.Send[UnitDuration](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitDuration creates a new UnitDuration instance.
func NewUnitDuration() UnitDuration {
	return unitDurationClass.New()
}





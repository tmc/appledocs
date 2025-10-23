// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitDuration] class.
var (
	UnitDurationClass     _UnitDurationClass
	UnitDurationClassOnce sync.Once
)

func getUnitDurationClass() _UnitDurationClass {
	UnitDurationClassOnce.Do(func() {
		UnitDurationClass = _UnitDurationClass{objc.GetClass("NSUnitDuration")}
	})
	return UnitDurationClass
}

type _UnitDurationClass struct {
	class objc.Class
}

// An interface definition for the [UnitDuration] class.
type IUnitDuration interface {
	IDimension
}

// A unit of measure for a duration of time.
//
// You typically use instances of to represent specific quantities of planar angle using the class.


// A unit of measure for a duration of time.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getUnitDurationClass().New()
}



// The hour unit of duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration/hours
func (uc _UnitDurationClass) Hours() UnitDuration {
	rv := objc.Send[UnitDuration](objc.ID(uc.class), objc.Sel("hours"))
	return rv
}

// The microsecond unit of duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration/microseconds
func (uc _UnitDurationClass) Microseconds() UnitDuration {
	rv := objc.Send[UnitDuration](objc.ID(uc.class), objc.Sel("microseconds"))
	return rv
}

// The millisecond unit of duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration/milliseconds
func (uc _UnitDurationClass) Milliseconds() UnitDuration {
	rv := objc.Send[UnitDuration](objc.ID(uc.class), objc.Sel("milliseconds"))
	return rv
}

// The minute unit of duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration/minutes
func (uc _UnitDurationClass) Minutes() UnitDuration {
	rv := objc.Send[UnitDuration](objc.ID(uc.class), objc.Sel("minutes"))
	return rv
}

// The nanosecond unit of duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration/nanoseconds
func (uc _UnitDurationClass) Nanoseconds() UnitDuration {
	rv := objc.Send[UnitDuration](objc.ID(uc.class), objc.Sel("nanoseconds"))
	return rv
}

// The picosecond unit of duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration/picoseconds
func (uc _UnitDurationClass) Picoseconds() UnitDuration {
	rv := objc.Send[UnitDuration](objc.ID(uc.class), objc.Sel("picoseconds"))
	return rv
}

// The second unit of duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration/seconds
func (uc _UnitDurationClass) Seconds() UnitDuration {
	rv := objc.Send[UnitDuration](objc.ID(uc.class), objc.Sel("seconds"))
	return rv
}

// The hour unit of duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration/hours
func (u_ UnitDuration) Hours() IUnitDuration {
	rv := objc.Send[UnitDuration](u_.ID, objc.Sel("hours"))
	return rv
}


// The microsecond unit of duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration/microseconds
func (u_ UnitDuration) Microseconds() IUnitDuration {
	rv := objc.Send[UnitDuration](u_.ID, objc.Sel("microseconds"))
	return rv
}


// The millisecond unit of duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration/milliseconds
func (u_ UnitDuration) Milliseconds() IUnitDuration {
	rv := objc.Send[UnitDuration](u_.ID, objc.Sel("milliseconds"))
	return rv
}


// The minute unit of duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration/minutes
func (u_ UnitDuration) Minutes() IUnitDuration {
	rv := objc.Send[UnitDuration](u_.ID, objc.Sel("minutes"))
	return rv
}


// The nanosecond unit of duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration/nanoseconds
func (u_ UnitDuration) Nanoseconds() IUnitDuration {
	rv := objc.Send[UnitDuration](u_.ID, objc.Sel("nanoseconds"))
	return rv
}


// The picosecond unit of duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration/picoseconds
func (u_ UnitDuration) Picoseconds() IUnitDuration {
	rv := objc.Send[UnitDuration](u_.ID, objc.Sel("picoseconds"))
	return rv
}


// The second unit of duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration/seconds
func (u_ UnitDuration) Seconds() IUnitDuration {
	rv := objc.Send[UnitDuration](u_.ID, objc.Sel("seconds"))
	return rv
}




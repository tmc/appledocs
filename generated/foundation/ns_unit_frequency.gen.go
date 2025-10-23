// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitFrequency] class.
var (
	UnitFrequencyClass     _UnitFrequencyClass
	UnitFrequencyClassOnce sync.Once
)

func getUnitFrequencyClass() _UnitFrequencyClass {
	UnitFrequencyClassOnce.Do(func() {
		UnitFrequencyClass = _UnitFrequencyClass{objc.GetClass("NSUnitFrequency")}
	})
	return UnitFrequencyClass
}

type _UnitFrequencyClass struct {
	class objc.Class
}

// An interface definition for the [UnitFrequency] class.
type IUnitFrequency interface {
	IDimension
}

// A unit of measure for frequency.
//
// You typically use instances of to represent specific quantities of frequency using the class.


// A unit of measure for frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency
type UnitFrequency struct {
	Dimension
}

// UnitFrequencyFrom constructs a [UnitFrequency] from an unsafe.Pointer.
//
// A unit of measure for frequency.
func UnitFrequencyFrom(ptr unsafe.Pointer) UnitFrequency {
	return UnitFrequency{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitFrequencyClass) Alloc() UnitFrequency {
	rv := objc.Send[UnitFrequency](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitFrequencyClass) New() UnitFrequency {
	rv := objc.Send[UnitFrequency](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitFrequency) Init() UnitFrequency {
	rv := objc.Send[UnitFrequency](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitFrequency) Autorelease() UnitFrequency {
	rv := objc.Send[UnitFrequency](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitFrequency creates a new UnitFrequency instance.
func NewUnitFrequency() UnitFrequency {
	return getUnitFrequencyClass().New()
}



// The frames per second unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/framesPerSecond
func (uc _UnitFrequencyClass) FramesPerSecond() UnitFrequency {
	rv := objc.Send[NSUnitFrequency](objc.ID(uc.class), objc.Sel("framesPerSecond"))
	return rv
}

// The gigahertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/gigahertz
func (uc _UnitFrequencyClass) Gigahertz() UnitFrequency {
	rv := objc.Send[NSUnitFrequency](objc.ID(uc.class), objc.Sel("gigahertz"))
	return rv
}

// The hertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/hertz
func (uc _UnitFrequencyClass) Hertz() UnitFrequency {
	rv := objc.Send[NSUnitFrequency](objc.ID(uc.class), objc.Sel("hertz"))
	return rv
}

// The kilohertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/kilohertz
func (uc _UnitFrequencyClass) Kilohertz() UnitFrequency {
	rv := objc.Send[NSUnitFrequency](objc.ID(uc.class), objc.Sel("kilohertz"))
	return rv
}

// The megahertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/megahertz
func (uc _UnitFrequencyClass) Megahertz() UnitFrequency {
	rv := objc.Send[NSUnitFrequency](objc.ID(uc.class), objc.Sel("megahertz"))
	return rv
}

// The microhertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/microhertz
func (uc _UnitFrequencyClass) Microhertz() UnitFrequency {
	rv := objc.Send[NSUnitFrequency](objc.ID(uc.class), objc.Sel("microhertz"))
	return rv
}

// The millihertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/millihertz
func (uc _UnitFrequencyClass) Millihertz() UnitFrequency {
	rv := objc.Send[NSUnitFrequency](objc.ID(uc.class), objc.Sel("millihertz"))
	return rv
}

// The nanohertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/nanohertz
func (uc _UnitFrequencyClass) Nanohertz() UnitFrequency {
	rv := objc.Send[NSUnitFrequency](objc.ID(uc.class), objc.Sel("nanohertz"))
	return rv
}

// The terahertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/terahertz
func (uc _UnitFrequencyClass) Terahertz() UnitFrequency {
	rv := objc.Send[NSUnitFrequency](objc.ID(uc.class), objc.Sel("terahertz"))
	return rv
}

// The frames per second unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/framesPerSecond
func (u_ UnitFrequency) FramesPerSecond() NSUnitFrequency {
	rv := objc.Send[NSUnitFrequency](u_.ID, objc.Sel("framesPerSecond"))
	return rv
}


// The gigahertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/gigahertz
func (u_ UnitFrequency) Gigahertz() NSUnitFrequency {
	rv := objc.Send[NSUnitFrequency](u_.ID, objc.Sel("gigahertz"))
	return rv
}


// The hertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/hertz
func (u_ UnitFrequency) Hertz() NSUnitFrequency {
	rv := objc.Send[NSUnitFrequency](u_.ID, objc.Sel("hertz"))
	return rv
}


// The kilohertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/kilohertz
func (u_ UnitFrequency) Kilohertz() NSUnitFrequency {
	rv := objc.Send[NSUnitFrequency](u_.ID, objc.Sel("kilohertz"))
	return rv
}


// The megahertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/megahertz
func (u_ UnitFrequency) Megahertz() NSUnitFrequency {
	rv := objc.Send[NSUnitFrequency](u_.ID, objc.Sel("megahertz"))
	return rv
}


// The microhertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/microhertz
func (u_ UnitFrequency) Microhertz() NSUnitFrequency {
	rv := objc.Send[NSUnitFrequency](u_.ID, objc.Sel("microhertz"))
	return rv
}


// The millihertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/millihertz
func (u_ UnitFrequency) Millihertz() NSUnitFrequency {
	rv := objc.Send[NSUnitFrequency](u_.ID, objc.Sel("millihertz"))
	return rv
}


// The nanohertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/nanohertz
func (u_ UnitFrequency) Nanohertz() NSUnitFrequency {
	rv := objc.Send[NSUnitFrequency](u_.ID, objc.Sel("nanohertz"))
	return rv
}


// The terahertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/terahertz
func (u_ UnitFrequency) Terahertz() NSUnitFrequency {
	rv := objc.Send[NSUnitFrequency](u_.ID, objc.Sel("terahertz"))
	return rv
}




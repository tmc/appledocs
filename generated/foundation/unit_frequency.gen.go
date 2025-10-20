// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitFrequency] class.
var (
	unitFrequencyClass     _UnitFrequencyClass
	unitFrequencyClassOnce sync.Once
)

func getUnitFrequencyClass() _UnitFrequencyClass {
	unitFrequencyClassOnce.Do(func() {
		unitFrequencyClass = _UnitFrequencyClass{objc.GetClass("NSUnitFrequency")}
	})
	return unitFrequencyClass
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
//
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





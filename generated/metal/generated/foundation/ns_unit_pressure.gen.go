// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitPressure] class.
var (
	UnitPressureClass     _UnitPressureClass
	UnitPressureClassOnce sync.Once
)

func getUnitPressureClass() _UnitPressureClass {
	UnitPressureClassOnce.Do(func() {
		UnitPressureClass = _UnitPressureClass{objc.GetClass("NSUnitPressure")}
	})
	return UnitPressureClass
}

type _UnitPressureClass struct {
	class objc.Class
}

// An interface definition for the [UnitPressure] class.
type IUnitPressure interface {
	IDimension
	// properties:
	// methods:
}

// A unit of measure for pressure.
//
// You typically use instances of to represent specific quantities of pressure using the class.


// A unit of measure for pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure
type UnitPressure struct {
	Dimension
}

// UnitPressureFrom constructs a [UnitPressure] from an unsafe.Pointer.
//
// A unit of measure for pressure.
func UnitPressureFrom(ptr unsafe.Pointer) UnitPressure {
	return UnitPressure{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitPressureClass) Alloc() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitPressureClass) New() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitPressure) Init() UnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitPressure) Autorelease() UnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitPressure creates a new UnitPressure instance.
func NewUnitPressure() UnitPressure {
	return getUnitPressureClass().New()
}





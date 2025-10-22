// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitAcceleration] class.
var (
	UnitAccelerationClass     _UnitAccelerationClass
	UnitAccelerationClassOnce sync.Once
)

func getUnitAccelerationClass() _UnitAccelerationClass {
	UnitAccelerationClassOnce.Do(func() {
		UnitAccelerationClass = _UnitAccelerationClass{objc.GetClass("NSUnitAcceleration")}
	})
	return UnitAccelerationClass
}

type _UnitAccelerationClass struct {
	class objc.Class
}

// An interface definition for the [UnitAcceleration] class.
type IUnitAcceleration interface {
	IDimension
}

// A unit of measure for acceleration.
//
// You typically use instances of to represent specific quantities of acceleration using the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitAcceleration
type UnitAcceleration struct {
	Dimension
}

// UnitAccelerationFrom constructs a [UnitAcceleration] from an unsafe.Pointer.
//
// A unit of measure for acceleration.
func UnitAccelerationFrom(ptr unsafe.Pointer) UnitAcceleration {
	return UnitAcceleration{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitAccelerationClass) Alloc() UnitAcceleration {
	rv := objc.Send[UnitAcceleration](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitAccelerationClass) New() UnitAcceleration {
	rv := objc.Send[UnitAcceleration](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitAcceleration) Init() UnitAcceleration {
	rv := objc.Send[UnitAcceleration](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitAcceleration) Autorelease() UnitAcceleration {
	rv := objc.Send[UnitAcceleration](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitAcceleration creates a new UnitAcceleration instance.
func NewUnitAcceleration() UnitAcceleration {
	return getUnitAccelerationClass().New()
}


// Returns the meter per second squared unit of acceleration.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitAcceleration/metersPerSecondSquared
func (uc _UnitAccelerationClass) MetersPerSecondSquared() UnitAcceleration {
	rv := objc.Send[NSUnitAcceleration](objc.ID(uc.class), objc.Sel("metersPerSecondSquared"))
	return rv
}
// Returns the meter per second squared unit of acceleration.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitAcceleration/metersPerSecondSquared
func (u_ UnitAcceleration) MetersPerSecondSquared() NSUnitAcceleration {
	rv := objc.Send[NSUnitAcceleration](u_.ID, objc.Sel("metersPerSecondSquared"))
	return rv
}




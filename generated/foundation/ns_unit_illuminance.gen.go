// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitIlluminance] class.
var (
	UnitIlluminanceClass     _UnitIlluminanceClass
	UnitIlluminanceClassOnce sync.Once
)

func getUnitIlluminanceClass() _UnitIlluminanceClass {
	UnitIlluminanceClassOnce.Do(func() {
		UnitIlluminanceClass = _UnitIlluminanceClass{objc.GetClass("NSUnitIlluminance")}
	})
	return UnitIlluminanceClass
}

type _UnitIlluminanceClass struct {
	class objc.Class
}

// An interface definition for the [UnitIlluminance] class.
type IUnitIlluminance interface {
	IDimension
}

// A unit of measure for illuminance.
//
// You typically use instances of to represent specific quantities of illuminance using the class.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitIlluminance
type UnitIlluminance struct {
	Dimension
}

// UnitIlluminanceFrom constructs a [UnitIlluminance] from an unsafe.Pointer.
//
// A unit of measure for illuminance.
func UnitIlluminanceFrom(ptr unsafe.Pointer) UnitIlluminance {
	return UnitIlluminance{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitIlluminanceClass) Alloc() UnitIlluminance {
	rv := objc.Send[UnitIlluminance](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitIlluminanceClass) New() UnitIlluminance {
	rv := objc.Send[UnitIlluminance](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitIlluminance) Init() UnitIlluminance {
	rv := objc.Send[UnitIlluminance](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitIlluminance) Autorelease() UnitIlluminance {
	rv := objc.Send[UnitIlluminance](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitIlluminance creates a new UnitIlluminance instance.
func NewUnitIlluminance() UnitIlluminance {
	return getUnitIlluminanceClass().New()
}


// The lux unit of illuminance.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitIlluminance/lux
func (uc _UnitIlluminanceClass) Lux() UnitIlluminance {
	rv := objc.Send[NSUnitIlluminance](objc.ID(uc.class), objc.Sel("lux"))
	return rv
}
// The lux unit of illuminance.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitIlluminance/lux
func (u_ UnitIlluminance) Lux() NSUnitIlluminance {
	rv := objc.Send[NSUnitIlluminance](u_.ID, objc.Sel("lux"))
	return rv
}




// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitIlluminance] class.
var unitIlluminanceClass = _UnitIlluminanceClass{objc.GetClass("NSUnitIlluminance")}

type _UnitIlluminanceClass struct {
	class objc.Class
}

// An interface definition for the [UnitIlluminance] class.
type IUnitIlluminance interface {
	IDimension
}

// A unit of measure for illuminance. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return unitIlluminanceClass.New()
}





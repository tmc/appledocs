// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitDispersion] class.
var (
	UnitDispersionClass     _UnitDispersionClass
	UnitDispersionClassOnce sync.Once
)

func getUnitDispersionClass() _UnitDispersionClass {
	UnitDispersionClassOnce.Do(func() {
		UnitDispersionClass = _UnitDispersionClass{objc.GetClass("NSUnitDispersion")}
	})
	return UnitDispersionClass
}

type _UnitDispersionClass struct {
	class objc.Class
}

// An interface definition for the [UnitDispersion] class.
type IUnitDispersion interface {
	IDimension
}

// A unit of measure for specific quantities of dispersion.
//
// You typically use instances of to represent specific quantities of dispersion using the class.


// A unit of measure for specific quantities of dispersion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDispersion
type UnitDispersion struct {
	Dimension
}

// UnitDispersionFrom constructs a [UnitDispersion] from an unsafe.Pointer.
//
// A unit of measure for specific quantities of dispersion.
func UnitDispersionFrom(ptr unsafe.Pointer) UnitDispersion {
	return UnitDispersion{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitDispersionClass) Alloc() UnitDispersion {
	rv := objc.Send[UnitDispersion](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitDispersionClass) New() UnitDispersion {
	rv := objc.Send[UnitDispersion](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitDispersion) Init() UnitDispersion {
	rv := objc.Send[UnitDispersion](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitDispersion) Autorelease() UnitDispersion {
	rv := objc.Send[UnitDispersion](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitDispersion creates a new UnitDispersion instance.
func NewUnitDispersion() UnitDispersion {
	return getUnitDispersionClass().New()
}



// The parts per million unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDispersion/partsPerMillion
func (uc _UnitDispersionClass) PartsPerMillion() UnitDispersion {
	rv := objc.Send[NSUnitDispersion](objc.ID(uc.class), objc.Sel("partsPerMillion"))
	return rv
}

// The parts per million unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDispersion/partsPerMillion
func (u_ UnitDispersion) PartsPerMillion() NSUnitDispersion {
	rv := objc.Send[NSUnitDispersion](u_.ID, objc.Sel("partsPerMillion"))
	return rv
}




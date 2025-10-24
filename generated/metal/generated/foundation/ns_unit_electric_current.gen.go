// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitElectricCurrent] class.
var (
	UnitElectricCurrentClass     _UnitElectricCurrentClass
	UnitElectricCurrentClassOnce sync.Once
)

func getUnitElectricCurrentClass() _UnitElectricCurrentClass {
	UnitElectricCurrentClassOnce.Do(func() {
		UnitElectricCurrentClass = _UnitElectricCurrentClass{objc.GetClass("NSUnitElectricCurrent")}
	})
	return UnitElectricCurrentClass
}

type _UnitElectricCurrentClass struct {
	class objc.Class
}

// An interface definition for the [UnitElectricCurrent] class.
type IUnitElectricCurrent interface {
	IDimension
	// properties:
	// methods:
}

// A unit of measure for electric current.
//
// You typically use instances of to represent specific quantities of electric current using the class.


// A unit of measure for electric current.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCurrent
type UnitElectricCurrent struct {
	Dimension
}

// UnitElectricCurrentFrom constructs a [UnitElectricCurrent] from an unsafe.Pointer.
//
// A unit of measure for electric current.
func UnitElectricCurrentFrom(ptr unsafe.Pointer) UnitElectricCurrent {
	return UnitElectricCurrent{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitElectricCurrentClass) Alloc() UnitElectricCurrent {
	rv := objc.Send[UnitElectricCurrent](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitElectricCurrentClass) New() UnitElectricCurrent {
	rv := objc.Send[UnitElectricCurrent](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitElectricCurrent) Init() UnitElectricCurrent {
	rv := objc.Send[UnitElectricCurrent](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitElectricCurrent) Autorelease() UnitElectricCurrent {
	rv := objc.Send[UnitElectricCurrent](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitElectricCurrent creates a new UnitElectricCurrent instance.
func NewUnitElectricCurrent() UnitElectricCurrent {
	return getUnitElectricCurrentClass().New()
}





// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitConverterLinear] class.
var (
	unitConverterLinearClass     _UnitConverterLinearClass
	unitConverterLinearClassOnce sync.Once
)

func getUnitConverterLinearClass() _UnitConverterLinearClass {
	unitConverterLinearClassOnce.Do(func() {
		unitConverterLinearClass = _UnitConverterLinearClass{objc.GetClass("NSUnitConverterLinear")}
	})
	return unitConverterLinearClass
}

type _UnitConverterLinearClass struct {
	class objc.Class
}

// An interface definition for the [UnitConverterLinear] class.
type IUnitConverterLinear interface {
	IUnitConverter
}

// A description of how to convert between units using a linear equation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitConverterLinear
type UnitConverterLinear struct {
	UnitConverter
}

// UnitConverterLinearFrom constructs a [UnitConverterLinear] from an unsafe.Pointer.
//
// A description of how to convert between units using a linear equation.
func UnitConverterLinearFrom(ptr unsafe.Pointer) UnitConverterLinear {
	return UnitConverterLinear{
		UnitConverter: UnitConverterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitConverterLinearClass) Alloc() UnitConverterLinear {
	rv := objc.Send[UnitConverterLinear](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitConverterLinearClass) New() UnitConverterLinear {
	rv := objc.Send[UnitConverterLinear](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitConverterLinear) Init() UnitConverterLinear {
	rv := objc.Send[UnitConverterLinear](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitConverterLinear) Autorelease() UnitConverterLinear {
	rv := objc.Send[UnitConverterLinear](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitConverterLinear creates a new UnitConverterLinear instance.
func NewUnitConverterLinear() UnitConverterLinear {
	return getUnitConverterLinearClass().New()
}





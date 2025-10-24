// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitConverterLinear] class.
var (
	UnitConverterLinearClass     _UnitConverterLinearClass
	UnitConverterLinearClassOnce sync.Once
)

func getUnitConverterLinearClass() _UnitConverterLinearClass {
	UnitConverterLinearClassOnce.Do(func() {
		UnitConverterLinearClass = _UnitConverterLinearClass{objc.GetClass("NSUnitConverterLinear")}
	})
	return UnitConverterLinearClass
}

type _UnitConverterLinearClass struct {
	class objc.Class
}

// An interface definition for the [UnitConverterLinear] class.
type IUnitConverterLinear interface {
	IUnitConverter
	// properties:
	Constant() float64
	Coefficient() float64
	SetCoefficient(value float64)
	// methods:
}

// A description of how to convert between units using a linear equation.
//
// A linear equation for unit conversion takes the form , such that the following is true: is the value in terms of the base unit of the dimension. is the known coefficient to use for this unit’s conversion. is the value in terms of the unit on which you call this method. is the known constant to use for this unit’s conversion. The method performs the conversion in the form of , where represents the value passed in and represents the value returned. The method performs the inverse conversion in the form of , where represents the value passed in and represents the value returned. For example, consider the unit that defines. The method calculates the value in the base unit, , using the formula . The method calculates the value in using the formula , where the is and the is . Units that perform conversion using only a scale factor have a equal to the scale factor and a equal to . For example, consider the unit defines. The method calculates the value in meters using the formula . The calculates the value in kilometers using the formula , where the coefficient is and the constant is .


// A description of how to convert between units using a linear equation.
//
// [Full Topic]
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



// The constant to use in the linear unit conversion calculation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitConverterLinear/constant
func (u_ UnitConverterLinear) Constant() float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("constant"))
	return rv
}


// The coefficient to use in the linear unit conversion calculation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/unitconverterlinear/coefficient
func (u_ UnitConverterLinear) Coefficient() float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("coefficient"))
	return rv
}


// The coefficient to use in the linear unit conversion calculation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/unitconverterlinear/coefficient
func (u_ UnitConverterLinear) SetCoefficient(value float64) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCoefficient:"), value)
}




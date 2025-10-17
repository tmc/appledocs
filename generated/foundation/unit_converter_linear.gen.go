// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitConverterLinear] class.
var unitConverterLinearClass = _UnitConverterLinearClass{objc.GetClass("NSUnitConverterLinear")}

type _UnitConverterLinearClass struct {
	class objc.Class
}

// An interface definition for the [UnitConverterLinear] class.
type IUnitConverterLinear interface {
	IUnitConverter
}

// A description of how to convert between units using a linear equation. [Full Topic]
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




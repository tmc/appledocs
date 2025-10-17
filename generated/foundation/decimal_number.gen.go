// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DecimalNumber] class.
var decimalNumberClass = _DecimalNumberClass{objc.GetClass("NSDecimalNumber")}

type _DecimalNumberClass struct {
	class objc.Class
}

// An object for representing and performing arithmetic on base-10 numbers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber

type DecimalNumber struct {
	Number
}

// DecimalNumberFrom constructs a [DecimalNumber] from an unsafe.Pointer.
//
// An object for representing and performing arithmetic on base-10 numbers.
func DecimalNumberFrom(ptr unsafe.Pointer) DecimalNumber {
	return DecimalNumber{
		Number: NumberFrom(ptr),
	}
}




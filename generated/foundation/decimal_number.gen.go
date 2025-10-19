// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DecimalNumber] class.
var (
	decimalNumberClass     _DecimalNumberClass
	decimalNumberClassOnce sync.Once
)

func getDecimalNumberClass() _DecimalNumberClass {
	decimalNumberClassOnce.Do(func() {
		decimalNumberClass = _DecimalNumberClass{objc.GetClass("NSDecimalNumber")}
	})
	return decimalNumberClass
}

type _DecimalNumberClass struct {
	class objc.Class
}

// An interface definition for the [DecimalNumber] class.
type IDecimalNumber interface {
	INumber
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

// Alloc allocates a new instance without initialization.
func (dc _DecimalNumberClass) Alloc() DecimalNumber {
	rv := objc.Send[DecimalNumber](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DecimalNumberClass) New() DecimalNumber {
	rv := objc.Send[DecimalNumber](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DecimalNumber) Init() DecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DecimalNumber) Autorelease() DecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDecimalNumber creates a new DecimalNumber instance.
func NewDecimalNumber() DecimalNumber {
	return getDecimalNumberClass().New()
}





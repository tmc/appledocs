// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DecimalNumber] class.
var (
	DecimalNumberClass     _DecimalNumberClass
	DecimalNumberClassOnce sync.Once
)

func getDecimalNumberClass() _DecimalNumberClass {
	DecimalNumberClassOnce.Do(func() {
		DecimalNumberClass = _DecimalNumberClass{objc.GetClass("NSDecimalNumber")}
	})
	return DecimalNumberClass
}

type _DecimalNumberClass struct {
	class objc.Class
}

// An interface definition for the [DecimalNumber] class.
type IDecimalNumber interface {
	INumber
	// properties:
	DecimalValue() unsafe.Pointer
	SetDecimalValue(value unsafe.Pointer)
	DoubleValue() float64 /* primitive/slice/pointer */
	SetDoubleValue(value float64 /* primitive/slice/pointer */)
	ObjCType() unsafe.Pointer
	SetObjCType(value unsafe.Pointer)
	// methods:
}

// An object for representing and performing arithmetic on base-10 numbers.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. , an immutable subclass of , provides an object-oriented wrapper for doing base-10 arithmetic. An instance can represent any number that can be expressed as where mantissa is a decimal integer up to 38 digits long, and exponent is an integer from –128 through 127.


// An object for representing and performing arithmetic on base-10 numbers.
//
// [Full Topic]
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



// The decimal number’s value, expressed as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/decimalvalue
func (d_ DecimalNumber) DecimalValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("decimalValue"))
	return rv
}


// The decimal number’s value, expressed as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/decimalvalue
func (d_ DecimalNumber) SetDecimalValue(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDecimalValue:"), value)
}


// The decimal number’s closest approximate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/doublevalue
func (d_ DecimalNumber) DoubleValue() float64 /* primitive/slice/pointer */ {
	rv := objc.Send[float64](d_.ID, objc.Sel("doubleValue"))
	return rv
}


// The decimal number’s closest approximate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/doublevalue
func (d_ DecimalNumber) SetDoubleValue(value float64 /* primitive/slice/pointer */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDoubleValue:"), value)
}


// A C string containing the Objective-C type for the data contained in the decimal number object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/objctype
func (d_ DecimalNumber) ObjCType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("objCType"))
	return rv
}


// A C string containing the Objective-C type for the data contained in the decimal number object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/objctype
func (d_ DecimalNumber) SetObjCType(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setObjCType:"), value)
}




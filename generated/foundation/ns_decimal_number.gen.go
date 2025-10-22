// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	DecimalNumberByAdding(decimalNumber IDecimalNumber) DecimalNumber
	DecimalNumberByAddingWithBehavior(decimalNumber IDecimalNumber, behavior objectivec.IObject) DecimalNumber
	Compare(decimalNumber INumber) ComparisonResult
	DescriptionWithLocale(locale objectivec.IObject) String
	DecimalNumberByDividingBy(decimalNumber IDecimalNumber) DecimalNumber
	DecimalNumberByDividingByWithBehavior(decimalNumber IDecimalNumber, behavior objectivec.IObject) DecimalNumber
	DecimalNumberByMultiplyingBy(decimalNumber IDecimalNumber) DecimalNumber
	DecimalNumberByMultiplyingByWithBehavior(decimalNumber IDecimalNumber, behavior objectivec.IObject) DecimalNumber
	DecimalNumberByMultiplyingByPowerOf10(power unsafe.Pointer) DecimalNumber
	DecimalNumberByMultiplyingByPowerOf10WithBehavior(power unsafe.Pointer, behavior objectivec.IObject) DecimalNumber
	DecimalNumberByRaisingToPower(power uint) DecimalNumber
	DecimalNumberByRaisingToPowerWithBehavior(power uint, behavior objectivec.IObject) DecimalNumber
	DecimalNumberByRoundingAccordingToBehavior(behavior objectivec.IObject) DecimalNumber
	DecimalNumberBySubtracting(decimalNumber IDecimalNumber) DecimalNumber
	DecimalNumberBySubtractingWithBehavior(decimalNumber IDecimalNumber, behavior objectivec.IObject) DecimalNumber
	DecimalValue() unsafe.Pointer
	DoubleValue() float64
	ObjCType() unsafe.Pointer
	SetObjCType(value unsafe.Pointer)
}

// An object for representing and performing arithmetic on base-10 numbers.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. , an immutable subclass of , provides an object-oriented wrapper for doing base-10 arithmetic. An instance can represent any number that can be expressed as where mantissa is a decimal integer up to 38 digits long, and exponent is an integer from –128 through 127.
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




// Initializes a decimal number to represent a given decimal.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/init(decimal:)
func NewDecimalNumberWithDecimal(dcm unsafe.Pointer) DecimalNumber {
	instance := getDecimalNumberClass().Alloc()
	rv := objc.Send[DecimalNumber](instance.ID, objc.Sel("initWithDecimal:"), dcm)
	rv.Autorelease()
	return rv
}



// Initializes a decimal number using the given mantissa, exponent, and sign.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/init(mantissa:exponent:isNegative:)
func NewDecimalNumberWithMantissaExponentIsNegative(mantissa uint64, exponent unsafe.Pointer, flag bool) DecimalNumber {
	instance := getDecimalNumberClass().Alloc()
	rv := objc.Send[DecimalNumber](instance.ID, objc.Sel("initWithMantissa:exponent:isNegative:"), mantissa, exponent, flag)
	rv.Autorelease()
	return rv
}



// Initializes a decimal number so that its value is equivalent to that in a given numeric string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/init(string:)
func NewDecimalNumberWithString(numberValue string) DecimalNumber {
	instance := getDecimalNumberClass().Alloc()
	rv := objc.Send[DecimalNumber](instance.ID, objc.Sel("initWithString:"), objc.String(numberValue))
	rv.Autorelease()
	return rv
}



// Initializes a decimal number so that its value is equivalent to that in a given numeric string, interpreted using a given locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/init(string:locale:)
func NewDecimalNumberWithStringLocale(numberValue string, locale objectivec.IObject) DecimalNumber {
	instance := getDecimalNumberClass().Alloc()
	rv := objc.Send[DecimalNumber](instance.ID, objc.Sel("initWithString:locale:"), objc.String(numberValue), locale)
	rv.Autorelease()
	return rv
}


// Creates and returns a decimal number equivalent to a given decimal structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/decimalNumberWithDecimal:
func (dc _DecimalNumberClass) DecimalNumberWithDecimal(dcm unsafe.Pointer) DecimalNumber {
	rv := objc.Send[DecimalNumber](objc.ID(dc.class), objc.Sel("decimalNumberWithDecimal:"), dcm)
	return rv
}

// Creates and returns a decimal number equivalent to the number specified by the arguments.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/decimalNumberWithMantissa:exponent:isNegative:
func (dc _DecimalNumberClass) DecimalNumberWithMantissaExponentIsNegative(mantissa uint64, exponent unsafe.Pointer, flag bool) DecimalNumber {
	rv := objc.Send[DecimalNumber](objc.ID(dc.class), objc.Sel("decimalNumberWithMantissa:exponent:isNegative:"), mantissa, exponent, flag)
	return rv
}

// Creates a decimal number whose value is equivalent to that in a given numeric string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/decimalNumberWithString:
func (dc _DecimalNumberClass) DecimalNumberWithString(numberValue string) DecimalNumber {
	rv := objc.Send[DecimalNumber](objc.ID(dc.class), objc.Sel("decimalNumberWithString:"), objc.String(numberValue))
	return rv
}

// Creates a decimal number whose value is equivalent to that in a given numeric string, interpreted using a given locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/decimalNumberWithString:locale:
func (dc _DecimalNumberClass) DecimalNumberWithStringLocale(numberValue string, locale objectivec.IObject) DecimalNumber {
	rv := objc.Send[DecimalNumber](objc.ID(dc.class), objc.Sel("decimalNumberWithString:locale:"), objc.String(numberValue), locale)
	return rv
}

// The way arithmetic methods round off and handle error conditions.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/defaultBehavior
func (dc _DecimalNumberClass) DefaultBehavior() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(dc.class), objc.Sel("defaultBehavior"))
	return rv
}
// Returns the largest possible value of a decimal number.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/maximum
func (dc _DecimalNumberClass) MaximumDecimalNumber() DecimalNumber {
	rv := objc.Send[NSDecimalNumber](objc.ID(dc.class), objc.Sel("maximumDecimalNumber"))
	return rv
}
// Returns the smallest possible value of a decimal number.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/minimum
func (dc _DecimalNumberClass) MinimumDecimalNumber() DecimalNumber {
	rv := objc.Send[NSDecimalNumber](objc.ID(dc.class), objc.Sel("minimumDecimalNumber"))
	return rv
}
// A decimal number that specifies no number.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/notANumber
func (dc _DecimalNumberClass) NotANumber() DecimalNumber {
	rv := objc.Send[NSDecimalNumber](objc.ID(dc.class), objc.Sel("notANumber"))
	return rv
}
// A decimal number equivalent to the number 1.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/one
func (dc _DecimalNumberClass) One() DecimalNumber {
	rv := objc.Send[NSDecimalNumber](objc.ID(dc.class), objc.Sel("one"))
	return rv
}
// A decimal number equivalent to the number 0.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/zero
func (dc _DecimalNumberClass) Zero() DecimalNumber {
	rv := objc.Send[NSDecimalNumber](objc.ID(dc.class), objc.Sel("zero"))
	return rv
}
// Adds this number to another given number.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/adding(_:)
func (d_ DecimalNumber) DecimalNumberByAdding(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByAdding:"), decimalNumber)
	return rv
}

// Adds this number to another given number using the specified behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/adding(_:withBehavior:)
func (d_ DecimalNumber) DecimalNumberByAddingWithBehavior(decimalNumber IDecimalNumber, behavior objectivec.IObject) DecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByAdding:withBehavior:"), decimalNumber, behavior)
	return rv
}

// Compares this decimal number and another.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/compare(_:)
func (d_ DecimalNumber) Compare(decimalNumber INumber) ComparisonResult {
	rv := objc.Send[ComparisonResult](d_.ID, objc.Sel("compare:"), decimalNumber)
	return rv
}

// Returns a string representation of the decimal number appropriate for the specified locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/description(withLocale:)
func (d_ DecimalNumber) DescriptionWithLocale(locale objectivec.IObject) String {
	rv := objc.Send[String](d_.ID, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}

// Divides the number by another given number.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/dividing(by:)
func (d_ DecimalNumber) DecimalNumberByDividingBy(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByDividingBy:"), decimalNumber)
	return rv
}

// Divides this number by another given number using the specified behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/dividing(by:withBehavior:)
func (d_ DecimalNumber) DecimalNumberByDividingByWithBehavior(decimalNumber IDecimalNumber, behavior objectivec.IObject) DecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByDividingBy:withBehavior:"), decimalNumber, behavior)
	return rv
}

// Multiplies the number by another given number.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/multiplying(by:)
func (d_ DecimalNumber) DecimalNumberByMultiplyingBy(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByMultiplyingBy:"), decimalNumber)
	return rv
}

// Multiplies this number by another given number using the specified behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/multiplying(by:withBehavior:)
func (d_ DecimalNumber) DecimalNumberByMultiplyingByWithBehavior(decimalNumber IDecimalNumber, behavior objectivec.IObject) DecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByMultiplyingBy:withBehavior:"), decimalNumber, behavior)
	return rv
}

// Multiplies the number by 10 raised to the given power.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/multiplying(byPowerOf10:)
func (d_ DecimalNumber) DecimalNumberByMultiplyingByPowerOf10(power unsafe.Pointer) DecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByMultiplyingByPowerOf10:"), power)
	return rv
}

// Multiplies the number by 10 raised to the given power using the specified behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/multiplying(byPowerOf10:withBehavior:)
func (d_ DecimalNumber) DecimalNumberByMultiplyingByPowerOf10WithBehavior(power unsafe.Pointer, behavior objectivec.IObject) DecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByMultiplyingByPowerOf10:withBehavior:"), power, behavior)
	return rv
}

// Raises the number to a given power.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/raising(toPower:)
func (d_ DecimalNumber) DecimalNumberByRaisingToPower(power uint) DecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByRaisingToPower:"), power)
	return rv
}

// Raises the number to a given power using the specified behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/raising(toPower:withBehavior:)
func (d_ DecimalNumber) DecimalNumberByRaisingToPowerWithBehavior(power uint, behavior objectivec.IObject) DecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByRaisingToPower:withBehavior:"), power, behavior)
	return rv
}

// Returns a rounded version of the decimal number using the specified rounding behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/rounding(accordingToBehavior:)
func (d_ DecimalNumber) DecimalNumberByRoundingAccordingToBehavior(behavior objectivec.IObject) DecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByRoundingAccordingToBehavior:"), behavior)
	return rv
}

// Subtracts another given number from this one.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/subtracting(_:)
func (d_ DecimalNumber) DecimalNumberBySubtracting(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberBySubtracting:"), decimalNumber)
	return rv
}

// Subtracts this a given number from this one using the specified behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/subtracting(_:withBehavior:)
func (d_ DecimalNumber) DecimalNumberBySubtractingWithBehavior(decimalNumber IDecimalNumber, behavior objectivec.IObject) DecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberBySubtracting:withBehavior:"), decimalNumber, behavior)
	return rv
}

// The decimal number’s value, expressed as an structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/decimalValue
func (d_ DecimalNumber) DecimalValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("decimalValue"))
	return rv
}

// The way arithmetic methods round off and handle error conditions.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/defaultBehavior
func (d_ DecimalNumber) DefaultBehavior() objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("defaultBehavior"))
	return rv
}


// SetDefaultBehavior sets the value of the defaultBehavior property.
// The way arithmetic methods round off and handle error conditions.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/defaultBehavior
func (d_ DecimalNumber) SetDefaultBehavior(value objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDefaultBehavior:"), value)
}

// The decimal number’s closest approximate value.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/doubleValue
func (d_ DecimalNumber) DoubleValue() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("doubleValue"))
	return rv
}

// Returns the largest possible value of a decimal number.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/maximum
func (d_ DecimalNumber) MaximumDecimalNumber() NSDecimalNumber {
	rv := objc.Send[NSDecimalNumber](d_.ID, objc.Sel("maximumDecimalNumber"))
	return rv
}

// Returns the smallest possible value of a decimal number.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/minimum
func (d_ DecimalNumber) MinimumDecimalNumber() NSDecimalNumber {
	rv := objc.Send[NSDecimalNumber](d_.ID, objc.Sel("minimumDecimalNumber"))
	return rv
}

// A decimal number that specifies no number.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/notANumber
func (d_ DecimalNumber) NotANumber() NSDecimalNumber {
	rv := objc.Send[NSDecimalNumber](d_.ID, objc.Sel("notANumber"))
	return rv
}

// A decimal number equivalent to the number 1.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/one
func (d_ DecimalNumber) One() NSDecimalNumber {
	rv := objc.Send[NSDecimalNumber](d_.ID, objc.Sel("one"))
	return rv
}

// A decimal number equivalent to the number 0.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/zero
func (d_ DecimalNumber) Zero() NSDecimalNumber {
	rv := objc.Send[NSDecimalNumber](d_.ID, objc.Sel("zero"))
	return rv
}

// A C string containing the Objective-C type for the data contained in the decimal number object.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/objctype
func (d_ DecimalNumber) ObjCType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("objCType"))
	return rv
}


// SetObjCType sets the value of the objCType property.
// A C string containing the Objective-C type for the data contained in the decimal number object.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdecimalnumber/objctype
func (d_ DecimalNumber) SetObjCType(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setObjCType:"), value)
}



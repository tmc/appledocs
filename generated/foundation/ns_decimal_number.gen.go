// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSDecimalNumber */


/* debug [class_header]: Header for NSDecimalNumber */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DecimalNumber */
// An interface definition for the [DecimalNumber] class.
type IDecimalNumber interface {
	INumber
	
/* debug [class_interface_properties]: Properties for DecimalNumber */
	// properties:
	DecimalValue() objc.IObject /* cross-framework: Decimal */
	DoubleValue() float64
	ObjCType() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DecimalNumber */
	// methods:
	DecimalNumberByAdding(decimalNumber IDecimalNumber) IDecimalNumber
	DecimalNumberByAddingWithBehavior(decimalNumber IDecimalNumber, behavior unsafe.Pointer) IDecimalNumber
	Compare(decimalNumber INumber) ComparisonResult
	DescriptionWithLocale(locale objc.IObject) IString
	DecimalNumberByDividingBy(decimalNumber IDecimalNumber) IDecimalNumber
	DecimalNumberByDividingByWithBehavior(decimalNumber IDecimalNumber, behavior unsafe.Pointer) IDecimalNumber
	DecimalNumberByMultiplyingBy(decimalNumber IDecimalNumber) IDecimalNumber
	DecimalNumberByMultiplyingByWithBehavior(decimalNumber IDecimalNumber, behavior unsafe.Pointer) IDecimalNumber
	DecimalNumberByMultiplyingByPowerOf10(power objectivec.IObject) IDecimalNumber
	DecimalNumberByMultiplyingByPowerOf10WithBehavior(power objectivec.IObject, behavior unsafe.Pointer) IDecimalNumber
	DecimalNumberByRaisingToPower(power uint) IDecimalNumber
	DecimalNumberByRaisingToPowerWithBehavior(power uint, behavior unsafe.Pointer) IDecimalNumber
	DecimalNumberByRoundingAccordingToBehavior(behavior unsafe.Pointer) IDecimalNumber
	DecimalNumberBySubtracting(decimalNumber IDecimalNumber) IDecimalNumber
	DecimalNumberBySubtractingWithBehavior(decimalNumber IDecimalNumber, behavior unsafe.Pointer) IDecimalNumber
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DecimalNumber */
// Alloc allocates a new instance without initialization.
func (dc _DecimalNumberClass) Alloc() DecimalNumber {
	rv := objc.Send[DecimalNumber](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DecimalNumber */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DecimalNumber */

// Initializes a decimal number to represent a given decimal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/init(decimal:)
func NewDecimalNumberWithDecimal(dcm objc.IObject /* cross-framework: Decimal */) DecimalNumber {
	instance := getDecimalNumberClass().Alloc()
	rv := objc.Send[DecimalNumber](instance.ID, objc.Sel("initWithDecimal:"), dcm)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDecimalNumberWithDecimal */


// Initializes a decimal number using the given mantissa, exponent, and sign.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/init(mantissa:exponent:isNegative:)
func NewDecimalNumberWithMantissaExponentIsNegative(mantissa uint64, exponent objectivec.IObject, flag bool) DecimalNumber {
	instance := getDecimalNumberClass().Alloc()
	rv := objc.Send[DecimalNumber](instance.ID, objc.Sel("initWithMantissa:exponent:isNegative:"), mantissa, exponent, flag)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDecimalNumberWithMantissaExponentIsNegative */


// Initializes a decimal number so that its value is equivalent to that in a given numeric string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/init(string:)
func NewDecimalNumberWithString(numberValue IString) DecimalNumber {
	instance := getDecimalNumberClass().Alloc()
	rv := objc.Send[DecimalNumber](instance.ID, objc.Sel("initWithString:"), numberValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDecimalNumberWithString */


// Initializes a decimal number so that its value is equivalent to that in a given numeric string, interpreted using a given locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/init(string:locale:)
func NewDecimalNumberWithStringLocale(numberValue IString, locale objc.IObject) DecimalNumber {
	instance := getDecimalNumberClass().Alloc()
	rv := objc.Send[DecimalNumber](instance.ID, objc.Sel("initWithString:locale:"), numberValue, locale)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDecimalNumberWithStringLocale */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DecimalNumber */

// Creates and returns a decimal number equivalent to a given decimal structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/decimalNumberWithDecimal:
func (dc _DecimalNumberClass) DecimalNumberWithDecimal(dcm objc.IObject /* cross-framework: Decimal */) IDecimalNumber {
	rv := objc.Send[DecimalNumber](objc.ID(dc.class), objc.Sel("decimalNumberWithDecimal:"), dcm)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DecimalNumberWithDecimal) */


// Creates and returns a decimal number equivalent to the number specified by the arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/decimalNumberWithMantissa:exponent:isNegative:
func (dc _DecimalNumberClass) DecimalNumberWithMantissaExponentIsNegative(mantissa uint64, exponent objectivec.IObject, flag bool) IDecimalNumber {
	rv := objc.Send[DecimalNumber](objc.ID(dc.class), objc.Sel("decimalNumberWithMantissa:exponent:isNegative:"), mantissa, exponent, flag)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DecimalNumberWithMantissaExponentIsNegative) */


// Creates a decimal number whose value is equivalent to that in a given numeric string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/decimalNumberWithString:
func (dc _DecimalNumberClass) DecimalNumberWithString(numberValue IString) IDecimalNumber {
	rv := objc.Send[DecimalNumber](objc.ID(dc.class), objc.Sel("decimalNumberWithString:"), numberValue)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DecimalNumberWithString) */


// Creates a decimal number whose value is equivalent to that in a given numeric string, interpreted using a given locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/decimalNumberWithString:locale:
func (dc _DecimalNumberClass) DecimalNumberWithStringLocale(numberValue IString, locale objc.IObject) IDecimalNumber {
	rv := objc.Send[DecimalNumber](objc.ID(dc.class), objc.Sel("decimalNumberWithString:locale:"), numberValue, locale)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DecimalNumberWithStringLocale) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DecimalNumber */

// The way arithmetic methods round off and handle error conditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/defaultBehavior
func (dc _DecimalNumberClass) DefaultBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("defaultBehavior"))
	return rv
}/* debug [class_properties_class/property]: defaultBehavior */

// Returns the largest possible value of a decimal number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/maximum
func (dc _DecimalNumberClass) MaximumDecimalNumber() DecimalNumber {
	rv := objc.Send[DecimalNumber](objc.ID(dc.class), objc.Sel("maximumDecimalNumber"))
	return rv
}/* debug [class_properties_class/property]: maximumDecimalNumber */

// Returns the smallest possible value of a decimal number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/minimum
func (dc _DecimalNumberClass) MinimumDecimalNumber() DecimalNumber {
	rv := objc.Send[DecimalNumber](objc.ID(dc.class), objc.Sel("minimumDecimalNumber"))
	return rv
}/* debug [class_properties_class/property]: minimumDecimalNumber */

// A decimal number that specifies no number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/notANumber
func (dc _DecimalNumberClass) NotANumber() DecimalNumber {
	rv := objc.Send[DecimalNumber](objc.ID(dc.class), objc.Sel("notANumber"))
	return rv
}/* debug [class_properties_class/property]: notANumber */

// A decimal number equivalent to the number 1.0.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/one
func (dc _DecimalNumberClass) One() DecimalNumber {
	rv := objc.Send[DecimalNumber](objc.ID(dc.class), objc.Sel("one"))
	return rv
}/* debug [class_properties_class/property]: one */

// A decimal number equivalent to the number 0.0.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/zero
func (dc _DecimalNumberClass) Zero() DecimalNumber {
	rv := objc.Send[DecimalNumber](objc.ID(dc.class), objc.Sel("zero"))
	return rv
}/* debug [class_properties_class/property]: zero */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DecimalNumber */

// Adds this number to another given number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/adding(_:)
func (d_ DecimalNumber) DecimalNumberByAdding(decimalNumber IDecimalNumber) IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByAdding:"), decimalNumber)
	return rv
}/* debug [instance_methods/method]: DecimalNumberByAdding */


// Adds this number to another given number using the specified behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/adding(_:withBehavior:)
func (d_ DecimalNumber) DecimalNumberByAddingWithBehavior(decimalNumber IDecimalNumber, behavior unsafe.Pointer) IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByAdding:withBehavior:"), decimalNumber, behavior)
	return rv
}/* debug [instance_methods/method]: DecimalNumberByAddingWithBehavior */


// Compares this decimal number and another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/compare(_:)
func (d_ DecimalNumber) Compare(decimalNumber INumber) ComparisonResult {
	rv := objc.Send[ComparisonResult](d_.ID, objc.Sel("compare:"), decimalNumber)
	return rv
}/* debug [instance_methods/method]: Compare */


// Returns a string representation of the decimal number appropriate for the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/description(withLocale:)
func (d_ DecimalNumber) DescriptionWithLocale(locale objc.IObject) IString {
	rv := objc.Send[String](d_.ID, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}/* debug [instance_methods/method]: DescriptionWithLocale */


// Divides the number by another given number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/dividing(by:)
func (d_ DecimalNumber) DecimalNumberByDividingBy(decimalNumber IDecimalNumber) IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByDividingBy:"), decimalNumber)
	return rv
}/* debug [instance_methods/method]: DecimalNumberByDividingBy */


// Divides this number by another given number using the specified behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/dividing(by:withBehavior:)
func (d_ DecimalNumber) DecimalNumberByDividingByWithBehavior(decimalNumber IDecimalNumber, behavior unsafe.Pointer) IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByDividingBy:withBehavior:"), decimalNumber, behavior)
	return rv
}/* debug [instance_methods/method]: DecimalNumberByDividingByWithBehavior */


// Multiplies the number by another given number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/multiplying(by:)
func (d_ DecimalNumber) DecimalNumberByMultiplyingBy(decimalNumber IDecimalNumber) IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByMultiplyingBy:"), decimalNumber)
	return rv
}/* debug [instance_methods/method]: DecimalNumberByMultiplyingBy */


// Multiplies this number by another given number using the specified behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/multiplying(by:withBehavior:)
func (d_ DecimalNumber) DecimalNumberByMultiplyingByWithBehavior(decimalNumber IDecimalNumber, behavior unsafe.Pointer) IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByMultiplyingBy:withBehavior:"), decimalNumber, behavior)
	return rv
}/* debug [instance_methods/method]: DecimalNumberByMultiplyingByWithBehavior */


// Multiplies the number by 10 raised to the given power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/multiplying(byPowerOf10:)
func (d_ DecimalNumber) DecimalNumberByMultiplyingByPowerOf10(power objectivec.IObject) IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByMultiplyingByPowerOf10:"), power)
	return rv
}/* debug [instance_methods/method]: DecimalNumberByMultiplyingByPowerOf10 */


// Multiplies the number by 10 raised to the given power using the specified behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/multiplying(byPowerOf10:withBehavior:)
func (d_ DecimalNumber) DecimalNumberByMultiplyingByPowerOf10WithBehavior(power objectivec.IObject, behavior unsafe.Pointer) IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByMultiplyingByPowerOf10:withBehavior:"), power, behavior)
	return rv
}/* debug [instance_methods/method]: DecimalNumberByMultiplyingByPowerOf10WithBehavior */


// Raises the number to a given power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/raising(toPower:)
func (d_ DecimalNumber) DecimalNumberByRaisingToPower(power uint) IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByRaisingToPower:"), power)
	return rv
}/* debug [instance_methods/method]: DecimalNumberByRaisingToPower */


// Raises the number to a given power using the specified behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/raising(toPower:withBehavior:)
func (d_ DecimalNumber) DecimalNumberByRaisingToPowerWithBehavior(power uint, behavior unsafe.Pointer) IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByRaisingToPower:withBehavior:"), power, behavior)
	return rv
}/* debug [instance_methods/method]: DecimalNumberByRaisingToPowerWithBehavior */


// Returns a rounded version of the decimal number using the specified rounding behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/rounding(accordingToBehavior:)
func (d_ DecimalNumber) DecimalNumberByRoundingAccordingToBehavior(behavior unsafe.Pointer) IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberByRoundingAccordingToBehavior:"), behavior)
	return rv
}/* debug [instance_methods/method]: DecimalNumberByRoundingAccordingToBehavior */


// Subtracts another given number from this one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/subtracting(_:)
func (d_ DecimalNumber) DecimalNumberBySubtracting(decimalNumber IDecimalNumber) IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberBySubtracting:"), decimalNumber)
	return rv
}/* debug [instance_methods/method]: DecimalNumberBySubtracting */


// Subtracts this a given number from this one using the specified behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/subtracting(_:withBehavior:)
func (d_ DecimalNumber) DecimalNumberBySubtractingWithBehavior(decimalNumber IDecimalNumber, behavior unsafe.Pointer) IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("decimalNumberBySubtracting:withBehavior:"), decimalNumber, behavior)
	return rv
}/* debug [instance_methods/method]: DecimalNumberBySubtractingWithBehavior */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DecimalNumber */

// The decimal number’s value, expressed as an structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/decimalValue
func (d_ DecimalNumber) DecimalValue() objc.IObject /* cross-framework: Decimal */ {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("decimalValue"))
	return rv
}/* debug [instance_properties/getter]: decimalValue */


// The way arithmetic methods round off and handle error conditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/defaultBehavior
func (d_ DecimalNumber) DefaultBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("defaultBehavior"))
	return rv
}/* debug [instance_properties/getter]: defaultBehavior */


// The way arithmetic methods round off and handle error conditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/defaultBehavior
func (d_ DecimalNumber) SetDefaultBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDefaultBehavior:"), value)
}/* debug [instance_properties/setter]: defaultBehavior */


// The decimal number’s closest approximate value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/doubleValue
func (d_ DecimalNumber) DoubleValue() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("doubleValue"))
	return rv
}/* debug [instance_properties/getter]: doubleValue */


// Returns the largest possible value of a decimal number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/maximum
func (d_ DecimalNumber) MaximumDecimalNumber() IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("maximumDecimalNumber"))
	return rv
}/* debug [instance_properties/getter]: maximumDecimalNumber */


// Returns the smallest possible value of a decimal number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/minimum
func (d_ DecimalNumber) MinimumDecimalNumber() IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("minimumDecimalNumber"))
	return rv
}/* debug [instance_properties/getter]: minimumDecimalNumber */


// A decimal number that specifies no number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/notANumber
func (d_ DecimalNumber) NotANumber() IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("notANumber"))
	return rv
}/* debug [instance_properties/getter]: notANumber */


// A C string containing the Objective-C type for the data contained in the decimal number object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/objCType
func (d_ DecimalNumber) ObjCType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("objCType"))
	return rv
}/* debug [instance_properties/getter]: objCType */


// A decimal number equivalent to the number 1.0.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/one
func (d_ DecimalNumber) One() IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("one"))
	return rv
}/* debug [instance_properties/getter]: one */


// A decimal number equivalent to the number 0.0.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/zero
func (d_ DecimalNumber) Zero() IDecimalNumber {
	rv := objc.Send[DecimalNumber](d_.ID, objc.Sel("zero"))
	return rv
}/* debug [instance_properties/getter]: zero */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDecimalNumber */



// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSNumber */


/* debug [class_header]: Header for NSNumber */
// The class instance for the [Number] class.
var (
	NumberClass     _NumberClass
	NumberClassOnce sync.Once
)

func getNumberClass() _NumberClass {
	NumberClassOnce.Do(func() {
		NumberClass = _NumberClass{objc.GetClass("NSNumber")}
	})
	return NumberClass
}

type _NumberClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Number */
// An interface definition for the [Number] class.
type INumber interface {
	IValue
	
/* debug [class_interface_properties]: Properties for Number */
	// properties:
	BoolValue() bool
	DecimalValue() objc.IObject /* cross-framework: Decimal */
	DoubleValue() float64
	FloatValue() float32
	ShortValue() objectivec.IObject
	IntValue() int
	LongLongValue() objectivec.IObject
	CharValue() objectivec.IObject
	IntegerValue() int
	LongValue() objectivec.IObject
	StringValue() IString
	UnsignedShortValue() objectivec.IObject
	UnsignedIntValue() objectivec.IObject
	UnsignedLongLongValue() uint64
	UnsignedCharValue() objectivec.IObject
	UnsignedIntegerValue() uint
	UnsignedLongValue() objectivec.IObject
	Int16Value() objectivec.IObject
	SetInt16Value(value objectivec.IObject)
	Int32Value() objectivec.IObject
	SetInt32Value(value objectivec.IObject)
	Int64Value() objectivec.IObject
	SetInt64Value(value objectivec.IObject)
	Int8Value() objectivec.IObject
	SetInt8Value(value objectivec.IObject)
	Uint16Value() objectivec.IObject
	SetUint16Value(value objectivec.IObject)
	Uint32Value() objectivec.IObject
	SetUint32Value(value objectivec.IObject)
	Uint64Value() uint64
	SetUint64Value(value uint64)
	Uint8Value() objectivec.IObject
	SetUint8Value(value objectivec.IObject)
	UintValue() uint
	SetUintValue(value uint)
	ObjCType() objectivec.IObject
	SetObjCType(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Number */
	// methods:
	Compare(otherNumber INumber) ComparisonResult
	DescriptionWithLocale(locale objc.IObject) IString
	IsEqualToNumber(number INumber) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Number */
// Alloc allocates a new instance without initialization.
func (nc _NumberClass) Alloc() Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NumberClass) New() Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ Number) Init() Number {
	rv := objc.Send[Number](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ Number) Autorelease() Number {
	rv := objc.Send[Number](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNumber creates a new Number instance.
func NewNumber() Number {
	return getNumberClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Number */
// An object wrapper for primitive scalar numeric values.
//
// is a subclass of that offers a value as any C scalar (numeric) type. It defines a set of methods specifically for setting and accessing the value as a signed or unsigned , , , , , , or or as a . (Note that number objects do not necessarily preserve the type they are created with.) It also defines a method to determine the ordering of two objects. is “toll-free bridged” with its Core Foundation counterparts: for integer and floating point values, and for Boolean values. See for more information on toll-free bridging.


// An object wrapper for primitive scalar numeric values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber
type Number struct {
	Value
}

// NumberFrom constructs a [Number] from an unsafe.Pointer.
//
// An object wrapper for primitive scalar numeric values.
func NumberFrom(ptr unsafe.Pointer) Number {
	return Number{
		Value: ValueFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Number */

// Returns an object initialized to contain a given value, treated as a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-1ojz2
func NewNumberWithBool(value bool) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithBool:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNumberWithBool */


// Returns an object initialized to contain a given value, treated as a signed .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-8krjs
func NewNumberWithChar(value objectivec.IObject) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithChar:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNumberWithChar */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(coder:)
func NewNumberWithCoder(coder ICoder) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNumberWithCoder */


// Returns an object initialized to contain , treated as a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-15chk
func NewNumberWithDouble(value float64) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithDouble:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNumberWithDouble */


// Returns an object initialized to contain a given value, treated as a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-2vlwk
func NewNumberWithFloat(value float32) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithFloat:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNumberWithFloat */


// Returns an object initialized to contain a given value, treated as a signed .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-7jvmg
func NewNumberWithInt(value int) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithInt:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNumberWithInt */


// Returns an object initialized to contain a given value, treated as an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-5jcjl
func NewNumberWithInteger(value int) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithInteger:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNumberWithInteger */


// Returns an object initialized to contain a given value, treated as a signed .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/initWithLong:
func NewNumberWithLong(value objectivec.IObject) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithLong:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNumberWithLong */


// Returns an object initialized to contain , treated as a signed .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-40ad0
func NewNumberWithLongLong(value objectivec.IObject) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithLongLong:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNumberWithLongLong */


// Returns an object initialized to contain a given value, treated as a signed .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-16drx
func NewNumberWithShort(value objectivec.IObject) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithShort:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNumberWithShort */


// Returns an object initialized to contain a given value, treated as an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-8se67
func NewNumberWithUnsignedChar(value objectivec.IObject) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithUnsignedChar:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNumberWithUnsignedChar */


// Returns an object initialized to contain a given value, treated as an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-47coa
func NewNumberWithUnsignedInt(value objectivec.IObject) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithUnsignedInt:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNumberWithUnsignedInt */


// Returns an object initialized to contain a given value, treated as an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-3l4ek
func NewNumberWithUnsignedInteger(value uint) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithUnsignedInteger:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNumberWithUnsignedInteger */


// Returns an object initialized to contain a given value, treated as an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/initWithUnsignedLong:
func NewNumberWithUnsignedLong(value objectivec.IObject) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithUnsignedLong:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNumberWithUnsignedLong */


// Returns an object initialized to contain a given value, treated as an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-43lc7
func NewNumberWithUnsignedLongLong(value uint64) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithUnsignedLongLong:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNumberWithUnsignedLongLong */


// Returns an object initialized to contain a given value, treated as an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-87y9m
func NewNumberWithUnsignedShort(value objectivec.IObject) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithUnsignedShort:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNumberWithUnsignedShort */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Number */

// Creates and returns an object containing a given value, treating it as a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithBool:
func (nc _NumberClass) NumberWithBool(value bool) INumber {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithBool:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NumberWithBool) */


// Creates and returns an object containing a given value, treating it as a signed .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithChar:
func (nc _NumberClass) NumberWithChar(value objectivec.IObject) INumber {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithChar:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NumberWithChar) */


// Creates and returns an object containing a given value, treating it as a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithDouble:
func (nc _NumberClass) NumberWithDouble(value float64) INumber {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithDouble:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NumberWithDouble) */


// Creates and returns an object containing a given value, treating it as a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithFloat:
func (nc _NumberClass) NumberWithFloat(value float32) INumber {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithFloat:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NumberWithFloat) */


// Creates and returns an object containing a given value, treating it as a signed .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithInt:
func (nc _NumberClass) NumberWithInt(value int) INumber {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithInt:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NumberWithInt) */


// Creates and returns an object containing a given value, treating it as an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithInteger:
func (nc _NumberClass) NumberWithInteger(value int) INumber {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithInteger:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NumberWithInteger) */


// Creates and returns an object containing a given value, treating it as a signed .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithLong:
func (nc _NumberClass) NumberWithLong(value objectivec.IObject) INumber {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithLong:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NumberWithLong) */


// Creates and returns an object containing a given value, treating it as a signed .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithLongLong:
func (nc _NumberClass) NumberWithLongLong(value objectivec.IObject) INumber {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithLongLong:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NumberWithLongLong) */


// Creates and returns an object containing , treating it as a signed .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithShort:
func (nc _NumberClass) NumberWithShort(value objectivec.IObject) INumber {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithShort:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NumberWithShort) */


// Creates and returns an object containing a given value, treating it as an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedChar:
func (nc _NumberClass) NumberWithUnsignedChar(value objectivec.IObject) INumber {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithUnsignedChar:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NumberWithUnsignedChar) */


// Creates and returns an object containing a given value, treating it as an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedInt:
func (nc _NumberClass) NumberWithUnsignedInt(value objectivec.IObject) INumber {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithUnsignedInt:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NumberWithUnsignedInt) */


// Creates and returns an object containing a given value, treating it as an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedInteger:
func (nc _NumberClass) NumberWithUnsignedInteger(value uint) INumber {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithUnsignedInteger:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NumberWithUnsignedInteger) */


// Creates and returns an object containing a given value, treating it as an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedLong:
func (nc _NumberClass) NumberWithUnsignedLong(value objectivec.IObject) INumber {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithUnsignedLong:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NumberWithUnsignedLong) */


// Creates and returns an object containing a given value, treating it as an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedLongLong:
func (nc _NumberClass) NumberWithUnsignedLongLong(value uint64) INumber {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithUnsignedLongLong:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NumberWithUnsignedLongLong) */


// Creates and returns an object containing a given value, treating it as an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedShort:
func (nc _NumberClass) NumberWithUnsignedShort(value objectivec.IObject) INumber {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithUnsignedShort:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NumberWithUnsignedShort) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Number */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Number */

// Returns an value that indicates whether the number object’s value is greater than, equal to, or less than a given number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/compare(_:)
func (n_ Number) Compare(otherNumber INumber) ComparisonResult {
	rv := objc.Send[ComparisonResult](n_.ID, objc.Sel("compare:"), otherNumber)
	return rv
}/* debug [instance_methods/method]: Compare */


// Returns a string that represents the contents of the number object for a given locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/description(withLocale:)
func (n_ Number) DescriptionWithLocale(locale objc.IObject) IString {
	rv := objc.Send[String](n_.ID, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}/* debug [instance_methods/method]: DescriptionWithLocale */


// Returns a Boolean value that indicates whether the number object’s value and a given number are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/isEqual(to:)
func (n_ Number) IsEqualToNumber(number INumber) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isEqualToNumber:"), number)
	return rv
}/* debug [instance_methods/method]: IsEqualToNumber */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Number */

// The number object’s value expressed as a Boolean value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/boolValue
func (n_ Number) BoolValue() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("boolValue"))
	return rv
}/* debug [instance_properties/getter]: boolValue */


// The number object’s value expressed as an structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/decimalValue
func (n_ Number) DecimalValue() objc.IObject /* cross-framework: Decimal */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("decimalValue"))
	return rv
}/* debug [instance_properties/getter]: decimalValue */


// The number object’s value expressed as a , converted as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/doubleValue
func (n_ Number) DoubleValue() float64 {
	rv := objc.Send[float64](n_.ID, objc.Sel("doubleValue"))
	return rv
}/* debug [instance_properties/getter]: doubleValue */


// The number object’s value expressed as a , converted as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/floatValue
func (n_ Number) FloatValue() float32 {
	rv := objc.Send[float32](n_.ID, objc.Sel("floatValue"))
	return rv
}/* debug [instance_properties/getter]: floatValue */


// The number object’s value expressed as a , converted as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/int16Value
func (n_ Number) ShortValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("shortValue"))
	return rv
}/* debug [instance_properties/getter]: shortValue */


// The number object’s value expressed as an , converted as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/int32Value
func (n_ Number) IntValue() int {
	rv := objc.Send[int](n_.ID, objc.Sel("intValue"))
	return rv
}/* debug [instance_properties/getter]: intValue */


// The number object’s value expressed as a , converted as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/int64Value
func (n_ Number) LongLongValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("longLongValue"))
	return rv
}/* debug [instance_properties/getter]: longLongValue */


// The number object’s value expressed as a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/int8Value
func (n_ Number) CharValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("charValue"))
	return rv
}/* debug [instance_properties/getter]: charValue */


// The number object’s value expressed as an object, converted as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/intValue-95zzp
func (n_ Number) IntegerValue() int {
	rv := objc.Send[int](n_.ID, objc.Sel("integerValue"))
	return rv
}/* debug [instance_properties/getter]: integerValue */


// The number object’s value expressed as a , converted as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/longValue
func (n_ Number) LongValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("longValue"))
	return rv
}/* debug [instance_properties/getter]: longValue */


// The number object’s value expressed as a human-readable string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/stringValue
func (n_ Number) StringValue() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("stringValue"))
	return rv
}/* debug [instance_properties/getter]: stringValue */


// The number object’s value expressed as an unsigned , converted as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/uint16Value
func (n_ Number) UnsignedShortValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("unsignedShortValue"))
	return rv
}/* debug [instance_properties/getter]: unsignedShortValue */


// The number object’s value expressed as an unsigned , converted as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/uint32Value
func (n_ Number) UnsignedIntValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("unsignedIntValue"))
	return rv
}/* debug [instance_properties/getter]: unsignedIntValue */


// The number object’s value expressed as an unsigned , converted as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/uint64Value
func (n_ Number) UnsignedLongLongValue() uint64 {
	rv := objc.Send[uint64](n_.ID, objc.Sel("unsignedLongLongValue"))
	return rv
}/* debug [instance_properties/getter]: unsignedLongLongValue */


// The number object’s value expressed as an unsigned , converted as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/uint8Value
func (n_ Number) UnsignedCharValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("unsignedCharValue"))
	return rv
}/* debug [instance_properties/getter]: unsignedCharValue */


// The number object’s value expressed as an object, converted as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/uintValue
func (n_ Number) UnsignedIntegerValue() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("unsignedIntegerValue"))
	return rv
}/* debug [instance_properties/getter]: unsignedIntegerValue */


// The number object’s value expressed as an unsigned , converted as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/unsignedLongValue
func (n_ Number) UnsignedLongValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("unsignedLongValue"))
	return rv
}/* debug [instance_properties/getter]: unsignedLongValue */


// The number object’s value expressed as a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int16value
func (n_ Number) Int16Value() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("int16Value"))
	return rv
}/* debug [instance_properties/getter]: int16Value */


// The number object’s value expressed as a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int16value
func (n_ Number) SetInt16Value(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInt16Value:"), value)
}/* debug [instance_properties/setter]: int16Value */


// The number object’s value expressed as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int32value
func (n_ Number) Int32Value() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("int32Value"))
	return rv
}/* debug [instance_properties/getter]: int32Value */


// The number object’s value expressed as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int32value
func (n_ Number) SetInt32Value(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInt32Value:"), value)
}/* debug [instance_properties/setter]: int32Value */


// The number object’s value expressed as a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int64value
func (n_ Number) Int64Value() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("int64Value"))
	return rv
}/* debug [instance_properties/getter]: int64Value */


// The number object’s value expressed as a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int64value
func (n_ Number) SetInt64Value(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInt64Value:"), value)
}/* debug [instance_properties/setter]: int64Value */


// The number object’s value expressed as a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int8value
func (n_ Number) Int8Value() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("int8Value"))
	return rv
}/* debug [instance_properties/getter]: int8Value */


// The number object’s value expressed as a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int8value
func (n_ Number) SetInt8Value(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInt8Value:"), value)
}/* debug [instance_properties/setter]: int8Value */


// The number object’s value expressed as an unsigned
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint16value
func (n_ Number) Uint16Value() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("uint16Value"))
	return rv
}/* debug [instance_properties/getter]: uint16Value */


// The number object’s value expressed as an unsigned
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint16value
func (n_ Number) SetUint16Value(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUint16Value:"), value)
}/* debug [instance_properties/setter]: uint16Value */


// The number object’s value expressed as an unsigned
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint32value
func (n_ Number) Uint32Value() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("uint32Value"))
	return rv
}/* debug [instance_properties/getter]: uint32Value */


// The number object’s value expressed as an unsigned
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint32value
func (n_ Number) SetUint32Value(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUint32Value:"), value)
}/* debug [instance_properties/setter]: uint32Value */


// The number object’s value expressed as an unsigned
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint64value
func (n_ Number) Uint64Value() uint64 {
	rv := objc.Send[uint64](n_.ID, objc.Sel("uint64Value"))
	return rv
}/* debug [instance_properties/getter]: uint64Value */


// The number object’s value expressed as an unsigned
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint64value
func (n_ Number) SetUint64Value(value uint64) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUint64Value:"), value)
}/* debug [instance_properties/setter]: uint64Value */


// The number object’s value expressed as an unsigned
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint8value
func (n_ Number) Uint8Value() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("uint8Value"))
	return rv
}/* debug [instance_properties/getter]: uint8Value */


// The number object’s value expressed as an unsigned
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint8value
func (n_ Number) SetUint8Value(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUint8Value:"), value)
}/* debug [instance_properties/setter]: uint8Value */


// The number object’s value expressed as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uintvalue
func (n_ Number) UintValue() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("uintValue"))
	return rv
}/* debug [instance_properties/getter]: uintValue */


// The number object’s value expressed as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uintvalue
func (n_ Number) SetUintValue(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUintValue:"), value)
}/* debug [instance_properties/setter]: uintValue */


// A C string containing the Objective-C type of the data contained in the value object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/objctype
func (n_ Number) ObjCType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("objCType"))
	return rv
}/* debug [instance_properties/getter]: objCType */


// A C string containing the Objective-C type of the data contained in the value object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/objctype
func (n_ Number) SetObjCType(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setObjCType:"), value)
}/* debug [instance_properties/setter]: objCType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSNumber */



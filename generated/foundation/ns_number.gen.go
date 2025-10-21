// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [Number] class.
type INumber interface {
	IValue
	Compare(otherNumber Number) unsafe.Pointer
	DescriptionWithLocale(locale objc.ID) string
	IsEqualToNumber(number Number) bool
}

// An object wrapper for primitive scalar numeric values.
//
// is a subclass of that offers a value as any C scalar (numeric) type. It defines a set of methods specifically for setting and accessing the value as a signed or unsigned , , , , , , or or as a . (Note that number objects do not necessarily preserve the type they are created with.) It also defines a method to determine the ordering of two objects. is “toll-free bridged” with its Core Foundation counterparts: for integer and floating point values, and for Boolean values. See for more information on toll-free bridging.
//
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

// Alloc allocates a new instance without initialization.
func (nc _NumberClass) Alloc() Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Returns an object initialized to contain a given value, treated as a .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-1ojz2
func NewNumberWithBool(value bool) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithBool:"), value)
	rv.Autorelease()
	return rv
}



// Returns an object initialized to contain a given value, treated as a signed .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-8krjs
func NewNumberWithChar(value unsafe.Pointer) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithChar:"), value)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(coder:)
func NewNumberWithCoder(coder unsafe.Pointer) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Returns an object initialized to contain , treated as a .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-15chk
func NewNumberWithDouble(value unsafe.Pointer) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithDouble:"), value)
	rv.Autorelease()
	return rv
}



// Returns an object initialized to contain a given value, treated as a .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-2vlwk
func NewNumberWithFloat(value unsafe.Pointer) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithFloat:"), value)
	rv.Autorelease()
	return rv
}



// Returns an object initialized to contain a given value, treated as a signed .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-7jvmg
func NewNumberWithInt(value unsafe.Pointer) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithInt:"), value)
	rv.Autorelease()
	return rv
}



// Returns an object initialized to contain a given value, treated as an .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-5jcjl
func NewNumberWithInteger(value int) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithInteger:"), value)
	rv.Autorelease()
	return rv
}



// Returns an object initialized to contain a given value, treated as a signed .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/initWithLong:
func NewNumberWithLong(value unsafe.Pointer) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithLong:"), value)
	rv.Autorelease()
	return rv
}



// Returns an object initialized to contain , treated as a signed .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-40ad0
func NewNumberWithLongLong(value unsafe.Pointer) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithLongLong:"), value)
	rv.Autorelease()
	return rv
}



// Returns an object initialized to contain a given value, treated as a signed .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-16drx
func NewNumberWithShort(value unsafe.Pointer) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithShort:"), value)
	rv.Autorelease()
	return rv
}



// Returns an object initialized to contain a given value, treated as an .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-8se67
func NewNumberWithUnsignedChar(value unsafe.Pointer) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithUnsignedChar:"), value)
	rv.Autorelease()
	return rv
}



// Returns an object initialized to contain a given value, treated as an .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-47coa
func NewNumberWithUnsignedInt(value unsafe.Pointer) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithUnsignedInt:"), value)
	rv.Autorelease()
	return rv
}



// Returns an object initialized to contain a given value, treated as an .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-3l4ek
func NewNumberWithUnsignedInteger(value uint) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithUnsignedInteger:"), value)
	rv.Autorelease()
	return rv
}



// Returns an object initialized to contain a given value, treated as an .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/initWithUnsignedLong:
func NewNumberWithUnsignedLong(value unsafe.Pointer) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithUnsignedLong:"), value)
	rv.Autorelease()
	return rv
}



// Returns an object initialized to contain a given value, treated as an .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-43lc7
func NewNumberWithUnsignedLongLong(value uint64) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithUnsignedLongLong:"), value)
	rv.Autorelease()
	return rv
}



// Returns an object initialized to contain a given value, treated as an .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-87y9m
func NewNumberWithUnsignedShort(value unsafe.Pointer) Number {
	instance := getNumberClass().Alloc()
	rv := objc.Send[Number](instance.ID, objc.Sel("initWithUnsignedShort:"), value)
	rv.Autorelease()
	return rv
}


// Creates and returns an object containing a given value, treating it as a .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithBool:
func (nc _NumberClass) NumberWithBool(value bool) Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithBool:"), value)
	return rv
}

// Creates and returns an object containing a given value, treating it as a signed .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithChar:
func (nc _NumberClass) NumberWithChar(value unsafe.Pointer) Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithChar:"), value)
	return rv
}

// Creates and returns an object containing a given value, treating it as a .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithDouble:
func (nc _NumberClass) NumberWithDouble(value unsafe.Pointer) Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithDouble:"), value)
	return rv
}

// Creates and returns an object containing a given value, treating it as a .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithFloat:
func (nc _NumberClass) NumberWithFloat(value unsafe.Pointer) Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithFloat:"), value)
	return rv
}

// Creates and returns an object containing a given value, treating it as a signed .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithInt:
func (nc _NumberClass) NumberWithInt(value unsafe.Pointer) Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithInt:"), value)
	return rv
}

// Creates and returns an object containing a given value, treating it as an .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithInteger:
func (nc _NumberClass) NumberWithInteger(value int) Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithInteger:"), value)
	return rv
}

// Creates and returns an object containing a given value, treating it as a signed .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithLong:
func (nc _NumberClass) NumberWithLong(value unsafe.Pointer) Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithLong:"), value)
	return rv
}

// Creates and returns an object containing a given value, treating it as a signed .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithLongLong:
func (nc _NumberClass) NumberWithLongLong(value unsafe.Pointer) Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithLongLong:"), value)
	return rv
}

// Creates and returns an object containing , treating it as a signed .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithShort:
func (nc _NumberClass) NumberWithShort(value unsafe.Pointer) Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithShort:"), value)
	return rv
}

// Creates and returns an object containing a given value, treating it as an .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedChar:
func (nc _NumberClass) NumberWithUnsignedChar(value unsafe.Pointer) Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithUnsignedChar:"), value)
	return rv
}

// Creates and returns an object containing a given value, treating it as an .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedInt:
func (nc _NumberClass) NumberWithUnsignedInt(value unsafe.Pointer) Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithUnsignedInt:"), value)
	return rv
}

// Creates and returns an object containing a given value, treating it as an .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedInteger:
func (nc _NumberClass) NumberWithUnsignedInteger(value uint) Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithUnsignedInteger:"), value)
	return rv
}

// Creates and returns an object containing a given value, treating it as an .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedLong:
func (nc _NumberClass) NumberWithUnsignedLong(value unsafe.Pointer) Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithUnsignedLong:"), value)
	return rv
}

// Creates and returns an object containing a given value, treating it as an .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedLongLong:
func (nc _NumberClass) NumberWithUnsignedLongLong(value uint64) Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithUnsignedLongLong:"), value)
	return rv
}

// Creates and returns an object containing a given value, treating it as an .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedShort:
func (nc _NumberClass) NumberWithUnsignedShort(value unsafe.Pointer) Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("numberWithUnsignedShort:"), value)
	return rv
}

// Returns an value that indicates whether the number object’s value is greater than, equal to, or less than a given number.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/compare(_:)
func (n_ Number) Compare(otherNumber Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("compare:"), otherNumber)
	return rv
}

// Returns a string that represents the contents of the number object for a given locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/description(withLocale:)
func (n_ Number) DescriptionWithLocale(locale objc.ID) string {
	rv := objc.Send[string](n_.ID, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}

// Returns a Boolean value that indicates whether the number object’s value and a given number are equal.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/isEqual(to:)
func (n_ Number) IsEqualToNumber(number Number) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isEqualToNumber:"), number)
	return rv
}

// The number object’s value expressed as an unsigned
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint8value
func (n_ Number) Uint8Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("uint8Value"))
	return rv
}


// SetUint8Value sets the value of the uint8Value property.
// The number object’s value expressed as an unsigned

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint8value
func (n_ Number) SetUint8Value(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUint8Value:"), value)
}

// The number object’s value expressed as an unsigned
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint16value
func (n_ Number) Uint16Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("uint16Value"))
	return rv
}


// SetUint16Value sets the value of the uint16Value property.
// The number object’s value expressed as an unsigned

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint16value
func (n_ Number) SetUint16Value(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUint16Value:"), value)
}

// The number object’s value expressed as an
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int32value
func (n_ Number) Int32Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("int32Value"))
	return rv
}


// SetInt32Value sets the value of the int32Value property.
// The number object’s value expressed as an

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int32value
func (n_ Number) SetInt32Value(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInt32Value:"), value)
}

// The number object’s value expressed as a
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int64value
func (n_ Number) Int64Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("int64Value"))
	return rv
}


// SetInt64Value sets the value of the int64Value property.
// The number object’s value expressed as a

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int64value
func (n_ Number) SetInt64Value(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInt64Value:"), value)
}

// The number object’s value expressed as an unsigned
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint32value
func (n_ Number) Uint32Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("uint32Value"))
	return rv
}


// SetUint32Value sets the value of the uint32Value property.
// The number object’s value expressed as an unsigned

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint32value
func (n_ Number) SetUint32Value(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUint32Value:"), value)
}

// A C string containing the Objective-C type of the data contained in the value object.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/objctype
func (n_ Number) ObjCType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("objCType"))
	return rv
}


// SetObjCType sets the value of the objCType property.
// A C string containing the Objective-C type of the data contained in the value object.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/objctype
func (n_ Number) SetObjCType(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setObjCType:"), value)
}

// The number object’s value expressed as a
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int16value
func (n_ Number) Int16Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("int16Value"))
	return rv
}


// SetInt16Value sets the value of the int16Value property.
// The number object’s value expressed as a

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int16value
func (n_ Number) SetInt16Value(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInt16Value:"), value)
}

// The number object’s value expressed as an unsigned
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint64value
func (n_ Number) Uint64Value() uint64 {
	rv := objc.Send[uint64](n_.ID, objc.Sel("uint64Value"))
	return rv
}


// SetUint64Value sets the value of the uint64Value property.
// The number object’s value expressed as an unsigned

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint64value
func (n_ Number) SetUint64Value(value uint64) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUint64Value:"), value)
}

// The number object’s value expressed as an
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uintvalue
func (n_ Number) UintValue() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("uintValue"))
	return rv
}


// SetUintValue sets the value of the uintValue property.
// The number object’s value expressed as an

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uintvalue
func (n_ Number) SetUintValue(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUintValue:"), value)
}

// The number object’s value expressed as a
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int8value
func (n_ Number) Int8Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("int8Value"))
	return rv
}


// SetInt8Value sets the value of the int8Value property.
// The number object’s value expressed as a

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int8value
func (n_ Number) SetInt8Value(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInt8Value:"), value)
}

// The number object’s value expressed as a Boolean value.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/boolValue
func (n_ Number) BoolValue() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("boolValue"))
	return rv
}

// The number object’s value expressed as an structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/decimalValue
func (n_ Number) DecimalValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("decimalValue"))
	return rv
}

// The number object’s value expressed as a , converted as necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/doubleValue
func (n_ Number) DoubleValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("doubleValue"))
	return rv
}

// The number object’s value expressed as a , converted as necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/floatValue
func (n_ Number) FloatValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("floatValue"))
	return rv
}

// The number object’s value expressed as a , converted as necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/int16Value
func (n_ Number) ShortValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("shortValue"))
	return rv
}

// The number object’s value expressed as an , converted as necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/int32Value
func (n_ Number) IntValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("intValue"))
	return rv
}

// The number object’s value expressed as a , converted as necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/int64Value
func (n_ Number) LongLongValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("longLongValue"))
	return rv
}

// The number object’s value expressed as a .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/int8Value
func (n_ Number) CharValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("charValue"))
	return rv
}

// The number object’s value expressed as an object, converted as necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/intValue-95zzp
func (n_ Number) IntegerValue() int {
	rv := objc.Send[int](n_.ID, objc.Sel("integerValue"))
	return rv
}

// The number object’s value expressed as a , converted as necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/longValue
func (n_ Number) LongValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("longValue"))
	return rv
}

// The number object’s value expressed as a human-readable string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/stringValue
func (n_ Number) StringValue() string {
	rv := objc.Send[string](n_.ID, objc.Sel("stringValue"))
	return rv
}

// The number object’s value expressed as an unsigned , converted as necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/uint16Value
func (n_ Number) UnsignedShortValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("unsignedShortValue"))
	return rv
}

// The number object’s value expressed as an unsigned , converted as necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/uint32Value
func (n_ Number) UnsignedIntValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("unsignedIntValue"))
	return rv
}

// The number object’s value expressed as an unsigned , converted as necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/uint64Value
func (n_ Number) UnsignedLongLongValue() uint64 {
	rv := objc.Send[uint64](n_.ID, objc.Sel("unsignedLongLongValue"))
	return rv
}

// The number object’s value expressed as an unsigned , converted as necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/uint8Value
func (n_ Number) UnsignedCharValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("unsignedCharValue"))
	return rv
}

// The number object’s value expressed as an object, converted as necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/uintValue
func (n_ Number) UnsignedIntegerValue() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("unsignedIntegerValue"))
	return rv
}

// The number object’s value expressed as an unsigned , converted as necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/unsignedLongValue
func (n_ Number) UnsignedLongValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("unsignedLongValue"))
	return rv
}



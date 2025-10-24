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
	// properties:
	BoolValue() bool
	SetBoolValue(value bool)
	DecimalValue() objc.IObject /* cross-framework: Decimal */
	SetDecimalValue(value objc.IObject /* cross-framework: Decimal */)
	DoubleValue() float64
	SetDoubleValue(value float64)
	FloatValue() float32
	SetFloatValue(value float32)
	Int16Value() unsafe.Pointer
	SetInt16Value(value unsafe.Pointer)
	Int32Value() unsafe.Pointer
	SetInt32Value(value unsafe.Pointer)
	Int64Value() unsafe.Pointer
	SetInt64Value(value unsafe.Pointer)
	Int8Value() unsafe.Pointer
	SetInt8Value(value unsafe.Pointer)
	IntValue() int
	SetIntValue(value int)
	StringValue() IString
	SetStringValue(value IString)
	Uint16Value() unsafe.Pointer
	SetUint16Value(value unsafe.Pointer)
	Uint32Value() unsafe.Pointer
	SetUint32Value(value unsafe.Pointer)
	Uint64Value() uint64
	SetUint64Value(value uint64)
	Uint8Value() unsafe.Pointer
	SetUint8Value(value unsafe.Pointer)
	UintValue() uint
	SetUintValue(value uint)
	ObjCType() unsafe.Pointer
	SetObjCType(value unsafe.Pointer)
	// methods:
}

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



// The number object’s value expressed as a Boolean value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/boolvalue
func (n_ Number) BoolValue() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("boolValue"))
	return rv
}


// The number object’s value expressed as a Boolean value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/boolvalue
func (n_ Number) SetBoolValue(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setBoolValue:"), value)
}


// The number object’s value expressed as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/decimalvalue
func (n_ Number) DecimalValue() objc.IObject /* cross-framework: Decimal */ {
	rv := objc.Send[Decimal](n_.ID, objc.Sel("decimalValue"))
	return rv
}


// The number object’s value expressed as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/decimalvalue
func (n_ Number) SetDecimalValue(value objc.IObject /* cross-framework: Decimal */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDecimalValue:"), value)
}


// The number object’s value expressed as a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/doublevalue
func (n_ Number) DoubleValue() float64 {
	rv := objc.Send[float64](n_.ID, objc.Sel("doubleValue"))
	return rv
}


// The number object’s value expressed as a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/doublevalue
func (n_ Number) SetDoubleValue(value float64) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDoubleValue:"), value)
}


// The number object’s value expressed as a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/floatvalue
func (n_ Number) FloatValue() float32 {
	rv := objc.Send[float32](n_.ID, objc.Sel("floatValue"))
	return rv
}


// The number object’s value expressed as a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/floatvalue
func (n_ Number) SetFloatValue(value float32) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFloatValue:"), value)
}


// The number object’s value expressed as a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int16value
func (n_ Number) Int16Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("int16Value"))
	return rv
}


// The number object’s value expressed as a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int16value
func (n_ Number) SetInt16Value(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInt16Value:"), value)
}


// The number object’s value expressed as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int32value
func (n_ Number) Int32Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("int32Value"))
	return rv
}


// The number object’s value expressed as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int32value
func (n_ Number) SetInt32Value(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInt32Value:"), value)
}


// The number object’s value expressed as a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int64value
func (n_ Number) Int64Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("int64Value"))
	return rv
}


// The number object’s value expressed as a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int64value
func (n_ Number) SetInt64Value(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInt64Value:"), value)
}


// The number object’s value expressed as a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int8value
func (n_ Number) Int8Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("int8Value"))
	return rv
}


// The number object’s value expressed as a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/int8value
func (n_ Number) SetInt8Value(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInt8Value:"), value)
}


// The number object’s value expressed as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/intvalue-95zzp
func (n_ Number) IntValue() int {
	rv := objc.Send[int](n_.ID, objc.Sel("intValue"))
	return rv
}


// The number object’s value expressed as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/intvalue-95zzp
func (n_ Number) SetIntValue(value int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIntValue:"), value)
}


// The number object’s value expressed as a human-readable string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/stringvalue
func (n_ Number) StringValue() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("stringValue"))
	return rv
}


// The number object’s value expressed as a human-readable string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/stringvalue
func (n_ Number) SetStringValue(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setStringValue:"), value)
}


// The number object’s value expressed as an unsigned
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint16value
func (n_ Number) Uint16Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("uint16Value"))
	return rv
}


// The number object’s value expressed as an unsigned
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint16value
func (n_ Number) SetUint16Value(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUint16Value:"), value)
}


// The number object’s value expressed as an unsigned
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint32value
func (n_ Number) Uint32Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("uint32Value"))
	return rv
}


// The number object’s value expressed as an unsigned
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint32value
func (n_ Number) SetUint32Value(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUint32Value:"), value)
}


// The number object’s value expressed as an unsigned
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint64value
func (n_ Number) Uint64Value() uint64 {
	rv := objc.Send[uint64](n_.ID, objc.Sel("uint64Value"))
	return rv
}


// The number object’s value expressed as an unsigned
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint64value
func (n_ Number) SetUint64Value(value uint64) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUint64Value:"), value)
}


// The number object’s value expressed as an unsigned
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint8value
func (n_ Number) Uint8Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("uint8Value"))
	return rv
}


// The number object’s value expressed as an unsigned
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uint8value
func (n_ Number) SetUint8Value(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUint8Value:"), value)
}


// The number object’s value expressed as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uintvalue
func (n_ Number) UintValue() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("uintValue"))
	return rv
}


// The number object’s value expressed as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnumber/uintvalue
func (n_ Number) SetUintValue(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUintValue:"), value)
}


// A C string containing the Objective-C type of the data contained in the value object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/objctype
func (n_ Number) ObjCType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("objCType"))
	return rv
}


// A C string containing the Objective-C type of the data contained in the value object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/objctype
func (n_ Number) SetObjCType(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setObjCType:"), value)
}




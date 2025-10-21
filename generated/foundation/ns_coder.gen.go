// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Coder] class.
var (
	CoderClass     _CoderClass
	CoderClassOnce sync.Once
)

func getCoderClass() _CoderClass {
	CoderClassOnce.Do(func() {
		CoderClass = _CoderClass{objc.GetClass("NSCoder")}
	})
	return CoderClass
}

type _CoderClass struct {
	class objc.Class
}

// An interface definition for the [Coder] class.
type ICoder interface {
	objectivec.IObject
	DecodeArrayOfObjCTypeCountAt(itemType unsafe.Pointer, count uint, array unsafe.Pointer)
	DecodeBoolForKey(key string) bool
	DecodeBytesForKeyReturnedLength(key string, lengthp unsafe.Pointer) unsafe.Pointer
	DecodeBytesWithReturnedLength(lengthp unsafe.Pointer)
	DecodeIntForKey(key string) unsafe.Pointer
	DecodeDataObject() unsafe.Pointer
	DecodeDoubleForKey(key string) unsafe.Pointer
	DecodeFloatForKey(key string) unsafe.Pointer
	DecodeInt32ForKey(key string) unsafe.Pointer
	DecodeInt64ForKey(key string) unsafe.Pointer
	DecodeIntegerForKey(key string) int
	DecodeNXObject() objc.ID
	DecodeObject() objc.ID
	DecodeObjectForKey(key string) objc.ID
	DecodeObjectOfClassForKey(aClass objc.Class, key string) objc.ID
	DecodePoint() Point
	EncodeDataObject(data unsafe.Pointer)
	EncodePoint(point Point)
	EncodeObjectForKey(object objc.ID, key string)
	EncodeIntegerForKey(value int, key string)
	EncodeInt32ForKey(value unsafe.Pointer, key string)
	EncodeBoolForKey(value bool, key string)
	EncodeFloatForKey(value unsafe.Pointer, key string)
	EncodeDoubleForKey(value unsafe.Pointer, key string)
	EncodeInt64ForKey(value unsafe.Pointer, key string)
	EncodeBytesLengthForKey(bytes unsafe.Pointer, length uint, key string)
	EncodeIntForKey(value unsafe.Pointer, key string)
	EncodeConditionalObjectForKey(object objc.ID, key string)
	EncodeNXObject(object objc.ID)
}

// An abstract class that serves as the basis for objects that enable archiving and distribution of other objects.
//
// declares the interface used by concrete subclasses to transfer objects and other values between memory and some other format. This capability provides the basis for archiving (storing objects and data on disk) and distribution (copying objects and data items between different processes or threads). The concrete subclasses provided by Foundation for these purposes are , , , , and . Concrete subclasses of are “coder classes”, and instances of these classes are “coder objects” (or simply “coders”). A coder that can only encode values is an “encoder”, and one that can only decode values is a “decoder”. operates on objects, scalars, C arrays, structures, strings, and on pointers to these types. It doesn’t handle types whose implementation varies across platforms, such as , , function pointers, and long chains of pointers. A coder stores object type information along with the data, so an object decoded from a stream of bytes is normally of the same class as the object that was originally encoded into the stream. An object can change its class when encoded, however; this is described in . The AVFoundation framework adds methods to the class to make it easier to create archives including Core Media time structures, and extract Core Media time structure from archives.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder
type Coder struct {
	objectivec.Object
}

// CoderFrom constructs a [Coder] from an unsafe.Pointer.
//
// An abstract class that serves as the basis for objects that enable archiving and distribution of other objects.
func CoderFrom(ptr unsafe.Pointer) Coder {
	return Coder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CoderClass) Alloc() Coder {
	rv := objc.Send[Coder](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CoderClass) New() Coder {
	rv := objc.Send[Coder](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Coder) Init() Coder {
	rv := objc.Send[Coder](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Coder) Autorelease() Coder {
	rv := objc.Send[Coder](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoder creates a new Coder instance.
func NewCoder() Coder {
	return getCoderClass().New()
}


// Decodes an array of items, whose Objective-C type is given by .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeArray(ofObjCType:count:at:)
func (c_ Coder) DecodeArrayOfObjCTypeCountAt(itemType unsafe.Pointer, count uint, array unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("decodeArrayOfObjCType:count:at:"), itemType, count, array)
}

// Decodes and returns a boolean value that was previously encoded with and associated with the string .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBool(forKey:)
func (c_ Coder) DecodeBoolForKey(key string) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("decodeBoolForKey:"), objc.String(key))
	return rv
}

// Decodes a buffer of data that was previously encoded with and associated with the string .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBytes(forKey:returnedLength:)
func (c_ Coder) DecodeBytesForKeyReturnedLength(key string, lengthp unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("decodeBytesForKey:returnedLength:"), objc.String(key), lengthp)
	return rv
}

// Decodes a buffer of data whose types are unspecified.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBytes(withReturnedLength:)
func (c_ Coder) DecodeBytesWithReturnedLength(lengthp unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("decodeBytesWithReturnedLength:"), lengthp)
}

// Decodes and returns an int value that was previously encoded with , , , or and associated with the string .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCInt(forKey:)
func (c_ Coder) DecodeIntForKey(key string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("decodeIntForKey:"), objc.String(key))
	return rv
}

// Decodes and returns an object that was previously encoded with . Subclasses must override this method.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeData()
func (c_ Coder) DecodeDataObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("decodeDataObject"))
	return rv
}

// Decodes and returns a double value that was previously encoded with either or and associated with the string .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeDouble(forKey:)
func (c_ Coder) DecodeDoubleForKey(key string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("decodeDoubleForKey:"), objc.String(key))
	return rv
}

// Decodes and returns a float value that was previously encoded with or and associated with the string .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeFloat(forKey:)
func (c_ Coder) DecodeFloatForKey(key string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("decodeFloatForKey:"), objc.String(key))
	return rv
}

// Decodes and returns a 32-bit integer value that was previously encoded with , , , or and associated with the string .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeInt32(forKey:)
func (c_ Coder) DecodeInt32ForKey(key string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("decodeInt32ForKey:"), objc.String(key))
	return rv
}

// Decodes and returns a 64-bit integer value that was previously encoded with , , , or and associated with the string .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeInt64(forKey:)
func (c_ Coder) DecodeInt64ForKey(key string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("decodeInt64ForKey:"), objc.String(key))
	return rv
}

// Decodes and returns an NSInteger value that was previously encoded with , , , or and associated with the string .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeInteger(forKey:)
func (c_ Coder) DecodeIntegerForKey(key string) int {
	rv := objc.Send[int](c_.ID, objc.Sel("decodeIntegerForKey:"), objc.String(key))
	return rv
}

// Decodes an object previously written with .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeNXObject
func (c_ Coder) DecodeNXObject() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeNXObject"))
	return rv
}

// Decodes and returns an object that was previously encoded with any of the methods.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObject()
func (c_ Coder) DecodeObject() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeObject"))
	return rv
}

// Decodes and returns a previously-encoded object that was previously encoded with or and associated with the string .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObject(forKey:)
func (c_ Coder) DecodeObjectForKey(key string) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeObjectForKey:"), objc.String(key))
	return rv
}

// Decodes an object for the key, restricted to the specified class.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObjectOfClass:forKey:
func (c_ Coder) DecodeObjectOfClassForKey(aClass objc.Class, key string) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeObjectOfClass:forKey:"), aClass, objc.String(key))
	return rv
}

// Decodes and returns an NSPoint structure that was previously encoded with .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodePoint()
func (c_ Coder) DecodePoint() Point {
	rv := objc.Send[Point](c_.ID, objc.Sel("decodePoint"))
	return rv
}

// Encodes a given data object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-1qd1e
func (c_ Coder) EncodeDataObject(data unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeDataObject:"), data)
}

// Encodes a point.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-75jv4
func (c_ Coder) EncodePoint(point Point) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodePoint:"), point)
}

// Encodes an object and associates it with the string key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-1mlmu
func (c_ Coder) EncodeObjectForKey(object objc.ID, key string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeObject:forKey:"), object, objc.String(key))
}

// Encodes an integer value and associates it with the string key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-2dprz
func (c_ Coder) EncodeIntegerForKey(value int, key string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeInteger:forKey:"), value, objc.String(key))
}

// Encodes a 32-bit integer value and associates it with the string key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-5sk4z
func (c_ Coder) EncodeInt32ForKey(value unsafe.Pointer, key string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeInt32:forKey:"), value, objc.String(key))
}

// Encodes a Boolean value and associates it with the string .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-7o6mu
func (c_ Coder) EncodeBoolForKey(value bool, key string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBool:forKey:"), value, objc.String(key))
}

// Encodes a floating point value and associates it with the string key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-84cez
func (c_ Coder) EncodeFloatForKey(value unsafe.Pointer, key string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeFloat:forKey:"), value, objc.String(key))
}

// Encodes a double-precision floating point value and associates it with the string key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-9xiiu
func (c_ Coder) EncodeDoubleForKey(value unsafe.Pointer, key string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeDouble:forKey:"), value, objc.String(key))
}

// Encodes a 64-bit integer value and associates it with the string key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-dixg
func (c_ Coder) EncodeInt64ForKey(value unsafe.Pointer, key string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeInt64:forKey:"), value, objc.String(key))
}

// Encodes a buffer of data, given its length and a pointer, and associates it with a string key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeBytes(_:length:forKey:)
func (c_ Coder) EncodeBytesLengthForKey(bytes unsafe.Pointer, length uint, key string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBytes:length:forKey:"), bytes, length, objc.String(key))
}

// Encodes a C integer value and associates it with the string key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeCInt(_:forKey:)
func (c_ Coder) EncodeIntForKey(value unsafe.Pointer, key string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeInt:forKey:"), value, objc.String(key))
}

// An encoding method for subclasses to override to conditionally encode an object, preserving common references to it, only if it has been unconditionally encoded.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeConditionalObject(_:forKey:)
func (c_ Coder) EncodeConditionalObjectForKey(object objc.ID, key string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeConditionalObject:forKey:"), object, objc.String(key))
}

// Encodes an old-style object onto the coder.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeNXObject:
func (c_ Coder) EncodeNXObject(object objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeNXObject:"), object)
}

// The set of coded classes allowed for secure coding.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/allowedclasses
func (c_ Coder) AllowedClasses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("allowedClasses"))
	return rv
}


// SetAllowedClasses sets the value of the allowedClasses property.
// The set of coded classes allowed for secure coding.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/allowedclasses
func (c_ Coder) SetAllowedClasses(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowedClasses:"), value)
}

// An error in the top-level encode.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/error
func (c_ Coder) Error_() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("error"))
	return rv
}


// SetError_ sets the value of the error property.
// An error in the top-level encode.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/error
func (c_ Coder) SetError_(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setError_:"), value)
}

// The system version in effect for the archive.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/systemversion
func (c_ Coder) SystemVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("systemVersion"))
	return rv
}


// SetSystemVersion sets the value of the systemVersion property.
// The system version in effect for the archive.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoder/systemversion
func (c_ Coder) SetSystemVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSystemVersion:"), value)
}

// The end of the range of error codes reserved for coder errors.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodererrormaximum-swift.var
func (c_ Coder) NSCoderErrorMaximum() int {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderErrorMaximum"))
	return rv
}


// SetNSCoderErrorMaximum sets the value of the NSCoderErrorMaximum property.
// The end of the range of error codes reserved for coder errors.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodererrormaximum-swift.var
func (c_ Coder) SetNSCoderErrorMaximum(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderErrorMaximum:"), value)
}

// The start of the range of error codes reserved for coder errors.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodererrorminimum-swift.var
func (c_ Coder) NSCoderErrorMinimum() int {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderErrorMinimum"))
	return rv
}


// SetNSCoderErrorMinimum sets the value of the NSCoderErrorMinimum property.
// The start of the range of error codes reserved for coder errors.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodererrorminimum-swift.var
func (c_ Coder) SetNSCoderErrorMinimum(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderErrorMinimum:"), value)
}

// Data wasn’t valid to encode.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoderinvalidvalueerror-swift.var
func (c_ Coder) NSCoderInvalidValueError() int {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderInvalidValueError"))
	return rv
}


// SetNSCoderInvalidValueError sets the value of the NSCoderInvalidValueError property.
// Data wasn’t valid to encode.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoderinvalidvalueerror-swift.var
func (c_ Coder) SetNSCoderInvalidValueError(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderInvalidValueError:"), value)
}

// Decoding failed due to corrupt data.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoderreadcorrupterror-swift.var
func (c_ Coder) NSCoderReadCorruptError() int {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderReadCorruptError"))
	return rv
}


// SetNSCoderReadCorruptError sets the value of the NSCoderReadCorruptError property.
// Decoding failed due to corrupt data.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoderreadcorrupterror-swift.var
func (c_ Coder) SetNSCoderReadCorruptError(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderReadCorruptError:"), value)
}

// The requested data wasn’t found.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodervaluenotfounderror-swift.var
func (c_ Coder) NSCoderValueNotFoundError() int {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderValueNotFoundError"))
	return rv
}


// SetNSCoderValueNotFoundError sets the value of the NSCoderValueNotFoundError property.
// The requested data wasn’t found.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodervaluenotfounderror-swift.var
func (c_ Coder) SetNSCoderValueNotFoundError(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderValueNotFoundError:"), value)
}

// A Boolean value that indicates whether the receiver supports keyed coding of objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/allowsKeyedCoding
func (c_ Coder) AllowsKeyedCoding() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsKeyedCoding"))
	return rv
}

// The action the coder should take when decoding fails.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodingFailurePolicy-swift.property
func (c_ Coder) DecodingFailurePolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("decodingFailurePolicy"))
	return rv
}

// Indicates whether the archiver requires all archived classes to resist object substitution attacks.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/requiresSecureCoding
func (c_ Coder) RequiresSecureCoding() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("requiresSecureCoding"))
	return rv
}




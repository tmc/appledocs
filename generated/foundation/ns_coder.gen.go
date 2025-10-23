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
	// properties:
	AllowedClasses() unsafe.Pointer
	AllowsKeyedCoding() bool /* primitive/slice/pointer. */
	DecodingFailurePolicy() DecodingFailurePolicy
	Error() IError
	RequiresSecureCoding() bool /* primitive/slice/pointer. */
	SystemVersion() unsafe.Pointer
	NSCoderErrorMaximum() int /* primitive/slice/pointer. */
	SetNSCoderErrorMaximum(value int /* primitive/slice/pointer. */)
	NSCoderErrorMinimum() int /* primitive/slice/pointer. */
	SetNSCoderErrorMinimum(value int /* primitive/slice/pointer. */)
	NSCoderInvalidValueError() int /* primitive/slice/pointer. */
	SetNSCoderInvalidValueError(value int /* primitive/slice/pointer. */)
	NSCoderReadCorruptError() int /* primitive/slice/pointer. */
	SetNSCoderReadCorruptError(value int /* primitive/slice/pointer. */)
	NSCoderValueNotFoundError() int /* primitive/slice/pointer. */
	SetNSCoderValueNotFoundError(value int /* primitive/slice/pointer. */)
	// methods:
	ContainsValueForKey(key IString) bool /* primitive/slice/pointer. */
	DecodeArrayOfObjCTypeCountAt(itemType unsafe.Pointer, count uint /* primitive/slice/pointer. */, array unsafe.Pointer)
	DecodeArrayOfObjectsOfClassForKey(cls objc.Class, key IString) IArray
	DecodeArrayOfObjectsOfClassesForKey(classes unsafe.Pointer, key IString) IArray
	DecodeBoolForKey(key IString) bool /* primitive/slice/pointer. */
	DecodeBytesForKeyMinimumLength(key IString, length uint /* primitive/slice/pointer. */) unsafe.Pointer
	DecodeBytesForKeyReturnedLength(key IString, lengthp UInteger /* not a class type */) unsafe.Pointer
	DecodeBytesWithMinimumLength(length uint /* primitive/slice/pointer. */)
	DecodeBytesWithReturnedLength(lengthp UInteger /* not a class type */)
	DecodeCGAffineTransformForKey(key IString) IAffineTransform
	DecodeCGPointForKey(key IString) objc.IObject /* cross-framework: Point */
	DecodeCGRectForKey(key IString) objc.IObject /* cross-framework: Rect */
	DecodeCGSizeForKey(key IString) objc.IObject /* cross-framework: Size */
	DecodeCGVectorForKey(key IString) objc.IObject /* cross-framework: Vector */
	DecodeIntForKey(key IString) int /* primitive/slice/pointer. */
	DecodeDataObject() IData
	DecodeDictionaryWithKeysOfClassObjectsOfClassForKey(keyCls objc.Class, objectCls objc.Class, key IString) IDictionary
	DecodeDictionaryWithKeysOfClassesObjectsOfClassesForKey(keyClasses unsafe.Pointer, objectClasses unsafe.Pointer, key IString) IDictionary
	DecodeDirectionalEdgeInsetsForKey(key IString) DirectionalEdgeInsets /* not a class type */
	DecodeDoubleForKey(key IString) float64 /* primitive/slice/pointer. */
	DecodeFloatForKey(key IString) float32 /* primitive/slice/pointer. */
	DecodeInt32ForKey(key IString) unsafe.Pointer
	DecodeInt64ForKey(key IString) unsafe.Pointer
	DecodeIntegerForKey(key IString) int /* primitive/slice/pointer. */
	DecodeObject() objc.ID
	DecodeObjectForKey(key IString) objc.ID
	DecodeObjectOfClassForKey(aClass objc.Class, key IString) objc.ID
	DecodeObjectOfClassesForKey(classes unsafe.Pointer, key IString) objc.ID
	DecodePoint() objc.IObject /* cross-framework: Point */
	DecodePointForKey(key IString) objc.IObject /* cross-framework: Point */
	DecodePropertyList() objc.ID
	DecodePropertyListForKey(key IString) objc.ID
	DecodeRect() objc.IObject /* cross-framework: Rect */
	DecodeRectForKey(key IString) objc.IObject /* cross-framework: Rect */
	DecodeSize() objc.IObject /* cross-framework: Size */
	DecodeSizeForKey(key IString) objc.IObject /* cross-framework: Size */
	DecodeCMTimeForKey(key IString) Time /* not a class type */
	DecodeCMTimeMappingForKey(key IString) TimeMapping /* not a class type */
	DecodeCMTimeRangeForKey(key IString) TimeRange /* not a class type */
	DecodeTopLevelObjectAndReturnError(error_ IError) objc.ID
	DecodeTopLevelObjectForKeyError(key IString, error_ IError) objc.ID
	DecodeTopLevelObjectOfClassForKeyError(aClass objc.Class, key IString, error_ IError) objc.ID
	DecodeTopLevelObjectOfClassesForKeyError(classes unsafe.Pointer, key IString, error_ IError) objc.ID
	DecodeUIEdgeInsetsForKey(key IString) objc.IObject /* cross-framework: EdgeInsets */
	DecodeUIOffsetForKey(key IString) Offset /* not a class type */
	DecodeValueOfObjCTypeAtSize(type_ unsafe.Pointer, data unsafe.Pointer, size uint /* primitive/slice/pointer. */)
	DecodeValuesOfObjCTypes(types unsafe.Pointer)
	EncodeDataObject(data IData)
	EncodeRect(rect objc.IObject /* cross-framework Rect */)
	EncodePoint(point objc.IObject /* cross-framework Point */)
	EncodeSize(size objc.IObject /* cross-framework Size */)
	EncodeObject(object objectivec.IObject)
	EncodeCGRectForKey(rect objc.IObject /* cross-framework Rect */, key IString)
	EncodeObjectForKey(object objectivec.IObject, key IString)
	EncodeCGVectorForKey(vector objc.IObject /* cross-framework Vector */, key IString)
	EncodePointForKey(point objc.IObject /* cross-framework Point */, key IString)
	EncodeCGAffineTransformForKey(transform IAffineTransform, key IString)
	EncodeIntegerForKey(value int /* primitive/slice/pointer. */, key IString)
	EncodeRectForKey(rect objc.IObject /* cross-framework Rect */, key IString)
	EncodeUIEdgeInsetsForKey(insets objc.IObject /* cross-framework EdgeInsets */, key IString)
	EncodeCMTimeRangeForKey(timeRange TimeRange /* not a class type */, key IString)
	EncodeInt32ForKey(value unsafe.Pointer, key IString)
	EncodeCMTimeForKey(time Time /* not a class type */, key IString)
	EncodeCGSizeForKey(size objc.IObject /* cross-framework Size */, key IString)
	EncodeBoolForKey(value bool /* primitive/slice/pointer. */, key IString)
	EncodeDirectionalEdgeInsetsForKey(insets DirectionalEdgeInsets /* not a class type */, key IString)
	EncodeCGPointForKey(point objc.IObject /* cross-framework Point */, key IString)
	EncodeFloatForKey(value float32 /* primitive/slice/pointer. */, key IString)
	EncodeCMTimeMappingForKey(timeMapping TimeMapping /* not a class type */, key IString)
	EncodeUIOffsetForKey(offset Offset /* not a class type */, key IString)
	EncodeSizeForKey(size objc.IObject /* cross-framework Size */, key IString)
	EncodeDoubleForKey(value float64 /* primitive/slice/pointer. */, key IString)
	EncodeInt64ForKey(value unsafe.Pointer, key IString)
	EncodeArrayOfObjCTypeCountAt(type_ unsafe.Pointer, count uint /* primitive/slice/pointer. */, array unsafe.Pointer)
	EncodeBycopyObject(anObject objectivec.IObject)
	EncodeByrefObject(anObject objectivec.IObject)
	EncodeBytesLength(byteaddr unsafe.Pointer, length uint /* primitive/slice/pointer. */)
	EncodeBytesLengthForKey(bytes unsafe.Pointer, length uint /* primitive/slice/pointer. */, key IString)
	EncodeIntForKey(value int /* primitive/slice/pointer. */, key IString)
	EncodeConditionalObject(object objectivec.IObject)
	EncodeConditionalObjectForKey(object objectivec.IObject, key IString)
	EncodePropertyList(aPropertyList objectivec.IObject)
	EncodeRootObject(rootObject objectivec.IObject)
	EncodeValueOfObjCTypeAt(type_ unsafe.Pointer, addr unsafe.Pointer)
	EncodeValuesOfObjCTypes(types unsafe.Pointer)
	FailWithError(error_ IError)
	ObjectZone() Zone /* not a class type */
	SetObjectZone(zone Zone /* not a class type */)
	VersionForClassName(className IString) int /* primitive/slice/pointer. */
}

// An abstract class that serves as the basis for objects that enable archiving and distribution of other objects.
//
// declares the interface used by concrete subclasses to transfer objects and other values between memory and some other format. This capability provides the basis for archiving (storing objects and data on disk) and distribution (copying objects and data items between different processes or threads). The concrete subclasses provided by Foundation for these purposes are , , , , and . Concrete subclasses of are “coder classes”, and instances of these classes are “coder objects” (or simply “coders”). A coder that can only encode values is an “encoder”, and one that can only decode values is a “decoder”. operates on objects, scalars, C arrays, structures, strings, and on pointers to these types. It doesn’t handle types whose implementation varies across platforms, such as , , function pointers, and long chains of pointers. A coder stores object type information along with the data, so an object decoded from a stream of bytes is normally of the same class as the object that was originally encoded into the stream. An object can change its class when encoded, however; this is described in . The AVFoundation framework adds methods to the class to make it easier to create archives including Core Media time structures, and extract Core Media time structure from archives.


// An abstract class that serves as the basis for objects that enable archiving and distribution of other objects.
//
// [Full Topic]
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



// Returns a Boolean value that indicates whether an encoded value is available for a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/containsValue(forKey:)
func (c_ Coder) ContainsValueForKey(key IString) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("containsValueForKey:"), key)
	return rv
}


// Decodes an array of items, whose Objective-C type is given by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeArray(ofObjCType:count:at:)
func (c_ Coder) DecodeArrayOfObjCTypeCountAt(itemType unsafe.Pointer, count uint /* primitive/slice/pointer. */, array unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("decodeArrayOfObjCType:count:at:"), itemType, count, array)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeArrayOfObjectsOfClass:forKey:
func (c_ Coder) DecodeArrayOfObjectsOfClassForKey(cls objc.Class, key IString) IArray {
	rv := objc.Send[Array](c_.ID, objc.Sel("decodeArrayOfObjectsOfClass:forKey:"), cls, key)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeArrayOfObjectsOfClasses:forKey:
func (c_ Coder) DecodeArrayOfObjectsOfClassesForKey(classes unsafe.Pointer, key IString) IArray {
	rv := objc.Send[Array](c_.ID, objc.Sel("decodeArrayOfObjectsOfClasses:forKey:"), classes, key)
	return rv
}


// Decodes and returns a boolean value that was previously encoded with and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBool(forKey:)
func (c_ Coder) DecodeBoolForKey(key IString) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("decodeBoolForKey:"), key)
	return rv
}


// Decode bytes from the decoder for a given key. The length of the bytes must be greater than or equal to the parameter. If the result exists, but is of insufficient length, then the decoder uses to fail the entire decode operation. The result of that is configurable on a per-NSCoder basis using .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBytes(forKey:minimumLength:)
func (c_ Coder) DecodeBytesForKeyMinimumLength(key IString, length uint /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("decodeBytesForKey:minimumLength:"), key, length)
	return rv
}


// Decodes a buffer of data that was previously encoded with and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBytes(forKey:returnedLength:)
func (c_ Coder) DecodeBytesForKeyReturnedLength(key IString, lengthp UInteger /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("decodeBytesForKey:returnedLength:"), key, lengthp)
	return rv
}


// Decode bytes from the decoder. The length of the bytes must be greater than or equal to the parameter. If the result exists, but is of insufficient length, then the decoder uses to fail the entire decode operation. The result of that is configurable on a per-NSCoder basis using .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBytes(withMinimumLength:)
func (c_ Coder) DecodeBytesWithMinimumLength(length uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("decodeBytesWithMinimumLength:"), length)
}


// Decodes a buffer of data whose types are unspecified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBytes(withReturnedLength:)
func (c_ Coder) DecodeBytesWithReturnedLength(lengthp UInteger /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("decodeBytesWithReturnedLength:"), lengthp)
}


// Decodes and returns the Core Graphics affine transform structure associated with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGAffineTransform(forKey:)
func (c_ Coder) DecodeCGAffineTransformForKey(key IString) IAffineTransform {
	rv := objc.Send[AffineTransform](c_.ID, objc.Sel("decodeCGAffineTransformForKey:"), key)
	return rv
}


// Decodes and returns the Core Graphics point structure associated with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGPoint(forKey:)
func (c_ Coder) DecodeCGPointForKey(key IString) objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[Point](c_.ID, objc.Sel("decodeCGPointForKey:"), key)
	return rv
}


// Decodes and returns the Core Graphics rectangle structure associated with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGRect(forKey:)
func (c_ Coder) DecodeCGRectForKey(key IString) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[Rect](c_.ID, objc.Sel("decodeCGRectForKey:"), key)
	return rv
}


// Decodes and returns the Core Graphics size structure associated with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGSize(forKey:)
func (c_ Coder) DecodeCGSizeForKey(key IString) objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("decodeCGSizeForKey:"), key)
	return rv
}


// Decodes and returns the Core Graphics vector data associated with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGVector(forKey:)
func (c_ Coder) DecodeCGVectorForKey(key IString) objc.IObject /* cross-framework: Vector */ {
	rv := objc.Send[Vector](c_.ID, objc.Sel("decodeCGVectorForKey:"), key)
	return rv
}


// Decodes and returns an int value that was previously encoded with , , , or and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCInt(forKey:)
func (c_ Coder) DecodeIntForKey(key IString) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](c_.ID, objc.Sel("decodeIntForKey:"), key)
	return rv
}


// Decodes and returns an object that was previously encoded with . Subclasses must override this method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeData()
func (c_ Coder) DecodeDataObject() IData {
	rv := objc.Send[Data](c_.ID, objc.Sel("decodeDataObject"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeDictionaryWithKeysOfClass:objectsOfClass:forKey:
func (c_ Coder) DecodeDictionaryWithKeysOfClassObjectsOfClassForKey(keyCls objc.Class, objectCls objc.Class, key IString) IDictionary {
	rv := objc.Send[Dictionary](c_.ID, objc.Sel("decodeDictionaryWithKeysOfClass:objectsOfClass:forKey:"), keyCls, objectCls, key)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeDictionaryWithKeysOfClasses:objectsOfClasses:forKey:
func (c_ Coder) DecodeDictionaryWithKeysOfClassesObjectsOfClassesForKey(keyClasses unsafe.Pointer, objectClasses unsafe.Pointer, key IString) IDictionary {
	rv := objc.Send[Dictionary](c_.ID, objc.Sel("decodeDictionaryWithKeysOfClasses:objectsOfClasses:forKey:"), keyClasses, objectClasses, key)
	return rv
}


// Decodes and returns the UIKit directional edge insets structure associated with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeDirectionalEdgeInsets(forKey:)
func (c_ Coder) DecodeDirectionalEdgeInsetsForKey(key IString) DirectionalEdgeInsets /* not a class type */ {
	rv := objc.Send[DirectionalEdgeInsets](c_.ID, objc.Sel("decodeDirectionalEdgeInsetsForKey:"), key)
	return rv
}


// Decodes and returns a double value that was previously encoded with either or and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeDouble(forKey:)
func (c_ Coder) DecodeDoubleForKey(key IString) float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("decodeDoubleForKey:"), key)
	return rv
}


// Decodes and returns a float value that was previously encoded with or and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeFloat(forKey:)
func (c_ Coder) DecodeFloatForKey(key IString) float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](c_.ID, objc.Sel("decodeFloatForKey:"), key)
	return rv
}


// Decodes and returns a 32-bit integer value that was previously encoded with , , , or and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeInt32(forKey:)
func (c_ Coder) DecodeInt32ForKey(key IString) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("decodeInt32ForKey:"), key)
	return rv
}


// Decodes and returns a 64-bit integer value that was previously encoded with , , , or and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeInt64(forKey:)
func (c_ Coder) DecodeInt64ForKey(key IString) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("decodeInt64ForKey:"), key)
	return rv
}


// Decodes and returns an NSInteger value that was previously encoded with , , , or and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeInteger(forKey:)
func (c_ Coder) DecodeIntegerForKey(key IString) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](c_.ID, objc.Sel("decodeIntegerForKey:"), key)
	return rv
}


// Decodes and returns an object that was previously encoded with any of the methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObject()
func (c_ Coder) DecodeObject() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeObject"))
	return rv
}


// Decodes and returns a previously-encoded object that was previously encoded with or and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObject(forKey:)
func (c_ Coder) DecodeObjectForKey(key IString) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeObjectForKey:"), key)
	return rv
}


// Decodes an object for the key, restricted to the specified class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObjectOfClass:forKey:
func (c_ Coder) DecodeObjectOfClassForKey(aClass objc.Class, key IString) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeObjectOfClass:forKey:"), aClass, key)
	return rv
}


// Decodes an object for the key, restricted to the specified classes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObjectOfClasses:forKey:
func (c_ Coder) DecodeObjectOfClassesForKey(classes unsafe.Pointer, key IString) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeObjectOfClasses:forKey:"), classes, key)
	return rv
}


// Decodes and returns an NSPoint structure that was previously encoded with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodePoint()
func (c_ Coder) DecodePoint() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[Point](c_.ID, objc.Sel("decodePoint"))
	return rv
}


// Decodes and returns an NSPoint structure that was previously encoded with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodePoint(forKey:)
func (c_ Coder) DecodePointForKey(key IString) objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[Point](c_.ID, objc.Sel("decodePointForKey:"), key)
	return rv
}


// Decodes a property list that was previously encoded with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodePropertyList()
func (c_ Coder) DecodePropertyList() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodePropertyList"))
	return rv
}


// Returns a decoded property list for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodePropertyList(forKey:)
func (c_ Coder) DecodePropertyListForKey(key IString) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodePropertyListForKey:"), key)
	return rv
}


// Decodes and returns an NSRect structure that was previously encoded with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeRect()
func (c_ Coder) DecodeRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[Rect](c_.ID, objc.Sel("decodeRect"))
	return rv
}


// Decodes and returns an NSRect structure that was previously encoded with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeRect(forKey:)
func (c_ Coder) DecodeRectForKey(key IString) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[Rect](c_.ID, objc.Sel("decodeRectForKey:"), key)
	return rv
}


// Decodes and returns an NSSize structure that was previously encoded with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeSize()
func (c_ Coder) DecodeSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("decodeSize"))
	return rv
}


// Decodes and returns an NSSize structure that was previously encoded with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeSize(forKey:)
func (c_ Coder) DecodeSizeForKey(key IString) objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("decodeSizeForKey:"), key)
	return rv
}


// Returns the Core Media time structure associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTime(forKey:)
func (c_ Coder) DecodeCMTimeForKey(key IString) Time /* not a class type */ {
	rv := objc.Send[Time](c_.ID, objc.Sel("decodeCMTimeForKey:"), key)
	return rv
}


// Returns the Core Media time mapping structure associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTimeMapping(forKey:)
func (c_ Coder) DecodeCMTimeMappingForKey(key IString) TimeMapping /* not a class type */ {
	rv := objc.Send[TimeMapping](c_.ID, objc.Sel("decodeCMTimeMappingForKey:"), key)
	return rv
}


// Returns the Core Media time range structure associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTimeRange(forKey:)
func (c_ Coder) DecodeCMTimeRangeForKey(key IString) TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](c_.ID, objc.Sel("decodeCMTimeRangeForKey:"), key)
	return rv
}


// Decodes a previously-encoded object, populating an error if decoding fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTopLevelObjectAndReturnError:
func (c_ Coder) DecodeTopLevelObjectAndReturnError(error_ IError) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeTopLevelObjectAndReturnError:"), error_)
	return rv
}


// Decodes the previously-encoded object associated by a key, populating an error if decoding fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTopLevelObjectForKey:error:
func (c_ Coder) DecodeTopLevelObjectForKeyError(key IString, error_ IError) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeTopLevelObjectForKey:error:"), key, error_)
	return rv
}


// Decode an object as an expected type, failing if the archived type does not match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTopLevelObjectOfClass:forKey:error:
func (c_ Coder) DecodeTopLevelObjectOfClassForKeyError(aClass objc.Class, key IString, error_ IError) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeTopLevelObjectOfClass:forKey:error:"), aClass, key, error_)
	return rv
}


// Decode an object as one of several expected types, failing if the archived type does not match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTopLevelObjectOfClasses:forKey:error:
func (c_ Coder) DecodeTopLevelObjectOfClassesForKeyError(classes unsafe.Pointer, key IString, error_ IError) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeTopLevelObjectOfClasses:forKey:error:"), classes, key, error_)
	return rv
}


// Decodes and returns the UIKit edge insets structure associated with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeUIEdgeInsets(forKey:)
func (c_ Coder) DecodeUIEdgeInsetsForKey(key IString) objc.IObject /* cross-framework: EdgeInsets */ {
	rv := objc.Send[EdgeInsets](c_.ID, objc.Sel("decodeUIEdgeInsetsForKey:"), key)
	return rv
}


// Decodes and returns the UIKit offset structure associated with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeUIOffset(forKey:)
func (c_ Coder) DecodeUIOffsetForKey(key IString) Offset /* not a class type */ {
	rv := objc.Send[Offset](c_.ID, objc.Sel("decodeUIOffsetForKey:"), key)
	return rv
}


// Decodes a single value of a known type from the specified data buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeValue(ofObjCType:at:size:)
func (c_ Coder) DecodeValueOfObjCTypeAtSize(type_ unsafe.Pointer, data unsafe.Pointer, size uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("decodeValueOfObjCType:at:size:"), type_, data, size)
}


// Decodes a series of potentially different Objective-C types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeValuesOfObjCTypes:
func (c_ Coder) DecodeValuesOfObjCTypes(types unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("decodeValuesOfObjCTypes:"), types)
}


// Encodes a given data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-1qd1e
func (c_ Coder) EncodeDataObject(data IData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeDataObject:"), data)
}


// Encodes a rectangle structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-3c1wz
func (c_ Coder) EncodeRect(rect objc.IObject /* cross-framework Rect */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeRect:"), rect)
}


// Encodes a point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-75jv4
func (c_ Coder) EncodePoint(point objc.IObject /* cross-framework Point */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodePoint:"), point)
}


// Encodes a size structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-82i7c
func (c_ Coder) EncodeSize(size objc.IObject /* cross-framework Size */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeSize:"), size)
}


// Encodes an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-9648d
func (c_ Coder) EncodeObject(object objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeObject:"), object)
}


// Encodes a rectangle and associates it with the specified key in the receiver’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-10qhm
func (c_ Coder) EncodeCGRectForKey(rect objc.IObject /* cross-framework Rect */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeCGRect:forKey:"), rect, key)
}


// Encodes an object and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-1mlmu
func (c_ Coder) EncodeObjectForKey(object objectivec.IObject, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeObject:forKey:"), object, key)
}


// Encodes vector data and associates it with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-26fxa
func (c_ Coder) EncodeCGVectorForKey(vector objc.IObject /* cross-framework Vector */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeCGVector:forKey:"), vector, key)
}


// Encodes a point and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-27lif
func (c_ Coder) EncodePointForKey(point objc.IObject /* cross-framework Point */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodePoint:forKey:"), point, key)
}


// Encodes an affine transform and associates it with the specified key in the receiver’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-29jyx
func (c_ Coder) EncodeCGAffineTransformForKey(transform IAffineTransform, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeCGAffineTransform:forKey:"), transform, key)
}


// Encodes an integer value and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-2dprz
func (c_ Coder) EncodeIntegerForKey(value int /* primitive/slice/pointer. */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeInteger:forKey:"), value, key)
}


// Encodes a rectangle structure and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-2knxx
func (c_ Coder) EncodeRectForKey(rect objc.IObject /* cross-framework Rect */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeRect:forKey:"), rect, key)
}


// Encodes edge inset data and associates it with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-44zsc
func (c_ Coder) EncodeUIEdgeInsetsForKey(insets objc.IObject /* cross-framework EdgeInsets */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeUIEdgeInsets:forKey:"), insets, key)
}


// Encodes a given Core Media time range structure and associates it with a specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-46lo8
func (c_ Coder) EncodeCMTimeRangeForKey(timeRange TimeRange /* not a class type */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeCMTimeRange:forKey:"), timeRange, key)
}


// Encodes a 32-bit integer value and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-5sk4z
func (c_ Coder) EncodeInt32ForKey(value unsafe.Pointer, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeInt32:forKey:"), value, key)
}


// Encodes a given Core Media time structure and associates it with a specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-6wbby
func (c_ Coder) EncodeCMTimeForKey(time Time /* not a class type */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeCMTime:forKey:"), time, key)
}


// Encodes size information and associates it with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-6wq3n
func (c_ Coder) EncodeCGSizeForKey(size objc.IObject /* cross-framework Size */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeCGSize:forKey:"), size, key)
}


// Encodes a Boolean value and associates it with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-7o6mu
func (c_ Coder) EncodeBoolForKey(value bool /* primitive/slice/pointer. */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBool:forKey:"), value, key)
}


// Encodes directional edge inset data and associates it with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-7oo2n
func (c_ Coder) EncodeDirectionalEdgeInsetsForKey(insets DirectionalEdgeInsets /* not a class type */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeDirectionalEdgeInsets:forKey:"), insets, key)
}


// Encodes a point and associates it with the specified key in the receiver’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-7z9kc
func (c_ Coder) EncodeCGPointForKey(point objc.IObject /* cross-framework Point */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeCGPoint:forKey:"), point, key)
}


// Encodes a floating point value and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-84cez
func (c_ Coder) EncodeFloatForKey(value float32 /* primitive/slice/pointer. */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeFloat:forKey:"), value, key)
}


// Encodes a given Core Media time mapping structure and associates it with a specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-8tefb
func (c_ Coder) EncodeCMTimeMappingForKey(timeMapping TimeMapping /* not a class type */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeCMTimeMapping:forKey:"), timeMapping, key)
}


// Encodes offset data and associates it with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-9d1qy
func (c_ Coder) EncodeUIOffsetForKey(offset Offset /* not a class type */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeUIOffset:forKey:"), offset, key)
}


// Encodes a size structure and associates it with the given string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-9imtu
func (c_ Coder) EncodeSizeForKey(size objc.IObject /* cross-framework Size */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeSize:forKey:"), size, key)
}


// Encodes a double-precision floating point value and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-9xiiu
func (c_ Coder) EncodeDoubleForKey(value float64 /* primitive/slice/pointer. */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeDouble:forKey:"), value, key)
}


// Encodes a 64-bit integer value and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-dixg
func (c_ Coder) EncodeInt64ForKey(value unsafe.Pointer, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeInt64:forKey:"), value, key)
}


// Encodes an array of the given Objective-C type, provided the number of items and a pointer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeArray(ofObjCType:count:at:)
func (c_ Coder) EncodeArrayOfObjCTypeCountAt(type_ unsafe.Pointer, count uint /* primitive/slice/pointer. */, array unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeArrayOfObjCType:count:at:"), type_, count, array)
}


// An encoding method for subclasses to override such that it creates a copy, rather than a proxy, when decoded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeBycopyObject(_:)
func (c_ Coder) EncodeBycopyObject(anObject objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBycopyObject:"), anObject)
}


// An encoding method for subclasses to override such that it creates a proxy, rather than a copy, when decoded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeByrefObject(_:)
func (c_ Coder) EncodeByrefObject(anObject objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeByrefObject:"), anObject)
}


// Encodes a buffer of data of an unspecified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeBytes(_:length:)
func (c_ Coder) EncodeBytesLength(byteaddr unsafe.Pointer, length uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBytes:length:"), byteaddr, length)
}


// Encodes a buffer of data, given its length and a pointer, and associates it with a string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeBytes(_:length:forKey:)
func (c_ Coder) EncodeBytesLengthForKey(bytes unsafe.Pointer, length uint /* primitive/slice/pointer. */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBytes:length:forKey:"), bytes, length, key)
}


// Encodes a C integer value and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeCInt(_:forKey:)
func (c_ Coder) EncodeIntForKey(value int /* primitive/slice/pointer. */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeInt:forKey:"), value, key)
}


// An encoding method for subclasses to override to conditionally encode an object, preserving common references to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeConditionalObject(_:)
func (c_ Coder) EncodeConditionalObject(object objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeConditionalObject:"), object)
}


// An encoding method for subclasses to override to conditionally encode an object, preserving common references to it, only if it has been unconditionally encoded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeConditionalObject(_:forKey:)
func (c_ Coder) EncodeConditionalObjectForKey(object objectivec.IObject, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeConditionalObject:forKey:"), object, key)
}


// Encodes a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodePropertyList(_:)
func (c_ Coder) EncodePropertyList(aPropertyList objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodePropertyList:"), aPropertyList)
}


// An encoding method for subclasses to override to encode an interconnected group of objects, starting with the provided root object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeRootObject(_:)
func (c_ Coder) EncodeRootObject(rootObject objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeRootObject:"), rootObject)
}


// Encodes a value of the given type at the given address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeValue(ofObjCType:at:)
func (c_ Coder) EncodeValueOfObjCTypeAt(type_ unsafe.Pointer, addr unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeValueOfObjCType:at:"), type_, addr)
}


// Encodes a series of values of potentially differing Objective-C types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeValuesOfObjCTypes:
func (c_ Coder) EncodeValuesOfObjCTypes(types unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeValuesOfObjCTypes:"), types)
}


// Signals to this coder that the decode operation has failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/failWithError(_:)
func (c_ Coder) FailWithError(error_ IError) {
	objc.Send[objc.ID](c_.ID, objc.Sel("failWithError:"), error_)
}


// This method is present for historical reasons and has no effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/objectZone
func (c_ Coder) ObjectZone() Zone /* not a class type */ {
	rv := objc.Send[Zone](c_.ID, objc.Sel("objectZone"))
	return rv
}


// This method is present for historical reasons and has no effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/setObjectZone:
func (c_ Coder) SetObjectZone(zone Zone /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObjectZone:"), zone)
}


// This method is present for historical reasons and is not used with keyed archivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/version(forClassName:)
func (c_ Coder) VersionForClassName(className IString) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](c_.ID, objc.Sel("versionForClassName:"), className)
	return rv
}


// The set of coded classes allowed for secure coding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/allowedClasses
func (c_ Coder) AllowedClasses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("allowedClasses"))
	return rv
}


// A Boolean value that indicates whether the receiver supports keyed coding of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/allowsKeyedCoding
func (c_ Coder) AllowsKeyedCoding() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsKeyedCoding"))
	return rv
}


// The action the coder should take when decoding fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodingFailurePolicy-swift.property
func (c_ Coder) DecodingFailurePolicy() DecodingFailurePolicy {
	rv := objc.Send[DecodingFailurePolicy](c_.ID, objc.Sel("decodingFailurePolicy"))
	return rv
}


// An error in the top-level encode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/error
func (c_ Coder) Error() IError {
	rv := objc.Send[Error](c_.ID, objc.Sel("error"))
	return rv
}


// Indicates whether the archiver requires all archived classes to resist object substitution attacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/requiresSecureCoding
func (c_ Coder) RequiresSecureCoding() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("requiresSecureCoding"))
	return rv
}


// The system version in effect for the archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/systemVersion
func (c_ Coder) SystemVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("systemVersion"))
	return rv
}


// The end of the range of error codes reserved for coder errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodererrormaximum-swift.var
func (c_ Coder) NSCoderErrorMaximum() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderErrorMaximum"))
	return rv
}


// The end of the range of error codes reserved for coder errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodererrormaximum-swift.var
func (c_ Coder) SetNSCoderErrorMaximum(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderErrorMaximum:"), value)
}


// The start of the range of error codes reserved for coder errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodererrorminimum-swift.var
func (c_ Coder) NSCoderErrorMinimum() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderErrorMinimum"))
	return rv
}


// The start of the range of error codes reserved for coder errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodererrorminimum-swift.var
func (c_ Coder) SetNSCoderErrorMinimum(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderErrorMinimum:"), value)
}


// Data wasn’t valid to encode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoderinvalidvalueerror-swift.var
func (c_ Coder) NSCoderInvalidValueError() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderInvalidValueError"))
	return rv
}


// Data wasn’t valid to encode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoderinvalidvalueerror-swift.var
func (c_ Coder) SetNSCoderInvalidValueError(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderInvalidValueError:"), value)
}


// Decoding failed due to corrupt data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoderreadcorrupterror-swift.var
func (c_ Coder) NSCoderReadCorruptError() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderReadCorruptError"))
	return rv
}


// Decoding failed due to corrupt data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoderreadcorrupterror-swift.var
func (c_ Coder) SetNSCoderReadCorruptError(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderReadCorruptError:"), value)
}


// The requested data wasn’t found.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodervaluenotfounderror-swift.var
func (c_ Coder) NSCoderValueNotFoundError() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderValueNotFoundError"))
	return rv
}


// The requested data wasn’t found.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodervaluenotfounderror-swift.var
func (c_ Coder) SetNSCoderValueNotFoundError(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderValueNotFoundError:"), value)
}




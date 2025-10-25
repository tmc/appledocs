// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCoder */


/* debug [class_header]: Header for NSCoder */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Coder */
// An interface definition for the [Coder] class.
type ICoder interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Coder */
	// properties:
	AllowedClasses() unsafe.Pointer
	AllowsKeyedCoding() bool
	DecodingFailurePolicy() DecodingFailurePolicy
	Error() IError
	RequiresSecureCoding() bool
	SystemVersion() objectivec.IObject
	NSCoderErrorMaximum() int
	SetNSCoderErrorMaximum(value int)
	NSCoderErrorMinimum() int
	SetNSCoderErrorMinimum(value int)
	NSCoderInvalidValueError() int
	SetNSCoderInvalidValueError(value int)
	NSCoderReadCorruptError() int
	SetNSCoderReadCorruptError(value int)
	NSCoderValueNotFoundError() int
	SetNSCoderValueNotFoundError(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Coder */
	// methods:
	ContainsValueForKey(key IString) bool
	DecodeArrayOfObjCTypeCountAt(itemType objectivec.IObject, count uint, array objectivec.IObject)
	DecodeArrayOfObjectsOfClassForKey(cls objc.Class, key IString) IArray
	DecodeArrayOfObjectsOfClassesForKey(classes unsafe.Pointer, key IString) IArray
	DecodeBoolForKey(key IString) bool
	DecodeBytesForKeyMinimumLength(key IString, length uint) uint8 /* not a class type */
	DecodeBytesForKeyReturnedLength(key IString, lengthp uint) uint8 /* not a class type */
	DecodeBytesWithMinimumLength(length uint)
	DecodeBytesWithReturnedLength(lengthp uint)
	DecodeIntForKey(key IString) int
	DecodeDataObject() IData
	DecodeDictionaryWithKeysOfClassObjectsOfClassForKey(keyCls objc.Class, objectCls objc.Class, key IString) IDictionary
	DecodeDictionaryWithKeysOfClassesObjectsOfClassesForKey(keyClasses unsafe.Pointer, objectClasses unsafe.Pointer, key IString) IDictionary
	DecodeDoubleForKey(key IString) float64
	DecodeFloatForKey(key IString) float32
	DecodeInt32ForKey(key IString) int32 /* not a class type */
	DecodeInt64ForKey(key IString) int64
	DecodeIntegerForKey(key IString) int
	DecodeObject() objc.ID
	DecodeObjectForKey(key IString) objc.ID
	DecodeObjectOfClassForKey(aClass objc.Class, key IString) objc.ID
	DecodeObjectOfClassesForKey(classes unsafe.Pointer, key IString) objc.ID
	DecodePoint() corefoundation.CGPoint
	DecodePointForKey(key IString) corefoundation.CGPoint
	DecodePropertyList() objc.ID
	DecodePropertyListForKey(key IString) objc.ID
	DecodeRect() corefoundation.CGRect
	DecodeRectForKey(key IString) corefoundation.CGRect
	DecodeSize() corefoundation.CGSize
	DecodeSizeForKey(key IString) corefoundation.CGSize
	DecodeCMTimeForKey(key IString) objectivec.IObject
	DecodeCMTimeMappingForKey(key IString) TimeMapping /* not a class type */
	DecodeCMTimeRangeForKey(key IString) TimeRange /* not a class type */
	DecodeTopLevelObjectAndReturnError(error_ IError) objc.ID
	DecodeTopLevelObjectForKeyError(key IString, error_ IError) objc.ID
	DecodeTopLevelObjectOfClassForKeyError(aClass objc.Class, key IString, error_ IError) objc.ID
	DecodeTopLevelObjectOfClassesForKeyError(classes unsafe.Pointer, key IString, error_ IError) objc.ID
	DecodeValueOfObjCTypeAtSize(type_ objectivec.IObject, data objectivec.IObject, size uint)
	DecodeValuesOfObjCTypes(types objectivec.IObject)
	EncodeDataObject(data IData)
	EncodeRect(rect corefoundation.CGRect)
	EncodePoint(point corefoundation.CGPoint)
	EncodeSize(size corefoundation.CGSize)
	EncodeObject(object objc.IObject)
	EncodeObjectForKey(object objc.IObject, key IString)
	EncodePointForKey(point corefoundation.CGPoint, key IString)
	EncodeIntegerForKey(value int, key IString)
	EncodeRectForKey(rect corefoundation.CGRect, key IString)
	EncodeCMTimeRangeForKey(timeRange TimeRange /* not a class type */, key IString)
	EncodeInt32ForKey(value int32 /* not a class type */, key IString)
	EncodeCMTimeForKey(time objectivec.IObject, key IString)
	EncodeBoolForKey(value bool, key IString)
	EncodeFloatForKey(value float32, key IString)
	EncodeCMTimeMappingForKey(timeMapping TimeMapping /* not a class type */, key IString)
	EncodeSizeForKey(size corefoundation.CGSize, key IString)
	EncodeDoubleForKey(value float64, key IString)
	EncodeInt64ForKey(value int64, key IString)
	EncodeArrayOfObjCTypeCountAt(type_ objectivec.IObject, count uint, array objectivec.IObject)
	EncodeBycopyObject(anObject objc.IObject)
	EncodeByrefObject(anObject objc.IObject)
	EncodeBytesLength(byteaddr objectivec.IObject, length uint)
	EncodeBytesLengthForKey(bytes objectivec.IObject, length uint, key IString)
	EncodeIntForKey(value int, key IString)
	EncodeConditionalObject(object objc.IObject)
	EncodeConditionalObjectForKey(object objc.IObject, key IString)
	EncodePropertyList(aPropertyList objc.IObject)
	EncodeRootObject(rootObject objc.IObject)
	EncodeValueOfObjCTypeAt(type_ objectivec.IObject, addr objectivec.IObject)
	EncodeValuesOfObjCTypes(types objectivec.IObject)
	FailWithError(error_ IError)
	ObjectZone() Zone /* not a class type */
	SetObjectZone(zone Zone /* not a class type */)
	VersionForClassName(className IString) int
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Coder */
// Alloc allocates a new instance without initialization.
func (cc _CoderClass) Alloc() Coder {
	rv := objc.Send[Coder](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Coder */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Coder *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Coder */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Coder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Coder */

// Returns a Boolean value that indicates whether an encoded value is available for a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/containsValue(forKey:)
func (c_ Coder) ContainsValueForKey(key IString) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("containsValueForKey:"), key)
	return rv
}/* debug [instance_methods/method]: ContainsValueForKey */


// Decodes an array of items, whose Objective-C type is given by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeArray(ofObjCType:count:at:)
func (c_ Coder) DecodeArrayOfObjCTypeCountAt(itemType objectivec.IObject, count uint, array objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("decodeArrayOfObjCType:count:at:"), itemType, count, array)
}/* debug [instance_methods/method]: DecodeArrayOfObjCTypeCountAt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeArrayOfObjectsOfClass:forKey:
func (c_ Coder) DecodeArrayOfObjectsOfClassForKey(cls objc.Class, key IString) IArray {
	rv := objc.Send[Array](c_.ID, objc.Sel("decodeArrayOfObjectsOfClass:forKey:"), cls, key)
	return rv
}/* debug [instance_methods/method]: DecodeArrayOfObjectsOfClassForKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeArrayOfObjectsOfClasses:forKey:
func (c_ Coder) DecodeArrayOfObjectsOfClassesForKey(classes unsafe.Pointer, key IString) IArray {
	rv := objc.Send[Array](c_.ID, objc.Sel("decodeArrayOfObjectsOfClasses:forKey:"), classes, key)
	return rv
}/* debug [instance_methods/method]: DecodeArrayOfObjectsOfClassesForKey */


// Decodes and returns a boolean value that was previously encoded with and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBool(forKey:)
func (c_ Coder) DecodeBoolForKey(key IString) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("decodeBoolForKey:"), key)
	return rv
}/* debug [instance_methods/method]: DecodeBoolForKey */


// Decode bytes from the decoder for a given key. The length of the bytes must be greater than or equal to the parameter. If the result exists, but is of insufficient length, then the decoder uses to fail the entire decode operation. The result of that is configurable on a per-NSCoder basis using .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBytes(forKey:minimumLength:)
func (c_ Coder) DecodeBytesForKeyMinimumLength(key IString, length uint) uint8 /* not a class type */ {
	rv := objc.Send[uint8](c_.ID, objc.Sel("decodeBytesForKey:minimumLength:"), key, length)
	return rv
}/* debug [instance_methods/method]: DecodeBytesForKeyMinimumLength */


// Decodes a buffer of data that was previously encoded with and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBytes(forKey:returnedLength:)
func (c_ Coder) DecodeBytesForKeyReturnedLength(key IString, lengthp uint) uint8 /* not a class type */ {
	rv := objc.Send[uint8](c_.ID, objc.Sel("decodeBytesForKey:returnedLength:"), key, lengthp)
	return rv
}/* debug [instance_methods/method]: DecodeBytesForKeyReturnedLength */


// Decode bytes from the decoder. The length of the bytes must be greater than or equal to the parameter. If the result exists, but is of insufficient length, then the decoder uses to fail the entire decode operation. The result of that is configurable on a per-NSCoder basis using .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBytes(withMinimumLength:)
func (c_ Coder) DecodeBytesWithMinimumLength(length uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("decodeBytesWithMinimumLength:"), length)
}/* debug [instance_methods/method]: DecodeBytesWithMinimumLength */


// Decodes a buffer of data whose types are unspecified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeBytes(withReturnedLength:)
func (c_ Coder) DecodeBytesWithReturnedLength(lengthp uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("decodeBytesWithReturnedLength:"), lengthp)
}/* debug [instance_methods/method]: DecodeBytesWithReturnedLength */


// Decodes and returns an int value that was previously encoded with , , , or and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCInt(forKey:)
func (c_ Coder) DecodeIntForKey(key IString) int {
	rv := objc.Send[int](c_.ID, objc.Sel("decodeIntForKey:"), key)
	return rv
}/* debug [instance_methods/method]: DecodeIntForKey */


// Decodes and returns an object that was previously encoded with . Subclasses must override this method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeData()
func (c_ Coder) DecodeDataObject() IData {
	rv := objc.Send[Data](c_.ID, objc.Sel("decodeDataObject"))
	return rv
}/* debug [instance_methods/method]: DecodeDataObject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeDictionaryWithKeysOfClass:objectsOfClass:forKey:
func (c_ Coder) DecodeDictionaryWithKeysOfClassObjectsOfClassForKey(keyCls objc.Class, objectCls objc.Class, key IString) IDictionary {
	rv := objc.Send[Dictionary](c_.ID, objc.Sel("decodeDictionaryWithKeysOfClass:objectsOfClass:forKey:"), keyCls, objectCls, key)
	return rv
}/* debug [instance_methods/method]: DecodeDictionaryWithKeysOfClassObjectsOfClassForKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeDictionaryWithKeysOfClasses:objectsOfClasses:forKey:
func (c_ Coder) DecodeDictionaryWithKeysOfClassesObjectsOfClassesForKey(keyClasses unsafe.Pointer, objectClasses unsafe.Pointer, key IString) IDictionary {
	rv := objc.Send[Dictionary](c_.ID, objc.Sel("decodeDictionaryWithKeysOfClasses:objectsOfClasses:forKey:"), keyClasses, objectClasses, key)
	return rv
}/* debug [instance_methods/method]: DecodeDictionaryWithKeysOfClassesObjectsOfClassesForKey */


// Decodes and returns a double value that was previously encoded with either or and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeDouble(forKey:)
func (c_ Coder) DecodeDoubleForKey(key IString) float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("decodeDoubleForKey:"), key)
	return rv
}/* debug [instance_methods/method]: DecodeDoubleForKey */


// Decodes and returns a float value that was previously encoded with or and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeFloat(forKey:)
func (c_ Coder) DecodeFloatForKey(key IString) float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("decodeFloatForKey:"), key)
	return rv
}/* debug [instance_methods/method]: DecodeFloatForKey */


// Decodes and returns a 32-bit integer value that was previously encoded with , , , or and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeInt32(forKey:)
func (c_ Coder) DecodeInt32ForKey(key IString) int32 /* not a class type */ {
	rv := objc.Send[int32](c_.ID, objc.Sel("decodeInt32ForKey:"), key)
	return rv
}/* debug [instance_methods/method]: DecodeInt32ForKey */


// Decodes and returns a 64-bit integer value that was previously encoded with , , , or and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeInt64(forKey:)
func (c_ Coder) DecodeInt64ForKey(key IString) int64 {
	rv := objc.Send[int64](c_.ID, objc.Sel("decodeInt64ForKey:"), key)
	return rv
}/* debug [instance_methods/method]: DecodeInt64ForKey */


// Decodes and returns an NSInteger value that was previously encoded with , , , or and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeInteger(forKey:)
func (c_ Coder) DecodeIntegerForKey(key IString) int {
	rv := objc.Send[int](c_.ID, objc.Sel("decodeIntegerForKey:"), key)
	return rv
}/* debug [instance_methods/method]: DecodeIntegerForKey */


// Decodes and returns an object that was previously encoded with any of the methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObject()
func (c_ Coder) DecodeObject() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeObject"))
	return rv
}/* debug [instance_methods/method]: DecodeObject */


// Decodes and returns a previously-encoded object that was previously encoded with or and associated with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObject(forKey:)
func (c_ Coder) DecodeObjectForKey(key IString) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeObjectForKey:"), key)
	return rv
}/* debug [instance_methods/method]: DecodeObjectForKey */


// Decodes an object for the key, restricted to the specified class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObjectOfClass:forKey:
func (c_ Coder) DecodeObjectOfClassForKey(aClass objc.Class, key IString) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeObjectOfClass:forKey:"), aClass, key)
	return rv
}/* debug [instance_methods/method]: DecodeObjectOfClassForKey */


// Decodes an object for the key, restricted to the specified classes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObjectOfClasses:forKey:
func (c_ Coder) DecodeObjectOfClassesForKey(classes unsafe.Pointer, key IString) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeObjectOfClasses:forKey:"), classes, key)
	return rv
}/* debug [instance_methods/method]: DecodeObjectOfClassesForKey */


// Decodes and returns an NSPoint structure that was previously encoded with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodePoint()
func (c_ Coder) DecodePoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c_.ID, objc.Sel("decodePoint"))
	return rv
}/* debug [instance_methods/method]: DecodePoint */


// Decodes and returns an NSPoint structure that was previously encoded with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodePoint(forKey:)
func (c_ Coder) DecodePointForKey(key IString) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c_.ID, objc.Sel("decodePointForKey:"), key)
	return rv
}/* debug [instance_methods/method]: DecodePointForKey */


// Decodes a property list that was previously encoded with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodePropertyList()
func (c_ Coder) DecodePropertyList() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodePropertyList"))
	return rv
}/* debug [instance_methods/method]: DecodePropertyList */


// Returns a decoded property list for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodePropertyList(forKey:)
func (c_ Coder) DecodePropertyListForKey(key IString) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodePropertyListForKey:"), key)
	return rv
}/* debug [instance_methods/method]: DecodePropertyListForKey */


// Decodes and returns an NSRect structure that was previously encoded with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeRect()
func (c_ Coder) DecodeRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("decodeRect"))
	return rv
}/* debug [instance_methods/method]: DecodeRect */


// Decodes and returns an NSRect structure that was previously encoded with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeRect(forKey:)
func (c_ Coder) DecodeRectForKey(key IString) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("decodeRectForKey:"), key)
	return rv
}/* debug [instance_methods/method]: DecodeRectForKey */


// Decodes and returns an NSSize structure that was previously encoded with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeSize()
func (c_ Coder) DecodeSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("decodeSize"))
	return rv
}/* debug [instance_methods/method]: DecodeSize */


// Decodes and returns an NSSize structure that was previously encoded with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeSize(forKey:)
func (c_ Coder) DecodeSizeForKey(key IString) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("decodeSizeForKey:"), key)
	return rv
}/* debug [instance_methods/method]: DecodeSizeForKey */


// Returns the Core Media time structure associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTime(forKey:)
func (c_ Coder) DecodeCMTimeForKey(key IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("decodeCMTimeForKey:"), key)
	return rv
}/* debug [instance_methods/method]: DecodeCMTimeForKey */


// Returns the Core Media time mapping structure associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTimeMapping(forKey:)
func (c_ Coder) DecodeCMTimeMappingForKey(key IString) TimeMapping /* not a class type */ {
	rv := objc.Send[TimeMapping](c_.ID, objc.Sel("decodeCMTimeMappingForKey:"), key)
	return rv
}/* debug [instance_methods/method]: DecodeCMTimeMappingForKey */


// Returns the Core Media time range structure associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTimeRange(forKey:)
func (c_ Coder) DecodeCMTimeRangeForKey(key IString) TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](c_.ID, objc.Sel("decodeCMTimeRangeForKey:"), key)
	return rv
}/* debug [instance_methods/method]: DecodeCMTimeRangeForKey */


// Decodes a previously-encoded object, populating an error if decoding fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTopLevelObjectAndReturnError:
func (c_ Coder) DecodeTopLevelObjectAndReturnError(error_ IError) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeTopLevelObjectAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: DecodeTopLevelObjectAndReturnError */


// Decodes the previously-encoded object associated by a key, populating an error if decoding fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTopLevelObjectForKey:error:
func (c_ Coder) DecodeTopLevelObjectForKeyError(key IString, error_ IError) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeTopLevelObjectForKey:error:"), key, error_)
	return rv
}/* debug [instance_methods/method]: DecodeTopLevelObjectForKeyError */


// Decode an object as an expected type, failing if the archived type does not match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTopLevelObjectOfClass:forKey:error:
func (c_ Coder) DecodeTopLevelObjectOfClassForKeyError(aClass objc.Class, key IString, error_ IError) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeTopLevelObjectOfClass:forKey:error:"), aClass, key, error_)
	return rv
}/* debug [instance_methods/method]: DecodeTopLevelObjectOfClassForKeyError */


// Decode an object as one of several expected types, failing if the archived type does not match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeTopLevelObjectOfClasses:forKey:error:
func (c_ Coder) DecodeTopLevelObjectOfClassesForKeyError(classes unsafe.Pointer, key IString, error_ IError) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeTopLevelObjectOfClasses:forKey:error:"), classes, key, error_)
	return rv
}/* debug [instance_methods/method]: DecodeTopLevelObjectOfClassesForKeyError */


// Decodes a single value of a known type from the specified data buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeValue(ofObjCType:at:size:)
func (c_ Coder) DecodeValueOfObjCTypeAtSize(type_ objectivec.IObject, data objectivec.IObject, size uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("decodeValueOfObjCType:at:size:"), type_, data, size)
}/* debug [instance_methods/method]: DecodeValueOfObjCTypeAtSize */


// Decodes a series of potentially different Objective-C types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeValuesOfObjCTypes:
func (c_ Coder) DecodeValuesOfObjCTypes(types objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("decodeValuesOfObjCTypes:"), types)
}/* debug [instance_methods/method]: DecodeValuesOfObjCTypes */


// Encodes a given data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-1qd1e
func (c_ Coder) EncodeDataObject(data IData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeDataObject:"), data)
}/* debug [instance_methods/method]: EncodeDataObject */


// Encodes a rectangle structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-3c1wz
func (c_ Coder) EncodeRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeRect:"), rect)
}/* debug [instance_methods/method]: EncodeRect */


// Encodes a point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-75jv4
func (c_ Coder) EncodePoint(point corefoundation.CGPoint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodePoint:"), point)
}/* debug [instance_methods/method]: EncodePoint */


// Encodes a size structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-82i7c
func (c_ Coder) EncodeSize(size corefoundation.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeSize:"), size)
}/* debug [instance_methods/method]: EncodeSize */


// Encodes an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:)-9648d
func (c_ Coder) EncodeObject(object objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeObject:"), object)
}/* debug [instance_methods/method]: EncodeObject */


// Encodes an object and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-1mlmu
func (c_ Coder) EncodeObjectForKey(object objc.IObject, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeObject:forKey:"), object, key)
}/* debug [instance_methods/method]: EncodeObjectForKey */


// Encodes a point and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-27lif
func (c_ Coder) EncodePointForKey(point corefoundation.CGPoint, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodePoint:forKey:"), point, key)
}/* debug [instance_methods/method]: EncodePointForKey */


// Encodes an integer value and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-2dprz
func (c_ Coder) EncodeIntegerForKey(value int, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeInteger:forKey:"), value, key)
}/* debug [instance_methods/method]: EncodeIntegerForKey */


// Encodes a rectangle structure and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-2knxx
func (c_ Coder) EncodeRectForKey(rect corefoundation.CGRect, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeRect:forKey:"), rect, key)
}/* debug [instance_methods/method]: EncodeRectForKey */


// Encodes a given Core Media time range structure and associates it with a specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-46lo8
func (c_ Coder) EncodeCMTimeRangeForKey(timeRange TimeRange /* not a class type */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeCMTimeRange:forKey:"), timeRange, key)
}/* debug [instance_methods/method]: EncodeCMTimeRangeForKey */


// Encodes a 32-bit integer value and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-5sk4z
func (c_ Coder) EncodeInt32ForKey(value int32 /* not a class type */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeInt32:forKey:"), value, key)
}/* debug [instance_methods/method]: EncodeInt32ForKey */


// Encodes a given Core Media time structure and associates it with a specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-6wbby
func (c_ Coder) EncodeCMTimeForKey(time objectivec.IObject, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeCMTime:forKey:"), time, key)
}/* debug [instance_methods/method]: EncodeCMTimeForKey */


// Encodes a Boolean value and associates it with the string .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-7o6mu
func (c_ Coder) EncodeBoolForKey(value bool, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBool:forKey:"), value, key)
}/* debug [instance_methods/method]: EncodeBoolForKey */


// Encodes a floating point value and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-84cez
func (c_ Coder) EncodeFloatForKey(value float32, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeFloat:forKey:"), value, key)
}/* debug [instance_methods/method]: EncodeFloatForKey */


// Encodes a given Core Media time mapping structure and associates it with a specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-8tefb
func (c_ Coder) EncodeCMTimeMappingForKey(timeMapping TimeMapping /* not a class type */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeCMTimeMapping:forKey:"), timeMapping, key)
}/* debug [instance_methods/method]: EncodeCMTimeMappingForKey */


// Encodes a size structure and associates it with the given string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-9imtu
func (c_ Coder) EncodeSizeForKey(size corefoundation.CGSize, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeSize:forKey:"), size, key)
}/* debug [instance_methods/method]: EncodeSizeForKey */


// Encodes a double-precision floating point value and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-9xiiu
func (c_ Coder) EncodeDoubleForKey(value float64, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeDouble:forKey:"), value, key)
}/* debug [instance_methods/method]: EncodeDoubleForKey */


// Encodes a 64-bit integer value and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-dixg
func (c_ Coder) EncodeInt64ForKey(value int64, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeInt64:forKey:"), value, key)
}/* debug [instance_methods/method]: EncodeInt64ForKey */


// Encodes an array of the given Objective-C type, provided the number of items and a pointer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeArray(ofObjCType:count:at:)
func (c_ Coder) EncodeArrayOfObjCTypeCountAt(type_ objectivec.IObject, count uint, array objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeArrayOfObjCType:count:at:"), type_, count, array)
}/* debug [instance_methods/method]: EncodeArrayOfObjCTypeCountAt */


// An encoding method for subclasses to override such that it creates a copy, rather than a proxy, when decoded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeBycopyObject(_:)
func (c_ Coder) EncodeBycopyObject(anObject objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBycopyObject:"), anObject)
}/* debug [instance_methods/method]: EncodeBycopyObject */


// An encoding method for subclasses to override such that it creates a proxy, rather than a copy, when decoded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeByrefObject(_:)
func (c_ Coder) EncodeByrefObject(anObject objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeByrefObject:"), anObject)
}/* debug [instance_methods/method]: EncodeByrefObject */


// Encodes a buffer of data of an unspecified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeBytes(_:length:)
func (c_ Coder) EncodeBytesLength(byteaddr objectivec.IObject, length uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBytes:length:"), byteaddr, length)
}/* debug [instance_methods/method]: EncodeBytesLength */


// Encodes a buffer of data, given its length and a pointer, and associates it with a string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeBytes(_:length:forKey:)
func (c_ Coder) EncodeBytesLengthForKey(bytes objectivec.IObject, length uint, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBytes:length:forKey:"), bytes, length, key)
}/* debug [instance_methods/method]: EncodeBytesLengthForKey */


// Encodes a C integer value and associates it with the string key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeCInt(_:forKey:)
func (c_ Coder) EncodeIntForKey(value int, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeInt:forKey:"), value, key)
}/* debug [instance_methods/method]: EncodeIntForKey */


// An encoding method for subclasses to override to conditionally encode an object, preserving common references to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeConditionalObject(_:)
func (c_ Coder) EncodeConditionalObject(object objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeConditionalObject:"), object)
}/* debug [instance_methods/method]: EncodeConditionalObject */


// An encoding method for subclasses to override to conditionally encode an object, preserving common references to it, only if it has been unconditionally encoded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeConditionalObject(_:forKey:)
func (c_ Coder) EncodeConditionalObjectForKey(object objc.IObject, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeConditionalObject:forKey:"), object, key)
}/* debug [instance_methods/method]: EncodeConditionalObjectForKey */


// Encodes a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodePropertyList(_:)
func (c_ Coder) EncodePropertyList(aPropertyList objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodePropertyList:"), aPropertyList)
}/* debug [instance_methods/method]: EncodePropertyList */


// An encoding method for subclasses to override to encode an interconnected group of objects, starting with the provided root object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeRootObject(_:)
func (c_ Coder) EncodeRootObject(rootObject objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeRootObject:"), rootObject)
}/* debug [instance_methods/method]: EncodeRootObject */


// Encodes a value of the given type at the given address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeValue(ofObjCType:at:)
func (c_ Coder) EncodeValueOfObjCTypeAt(type_ objectivec.IObject, addr objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeValueOfObjCType:at:"), type_, addr)
}/* debug [instance_methods/method]: EncodeValueOfObjCTypeAt */


// Encodes a series of values of potentially differing Objective-C types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encodeValuesOfObjCTypes:
func (c_ Coder) EncodeValuesOfObjCTypes(types objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeValuesOfObjCTypes:"), types)
}/* debug [instance_methods/method]: EncodeValuesOfObjCTypes */


// Signals to this coder that the decode operation has failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/failWithError(_:)
func (c_ Coder) FailWithError(error_ IError) {
	objc.Send[objc.ID](c_.ID, objc.Sel("failWithError:"), error_)
}/* debug [instance_methods/method]: FailWithError */


// This method is present for historical reasons and has no effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/objectZone
func (c_ Coder) ObjectZone() Zone /* not a class type */ {
	rv := objc.Send[Zone](c_.ID, objc.Sel("objectZone"))
	return rv
}/* debug [instance_methods/method]: ObjectZone */


// This method is present for historical reasons and has no effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/setObjectZone:
func (c_ Coder) SetObjectZone(zone Zone /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObjectZone:"), zone)
}/* debug [instance_methods/method]: SetObjectZone */


// This method is present for historical reasons and is not used with keyed archivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/version(forClassName:)
func (c_ Coder) VersionForClassName(className IString) int {
	rv := objc.Send[int](c_.ID, objc.Sel("versionForClassName:"), className)
	return rv
}/* debug [instance_methods/method]: VersionForClassName */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Coder */

// The set of coded classes allowed for secure coding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/allowedClasses
func (c_ Coder) AllowedClasses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("allowedClasses"))
	return rv
}/* debug [instance_properties/getter]: allowedClasses */


// A Boolean value that indicates whether the receiver supports keyed coding of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/allowsKeyedCoding
func (c_ Coder) AllowsKeyedCoding() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsKeyedCoding"))
	return rv
}/* debug [instance_properties/getter]: allowsKeyedCoding */


// The action the coder should take when decoding fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodingFailurePolicy-swift.property
func (c_ Coder) DecodingFailurePolicy() DecodingFailurePolicy {
	rv := objc.Send[DecodingFailurePolicy](c_.ID, objc.Sel("decodingFailurePolicy"))
	return rv
}/* debug [instance_properties/getter]: decodingFailurePolicy */


// An error in the top-level encode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/error
func (c_ Coder) Error() IError {
	rv := objc.Send[Error](c_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// Indicates whether the archiver requires all archived classes to resist object substitution attacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/requiresSecureCoding
func (c_ Coder) RequiresSecureCoding() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("requiresSecureCoding"))
	return rv
}/* debug [instance_properties/getter]: requiresSecureCoding */


// The system version in effect for the archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/systemVersion
func (c_ Coder) SystemVersion() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("systemVersion"))
	return rv
}/* debug [instance_properties/getter]: systemVersion */


// The end of the range of error codes reserved for coder errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodererrormaximum-swift.var
func (c_ Coder) NSCoderErrorMaximum() int {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderErrorMaximum"))
	return rv
}/* debug [instance_properties/getter]: NSCoderErrorMaximum */


// The end of the range of error codes reserved for coder errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodererrormaximum-swift.var
func (c_ Coder) SetNSCoderErrorMaximum(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderErrorMaximum:"), value)
}/* debug [instance_properties/setter]: NSCoderErrorMaximum */


// The start of the range of error codes reserved for coder errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodererrorminimum-swift.var
func (c_ Coder) NSCoderErrorMinimum() int {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderErrorMinimum"))
	return rv
}/* debug [instance_properties/getter]: NSCoderErrorMinimum */


// The start of the range of error codes reserved for coder errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodererrorminimum-swift.var
func (c_ Coder) SetNSCoderErrorMinimum(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderErrorMinimum:"), value)
}/* debug [instance_properties/setter]: NSCoderErrorMinimum */


// Data wasn’t valid to encode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoderinvalidvalueerror-swift.var
func (c_ Coder) NSCoderInvalidValueError() int {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderInvalidValueError"))
	return rv
}/* debug [instance_properties/getter]: NSCoderInvalidValueError */


// Data wasn’t valid to encode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoderinvalidvalueerror-swift.var
func (c_ Coder) SetNSCoderInvalidValueError(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderInvalidValueError:"), value)
}/* debug [instance_properties/setter]: NSCoderInvalidValueError */


// Decoding failed due to corrupt data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoderreadcorrupterror-swift.var
func (c_ Coder) NSCoderReadCorruptError() int {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderReadCorruptError"))
	return rv
}/* debug [instance_properties/getter]: NSCoderReadCorruptError */


// Decoding failed due to corrupt data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscoderreadcorrupterror-swift.var
func (c_ Coder) SetNSCoderReadCorruptError(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderReadCorruptError:"), value)
}/* debug [instance_properties/setter]: NSCoderReadCorruptError */


// The requested data wasn’t found.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodervaluenotfounderror-swift.var
func (c_ Coder) NSCoderValueNotFoundError() int {
	rv := objc.Send[int](c_.ID, objc.Sel("NSCoderValueNotFoundError"))
	return rv
}/* debug [instance_properties/getter]: NSCoderValueNotFoundError */


// The requested data wasn’t found.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscodervaluenotfounderror-swift.var
func (c_ Coder) SetNSCoderValueNotFoundError(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNSCoderValueNotFoundError:"), value)
}/* debug [instance_properties/setter]: NSCoderValueNotFoundError */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCoder */



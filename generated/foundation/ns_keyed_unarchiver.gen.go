// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [KeyedUnarchiver] class.
var (
	KeyedUnarchiverClass     _KeyedUnarchiverClass
	KeyedUnarchiverClassOnce sync.Once
)

func getKeyedUnarchiverClass() _KeyedUnarchiverClass {
	KeyedUnarchiverClassOnce.Do(func() {
		KeyedUnarchiverClass = _KeyedUnarchiverClass{objc.GetClass("NSKeyedUnarchiver")}
	})
	return KeyedUnarchiverClass
}

type _KeyedUnarchiverClass struct {
	class objc.Class
}

// An interface definition for the [KeyedUnarchiver] class.
type IKeyedUnarchiver interface {
	ICoder
	// properties:
	DecodingFailurePolicy() DecodingFailurePolicy
	SetDecodingFailurePolicy(value DecodingFailurePolicy)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	RequiresSecureCoding() bool
	SetRequiresSecureCoding(value bool)
	// methods:
	ClassForClassName(codedName IString) objc.Class
	ContainsValueForKey(key IString) bool
	DecodeBoolForKey(key IString) bool
	DecodeBytesForKeyReturnedLength(key IString, lengthp uint) uint8 /* not a class type */
	DecodeDoubleForKey(key IString) float64
	DecodeFloatForKey(key IString) float32
	DecodeInt32ForKey(key IString) int32 /* not a class type */
	DecodeInt64ForKey(key IString) int64
	DecodeIntForKey(key IString) int
	DecodeObjectForKey(key IString) objc.ID
	FinishDecoding()
	SetClassForClassName(cls objc.Class, codedName IString)
}

// A decoder that restores data from an archive referenced by keys.
//
// is a concrete subclass of that defines methods for decoding a set of named objects (and scalar values) from a keyed archive. The class produces archives that this class can decode. The archiver creates keyed archive as a hierarchy of objects. The archiver treats each object as a namespace into which it can encode other objects. This means that an unarchiver can only decode objects encoded within the immediate scope of their parent object. Objects encoded elsewhere in the hierarchy — whether higher than, lower than, or parallel to this particular object — aren’t accessible. In this way, the keys used by a particular object to encode its instance variables need to be unique only within the scope of that object. If you invoke one of the -prefixed methods of this class using a key that does not exist in the archive, the return value indicates failure. This value varies by decoded type. For example, if a key does not exist in an archive, returns , returns , and returns . supports limited type coercion for numeric types. You can use any of the integer decode methods to decode a value encoded as any type of integer, whether a standard or an explicit 32-bit or 64-bit integer. Likewise, you can use the - or -returning decode methods to handle value encoded as a or . If an encoded value is too large to fit within the coerced type, the decoding method throws a . Further, when trying to coerce a value to an incompatible type — for example decoding an as a — the decoding method throws an .


// A decoder that restores data from an archive referenced by keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver
type KeyedUnarchiver struct {
	Coder
}

// KeyedUnarchiverFrom constructs a [KeyedUnarchiver] from an unsafe.Pointer.
//
// A decoder that restores data from an archive referenced by keys.
func KeyedUnarchiverFrom(ptr unsafe.Pointer) KeyedUnarchiver {
	return KeyedUnarchiver{
		Coder: CoderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (kc _KeyedUnarchiverClass) Alloc() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (kc _KeyedUnarchiverClass) New() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](objc.ID(kc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (k_ KeyedUnarchiver) Init() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](k_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k_ KeyedUnarchiver) Autorelease() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](k_.ID, objc.Sel("autorelease"))
	return rv
}

// NewKeyedUnarchiver creates a new KeyedUnarchiver instance.
func NewKeyedUnarchiver() KeyedUnarchiver {
	return getKeyedUnarchiverClass().New()
}



// Initializes an archiver to decode data from the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/init(forReadingFrom:)
func NewKeyedUnarchiverForReadingFromDataError(data IData, error_ IError) KeyedUnarchiver {
	instance := getKeyedUnarchiverClass().Alloc()
	rv := objc.Send[KeyedUnarchiver](instance.ID, objc.Sel("initForReadingFromData:error:"), data, error_)
	rv.Autorelease()
	return rv
}


// Initializes an archiver to decode data from the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/init(forReadingWith:)
func NewKeyedUnarchiverForReadingWithData(data IData) KeyedUnarchiver {
	instance := getKeyedUnarchiverClass().Alloc()
	rv := objc.Send[KeyedUnarchiver](instance.ID, objc.Sel("initForReadingWithData:"), data)
	rv.Autorelease()
	return rv
}



// Returns the class from which this unarchiver instantiates an encoded object with a given class name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/class(forClassName:)-swift.type.method
func (kc _KeyedUnarchiverClass) ClassForClassName(codedName IString) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(kc.class), objc.Sel("classForClassName:"), codedName)
	return rv
}


// Sets a global translation mapping to decode objects encoded with a given class name as instances of a given class instead.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/setClass(_:forClassName:)-swift.type.method
func (kc _KeyedUnarchiverClass) SetClassForClassName(cls objc.Class, codedName IString) {
	objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("setClass:forClassName:"), cls, codedName)
}


// Decodes and returns the object graph previously encoded by and stored in a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/unarchiveObject(with:)
func (kc _KeyedUnarchiverClass) UnarchiveObjectWithData(data IData) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchiveObjectWithData:"), data)
	return rv
}


// Decodes and returns the object graph previously encoded by written to the file at a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/unarchiveObject(withFile:)
func (kc _KeyedUnarchiverClass) UnarchiveObjectWithFile(path IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchiveObjectWithFile:"), path)
	return rv
}


// Decodes a previously-archived object graph, returning the root object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/unarchiveTopLevelObjectWithData:error:
func (kc _KeyedUnarchiverClass) UnarchiveTopLevelObjectWithDataError(data IData, error_ IError) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchiveTopLevelObjectWithData:error:"), data, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/unarchivedArrayOfObjectsOfClass:fromData:error:
func (kc _KeyedUnarchiverClass) UnarchivedArrayOfObjectsOfClassFromDataError(cls objc.Class, data IData, error_ IError) IArray {
	rv := objc.Send[Array](objc.ID(kc.class), objc.Sel("unarchivedArrayOfObjectsOfClass:fromData:error:"), cls, data, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/unarchivedArrayOfObjectsOfClasses:fromData:error:
func (kc _KeyedUnarchiverClass) UnarchivedArrayOfObjectsOfClassesFromDataError(classes unsafe.Pointer, data IData, error_ IError) IArray {
	rv := objc.Send[Array](objc.ID(kc.class), objc.Sel("unarchivedArrayOfObjectsOfClasses:fromData:error:"), classes, data, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/unarchivedDictionaryWithKeysOfClass:objectsOfClass:fromData:error:
func (kc _KeyedUnarchiverClass) UnarchivedDictionaryWithKeysOfClassObjectsOfClassFromDataError(keyCls objc.Class, valueCls objc.Class, data IData, error_ IError) IDictionary {
	rv := objc.Send[Dictionary](objc.ID(kc.class), objc.Sel("unarchivedDictionaryWithKeysOfClass:objectsOfClass:fromData:error:"), keyCls, valueCls, data, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/unarchivedDictionaryWithKeysOfClasses:objectsOfClasses:fromData:error:
func (kc _KeyedUnarchiverClass) UnarchivedDictionaryWithKeysOfClassesObjectsOfClassesFromDataError(keyClasses unsafe.Pointer, valueClasses unsafe.Pointer, data IData, error_ IError) IDictionary {
	rv := objc.Send[Dictionary](objc.ID(kc.class), objc.Sel("unarchivedDictionaryWithKeysOfClasses:objectsOfClasses:fromData:error:"), keyClasses, valueClasses, data, error_)
	return rv
}


// Decodes a previously-archived object graph, returning the root object as one of the specified classes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/unarchivedObject(ofClasses:from:)-b9t5
func (kc _KeyedUnarchiverClass) UnarchivedObjectOfClassesFromDataError(classes unsafe.Pointer, data IData, error_ IError) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchivedObjectOfClasses:fromData:error:"), classes, data, error_)
	return rv
}


// Decodes a previously-archived object graph, that returns the root object as the specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/unarchivedObjectOfClass:fromData:error:
func (kc _KeyedUnarchiverClass) UnarchivedObjectOfClassFromDataError(cls objc.Class, data IData, error_ IError) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchivedObjectOfClass:fromData:error:"), cls, data, error_)
	return rv
}


// Returns the class from which this unarchiver instantiates an encoded object with a given class name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/class(forClassName:)-swift.method
func (k_ KeyedUnarchiver) ClassForClassName(codedName IString) objc.Class {
	rv := objc.Send[objc.Class](k_.ID, objc.Sel("classForClassName:"), codedName)
	return rv
}


// Returns a Boolean value that indicates whether the archive contains a value for a given key within the current decoding scope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/containsValue(forKey:)
func (k_ KeyedUnarchiver) ContainsValueForKey(key IString) bool {
	rv := objc.Send[bool](k_.ID, objc.Sel("containsValueForKey:"), key)
	return rv
}


// Decodes a Boolean value associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/decodeBool(forKey:)
func (k_ KeyedUnarchiver) DecodeBoolForKey(key IString) bool {
	rv := objc.Send[bool](k_.ID, objc.Sel("decodeBoolForKey:"), key)
	return rv
}


// Decodes a stream of bytes associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/decodeBytes(forKey:returnedLength:)
func (k_ KeyedUnarchiver) DecodeBytesForKeyReturnedLength(key IString, lengthp uint) uint8 /* not a class type */ {
	rv := objc.Send[uint8](k_.ID, objc.Sel("decodeBytesForKey:returnedLength:"), key, lengthp)
	return rv
}


// Decodes a double-precision floating-point value associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/decodeDouble(forKey:)
func (k_ KeyedUnarchiver) DecodeDoubleForKey(key IString) float64 {
	rv := objc.Send[float64](k_.ID, objc.Sel("decodeDoubleForKey:"), key)
	return rv
}


// Decodes a single-precision floating-point value associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/decodeFloat(forKey:)
func (k_ KeyedUnarchiver) DecodeFloatForKey(key IString) float32 {
	rv := objc.Send[float32](k_.ID, objc.Sel("decodeFloatForKey:"), key)
	return rv
}


// Decodes a 32-bit integer value associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/decodeInt32(forKey:)
func (k_ KeyedUnarchiver) DecodeInt32ForKey(key IString) int32 /* not a class type */ {
	rv := objc.Send[int32](k_.ID, objc.Sel("decodeInt32ForKey:"), key)
	return rv
}


// Decodes a 64-bit integer value associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/decodeInt64(forKey:)
func (k_ KeyedUnarchiver) DecodeInt64ForKey(key IString) int64 {
	rv := objc.Send[int64](k_.ID, objc.Sel("decodeInt64ForKey:"), key)
	return rv
}


// Decodes an integer value associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/decodeIntForKey:
func (k_ KeyedUnarchiver) DecodeIntForKey(key IString) int {
	rv := objc.Send[int](k_.ID, objc.Sel("decodeIntForKey:"), key)
	return rv
}


// Decodes and returns an object associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/decodeObject(forKey:)
func (k_ KeyedUnarchiver) DecodeObjectForKey(key IString) objc.ID {
	rv := objc.Send[objc.ID](k_.ID, objc.Sel("decodeObjectForKey:"), key)
	return rv
}


// Tells the receiver that you are finished decoding objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/finishDecoding()
func (k_ KeyedUnarchiver) FinishDecoding() {
	objc.Send[objc.ID](k_.ID, objc.Sel("finishDecoding"))
}


// Sets a translation mapping on this unarchiver to decode objects encoded with a given class name as instances of a given class instead.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/setClass(_:forClassName:)-swift.method
func (k_ KeyedUnarchiver) SetClassForClassName(cls objc.Class, codedName IString) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setClass:forClassName:"), cls, codedName)
}


// The action to take when this unarchiver fails to decode an entry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/decodingFailurePolicy
func (k_ KeyedUnarchiver) DecodingFailurePolicy() DecodingFailurePolicy {
	rv := objc.Send[DecodingFailurePolicy](k_.ID, objc.Sel("decodingFailurePolicy"))
	return rv
}


// The action to take when this unarchiver fails to decode an entry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/decodingFailurePolicy
func (k_ KeyedUnarchiver) SetDecodingFailurePolicy(value DecodingFailurePolicy) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setDecodingFailurePolicy:"), value)
}


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/delegate
func (k_ KeyedUnarchiver) Delegate() objc.ID {
	rv := objc.Send[objc.ID](k_.ID, objc.Sel("delegate"))
	return rv
}


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/delegate
func (k_ KeyedUnarchiver) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setDelegate:"), value)
}


// Indicates whether the receiver requires all unarchived classes to conform to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/requiresSecureCoding
func (k_ KeyedUnarchiver) RequiresSecureCoding() bool {
	rv := objc.Send[bool](k_.ID, objc.Sel("requiresSecureCoding"))
	return rv
}


// Indicates whether the receiver requires all unarchived classes to conform to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/requiresSecureCoding
func (k_ KeyedUnarchiver) SetRequiresSecureCoding(value bool) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setRequiresSecureCoding:"), value)
}



// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSKeyedArchiver */


/* debug [class_header]: Header for NSKeyedArchiver */
// The class instance for the [KeyedArchiver] class.
var (
	KeyedArchiverClass     _KeyedArchiverClass
	KeyedArchiverClassOnce sync.Once
)

func getKeyedArchiverClass() _KeyedArchiverClass {
	KeyedArchiverClassOnce.Do(func() {
		KeyedArchiverClass = _KeyedArchiverClass{objc.GetClass("NSKeyedArchiver")}
	})
	return KeyedArchiverClass
}

type _KeyedArchiverClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for KeyedArchiver */
// An interface definition for the [KeyedArchiver] class.
type IKeyedArchiver interface {
	ICoder
	
/* debug [class_interface_properties]: Properties for KeyedArchiver */
	// properties:
	EncodedData() IData
	OutputFormat() PropertyListFormat
	SetOutputFormat(value PropertyListFormat)
	RequiresSecureCoding() bool
	SetRequiresSecureCoding(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for KeyedArchiver */
	// methods:
	ClassNameForClass(cls objc.Class) IString
	EncodeDoubleForKey(value float64, key IString)
	EncodeInt32ForKey(value int32 /* not a class type */, key IString)
	EncodeFloatForKey(value float32, key IString)
	EncodeObjectForKey(object objc.IObject, key IString)
	EncodeBoolForKey(value bool, key IString)
	EncodeInt64ForKey(value int64, key IString)
	EncodeBytesLengthForKey(bytes objectivec.IObject, length uint, key IString)
	EncodeConditionalObjectForKey(object objc.IObject, key IString)
	EncodeIntForKey(value int, key IString)
	FinishEncoding()
	SetClassNameForClass(codedName IString, cls objc.Class)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for KeyedArchiver */
// Alloc allocates a new instance without initialization.
func (kc _KeyedArchiverClass) Alloc() KeyedArchiver {
	rv := objc.Send[KeyedArchiver](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (kc _KeyedArchiverClass) New() KeyedArchiver {
	rv := objc.Send[KeyedArchiver](objc.ID(kc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (k_ KeyedArchiver) Init() KeyedArchiver {
	rv := objc.Send[KeyedArchiver](k_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k_ KeyedArchiver) Autorelease() KeyedArchiver {
	rv := objc.Send[KeyedArchiver](k_.ID, objc.Sel("autorelease"))
	return rv
}

// NewKeyedArchiver creates a new KeyedArchiver instance.
func NewKeyedArchiver() KeyedArchiver {
	return getKeyedArchiverClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for KeyedArchiver */
// An encoder that stores an object’s data to an archive referenced by keys.
//
// , a concrete subclass of , provides a way to encode objects (and scalar values) into an architecture-independent format suitable for storage in a file. When you archive a set of objects, the archiver writes the class information and instance variables for each object to the archive. The companion class decodes the data in an archive and creates a set of objects equivalent to the original set. A keyed archive differs from a non-keyed archive in that all the objects and values encoded into the archive have names, or keys. When decoding a non-keyed archive, the decoder must decode values in the same order the original encoder used. When decoding a keyed archive, the decoder requests values by name, meaning it can decode values out of sequence or not at all. Keyed archives, therefore, provide better support for forward and backward compatibility. The keys given to encoded values must be unique only within the scope of the currently-encoding object. A keyed archive is hierarchical, so the keys used by object A to encode its instance variables don’t conflict with the keys used by object B. This is true even if A and B are instances of the same class. Within a single object, however, the keys used by a subclass can conflict with keys used in its superclasses. An object can write the archive data to a file or to a mutable-data object (an instance of ) that you provide.


// An encoder that stores an object’s data to an archive referenced by keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver
type KeyedArchiver struct {
	Coder
}

// KeyedArchiverFrom constructs a [KeyedArchiver] from an unsafe.Pointer.
//
// An encoder that stores an object’s data to an archive referenced by keys.
func KeyedArchiverFrom(ptr unsafe.Pointer) KeyedArchiver {
	return KeyedArchiver{
		Coder: CoderFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for KeyedArchiver */

// Initializes an archiver to encode data into a given a mutable-data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/init(forWritingWith:)
func NewKeyedArchiverForWritingWithMutableData(data IMutableData) KeyedArchiver {
	instance := getKeyedArchiverClass().Alloc()
	rv := objc.Send[KeyedArchiver](instance.ID, objc.Sel("initForWritingWithMutableData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewKeyedArchiverForWritingWithMutableData */


// Creates an archiver to encode data, and optionally disables secure coding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/init(requiringSecureCoding:)
func NewKeyedArchiverRequiringSecureCoding(requiresSecureCoding bool) KeyedArchiver {
	instance := getKeyedArchiverClass().Alloc()
	rv := objc.Send[KeyedArchiver](instance.ID, objc.Sel("initRequiringSecureCoding:"), requiresSecureCoding)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewKeyedArchiverRequiringSecureCoding */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for KeyedArchiver */

// Archives an object graph rooted at a given object to a file at a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/archiveRootObject(_:toFile:)
func (kc _KeyedArchiverClass) ArchiveRootObjectToFile(rootObject objc.IObject, path IString) bool {
	rv := objc.Send[bool](objc.ID(kc.class), objc.Sel("archiveRootObject:toFile:"), rootObject, path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ArchiveRootObjectToFile) */


// Returns a data object that contains the encoded form of the object graph formed by the given root object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/archivedData(withRootObject:)
func (kc _KeyedArchiverClass) ArchivedDataWithRootObject(rootObject objc.IObject) IData {
	rv := objc.Send[Data](objc.ID(kc.class), objc.Sel("archivedDataWithRootObject:"), rootObject)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ArchivedDataWithRootObject) */


// Encodes an object graph with the given root object into a data representation, optionally requiring secure coding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/archivedData(withRootObject:requiringSecureCoding:)
func (kc _KeyedArchiverClass) ArchivedDataWithRootObjectRequiringSecureCodingError(object objc.IObject, requiresSecureCoding bool, error_ IError) IData {
	rv := objc.Send[Data](objc.ID(kc.class), objc.Sel("archivedDataWithRootObject:requiringSecureCoding:error:"), object, requiresSecureCoding, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ArchivedDataWithRootObjectRequiringSecureCodingError) */


// Returns the class name with which the archiver class encodes instances of a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/className(for:)-swift.type.method
func (kc _KeyedArchiverClass) ClassNameForClass(cls objc.Class) IString {
	rv := objc.Send[String](objc.ID(kc.class), objc.Sel("classNameForClass:"), cls)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ClassNameForClass) */


// Sets a global translation mapping to encode instances of a given class with the provided name, rather than their real name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/setClassName(_:for:)-swift.type.method
func (kc _KeyedArchiverClass) SetClassNameForClass(codedName IString, cls objc.Class) {
	objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("setClassName:forClass:"), codedName, cls)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetClassNameForClass) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for KeyedArchiver */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for KeyedArchiver */

// Returns the class name with which this archiver encodes instances of a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/className(for:)-swift.method
func (k_ KeyedArchiver) ClassNameForClass(cls objc.Class) IString {
	rv := objc.Send[String](k_.ID, objc.Sel("classNameForClass:"), cls)
	return rv
}/* debug [instance_methods/method]: ClassNameForClass */


// Encodes a given value and associates it with a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/encode(_:forKey:)-1mkfl
func (k_ KeyedArchiver) EncodeDoubleForKey(value float64, key IString) {
	objc.Send[objc.ID](k_.ID, objc.Sel("encodeDouble:forKey:"), value, key)
}/* debug [instance_methods/method]: EncodeDoubleForKey */


// Encodes a given 32-bit integer value and associates it with a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/encode(_:forKey:)-5i7tc
func (k_ KeyedArchiver) EncodeInt32ForKey(value int32 /* not a class type */, key IString) {
	objc.Send[objc.ID](k_.ID, objc.Sel("encodeInt32:forKey:"), value, key)
}/* debug [instance_methods/method]: EncodeInt32ForKey */


// Encodes a given value and associates it with a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/encode(_:forKey:)-67rcs
func (k_ KeyedArchiver) EncodeFloatForKey(value float32, key IString) {
	objc.Send[objc.ID](k_.ID, objc.Sel("encodeFloat:forKey:"), value, key)
}/* debug [instance_methods/method]: EncodeFloatForKey */


// Encodes a given object and associates it with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/encode(_:forKey:)-9f4n9
func (k_ KeyedArchiver) EncodeObjectForKey(object objc.IObject, key IString) {
	objc.Send[objc.ID](k_.ID, objc.Sel("encodeObject:forKey:"), object, key)
}/* debug [instance_methods/method]: EncodeObjectForKey */


// Encodes a given Boolean value and associates it with a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/encode(_:forKey:)-9pxhm
func (k_ KeyedArchiver) EncodeBoolForKey(value bool, key IString) {
	objc.Send[objc.ID](k_.ID, objc.Sel("encodeBool:forKey:"), value, key)
}/* debug [instance_methods/method]: EncodeBoolForKey */


// Encodes a given 64-bit integer value and associates it with a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/encode(_:forKey:)-ycdd
func (k_ KeyedArchiver) EncodeInt64ForKey(value int64, key IString) {
	objc.Send[objc.ID](k_.ID, objc.Sel("encodeInt64:forKey:"), value, key)
}/* debug [instance_methods/method]: EncodeInt64ForKey */


// Encodes a given number of bytes from a given C array of bytes and associates them with a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/encodeBytes(_:length:forKey:)
func (k_ KeyedArchiver) EncodeBytesLengthForKey(bytes objectivec.IObject, length uint, key IString) {
	objc.Send[objc.ID](k_.ID, objc.Sel("encodeBytes:length:forKey:"), bytes, length, key)
}/* debug [instance_methods/method]: EncodeBytesLengthForKey */


// Encodes a reference to a given object and associates it with a key only if it has been unconditionally encoded elsewhere in the archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/encodeConditionalObject(_:forKey:)
func (k_ KeyedArchiver) EncodeConditionalObjectForKey(object objc.IObject, key IString) {
	objc.Send[objc.ID](k_.ID, objc.Sel("encodeConditionalObject:forKey:"), object, key)
}/* debug [instance_methods/method]: EncodeConditionalObjectForKey */


// Encodes a given value and associates it with a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/encodeInt:forKey:
func (k_ KeyedArchiver) EncodeIntForKey(value int, key IString) {
	objc.Send[objc.ID](k_.ID, objc.Sel("encodeInt:forKey:"), value, key)
}/* debug [instance_methods/method]: EncodeIntForKey */


// Instructs the receiver to construct the final data stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/finishEncoding()
func (k_ KeyedArchiver) FinishEncoding() {
	objc.Send[objc.ID](k_.ID, objc.Sel("finishEncoding"))
}/* debug [instance_methods/method]: FinishEncoding */


// Sets a mapping for this archiver to encode instances of a given class with the provided name, rather than their real name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/setClassName(_:for:)-swift.method
func (k_ KeyedArchiver) SetClassNameForClass(codedName IString, cls objc.Class) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setClassName:forClass:"), codedName, cls)
}/* debug [instance_methods/method]: SetClassNameForClass */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for KeyedArchiver */

// The encoded data for the archiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/encodedData
func (k_ KeyedArchiver) EncodedData() IData {
	rv := objc.Send[Data](k_.ID, objc.Sel("encodedData"))
	return rv
}/* debug [instance_properties/getter]: encodedData */


// The format in which the receiver encodes its data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/outputFormat
func (k_ KeyedArchiver) OutputFormat() PropertyListFormat {
	rv := objc.Send[PropertyListFormat](k_.ID, objc.Sel("outputFormat"))
	return rv
}/* debug [instance_properties/getter]: outputFormat */


// The format in which the receiver encodes its data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/outputFormat
func (k_ KeyedArchiver) SetOutputFormat(value PropertyListFormat) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setOutputFormat:"), value)
}/* debug [instance_properties/setter]: outputFormat */


// Indicates whether the archiver requires all archived classes to resist object substitution attacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/requiresSecureCoding
func (k_ KeyedArchiver) RequiresSecureCoding() bool {
	rv := objc.Send[bool](k_.ID, objc.Sel("requiresSecureCoding"))
	return rv
}/* debug [instance_properties/getter]: requiresSecureCoding */


// Indicates whether the archiver requires all archived classes to resist object substitution attacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/requiresSecureCoding
func (k_ KeyedArchiver) SetRequiresSecureCoding(value bool) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setRequiresSecureCoding:"), value)
}/* debug [instance_properties/setter]: requiresSecureCoding */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSKeyedArchiver */



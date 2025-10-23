// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Dictionary] class.
var (
	DictionaryClass     _DictionaryClass
	DictionaryClassOnce sync.Once
)

func getDictionaryClass() _DictionaryClass {
	DictionaryClassOnce.Do(func() {
		DictionaryClass = _DictionaryClass{objc.GetClass("NSDictionary")}
	})
	return DictionaryClass
}

type _DictionaryClass struct {
	class objc.Class
}

// An interface definition for the [Dictionary] class.
type IDictionary interface {
	objectivec.IObject
	// properties:
	Count() uint /* primitive/slice/pointer. */
	AllKeys() unsafe.Pointer
	SetAllKeys(value unsafe.Pointer)
	AllValues() unsafe.Pointer
	SetAllValues(value unsafe.Pointer)
	Description() IString
	SetDescription(value IString)
	DescriptionInStringsFileFormat() IString
	SetDescriptionInStringsFileFormat(value IString)
	// methods:
	EnumerateKeysAndObjectsUsingBlock(block unsafe.Pointer)
	FileHFSTypeCode() unsafe.Pointer
	FileIsImmutable() bool /* primitive/slice/pointer. */
	KeyEnumerator() unsafe.Pointer
	ObjectForKey(aKey unsafe.Pointer) unsafe.Pointer
}

// A static collection of objects associated with unique keys.
//
// You can use this type in Swift instead of a in cases that require reference semantics. The class declares the programmatic interface to objects that manage immutable associations of keys and values. For example, an interactive form could be represented as a dictionary, with the field names as keys, corresponding to user-entered values. Use this class or its subclass when you need a convenient and efficient way to retrieve data associated with an arbitrary key. creates static dictionaries, and creates dynamic dictionaries. (For convenience, the term refers to any instance of one of these classes without specifying its exact class membership.) A key-value pair within a dictionary is called an entry. Each entry consists of one object that represents the key and a second object that is that key’s value. Within a dictionary, the keys are unique. That is, no two keys in a single dictionary are equal (as determined by ). In general, a key can be any object (provided that it conforms to the protocol—see below), but note that when using key-value coding the key must be a string (see ). Neither a key nor a value can be ; if you need to represent a null value in a dictionary, you should use . is “toll-free bridged” with its Core Foundation counterpart, . See for more information on toll-free bridging.


// A static collection of objects associated with unique keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary
type Dictionary struct {
	objectivec.Object
}

// DictionaryFrom constructs a [Dictionary] from an unsafe.Pointer.
//
// A static collection of objects associated with unique keys.
func DictionaryFrom(ptr unsafe.Pointer) Dictionary {
	return Dictionary{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DictionaryClass) Alloc() Dictionary {
	rv := objc.Send[Dictionary](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DictionaryClass) New() Dictionary {
	rv := objc.Send[Dictionary](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ Dictionary) Init() Dictionary {
	rv := objc.Send[Dictionary](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ Dictionary) Autorelease() Dictionary {
	rv := objc.Send[Dictionary](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDictionary creates a new Dictionary instance.
func NewDictionary() Dictionary {
	return getDictionaryClass().New()
}



// Initializes a newly allocated dictionary using the keys and values found in a file at a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(contentsOfFile:)
func NewDictionaryWithContentsOfFile(path IString) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithContentsOfFile:"), path)
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated dictionary using the keys and values found at a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(contentsOfURL:)-4pv16
func NewDictionaryWithContentsOfURL(url objc.IObject /* cross-framework NSURL */) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated dictionary with entries constructed from the specified set of values and keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/initWithObjectsAndKeys:
func NewDictionaryWithObjectsAndKeys(firstObject objectivec.IObject) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithObjectsAndKeys:"), firstObject)
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated dictionary with key-value pairs constructed from the provided arrays of keys and objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(objects:forKeys:)
func NewDictionaryWithObjectsForKeys(objects []objc.ID /* already interface */, keys []objc.ID /* already interface */) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithObjects:forKeys:"), objects, keys)
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated dictionary with the specified number of key-value pairs constructed from the provided C arrays of keys and objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(objects:forKeys:count:)
func NewDictionaryWithObjectsForKeysCount(objects []unsafe.Pointer /* not a class type */, keys []objc.ID /* already interface */, cnt uint /* primitive/slice/pointer. */) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithObjects:forKeys:count:"), objects, keys, cnt)
	rv.Autorelease()
	return rv
}



// Creates a dictionary containing entries constructed from the contents of an array of keys and an array of values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithObjects:forKeys:
func (dc _DictionaryClass) DictionaryWithObjectsForKeys(objects []objc.ID /* already interface */, keys []objc.ID /* already interface */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dictionaryWithObjects:forKeys:"), objects, keys)
	return rv
}


// Creates a dictionary containing a specified number of objects from a C array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithObjects:forKeys:count:
func (dc _DictionaryClass) DictionaryWithObjectsForKeysCount(objects []unsafe.Pointer /* not a class type */, keys []objc.ID /* already interface */, cnt uint /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dictionaryWithObjects:forKeys:count:"), objects, keys, cnt)
	return rv
}


// Creates a dictionary containing entries constructed from the specified set of values and keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithObjectsAndKeys:
func (dc _DictionaryClass) DictionaryWithObjectsAndKeys(firstObject objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dictionaryWithObjectsAndKeys:"), firstObject)
	return rv
}


// Applies a given block object to the entries of the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/enumerateKeysAndObjects(_:)
func (d_ Dictionary) EnumerateKeysAndObjectsUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("enumerateKeysAndObjectsUsingBlock:"), block)
}


// Returns file’s HFS type code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileHFSTypeCode()
func (d_ Dictionary) FileHFSTypeCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileHFSTypeCode"))
	return rv
}


// Returns a Boolean value indicating whether the file is immutable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileIsImmutable()
func (d_ Dictionary) FileIsImmutable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("fileIsImmutable"))
	return rv
}


// Provides an enumerator to access the keys in the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/keyEnumerator()
func (d_ Dictionary) KeyEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("keyEnumerator"))
	return rv
}


// Returns the value associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/object(forKey:)
func (d_ Dictionary) ObjectForKey(aKey unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("objectForKey:"), aKey)
	return rv
}


// Returns the value associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/value(forKey:)
func (d_ Dictionary) ValueForKey(key IString) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("valueForKey:"), key)
	return rv
}


// The number of entries in the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/count
func (d_ Dictionary) Count() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](d_.ID, objc.Sel("count"))
	return rv
}


// A new array containing the dictionary’s keys, or an empty array if the dictionary has no entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/allkeys
func (d_ Dictionary) AllKeys() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("allKeys"))
	return rv
}


// A new array containing the dictionary’s keys, or an empty array if the dictionary has no entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/allkeys
func (d_ Dictionary) SetAllKeys(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAllKeys:"), value)
}


// A new array containing the dictionary’s values, or an empty array if the dictionary has no entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/allvalues
func (d_ Dictionary) AllValues() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("allValues"))
	return rv
}


// A new array containing the dictionary’s values, or an empty array if the dictionary has no entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/allvalues
func (d_ Dictionary) SetAllValues(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAllValues:"), value)
}


// A string that represents the contents of the dictionary, formatted as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/description
func (d_ Dictionary) Description() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("description"))
	return rv
}


// A string that represents the contents of the dictionary, formatted as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/description
func (d_ Dictionary) SetDescription(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDescription:"), value)
}


// A string that represents the contents of the dictionary, formatted in
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/descriptioninstringsfileformat
func (d_ Dictionary) DescriptionInStringsFileFormat() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("descriptionInStringsFileFormat"))
	return rv
}


// A string that represents the contents of the dictionary, formatted in
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/descriptioninstringsfileformat
func (d_ Dictionary) SetDescriptionInStringsFileFormat(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDescriptionInStringsFileFormat:"), value)
}



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
	AllKeys() unsafe.Pointer
	SetAllKeys(value unsafe.Pointer)
	AllValues() unsafe.Pointer
	SetAllValues(value unsafe.Pointer)
	Count() int /* primitive/slice/pointer */
	SetCount(value int /* primitive/slice/pointer */)
	Description() string /* primitive/slice/pointer */
	SetDescription(value string /* primitive/slice/pointer */)
	DescriptionInStringsFileFormat() string /* primitive/slice/pointer */
	SetDescriptionInStringsFileFormat(value string /* primitive/slice/pointer */)
	// methods:
	KeyEnumerator() unsafe.Pointer
	ObjectEnumerator() unsafe.Pointer
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



// Provides an enumerator to access the keys in the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/keyEnumerator()
func (d_ Dictionary) KeyEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("keyEnumerator"))
	return rv
}


// Returns an enumerator object that lets you access each value in the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/objectEnumerator()
func (d_ Dictionary) ObjectEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("objectEnumerator"))
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


// The number of entries in the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/count
func (d_ Dictionary) Count() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](d_.ID, objc.Sel("count"))
	return rv
}


// The number of entries in the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/count
func (d_ Dictionary) SetCount(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCount:"), value)
}


// A string that represents the contents of the dictionary, formatted as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/description
func (d_ Dictionary) Description() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](d_.ID, objc.Sel("description"))
	return rv
}


// A string that represents the contents of the dictionary, formatted as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/description
func (d_ Dictionary) SetDescription(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDescription:"), objc.String(value))
}


// A string that represents the contents of the dictionary, formatted in
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/descriptioninstringsfileformat
func (d_ Dictionary) DescriptionInStringsFileFormat() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](d_.ID, objc.Sel("descriptionInStringsFileFormat"))
	return rv
}


// A string that represents the contents of the dictionary, formatted in
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/descriptioninstringsfileformat
func (d_ Dictionary) SetDescriptionInStringsFileFormat(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDescriptionInStringsFileFormat:"), objc.String(value))
}



// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MutableDictionary] class.
var (
	MutableDictionaryClass     _MutableDictionaryClass
	MutableDictionaryClassOnce sync.Once
)

func getMutableDictionaryClass() _MutableDictionaryClass {
	MutableDictionaryClassOnce.Do(func() {
		MutableDictionaryClass = _MutableDictionaryClass{objc.GetClass("NSMutableDictionary")}
	})
	return MutableDictionaryClass
}

type _MutableDictionaryClass struct {
	class objc.Class
}

// An interface definition for the [MutableDictionary] class.
type IMutableDictionary interface {
	IDictionary
	// properties:
	// methods:
	AddApplicationParameterHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32 /* foo */) unsafe.Pointer
	AddAuthorizationResponseHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32 /* foo */) unsafe.Pointer
	AddByteSequenceHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32 /* foo */) unsafe.Pointer
	AddDescriptionHeader(inDescriptionString string /* primitive/slice/pointer */) unsafe.Pointer
	AddEntriesFromDictionary(otherDictionary IDictionary /* already interface */)
	AddHTTPHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32 /* foo */) unsafe.Pointer
	AddImageDescriptorHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32 /* foo */) unsafe.Pointer
	AddImageHandleHeader(type_ string /* primitive/slice/pointer */) unsafe.Pointer
	AddTargetHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32 /* foo */) unsafe.Pointer
	AddTime4ByteHeader(time4Byte uint32 /* foo */) unsafe.Pointer
	GetHeaderBytes() IMutableData
	RemoveObjectForKey(aKey unsafe.Pointer)
	SetObjectForKey(anObject unsafe.Pointer, aKey objectivec.IObject)
	SetObjectForKeyedSubscript(obj unsafe.Pointer, key objectivec.IObject)
}

// A dynamic collection of objects associated with unique keys.
//
// In Swift, you can use this type instead of a variable in cases that require reference semantics. The class declares the programmatic interface to objects that manage mutable associations of keys and values. It adds modification operations to the basic operations it inherits from . is “toll-free bridged” with its Core Foundation counterpart, . See for more information on toll-free bridging.


// A dynamic collection of objects associated with unique keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary
type MutableDictionary struct {
	Dictionary
}

// MutableDictionaryFrom constructs a [MutableDictionary] from an unsafe.Pointer.
//
// A dynamic collection of objects associated with unique keys.
func MutableDictionaryFrom(ptr unsafe.Pointer) MutableDictionary {
	return MutableDictionary{
		Dictionary: DictionaryFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableDictionaryClass) Alloc() MutableDictionary {
	rv := objc.Send[MutableDictionary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableDictionaryClass) New() MutableDictionary {
	rv := objc.Send[MutableDictionary](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableDictionary) Init() MutableDictionary {
	rv := objc.Send[MutableDictionary](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableDictionary) Autorelease() MutableDictionary {
	rv := objc.Send[MutableDictionary](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableDictionary creates a new MutableDictionary instance.
func NewMutableDictionary() MutableDictionary {
	return getMutableDictionaryClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(coder:)
func NewMutableDictionaryWithCoder(coder ICoder) MutableDictionary {
	instance := getMutableDictionaryClass().Alloc()
	rv := objc.Send[MutableDictionary](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(OBEXHeadersData:headersDataSize:)
func NewMutableDictionaryWithOBEXHeadersDataHeadersDataSize(inHeadersData unsafe.Pointer, inDataSize uintptr /* foo */) MutableDictionary {
	rv := objc.Send[MutableDictionary](objc.ID(getMutableDictionaryClass().class), objc.Sel("dictionaryWithOBEXHeadersData:headersDataSize:"), inHeadersData, inDataSize)
	return rv
}



// Creates and returns a mutable dictionary, initially giving it enough allocated memory to hold a given number of entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/dictionaryWithCapacity:
func (mc _MutableDictionaryClass) DictionaryWithCapacity(numItems uint /* primitive/slice/pointer */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("dictionaryWithCapacity:"), numItems)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(OBEXHeadersData:headersDataSize:)
func (mc _MutableDictionaryClass) DictionaryWithOBEXHeadersDataHeadersDataSize(inHeadersData unsafe.Pointer, inDataSize uintptr /* foo */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("dictionaryWithOBEXHeadersData:headersDataSize:"), inHeadersData, inDataSize)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/withOBEXHeadersData:headersDataSize:
func (mc _MutableDictionaryClass) WithOBEXHeadersDataHeadersDataSize(inHeadersData unsafe.Pointer, inDataSize uintptr /* foo */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("withOBEXHeadersData:headersDataSize:"), inHeadersData, inDataSize)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addApplicationParameterHeader(_:length:)
func (m_ MutableDictionary) AddApplicationParameterHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32 /* foo */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("addApplicationParameterHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addAuthorizationResponseHeader(_:length:)
func (m_ MutableDictionary) AddAuthorizationResponseHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32 /* foo */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("addAuthorizationResponseHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addByteSequenceHeader(_:length:)
func (m_ MutableDictionary) AddByteSequenceHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32 /* foo */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("addByteSequenceHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addDescriptionHeader(_:)
func (m_ MutableDictionary) AddDescriptionHeader(inDescriptionString string /* primitive/slice/pointer */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("addDescriptionHeader:"), objc.String(inDescriptionString))
	return rv
}


// Adds to the receiving dictionary the entries from another dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addEntries(from:)
func (m_ MutableDictionary) AddEntriesFromDictionary(otherDictionary IDictionary /* already interface */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addEntriesFromDictionary:"), otherDictionary)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addHTTPHeader(_:length:)
func (m_ MutableDictionary) AddHTTPHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32 /* foo */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("addHTTPHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addImageDescriptorHeader(_:length:)
func (m_ MutableDictionary) AddImageDescriptorHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32 /* foo */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("addImageDescriptorHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addImageHandleHeader(_:)
func (m_ MutableDictionary) AddImageHandleHeader(type_ string /* primitive/slice/pointer */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("addImageHandleHeader:"), objc.String(type_))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addTargetHeader(_:length:)
func (m_ MutableDictionary) AddTargetHeaderLength(inHeaderData unsafe.Pointer, inHeaderDataLength uint32 /* foo */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("addTargetHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addTime4ByteHeader(_:)
func (m_ MutableDictionary) AddTime4ByteHeader(time4Byte uint32 /* foo */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("addTime4ByteHeader:"), time4Byte)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/getHeaderBytes()
func (m_ MutableDictionary) GetHeaderBytes() IMutableData {
	rv := objc.Send[MutableData](m_.ID, objc.Sel("getHeaderBytes"))
	return rv
}


// Removes a given key and its associated value from the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/removeObject(forKey:)
func (m_ MutableDictionary) RemoveObjectForKey(aKey unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectForKey:"), aKey)
}


// Adds a given key-value pair to the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/setObject(_:forKey:)
func (m_ MutableDictionary) SetObjectForKey(anObject unsafe.Pointer, aKey objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:forKey:"), anObject, aKey)
}


// Adds a given key-value pair to the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/setObject:forKeyedSubscript:
func (m_ MutableDictionary) SetObjectForKeyedSubscript(obj unsafe.Pointer, key objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:forKeyedSubscript:"), obj, key)
}


// Adds a given key-value pair to the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/setValue(_:forKey:)
func (m_ MutableDictionary) SetValueForKey(value unsafe.Pointer, key string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:forKey:"), value, objc.String(key))
}



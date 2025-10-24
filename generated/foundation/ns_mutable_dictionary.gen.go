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
	AddApplicationParameterHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject
	AddAuthorizationChallengeHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject
	AddAuthorizationResponseHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject
	AddBodyHeaderLengthEndOfBody(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */, isEndOfBody bool) objectivec.IObject
	AddByteSequenceHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject
	AddConnectionIDHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject
	AddCountHeader(inCount uint32 /* not a class type */) objectivec.IObject
	AddDescriptionHeader(inDescriptionString IString) objectivec.IObject
	AddEntriesFromDictionary(otherDictionary IDictionary)
	AddHTTPHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject
	AddImageDescriptorHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject
	AddImageHandleHeader(type_ IString) objectivec.IObject
	AddLengthHeader(length uint32 /* not a class type */) objectivec.IObject
	AddNameHeader(inNameString IString) objectivec.IObject
	AddObjectClassHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject
	AddTargetHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject
	AddTime4ByteHeader(time4Byte uint32 /* not a class type */) objectivec.IObject
	AddTimeISOHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject
	AddTypeHeader(type_ IString) objectivec.IObject
	AddUserDefinedHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject
	AddWhoHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject
	GetHeaderBytes() IMutableData
	RemoveAllObjects()
	RemoveObjectForKey(aKey objectivec.IObject)
	RemoveObjectsForKeys(keyArray []objc.ID)
	SetDictionary(otherDictionary IDictionary)
	SetObjectForKey(anObject objectivec.IObject, aKey unsafe.Pointer)
	SetObjectForKeyedSubscript(obj objectivec.IObject, key unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (mc _MutableDictionaryClass) Alloc() MutableDictionary {
	rv := objc.Send[MutableDictionary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Initializes a newly allocated mutable dictionary, allocating enough memory to hold entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(capacity:)
func NewMutableDictionaryWithCapacity(numItems uint) MutableDictionary {
	instance := getMutableDictionaryClass().Alloc()
	rv := objc.Send[MutableDictionary](instance.ID, objc.Sel("initWithCapacity:"), numItems)
	rv.Autorelease()
	return rv
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/initWithContentsOfFile:
func NewMutableDictionaryWithContentsOfFile(path IString) MutableDictionary {
	instance := getMutableDictionaryClass().Alloc()
	rv := objc.Send[MutableDictionary](instance.ID, objc.Sel("initWithContentsOfFile:"), path)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/initWithContentsOfURL:
func NewMutableDictionaryWithContentsOfURL(url IURL) MutableDictionary {
	instance := getMutableDictionaryClass().Alloc()
	rv := objc.Send[MutableDictionary](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(OBEXHeadersData:)
func NewMutableDictionaryWithOBEXHeadersData(inHeadersData IData) MutableDictionary {
	rv := objc.Send[MutableDictionary](objc.ID(getMutableDictionaryClass().class), objc.Sel("dictionaryWithOBEXHeadersData:"), inHeadersData)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(OBEXHeadersData:headersDataSize:)
func NewMutableDictionaryWithOBEXHeadersDataHeadersDataSize(inHeadersData objectivec.IObject, inDataSize uintptr /* not a class type */) MutableDictionary {
	rv := objc.Send[MutableDictionary](objc.ID(getMutableDictionaryClass().class), objc.Sel("dictionaryWithOBEXHeadersData:headersDataSize:"), inHeadersData, inDataSize)
	return rv
}


// Creates a mutable dictionary which is optimized for dealing with a known set of keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(sharedKeySet:)
func NewMutableDictionaryWithSharedKeySet(keyset objc.IObject) MutableDictionary {
	rv := objc.Send[MutableDictionary](objc.ID(getMutableDictionaryClass().class), objc.Sel("dictionaryWithSharedKeySet:"), keyset)
	return rv
}







// Creates and returns a mutable dictionary, initially giving it enough allocated memory to hold a given number of entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/dictionaryWithCapacity:
func (mc _MutableDictionaryClass) DictionaryWithCapacity(numItems uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("dictionaryWithCapacity:"), numItems)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/dictionaryWithContentsOfFile:
func (mc _MutableDictionaryClass) DictionaryWithContentsOfFile(path IString) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("dictionaryWithContentsOfFile:"), path)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(OBEXHeadersData:)
func (mc _MutableDictionaryClass) DictionaryWithOBEXHeadersData(inHeadersData IData) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("dictionaryWithOBEXHeadersData:"), inHeadersData)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(OBEXHeadersData:headersDataSize:)
func (mc _MutableDictionaryClass) DictionaryWithOBEXHeadersDataHeadersDataSize(inHeadersData objectivec.IObject, inDataSize uintptr /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("dictionaryWithOBEXHeadersData:headersDataSize:"), inHeadersData, inDataSize)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(contentsOfURL:)
func (mc _MutableDictionaryClass) DictionaryWithContentsOfURL(url IURL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("dictionaryWithContentsOfURL:"), url)
	return rv
}


// Creates a mutable dictionary which is optimized for dealing with a known set of keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/init(sharedKeySet:)
func (mc _MutableDictionaryClass) DictionaryWithSharedKeySet(keyset objc.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("dictionaryWithSharedKeySet:"), keyset)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/withOBEXHeadersData:headersDataSize:
func (mc _MutableDictionaryClass) WithOBEXHeadersDataHeadersDataSize(inHeadersData objectivec.IObject, inDataSize uintptr /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("withOBEXHeadersData:headersDataSize:"), inHeadersData, inDataSize)
	return rv
}












// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addApplicationParameterHeader(_:length:)
func (m_ MutableDictionary) AddApplicationParameterHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addApplicationParameterHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addAuthorizationChallengeHeader(_:length:)
func (m_ MutableDictionary) AddAuthorizationChallengeHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addAuthorizationChallengeHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addAuthorizationResponseHeader(_:length:)
func (m_ MutableDictionary) AddAuthorizationResponseHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addAuthorizationResponseHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addBodyHeader(_:length:endOfBody:)
func (m_ MutableDictionary) AddBodyHeaderLengthEndOfBody(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */, isEndOfBody bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addBodyHeader:length:endOfBody:"), inHeaderData, inHeaderDataLength, isEndOfBody)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addByteSequenceHeader(_:length:)
func (m_ MutableDictionary) AddByteSequenceHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addByteSequenceHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addConnectionIDHeader(_:length:)
func (m_ MutableDictionary) AddConnectionIDHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addConnectionIDHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addCountHeader(_:)
func (m_ MutableDictionary) AddCountHeader(inCount uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addCountHeader:"), inCount)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addDescriptionHeader(_:)
func (m_ MutableDictionary) AddDescriptionHeader(inDescriptionString IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addDescriptionHeader:"), inDescriptionString)
	return rv
}


// Adds to the receiving dictionary the entries from another dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addEntries(from:)
func (m_ MutableDictionary) AddEntriesFromDictionary(otherDictionary IDictionary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addEntriesFromDictionary:"), otherDictionary)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addHTTPHeader(_:length:)
func (m_ MutableDictionary) AddHTTPHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addHTTPHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addImageDescriptorHeader(_:length:)
func (m_ MutableDictionary) AddImageDescriptorHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addImageDescriptorHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addImageHandleHeader(_:)
func (m_ MutableDictionary) AddImageHandleHeader(type_ IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addImageHandleHeader:"), type_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addLengthHeader(_:)
func (m_ MutableDictionary) AddLengthHeader(length uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addLengthHeader:"), length)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addNameHeader(_:)
func (m_ MutableDictionary) AddNameHeader(inNameString IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addNameHeader:"), inNameString)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addObjectClassHeader(_:length:)
func (m_ MutableDictionary) AddObjectClassHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addObjectClassHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addTargetHeader(_:length:)
func (m_ MutableDictionary) AddTargetHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addTargetHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addTime4ByteHeader(_:)
func (m_ MutableDictionary) AddTime4ByteHeader(time4Byte uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addTime4ByteHeader:"), time4Byte)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addTimeISOHeader(_:length:)
func (m_ MutableDictionary) AddTimeISOHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addTimeISOHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addTypeHeader(_:)
func (m_ MutableDictionary) AddTypeHeader(type_ IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addTypeHeader:"), type_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addUserDefinedHeader(_:length:)
func (m_ MutableDictionary) AddUserDefinedHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addUserDefinedHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/addWhoHeader(_:length:)
func (m_ MutableDictionary) AddWhoHeaderLength(inHeaderData objectivec.IObject, inHeaderDataLength uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addWhoHeader:length:"), inHeaderData, inHeaderDataLength)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/getHeaderBytes()
func (m_ MutableDictionary) GetHeaderBytes() IMutableData {
	rv := objc.Send[MutableData](m_.ID, objc.Sel("getHeaderBytes"))
	return rv
}


// Empties the dictionary of its entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/removeAllObjects()
func (m_ MutableDictionary) RemoveAllObjects() {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAllObjects"))
}


// Removes a given key and its associated value from the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/removeObject(forKey:)
func (m_ MutableDictionary) RemoveObjectForKey(aKey objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectForKey:"), aKey)
}


// Removes from the dictionary entries specified by elements in a given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/removeObjects(forKeys:)
func (m_ MutableDictionary) RemoveObjectsForKeys(keyArray []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectsForKeys:"), keyArray)
}


// Sets the contents of the receiving dictionary to entries in a given dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/setDictionary(_:)
func (m_ MutableDictionary) SetDictionary(otherDictionary IDictionary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDictionary:"), otherDictionary)
}


// Adds a given key-value pair to the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/setObject(_:forKey:)
func (m_ MutableDictionary) SetObjectForKey(anObject objectivec.IObject, aKey unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:forKey:"), anObject, aKey)
}


// Adds a given key-value pair to the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/setObject:forKeyedSubscript:
func (m_ MutableDictionary) SetObjectForKeyedSubscript(obj objectivec.IObject, key unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:forKeyedSubscript:"), obj, key)
}


// Adds a given key-value pair to the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableDictionary/setValue(_:forKey:)
func (m_ MutableDictionary) SetValueForKey(value objectivec.IObject, key IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:forKey:"), value, key)
}













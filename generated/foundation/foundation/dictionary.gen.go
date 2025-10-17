// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Dictionary] class.
var DictionaryClass objc.Class

func init() {
	DictionaryClass = objc.GetClass("NSDictionary")
}

type Dictionary struct {
	objc.ID
}

func DictionaryFrom(ptr unsafe.Pointer) Dictionary {
	return Dictionary{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc Dictionary) Alloc() Dictionary {
	ret := objc.ID(DictionaryClass).Send(objc.RegisterName("alloc"))
	return Dictionary{ret}
}

// Init initializes the instance.
func (d_ Dictionary) Init() Dictionary {
	ret := d_.ID.Send(objc.RegisterName("init"))
	return Dictionary{ret}
}
// Initializes a newly allocated dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/init()
func NewDictionary() Dictionary {
	instance := Dictionary{}.Alloc()
	instance = instance.Init()
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Creates a dictionary initialized from data in the provided unarchiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/init(coder:)
func NewDictionaryWithCoder(coder unsafe.Pointer) Dictionary {
	instance := Dictionary{}.Alloc()
	sel := objc.RegisterName("initWithCoder:")
	ret := instance.ID.Send(sel, coder)
	instance = Dictionary{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly allocated dictionary using the keys and values found in a file at a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/init(contentsOfFile:)
func NewDictionaryWithContentsOfFile(path string) Dictionary {
	instance := Dictionary{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfFile:")
	ret := instance.ID.Send(sel, path)
	instance = Dictionary{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly allocated dictionary using the keys and values found at a given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/init(contentsOfURL:)-4pv16
func NewDictionaryWithContentsOfURL(url unsafe.Pointer) Dictionary {
	instance := Dictionary{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfURL:")
	ret := instance.ID.Send(sel, url)
	instance = Dictionary{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly allocated dictionary using the keys and values found at a given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/init(contentsOfURL:error:)
func NewDictionaryWithContentsOfURLError(url unsafe.Pointer, error unsafe.Pointer) Dictionary {
	instance := Dictionary{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfURL:error:")
	ret := instance.ID.Send(sel, url, error)
	instance = Dictionary{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly allocated dictionary by placing in it the keys and values contained in another given dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/init(dictionary:)-9fw1u
func NewDictionaryWithDictionary(otherDictionary unsafe.Pointer) Dictionary {
	instance := Dictionary{}.Alloc()
	sel := objc.RegisterName("initWithDictionary:")
	ret := instance.ID.Send(sel, otherDictionary)
	instance = Dictionary{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly allocated dictionary using the objects contained in another given dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/init(dictionary:copyItems:)
func NewDictionaryWithDictionaryCopyItems(otherDictionary unsafe.Pointer, flag bool) Dictionary {
	instance := Dictionary{}.Alloc()
	sel := objc.RegisterName("initWithDictionary:copyItems:")
	ret := instance.ID.Send(sel, otherDictionary, flag)
	instance = Dictionary{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly allocated dictionary with key-value pairs constructed from the provided arrays of keys and objects. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/init(objects:forKeys:)
func NewDictionaryWithObjectsForKeys(objects unsafe.Pointer, keys unsafe.Pointer) Dictionary {
	instance := Dictionary{}.Alloc()
	sel := objc.RegisterName("initWithObjects:forKeys:")
	ret := instance.ID.Send(sel, objects, keys)
	instance = Dictionary{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly allocated dictionary with the specified number of key-value pairs constructed from the provided C arrays of keys and objects. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/init(objects:forKeys:count:)
func NewDictionaryWithObjectsForKeysCount(objects unsafe.Pointer, keys unsafe.Pointer, cnt uint) Dictionary {
	instance := Dictionary{}.Alloc()
	sel := objc.RegisterName("initWithObjects:forKeys:count:")
	ret := instance.ID.Send(sel, objects, keys, cnt)
	instance = Dictionary{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly allocated dictionary with entries constructed from the specified set of values and keys. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/initWithObjectsAndKeys:
func NewDictionaryWithObjectsAndKeys(firstObject objc.ID) Dictionary {
	instance := Dictionary{}.Alloc()
	sel := objc.RegisterName("initWithObjectsAndKeys:")
	ret := instance.ID.Send(sel, firstObject)
	instance = Dictionary{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Creates an empty dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/dictionary
func (dc Dictionary) Dictionary() unsafe.Pointer {
	sel := objc.RegisterName("dictionary")
	ret := objc.ID(DictionaryClass).Send(sel)
	return unsafe.Pointer(ret)
}
// Creates a dictionary using the keys and values found in a file specified by a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/dictionaryWithContentsOfFile:
func (dc Dictionary) DictionaryWithContentsOfFile(path string) unsafe.Pointer {
	sel := objc.RegisterName("dictionaryWithContentsOfFile:")
	ret := objc.ID(DictionaryClass).Send(sel, path)
	return unsafe.Pointer(ret)
}
// Creates a dictionary using the keys and values found in a resource specified by a given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/dictionaryWithContentsOfURL:error:
func (dc Dictionary) DictionaryWithContentsOfURLError(url unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dictionaryWithContentsOfURL:error:")
	ret := objc.ID(DictionaryClass).Send(sel, url, error)
	return unsafe.Pointer(ret)
}
// Creates a dictionary containing the keys and values from another given dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/dictionaryWithDictionary:
func (dc Dictionary) DictionaryWithDictionary(dict unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dictionaryWithDictionary:")
	ret := objc.ID(DictionaryClass).Send(sel, dict)
	return unsafe.Pointer(ret)
}
// Creates a dictionary containing entries constructed from the contents of an array of keys and an array of values. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/dictionaryWithObjects:forKeys:
func (dc Dictionary) DictionaryWithObjectsForKeys(objects unsafe.Pointer, keys unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dictionaryWithObjects:forKeys:")
	ret := objc.ID(DictionaryClass).Send(sel, objects, keys)
	return unsafe.Pointer(ret)
}
// Creates a dictionary containing a specified number of objects from a C array. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/dictionaryWithObjects:forKeys:count:
func (dc Dictionary) DictionaryWithObjectsForKeysCount(objects unsafe.Pointer, keys unsafe.Pointer, cnt uint) unsafe.Pointer {
	sel := objc.RegisterName("dictionaryWithObjects:forKeys:count:")
	ret := objc.ID(DictionaryClass).Send(sel, objects, keys, cnt)
	return unsafe.Pointer(ret)
}
// Creates a dictionary containing entries constructed from the specified set of values and keys. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/dictionaryWithObjectsAndKeys:
func (dc Dictionary) DictionaryWithObjectsAndKeys(firstObject objc.ID) unsafe.Pointer {
	sel := objc.RegisterName("dictionaryWithObjectsAndKeys:")
	ret := objc.ID(DictionaryClass).Send(sel, firstObject)
	return unsafe.Pointer(ret)
}
// Creates a dictionary using the keys and values found in a resource specified by a given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/init(contentsOfURL:)-98pl3
func (dc Dictionary) DictionaryWithContentsOfURL(url unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dictionaryWithContentsOfURL:")
	ret := objc.ID(DictionaryClass).Send(sel, url)
	return unsafe.Pointer(ret)
}
// Creates a dictionary containing a given key and value. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/init(object:forKey:)
func (dc Dictionary) DictionaryWithObjectForKey(object unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dictionaryWithObject:forKey:")
	ret := objc.ID(DictionaryClass).Send(sel, object, key)
	return unsafe.Pointer(ret)
}
// Creates a shared key set object for the specified keys. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/sharedKeySet(forKeys:)
func (dc Dictionary) SharedKeySetForKeys(keys unsafe.Pointer) objc.ID {
	sel := objc.RegisterName("sharedKeySetForKeys:")
	ret := objc.ID(DictionaryClass).Send(sel, keys)
	return ret
}
// Returns a new array containing the keys corresponding to all occurrences of a given object in the dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/allKeys(for:)
func (d_ Dictionary) AllKeysForObject(anObject unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("allKeysForObject:")
	ret := d_.ID.Send(sel, anObject)
	return unsafe.Pointer(ret)
}
// Returns by reference a C array of objects over which the sender should iterate. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/countByEnumeratingWithState:objects:count:
func (d_ Dictionary) CountByEnumeratingWithStateObjectsCount(state unsafe.Pointer, buffer unsafe.Pointer, len uint) uint {
	sel := objc.RegisterName("countByEnumeratingWithState:objects:count:")
	ret := d_.ID.Send(sel, state, buffer, len)
	return uint(ret)
}
// Returns a string object that represents the contents of the dictionary, formatted as a property list. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/description(withLocale:)
func (d_ Dictionary) DescriptionWithLocale(locale objc.ID) unsafe.Pointer {
	sel := objc.RegisterName("descriptionWithLocale:")
	ret := d_.ID.Send(sel, locale)
	return unsafe.Pointer(ret)
}
// Returns a string object that represents the contents of the dictionary, formatted as a property list. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/description(withLocale:indent:)
func (d_ Dictionary) DescriptionWithLocaleIndent(locale objc.ID, level uint) unsafe.Pointer {
	sel := objc.RegisterName("descriptionWithLocale:indent:")
	ret := d_.ID.Send(sel, locale, level)
	return unsafe.Pointer(ret)
}
// Applies a given block object to the entries of the dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/enumerateKeysAndObjects(_:)
func (d_ Dictionary) EnumerateKeysAndObjectsUsingBlock(block unsafe.Pointer) {
	sel := objc.RegisterName("enumerateKeysAndObjectsUsingBlock:")
	d_.ID.Send(sel, block)
}
// Applies a given block object to the entries of the dictionary, with options specifying how the enumeration is performed. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/enumerateKeysAndObjects(options:using:)
func (d_ Dictionary) EnumerateKeysAndObjectsWithOptionsUsingBlock(opts unsafe.Pointer, block unsafe.Pointer) {
	sel := objc.RegisterName("enumerateKeysAndObjectsWithOptions:usingBlock:")
	d_.ID.Send(sel, opts, block)
}
// Returns the file’s creation date. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/fileCreationDate()
func (d_ Dictionary) FileCreationDate() unsafe.Pointer {
	sel := objc.RegisterName("fileCreationDate")
	ret := d_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns a Boolean value indicating whether the file hides its extension. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/fileExtensionHidden()
func (d_ Dictionary) FileExtensionHidden() bool {
	sel := objc.RegisterName("fileExtensionHidden")
	ret := d_.ID.Send(sel)
	return ret != 0
}
// Returns file’s group owner account ID. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/fileGroupOwnerAccountID()
func (d_ Dictionary) FileGroupOwnerAccountID() unsafe.Pointer {
	sel := objc.RegisterName("fileGroupOwnerAccountID")
	ret := d_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns the file’s group owner account name. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/fileGroupOwnerAccountName()
func (d_ Dictionary) FileGroupOwnerAccountName() unsafe.Pointer {
	sel := objc.RegisterName("fileGroupOwnerAccountName")
	ret := d_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns the file’s HFS creator code. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/fileHFSCreatorCode()
func (d_ Dictionary) FileHFSCreatorCode() unsafe.Pointer {
	sel := objc.RegisterName("fileHFSCreatorCode")
	ret := d_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns file’s HFS type code. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/fileHFSTypeCode()
func (d_ Dictionary) FileHFSTypeCode() unsafe.Pointer {
	sel := objc.RegisterName("fileHFSTypeCode")
	ret := d_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns a Boolean value indicating whether the file is append only. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/fileIsAppendOnly()
func (d_ Dictionary) FileIsAppendOnly() bool {
	sel := objc.RegisterName("fileIsAppendOnly")
	ret := d_.ID.Send(sel)
	return ret != 0
}
// Returns a Boolean value indicating whether the file is immutable. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/fileIsImmutable()
func (d_ Dictionary) FileIsImmutable() bool {
	sel := objc.RegisterName("fileIsImmutable")
	ret := d_.ID.Send(sel)
	return ret != 0
}
// Returns file’s modification date. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/fileModificationDate()
func (d_ Dictionary) FileModificationDate() unsafe.Pointer {
	sel := objc.RegisterName("fileModificationDate")
	ret := d_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns the file’s owner account ID. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/fileOwnerAccountID()
func (d_ Dictionary) FileOwnerAccountID() unsafe.Pointer {
	sel := objc.RegisterName("fileOwnerAccountID")
	ret := d_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns the file’s owner account name. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/fileOwnerAccountName()
func (d_ Dictionary) FileOwnerAccountName() unsafe.Pointer {
	sel := objc.RegisterName("fileOwnerAccountName")
	ret := d_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns the file’s POSIX permissions. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/filePosixPermissions()
func (d_ Dictionary) FilePosixPermissions() uint {
	sel := objc.RegisterName("filePosixPermissions")
	ret := d_.ID.Send(sel)
	return uint(ret)
}
// Returns the file’s size, in bytes. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/fileSize()
func (d_ Dictionary) FileSize() unsafe.Pointer {
	sel := objc.RegisterName("fileSize")
	ret := d_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns the filesystem file number. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/fileSystemFileNumber()
func (d_ Dictionary) FileSystemFileNumber() uint {
	sel := objc.RegisterName("fileSystemFileNumber")
	ret := d_.ID.Send(sel)
	return uint(ret)
}
// Returns the filesystem number. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/fileSystemNumber()
func (d_ Dictionary) FileSystemNumber() int {
	sel := objc.RegisterName("fileSystemNumber")
	ret := d_.ID.Send(sel)
	return int(ret)
}
// Returns the file type. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/fileType()
func (d_ Dictionary) FileType() unsafe.Pointer {
	sel := objc.RegisterName("fileType")
	ret := d_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns by reference C arrays of the keys and values in the dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/getObjects:andKeys:
func (d_ Dictionary) GetObjectsAndKeys(objects unsafe.Pointer, keys unsafe.Pointer) {
	sel := objc.RegisterName("getObjects:andKeys:")
	d_.ID.Send(sel, objects, keys)
}
// Returns by reference C arrays of the keys and values in the dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/getObjects:andKeys:count:
func (d_ Dictionary) GetObjectsAndKeysCount(objects unsafe.Pointer, keys unsafe.Pointer, count uint) {
	sel := objc.RegisterName("getObjects:andKeys:count:")
	d_.ID.Send(sel, objects, keys, count)
}
// Returns a Boolean value that indicates whether the contents of the receiving dictionary are equal to the contents of another given dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/isEqual(to:)
func (d_ Dictionary) IsEqualToDictionary(otherDictionary unsafe.Pointer) bool {
	sel := objc.RegisterName("isEqualToDictionary:")
	ret := d_.ID.Send(sel, otherDictionary)
	return ret != 0
}
// Provides an enumerator to access the keys in the dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/keyEnumerator()
func (d_ Dictionary) KeyEnumerator() unsafe.Pointer {
	sel := objc.RegisterName("keyEnumerator")
	ret := d_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns the set of keys whose corresponding value satisfies a constraint described by a block object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/keysOfEntries(options:passingTest:)
func (d_ Dictionary) KeysOfEntriesWithOptionsPassingTest(opts unsafe.Pointer, predicate unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("keysOfEntriesWithOptions:passingTest:")
	ret := d_.ID.Send(sel, opts, predicate)
	return unsafe.Pointer(ret)
}
// Returns the set of keys whose corresponding value satisfies a constraint described by a block object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/keysOfEntries(passingTest:)
func (d_ Dictionary) KeysOfEntriesPassingTest(predicate unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("keysOfEntriesPassingTest:")
	ret := d_.ID.Send(sel, predicate)
	return unsafe.Pointer(ret)
}
// Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values using a given comparator block. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/keysSortedByValue(comparator:)
func (d_ Dictionary) KeysSortedByValueUsingComparator(cmptr unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("keysSortedByValueUsingComparator:")
	ret := d_.ID.Send(sel, cmptr)
	return unsafe.Pointer(ret)
}
// Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values using a given comparator block and a specified set of options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/keysSortedByValue(options:usingComparator:)
func (d_ Dictionary) KeysSortedByValueWithOptionsUsingComparator(opts unsafe.Pointer, cmptr unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("keysSortedByValueWithOptions:usingComparator:")
	ret := d_.ID.Send(sel, opts, cmptr)
	return unsafe.Pointer(ret)
}
// Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/keysSortedByValue(using:)
func (d_ Dictionary) KeysSortedByValueUsingSelector(comparator objc.SEL) unsafe.Pointer {
	sel := objc.RegisterName("keysSortedByValueUsingSelector:")
	ret := d_.ID.Send(sel, comparator)
	return unsafe.Pointer(ret)
}
// Returns the value associated with a given key. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/object(forKey:)
func (d_ Dictionary) ObjectForKey(aKey unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("objectForKey:")
	ret := d_.ID.Send(sel, aKey)
	return unsafe.Pointer(ret)
}
// Returns an enumerator object that lets you access each value in the dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/objectEnumerator()
func (d_ Dictionary) ObjectEnumerator() unsafe.Pointer {
	sel := objc.RegisterName("objectEnumerator")
	ret := d_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns as a static array the set of objects from the dictionary that corresponds to the specified keys. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/objects(forKeys:notFoundMarker:)
func (d_ Dictionary) ObjectsForKeysNotFoundMarker(keys unsafe.Pointer, marker unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("objectsForKeys:notFoundMarker:")
	ret := d_.ID.Send(sel, keys, marker)
	return unsafe.Pointer(ret)
}
// Returns the value associated with a given key. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/subscript(_:)-52n56
func (d_ Dictionary) ObjectForKeyedSubscript(key unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("objectForKeyedSubscript:")
	ret := d_.ID.Send(sel, key)
	return unsafe.Pointer(ret)
}
// Returns the value associated with a given key. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/value(forKey:)
func (d_ Dictionary) ValueForKey(key string) unsafe.Pointer {
	sel := objc.RegisterName("valueForKey:")
	ret := d_.ID.Send(sel, key)
	return unsafe.Pointer(ret)
}
// Writes a property list representation of the contents of the dictionary to a given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/write(to:)
func (d_ Dictionary) WriteToURLError(url unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("writeToURL:error:")
	ret := d_.ID.Send(sel, url, error)
	return ret != 0
}
// Writes a property list representation of the contents of the dictionary to a given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/write(to:atomically:)
func (d_ Dictionary) WriteToURLAtomically(url unsafe.Pointer, atomically bool) bool {
	sel := objc.RegisterName("writeToURL:atomically:")
	ret := d_.ID.Send(sel, url, atomically)
	return ret != 0
}
// Writes a property list representation of the contents of the dictionary to a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSDictionary/write(toFile:atomically:)
func (d_ Dictionary) WriteToFileAtomically(path string, useAuxiliaryFile bool) bool {
	sel := objc.RegisterName("writeToFile:atomically:")
	ret := d_.ID.Send(sel, path, useAuxiliaryFile)
	return ret != 0
}


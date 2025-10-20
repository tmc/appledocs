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
	AllKeysForObject(anObject unsafe.Pointer) []objc.ID
	CountByEnumeratingWithStateObjectsCount(state unsafe.Pointer, buffer unsafe.Pointer, len uint) uint
	DescriptionWithLocale(locale objc.ID) unsafe.Pointer
	DescriptionWithLocaleIndent(locale objc.ID, level uint) unsafe.Pointer
	EnumerateKeysAndObjectsUsingBlock(block unsafe.Pointer)
	EnumerateKeysAndObjectsWithOptionsUsingBlock(opts unsafe.Pointer, block unsafe.Pointer)
	FileCreationDate() unsafe.Pointer
	FileExtensionHidden() bool
	FileGroupOwnerAccountID() unsafe.Pointer
	FileGroupOwnerAccountName() unsafe.Pointer
	FileHFSCreatorCode() unsafe.Pointer
	FileHFSTypeCode() unsafe.Pointer
	FileIsAppendOnly() bool
	FileIsImmutable() bool
	FileModificationDate() unsafe.Pointer
	FileOwnerAccountID() unsafe.Pointer
	FileOwnerAccountName() unsafe.Pointer
	FilePosixPermissions() uint
	FileSize() unsafe.Pointer
	FileSystemFileNumber() uint
	FileSystemNumber() int
	FileType() unsafe.Pointer
	GetObjectsAndKeys(objects unsafe.Pointer, keys unsafe.Pointer)
	GetObjectsAndKeysCount(objects unsafe.Pointer, keys unsafe.Pointer, count uint)
	IsEqualToDictionary(otherDictionary unsafe.Pointer) bool
	KeyEnumerator() unsafe.Pointer
	KeysOfEntriesWithOptionsPassingTest(opts unsafe.Pointer, predicate unsafe.Pointer) unsafe.Pointer
	KeysOfEntriesPassingTest(predicate unsafe.Pointer) unsafe.Pointer
	KeysSortedByValueUsingComparator(cmptr unsafe.Pointer) []objc.ID
	KeysSortedByValueWithOptionsUsingComparator(opts unsafe.Pointer, cmptr unsafe.Pointer) []objc.ID
	KeysSortedByValueUsingSelector(comparator objc.SEL) []objc.ID
	ObjectForKey(aKey unsafe.Pointer) unsafe.Pointer
	ObjectEnumerator() unsafe.Pointer
	ObjectsForKeysNotFoundMarker(keys unsafe.Pointer, marker unsafe.Pointer) []objc.ID
	ObjectForKeyedSubscript(key unsafe.Pointer) unsafe.Pointer
	WriteToURLError(url unsafe.Pointer, error unsafe.Pointer) bool
	WriteToURLAtomically(url unsafe.Pointer, atomically bool) bool
	WriteToFileAtomically(path string, useAuxiliaryFile bool) bool
}

// A static collection of objects associated with unique keys.
//
// You can use this type in Swift instead of a in cases that require reference semantics. The class declares the programmatic interface to objects that manage immutable associations of keys and values. For example, an interactive form could be represented as a dictionary, with the field names as keys, corresponding to user-entered values. Use this class or its subclass when you need a convenient and efficient way to retrieve data associated with an arbitrary key. creates static dictionaries, and creates dynamic dictionaries. (For convenience, the term refers to any instance of one of these classes without specifying its exact class membership.) A key-value pair within a dictionary is called an entry. Each entry consists of one object that represents the key and a second object that is that key’s value. Within a dictionary, the keys are unique. That is, no two keys in a single dictionary are equal (as determined by ). In general, a key can be any object (provided that it conforms to the protocol—see below), but note that when using key-value coding the key must be a string (see ). Neither a key nor a value can be ; if you need to represent a null value in a dictionary, you should use . is “toll-free bridged” with its Core Foundation counterpart, . See for more information on toll-free bridging.
//
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(objects:forKeys:)
func NewDictionaryWithObjectsForKeys(objects unsafe.Pointer, keys unsafe.Pointer) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithObjects:forKeys:"), objects, keys)
	rv.Autorelease()
	return rv
}

// Initializes a newly allocated dictionary with the specified number of key-value pairs constructed from the provided C arrays of keys and objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(objects:forKeys:count:)
func NewDictionaryWithObjectsForKeysCount(objects unsafe.Pointer, keys objc.ID, cnt uint) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithObjects:forKeys:count:"), objects, keys, cnt)
	rv.Autorelease()
	return rv
}

// Initializes a newly allocated dictionary with entries constructed from the specified set of values and keys.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/initWithObjectsAndKeys:
func NewDictionaryWithObjectsAndKeys(firstObject objc.ID) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithObjectsAndKeys:"), firstObject)
	rv.Autorelease()
	return rv
}

// Initializes a newly allocated dictionary using the keys and values found at a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(contentsOfURL:)-4pv16
func NewDictionaryWithContentsOfURL(url unsafe.Pointer) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}

// Initializes a newly allocated dictionary using the keys and values found at a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(contentsOfURL:error:)
func NewDictionaryWithContentsOfURLError(url unsafe.Pointer, error unsafe.Pointer) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithContentsOfURL:error:"), url, error)
	rv.Autorelease()
	return rv
}

// Creates a dictionary containing a given key and value.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(object:forKey:)
func NewDictionaryWithObjectForKey(object unsafe.Pointer, key objc.ID) Dictionary {
	rv := objc.Send[Dictionary](objc.ID(getDictionaryClass().class), objc.Sel("dictionaryWithObject:forKey:"), object, key)
	return rv
}

// Creates a dictionary initialized from data in the provided unarchiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(coder:)
func NewDictionaryWithCoder(coder unsafe.Pointer) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}

// Initializes a newly allocated dictionary using the keys and values found in a file at a given path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(contentsOfFile:)
func NewDictionaryWithContentsOfFile(path string) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithContentsOfFile:"), objc.String(path))
	rv.Autorelease()
	return rv
}

// Initializes a newly allocated dictionary by placing in it the keys and values contained in another given dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(dictionary:)-9fw1u
func NewDictionaryWithDictionary(otherDictionary unsafe.Pointer) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithDictionary:"), otherDictionary)
	rv.Autorelease()
	return rv
}

// Initializes a newly allocated dictionary using the objects contained in another given dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(dictionary:copyItems:)
func NewDictionaryWithDictionaryCopyItems(otherDictionary unsafe.Pointer, flag bool) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithDictionary:copyItems:"), otherDictionary, flag)
	rv.Autorelease()
	return rv
}


// Creates an empty dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionary
func (dc _DictionaryClass) Dictionary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dictionary"))
	return rv
}

// Creates a dictionary using the keys and values found in a file specified by a given path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithContentsOfFile:
func (dc _DictionaryClass) DictionaryWithContentsOfFile(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dictionaryWithContentsOfFile:"), objc.String(path))
	return rv
}

// Creates a dictionary using the keys and values found in a resource specified by a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithContentsOfURL:error:
func (dc _DictionaryClass) DictionaryWithContentsOfURLError(url unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dictionaryWithContentsOfURL:error:"), url, error)
	return rv
}

// Creates a dictionary containing the keys and values from another given dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithDictionary:
func (dc _DictionaryClass) DictionaryWithDictionary(dict unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dictionaryWithDictionary:"), dict)
	return rv
}

// Creates a dictionary containing entries constructed from the contents of an array of keys and an array of values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithObjects:forKeys:
func (dc _DictionaryClass) DictionaryWithObjectsForKeys(objects unsafe.Pointer, keys unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dictionaryWithObjects:forKeys:"), objects, keys)
	return rv
}

// Creates a dictionary containing a specified number of objects from a C array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithObjects:forKeys:count:
func (dc _DictionaryClass) DictionaryWithObjectsForKeysCount(objects unsafe.Pointer, keys objc.ID, cnt uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dictionaryWithObjects:forKeys:count:"), objects, keys, cnt)
	return rv
}

// Creates a dictionary containing entries constructed from the specified set of values and keys.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithObjectsAndKeys:
func (dc _DictionaryClass) DictionaryWithObjectsAndKeys(firstObject objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dictionaryWithObjectsAndKeys:"), firstObject)
	return rv
}

// Creates a dictionary using the keys and values found in a resource specified by a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(contentsOfURL:)-98pl3
func (dc _DictionaryClass) DictionaryWithContentsOfURL(url unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dictionaryWithContentsOfURL:"), url)
	return rv
}

// Creates a dictionary containing a given key and value.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(object:forKey:)
func (dc _DictionaryClass) DictionaryWithObjectForKey(object unsafe.Pointer, key objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("dictionaryWithObject:forKey:"), object, key)
	return rv
}

// Creates a shared key set object for the specified keys.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/sharedKeySet(forKeys:)
func (dc _DictionaryClass) SharedKeySetForKeys(keys unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(dc.class), objc.Sel("sharedKeySetForKeys:"), keys)
	return rv
}

// Returns a new array containing the keys corresponding to all occurrences of a given object in the dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/allKeys(for:)
func (d_ Dictionary) AllKeysForObject(anObject unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("allKeysForObject:"), anObject)
	return rv
}

// Returns by reference a C array of objects over which the sender should iterate.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/countByEnumeratingWithState:objects:count:
func (d_ Dictionary) CountByEnumeratingWithStateObjectsCount(state unsafe.Pointer, buffer unsafe.Pointer, len uint) uint {
	rv := objc.Send[uint](d_.ID, objc.Sel("countByEnumeratingWithState:objects:count:"), state, buffer, len)
	return rv
}

// Returns a string object that represents the contents of the dictionary, formatted as a property list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/description(withLocale:)
func (d_ Dictionary) DescriptionWithLocale(locale objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}

// Returns a string object that represents the contents of the dictionary, formatted as a property list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/description(withLocale:indent:)
func (d_ Dictionary) DescriptionWithLocaleIndent(locale objc.ID, level uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("descriptionWithLocale:indent:"), locale, level)
	return rv
}

// Applies a given block object to the entries of the dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/enumerateKeysAndObjects(_:)
func (d_ Dictionary) EnumerateKeysAndObjectsUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("enumerateKeysAndObjectsUsingBlock:"), block)
}

// Applies a given block object to the entries of the dictionary, with options specifying how the enumeration is performed.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/enumerateKeysAndObjects(options:using:)
func (d_ Dictionary) EnumerateKeysAndObjectsWithOptionsUsingBlock(opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("enumerateKeysAndObjectsWithOptions:usingBlock:"), opts, block)
}

// Returns the file’s creation date.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileCreationDate()
func (d_ Dictionary) FileCreationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileCreationDate"))
	return rv
}

// Returns a Boolean value indicating whether the file hides its extension.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileExtensionHidden()
func (d_ Dictionary) FileExtensionHidden() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("fileExtensionHidden"))
	return rv
}

// Returns file’s group owner account ID.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileGroupOwnerAccountID()
func (d_ Dictionary) FileGroupOwnerAccountID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileGroupOwnerAccountID"))
	return rv
}

// Returns the file’s group owner account name.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileGroupOwnerAccountName()
func (d_ Dictionary) FileGroupOwnerAccountName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileGroupOwnerAccountName"))
	return rv
}

// Returns the file’s HFS creator code.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileHFSCreatorCode()
func (d_ Dictionary) FileHFSCreatorCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileHFSCreatorCode"))
	return rv
}

// Returns file’s HFS type code.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileHFSTypeCode()
func (d_ Dictionary) FileHFSTypeCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileHFSTypeCode"))
	return rv
}

// Returns a Boolean value indicating whether the file is append only.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileIsAppendOnly()
func (d_ Dictionary) FileIsAppendOnly() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("fileIsAppendOnly"))
	return rv
}

// Returns a Boolean value indicating whether the file is immutable.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileIsImmutable()
func (d_ Dictionary) FileIsImmutable() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("fileIsImmutable"))
	return rv
}

// Returns file’s modification date.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileModificationDate()
func (d_ Dictionary) FileModificationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileModificationDate"))
	return rv
}

// Returns the file’s owner account ID.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileOwnerAccountID()
func (d_ Dictionary) FileOwnerAccountID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileOwnerAccountID"))
	return rv
}

// Returns the file’s owner account name.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileOwnerAccountName()
func (d_ Dictionary) FileOwnerAccountName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileOwnerAccountName"))
	return rv
}

// Returns the file’s POSIX permissions.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/filePosixPermissions()
func (d_ Dictionary) FilePosixPermissions() uint {
	rv := objc.Send[uint](d_.ID, objc.Sel("filePosixPermissions"))
	return rv
}

// Returns the file’s size, in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileSize()
func (d_ Dictionary) FileSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileSize"))
	return rv
}

// Returns the filesystem file number.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileSystemFileNumber()
func (d_ Dictionary) FileSystemFileNumber() uint {
	rv := objc.Send[uint](d_.ID, objc.Sel("fileSystemFileNumber"))
	return rv
}

// Returns the filesystem number.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileSystemNumber()
func (d_ Dictionary) FileSystemNumber() int {
	rv := objc.Send[int](d_.ID, objc.Sel("fileSystemNumber"))
	return rv
}

// Returns the file type.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileType()
func (d_ Dictionary) FileType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileType"))
	return rv
}

// Returns by reference C arrays of the keys and values in the dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/getObjects:andKeys:
func (d_ Dictionary) GetObjectsAndKeys(objects unsafe.Pointer, keys unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("getObjects:andKeys:"), objects, keys)
}

// Returns by reference C arrays of the keys and values in the dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/getObjects:andKeys:count:
func (d_ Dictionary) GetObjectsAndKeysCount(objects unsafe.Pointer, keys unsafe.Pointer, count uint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("getObjects:andKeys:count:"), objects, keys, count)
}

// Returns a Boolean value that indicates whether the contents of the receiving dictionary are equal to the contents of another given dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/isEqual(to:)
func (d_ Dictionary) IsEqualToDictionary(otherDictionary unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isEqualToDictionary:"), otherDictionary)
	return rv
}

// Provides an enumerator to access the keys in the dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/keyEnumerator()
func (d_ Dictionary) KeyEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("keyEnumerator"))
	return rv
}

// Returns the set of keys whose corresponding value satisfies a constraint described by a block object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/keysOfEntries(options:passingTest:)
func (d_ Dictionary) KeysOfEntriesWithOptionsPassingTest(opts unsafe.Pointer, predicate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("keysOfEntriesWithOptions:passingTest:"), opts, predicate)
	return rv
}

// Returns the set of keys whose corresponding value satisfies a constraint described by a block object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/keysOfEntries(passingTest:)
func (d_ Dictionary) KeysOfEntriesPassingTest(predicate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("keysOfEntriesPassingTest:"), predicate)
	return rv
}

// Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values using a given comparator block.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/keysSortedByValue(comparator:)
func (d_ Dictionary) KeysSortedByValueUsingComparator(cmptr unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("keysSortedByValueUsingComparator:"), cmptr)
	return rv
}

// Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values using a given comparator block and a specified set of options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/keysSortedByValue(options:usingComparator:)
func (d_ Dictionary) KeysSortedByValueWithOptionsUsingComparator(opts unsafe.Pointer, cmptr unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("keysSortedByValueWithOptions:usingComparator:"), opts, cmptr)
	return rv
}

// Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/keysSortedByValue(using:)
func (d_ Dictionary) KeysSortedByValueUsingSelector(comparator objc.SEL) []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("keysSortedByValueUsingSelector:"), comparator)
	return rv
}

// Returns the value associated with a given key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/object(forKey:)
func (d_ Dictionary) ObjectForKey(aKey unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("objectForKey:"), aKey)
	return rv
}

// Returns an enumerator object that lets you access each value in the dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/objectEnumerator()
func (d_ Dictionary) ObjectEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("objectEnumerator"))
	return rv
}

// Returns as a static array the set of objects from the dictionary that corresponds to the specified keys.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/objects(forKeys:notFoundMarker:)
func (d_ Dictionary) ObjectsForKeysNotFoundMarker(keys unsafe.Pointer, marker unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("objectsForKeys:notFoundMarker:"), keys, marker)
	return rv
}

// Returns the value associated with a given key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/subscript(_:)-52n56
func (d_ Dictionary) ObjectForKeyedSubscript(key unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}

// Returns the value associated with a given key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/value(forKey:)
func (d_ Dictionary) ValueForKey(key string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("valueForKey:"), objc.String(key))
	return rv
}

// Writes a property list representation of the contents of the dictionary to a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/write(to:)
func (d_ Dictionary) WriteToURLError(url unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeToURL:error:"), url, error)
	return rv
}

// Writes a property list representation of the contents of the dictionary to a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/write(to:atomically:)
func (d_ Dictionary) WriteToURLAtomically(url unsafe.Pointer, atomically bool) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeToURL:atomically:"), url, atomically)
	return rv
}

// Writes a property list representation of the contents of the dictionary to a given path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/write(toFile:atomically:)
func (d_ Dictionary) WriteToFileAtomically(path string, useAuxiliaryFile bool) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeToFile:atomically:"), objc.String(path), useAuxiliaryFile)
	return rv
}

// A new array containing the dictionary’s keys, or an empty array if the dictionary has no entries.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/allKeys
func (d_ Dictionary) AllKeys() []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("allKeys"))
	return rv
}

// A new array containing the dictionary’s values, or an empty array if the dictionary has no entries.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/allValues
func (d_ Dictionary) AllValues() []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("allValues"))
	return rv
}

// The number of entries in the dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/count
func (d_ Dictionary) Count() uint {
	rv := objc.Send[uint](d_.ID, objc.Sel("count"))
	return rv
}

// A string that represents the contents of the dictionary, formatted as a property list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/description
func (d_ Dictionary) Description() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("description"))
	return rv
}

// A string that represents the contents of the dictionary, formatted in file format.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/descriptionInStringsFileFormat
func (d_ Dictionary) DescriptionInStringsFileFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("descriptionInStringsFileFormat"))
	return rv
}



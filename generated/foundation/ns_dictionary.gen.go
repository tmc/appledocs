// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSDictionary */


/* debug [class_header]: Header for NSDictionary */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Dictionary */
// An interface definition for the [Dictionary] class.
type IDictionary interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Dictionary */
	// properties:
	AllKeys() []objc.ID
	AllValues() []objc.ID
	Description() IString
	DescriptionInStringsFileFormat() IString
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Dictionary */
	// methods:
	AllKeysForObject(anObject objectivec.IObject) []objc.ID
	CountByEnumeratingWithStateObjectsCount(state objc.IObject /* cross-framework: FastEnumerationState */, buffer []objc.ID, len_ uint) uint
	DescriptionWithLocale(locale objc.IObject) IString
	DescriptionWithLocaleIndent(locale objc.IObject, level uint) IString
	EnumerateKeysAndObjectsUsingBlock(block unsafe.Pointer)
	EnumerateKeysAndObjectsWithOptionsUsingBlock(opts EnumerationOptions, block unsafe.Pointer)
	FileCreationDate() IDate
	FileExtensionHidden() bool
	FileGroupOwnerAccountID() INumber
	FileGroupOwnerAccountName() IString
	FileHFSCreatorCode() uint32 /* not a class type */
	FileHFSTypeCode() uint32 /* not a class type */
	FileIsAppendOnly() bool
	FileIsImmutable() bool
	FileModificationDate() IDate
	FileOwnerAccountID() INumber
	FileOwnerAccountName() IString
	FilePosixPermissions() uint
	FileSize() uint64
	FileSystemFileNumber() uint
	FileSystemNumber() int
	FileType() IString
	GetObjectsAndKeysCount(objects []objc.ID, keys []objc.ID, count uint)
	IsEqualToDictionary(otherDictionary IDictionary) bool
	KeyEnumerator() unsafe.Pointer
	KeysOfEntriesWithOptionsPassingTest(opts EnumerationOptions, predicate unsafe.Pointer) unsafe.Pointer
	KeysOfEntriesPassingTest(predicate unsafe.Pointer) unsafe.Pointer
	KeysSortedByValueUsingComparator(cmptr Comparator /* not a class type */) []objc.ID
	KeysSortedByValueWithOptionsUsingComparator(opts SortOptions, cmptr Comparator /* not a class type */) []objc.ID
	KeysSortedByValueUsingSelector(comparator objc.SEL) []objc.ID
	ObjectForKey(aKey objectivec.IObject) objectivec.IObject
	ObjectEnumerator() unsafe.Pointer
	ObjectsForKeysNotFoundMarker(keys []objc.ID, marker objectivec.IObject) []objc.ID
	ObjectForKeyedSubscript(key objectivec.IObject) objectivec.IObject
	WriteToURLError(url IURL, error_ IError) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Dictionary */
// Alloc allocates a new instance without initialization.
func (dc _DictionaryClass) Alloc() Dictionary {
	rv := objc.Send[Dictionary](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Dictionary */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Dictionary */

// Creates a dictionary initialized from data in the provided unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(coder:)
func NewDictionaryWithCoder(coder ICoder) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDictionaryWithCoder */


// Initializes a newly allocated dictionary using the keys and values found in a file at a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(contentsOfFile:)
func NewDictionaryWithContentsOfFile(path IString) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithContentsOfFile:"), path)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDictionaryWithContentsOfFile */


// Initializes a newly allocated dictionary using the keys and values found at a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(contentsOfURL:)-4pv16
func NewDictionaryWithContentsOfURL(url IURL) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDictionaryWithContentsOfURL */


// Initializes a newly allocated dictionary using the keys and values found at a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(contentsOfURL:error:)
func NewDictionaryWithContentsOfURLError(url IURL, error_ IError) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithContentsOfURL:error:"), url, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDictionaryWithContentsOfURLError */


// Initializes a newly allocated dictionary by placing in it the keys and values contained in another given dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(dictionary:)-9fw1u
func NewDictionaryWithDictionary(otherDictionary IDictionary) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithDictionary:"), otherDictionary)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDictionaryWithDictionary */


// Initializes a newly allocated dictionary using the objects contained in another given dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(dictionary:copyItems:)
func NewDictionaryWithDictionaryCopyItems(otherDictionary IDictionary, flag bool) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithDictionary:copyItems:"), otherDictionary, flag)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDictionaryWithDictionaryCopyItems */


// Creates a dictionary containing a given key and value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(object:forKey:)
func NewDictionaryWithObjectForKey(object objectivec.IObject, key unsafe.Pointer) Dictionary {
	rv := objc.Send[Dictionary](objc.ID(getDictionaryClass().class), objc.Sel("dictionaryWithObject:forKey:"), object, key)
	return rv
}/* debug [class_init_methods/constructor]: NewDictionaryWithObjectForKey */


// Initializes a newly allocated dictionary with entries constructed from the specified set of values and keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/initWithObjectsAndKeys:
func NewDictionaryWithObjectsAndKeys(firstObject objc.IObject) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithObjectsAndKeys:"), firstObject)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDictionaryWithObjectsAndKeys */


// Initializes a newly allocated dictionary with key-value pairs constructed from the provided arrays of keys and objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(objects:forKeys:)
func NewDictionaryWithObjectsForKeys(objects []objc.ID, keys []objc.ID) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithObjects:forKeys:"), objects, keys)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDictionaryWithObjectsForKeys */


// Initializes a newly allocated dictionary with the specified number of key-value pairs constructed from the provided C arrays of keys and objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(objects:forKeys:count:)
func NewDictionaryWithObjectsForKeysCount(objects []objc.ID, keys []objc.ID, cnt uint) Dictionary {
	instance := getDictionaryClass().Alloc()
	rv := objc.Send[Dictionary](instance.ID, objc.Sel("initWithObjects:forKeys:count:"), objects, keys, cnt)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDictionaryWithObjectsForKeysCount */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Dictionary */

// Creates an empty dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionary
func (dc _DictionaryClass) Dictionary() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("dictionary"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Dictionary) */


// Creates a dictionary using the keys and values found in a file specified by a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithContentsOfFile:
func (dc _DictionaryClass) DictionaryWithContentsOfFile(path IString) IDictionary {
	rv := objc.Send[Dictionary](objc.ID(dc.class), objc.Sel("dictionaryWithContentsOfFile:"), path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DictionaryWithContentsOfFile) */


// Creates a dictionary using the keys and values found in a resource specified by a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithContentsOfURL:error:
func (dc _DictionaryClass) DictionaryWithContentsOfURLError(url IURL, error_ IError) IDictionary {
	rv := objc.Send[Dictionary](objc.ID(dc.class), objc.Sel("dictionaryWithContentsOfURL:error:"), url, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DictionaryWithContentsOfURLError) */


// Creates a dictionary containing the keys and values from another given dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithDictionary:
func (dc _DictionaryClass) DictionaryWithDictionary(dict IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("dictionaryWithDictionary:"), dict)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DictionaryWithDictionary) */


// Creates a dictionary containing entries constructed from the contents of an array of keys and an array of values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithObjects:forKeys:
func (dc _DictionaryClass) DictionaryWithObjectsForKeys(objects []objc.ID, keys []objc.ID) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("dictionaryWithObjects:forKeys:"), objects, keys)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DictionaryWithObjectsForKeys) */


// Creates a dictionary containing a specified number of objects from a C array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithObjects:forKeys:count:
func (dc _DictionaryClass) DictionaryWithObjectsForKeysCount(objects []objc.ID, keys []objc.ID, cnt uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("dictionaryWithObjects:forKeys:count:"), objects, keys, cnt)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DictionaryWithObjectsForKeysCount) */


// Creates a dictionary containing entries constructed from the specified set of values and keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/dictionaryWithObjectsAndKeys:
func (dc _DictionaryClass) DictionaryWithObjectsAndKeys(firstObject objc.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("dictionaryWithObjectsAndKeys:"), firstObject)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DictionaryWithObjectsAndKeys) */


// Creates a dictionary using the keys and values found in a resource specified by a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(contentsOfURL:)-98pl3
func (dc _DictionaryClass) DictionaryWithContentsOfURL(url IURL) IDictionary {
	rv := objc.Send[Dictionary](objc.ID(dc.class), objc.Sel("dictionaryWithContentsOfURL:"), url)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DictionaryWithContentsOfURL) */


// Creates a dictionary containing a given key and value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/init(object:forKey:)
func (dc _DictionaryClass) DictionaryWithObjectForKey(object objectivec.IObject, key unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("dictionaryWithObject:forKey:"), object, key)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DictionaryWithObjectForKey) */


// Creates a shared key set object for the specified keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/sharedKeySet(forKeys:)
func (dc _DictionaryClass) SharedKeySetForKeys(keys []objc.ID) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(dc.class), objc.Sel("sharedKeySetForKeys:"), keys)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedKeySetForKeys) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Dictionary */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Dictionary */

// Returns a new array containing the keys corresponding to all occurrences of a given object in the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/allKeys(for:)
func (d_ Dictionary) AllKeysForObject(anObject objectivec.IObject) []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("allKeysForObject:"), anObject)
	return rv
}/* debug [instance_methods/method]: AllKeysForObject */


// Returns by reference a C array of objects over which the sender should iterate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/countByEnumeratingWithState:objects:count:
func (d_ Dictionary) CountByEnumeratingWithStateObjectsCount(state objc.IObject /* cross-framework: FastEnumerationState */, buffer []objc.ID, len_ uint) uint {
	rv := objc.Send[uint](d_.ID, objc.Sel("countByEnumeratingWithState:objects:count:"), state, buffer, len_)
	return rv
}/* debug [instance_methods/method]: CountByEnumeratingWithStateObjectsCount */


// Returns a string object that represents the contents of the dictionary, formatted as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/description(withLocale:)
func (d_ Dictionary) DescriptionWithLocale(locale objc.IObject) IString {
	rv := objc.Send[String](d_.ID, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}/* debug [instance_methods/method]: DescriptionWithLocale */


// Returns a string object that represents the contents of the dictionary, formatted as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/description(withLocale:indent:)
func (d_ Dictionary) DescriptionWithLocaleIndent(locale objc.IObject, level uint) IString {
	rv := objc.Send[String](d_.ID, objc.Sel("descriptionWithLocale:indent:"), locale, level)
	return rv
}/* debug [instance_methods/method]: DescriptionWithLocaleIndent */


// Applies a given block object to the entries of the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/enumerateKeysAndObjects(_:)
func (d_ Dictionary) EnumerateKeysAndObjectsUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("enumerateKeysAndObjectsUsingBlock:"), block)
}/* debug [instance_methods/method]: EnumerateKeysAndObjectsUsingBlock */


// Applies a given block object to the entries of the dictionary, with options specifying how the enumeration is performed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/enumerateKeysAndObjects(options:using:)
func (d_ Dictionary) EnumerateKeysAndObjectsWithOptionsUsingBlock(opts EnumerationOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("enumerateKeysAndObjectsWithOptions:usingBlock:"), opts, block)
}/* debug [instance_methods/method]: EnumerateKeysAndObjectsWithOptionsUsingBlock */


// Returns the file’s creation date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileCreationDate()
func (d_ Dictionary) FileCreationDate() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("fileCreationDate"))
	return rv
}/* debug [instance_methods/method]: FileCreationDate */


// Returns a Boolean value indicating whether the file hides its extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileExtensionHidden()
func (d_ Dictionary) FileExtensionHidden() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("fileExtensionHidden"))
	return rv
}/* debug [instance_methods/method]: FileExtensionHidden */


// Returns file’s group owner account ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileGroupOwnerAccountID()
func (d_ Dictionary) FileGroupOwnerAccountID() INumber {
	rv := objc.Send[Number](d_.ID, objc.Sel("fileGroupOwnerAccountID"))
	return rv
}/* debug [instance_methods/method]: FileGroupOwnerAccountID */


// Returns the file’s group owner account name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileGroupOwnerAccountName()
func (d_ Dictionary) FileGroupOwnerAccountName() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("fileGroupOwnerAccountName"))
	return rv
}/* debug [instance_methods/method]: FileGroupOwnerAccountName */


// Returns the file’s HFS creator code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileHFSCreatorCode()
func (d_ Dictionary) FileHFSCreatorCode() uint32 /* not a class type */ {
	rv := objc.Send[uint32](d_.ID, objc.Sel("fileHFSCreatorCode"))
	return rv
}/* debug [instance_methods/method]: FileHFSCreatorCode */


// Returns file’s HFS type code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileHFSTypeCode()
func (d_ Dictionary) FileHFSTypeCode() uint32 /* not a class type */ {
	rv := objc.Send[uint32](d_.ID, objc.Sel("fileHFSTypeCode"))
	return rv
}/* debug [instance_methods/method]: FileHFSTypeCode */


// Returns a Boolean value indicating whether the file is append only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileIsAppendOnly()
func (d_ Dictionary) FileIsAppendOnly() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("fileIsAppendOnly"))
	return rv
}/* debug [instance_methods/method]: FileIsAppendOnly */


// Returns a Boolean value indicating whether the file is immutable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileIsImmutable()
func (d_ Dictionary) FileIsImmutable() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("fileIsImmutable"))
	return rv
}/* debug [instance_methods/method]: FileIsImmutable */


// Returns file’s modification date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileModificationDate()
func (d_ Dictionary) FileModificationDate() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("fileModificationDate"))
	return rv
}/* debug [instance_methods/method]: FileModificationDate */


// Returns the file’s owner account ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileOwnerAccountID()
func (d_ Dictionary) FileOwnerAccountID() INumber {
	rv := objc.Send[Number](d_.ID, objc.Sel("fileOwnerAccountID"))
	return rv
}/* debug [instance_methods/method]: FileOwnerAccountID */


// Returns the file’s owner account name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileOwnerAccountName()
func (d_ Dictionary) FileOwnerAccountName() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("fileOwnerAccountName"))
	return rv
}/* debug [instance_methods/method]: FileOwnerAccountName */


// Returns the file’s POSIX permissions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/filePosixPermissions()
func (d_ Dictionary) FilePosixPermissions() uint {
	rv := objc.Send[uint](d_.ID, objc.Sel("filePosixPermissions"))
	return rv
}/* debug [instance_methods/method]: FilePosixPermissions */


// Returns the file’s size, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileSize()
func (d_ Dictionary) FileSize() uint64 {
	rv := objc.Send[uint64](d_.ID, objc.Sel("fileSize"))
	return rv
}/* debug [instance_methods/method]: FileSize */


// Returns the filesystem file number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileSystemFileNumber()
func (d_ Dictionary) FileSystemFileNumber() uint {
	rv := objc.Send[uint](d_.ID, objc.Sel("fileSystemFileNumber"))
	return rv
}/* debug [instance_methods/method]: FileSystemFileNumber */


// Returns the filesystem number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileSystemNumber()
func (d_ Dictionary) FileSystemNumber() int {
	rv := objc.Send[int](d_.ID, objc.Sel("fileSystemNumber"))
	return rv
}/* debug [instance_methods/method]: FileSystemNumber */


// Returns the file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/fileType()
func (d_ Dictionary) FileType() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("fileType"))
	return rv
}/* debug [instance_methods/method]: FileType */


// Returns by reference C arrays of the keys and values in the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/getObjects:andKeys:count:
func (d_ Dictionary) GetObjectsAndKeysCount(objects []objc.ID, keys []objc.ID, count uint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("getObjects:andKeys:count:"), objects, keys, count)
}/* debug [instance_methods/method]: GetObjectsAndKeysCount */


// Returns a Boolean value that indicates whether the contents of the receiving dictionary are equal to the contents of another given dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/isEqual(to:)
func (d_ Dictionary) IsEqualToDictionary(otherDictionary IDictionary) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isEqualToDictionary:"), otherDictionary)
	return rv
}/* debug [instance_methods/method]: IsEqualToDictionary */


// Provides an enumerator to access the keys in the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/keyEnumerator()
func (d_ Dictionary) KeyEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("keyEnumerator"))
	return rv
}/* debug [instance_methods/method]: KeyEnumerator */


// Returns the set of keys whose corresponding value satisfies a constraint described by a block object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/keysOfEntries(options:passingTest:)
func (d_ Dictionary) KeysOfEntriesWithOptionsPassingTest(opts EnumerationOptions, predicate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("keysOfEntriesWithOptions:passingTest:"), opts, predicate)
	return rv
}/* debug [instance_methods/method]: KeysOfEntriesWithOptionsPassingTest */


// Returns the set of keys whose corresponding value satisfies a constraint described by a block object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/keysOfEntries(passingTest:)
func (d_ Dictionary) KeysOfEntriesPassingTest(predicate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("keysOfEntriesPassingTest:"), predicate)
	return rv
}/* debug [instance_methods/method]: KeysOfEntriesPassingTest */


// Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values using a given comparator block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/keysSortedByValue(comparator:)
func (d_ Dictionary) KeysSortedByValueUsingComparator(cmptr Comparator /* not a class type */) []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("keysSortedByValueUsingComparator:"), cmptr)
	return rv
}/* debug [instance_methods/method]: KeysSortedByValueUsingComparator */


// Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values using a given comparator block and a specified set of options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/keysSortedByValue(options:usingComparator:)
func (d_ Dictionary) KeysSortedByValueWithOptionsUsingComparator(opts SortOptions, cmptr Comparator /* not a class type */) []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("keysSortedByValueWithOptions:usingComparator:"), opts, cmptr)
	return rv
}/* debug [instance_methods/method]: KeysSortedByValueWithOptionsUsingComparator */


// Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/keysSortedByValue(using:)
func (d_ Dictionary) KeysSortedByValueUsingSelector(comparator objc.SEL) []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("keysSortedByValueUsingSelector:"), comparator)
	return rv
}/* debug [instance_methods/method]: KeysSortedByValueUsingSelector */


// Returns the value associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/object(forKey:)
func (d_ Dictionary) ObjectForKey(aKey objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("objectForKey:"), aKey)
	return rv
}/* debug [instance_methods/method]: ObjectForKey */


// Returns an enumerator object that lets you access each value in the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/objectEnumerator()
func (d_ Dictionary) ObjectEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("objectEnumerator"))
	return rv
}/* debug [instance_methods/method]: ObjectEnumerator */


// Returns as a static array the set of objects from the dictionary that corresponds to the specified keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/objects(forKeys:notFoundMarker:)
func (d_ Dictionary) ObjectsForKeysNotFoundMarker(keys []objc.ID, marker objectivec.IObject) []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("objectsForKeys:notFoundMarker:"), keys, marker)
	return rv
}/* debug [instance_methods/method]: ObjectsForKeysNotFoundMarker */


// Returns the value associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/subscript(_:)-52n56
func (d_ Dictionary) ObjectForKeyedSubscript(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}/* debug [instance_methods/method]: ObjectForKeyedSubscript */


// Returns the value associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/value(forKey:)
func (d_ Dictionary) ValueForKey(key IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("valueForKey:"), key)
	return rv
}/* debug [instance_methods/method]: ValueForKey */


// Writes a property list representation of the contents of the dictionary to a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/write(to:)
func (d_ Dictionary) WriteToURLError(url IURL, error_ IError) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeToURL:error:"), url, error_)
	return rv
}/* debug [instance_methods/method]: WriteToURLError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Dictionary */

// A new array containing the dictionary’s keys, or an empty array if the dictionary has no entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/allKeys
func (d_ Dictionary) AllKeys() []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("allKeys"))
	return rv
}/* debug [instance_properties/getter]: allKeys */


// A new array containing the dictionary’s values, or an empty array if the dictionary has no entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/allValues
func (d_ Dictionary) AllValues() []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("allValues"))
	return rv
}/* debug [instance_properties/getter]: allValues */


// A string that represents the contents of the dictionary, formatted as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/description
func (d_ Dictionary) Description() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("description"))
	return rv
}/* debug [instance_properties/getter]: description */


// A string that represents the contents of the dictionary, formatted in file format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDictionary/descriptionInStringsFileFormat
func (d_ Dictionary) DescriptionInStringsFileFormat() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("descriptionInStringsFileFormat"))
	return rv
}/* debug [instance_properties/getter]: descriptionInStringsFileFormat */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDictionary */



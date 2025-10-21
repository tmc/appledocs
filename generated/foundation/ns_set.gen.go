// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Set] class.
var (
	SetClass     _SetClass
	SetClassOnce sync.Once
)

func getSetClass() _SetClass {
	SetClassOnce.Do(func() {
		SetClass = _SetClass{objc.GetClass("NSSet")}
	})
	return SetClass
}

type _SetClass struct {
	class objc.Class
}

// An interface definition for the [Set] class.
type ISet interface {
	objectivec.IObject
	SetByAddingObject(anObject unsafe.Pointer) unsafe.Pointer
	SetByAddingObjectsFromSet(other unsafe.Pointer) unsafe.Pointer
	SetByAddingObjectsFromArray(other unsafe.Pointer) unsafe.Pointer
	AnyObject() unsafe.Pointer
	ContainsObject(anObject unsafe.Pointer) bool
	DescriptionWithLocale(locale objc.ID) string
	EnumerateIndexPathsWithOptionsUsingBlock(opts unsafe.Pointer, block unsafe.Pointer)
	EnumerateObjectsUsingBlock(block unsafe.Pointer)
	EnumerateObjectsWithOptionsUsingBlock(opts unsafe.Pointer, block unsafe.Pointer)
	FilteredSetUsingPredicate(predicate unsafe.Pointer) unsafe.Pointer
	IntersectsSet(otherSet unsafe.Pointer) bool
	IsEqualToSet(otherSet unsafe.Pointer) bool
	IsSubsetOfSet(otherSet unsafe.Pointer) bool
	MakeObjectsPerformSelector(aSelector objc.SEL)
	MakeObjectsPerformSelectorWithObject(aSelector objc.SEL, argument objc.ID)
	Member(object unsafe.Pointer) unsafe.Pointer
	ObjectEnumerator() unsafe.Pointer
	ObjectsWithOptionsPassingTest(opts unsafe.Pointer, predicate unsafe.Pointer) unsafe.Pointer
	ObjectsPassingTest(predicate unsafe.Pointer) unsafe.Pointer
	SortedArrayUsingDescriptors(sortDescriptors unsafe.Pointer) []objc.ID
}

// A static, unordered collection of unique objects.
//
// The , , and classes declare the programmatic interface to an unordered collection of objects. declares the programmatic interface for static sets of distinct objects. You establish a static set’s entries when it’s created, and can’t modify the entries after that. , on the other hand, declares a programmatic interface for dynamic sets of distinct objects. A dynamic — or mutable — set allows the addition and deletion of entries at any time, automatically allocating memory as needed. Use sets as an alternative to arrays when the order of elements isn’t important and you need to consider performance in testing whether the set contains an object. With an array, testing for membership is slower than with sets. is “toll-free bridged” with its Core Foundation counterpart, . See for more information on toll-free bridging. In Swift, use this class instead of a constant in cases where you require reference semantics.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet
type Set struct {
	objectivec.Object
}

// SetFrom constructs a [Set] from an unsafe.Pointer.
//
// A static, unordered collection of unique objects.
func SetFrom(ptr unsafe.Pointer) Set {
	return Set{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SetClass) Alloc() Set {
	rv := objc.Send[Set](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SetClass) New() Set {
	rv := objc.Send[Set](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Set) Init() Set {
	rv := objc.Send[Set](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Set) Autorelease() Set {
	rv := objc.Send[Set](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSet creates a new Set instance.
func NewSet() Set {
	return getSetClass().New()
}


// Initializes a newly allocated set with the objects that are contained in a given array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/init(array:)
func NewSetWithArray(array unsafe.Pointer) Set {
	instance := getSetClass().Alloc()
	rv := objc.Send[Set](instance.ID, objc.Sel("initWithArray:"), array)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/init(coder:)
func NewSetWithCoder(coder unsafe.Pointer) Set {
	instance := getSetClass().Alloc()
	rv := objc.Send[Set](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/init(collectionViewIndexPath:)
func NewSetWithCollectionViewIndexPath(indexPath unsafe.Pointer) Set {
	rv := objc.Send[Set](objc.ID(getSetClass().class), objc.Sel("setWithCollectionViewIndexPath:"), indexPath)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/init(collectionViewIndexPaths:)
func NewSetWithCollectionViewIndexPaths(indexPaths unsafe.Pointer) Set {
	rv := objc.Send[Set](objc.ID(getSetClass().class), objc.Sel("setWithCollectionViewIndexPaths:"), indexPaths)
	return rv
}

// Creates and returns a set that contains a single given object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/init(object:)
func NewSetWithObject(object unsafe.Pointer) Set {
	rv := objc.Send[Set](objc.ID(getSetClass().class), objc.Sel("setWithObject:"), object)
	return rv
}

// Initializes a newly allocated set with members taken from the specified list of objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/initWithObjects:
func NewSetWithObjects(firstObj unsafe.Pointer) Set {
	instance := getSetClass().Alloc()
	rv := objc.Send[Set](instance.ID, objc.Sel("initWithObjects:"), firstObj)
	rv.Autorelease()
	return rv
}

// Initializes a newly allocated set with a specified number of objects from a given C array of objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/init(objects:count:)-7kift
func NewSetWithObjectsCount(objects unsafe.Pointer, cnt uint) Set {
	instance := getSetClass().Alloc()
	rv := objc.Send[Set](instance.ID, objc.Sel("initWithObjects:count:"), objects, cnt)
	rv.Autorelease()
	return rv
}

// Initializes a newly allocated set and adds to it objects from another given set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/init(set:)-1xovx
func NewSetWithSet(set unsafe.Pointer) Set {
	instance := getSetClass().Alloc()
	rv := objc.Send[Set](instance.ID, objc.Sel("initWithSet:"), set)
	rv.Autorelease()
	return rv
}

// Initializes a newly allocated set and adds to it members of another given set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/init(set:copyItems:)
func NewSetWithSetCopyItems(set unsafe.Pointer, flag bool) Set {
	instance := getSetClass().Alloc()
	rv := objc.Send[Set](instance.ID, objc.Sel("initWithSet:copyItems:"), set, flag)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/init(collectionViewIndexPath:)
func (sc _SetClass) SetWithCollectionViewIndexPath(indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("setWithCollectionViewIndexPath:"), indexPath)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/init(collectionViewIndexPaths:)
func (sc _SetClass) SetWithCollectionViewIndexPaths(indexPaths unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("setWithCollectionViewIndexPaths:"), indexPaths)
	return rv
}

// Creates and returns a set that contains a single given object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/init(object:)
func (sc _SetClass) SetWithObject(object unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("setWithObject:"), object)
	return rv
}

// Creates and returns a set containing a specified number of objects from a given C array of objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/init(objects:count:)-65ni4
func (sc _SetClass) SetWithObjectsCount(objects unsafe.Pointer, cnt uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("setWithObjects:count:"), objects, cnt)
	return rv
}

// Creates and returns an empty set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/set
func (sc _SetClass) Set() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("set"))
	return rv
}

// Creates and returns a set containing a uniqued collection of the objects contained in a given array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/setWithArray:
func (sc _SetClass) SetWithArray(array unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("setWithArray:"), array)
	return rv
}

// Creates and returns a set containing the objects in a given argument list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/setWithObjects:
func (sc _SetClass) SetWithObjects(firstObj unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("setWithObjects:"), firstObj)
	return rv
}

// Creates and returns a set containing the objects from another set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/setWithSet:
func (sc _SetClass) SetWithSet(set unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("setWithSet:"), set)
	return rv
}

// Raises an exception.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/addObserver(_:forKeyPath:options:context:)
func (s_ Set) AddObserverForKeyPathOptionsContext(observer unsafe.Pointer, keyPath string, options unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addObserver:forKeyPath:options:context:"), observer, objc.String(keyPath), options, context)
}

// Returns a new set formed by adding a given object to the receiving set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/adding(_:)
func (s_ Set) SetByAddingObject(anObject unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("setByAddingObject:"), anObject)
	return rv
}

// Returns a new set formed by adding the objects in a given set to the receiving set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/addingObjects(from:)-2i31h
func (s_ Set) SetByAddingObjectsFromSet(other unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("setByAddingObjectsFromSet:"), other)
	return rv
}

// Returns a new set formed by adding the objects in a given array to the receiving set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/addingObjects(from:)-544m9
func (s_ Set) SetByAddingObjectsFromArray(other unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("setByAddingObjectsFromArray:"), other)
	return rv
}

// Returns one of the objects in the set, or if the set contains no objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/anyObject()
func (s_ Set) AnyObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("anyObject"))
	return rv
}

// Returns a Boolean value that indicates whether a given object is present in the set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/contains(_:)
func (s_ Set) ContainsObject(anObject unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("containsObject:"), anObject)
	return rv
}

// Returns a string that represents the contents of the set, formatted as a property list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/description(withLocale:)
func (s_ Set) DescriptionWithLocale(locale objc.ID) string {
	rv := objc.Send[string](s_.ID, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/enumerateIndexPaths(options:using:)
func (s_ Set) EnumerateIndexPathsWithOptionsUsingBlock(opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateIndexPathsWithOptions:usingBlock:"), opts, block)
}

// Executes a given block using each object in the set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/enumerateObjects(_:)
func (s_ Set) EnumerateObjectsUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateObjectsUsingBlock:"), block)
}

// Executes a given block using each object in the set, using the specified enumeration options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/enumerateObjects(options:using:)
func (s_ Set) EnumerateObjectsWithOptionsUsingBlock(opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateObjectsWithOptions:usingBlock:"), opts, block)
}

// Evaluates a given predicate against each object in the receiving set and returns a new set containing the objects for which the predicate returns true.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/filtered(using:)
func (s_ Set) FilteredSetUsingPredicate(predicate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("filteredSetUsingPredicate:"), predicate)
	return rv
}

// Returns a Boolean value that indicates whether at least one object in the receiving set is also present in another given set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/intersects(_:)
func (s_ Set) IntersectsSet(otherSet unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("intersectsSet:"), otherSet)
	return rv
}

// Compares the receiving set to another set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/isEqual(to:)
func (s_ Set) IsEqualToSet(otherSet unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEqualToSet:"), otherSet)
	return rv
}

// Returns a Boolean value that indicates whether every object in the receiving set is also present in another given set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/isSubset(of:)
func (s_ Set) IsSubsetOfSet(otherSet unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSubsetOfSet:"), otherSet)
	return rv
}

// Sends a message specified by a given selector to each object in the set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/makeObjectsPerformSelector:
func (s_ Set) MakeObjectsPerformSelector(aSelector objc.SEL) {
	objc.Send[objc.ID](s_.ID, objc.Sel("makeObjectsPerformSelector:"), aSelector)
}

// Sends a message specified by a given selector to each object in the set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/makeObjectsPerformSelector:withObject:
func (s_ Set) MakeObjectsPerformSelectorWithObject(aSelector objc.SEL, argument objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("makeObjectsPerformSelector:withObject:"), aSelector, argument)
}

// Determines whether a given object is present in the set, and returns that object if it is.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/member(_:)
func (s_ Set) Member(object unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("member:"), object)
	return rv
}

// Returns an enumerator object that lets you access each object in the set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/objectEnumerator()
func (s_ Set) ObjectEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("objectEnumerator"))
	return rv
}

// Returns a set of objects that pass a test in a given block, using the specified enumeration options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/objects(options:passingTest:)
func (s_ Set) ObjectsWithOptionsPassingTest(opts unsafe.Pointer, predicate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("objectsWithOptions:passingTest:"), opts, predicate)
	return rv
}

// Returns a set of objects that pass a test in a given block.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/objects(passingTest:)
func (s_ Set) ObjectsPassingTest(predicate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("objectsPassingTest:"), predicate)
	return rv
}

// Raises an exception.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/removeObserver(_:forKeyPath:)
func (s_ Set) RemoveObserverForKeyPath(observer unsafe.Pointer, keyPath string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeObserver:forKeyPath:"), observer, objc.String(keyPath))
}

// Raises an exception.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/removeObserver(_:forKeyPath:context:)
func (s_ Set) RemoveObserverForKeyPathContext(observer unsafe.Pointer, keyPath string, context unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeObserver:forKeyPath:context:"), observer, objc.String(keyPath), context)
}

// Invokes on each of the set’s members.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/setValue(_:forKey:)
func (s_ Set) SetValueForKey(value objc.ID, key string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setValue:forKey:"), value, objc.String(key))
}

// Returns an array of the set’s content sorted as specified by a given array of sort descriptors.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/sortedArray(using:)
func (s_ Set) SortedArrayUsingDescriptors(sortDescriptors unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](s_.ID, objc.Sel("sortedArrayUsingDescriptors:"), sortDescriptors)
	return rv
}

// Return a set containing the results of invoking on each of the receiving set’s members.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/value(forKey:)
func (s_ Set) ValueForKey(key string) objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("valueForKey:"), objc.String(key))
	return rv
}

// An array containing the set’s members, or an empty array if the set has no members.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/allObjects
func (s_ Set) AllObjects() []objc.ID {
	rv := objc.Send[[]objc.ID](s_.ID, objc.Sel("allObjects"))
	return rv
}

// The number of members in the set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/count
func (s_ Set) Count() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("count"))
	return rv
}

// A string that represents the contents of the set, formatted as a property list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet/description
func (s_ Set) Description() string {
	rv := objc.Send[string](s_.ID, objc.Sel("description"))
	return rv
}



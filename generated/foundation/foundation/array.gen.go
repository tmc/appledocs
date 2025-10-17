// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Array] class.
var ArrayClass objc.Class

func init() {
	ArrayClass = objc.GetClass("NSArray")
}

type Array struct {
	objc.ID
}

func ArrayFrom(ptr unsafe.Pointer) Array {
	return Array{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac Array) Alloc() Array {
	ret := objc.ID(ArrayClass).Send(objc.RegisterName("alloc"))
	return Array{ret}
}

// Init initializes the instance.
func (a_ Array) Init() Array {
	ret := a_.ID.Send(objc.RegisterName("init"))
	return Array{ret}
}
// Initializes a newly allocated array. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/init()
func NewArray() Array {
	instance := Array{}.Alloc()
	instance = instance.Init()
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly allocated array by placing in it the objects contained in a given array. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/init(array:)-o72h
func NewArrayWithArray(array unsafe.Pointer) Array {
	instance := Array{}.Alloc()
	sel := objc.RegisterName("initWithArray:")
	ret := instance.ID.Send(sel, array)
	instance = Array{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly allocated array using   as the source of data objects for the array. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/init(array:copyItems:)
func NewArrayWithArrayCopyItems(array unsafe.Pointer, flag bool) Array {
	instance := Array{}.Alloc()
	sel := objc.RegisterName("initWithArray:copyItems:")
	ret := instance.ID.Send(sel, array, flag)
	instance = Array{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/init(coder:)
func NewArrayWithCoder(coder unsafe.Pointer) Array {
	instance := Array{}.Alloc()
	sel := objc.RegisterName("initWithCoder:")
	ret := instance.ID.Send(sel, coder)
	instance = Array{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly allocated array with the contents of the file specified by a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/init(contentsOfFile:)
func NewArrayWithContentsOfFile(path string) Array {
	instance := Array{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfFile:")
	ret := instance.ID.Send(sel, path)
	instance = Array{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly allocated array with the contents of the location specified by a given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/init(contentsOfURL:)-5lo2y
func NewArrayWithContentsOfURL(url unsafe.Pointer) Array {
	instance := Array{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfURL:")
	ret := instance.ID.Send(sel, url)
	instance = Array{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/init(contentsOfURL:error:)
func NewArrayWithContentsOfURLError(url unsafe.Pointer, error unsafe.Pointer) Array {
	instance := Array{}.Alloc()
	sel := objc.RegisterName("initWithContentsOfURL:error:")
	ret := instance.ID.Send(sel, url, error)
	instance = Array{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly allocated array to include a given number of objects from a given C array. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/init(objects:count:)-5odxv
func NewArrayWithObjectsCount(objects unsafe.Pointer, cnt uint) Array {
	instance := Array{}.Alloc()
	sel := objc.RegisterName("initWithObjects:count:")
	ret := instance.ID.Send(sel, objects, cnt)
	instance = Array{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a newly allocated array by placing in it the objects in the argument list. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/initWithObjects:
func NewArrayWithObjects(firstObj unsafe.Pointer) Array {
	instance := Array{}.Alloc()
	sel := objc.RegisterName("initWithObjects:")
	ret := instance.ID.Send(sel, firstObj)
	instance = Array{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Creates and returns an empty array. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/array
func (ac Array) Array() unsafe.Pointer {
	sel := objc.RegisterName("array")
	ret := objc.ID(ArrayClass).Send(sel)
	return unsafe.Pointer(ret)
}
// Creates and returns an array containing the objects in another given array. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/arrayWithArray:
func (ac Array) ArrayWithArray(array unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("arrayWithArray:")
	ret := objc.ID(ArrayClass).Send(sel, array)
	return unsafe.Pointer(ret)
}
// Creates and returns an array containing the contents of the file specified by a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/arrayWithContentsOfFile:
func (ac Array) ArrayWithContentsOfFile(path string) unsafe.Pointer {
	sel := objc.RegisterName("arrayWithContentsOfFile:")
	ret := objc.ID(ArrayClass).Send(sel, path)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/arrayWithContentsOfURL:error:
func (ac Array) ArrayWithContentsOfURLError(url unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("arrayWithContentsOfURL:error:")
	ret := objc.ID(ArrayClass).Send(sel, url, error)
	return unsafe.Pointer(ret)
}
// Creates and returns an array containing the objects in the argument list. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/arrayWithObjects:
func (ac Array) ArrayWithObjects(firstObj unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("arrayWithObjects:")
	ret := objc.ID(ArrayClass).Send(sel, firstObj)
	return unsafe.Pointer(ret)
}
// Creates and returns an array containing the contents specified by a given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/init(contentsOfURL:)-fk8x
func (ac Array) ArrayWithContentsOfURL(url unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("arrayWithContentsOfURL:")
	ret := objc.ID(ArrayClass).Send(sel, url)
	return unsafe.Pointer(ret)
}
// Creates and returns an array containing a given object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/init(object:)
func (ac Array) ArrayWithObject(anObject unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("arrayWithObject:")
	ret := objc.ID(ArrayClass).Send(sel, anObject)
	return unsafe.Pointer(ret)
}
// Creates and returns an array that includes a given number of objects from a given C array. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/init(objects:count:)-7dct1
func (ac Array) ArrayWithObjectsCount(objects unsafe.Pointer, cnt uint) unsafe.Pointer {
	sel := objc.RegisterName("arrayWithObjects:count:")
	ret := objc.ID(ArrayClass).Send(sel, objects, cnt)
	return unsafe.Pointer(ret)
}
// Raises an exception. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/addObserver(_:forKeyPath:options:context:)
func (a_ Array) AddObserverForKeyPathOptionsContext(observer unsafe.Pointer, keyPath string, options unsafe.Pointer, context unsafe.Pointer) {
	sel := objc.RegisterName("addObserver:forKeyPath:options:context:")
	a_.ID.Send(sel, observer, keyPath, options, context)
}
// Registers an observer to receive key value observer notifications for the specified key-path relative to the objects at the indexes. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/addObserver(_:toObjectsAt:forKeyPath:options:context:)
func (a_ Array) AddObserverToObjectsAtIndexesForKeyPathOptionsContext(observer unsafe.Pointer, indexes unsafe.Pointer, keyPath string, options unsafe.Pointer, context unsafe.Pointer) {
	sel := objc.RegisterName("addObserver:toObjectsAtIndexes:forKeyPath:options:context:")
	a_.ID.Send(sel, observer, indexes, keyPath, options, context)
}
// Returns a new array that is a copy of the receiving array with a given object added to the end. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/adding(_:)
func (a_ Array) ArrayByAddingObject(anObject unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("arrayByAddingObject:")
	ret := a_.ID.Send(sel, anObject)
	return unsafe.Pointer(ret)
}
// Returns a new array that is a copy of the receiving array with the objects contained in another array added to the end. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/addingObjects(from:)
func (a_ Array) ArrayByAddingObjectsFromArray(otherArray unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("arrayByAddingObjectsFromArray:")
	ret := a_.ID.Send(sel, otherArray)
	return unsafe.Pointer(ret)
}
// Creates a new array by applying a difference object to an existing array. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/arrayByApplyingDifference:
func (a_ Array) ArrayByApplyingDifference(difference unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("arrayByApplyingDifference:")
	ret := a_.ID.Send(sel, difference)
	return unsafe.Pointer(ret)
}
// Constructs and returns an   object that is the result of interposing a given separator between the elements of the array. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/componentsJoined(by:)
func (a_ Array) ComponentsJoinedByString(separator string) unsafe.Pointer {
	sel := objc.RegisterName("componentsJoinedByString:")
	ret := a_.ID.Send(sel, separator)
	return unsafe.Pointer(ret)
}
// Returns a Boolean value that indicates whether a given object is present in the array. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/contains(_:)
func (a_ Array) ContainsObject(anObject unsafe.Pointer) bool {
	sel := objc.RegisterName("containsObject:")
	ret := a_.ID.Send(sel, anObject)
	return ret != 0
}
// Returns a string that represents the contents of the array, formatted as a property list. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/description(withLocale:)
func (a_ Array) DescriptionWithLocale(locale objc.ID) unsafe.Pointer {
	sel := objc.RegisterName("descriptionWithLocale:")
	ret := a_.ID.Send(sel, locale)
	return unsafe.Pointer(ret)
}
// Returns a string that represents the contents of the array, formatted as a property list. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/description(withLocale:indent:)
func (a_ Array) DescriptionWithLocaleIndent(locale objc.ID, level uint) unsafe.Pointer {
	sel := objc.RegisterName("descriptionWithLocale:indent:")
	ret := a_.ID.Send(sel, locale, level)
	return unsafe.Pointer(ret)
}
// Compares two arrays to create a difference object that represents the changes between them. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/differenceFromArray:
func (a_ Array) DifferenceFromArray(other unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("differenceFromArray:")
	ret := a_.ID.Send(sel, other)
	return unsafe.Pointer(ret)
}
// Compares two arrays, with options, to create a difference object that represents the changes between them. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/differenceFromArray:withOptions:
func (a_ Array) DifferenceFromArrayWithOptions(other unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("differenceFromArray:withOptions:")
	ret := a_.ID.Send(sel, other, options)
	return unsafe.Pointer(ret)
}
// Compares two arrays, using the provided block and with options, to create a difference object that represents the changes between them. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/differenceFromArray:withOptions:usingEquivalenceTest:
func (a_ Array) DifferenceFromArrayWithOptionsUsingEquivalenceTest(other unsafe.Pointer, options unsafe.Pointer, block unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("differenceFromArray:withOptions:usingEquivalenceTest:")
	ret := a_.ID.Send(sel, other, options, block)
	return unsafe.Pointer(ret)
}
// Executes a given closure or block using each object in the array, starting with the first object and continuing through the array to the last object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/enumerateObjects(_:)
func (a_ Array) EnumerateObjectsUsingBlock(block unsafe.Pointer) {
	sel := objc.RegisterName("enumerateObjectsUsingBlock:")
	a_.ID.Send(sel, block)
}
// Executes a given block using the objects in the array at the specified indexes. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/enumerateObjects(at:options:using:)
func (a_ Array) EnumerateObjectsAtIndexesOptionsUsingBlock(s unsafe.Pointer, opts unsafe.Pointer, block unsafe.Pointer) {
	sel := objc.RegisterName("enumerateObjectsAtIndexes:options:usingBlock:")
	a_.ID.Send(sel, s, opts, block)
}
// Executes a given closure or block using each object in the array with the specified options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/enumerateObjects(options:using:)
func (a_ Array) EnumerateObjectsWithOptionsUsingBlock(opts unsafe.Pointer, block unsafe.Pointer) {
	sel := objc.RegisterName("enumerateObjectsWithOptions:usingBlock:")
	a_.ID.Send(sel, opts, block)
}
// Evaluates a given predicate against each object in the receiving array and returns a new array containing the objects for which the predicate returns true. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/filtered(using:)
func (a_ Array) FilteredArrayUsingPredicate(predicate unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("filteredArrayUsingPredicate:")
	ret := a_.ID.Send(sel, predicate)
	return unsafe.Pointer(ret)
}
// Returns the first object contained in the receiving array that’s equal to an object in another given array. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/firstObjectCommon(with:)
func (a_ Array) FirstObjectCommonWithArray(otherArray unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("firstObjectCommonWithArray:")
	ret := a_.ID.Send(sel, otherArray)
	return unsafe.Pointer(ret)
}
// Copies all the objects contained in the array to  . [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/getObjects:
func (a_ Array) GetObjects(objects unsafe.Pointer) {
	sel := objc.RegisterName("getObjects:")
	a_.ID.Send(sel, objects)
}
// Copies references to objects contained in the array that fall within the specified range to  . [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/getObjects:range:
func (a_ Array) GetObjectsRange(objects unsafe.Pointer, range_ unsafe.Pointer) {
	sel := objc.RegisterName("getObjects:range:")
	a_.ID.Send(sel, objects, range_)
}
// Returns the lowest index whose corresponding array value is equal to a given object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/index(of:)
func (a_ Array) IndexOfObject(anObject unsafe.Pointer) uint {
	sel := objc.RegisterName("indexOfObject:")
	ret := a_.ID.Send(sel, anObject)
	return uint(ret)
}
// Returns the lowest index within a specified range whose corresponding array value is equal to a given object . [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/index(of:in:)
func (a_ Array) IndexOfObjectInRange(anObject unsafe.Pointer, range_ unsafe.Pointer) uint {
	sel := objc.RegisterName("indexOfObject:inRange:")
	ret := a_.ID.Send(sel, anObject, range_)
	return uint(ret)
}
// Returns the index, within a specified range, of an object compared with elements in the array using a given   block. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/index(of:inSortedRange:options:usingComparator:)
func (a_ Array) IndexOfObjectInSortedRangeOptionsUsingComparator(obj unsafe.Pointer, r unsafe.Pointer, opts unsafe.Pointer, cmp unsafe.Pointer) uint {
	sel := objc.RegisterName("indexOfObject:inSortedRange:options:usingComparator:")
	ret := a_.ID.Send(sel, obj, r, opts, cmp)
	return uint(ret)
}
// Returns the index, from a given set of indexes, of the first object in the array that passes a test in a given block for a given set of enumeration options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/indexOfObject(at:options:passingTest:)
func (a_ Array) IndexOfObjectAtIndexesOptionsPassingTest(s unsafe.Pointer, opts unsafe.Pointer, predicate unsafe.Pointer) uint {
	sel := objc.RegisterName("indexOfObjectAtIndexes:options:passingTest:")
	ret := a_.ID.Send(sel, s, opts, predicate)
	return uint(ret)
}
// Returns the index of an object in the array that passes a test in a given block for a given set of enumeration options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/indexOfObject(options:passingTest:)
func (a_ Array) IndexOfObjectWithOptionsPassingTest(opts unsafe.Pointer, predicate unsafe.Pointer) uint {
	sel := objc.RegisterName("indexOfObjectWithOptions:passingTest:")
	ret := a_.ID.Send(sel, opts, predicate)
	return uint(ret)
}
// Returns the index of the first object in the array that passes a test in a given block. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/indexOfObject(passingTest:)
func (a_ Array) IndexOfObjectPassingTest(predicate unsafe.Pointer) uint {
	sel := objc.RegisterName("indexOfObjectPassingTest:")
	ret := a_.ID.Send(sel, predicate)
	return uint(ret)
}
// Returns the lowest index whose corresponding array value is identical to a given object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/indexOfObjectIdentical(to:)
func (a_ Array) IndexOfObjectIdenticalTo(anObject unsafe.Pointer) uint {
	sel := objc.RegisterName("indexOfObjectIdenticalTo:")
	ret := a_.ID.Send(sel, anObject)
	return uint(ret)
}
// Returns the lowest index within a specified range whose corresponding array value is equal to a given object . [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/indexOfObjectIdentical(to:in:)
func (a_ Array) IndexOfObjectIdenticalToInRange(anObject unsafe.Pointer, range_ unsafe.Pointer) uint {
	sel := objc.RegisterName("indexOfObjectIdenticalTo:inRange:")
	ret := a_.ID.Send(sel, anObject, range_)
	return uint(ret)
}
// Returns the indexes, from a given set of indexes, of objects in the array that pass a test in a given block for a given set of enumeration options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/indexesOfObjects(at:options:passingTest:)
func (a_ Array) IndexesOfObjectsAtIndexesOptionsPassingTest(s unsafe.Pointer, opts unsafe.Pointer, predicate unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("indexesOfObjectsAtIndexes:options:passingTest:")
	ret := a_.ID.Send(sel, s, opts, predicate)
	return unsafe.Pointer(ret)
}
// Returns the indexes of objects in the array that pass a test in a given block for a given set of enumeration options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/indexesOfObjects(options:passingTest:)
func (a_ Array) IndexesOfObjectsWithOptionsPassingTest(opts unsafe.Pointer, predicate unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("indexesOfObjectsWithOptions:passingTest:")
	ret := a_.ID.Send(sel, opts, predicate)
	return unsafe.Pointer(ret)
}
// Returns the indexes of objects in the array that pass a test in a given block. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/indexesOfObjects(passingTest:)
func (a_ Array) IndexesOfObjectsPassingTest(predicate unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("indexesOfObjectsPassingTest:")
	ret := a_.ID.Send(sel, predicate)
	return unsafe.Pointer(ret)
}
// Compares the receiving array to another array. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/isEqual(to:)
func (a_ Array) IsEqualToArray(otherArray unsafe.Pointer) bool {
	sel := objc.RegisterName("isEqualToArray:")
	ret := a_.ID.Send(sel, otherArray)
	return ret != 0
}
// Sends to each object in the array the message identified by a given selector, starting with the first object and continuing through the array to the last object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/makeObjectsPerformSelector:
func (a_ Array) MakeObjectsPerformSelector(aSelector objc.SEL) {
	sel := objc.RegisterName("makeObjectsPerformSelector:")
	a_.ID.Send(sel, aSelector)
}
// Sends the   message to each object in the array, starting with the first object and continuing through the array to the last object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/makeObjectsPerformSelector:withObject:
func (a_ Array) MakeObjectsPerformSelectorWithObject(aSelector objc.SEL, argument objc.ID) {
	sel := objc.RegisterName("makeObjectsPerformSelector:withObject:")
	a_.ID.Send(sel, aSelector, argument)
}
// Returns the object located at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/object(at:)
func (a_ Array) ObjectAtIndex(index uint) unsafe.Pointer {
	sel := objc.RegisterName("objectAtIndex:")
	ret := a_.ID.Send(sel, index)
	return unsafe.Pointer(ret)
}
// Returns an enumerator object that lets you access each object in the array. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/objectEnumerator()
func (a_ Array) ObjectEnumerator() unsafe.Pointer {
	sel := objc.RegisterName("objectEnumerator")
	ret := a_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns an array containing the objects in the array at the indexes specified by a given index set. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/objects(at:)
func (a_ Array) ObjectsAtIndexes(indexes unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("objectsAtIndexes:")
	ret := a_.ID.Send(sel, indexes)
	return unsafe.Pointer(ret)
}
// Returns an array containing all the pathname elements in the receiving array that have filename extensions from a given array. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/pathsMatchingExtensions(_:)
func (a_ Array) PathsMatchingExtensions(filterTypes unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("pathsMatchingExtensions:")
	ret := a_.ID.Send(sel, filterTypes)
	return unsafe.Pointer(ret)
}
// Raises an exception. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/removeObserver(_:forKeyPath:)
func (a_ Array) RemoveObserverForKeyPath(observer unsafe.Pointer, keyPath string) {
	sel := objc.RegisterName("removeObserver:forKeyPath:")
	a_.ID.Send(sel, observer, keyPath)
}
// Raises an exception. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/removeObserver(_:forKeyPath:context:)
func (a_ Array) RemoveObserverForKeyPathContext(observer unsafe.Pointer, keyPath string, context unsafe.Pointer) {
	sel := objc.RegisterName("removeObserver:forKeyPath:context:")
	a_.ID.Send(sel, observer, keyPath, context)
}
// Removes   from all key value observer notifications associated with the specified   relative to the array’s objects at  . [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/removeObserver(_:fromObjectsAt:forKeyPath:)
func (a_ Array) RemoveObserverFromObjectsAtIndexesForKeyPath(observer unsafe.Pointer, indexes unsafe.Pointer, keyPath string) {
	sel := objc.RegisterName("removeObserver:fromObjectsAtIndexes:forKeyPath:")
	a_.ID.Send(sel, observer, indexes, keyPath)
}
// Raises an exception. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/removeObserver(_:fromObjectsAt:forKeyPath:context:)
func (a_ Array) RemoveObserverFromObjectsAtIndexesForKeyPathContext(observer unsafe.Pointer, indexes unsafe.Pointer, keyPath string, context unsafe.Pointer) {
	sel := objc.RegisterName("removeObserver:fromObjectsAtIndexes:forKeyPath:context:")
	a_.ID.Send(sel, observer, indexes, keyPath, context)
}
// Returns an enumerator object that lets you access each object in the array, in reverse order. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/reverseObjectEnumerator()
func (a_ Array) ReverseObjectEnumerator() unsafe.Pointer {
	sel := objc.RegisterName("reverseObjectEnumerator")
	ret := a_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Invokes   on each of the array’s items using the specified   and  . [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/setValue(_:forKey:)
func (a_ Array) SetValueForKey(value objc.ID, key string) {
	sel := objc.RegisterName("setValue:forKey:")
	a_.ID.Send(sel, value, key)
}
// Returns a new array that lists this array’s elements in a random order. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/shuffled()
func (a_ Array) ShuffledArray() unsafe.Pointer {
	sel := objc.RegisterName("shuffledArray")
	ret := a_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns a new array that lists this array’s elements in a random order, using the specified random source. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/shuffled(using:)
func (a_ Array) ShuffledArrayWithRandomSource(randomSource unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("shuffledArrayWithRandomSource:")
	ret := a_.ID.Send(sel, randomSource)
	return unsafe.Pointer(ret)
}
// Returns a new array that lists the receiving array’s elements in ascending order as defined by the comparison function  . [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/sortedArray(_:context:)
func (a_ Array) SortedArrayUsingFunctionContext(comparator unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("sortedArrayUsingFunction:context:")
	ret := a_.ID.Send(sel, comparator, context)
	return unsafe.Pointer(ret)
}
// Returns a new array that lists the receiving array’s elements in ascending order as defined by the comparison function  . [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/sortedArray(_:context:hint:)
func (a_ Array) SortedArrayUsingFunctionContextHint(comparator unsafe.Pointer, context unsafe.Pointer, hint unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("sortedArrayUsingFunction:context:hint:")
	ret := a_.ID.Send(sel, comparator, context, hint)
	return unsafe.Pointer(ret)
}
// Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given   block. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/sortedArray(comparator:)
func (a_ Array) SortedArrayUsingComparator(cmptr unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("sortedArrayUsingComparator:")
	ret := a_.ID.Send(sel, cmptr)
	return unsafe.Pointer(ret)
}
// Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given   block. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/sortedArray(options:usingComparator:)
func (a_ Array) SortedArrayWithOptionsUsingComparator(opts unsafe.Pointer, cmptr unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("sortedArrayWithOptions:usingComparator:")
	ret := a_.ID.Send(sel, opts, cmptr)
	return unsafe.Pointer(ret)
}
// Returns a copy of the receiving array sorted as specified by a given array of sort descriptors. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/sortedArray(using:)-82wi1
func (a_ Array) SortedArrayUsingDescriptors(sortDescriptors unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("sortedArrayUsingDescriptors:")
	ret := a_.ID.Send(sel, sortDescriptors)
	return unsafe.Pointer(ret)
}
// Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given selector. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/sortedArray(using:)-9nhh9
func (a_ Array) SortedArrayUsingSelector(comparator objc.SEL) unsafe.Pointer {
	sel := objc.RegisterName("sortedArrayUsingSelector:")
	ret := a_.ID.Send(sel, comparator)
	return unsafe.Pointer(ret)
}
// Returns a new array containing the receiving array’s elements that fall within the limits specified by a given range. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/subarray(with:)
func (a_ Array) SubarrayWithRange(range_ unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("subarrayWithRange:")
	ret := a_.ID.Send(sel, range_)
	return unsafe.Pointer(ret)
}
// Returns the object at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/subscript(_:)
func (a_ Array) ObjectAtIndexedSubscript(idx uint) unsafe.Pointer {
	sel := objc.RegisterName("objectAtIndexedSubscript:")
	ret := a_.ID.Send(sel, idx)
	return unsafe.Pointer(ret)
}
// Returns an array containing the results of invoking   using   on each of the array’s objects. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/value(forKey:)
func (a_ Array) ValueForKey(key string) objc.ID {
	sel := objc.RegisterName("valueForKey:")
	ret := a_.ID.Send(sel, key)
	return ret
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/write(to:)
func (a_ Array) WriteToURLError(url unsafe.Pointer, error unsafe.Pointer) bool {
	sel := objc.RegisterName("writeToURL:error:")
	ret := a_.ID.Send(sel, url, error)
	return ret != 0
}
// Writes the contents of the array to the location specified by a given URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/write(to:atomically:)
func (a_ Array) WriteToURLAtomically(url unsafe.Pointer, atomically bool) bool {
	sel := objc.RegisterName("writeToURL:atomically:")
	ret := a_.ID.Send(sel, url, atomically)
	return ret != 0
}
// Writes the contents of the array to a file at a given path. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSArray/write(toFile:atomically:)
func (a_ Array) WriteToFileAtomically(path string, useAuxiliaryFile bool) bool {
	sel := objc.RegisterName("writeToFile:atomically:")
	ret := a_.ID.Send(sel, path, useAuxiliaryFile)
	return ret != 0
}


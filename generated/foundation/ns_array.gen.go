// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Array] class.
var (
	ArrayClass     _ArrayClass
	ArrayClassOnce sync.Once
)

func getArrayClass() _ArrayClass {
	ArrayClassOnce.Do(func() {
		ArrayClass = _ArrayClass{objc.GetClass("NSArray")}
	})
	return ArrayClass
}

type _ArrayClass struct {
	class objc.Class
}

// An interface definition for the [Array] class.
type IArray interface {
	objectivec.IObject
	AddObserverToObjectsAtIndexesForKeyPathOptionsContext(observer unsafe.Pointer, indexes unsafe.Pointer, keyPath string, options unsafe.Pointer, context unsafe.Pointer)
	ArrayByAddingObject(anObject unsafe.Pointer) []objc.ID
	ArrayByAddingObjectsFromArray(otherArray unsafe.Pointer) []objc.ID
	ArrayByApplyingDifference(difference unsafe.Pointer) []objc.ID
	ComponentsJoinedByString(separator string) string
	ContainsObject(anObject unsafe.Pointer) bool
	DescriptionWithLocale(locale objc.ID) string
	DescriptionWithLocaleIndent(locale objc.ID, level uint) string
	DifferenceFromArray(other unsafe.Pointer) unsafe.Pointer
	DifferenceFromArrayWithOptions(other unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer
	DifferenceFromArrayWithOptionsUsingEquivalenceTest(other unsafe.Pointer, options unsafe.Pointer, block unsafe.Pointer) unsafe.Pointer
	EnumerateObjectsUsingBlock(block unsafe.Pointer)
	EnumerateObjectsAtIndexesOptionsUsingBlock(s unsafe.Pointer, opts unsafe.Pointer, block unsafe.Pointer)
	EnumerateObjectsWithOptionsUsingBlock(opts unsafe.Pointer, block unsafe.Pointer)
	FilteredArrayUsingPredicate(predicate unsafe.Pointer) []objc.ID
	FirstObjectCommonWithArray(otherArray unsafe.Pointer) unsafe.Pointer
	GetObjects(objects unsafe.Pointer)
	GetObjectsRange(objects unsafe.Pointer, range_ Range)
	IndexOfObject(anObject unsafe.Pointer) uint
	IndexOfObjectInRange(anObject unsafe.Pointer, range_ Range) uint
	IndexOfObjectInSortedRangeOptionsUsingComparator(obj unsafe.Pointer, r Range, opts unsafe.Pointer, cmp unsafe.Pointer) uint
	IndexOfObjectAtIndexesOptionsPassingTest(s unsafe.Pointer, opts unsafe.Pointer, predicate unsafe.Pointer) uint
	IndexOfObjectWithOptionsPassingTest(opts unsafe.Pointer, predicate unsafe.Pointer) uint
	IndexOfObjectPassingTest(predicate unsafe.Pointer) uint
	IndexOfObjectIdenticalTo(anObject unsafe.Pointer) uint
	IndexOfObjectIdenticalToInRange(anObject unsafe.Pointer, range_ Range) uint
	IndexesOfObjectsAtIndexesOptionsPassingTest(s unsafe.Pointer, opts unsafe.Pointer, predicate unsafe.Pointer) unsafe.Pointer
	IndexesOfObjectsWithOptionsPassingTest(opts unsafe.Pointer, predicate unsafe.Pointer) unsafe.Pointer
	IndexesOfObjectsPassingTest(predicate unsafe.Pointer) unsafe.Pointer
	IsEqualToArray(otherArray unsafe.Pointer) bool
	MakeObjectsPerformSelector(aSelector objc.SEL)
	MakeObjectsPerformSelectorWithObject(aSelector objc.SEL, argument objc.ID)
	ObjectAtIndex(index uint) unsafe.Pointer
	ObjectEnumerator() unsafe.Pointer
	ObjectsAtIndexes(indexes unsafe.Pointer) []objc.ID
	PathsMatchingExtensions(filterTypes unsafe.Pointer) []string
	RemoveObserverFromObjectsAtIndexesForKeyPath(observer unsafe.Pointer, indexes unsafe.Pointer, keyPath string)
	RemoveObserverFromObjectsAtIndexesForKeyPathContext(observer unsafe.Pointer, indexes unsafe.Pointer, keyPath string, context unsafe.Pointer)
	ReverseObjectEnumerator() unsafe.Pointer
	ShuffledArray() []objc.ID
	ShuffledArrayWithRandomSource(randomSource unsafe.Pointer) []objc.ID
	SortedArrayUsingFunctionContext(comparator unsafe.Pointer, context unsafe.Pointer) []objc.ID
	SortedArrayUsingFunctionContextHint(comparator unsafe.Pointer, context unsafe.Pointer, hint unsafe.Pointer) []objc.ID
	SortedArrayUsingComparator(cmptr unsafe.Pointer) []objc.ID
	SortedArrayWithOptionsUsingComparator(opts unsafe.Pointer, cmptr unsafe.Pointer) []objc.ID
	SortedArrayUsingDescriptors(sortDescriptors unsafe.Pointer) []objc.ID
	SortedArrayUsingSelector(comparator objc.SEL) []objc.ID
	SubarrayWithRange(range_ Range) []objc.ID
	ObjectAtIndexedSubscript(idx uint) unsafe.Pointer
	WriteToURLError(url unsafe.Pointer, error_ unsafe.Pointer) bool
	WriteToURLAtomically(url unsafe.Pointer, atomically bool) bool
	WriteToFileAtomically(path string, useAuxiliaryFile bool) bool
}

// A static ordered collection of objects.
//
// You can use this type in Swift instead of an constant in cases that require reference semantics. and its subclass manage ordered collections of objects called . creates static arrays, and creates dynamic arrays. You can use arrays when you need an ordered collection of objects. is “toll-free bridged” with its Core Foundation counterpart, . See for more information on toll-free bridging.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray
type Array struct {
	objectivec.Object
}

// ArrayFrom constructs a [Array] from an unsafe.Pointer.
//
// A static ordered collection of objects.
func ArrayFrom(ptr unsafe.Pointer) Array {
	return Array{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _ArrayClass) Alloc() Array {
	rv := objc.Send[Array](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ArrayClass) New() Array {
	rv := objc.Send[Array](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Array) Init() Array {
	rv := objc.Send[Array](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Array) Autorelease() Array {
	rv := objc.Send[Array](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewArray creates a new Array instance.
func NewArray() Array {
	return getArrayClass().New()
}




// Initializes a newly allocated array by placing in it the objects contained in a given array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/init(array:)-o72h
func NewArrayWithArray(array unsafe.Pointer) Array {
	instance := getArrayClass().Alloc()
	rv := objc.Send[Array](instance.ID, objc.Sel("initWithArray:"), array)
	rv.Autorelease()
	return rv
}



// Initializes a newly allocated array using as the source of data objects for the array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/init(array:copyItems:)
func NewArrayWithArrayCopyItems(array unsafe.Pointer, flag bool) Array {
	instance := getArrayClass().Alloc()
	rv := objc.Send[Array](instance.ID, objc.Sel("initWithArray:copyItems:"), array, flag)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/init(coder:)
func NewArrayWithCoder(coder unsafe.Pointer) Array {
	instance := getArrayClass().Alloc()
	rv := objc.Send[Array](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Initializes a newly allocated array with the contents of the file specified by a given path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/init(contentsOfFile:)
func NewArrayWithContentsOfFile(path string) Array {
	instance := getArrayClass().Alloc()
	rv := objc.Send[Array](instance.ID, objc.Sel("initWithContentsOfFile:"), objc.String(path))
	rv.Autorelease()
	return rv
}



// Initializes a newly allocated array with the contents of the location specified by a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/init(contentsOfURL:)-5lo2y
func NewArrayWithContentsOfURL(url unsafe.Pointer) Array {
	instance := getArrayClass().Alloc()
	rv := objc.Send[Array](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/init(contentsOfURL:error:)
func NewArrayWithContentsOfURLError(url unsafe.Pointer, error_ unsafe.Pointer) Array {
	instance := getArrayClass().Alloc()
	rv := objc.Send[Array](instance.ID, objc.Sel("initWithContentsOfURL:error:"), url, error_)
	rv.Autorelease()
	return rv
}



// Creates and returns an array containing a given object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/init(object:)
func NewArrayWithObject(anObject unsafe.Pointer) Array {
	rv := objc.Send[Array](objc.ID(getArrayClass().class), objc.Sel("arrayWithObject:"), anObject)
	return rv
}



// Initializes a newly allocated array by placing in it the objects in the argument list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/initWithObjects:
func NewArrayWithObjects(firstObj unsafe.Pointer) Array {
	instance := getArrayClass().Alloc()
	rv := objc.Send[Array](instance.ID, objc.Sel("initWithObjects:"), firstObj)
	rv.Autorelease()
	return rv
}



// Initializes a newly allocated array to include a given number of objects from a given C array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/init(objects:count:)-5odxv
func NewArrayWithObjectsCount(objects unsafe.Pointer, cnt uint) Array {
	instance := getArrayClass().Alloc()
	rv := objc.Send[Array](instance.ID, objc.Sel("initWithObjects:count:"), objects, cnt)
	rv.Autorelease()
	return rv
}


// Creates and returns an empty array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/array
func (ac _ArrayClass) Array() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("array"))
	return rv
}

// Creates and returns an array containing the objects in another given array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/arrayWithArray:
func (ac _ArrayClass) ArrayWithArray(array unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("arrayWithArray:"), array)
	return rv
}

// Creates and returns an array containing the contents of the file specified by a given path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/arrayWithContentsOfFile:
func (ac _ArrayClass) ArrayWithContentsOfFile(path string) []objc.ID {
	rv := objc.Send[[]objc.ID](objc.ID(ac.class), objc.Sel("arrayWithContentsOfFile:"), objc.String(path))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/arrayWithContentsOfURL:error:
func (ac _ArrayClass) ArrayWithContentsOfURLError(url unsafe.Pointer, error_ unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](objc.ID(ac.class), objc.Sel("arrayWithContentsOfURL:error:"), url, error_)
	return rv
}

// Creates and returns an array containing the objects in the argument list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/arrayWithObjects:
func (ac _ArrayClass) ArrayWithObjects(firstObj unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("arrayWithObjects:"), firstObj)
	return rv
}

// Creates and returns an array containing the contents specified by a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/init(contentsOfURL:)-fk8x
func (ac _ArrayClass) ArrayWithContentsOfURL(url unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](objc.ID(ac.class), objc.Sel("arrayWithContentsOfURL:"), url)
	return rv
}

// Creates and returns an array containing a given object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/init(object:)
func (ac _ArrayClass) ArrayWithObject(anObject unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("arrayWithObject:"), anObject)
	return rv
}

// Creates and returns an array that includes a given number of objects from a given C array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/init(objects:count:)-7dct1
func (ac _ArrayClass) ArrayWithObjectsCount(objects unsafe.Pointer, cnt uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("arrayWithObjects:count:"), objects, cnt)
	return rv
}

// Raises an exception.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/addObserver(_:forKeyPath:options:context:)
func (a_ Array) AddObserverForKeyPathOptionsContext(observer unsafe.Pointer, keyPath string, options unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addObserver:forKeyPath:options:context:"), observer, objc.String(keyPath), options, context)
}

// Registers an observer to receive key value observer notifications for the specified key-path relative to the objects at the indexes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/addObserver(_:toObjectsAt:forKeyPath:options:context:)
func (a_ Array) AddObserverToObjectsAtIndexesForKeyPathOptionsContext(observer unsafe.Pointer, indexes unsafe.Pointer, keyPath string, options unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addObserver:toObjectsAtIndexes:forKeyPath:options:context:"), observer, indexes, objc.String(keyPath), options, context)
}

// Returns a new array that is a copy of the receiving array with a given object added to the end.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/adding(_:)
func (a_ Array) ArrayByAddingObject(anObject unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("arrayByAddingObject:"), anObject)
	return rv
}

// Returns a new array that is a copy of the receiving array with the objects contained in another array added to the end.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/addingObjects(from:)
func (a_ Array) ArrayByAddingObjectsFromArray(otherArray unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("arrayByAddingObjectsFromArray:"), otherArray)
	return rv
}

// Creates a new array by applying a difference object to an existing array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/arrayByApplyingDifference:
func (a_ Array) ArrayByApplyingDifference(difference unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("arrayByApplyingDifference:"), difference)
	return rv
}

// Constructs and returns an object that is the result of interposing a given separator between the elements of the array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/componentsJoined(by:)
func (a_ Array) ComponentsJoinedByString(separator string) string {
	rv := objc.Send[string](a_.ID, objc.Sel("componentsJoinedByString:"), objc.String(separator))
	return rv
}

// Returns a Boolean value that indicates whether a given object is present in the array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/contains(_:)
func (a_ Array) ContainsObject(anObject unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("containsObject:"), anObject)
	return rv
}

// Returns a string that represents the contents of the array, formatted as a property list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/description(withLocale:)
func (a_ Array) DescriptionWithLocale(locale objc.ID) string {
	rv := objc.Send[string](a_.ID, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}

// Returns a string that represents the contents of the array, formatted as a property list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/description(withLocale:indent:)
func (a_ Array) DescriptionWithLocaleIndent(locale objc.ID, level uint) string {
	rv := objc.Send[string](a_.ID, objc.Sel("descriptionWithLocale:indent:"), locale, level)
	return rv
}

// Compares two arrays to create a difference object that represents the changes between them.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/differenceFromArray:
func (a_ Array) DifferenceFromArray(other unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("differenceFromArray:"), other)
	return rv
}

// Compares two arrays, with options, to create a difference object that represents the changes between them.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/differenceFromArray:withOptions:
func (a_ Array) DifferenceFromArrayWithOptions(other unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("differenceFromArray:withOptions:"), other, options)
	return rv
}

// Compares two arrays, using the provided block and with options, to create a difference object that represents the changes between them.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/differenceFromArray:withOptions:usingEquivalenceTest:
func (a_ Array) DifferenceFromArrayWithOptionsUsingEquivalenceTest(other unsafe.Pointer, options unsafe.Pointer, block unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("differenceFromArray:withOptions:usingEquivalenceTest:"), other, options, block)
	return rv
}

// Executes a given closure or block using each object in the array, starting with the first object and continuing through the array to the last object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/enumerateObjects(_:)
func (a_ Array) EnumerateObjectsUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("enumerateObjectsUsingBlock:"), block)
}

// Executes a given block using the objects in the array at the specified indexes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/enumerateObjects(at:options:using:)
func (a_ Array) EnumerateObjectsAtIndexesOptionsUsingBlock(s unsafe.Pointer, opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("enumerateObjectsAtIndexes:options:usingBlock:"), s, opts, block)
}

// Executes a given closure or block using each object in the array with the specified options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/enumerateObjects(options:using:)
func (a_ Array) EnumerateObjectsWithOptionsUsingBlock(opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("enumerateObjectsWithOptions:usingBlock:"), opts, block)
}

// Evaluates a given predicate against each object in the receiving array and returns a new array containing the objects for which the predicate returns true.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/filtered(using:)
func (a_ Array) FilteredArrayUsingPredicate(predicate unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("filteredArrayUsingPredicate:"), predicate)
	return rv
}

// Returns the first object contained in the receiving array that’s equal to an object in another given array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/firstObjectCommon(with:)
func (a_ Array) FirstObjectCommonWithArray(otherArray unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("firstObjectCommonWithArray:"), otherArray)
	return rv
}

// Copies all the objects contained in the array to .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/getObjects:
func (a_ Array) GetObjects(objects unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("getObjects:"), objects)
}

// Copies references to objects contained in the array that fall within the specified range to .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/getObjects:range:
func (a_ Array) GetObjectsRange(objects unsafe.Pointer, range_ Range) {
	objc.Send[objc.ID](a_.ID, objc.Sel("getObjects:range:"), objects, range_)
}

// Returns the lowest index whose corresponding array value is equal to a given object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/index(of:)
func (a_ Array) IndexOfObject(anObject unsafe.Pointer) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexOfObject:"), anObject)
	return rv
}

// Returns the lowest index within a specified range whose corresponding array value is equal to a given object .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/index(of:in:)
func (a_ Array) IndexOfObjectInRange(anObject unsafe.Pointer, range_ Range) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexOfObject:inRange:"), anObject, range_)
	return rv
}

// Returns the index, within a specified range, of an object compared with elements in the array using a given block.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/index(of:inSortedRange:options:usingComparator:)
func (a_ Array) IndexOfObjectInSortedRangeOptionsUsingComparator(obj unsafe.Pointer, r Range, opts unsafe.Pointer, cmp unsafe.Pointer) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexOfObject:inSortedRange:options:usingComparator:"), obj, r, opts, cmp)
	return rv
}

// Returns the index, from a given set of indexes, of the first object in the array that passes a test in a given block for a given set of enumeration options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/indexOfObject(at:options:passingTest:)
func (a_ Array) IndexOfObjectAtIndexesOptionsPassingTest(s unsafe.Pointer, opts unsafe.Pointer, predicate unsafe.Pointer) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexOfObjectAtIndexes:options:passingTest:"), s, opts, predicate)
	return rv
}

// Returns the index of an object in the array that passes a test in a given block for a given set of enumeration options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/indexOfObject(options:passingTest:)
func (a_ Array) IndexOfObjectWithOptionsPassingTest(opts unsafe.Pointer, predicate unsafe.Pointer) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexOfObjectWithOptions:passingTest:"), opts, predicate)
	return rv
}

// Returns the index of the first object in the array that passes a test in a given block.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/indexOfObject(passingTest:)
func (a_ Array) IndexOfObjectPassingTest(predicate unsafe.Pointer) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexOfObjectPassingTest:"), predicate)
	return rv
}

// Returns the lowest index whose corresponding array value is identical to a given object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/indexOfObjectIdentical(to:)
func (a_ Array) IndexOfObjectIdenticalTo(anObject unsafe.Pointer) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexOfObjectIdenticalTo:"), anObject)
	return rv
}

// Returns the lowest index within a specified range whose corresponding array value is equal to a given object .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/indexOfObjectIdentical(to:in:)
func (a_ Array) IndexOfObjectIdenticalToInRange(anObject unsafe.Pointer, range_ Range) uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexOfObjectIdenticalTo:inRange:"), anObject, range_)
	return rv
}

// Returns the indexes, from a given set of indexes, of objects in the array that pass a test in a given block for a given set of enumeration options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/indexesOfObjects(at:options:passingTest:)
func (a_ Array) IndexesOfObjectsAtIndexesOptionsPassingTest(s unsafe.Pointer, opts unsafe.Pointer, predicate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("indexesOfObjectsAtIndexes:options:passingTest:"), s, opts, predicate)
	return rv
}

// Returns the indexes of objects in the array that pass a test in a given block for a given set of enumeration options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/indexesOfObjects(options:passingTest:)
func (a_ Array) IndexesOfObjectsWithOptionsPassingTest(opts unsafe.Pointer, predicate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("indexesOfObjectsWithOptions:passingTest:"), opts, predicate)
	return rv
}

// Returns the indexes of objects in the array that pass a test in a given block.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/indexesOfObjects(passingTest:)
func (a_ Array) IndexesOfObjectsPassingTest(predicate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("indexesOfObjectsPassingTest:"), predicate)
	return rv
}

// Compares the receiving array to another array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/isEqual(to:)
func (a_ Array) IsEqualToArray(otherArray unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEqualToArray:"), otherArray)
	return rv
}

// Sends to each object in the array the message identified by a given selector, starting with the first object and continuing through the array to the last object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/makeObjectsPerformSelector:
func (a_ Array) MakeObjectsPerformSelector(aSelector objc.SEL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("makeObjectsPerformSelector:"), aSelector)
}

// Sends the message to each object in the array, starting with the first object and continuing through the array to the last object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/makeObjectsPerformSelector:withObject:
func (a_ Array) MakeObjectsPerformSelectorWithObject(aSelector objc.SEL, argument objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("makeObjectsPerformSelector:withObject:"), aSelector, argument)
}

// Returns the object located at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/object(at:)
func (a_ Array) ObjectAtIndex(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("objectAtIndex:"), index)
	return rv
}

// Returns an enumerator object that lets you access each object in the array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/objectEnumerator()
func (a_ Array) ObjectEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("objectEnumerator"))
	return rv
}

// Returns an array containing the objects in the array at the indexes specified by a given index set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/objects(at:)
func (a_ Array) ObjectsAtIndexes(indexes unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("objectsAtIndexes:"), indexes)
	return rv
}

// Returns an array containing all the pathname elements in the receiving array that have filename extensions from a given array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/pathsMatchingExtensions(_:)
func (a_ Array) PathsMatchingExtensions(filterTypes unsafe.Pointer) []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("pathsMatchingExtensions:"), filterTypes)
	return rv
}

// Raises an exception.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/removeObserver(_:forKeyPath:)
func (a_ Array) RemoveObserverForKeyPath(observer unsafe.Pointer, keyPath string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeObserver:forKeyPath:"), observer, objc.String(keyPath))
}

// Raises an exception.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/removeObserver(_:forKeyPath:context:)
func (a_ Array) RemoveObserverForKeyPathContext(observer unsafe.Pointer, keyPath string, context unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeObserver:forKeyPath:context:"), observer, objc.String(keyPath), context)
}

// Removes from all key value observer notifications associated with the specified relative to the array’s objects at .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/removeObserver(_:fromObjectsAt:forKeyPath:)
func (a_ Array) RemoveObserverFromObjectsAtIndexesForKeyPath(observer unsafe.Pointer, indexes unsafe.Pointer, keyPath string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeObserver:fromObjectsAtIndexes:forKeyPath:"), observer, indexes, objc.String(keyPath))
}

// Raises an exception.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/removeObserver(_:fromObjectsAt:forKeyPath:context:)
func (a_ Array) RemoveObserverFromObjectsAtIndexesForKeyPathContext(observer unsafe.Pointer, indexes unsafe.Pointer, keyPath string, context unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeObserver:fromObjectsAtIndexes:forKeyPath:context:"), observer, indexes, objc.String(keyPath), context)
}

// Returns an enumerator object that lets you access each object in the array, in reverse order.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/reverseObjectEnumerator()
func (a_ Array) ReverseObjectEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("reverseObjectEnumerator"))
	return rv
}

// Invokes on each of the array’s items using the specified and .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/setValue(_:forKey:)
func (a_ Array) SetValueForKey(value objc.ID, key string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setValue:forKey:"), value, objc.String(key))
}

// Returns a new array that lists this array’s elements in a random order.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/shuffled()
func (a_ Array) ShuffledArray() []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("shuffledArray"))
	return rv
}

// Returns a new array that lists this array’s elements in a random order, using the specified random source.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/shuffled(using:)
func (a_ Array) ShuffledArrayWithRandomSource(randomSource unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("shuffledArrayWithRandomSource:"), randomSource)
	return rv
}

// Returns a new array that lists the receiving array’s elements in ascending order as defined by the comparison function .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/sortedArray(_:context:)
func (a_ Array) SortedArrayUsingFunctionContext(comparator unsafe.Pointer, context unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("sortedArrayUsingFunction:context:"), comparator, context)
	return rv
}

// Returns a new array that lists the receiving array’s elements in ascending order as defined by the comparison function .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/sortedArray(_:context:hint:)
func (a_ Array) SortedArrayUsingFunctionContextHint(comparator unsafe.Pointer, context unsafe.Pointer, hint unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("sortedArrayUsingFunction:context:hint:"), comparator, context, hint)
	return rv
}

// Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given block.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/sortedArray(comparator:)
func (a_ Array) SortedArrayUsingComparator(cmptr unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("sortedArrayUsingComparator:"), cmptr)
	return rv
}

// Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given block.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/sortedArray(options:usingComparator:)
func (a_ Array) SortedArrayWithOptionsUsingComparator(opts unsafe.Pointer, cmptr unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("sortedArrayWithOptions:usingComparator:"), opts, cmptr)
	return rv
}

// Returns a copy of the receiving array sorted as specified by a given array of sort descriptors.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/sortedArray(using:)-82wi1
func (a_ Array) SortedArrayUsingDescriptors(sortDescriptors unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("sortedArrayUsingDescriptors:"), sortDescriptors)
	return rv
}

// Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given selector.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/sortedArray(using:)-9nhh9
func (a_ Array) SortedArrayUsingSelector(comparator objc.SEL) []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("sortedArrayUsingSelector:"), comparator)
	return rv
}

// Returns a new array containing the receiving array’s elements that fall within the limits specified by a given range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/subarray(with:)
func (a_ Array) SubarrayWithRange(range_ Range) []objc.ID {
	rv := objc.Send[[]objc.ID](a_.ID, objc.Sel("subarrayWithRange:"), range_)
	return rv
}

// Returns the object at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/subscript(_:)
func (a_ Array) ObjectAtIndexedSubscript(idx uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("objectAtIndexedSubscript:"), idx)
	return rv
}

// Returns an array containing the results of invoking using on each of the array’s objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/value(forKey:)
func (a_ Array) ValueForKey(key string) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("valueForKey:"), objc.String(key))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/write(to:)
func (a_ Array) WriteToURLError(url unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("writeToURL:error:"), url, error_)
	return rv
}

// Writes the contents of the array to the location specified by a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/write(to:atomically:)
func (a_ Array) WriteToURLAtomically(url unsafe.Pointer, atomically bool) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("writeToURL:atomically:"), url, atomically)
	return rv
}

// Writes the contents of the array to a file at a given path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/write(toFile:atomically:)
func (a_ Array) WriteToFileAtomically(path string, useAuxiliaryFile bool) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("writeToFile:atomically:"), objc.String(path), useAuxiliaryFile)
	return rv
}

// The number of objects in the array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/count
func (a_ Array) Count() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("count"))
	return rv
}

// A string that represents the contents of the array, formatted as a property list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/description
func (a_ Array) Description() string {
	rv := objc.Send[string](a_.ID, objc.Sel("description"))
	return rv
}

// The first object in the array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/firstObject
func (a_ Array) FirstObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("firstObject"))
	return rv
}

// The last object in the array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/lastObject
func (a_ Array) LastObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("lastObject"))
	return rv
}

// Analyzes the array and returns a “hint” that speeds the sorting of the array when the hint is supplied to .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/sortedArrayHint
func (a_ Array) SortedArrayHint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("sortedArrayHint"))
	return rv
}



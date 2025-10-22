// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableArray] class.
var (
	MutableArrayClass     _MutableArrayClass
	MutableArrayClassOnce sync.Once
)

func getMutableArrayClass() _MutableArrayClass {
	MutableArrayClassOnce.Do(func() {
		MutableArrayClass = _MutableArrayClass{objc.GetClass("NSMutableArray")}
	})
	return MutableArrayClass
}

type _MutableArrayClass struct {
	class objc.Class
}

// An interface definition for the [MutableArray] class.
type IMutableArray interface {
	IArray
	AddObject(anObject unsafe.Pointer)
	FilterUsingPredicate(predicate IPredicate)
	InsertObjectAtIndex(anObject unsafe.Pointer, index uint)
	RemoveObject(anObject unsafe.Pointer)
	RemoveAllObjects()
	RemoveLastObject()
	RemoveObjectIdenticalTo(anObject unsafe.Pointer)
	RemoveObjectIdenticalToInRange(anObject unsafe.Pointer, range_ IRange)
	RemoveObjectsAtIndexes(indexes IIndexSet)
	RemoveObjectsFromIndicesNumIndices(indices unsafe.Pointer, cnt uint)
	ReplaceObjectAtIndexWithObject(index uint, anObject unsafe.Pointer)
	ReplaceObjectsAtIndexesWithObjects(indexes IIndexSet, objects []objc.ID)
	ReplaceObjectsInRangeWithObjectsFromArray(range_ IRange, otherArray []objc.ID)
	SetObjectAtIndexedSubscript(obj unsafe.Pointer, idx uint)
	SortUsingFunctionContext(compare unsafe.Pointer, context unsafe.Pointer)
	SortUsingComparator(cmptr unsafe.Pointer)
	SortUsingDescriptors(sortDescriptors []SortDescriptor)
}

// A dynamic ordered collection of objects.
//
// You can use this type in Swift instead of an variable in cases that require reference semantics. The class declares the programmatic interface to objects that manage a modifiable array of objects. This class adds insertion and deletion operations to the basic array-handling behavior inherited from . NSMutableArray is “toll-free bridged” with its Core Foundation counterpart, . See for more information.


// A dynamic ordered collection of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray

type MutableArray struct {
	Array
}

// MutableArrayFrom constructs a [MutableArray] from an unsafe.Pointer.
//
// A dynamic ordered collection of objects.
func MutableArrayFrom(ptr unsafe.Pointer) MutableArray {
	return MutableArray{
		Array: ArrayFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableArrayClass) Alloc() MutableArray {
	rv := objc.Send[MutableArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableArrayClass) New() MutableArray {
	rv := objc.Send[MutableArray](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableArray) Init() MutableArray {
	rv := objc.Send[MutableArray](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableArray) Autorelease() MutableArray {
	rv := objc.Send[MutableArray](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableArray creates a new MutableArray instance.
func NewMutableArray() MutableArray {
	return getMutableArrayClass().New()
}




// Initializes a newly allocated mutable array with the contents of the file specified by a given path
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/initWithContentsOfFile:

func NewMutableArrayWithContentsOfFile(path string) MutableArray {
	instance := getMutableArrayClass().Alloc()
	rv := objc.Send[MutableArray](instance.ID, objc.Sel("initWithContentsOfFile:"), objc.String(path))
	rv.Autorelease()
	return rv
}



// Initialized a newly allocated mutable array with the contents of the location specified by a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/initWithContentsOfURL:

func NewMutableArrayWithContentsOfURL(url IURL) MutableArray {
	instance := getMutableArrayClass().Alloc()
	rv := objc.Send[MutableArray](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}



// Creates and returns a mutable array containing the contents specified by a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/init(contentsOfURL:)

func (mc _MutableArrayClass) ArrayWithContentsOfURL(url IURL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("arrayWithContentsOfURL:"), url)
	return rv
}



// Inserts a given object at the end of the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/add(_:)

func (m_ MutableArray) AddObject(anObject unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addObject:"), anObject)
}



// Evaluates a given predicate against the array’s content and leaves only objects that match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/filter(using:)

func (m_ MutableArray) FilterUsingPredicate(predicate IPredicate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("filterUsingPredicate:"), predicate)
}



// Inserts a given object into the array’s contents at a given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/insert(_:at:)-5dbx5

func (m_ MutableArray) InsertObjectAtIndex(anObject unsafe.Pointer, index uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertObject:atIndex:"), anObject, index)
}



// Removes all occurrences in the array of a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/remove(_:)

func (m_ MutableArray) RemoveObject(anObject unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObject:"), anObject)
}



// Empties the array of all its elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeAllObjects()

func (m_ MutableArray) RemoveAllObjects() {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAllObjects"))
}



// Removes the object with the highest-valued index in the array
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeLastObject()

func (m_ MutableArray) RemoveLastObject() {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeLastObject"))
}



// Removes all occurrences of a given object in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObject(identicalTo:)

func (m_ MutableArray) RemoveObjectIdenticalTo(anObject unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectIdenticalTo:"), anObject)
}



// Removes all occurrences of within the specified range in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObject(identicalTo:in:)

func (m_ MutableArray) RemoveObjectIdenticalToInRange(anObject unsafe.Pointer, range_ IRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectIdenticalTo:inRange:"), anObject, range_)
}



// Removes the objects at the specified indexes from the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObjects(at:)

func (m_ MutableArray) RemoveObjectsAtIndexes(indexes IIndexSet) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectsAtIndexes:"), indexes)
}



// Removes the specified number of objects from the array, beginning at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObjects(fromIndices:numIndices:)

func (m_ MutableArray) RemoveObjectsFromIndicesNumIndices(indices unsafe.Pointer, cnt uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectsFromIndices:numIndices:"), indices, cnt)
}



// Replaces the object at with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/replaceObject(at:with:)

func (m_ MutableArray) ReplaceObjectAtIndexWithObject(index uint, anObject unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceObjectAtIndex:withObject:"), index, anObject)
}



// Replaces the objects in the receiving array at locations specified with the objects from a given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/replaceObjects(at:with:)

func (m_ MutableArray) ReplaceObjectsAtIndexesWithObjects(indexes IIndexSet, objects []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceObjectsAtIndexes:withObjects:"), indexes, objects)
}



// Replaces the objects in the receiving array specified by a given range with all of the objects from a given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/replaceObjects(in:withObjectsFrom:)

func (m_ MutableArray) ReplaceObjectsInRangeWithObjectsFromArray(range_ IRange, otherArray []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceObjectsInRange:withObjectsFromArray:"), range_, otherArray)
}



// Replaces the object at the index with the new object, possibly adding the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/setObject:atIndexedSubscript:

func (m_ MutableArray) SetObjectAtIndexedSubscript(obj unsafe.Pointer, idx uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:atIndexedSubscript:"), obj, idx)
}



// Sorts the receiver in ascending order as defined by the comparison function .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(_:context:)

func (m_ MutableArray) SortUsingFunctionContext(compare unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingFunction:context:"), compare, context)
}



// Sorts the receiver in ascending order using the comparison method specified by a given block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(comparator:)

func (m_ MutableArray) SortUsingComparator(cmptr unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingComparator:"), cmptr)
}



// Sorts the receiver using a given array of sort descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(using:)-4eh07

func (m_ MutableArray) SortUsingDescriptors(sortDescriptors []SortDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingDescriptors:"), sortDescriptors)
}



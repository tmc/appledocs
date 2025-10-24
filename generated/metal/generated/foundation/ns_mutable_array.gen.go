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
	// properties:
	// methods:
	AddObject(anObject unsafe.Pointer)
	AddObjectsFromArray(otherArray []objc.ID /* already interface */)
	ApplyDifference(difference unsafe.Pointer)
	ExchangeObjectAtIndexWithObjectAtIndex(idx1 uint /* primitive/slice/pointer. */, idx2 uint /* primitive/slice/pointer. */)
	FilterUsingPredicate(predicate IPredicate)
	InsertObjectAtIndex(anObject unsafe.Pointer, index uint /* primitive/slice/pointer. */)
	InsertObjectsAtIndexes(objects []objc.ID /* already interface */, indexes IIndexSet)
	RemoveObject(anObject unsafe.Pointer)
	RemoveObjectInRange(anObject unsafe.Pointer, range_ Range /* not a class type */)
	RemoveAllObjects()
	RemoveLastObject()
	RemoveObjectAtIndex(index uint /* primitive/slice/pointer. */)
	RemoveObjectIdenticalTo(anObject unsafe.Pointer)
	RemoveObjectIdenticalToInRange(anObject unsafe.Pointer, range_ Range /* not a class type */)
	RemoveObjectsAtIndexes(indexes IIndexSet)
	RemoveObjectsInRange(range_ Range /* not a class type */)
	RemoveObjectsInArray(otherArray []objc.ID /* already interface */)
	ReplaceObjectAtIndexWithObject(index uint /* primitive/slice/pointer. */, anObject unsafe.Pointer)
	ReplaceObjectsAtIndexesWithObjects(indexes IIndexSet, objects []objc.ID /* already interface */)
	ReplaceObjectsInRangeWithObjectsFromArray(range_ Range /* not a class type */, otherArray []objc.ID /* already interface */)
	ReplaceObjectsInRangeWithObjectsFromArrayRange(range_ Range /* not a class type */, otherArray []objc.ID /* already interface */, otherRange Range /* not a class type */)
	SetArray(otherArray []objc.ID /* already interface */)
	SetObjectAtIndexedSubscript(obj unsafe.Pointer, idx uint /* primitive/slice/pointer. */)
	SortUsingFunctionContext(compare unsafe.Pointer, context unsafe.Pointer)
	SortUsingComparator(cmptr Comparator /* not a class type */)
	SortWithOptionsUsingComparator(opts SortOptions, cmptr Comparator /* not a class type */)
	SortUsingDescriptors(sortDescriptors []SortDescriptor /* primitive/slice/pointer. */)
	SortUsingSelector(comparator objc.SEL)
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



// Returns an array, initialized with enough memory to initially hold a given number of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/init(capacity:)
func NewMutableArrayWithCapacity(numItems uint /* primitive/slice/pointer. */) MutableArray {
	instance := getMutableArrayClass().Alloc()
	rv := objc.Send[MutableArray](instance.ID, objc.Sel("initWithCapacity:"), numItems)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/init(coder:)
func NewMutableArrayWithCoder(coder ICoder) MutableArray {
	instance := getMutableArrayClass().Alloc()
	rv := objc.Send[MutableArray](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated mutable array with the contents of the file specified by a given path
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/initWithContentsOfFile:
func NewMutableArrayWithContentsOfFile(path IString) MutableArray {
	instance := getMutableArrayClass().Alloc()
	rv := objc.Send[MutableArray](instance.ID, objc.Sel("initWithContentsOfFile:"), path)
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



// Creates and returns an object with enough allocated memory to initially hold a given number of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/arrayWithCapacity:
func (mc _MutableArrayClass) ArrayWithCapacity(numItems uint /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("arrayWithCapacity:"), numItems)
	return rv
}


// Creates and returns a mutable array containing the contents of the file specified by the given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/arrayWithContentsOfFile:
func (mc _MutableArrayClass) ArrayWithContentsOfFile(path IString) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("arrayWithContentsOfFile:"), path)
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


// Adds the objects contained in another given array to the end of the receiving array’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/addObjects(from:)
func (m_ MutableArray) AddObjectsFromArray(otherArray []objc.ID /* already interface */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addObjectsFromArray:"), otherArray)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/applyDifference:
func (m_ MutableArray) ApplyDifference(difference unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("applyDifference:"), difference)
}


// Exchanges the objects in the array at given indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/exchangeObject(at:withObjectAt:)
func (m_ MutableArray) ExchangeObjectAtIndexWithObjectAtIndex(idx1 uint /* primitive/slice/pointer. */, idx2 uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("exchangeObjectAtIndex:withObjectAtIndex:"), idx1, idx2)
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
func (m_ MutableArray) InsertObjectAtIndex(anObject unsafe.Pointer, index uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertObject:atIndex:"), anObject, index)
}


// Inserts the objects in the provided array into the receiving array at the specified indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/insert(_:at:)-73pln
func (m_ MutableArray) InsertObjectsAtIndexes(objects []objc.ID /* already interface */, indexes IIndexSet) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertObjects:atIndexes:"), objects, indexes)
}


// Removes all occurrences in the array of a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/remove(_:)
func (m_ MutableArray) RemoveObject(anObject unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObject:"), anObject)
}


// Removes all occurrences within a specified range in the array of a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/remove(_:in:)
func (m_ MutableArray) RemoveObjectInRange(anObject unsafe.Pointer, range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObject:inRange:"), anObject, range_)
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


// Removes the object at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObject(at:)
func (m_ MutableArray) RemoveObjectAtIndex(index uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectAtIndex:"), index)
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
func (m_ MutableArray) RemoveObjectIdenticalToInRange(anObject unsafe.Pointer, range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectIdenticalTo:inRange:"), anObject, range_)
}


// Removes the objects at the specified indexes from the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObjects(at:)
func (m_ MutableArray) RemoveObjectsAtIndexes(indexes IIndexSet) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectsAtIndexes:"), indexes)
}


// Removes from the array each of the objects within a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObjects(in:)-1udmn
func (m_ MutableArray) RemoveObjectsInRange(range_ Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectsInRange:"), range_)
}


// Removes from the receiving array the objects in another given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObjects(in:)-4yb26
func (m_ MutableArray) RemoveObjectsInArray(otherArray []objc.ID /* already interface */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectsInArray:"), otherArray)
}


// Replaces the object at with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/replaceObject(at:with:)
func (m_ MutableArray) ReplaceObjectAtIndexWithObject(index uint /* primitive/slice/pointer. */, anObject unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceObjectAtIndex:withObject:"), index, anObject)
}


// Replaces the objects in the receiving array at locations specified with the objects from a given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/replaceObjects(at:with:)
func (m_ MutableArray) ReplaceObjectsAtIndexesWithObjects(indexes IIndexSet, objects []objc.ID /* already interface */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceObjectsAtIndexes:withObjects:"), indexes, objects)
}


// Replaces the objects in the receiving array specified by a given range with all of the objects from a given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/replaceObjects(in:withObjectsFrom:)
func (m_ MutableArray) ReplaceObjectsInRangeWithObjectsFromArray(range_ Range /* not a class type */, otherArray []objc.ID /* already interface */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceObjectsInRange:withObjectsFromArray:"), range_, otherArray)
}


// Replaces the objects in the receiving array specified by one given range with the objects in another array specified by another range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/replaceObjects(in:withObjectsFrom:range:)
func (m_ MutableArray) ReplaceObjectsInRangeWithObjectsFromArrayRange(range_ Range /* not a class type */, otherArray []objc.ID /* already interface */, otherRange Range /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceObjectsInRange:withObjectsFromArray:range:"), range_, otherArray, otherRange)
}


// Sets the receiving array’s elements to those in another given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/setArray(_:)
func (m_ MutableArray) SetArray(otherArray []objc.ID /* already interface */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArray:"), otherArray)
}


// Replaces the object at the index with the new object, possibly adding the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/setObject:atIndexedSubscript:
func (m_ MutableArray) SetObjectAtIndexedSubscript(obj unsafe.Pointer, idx uint /* primitive/slice/pointer. */) {
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
func (m_ MutableArray) SortUsingComparator(cmptr Comparator /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingComparator:"), cmptr)
}


// Sorts the receiver in ascending order using the specified options and the comparison method specified by a given block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(options:usingComparator:)
func (m_ MutableArray) SortWithOptionsUsingComparator(opts SortOptions, cmptr Comparator /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortWithOptions:usingComparator:"), opts, cmptr)
}


// Sorts the receiver using a given array of sort descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(using:)-4eh07
func (m_ MutableArray) SortUsingDescriptors(sortDescriptors []SortDescriptor /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingDescriptors:"), sortDescriptors)
}


// Sorts the receiver in ascending order, as determined by the comparison method specified by a given selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(using:)-537vs
func (m_ MutableArray) SortUsingSelector(comparator objc.SEL) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingSelector:"), comparator)
}



// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableOrderedSet] class.
var (
	MutableOrderedSetClass     _MutableOrderedSetClass
	MutableOrderedSetClassOnce sync.Once
)

func getMutableOrderedSetClass() _MutableOrderedSetClass {
	MutableOrderedSetClassOnce.Do(func() {
		MutableOrderedSetClass = _MutableOrderedSetClass{objc.GetClass("NSMutableOrderedSet")}
	})
	return MutableOrderedSetClass
}

type _MutableOrderedSetClass struct {
	class objc.Class
}

// An interface definition for the [MutableOrderedSet] class.
type IMutableOrderedSet interface {
	IOrderedSet
	// properties:
	// methods:
	AddObject(object unsafe.Pointer)
	AddObjectsCount(objects []unsafe.Pointer, count uint)
	AddObjectsFromArray(array []objc.ID)
	ApplyDifference(difference unsafe.Pointer)
	ExchangeObjectAtIndexWithObjectAtIndex(idx1 uint, idx2 uint)
	FilterUsingPredicate(p IPredicate)
	InsertObjectsAtIndexes(objects []objc.ID, indexes IIndexSet)
	InsertObjectAtIndex(object unsafe.Pointer, idx uint)
	IntersectOrderedSet(other unsafe.Pointer)
	IntersectSet(other unsafe.Pointer)
	MinusOrderedSet(other unsafe.Pointer)
	MinusSet(other unsafe.Pointer)
	MoveObjectsAtIndexesToIndex(indexes IIndexSet, idx uint)
	RemoveObject(object unsafe.Pointer)
	RemoveAllObjects()
	RemoveObjectAtIndex(idx uint)
	RemoveObjectsAtIndexes(indexes IIndexSet)
	RemoveObjectsInArray(array []objc.ID)
	RemoveObjectsInRange(range_ objc.IObject /* cross-framework: Range */)
	ReplaceObjectAtIndexWithObject(idx uint, object unsafe.Pointer)
	ReplaceObjectsAtIndexesWithObjects(indexes IIndexSet, objects []objc.ID)
	ReplaceObjectsInRangeWithObjectsCount(range_ objc.IObject /* cross-framework: Range */, objects []unsafe.Pointer, count uint)
	SetObjectAtIndex(obj unsafe.Pointer, idx uint)
	SetObjectAtIndexedSubscript(obj unsafe.Pointer, idx uint)
	SortUsingComparator(cmptr Comparator /* not a class type */)
	SortWithOptionsUsingComparator(opts SortOptions, cmptr Comparator /* not a class type */)
	SortUsingDescriptors(sortDescriptors []ISortDescriptor)
	SortRangeOptionsUsingComparator(range_ objc.IObject /* cross-framework: Range */, opts SortOptions, cmptr Comparator /* not a class type */)
	UnionOrderedSet(other unsafe.Pointer)
	UnionSet(other unsafe.Pointer)
}

// A dynamic, ordered collection of unique objects.
//
// objects are not like C arrays. That is, even though you may specify a size when you create a mutable ordered set, the specified size is regarded as a “hint”; the actual size of the set is still 0. This means that you cannot insert an object at an index greater than the current count of an set. For example, if a set contains two objects, its size is 2, so you can add objects at indices 0, 1, or 2. Index 3 is illegal and out of bounds; if you try to add an object at index 3 (when the size of the array is 2), raises an exception.


// A dynamic, ordered collection of unique objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet
type MutableOrderedSet struct {
	OrderedSet
}

// MutableOrderedSetFrom constructs a [MutableOrderedSet] from an unsafe.Pointer.
//
// A dynamic, ordered collection of unique objects.
func MutableOrderedSetFrom(ptr unsafe.Pointer) MutableOrderedSet {
	return MutableOrderedSet{
		OrderedSet: OrderedSetFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableOrderedSetClass) Alloc() MutableOrderedSet {
	rv := objc.Send[MutableOrderedSet](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableOrderedSetClass) New() MutableOrderedSet {
	rv := objc.Send[MutableOrderedSet](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableOrderedSet) Init() MutableOrderedSet {
	rv := objc.Send[MutableOrderedSet](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableOrderedSet) Autorelease() MutableOrderedSet {
	rv := objc.Send[MutableOrderedSet](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableOrderedSet creates a new MutableOrderedSet instance.
func NewMutableOrderedSet() MutableOrderedSet {
	return getMutableOrderedSetClass().New()
}



// Returns an initialized mutable ordered set with a given initial capacity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/init(capacity:)
func NewMutableOrderedSetWithCapacity(numItems uint) MutableOrderedSet {
	instance := getMutableOrderedSetClass().Alloc()
	rv := objc.Send[MutableOrderedSet](instance.ID, objc.Sel("initWithCapacity:"), numItems)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/init(coder:)
func NewMutableOrderedSetWithCoder(coder ICoder) MutableOrderedSet {
	instance := getMutableOrderedSetClass().Alloc()
	rv := objc.Send[MutableOrderedSet](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Creates and returns an mutable ordered set with a given initial capacity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/orderedSetWithCapacity:
func (mc _MutableOrderedSetClass) OrderedSetWithCapacity(numItems uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("orderedSetWithCapacity:"), numItems)
	return rv
}


// Appends a given object to the end of the mutable ordered set, if it is not already a member.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/add(_:)
func (m_ MutableOrderedSet) AddObject(object unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addObject:"), object)
}


// Appends the given number of objects from a given C array to the end of the mutable ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/add(_:count:)
func (m_ MutableOrderedSet) AddObjectsCount(objects []unsafe.Pointer, count uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addObjects:count:"), objects, count)
}


// Appends to the end of the mutable ordered set each object contained in a given array that is not already a member.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/addObjects(from:)
func (m_ MutableOrderedSet) AddObjectsFromArray(array []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addObjectsFromArray:"), array)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/applyDifference:
func (m_ MutableOrderedSet) ApplyDifference(difference unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("applyDifference:"), difference)
}


// Exchanges the object at the specified index with the object at the other index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/exchangeObject(at:withObjectAt:)
func (m_ MutableOrderedSet) ExchangeObjectAtIndexWithObjectAtIndex(idx1 uint, idx2 uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("exchangeObjectAtIndex:withObjectAtIndex:"), idx1, idx2)
}


// Evaluates a given predicate against the mutable ordered set’s content and leaves only objects that match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/filter(using:)
func (m_ MutableOrderedSet) FilterUsingPredicate(p IPredicate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("filterUsingPredicate:"), p)
}


// Inserts the objects in the array at the specified indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/insert(_:at:)-3ncnm
func (m_ MutableOrderedSet) InsertObjectsAtIndexes(objects []objc.ID, indexes IIndexSet) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertObjects:atIndexes:"), objects, indexes)
}


// Inserts the given object at the specified index of the mutable ordered set, if it is not already a member.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/insert(_:at:)-7qg51
func (m_ MutableOrderedSet) InsertObjectAtIndex(object unsafe.Pointer, idx uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertObject:atIndex:"), object, idx)
}


// Removes from the receiving ordered set each object that isn’t a member of another ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/intersect(_:)
func (m_ MutableOrderedSet) IntersectOrderedSet(other unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("intersectOrderedSet:"), other)
}


// Removes from the receiving ordered set each object that isn’t a member of another set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/intersectSet(_:)
func (m_ MutableOrderedSet) IntersectSet(other unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("intersectSet:"), other)
}


// Removes each object in another given ordered set from the receiving mutable ordered set, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/minus(_:)
func (m_ MutableOrderedSet) MinusOrderedSet(other unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("minusOrderedSet:"), other)
}


// Removes each object in another given set from the receiving mutable ordered set, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/minusSet(_:)
func (m_ MutableOrderedSet) MinusSet(other unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("minusSet:"), other)
}


// Moves the objects at the specified indexes to the new location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/moveObjects(at:to:)
func (m_ MutableOrderedSet) MoveObjectsAtIndexesToIndex(indexes IIndexSet, idx uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("moveObjectsAtIndexes:toIndex:"), indexes, idx)
}


// Removes a given object from the mutable ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/remove(_:)
func (m_ MutableOrderedSet) RemoveObject(object unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObject:"), object)
}


// Removes all the objects from the mutable ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/removeAllObjects()
func (m_ MutableOrderedSet) RemoveAllObjects() {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAllObjects"))
}


// Removes a the object at the specified index from the mutable ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/removeObject(at:)
func (m_ MutableOrderedSet) RemoveObjectAtIndex(idx uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectAtIndex:"), idx)
}


// Removes the objects at the specified indexes from the mutable ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/removeObjects(at:)
func (m_ MutableOrderedSet) RemoveObjectsAtIndexes(indexes IIndexSet) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectsAtIndexes:"), indexes)
}


// Removes the objects in the array from the mutable ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/removeObjects(in:)-8h2kh
func (m_ MutableOrderedSet) RemoveObjectsInArray(array []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectsInArray:"), array)
}


// Removes from the mutable ordered set each of the objects within a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/removeObjects(in:)-9jkis
func (m_ MutableOrderedSet) RemoveObjectsInRange(range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectsInRange:"), range_)
}


// Replaces the object at the specified index with the new object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/replaceObject(at:with:)
func (m_ MutableOrderedSet) ReplaceObjectAtIndexWithObject(idx uint, object unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceObjectAtIndex:withObject:"), idx, object)
}


// Replaces the objects at the specified indexes with the new objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/replaceObjects(at:with:)
func (m_ MutableOrderedSet) ReplaceObjectsAtIndexesWithObjects(indexes IIndexSet, objects []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceObjectsAtIndexes:withObjects:"), indexes, objects)
}


// Replaces the objects in the receiving mutable ordered set at the range with the specified number of objects from a given C array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/replaceObjects(in:with:count:)
func (m_ MutableOrderedSet) ReplaceObjectsInRangeWithObjectsCount(range_ objc.IObject /* cross-framework: Range */, objects []unsafe.Pointer, count uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceObjectsInRange:withObjects:count:"), range_, objects, count)
}


// Appends or replaces the object at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/setObject(_:at:)
func (m_ MutableOrderedSet) SetObjectAtIndex(obj unsafe.Pointer, idx uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:atIndex:"), obj, idx)
}


// Replaces the given object at the specified index of the mutable ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/setObject:atIndexedSubscript:
func (m_ MutableOrderedSet) SetObjectAtIndexedSubscript(obj unsafe.Pointer, idx uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:atIndexedSubscript:"), obj, idx)
}


// Sorts the mutable ordered set using the comparison method specified by the comparator block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/sort(comparator:)
func (m_ MutableOrderedSet) SortUsingComparator(cmptr Comparator /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingComparator:"), cmptr)
}


// Sorts the mutable ordered set using the specified options and the comparison method specified by a given comparator block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/sort(options:usingComparator:)
func (m_ MutableOrderedSet) SortWithOptionsUsingComparator(opts SortOptions, cmptr Comparator /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortWithOptions:usingComparator:"), opts, cmptr)
}


// Sorts the receiving ordered set using a given array of sort descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/sort(using:)
func (m_ MutableOrderedSet) SortUsingDescriptors(sortDescriptors []ISortDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingDescriptors:"), sortDescriptors)
}


// Sorts the specified range of the mutable ordered set using the specified options and the comparison method specified by a given comparator block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/sortRange(_:options:usingComparator:)
func (m_ MutableOrderedSet) SortRangeOptionsUsingComparator(range_ objc.IObject /* cross-framework: Range */, opts SortOptions, cmptr Comparator /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortRange:options:usingComparator:"), range_, opts, cmptr)
}


// Adds each object in another given ordered set to the receiving mutable ordered set, if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/union(_:)
func (m_ MutableOrderedSet) UnionOrderedSet(other unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("unionOrderedSet:"), other)
}


// Adds each object in another given set to the receiving mutable ordered set, if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/unionSet(_:)
func (m_ MutableOrderedSet) UnionSet(other unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("unionSet:"), other)
}



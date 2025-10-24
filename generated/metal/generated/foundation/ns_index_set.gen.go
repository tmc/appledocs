// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [IndexSet] class.
var (
	IndexSetClass     _IndexSetClass
	IndexSetClassOnce sync.Once
)

func getIndexSetClass() _IndexSetClass {
	IndexSetClassOnce.Do(func() {
		IndexSetClass = _IndexSetClass{objc.GetClass("NSIndexSet")}
	})
	return IndexSetClass
}

type _IndexSetClass struct {
	class objc.Class
}

// An interface definition for the [IndexSet] class.
type IIndexSet interface {
	objectivec.IObject
	// properties:
	Count() uint /* primitive/slice/pointer. */
	FirstIndex() uint /* primitive/slice/pointer. */
	LastIndex() uint /* primitive/slice/pointer. */
	// methods:
	ContainsIndexes(indexSet IIndexSet) bool /* primitive/slice/pointer. */
	ContainsIndex(value uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	ContainsIndexesInRange(range_ Range /* not a class type */) bool /* primitive/slice/pointer. */
	CountOfIndexesInRange(range_ Range /* not a class type */) uint /* primitive/slice/pointer. */
	EnumerateIndexesUsingBlock(block unsafe.Pointer)
	EnumerateIndexesInRangeOptionsUsingBlock(range_ Range /* not a class type */, opts EnumerationOptions, block unsafe.Pointer)
	EnumerateIndexesWithOptionsUsingBlock(opts EnumerationOptions, block unsafe.Pointer)
	EnumerateRangesUsingBlock(block unsafe.Pointer)
	EnumerateRangesInRangeOptionsUsingBlock(range_ Range /* not a class type */, opts EnumerationOptions, block unsafe.Pointer)
	EnumerateRangesWithOptionsUsingBlock(opts EnumerationOptions, block unsafe.Pointer)
	GetIndexesMaxCountInIndexRange(indexBuffer UInteger /* not a class type */, bufferSize uint /* primitive/slice/pointer. */, range_ objc.IObject /* cross-framework: RangePointer */) uint /* primitive/slice/pointer. */
	IndexInRangeOptionsPassingTest(range_ Range /* not a class type */, opts EnumerationOptions, predicate unsafe.Pointer) uint /* primitive/slice/pointer. */
	IndexWithOptionsPassingTest(opts EnumerationOptions, predicate unsafe.Pointer) uint /* primitive/slice/pointer. */
	IndexPassingTest(predicate unsafe.Pointer) uint /* primitive/slice/pointer. */
	IndexGreaterThanIndex(value uint /* primitive/slice/pointer. */) uint /* primitive/slice/pointer. */
	IndexGreaterThanOrEqualToIndex(value uint /* primitive/slice/pointer. */) uint /* primitive/slice/pointer. */
	IndexLessThanIndex(value uint /* primitive/slice/pointer. */) uint /* primitive/slice/pointer. */
	IndexLessThanOrEqualToIndex(value uint /* primitive/slice/pointer. */) uint /* primitive/slice/pointer. */
	IndexesInRangeOptionsPassingTest(range_ Range /* not a class type */, opts EnumerationOptions, predicate unsafe.Pointer) IIndexSet
	IndexesWithOptionsPassingTest(opts EnumerationOptions, predicate unsafe.Pointer) IIndexSet
	IndexesPassingTest(predicate unsafe.Pointer) IIndexSet
	IntersectsIndexesInRange(range_ Range /* not a class type */) bool /* primitive/slice/pointer. */
	IsEqualToIndexSet(indexSet IIndexSet) bool /* primitive/slice/pointer. */
}

// An immutable collection of unique integer values that represent indexes in another collection.
//
// In Swift, this type bridges to ; use when you need reference semantics or other Foundation-specific behavior. The class represents an immutable collection of unique unsigned integers, known as because of the way they are used. This collection is referred to as an . Indexes must be in the range . You use index sets in your code to store indexes into some other data structure. For example, given an object, you could use an index set to identify a subset of objects in that array. You should not use index sets to store an arbitrary collection of integer values because index sets store indexes as sorted ranges. This makes them more efficient than storing a collection of individual integers. It also means that each index value can only appear once in the index set. The designated initializers of the class are: , , and . You must not subclass the class. The mutable subclass of is .


// An immutable collection of unique integer values that represent indexes in another collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet
type IndexSet struct {
	objectivec.Object
}

// IndexSetFrom constructs a [IndexSet] from an unsafe.Pointer.
//
// An immutable collection of unique integer values that represent indexes in another collection.
func IndexSetFrom(ptr unsafe.Pointer) IndexSet {
	return IndexSet{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _IndexSetClass) Alloc() IndexSet {
	rv := objc.Send[IndexSet](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IndexSetClass) New() IndexSet {
	rv := objc.Send[IndexSet](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IndexSet) Init() IndexSet {
	rv := objc.Send[IndexSet](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IndexSet) Autorelease() IndexSet {
	rv := objc.Send[IndexSet](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIndexSet creates a new IndexSet instance.
func NewIndexSet() IndexSet {
	return getIndexSetClass().New()
}



// Initializes an allocated object with an index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/init(index:)
func NewIndexSetWithIndex(value uint /* primitive/slice/pointer. */) IndexSet {
	instance := getIndexSetClass().Alloc()
	rv := objc.Send[IndexSet](instance.ID, objc.Sel("initWithIndex:"), value)
	rv.Autorelease()
	return rv
}


// Initializes an allocated object with an index set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/init(indexSet:)
func NewIndexSetWithIndexSet(indexSet IIndexSet) IndexSet {
	instance := getIndexSetClass().Alloc()
	rv := objc.Send[IndexSet](instance.ID, objc.Sel("initWithIndexSet:"), indexSet)
	rv.Autorelease()
	return rv
}


// Initializes an allocated object with an index range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/init(indexesIn:)
func NewIndexSetWithIndexesInRange(range_ Range /* not a class type */) IndexSet {
	instance := getIndexSetClass().Alloc()
	rv := objc.Send[IndexSet](instance.ID, objc.Sel("initWithIndexesInRange:"), range_)
	rv.Autorelease()
	return rv
}



// Creates an empty index set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexSet
func (ic _IndexSetClass) IndexSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("indexSet"))
	return rv
}


// Creates an index set with an index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexSetWithIndex:
func (ic _IndexSetClass) IndexSetWithIndex(value uint /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("indexSetWithIndex:"), value)
	return rv
}


// Creates an index set with an index range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexSetWithIndexesInRange:
func (ic _IndexSetClass) IndexSetWithIndexesInRange(range_ Range /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("indexSetWithIndexesInRange:"), range_)
	return rv
}


// Indicates whether the receiving index set contains a superset of the indexes in another index set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/contains(_:)-5j2kh
func (i_ IndexSet) ContainsIndexes(indexSet IIndexSet) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("containsIndexes:"), indexSet)
	return rv
}


// Indicates whether the index set contains a specific index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/contains(_:)-bb19
func (i_ IndexSet) ContainsIndex(value uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("containsIndex:"), value)
	return rv
}


// Indicates whether the index set contains the indexes represented by an index range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/contains(in:)
func (i_ IndexSet) ContainsIndexesInRange(range_ Range /* not a class type */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("containsIndexesInRange:"), range_)
	return rv
}


// Returns the number of indexes in the index set that are members of a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/countOfIndexes(in:)
func (i_ IndexSet) CountOfIndexesInRange(range_ Range /* not a class type */) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("countOfIndexesInRange:"), range_)
	return rv
}


// Executes a given Block using each object in the index set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/enumerate(_:)
func (i_ IndexSet) EnumerateIndexesUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("enumerateIndexesUsingBlock:"), block)
}


// Executes a given Block using the indexes in the specified range, using the specified enumeration options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/enumerate(in:options:using:)
func (i_ IndexSet) EnumerateIndexesInRangeOptionsUsingBlock(range_ Range /* not a class type */, opts EnumerationOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("enumerateIndexesInRange:options:usingBlock:"), range_, opts, block)
}


// Executes a given Block over the index set’s indexes, using the specified enumeration options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/enumerate(options:using:)
func (i_ IndexSet) EnumerateIndexesWithOptionsUsingBlock(opts EnumerationOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("enumerateIndexesWithOptions:usingBlock:"), opts, block)
}


// Executes a given block using each object in the index set, in the specified ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/enumerateRanges(_:)
func (i_ IndexSet) EnumerateRangesUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("enumerateRangesUsingBlock:"), block)
}


// Enumerates over the ranges in the range of objects using the block
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/enumerateRanges(in:options:using:)
func (i_ IndexSet) EnumerateRangesInRangeOptionsUsingBlock(range_ Range /* not a class type */, opts EnumerationOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("enumerateRangesInRange:options:usingBlock:"), range_, opts, block)
}


// Executes a given block using each object in the index set, in the specified ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/enumerateRanges(options:using:)
func (i_ IndexSet) EnumerateRangesWithOptionsUsingBlock(opts EnumerationOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("enumerateRangesWithOptions:usingBlock:"), opts, block)
}


// The index set fills an index buffer with the indexes contained both in the index set and in an index range, returning the number of indexes copied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/getIndexes(_:maxCount:inIndexRange:)
func (i_ IndexSet) GetIndexesMaxCountInIndexRange(indexBuffer UInteger /* not a class type */, bufferSize uint /* primitive/slice/pointer. */, range_ objc.IObject /* cross-framework: RangePointer */) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("getIndexes:maxCount:inIndexRange:"), indexBuffer, bufferSize, range_)
	return rv
}


// Returns the index of the first object in the specified range that passes the predicate Block test.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/index(in:options:passingTest:)
func (i_ IndexSet) IndexInRangeOptionsPassingTest(range_ Range /* not a class type */, opts EnumerationOptions, predicate unsafe.Pointer) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("indexInRange:options:passingTest:"), range_, opts, predicate)
	return rv
}


// Returns the index of the first object that passes the predicate Block test using the specified enumeration options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/index(options:passingTest:)
func (i_ IndexSet) IndexWithOptionsPassingTest(opts EnumerationOptions, predicate unsafe.Pointer) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("indexWithOptions:passingTest:"), opts, predicate)
	return rv
}


// Returns the index of the first object that passes the predicate Block test.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/index(passingTest:)
func (i_ IndexSet) IndexPassingTest(predicate unsafe.Pointer) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("indexPassingTest:"), predicate)
	return rv
}


// Returns either the closest index in the index set that is greater than a specific index or the not-found indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexGreaterThanIndex(_:)
func (i_ IndexSet) IndexGreaterThanIndex(value uint /* primitive/slice/pointer. */) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("indexGreaterThanIndex:"), value)
	return rv
}


// Returns either the closest index in the index set that is greater than or equal to a specific index or the not-found indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexGreaterThanOrEqual(to:)
func (i_ IndexSet) IndexGreaterThanOrEqualToIndex(value uint /* primitive/slice/pointer. */) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("indexGreaterThanOrEqualToIndex:"), value)
	return rv
}


// Returns either the closest index in the index set that is less than a specific index or the not-found indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexLessThanIndex(_:)
func (i_ IndexSet) IndexLessThanIndex(value uint /* primitive/slice/pointer. */) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("indexLessThanIndex:"), value)
	return rv
}


// Returns either the closest index in the index set that is less than or equal to a specific index or the not-found indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexLessThanOrEqual(to:)
func (i_ IndexSet) IndexLessThanOrEqualToIndex(value uint /* primitive/slice/pointer. */) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("indexLessThanOrEqualToIndex:"), value)
	return rv
}


// Returns an containing the receiving index set’s objects in the specified range that pass the Block test.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexes(in:options:passingTest:)
func (i_ IndexSet) IndexesInRangeOptionsPassingTest(range_ Range /* not a class type */, opts EnumerationOptions, predicate unsafe.Pointer) IIndexSet {
	rv := objc.Send[IndexSet](i_.ID, objc.Sel("indexesInRange:options:passingTest:"), range_, opts, predicate)
	return rv
}


// Returns an containing the receiving index set’s objects that pass the Block test using the specified enumeration options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexes(options:passingTest:)
func (i_ IndexSet) IndexesWithOptionsPassingTest(opts EnumerationOptions, predicate unsafe.Pointer) IIndexSet {
	rv := objc.Send[IndexSet](i_.ID, objc.Sel("indexesWithOptions:passingTest:"), opts, predicate)
	return rv
}


// Returns an containing the receiving index set’s objects that pass the Block test.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/indexes(passingTest:)
func (i_ IndexSet) IndexesPassingTest(predicate unsafe.Pointer) IIndexSet {
	rv := objc.Send[IndexSet](i_.ID, objc.Sel("indexesPassingTest:"), predicate)
	return rv
}


// Indicates whether the index set contains any of the indexes in a range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/intersects(in:)
func (i_ IndexSet) IntersectsIndexesInRange(range_ Range /* not a class type */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("intersectsIndexesInRange:"), range_)
	return rv
}


// Indicates whether the indexes in the receiving index set are the same indexes contained in another index set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/isEqual(to:)
func (i_ IndexSet) IsEqualToIndexSet(indexSet IIndexSet) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("isEqualToIndexSet:"), indexSet)
	return rv
}


// The number of indexes in the index set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/count
func (i_ IndexSet) Count() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("count"))
	return rv
}


// The first index in the index set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/firstIndex
func (i_ IndexSet) FirstIndex() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("firstIndex"))
	return rv
}


// The last index in the index set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/lastIndex
func (i_ IndexSet) LastIndex() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("lastIndex"))
	return rv
}



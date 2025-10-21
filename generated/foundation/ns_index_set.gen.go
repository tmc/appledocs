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
	EnumerateIndexesUsingBlock(block unsafe.Pointer)
	EnumerateIndexesInRangeOptionsUsingBlock(range_ Range, opts unsafe.Pointer, block unsafe.Pointer)
	EnumerateIndexesWithOptionsUsingBlock(opts unsafe.Pointer, block unsafe.Pointer)
}

// An immutable collection of unique integer values that represent indexes in another collection.
//
// In Swift, this type bridges to ; use when you need reference semantics or other Foundation-specific behavior. The class represents an immutable collection of unique unsigned integers, known as because of the way they are used. This collection is referred to as an . Indexes must be in the range . You use index sets in your code to store indexes into some other data structure. For example, given an object, you could use an index set to identify a subset of objects in that array. You should not use index sets to store an arbitrary collection of integer values because index sets store indexes as sorted ranges. This makes them more efficient than storing a collection of individual integers. It also means that each index value can only appear once in the index set. The designated initializers of the class are: , , and . You must not subclass the class. The mutable subclass of is .
//
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/init(index:)
func NewIndexSetWithIndex(value uint) IndexSet {
	instance := getIndexSetClass().Alloc()
	rv := objc.Send[IndexSet](instance.ID, objc.Sel("initWithIndex:"), value)
	rv.Autorelease()
	return rv
}


// Executes a given Block using each object in the index set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/enumerate(_:)
func (i_ IndexSet) EnumerateIndexesUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("enumerateIndexesUsingBlock:"), block)
}

// Executes a given Block using the indexes in the specified range, using the specified enumeration options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/enumerate(in:options:using:)
func (i_ IndexSet) EnumerateIndexesInRangeOptionsUsingBlock(range_ Range, opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("enumerateIndexesInRange:options:usingBlock:"), range_, opts, block)
}

// Executes a given Block over the index set’s indexes, using the specified enumeration options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet/enumerate(options:using:)
func (i_ IndexSet) EnumerateIndexesWithOptionsUsingBlock(opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("enumerateIndexesWithOptions:usingBlock:"), opts, block)
}



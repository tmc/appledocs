// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [IndexPath] class.
var (
	IndexPathClass     _IndexPathClass
	IndexPathClassOnce sync.Once
)

func getIndexPathClass() _IndexPathClass {
	IndexPathClassOnce.Do(func() {
		IndexPathClass = _IndexPathClass{objc.GetClass("NSIndexPath")}
	})
	return IndexPathClass
}

type _IndexPathClass struct {
	class objc.Class
}

// An interface definition for the [IndexPath] class.
type IIndexPath interface {
	objectivec.IObject
	// properties:
	Item() int /* primitive/slice/pointer */
	Length() uint /* primitive/slice/pointer */
	Row() int /* primitive/slice/pointer */
	Section() int /* primitive/slice/pointer */
	// methods:
	IndexPathByAddingIndex(index uint /* primitive/slice/pointer */) IIndexPath
	Compare(otherObject IIndexPath) ComparisonResult /* foo */
	GetIndexesRange(indexes UInteger /* foo */, positionRange Range /* foo */)
	IndexAtPosition(position uint /* primitive/slice/pointer */) uint /* primitive/slice/pointer */
	IndexPathByRemovingLastIndex() IIndexPath
}

// A list of indexes that together represent the path to a specific location in a tree of nested arrays.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. Each index in an index path represents the index into an array of children from one node in the tree to another, deeper, node. For example, the index path specifies the path shown in .


// A list of indexes that together represent the path to a specific location in a tree of nested arrays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath
type IndexPath struct {
	objectivec.Object
}

// IndexPathFrom constructs a [IndexPath] from an unsafe.Pointer.
//
// A list of indexes that together represent the path to a specific location in a tree of nested arrays.
func IndexPathFrom(ptr unsafe.Pointer) IndexPath {
	return IndexPath{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _IndexPathClass) Alloc() IndexPath {
	rv := objc.Send[IndexPath](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IndexPathClass) New() IndexPath {
	rv := objc.Send[IndexPath](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IndexPath) Init() IndexPath {
	rv := objc.Send[IndexPath](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IndexPath) Autorelease() IndexPath {
	rv := objc.Send[IndexPath](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIndexPath creates a new IndexPath instance.
func NewIndexPath() IndexPath {
	return getIndexPathClass().New()
}



// Initializes an index path with the indexes of a specific item and section in a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(forItem:inSection:)
func NewIndexPathForItemInSection(item int /* primitive/slice/pointer */, section int /* primitive/slice/pointer */) IndexPath {
	rv := objc.Send[IndexPath](objc.ID(getIndexPathClass().class), objc.Sel("indexPathForItem:inSection:"), item, section)
	return rv
}


// Initializes an index path with the indexes of a specific row and section in a table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(forRow:inSection:)
func NewIndexPathForRowInSection(row int /* primitive/slice/pointer */, section int /* primitive/slice/pointer */) IndexPath {
	rv := objc.Send[IndexPath](objc.ID(getIndexPathClass().class), objc.Sel("indexPathForRow:inSection:"), row, section)
	return rv
}


// Initializes an index path with a single node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(index:)
func NewIndexPathWithIndex(index uint /* primitive/slice/pointer */) IndexPath {
	instance := getIndexPathClass().Alloc()
	rv := objc.Send[IndexPath](instance.ID, objc.Sel("initWithIndex:"), index)
	rv.Autorelease()
	return rv
}


// Initializes an index path with the given nodes and length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(indexes:length:)
func NewIndexPathWithIndexesLength(indexes UInteger [] /* foo */, length uint /* primitive/slice/pointer */) IndexPath {
	instance := getIndexPathClass().Alloc()
	rv := objc.Send[IndexPath](instance.ID, objc.Sel("initWithIndexes:length:"), indexes, length)
	rv.Autorelease()
	return rv
}



// Initializes an index path with the indexes of a specific item and section in a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(forItem:inSection:)
func (ic _IndexPathClass) IndexPathForItemInSection(item int /* primitive/slice/pointer */, section int /* primitive/slice/pointer */) IIndexPath {
	rv := objc.Send[IndexPath](objc.ID(ic.class), objc.Sel("indexPathForItem:inSection:"), item, section)
	return rv
}


// Initializes an index path with the indexes of a specific row and section in a table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(forRow:inSection:)
func (ic _IndexPathClass) IndexPathForRowInSection(row int /* primitive/slice/pointer */, section int /* primitive/slice/pointer */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("indexPathForRow:inSection:"), row, section)
	return rv
}


// Returns an index path containing the nodes in the receiving index path plus another given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/adding(_:)
func (i_ IndexPath) IndexPathByAddingIndex(index uint /* primitive/slice/pointer */) IIndexPath {
	rv := objc.Send[IndexPath](i_.ID, objc.Sel("indexPathByAddingIndex:"), index)
	return rv
}


// Indicates the depth-first traversal order of the receiving index path and another index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/compare(_:)
func (i_ IndexPath) Compare(otherObject IIndexPath) ComparisonResult /* foo */ {
	rv := objc.Send[ComparisonResult](i_.ID, objc.Sel("compare:"), otherObject)
	return rv
}


// Copies the indexes stored in the index path from the positions specified by the position range into the specified indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/getIndexes(_:range:)
func (i_ IndexPath) GetIndexesRange(indexes UInteger /* foo */, positionRange Range /* foo */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("getIndexes:range:"), indexes, positionRange)
}


// Provides the value at a particular node in the index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/index(atPosition:)
func (i_ IndexPath) IndexAtPosition(position uint /* primitive/slice/pointer */) uint /* primitive/slice/pointer */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("indexAtPosition:"), position)
	return rv
}


// Returns an index path with the nodes in the receiving index path, excluding the last one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/removingLastIndex()
func (i_ IndexPath) IndexPathByRemovingLastIndex() IIndexPath {
	rv := objc.Send[IndexPath](i_.ID, objc.Sel("indexPathByRemovingLastIndex"))
	return rv
}


// An index number identifying an item in a section of a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/item
func (i_ IndexPath) Item() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](i_.ID, objc.Sel("item"))
	return rv
}


// The number of nodes in the index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/length
func (i_ IndexPath) Length() uint /* primitive/slice/pointer */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("length"))
	return rv
}


// An index number identifying a row in a section of a table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/row
func (i_ IndexPath) Row() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](i_.ID, objc.Sel("row"))
	return rv
}


// An index number identifying a section in a table view or collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/section
func (i_ IndexPath) Section() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](i_.ID, objc.Sel("section"))
	return rv
}



// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSIndexPath */


/* debug [class_header]: Header for NSIndexPath */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IndexPath */
// An interface definition for the [IndexPath] class.
type IIndexPath interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for IndexPath */
	// properties:
	Item() int
	Length() uint
	Section() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IndexPath */
	// methods:
	IndexPathByAddingIndex(index uint) IIndexPath
	Compare(otherObject IIndexPath) ComparisonResult
	GetIndexesRange(indexes uint, positionRange objc.IObject /* cross-framework: Range */)
	IndexAtPosition(position uint) uint
	IndexPathByRemovingLastIndex() IIndexPath
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IndexPath */
// Alloc allocates a new instance without initialization.
func (ic _IndexPathClass) Alloc() IndexPath {
	rv := objc.Send[IndexPath](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IndexPath */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IndexPath */

// Initializes an index path with the indexes of a specific item and section in a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(forItem:inSection:)
func NewIndexPathForItemInSection(item int, section int) IndexPath {
	rv := objc.Send[IndexPath](objc.ID(getIndexPathClass().class), objc.Sel("indexPathForItem:inSection:"), item, section)
	return rv
}/* debug [class_init_methods/constructor]: NewIndexPathForItemInSection */


// Initializes an index path with the indexes of a specific row and section in a table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(forRow:inSection:)
func NewIndexPathForRowInSection(row int, section int) IndexPath {
	rv := objc.Send[IndexPath](objc.ID(getIndexPathClass().class), objc.Sel("indexPathForRow:inSection:"), row, section)
	return rv
}/* debug [class_init_methods/constructor]: NewIndexPathForRowInSection */


// Initializes an index path with a single node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(index:)
func NewIndexPathWithIndex(index uint) IndexPath {
	instance := getIndexPathClass().Alloc()
	rv := objc.Send[IndexPath](instance.ID, objc.Sel("initWithIndex:"), index)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewIndexPathWithIndex */


// Initializes an index path with the given nodes and length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(indexes:length:)
func NewIndexPathWithIndexesLength(indexes []uint, length uint) IndexPath {
	instance := getIndexPathClass().Alloc()
	rv := objc.Send[IndexPath](instance.ID, objc.Sel("initWithIndexes:length:"), indexes, length)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewIndexPathWithIndexesLength */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IndexPath */

// Creates a one-node index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/indexPathWithIndex:
func (ic _IndexPathClass) IndexPathWithIndex(index uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("indexPathWithIndex:"), index)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IndexPathWithIndex) */


// Creates an index path with one or more nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/indexPathWithIndexes:length:
func (ic _IndexPathClass) IndexPathWithIndexesLength(indexes uint, length uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("indexPathWithIndexes:length:"), indexes, length)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IndexPathWithIndexesLength) */


// Initializes an index path with the indexes of a specific item and section in a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(forItem:inSection:)
func (ic _IndexPathClass) IndexPathForItemInSection(item int, section int) IIndexPath {
	rv := objc.Send[IndexPath](objc.ID(ic.class), objc.Sel("indexPathForItem:inSection:"), item, section)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IndexPathForItemInSection) */


// Initializes an index path with the indexes of a specific row and section in a table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(forRow:inSection:)
func (ic _IndexPathClass) IndexPathForRowInSection(row int, section int) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("indexPathForRow:inSection:"), row, section)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IndexPathForRowInSection) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IndexPath */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IndexPath */

// Returns an index path containing the nodes in the receiving index path plus another given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/adding(_:)
func (i_ IndexPath) IndexPathByAddingIndex(index uint) IIndexPath {
	rv := objc.Send[IndexPath](i_.ID, objc.Sel("indexPathByAddingIndex:"), index)
	return rv
}/* debug [instance_methods/method]: IndexPathByAddingIndex */


// Indicates the depth-first traversal order of the receiving index path and another index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/compare(_:)
func (i_ IndexPath) Compare(otherObject IIndexPath) ComparisonResult {
	rv := objc.Send[ComparisonResult](i_.ID, objc.Sel("compare:"), otherObject)
	return rv
}/* debug [instance_methods/method]: Compare */


// Copies the indexes stored in the index path from the positions specified by the position range into the specified indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/getIndexes(_:range:)
func (i_ IndexPath) GetIndexesRange(indexes uint, positionRange objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("getIndexes:range:"), indexes, positionRange)
}/* debug [instance_methods/method]: GetIndexesRange */


// Provides the value at a particular node in the index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/index(atPosition:)
func (i_ IndexPath) IndexAtPosition(position uint) uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("indexAtPosition:"), position)
	return rv
}/* debug [instance_methods/method]: IndexAtPosition */


// Returns an index path with the nodes in the receiving index path, excluding the last one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/removingLastIndex()
func (i_ IndexPath) IndexPathByRemovingLastIndex() IIndexPath {
	rv := objc.Send[IndexPath](i_.ID, objc.Sel("indexPathByRemovingLastIndex"))
	return rv
}/* debug [instance_methods/method]: IndexPathByRemovingLastIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IndexPath */

// An index number identifying an item in a section of a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/item
func (i_ IndexPath) Item() int {
	rv := objc.Send[int](i_.ID, objc.Sel("item"))
	return rv
}/* debug [instance_properties/getter]: item */


// The number of nodes in the index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/length
func (i_ IndexPath) Length() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */


// An index number identifying a section in a table view or collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/section
func (i_ IndexPath) Section() int {
	rv := objc.Send[int](i_.ID, objc.Sel("section"))
	return rv
}/* debug [instance_properties/getter]: section */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSIndexPath */



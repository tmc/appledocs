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
}

// A list of indexes that together represent the path to a specific location in a tree of nested arrays.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. Each index in an index path represents the index into an array of children from one node in the tree to another, deeper, node. For example, the index path specifies the path shown in .
//
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


// An index number identifying an item in a section of a collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath/item
func (i_ IndexPath) Item() int {
	rv := objc.Send[int](i_.ID, objc.Sel("item"))
	return rv
}


// SetItem sets the value of the item property.
// An index number identifying an item in a section of a collection view.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath/item
func (i_ IndexPath) SetItem(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setItem:"), value)
}

// The number of nodes in the index path.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath/length
func (i_ IndexPath) Length() int {
	rv := objc.Send[int](i_.ID, objc.Sel("length"))
	return rv
}


// SetLength sets the value of the length property.
// The number of nodes in the index path.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath/length
func (i_ IndexPath) SetLength(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLength:"), value)
}

// An index number identifying a row in a section of a table view.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/row
func (i_ IndexPath) Row() int {
	rv := objc.Send[int](i_.ID, objc.Sel("row"))
	return rv
}

// An index number identifying a section in a table view or collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/section
func (i_ IndexPath) Section() int {
	rv := objc.Send[int](i_.ID, objc.Sel("section"))
	return rv
}




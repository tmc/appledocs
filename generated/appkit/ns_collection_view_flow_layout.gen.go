// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CollectionViewFlowLayout] class.
var (
	CollectionViewFlowLayoutClass     _CollectionViewFlowLayoutClass
	CollectionViewFlowLayoutClassOnce sync.Once
)

func getCollectionViewFlowLayoutClass() _CollectionViewFlowLayoutClass {
	CollectionViewFlowLayoutClassOnce.Do(func() {
		CollectionViewFlowLayoutClass = _CollectionViewFlowLayoutClass{objc.GetClass("NSCollectionViewFlowLayout")}
	})
	return CollectionViewFlowLayoutClass
}

type _CollectionViewFlowLayoutClass struct {
	class objc.Class
}

// An interface definition for the [CollectionViewFlowLayout] class.
type ICollectionViewFlowLayout interface {
	ICollectionViewLayout
}

// A layout that organizes items into a flexible and configurable arrangement.
//
// In a flow layout, the first item is positioned in the top-left corner and other items are laid out either horizontally or vertically based on the scroll direction, which is configurable. Items may be the same size or different sizes, and you may use the flow layout object or the collection view’s delegate object to specify the size of items and the spacing around them. The flow layout also lets you specify custom header and footer views for each section. You can use an object as-is or subclass it to modify more aspects of the layout behavior. There are several ways to customize the basic layout behavior that do not require subclassing. For example, you can use a delegate object to change the size and spacing of items dynamically. Subclassing is appropriate for more advanced layout changes, such as adding supplementary views or decoration views, supporting custom layout attributes, or customizing the layout animations when inserting or deleting items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout
type CollectionViewFlowLayout struct {
	CollectionViewLayout
}

// CollectionViewFlowLayoutFrom constructs a [CollectionViewFlowLayout] from an unsafe.Pointer.
//
// A layout that organizes items into a flexible and configurable arrangement.
func CollectionViewFlowLayoutFrom(ptr unsafe.Pointer) CollectionViewFlowLayout {
	return CollectionViewFlowLayout{
		CollectionViewLayout: CollectionViewLayoutFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CollectionViewFlowLayoutClass) Alloc() CollectionViewFlowLayout {
	rv := objc.Send[CollectionViewFlowLayout](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CollectionViewFlowLayoutClass) New() CollectionViewFlowLayout {
	rv := objc.Send[CollectionViewFlowLayout](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionViewFlowLayout) Init() CollectionViewFlowLayout {
	rv := objc.Send[CollectionViewFlowLayout](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionViewFlowLayout) Autorelease() CollectionViewFlowLayout {
	rv := objc.Send[CollectionViewFlowLayout](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionViewFlowLayout creates a new CollectionViewFlowLayout instance.
func NewCollectionViewFlowLayout() CollectionViewFlowLayout {
	return getCollectionViewFlowLayoutClass().New()
}





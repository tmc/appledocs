// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CollectionViewItem] class.
var collectionViewItemClass = _CollectionViewItemClass{objc.GetClass("NSCollectionViewItem")}

type _CollectionViewItemClass struct {
	class objc.Class
}

// An interface definition for the [CollectionViewItem] class.
type ICollectionViewItem interface {
	IViewController
}

// The visual representation for a single data element in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem

type CollectionViewItem struct {
	ViewController
}

// CollectionViewItemFrom constructs a [CollectionViewItem] from an unsafe.Pointer.
//
// The visual representation for a single data element in a collection view.
func CollectionViewItemFrom(ptr unsafe.Pointer) CollectionViewItem {
	return CollectionViewItem{
		ViewController: ViewControllerFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (cc _CollectionViewItemClass) Alloc() CollectionViewItem {
	rv := objc.Send[CollectionViewItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (cc _CollectionViewItemClass) New() CollectionViewItem {
	rv := objc.Send[CollectionViewItem](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionViewItem) Init() CollectionViewItem {
	rv := objc.Send[CollectionViewItem](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionViewItem) Autorelease() CollectionViewItem {
	rv := objc.Send[CollectionViewItem](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionViewItem creates a new CollectionViewItem instance.
func NewCollectionViewItem() CollectionViewItem {
	return collectionViewItemClass.New()
}





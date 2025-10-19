// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CollectionViewItem] class.
var (
	collectionViewItemClass     _CollectionViewItemClass
	collectionViewItemClassOnce sync.Once
)

func getCollectionViewItemClass() _CollectionViewItemClass {
	collectionViewItemClassOnce.Do(func() {
		collectionViewItemClass = _CollectionViewItemClass{objc.GetClass("NSCollectionViewItem")}
	})
	return collectionViewItemClass
}

type _CollectionViewItemClass struct {
	class objc.Class
}

// An interface definition for the [CollectionViewItem] class.
type ICollectionViewItem interface {
	IViewController
}

// The visual representation for a single data element in a collection view.
//
// Item objects are view controllers, and you use their view hierarchies to display your content. The default implementation of this class supports the creation of a simple item that displays a single image or string. If the appearance or layout of your items is more sophisticated, you can subclass and configure the view hierarchy based on your needs. Items are the most common types of elements displayed by a collection view, and every collection view must have at least one type of item. You use items to represent the main content of your collection view interface. For example, a photo browser app would use items to display individual photos. Remember that items are only the visual interpretation of your app’s underlying data. The actual data is always managed by your app and exposed to the collection view through the data source object, which uses the data to configure the items that are displayed. The use of items with a collection view requires doing the following: Define the visual appearance of your items by specifying what views they contain and how those views are arranged. When your interface is first loaded, register your items with the collection view. (You must register your items before the collection view tries to display any content.) In your data source object, create and configure items when the collection view asks for them; see . At runtime, items merely present the data they are given. Your app’s data structures are always the original source of content, and the item contains only a copy of that data to present to the user. When the underlying data associated with an item changes, the data source should invalidate the item by calling the method of the collection view. Invalidating an item forces the collection view to dispose of it so that the collection view can create a new one with the updated content. For information about how the collection view displays items to the user, see .
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getCollectionViewItemClass().New()
}





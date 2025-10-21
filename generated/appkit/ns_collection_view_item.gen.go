// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CollectionViewItem] class.
var (
	CollectionViewItemClass     _CollectionViewItemClass
	CollectionViewItemClassOnce sync.Once
)

func getCollectionViewItemClass() _CollectionViewItemClass {
	CollectionViewItemClassOnce.Do(func() {
		CollectionViewItemClass = _CollectionViewItemClass{objc.GetClass("NSCollectionViewItem")}
	})
	return CollectionViewItemClass
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


// The collection view that owns the item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/collectionView
func (c_ CollectionViewItem) CollectionView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("collectionView"))
	return rv
}

// A text field outlet that you can use to display a string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/textField
func (c_ CollectionViewItem) TextField() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("textField"))
	return rv
}


// SetTextField sets the value of the textField property.
// A text field outlet that you can use to display a string.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/textField
func (c_ CollectionViewItem) SetTextField(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTextField:"), value)
}

// The receiver’s collection view item prototype.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/itemprototype
func (c_ CollectionViewItem) ItemPrototype() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("itemPrototype"))
	return rv
}


// SetItemPrototype sets the value of the itemPrototype property.
// The receiver’s collection view item prototype.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/itemprototype
func (c_ CollectionViewItem) SetItemPrototype(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setItemPrototype:"), value)
}

// Dragging images for multi-image drag and drop support.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/draggingimagecomponents
func (c_ CollectionViewItem) DraggingImageComponents() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("draggingImageComponents"))
	return rv
}


// SetDraggingImageComponents sets the value of the draggingImageComponents property.
// Dragging images for multi-image drag and drop support.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/draggingimagecomponents
func (c_ CollectionViewItem) SetDraggingImageComponents(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDraggingImageComponents:"), value)
}

// The highlight state currently applied to the item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/highlightstate-swift.property
func (c_ CollectionViewItem) HighlightState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("highlightState"))
	return rv
}


// SetHighlightState sets the value of the highlightState property.
// The highlight state currently applied to the item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/highlightstate-swift.property
func (c_ CollectionViewItem) SetHighlightState(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHighlightState:"), value)
}

// An image view outlet that you can use to display images.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/imageview
func (c_ CollectionViewItem) ImageView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("imageView"))
	return rv
}


// SetImageView sets the value of the imageView property.
// An image view outlet that you can use to display images.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/imageview
func (c_ CollectionViewItem) SetImageView(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImageView:"), value)
}

// A Boolean indicating whether the item is currently selected.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/isselected
func (c_ CollectionViewItem) IsSelected() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSelected"))
	return rv
}


// SetIsSelected sets the value of the isSelected property.
// A Boolean indicating whether the item is currently selected.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/isselected
func (c_ CollectionViewItem) SetIsSelected(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSelected:"), value)
}

// The view controller’s primary view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/view
func (c_ CollectionViewItem) View() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("view"))
	return rv
}


// SetView sets the value of the view property.
// The view controller’s primary view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/view
func (c_ CollectionViewItem) SetView(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setView:"), value)
}




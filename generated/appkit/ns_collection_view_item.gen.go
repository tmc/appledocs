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
	// properties:
	HighlightState() CollectionViewItemHighlightState /* not a class type */
	SetHighlightState(value CollectionViewItemHighlightState /* not a class type */)
	TextField() objc.IObject /* cross-framework: TextField */
	SetTextField(value objc.IObject /* cross-framework: TextField */)
	ItemPrototype() ICollectionViewItem
	SetItemPrototype(value ICollectionViewItem)
	CollectionView() ICollectionView
	SetCollectionView(value ICollectionView)
	DraggingImageComponents() IDraggingImageComponent
	SetDraggingImageComponents(value IDraggingImageComponent)
	ImageView() IImageView
	SetImageView(value IImageView)
	IsSelected() bool /* primitive/slice/pointer. */
	SetIsSelected(value bool /* primitive/slice/pointer. */)
	View() IView
	SetView(value IView)
	// methods:
}

// The visual representation for a single data element in a collection view.
//
// Item objects are view controllers, and you use their view hierarchies to display your content. The default implementation of this class supports the creation of a simple item that displays a single image or string. If the appearance or layout of your items is more sophisticated, you can subclass and configure the view hierarchy based on your needs. Items are the most common types of elements displayed by a collection view, and every collection view must have at least one type of item. You use items to represent the main content of your collection view interface. For example, a photo browser app would use items to display individual photos. Remember that items are only the visual interpretation of your app’s underlying data. The actual data is always managed by your app and exposed to the collection view through the data source object, which uses the data to configure the items that are displayed. The use of items with a collection view requires doing the following: Define the visual appearance of your items by specifying what views they contain and how those views are arranged. When your interface is first loaded, register your items with the collection view. (You must register your items before the collection view tries to display any content.) In your data source object, create and configure items when the collection view asks for them; see . At runtime, items merely present the data they are given. Your app’s data structures are always the original source of content, and the item contains only a copy of that data to present to the user. When the underlying data associated with an item changes, the data source should invalidate the item by calling the method of the collection view. Invalidating an item forces the collection view to dispose of it so that the collection view can create a new one with the updated content. For information about how the collection view displays items to the user, see .


// The visual representation for a single data element in a collection view.
//
// [Full Topic]
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



// The highlight state currently applied to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/highlightState-swift.property
func (c_ CollectionViewItem) HighlightState() CollectionViewItemHighlightState /* not a class type */ {
	rv := objc.Send[CollectionViewItemHighlightState](c_.ID, objc.Sel("highlightState"))
	return rv
}


// The highlight state currently applied to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/highlightState-swift.property
func (c_ CollectionViewItem) SetHighlightState(value CollectionViewItemHighlightState /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHighlightState:"), value)
}


// A text field outlet that you can use to display a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/textField
func (c_ CollectionViewItem) TextField() objc.IObject /* cross-framework: TextField */ {
	rv := objc.Send[TextField](c_.ID, objc.Sel("textField"))
	return rv
}


// A text field outlet that you can use to display a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/textField
func (c_ CollectionViewItem) SetTextField(value objc.IObject /* cross-framework: TextField */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTextField:"), value)
}


// The receiver’s collection view item prototype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/itemprototype
func (c_ CollectionViewItem) ItemPrototype() ICollectionViewItem {
	rv := objc.Send[CollectionViewItem](c_.ID, objc.Sel("itemPrototype"))
	return rv
}


// The receiver’s collection view item prototype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/itemprototype
func (c_ CollectionViewItem) SetItemPrototype(value ICollectionViewItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setItemPrototype:"), value)
}


// The collection view that owns the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/collectionview
func (c_ CollectionViewItem) CollectionView() ICollectionView {
	rv := objc.Send[CollectionView](c_.ID, objc.Sel("collectionView"))
	return rv
}


// The collection view that owns the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/collectionview
func (c_ CollectionViewItem) SetCollectionView(value ICollectionView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCollectionView:"), value)
}


// Dragging images for multi-image drag and drop support.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/draggingimagecomponents
func (c_ CollectionViewItem) DraggingImageComponents() IDraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](c_.ID, objc.Sel("draggingImageComponents"))
	return rv
}


// Dragging images for multi-image drag and drop support.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/draggingimagecomponents
func (c_ CollectionViewItem) SetDraggingImageComponents(value IDraggingImageComponent) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDraggingImageComponents:"), value)
}


// An image view outlet that you can use to display images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/imageview
func (c_ CollectionViewItem) ImageView() IImageView {
	rv := objc.Send[ImageView](c_.ID, objc.Sel("imageView"))
	return rv
}


// An image view outlet that you can use to display images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/imageview
func (c_ CollectionViewItem) SetImageView(value IImageView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImageView:"), value)
}


// A Boolean indicating whether the item is currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/isselected
func (c_ CollectionViewItem) IsSelected() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSelected"))
	return rv
}


// A Boolean indicating whether the item is currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/isselected
func (c_ CollectionViewItem) SetIsSelected(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSelected:"), value)
}


// The view controller’s primary view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/view
func (c_ CollectionViewItem) View() IView {
	rv := objc.Send[View](c_.ID, objc.Sel("view"))
	return rv
}


// The view controller’s primary view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/view
func (c_ CollectionViewItem) SetView(value IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setView:"), value)
}




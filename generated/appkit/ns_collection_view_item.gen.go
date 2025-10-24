// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSCollectionViewItem */


/* debug [class_header]: Header for NSCollectionViewItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionViewItem */
// An interface definition for the [CollectionViewItem] class.
type ICollectionViewItem interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for CollectionViewItem */
	// properties:
	CollectionView() objc.IObject /* cross-framework: CollectionView */
	DraggingImageComponents() []DraggingImageComponent
	HighlightState() CollectionViewItemHighlightState
	SetHighlightState(value CollectionViewItemHighlightState)
	ImageView() IImageView
	SetImageView(value IImageView)
	Selected() bool
	SetSelected(value bool)
	TextField() ITextField
	SetTextField(value ITextField)
	ItemPrototype() ICollectionViewItem
	SetItemPrototype(value ICollectionViewItem)
	IsSelected() bool
	SetIsSelected(value bool)
	View() IView
	SetView(value IView)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionViewItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionViewItem */
// Alloc allocates a new instance without initialization.
func (cc _CollectionViewItemClass) Alloc() CollectionViewItem {
	rv := objc.Send[CollectionViewItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionViewItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionViewItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionViewItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionViewItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionViewItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionViewItem */

// The collection view that owns the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/collectionView
func (c_ CollectionViewItem) CollectionView() objc.IObject /* cross-framework: CollectionView */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("collectionView"))
	return rv
}/* debug [instance_properties/getter]: collectionView */


// Dragging images for multi-image drag and drop support.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/draggingImageComponents
func (c_ CollectionViewItem) DraggingImageComponents() []DraggingImageComponent {
	rv := objc.Send[[]DraggingImageComponent](c_.ID, objc.Sel("draggingImageComponents"))
	return rv
}/* debug [instance_properties/getter]: draggingImageComponents */


// The highlight state currently applied to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/highlightState-swift.property
func (c_ CollectionViewItem) HighlightState() CollectionViewItemHighlightState {
	rv := objc.Send[CollectionViewItemHighlightState](c_.ID, objc.Sel("highlightState"))
	return rv
}/* debug [instance_properties/getter]: highlightState */


// The highlight state currently applied to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/highlightState-swift.property
func (c_ CollectionViewItem) SetHighlightState(value CollectionViewItemHighlightState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHighlightState:"), value)
}/* debug [instance_properties/setter]: highlightState */


// An image view outlet that you can use to display images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/imageView
func (c_ CollectionViewItem) ImageView() IImageView {
	rv := objc.Send[ImageView](c_.ID, objc.Sel("imageView"))
	return rv
}/* debug [instance_properties/getter]: imageView */


// An image view outlet that you can use to display images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/imageView
func (c_ CollectionViewItem) SetImageView(value IImageView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImageView:"), value)
}/* debug [instance_properties/setter]: imageView */


// A Boolean indicating whether the item is currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/isSelected
func (c_ CollectionViewItem) Selected() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("selected"))
	return rv
}/* debug [instance_properties/getter]: selected */


// A Boolean indicating whether the item is currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/isSelected
func (c_ CollectionViewItem) SetSelected(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSelected:"), value)
}/* debug [instance_properties/setter]: selected */


// A text field outlet that you can use to display a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/textField
func (c_ CollectionViewItem) TextField() ITextField {
	rv := objc.Send[TextField](c_.ID, objc.Sel("textField"))
	return rv
}/* debug [instance_properties/getter]: textField */


// A text field outlet that you can use to display a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/textField
func (c_ CollectionViewItem) SetTextField(value ITextField) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTextField:"), value)
}/* debug [instance_properties/setter]: textField */


// The receiver’s collection view item prototype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/itemprototype
func (c_ CollectionViewItem) ItemPrototype() ICollectionViewItem {
	rv := objc.Send[CollectionViewItem](c_.ID, objc.Sel("itemPrototype"))
	return rv
}/* debug [instance_properties/getter]: itemPrototype */


// The receiver’s collection view item prototype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/itemprototype
func (c_ CollectionViewItem) SetItemPrototype(value ICollectionViewItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setItemPrototype:"), value)
}/* debug [instance_properties/setter]: itemPrototype */


// A Boolean indicating whether the item is currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/isselected
func (c_ CollectionViewItem) IsSelected() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSelected"))
	return rv
}/* debug [instance_properties/getter]: isSelected */


// A Boolean indicating whether the item is currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/isselected
func (c_ CollectionViewItem) SetIsSelected(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSelected:"), value)
}/* debug [instance_properties/setter]: isSelected */


// The view controller’s primary view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/view
func (c_ CollectionViewItem) View() IView {
	rv := objc.Send[View](c_.ID, objc.Sel("view"))
	return rv
}/* debug [instance_properties/getter]: view */


// The view controller’s primary view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/view
func (c_ CollectionViewItem) SetView(value IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setView:"), value)
}/* debug [instance_properties/setter]: view */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionViewItem */




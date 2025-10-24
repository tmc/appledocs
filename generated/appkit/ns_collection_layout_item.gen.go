// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCollectionLayoutItem */


/* debug [class_header]: Header for NSCollectionLayoutItem */
// The class instance for the [CollectionLayoutItem] class.
var (
	CollectionLayoutItemClass     _CollectionLayoutItemClass
	CollectionLayoutItemClassOnce sync.Once
)

func getCollectionLayoutItemClass() _CollectionLayoutItemClass {
	CollectionLayoutItemClassOnce.Do(func() {
		CollectionLayoutItemClass = _CollectionLayoutItemClass{objc.GetClass("NSCollectionLayoutItem")}
	})
	return CollectionLayoutItemClass
}

type _CollectionLayoutItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionLayoutItem */
// An interface definition for the [CollectionLayoutItem] class.
type ICollectionLayoutItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CollectionLayoutItem */
	// properties:
	ContentInsets() objc.IObject /* cross-framework: DirectionalEdgeInsets */
	SetContentInsets(value objc.IObject /* cross-framework: DirectionalEdgeInsets */)
	EdgeSpacing() ICollectionLayoutEdgeSpacing
	SetEdgeSpacing(value ICollectionLayoutEdgeSpacing)
	LayoutSize() ICollectionLayoutSize
	SupplementaryItems() []CollectionLayoutSupplementaryItem
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionLayoutItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionLayoutItem */
// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutItemClass) Alloc() CollectionLayoutItem {
	rv := objc.Send[CollectionLayoutItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionLayoutItemClass) New() CollectionLayoutItem {
	rv := objc.Send[CollectionLayoutItem](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionLayoutItem) Init() CollectionLayoutItem {
	rv := objc.Send[CollectionLayoutItem](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionLayoutItem) Autorelease() CollectionLayoutItem {
	rv := objc.Send[CollectionLayoutItem](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionLayoutItem creates a new CollectionLayoutItem instance.
func NewCollectionLayoutItem() CollectionLayoutItem {
	return getCollectionLayoutItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionLayoutItem */
// The most basic component of a collection view’s layout.
//
// An item is a blueprint for how to size, space, and arrange an individual piece of content in your collection view. An item represents a single view that’s rendered onscreen. Generally, an item is a cell, but items can be supplementary views like headers, footers, and other decorations. For example, in the Photos app, an item might represent a single photo. In the App Store app, an item might be a cell displaying information about an individual app in a list of featured apps, such as the app icon, app name, tagline, and download button. Each item specifies its own size in terms of a width dimension and a height dimension. Items can express their dimensions relative to their container, as an absolute value, or as an estimated value that might change at runtime, like in response to a change in system font size. For more information, see . You combine items into groups that determine how those items are arranged in relation to each other. For more information, see .


// The most basic component of a collection view’s layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem
type CollectionLayoutItem struct {
	objectivec.Object
}

// CollectionLayoutItemFrom constructs a [CollectionLayoutItem] from an unsafe.Pointer.
//
// The most basic component of a collection view’s layout.
func CollectionLayoutItemFrom(ptr unsafe.Pointer) CollectionLayoutItem {
	return CollectionLayoutItem{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionLayoutItem */

// Creates an item of the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/init(layoutSize:)
func NewCollectionLayoutItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutItem {
	rv := objc.Send[CollectionLayoutItem](objc.ID(getCollectionLayoutItemClass().class), objc.Sel("itemWithLayoutSize:"), layoutSize)
	return rv
}/* debug [class_init_methods/constructor]: NewCollectionLayoutItemWithLayoutSize */


// Creates an item of the specified size with an array of supplementary items to attach to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/init(layoutSize:supplementaryItems:)
func NewCollectionLayoutItemWithLayoutSizeSupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []CollectionLayoutSupplementaryItem) CollectionLayoutItem {
	rv := objc.Send[CollectionLayoutItem](objc.ID(getCollectionLayoutItemClass().class), objc.Sel("itemWithLayoutSize:supplementaryItems:"), layoutSize, supplementaryItems)
	return rv
}/* debug [class_init_methods/constructor]: NewCollectionLayoutItemWithLayoutSizeSupplementaryItems */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionLayoutItem */

// Creates an item of the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/init(layoutSize:)
func (cc _CollectionLayoutItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("itemWithLayoutSize:"), layoutSize)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ItemWithLayoutSize) */


// Creates an item of the specified size with an array of supplementary items to attach to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/init(layoutSize:supplementaryItems:)
func (cc _CollectionLayoutItemClass) ItemWithLayoutSizeSupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []CollectionLayoutSupplementaryItem) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("itemWithLayoutSize:supplementaryItems:"), layoutSize, supplementaryItems)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ItemWithLayoutSizeSupplementaryItems) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionLayoutItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionLayoutItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionLayoutItem */

// The amount of space added around the content of the item to adjust its final size after its position is computed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/contentInsets
func (c_ CollectionLayoutItem) ContentInsets() objc.IObject /* cross-framework: DirectionalEdgeInsets */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("contentInsets"))
	return rv
}/* debug [instance_properties/getter]: contentInsets */


// The amount of space added around the content of the item to adjust its final size after its position is computed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/contentInsets
func (c_ CollectionLayoutItem) SetContentInsets(value objc.IObject /* cross-framework: DirectionalEdgeInsets */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentInsets:"), value)
}/* debug [instance_properties/setter]: contentInsets */


// The amount of space added around the boundaries of the item between other items and this item’s container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/edgeSpacing
func (c_ CollectionLayoutItem) EdgeSpacing() ICollectionLayoutEdgeSpacing {
	rv := objc.Send[CollectionLayoutEdgeSpacing](c_.ID, objc.Sel("edgeSpacing"))
	return rv
}/* debug [instance_properties/getter]: edgeSpacing */


// The amount of space added around the boundaries of the item between other items and this item’s container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/edgeSpacing
func (c_ CollectionLayoutItem) SetEdgeSpacing(value ICollectionLayoutEdgeSpacing) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEdgeSpacing:"), value)
}/* debug [instance_properties/setter]: edgeSpacing */


// The item’s size expressed in width and height dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/layoutSize
func (c_ CollectionLayoutItem) LayoutSize() ICollectionLayoutSize {
	rv := objc.Send[CollectionLayoutSize](c_.ID, objc.Sel("layoutSize"))
	return rv
}/* debug [instance_properties/getter]: layoutSize */


// An array of the supplementary items attached to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/supplementaryItems
func (c_ CollectionLayoutItem) SupplementaryItems() []CollectionLayoutSupplementaryItem {
	rv := objc.Send[[]CollectionLayoutSupplementaryItem](c_.ID, objc.Sel("supplementaryItems"))
	return rv
}/* debug [instance_properties/getter]: supplementaryItems */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionLayoutItem */



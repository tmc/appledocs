// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCollectionLayoutGroup */


/* debug [class_header]: Header for NSCollectionLayoutGroup */
// The class instance for the [CollectionLayoutGroup] class.
var (
	CollectionLayoutGroupClass     _CollectionLayoutGroupClass
	CollectionLayoutGroupClassOnce sync.Once
)

func getCollectionLayoutGroupClass() _CollectionLayoutGroupClass {
	CollectionLayoutGroupClassOnce.Do(func() {
		CollectionLayoutGroupClass = _CollectionLayoutGroupClass{objc.GetClass("NSCollectionLayoutGroup")}
	})
	return CollectionLayoutGroupClass
}

type _CollectionLayoutGroupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionLayoutGroup */
// An interface definition for the [CollectionLayoutGroup] class.
type ICollectionLayoutGroup interface {
	ICollectionLayoutItem
	
/* debug [class_interface_properties]: Properties for CollectionLayoutGroup */
	// properties:
	InterItemSpacing() ICollectionLayoutSpacing
	SetInterItemSpacing(value ICollectionLayoutSpacing)
	Subitems() []CollectionLayoutItem
	SupplementaryItems() []CollectionLayoutSupplementaryItem
	SetSupplementaryItems(value []CollectionLayoutSupplementaryItem)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionLayoutGroup */
	// methods:
	VisualDescription() foundation.String
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionLayoutGroup */
// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutGroupClass) Alloc() CollectionLayoutGroup {
	rv := objc.Send[CollectionLayoutGroup](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionLayoutGroupClass) New() CollectionLayoutGroup {
	rv := objc.Send[CollectionLayoutGroup](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionLayoutGroup) Init() CollectionLayoutGroup {
	rv := objc.Send[CollectionLayoutGroup](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionLayoutGroup) Autorelease() CollectionLayoutGroup {
	rv := objc.Send[CollectionLayoutGroup](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionLayoutGroup creates a new CollectionLayoutGroup instance.
func NewCollectionLayoutGroup() CollectionLayoutGroup {
	return getCollectionLayoutGroupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionLayoutGroup */
// A container for a set of items that lays out the items along a path.
//
// Groups determine how the items in a collection view lay out in relation to each other. A group might lay out its items in a horizontal row, a vertical column, or a custom arrangement. A group determines the rules for how items are rendered in relation to each other, but in itself doesn’t render any content. For example, in the Photos app, a group of items is a row of photos. In the App Store app, a group might be a single column of cells (items) arranged in a vertical column. Each group specifies its own size in terms of a width dimension and a height dimension. Groups can express their dimensions relative to their container, as an absolute value, or as an estimated value that might change at runtime, like in response to a change in system font size. For more information, see . Because a group is a subclass of , it behaves like an item. You can combine a group with other items and groups into more complex layouts. After you configure a group, you must initialize a section ( ) of your collection view layout with that group.


// A container for a set of items that lays out the items along a path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup
type CollectionLayoutGroup struct {
	CollectionLayoutItem
}

// CollectionLayoutGroupFrom constructs a [CollectionLayoutGroup] from an unsafe.Pointer.
//
// A container for a set of items that lays out the items along a path.
func CollectionLayoutGroupFrom(ptr unsafe.Pointer) CollectionLayoutGroup {
	return CollectionLayoutGroup{
		CollectionLayoutItem: CollectionLayoutItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionLayoutGroup *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionLayoutGroup */

// Creates a group of the specified size, with an item provider that creates a custom arrangement for those items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/custom(layoutSize:itemProvider:)
func (cc _CollectionLayoutGroupClass) CustomGroupWithLayoutSizeItemProvider(layoutSize ICollectionLayoutSize, itemProvider CollectionLayoutGroupCustomItemProvider /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("customGroupWithLayoutSize:itemProvider:"), layoutSize, itemProvider)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CustomGroupWithLayoutSizeItemProvider) */


// Creates a group of the specified size, containing an array of equally sized items arranged in a horizontal line up to the number specified by count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/horizontal(layoutSize:subitem:count:)
func (cc _CollectionLayoutGroupClass) HorizontalGroupWithLayoutSizeSubitemCount(layoutSize ICollectionLayoutSize, subitem ICollectionLayoutItem, count int) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("horizontalGroupWithLayoutSize:subitem:count:"), layoutSize, subitem, count)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HorizontalGroupWithLayoutSizeSubitemCount) */


// Creates a group of the specified size, containing an array of items arranged in a horizontal line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/horizontal(layoutSize:subitems:)
func (cc _CollectionLayoutGroupClass) HorizontalGroupWithLayoutSizeSubitems(layoutSize ICollectionLayoutSize, subitems []CollectionLayoutItem) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("horizontalGroupWithLayoutSize:subitems:"), layoutSize, subitems)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HorizontalGroupWithLayoutSizeSubitems) */


// Creates a group of the specified size, containing an array of equally sized items arranged in a vertical line up to the number specified by count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/vertical(layoutSize:subitem:count:)
func (cc _CollectionLayoutGroupClass) VerticalGroupWithLayoutSizeSubitemCount(layoutSize ICollectionLayoutSize, subitem ICollectionLayoutItem, count int) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("verticalGroupWithLayoutSize:subitem:count:"), layoutSize, subitem, count)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VerticalGroupWithLayoutSizeSubitemCount) */


// Creates a group of the specified size, containing an array of items arranged in a vertical line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/vertical(layoutSize:subitems:)
func (cc _CollectionLayoutGroupClass) VerticalGroupWithLayoutSizeSubitems(layoutSize ICollectionLayoutSize, subitems []CollectionLayoutItem) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("verticalGroupWithLayoutSize:subitems:"), layoutSize, subitems)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VerticalGroupWithLayoutSizeSubitems) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionLayoutGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionLayoutGroup */

// Returns a string with an ASCII representation of the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/visualDescription()
func (c_ CollectionLayoutGroup) VisualDescription() foundation.String {
	rv := objc.Send[foundation.String](c_.ID, objc.Sel("visualDescription"))
	return rv
}/* debug [instance_methods/method]: VisualDescription */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionLayoutGroup */

// The amount of space between the items in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/interItemSpacing
func (c_ CollectionLayoutGroup) InterItemSpacing() ICollectionLayoutSpacing {
	rv := objc.Send[CollectionLayoutSpacing](c_.ID, objc.Sel("interItemSpacing"))
	return rv
}/* debug [instance_properties/getter]: interItemSpacing */


// The amount of space between the items in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/interItemSpacing
func (c_ CollectionLayoutGroup) SetInterItemSpacing(value ICollectionLayoutSpacing) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInterItemSpacing:"), value)
}/* debug [instance_properties/setter]: interItemSpacing */


// An array of the items contained in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/subitems
func (c_ CollectionLayoutGroup) Subitems() []CollectionLayoutItem {
	rv := objc.Send[[]CollectionLayoutItem](c_.ID, objc.Sel("subitems"))
	return rv
}/* debug [instance_properties/getter]: subitems */


// An array of the supplementary items that are anchored to the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/supplementaryItems
func (c_ CollectionLayoutGroup) SupplementaryItems() []CollectionLayoutSupplementaryItem {
	rv := objc.Send[[]CollectionLayoutSupplementaryItem](c_.ID, objc.Sel("supplementaryItems"))
	return rv
}/* debug [instance_properties/getter]: supplementaryItems */


// An array of the supplementary items that are anchored to the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/supplementaryItems
func (c_ CollectionLayoutGroup) SetSupplementaryItems(value []CollectionLayoutSupplementaryItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupplementaryItems:"), nsArray)
}/* debug [instance_properties/setter]: supplementaryItems */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionLayoutGroup */




// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCollectionLayoutDecorationItem */


/* debug [class_header]: Header for NSCollectionLayoutDecorationItem */
// The class instance for the [CollectionLayoutDecorationItem] class.
var (
	CollectionLayoutDecorationItemClass     _CollectionLayoutDecorationItemClass
	CollectionLayoutDecorationItemClassOnce sync.Once
)

func getCollectionLayoutDecorationItemClass() _CollectionLayoutDecorationItemClass {
	CollectionLayoutDecorationItemClassOnce.Do(func() {
		CollectionLayoutDecorationItemClass = _CollectionLayoutDecorationItemClass{objc.GetClass("NSCollectionLayoutDecorationItem")}
	})
	return CollectionLayoutDecorationItemClass
}

type _CollectionLayoutDecorationItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionLayoutDecorationItem */
// An interface definition for the [CollectionLayoutDecorationItem] class.
type ICollectionLayoutDecorationItem interface {
	ICollectionLayoutItem
	
/* debug [class_interface_properties]: Properties for CollectionLayoutDecorationItem */
	// properties:
	ElementKind() objc.IObject /* cross-framework: NSString */
	ZIndex() int
	SetZIndex(value int)
	DecorationItems() ICollectionLayoutDecorationItem
	SetDecorationItems(value ICollectionLayoutDecorationItem)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionLayoutDecorationItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionLayoutDecorationItem */
// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutDecorationItemClass) Alloc() CollectionLayoutDecorationItem {
	rv := objc.Send[CollectionLayoutDecorationItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionLayoutDecorationItemClass) New() CollectionLayoutDecorationItem {
	rv := objc.Send[CollectionLayoutDecorationItem](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionLayoutDecorationItem) Init() CollectionLayoutDecorationItem {
	rv := objc.Send[CollectionLayoutDecorationItem](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionLayoutDecorationItem) Autorelease() CollectionLayoutDecorationItem {
	rv := objc.Send[CollectionLayoutDecorationItem](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionLayoutDecorationItem creates a new CollectionLayoutDecorationItem instance.
func NewCollectionLayoutDecorationItem() CollectionLayoutDecorationItem {
	return getCollectionLayoutDecorationItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionLayoutDecorationItem */
// An object used to add a background to a section of a collection view.
//
// Each type of decoration item must have a unique element kind. Consider tracking these strings together in a way that makes it straightforward to identify each element, for example: Add a background to a section by setting that section’s property:


// An object used to add a background to a section of a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDecorationItem
type CollectionLayoutDecorationItem struct {
	CollectionLayoutItem
}

// CollectionLayoutDecorationItemFrom constructs a [CollectionLayoutDecorationItem] from an unsafe.Pointer.
//
// An object used to add a background to a section of a collection view.
func CollectionLayoutDecorationItemFrom(ptr unsafe.Pointer) CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItem{
		CollectionLayoutItem: CollectionLayoutItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionLayoutDecorationItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionLayoutDecorationItem */

// Creates a section background with a string to identify the element kind.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDecorationItem/background(elementKind:)
func (cc _CollectionLayoutDecorationItemClass) BackgroundDecorationItemWithElementKind(elementKind objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("backgroundDecorationItemWithElementKind:"), elementKind)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BackgroundDecorationItemWithElementKind) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionLayoutDecorationItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionLayoutDecorationItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionLayoutDecorationItem */

// A string that identifies the type of decoration item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDecorationItem/elementKind
func (c_ CollectionLayoutDecorationItem) ElementKind() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("elementKind"))
	return rv
}/* debug [instance_properties/getter]: elementKind */


// The vertical stacking order of the decoration item in relation to other items in the section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDecorationItem/zIndex
func (c_ CollectionLayoutDecorationItem) ZIndex() int {
	rv := objc.Send[int](c_.ID, objc.Sel("zIndex"))
	return rv
}/* debug [instance_properties/getter]: zIndex */


// The vertical stacking order of the decoration item in relation to other items in the section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDecorationItem/zIndex
func (c_ CollectionLayoutDecorationItem) SetZIndex(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZIndex:"), value)
}/* debug [instance_properties/setter]: zIndex */


// An array of the decoration items that are anchored to the section, such as background decoration views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutsection/decorationitems
func (c_ CollectionLayoutDecorationItem) DecorationItems() ICollectionLayoutDecorationItem {
	rv := objc.Send[CollectionLayoutDecorationItem](c_.ID, objc.Sel("decorationItems"))
	return rv
}/* debug [instance_properties/getter]: decorationItems */


// An array of the decoration items that are anchored to the section, such as background decoration views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutsection/decorationitems
func (c_ CollectionLayoutDecorationItem) SetDecorationItems(value ICollectionLayoutDecorationItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDecorationItems:"), value)
}/* debug [instance_properties/setter]: decorationItems */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionLayoutDecorationItem */




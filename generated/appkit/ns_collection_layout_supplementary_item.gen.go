// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CollectionLayoutSupplementaryItem] class.
var (
	CollectionLayoutSupplementaryItemClass     _CollectionLayoutSupplementaryItemClass
	CollectionLayoutSupplementaryItemClassOnce sync.Once
)

func getCollectionLayoutSupplementaryItemClass() _CollectionLayoutSupplementaryItemClass {
	CollectionLayoutSupplementaryItemClassOnce.Do(func() {
		CollectionLayoutSupplementaryItemClass = _CollectionLayoutSupplementaryItemClass{objc.GetClass("NSCollectionLayoutSupplementaryItem")}
	})
	return CollectionLayoutSupplementaryItemClass
}

type _CollectionLayoutSupplementaryItemClass struct {
	class objc.Class
}

// An interface definition for the [CollectionLayoutSupplementaryItem] class.
type ICollectionLayoutSupplementaryItem interface {
	ICollectionLayoutItem
	// properties:
	ContainerAnchor() ICollectionLayoutAnchor
	ElementKind() objc.IObject /* cross-framework: NSString */
	ItemAnchor() ICollectionLayoutAnchor
	ZIndex() int
	SetZIndex(value int)
	// methods:
}

// An object used to add an extra visual decoration to an item in a collection view.
//
// You use supplementary items to attach additional views to your content. For example, you might attach a badge to an item or a frame around a group. A supplementary item follows the index path of the item it’s attached to. If you want to create a header or footer for your layout or its sections, use a boundary supplementary item (<````NSCollectionLayoutBoundarySupplementaryItem``>) instead. Each type of supplementary item must have a unique element kind. Consider tracking these strings together in a way that makes it straightforward to identify each element, for example: Add supplementary items to an item by passing in an array of supplementary items when you construct the item:


// An object used to add an extra visual decoration to an item in a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem
type CollectionLayoutSupplementaryItem struct {
	CollectionLayoutItem
}

// CollectionLayoutSupplementaryItemFrom constructs a [CollectionLayoutSupplementaryItem] from an unsafe.Pointer.
//
// An object used to add an extra visual decoration to an item in a collection view.
func CollectionLayoutSupplementaryItemFrom(ptr unsafe.Pointer) CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItem{
		CollectionLayoutItem: CollectionLayoutItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutSupplementaryItemClass) Alloc() CollectionLayoutSupplementaryItem {
	rv := objc.Send[CollectionLayoutSupplementaryItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CollectionLayoutSupplementaryItemClass) New() CollectionLayoutSupplementaryItem {
	rv := objc.Send[CollectionLayoutSupplementaryItem](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionLayoutSupplementaryItem) Init() CollectionLayoutSupplementaryItem {
	rv := objc.Send[CollectionLayoutSupplementaryItem](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionLayoutSupplementaryItem) Autorelease() CollectionLayoutSupplementaryItem {
	rv := objc.Send[CollectionLayoutSupplementaryItem](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionLayoutSupplementaryItem creates a new CollectionLayoutSupplementaryItem instance.
func NewCollectionLayoutSupplementaryItem() CollectionLayoutSupplementaryItem {
	return getCollectionLayoutSupplementaryItemClass().New()
}



// Creates a supplementary item of the specified size and element kind, with an anchor relative to a container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/init(layoutSize:elementKind:containerAnchor:)
func NewCollectionLayoutSupplementaryItemWithLayoutSizeElementKindContainerAnchor(layoutSize ICollectionLayoutSize, elementKind objc.IObject /* cross-framework: NSString */, containerAnchor ICollectionLayoutAnchor) CollectionLayoutSupplementaryItem {
	rv := objc.Send[CollectionLayoutSupplementaryItem](objc.ID(getCollectionLayoutSupplementaryItemClass().class), objc.Sel("supplementaryItemWithLayoutSize:elementKind:containerAnchor:"), layoutSize, elementKind, containerAnchor)
	return rv
}


// Creates a supplementary item of the specified size and element kind, an anchor relative to a container, and an anchor relative to an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/init(layoutSize:elementKind:containerAnchor:itemAnchor:)
func NewCollectionLayoutSupplementaryItemWithLayoutSizeElementKindContainerAnchorItemAnchor(layoutSize ICollectionLayoutSize, elementKind objc.IObject /* cross-framework: NSString */, containerAnchor ICollectionLayoutAnchor, itemAnchor ICollectionLayoutAnchor) CollectionLayoutSupplementaryItem {
	rv := objc.Send[CollectionLayoutSupplementaryItem](objc.ID(getCollectionLayoutSupplementaryItemClass().class), objc.Sel("supplementaryItemWithLayoutSize:elementKind:containerAnchor:itemAnchor:"), layoutSize, elementKind, containerAnchor, itemAnchor)
	return rv
}



// Creates a supplementary item of the specified size and element kind, with an anchor relative to a container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/init(layoutSize:elementKind:containerAnchor:)
func (cc _CollectionLayoutSupplementaryItemClass) SupplementaryItemWithLayoutSizeElementKindContainerAnchor(layoutSize ICollectionLayoutSize, elementKind objc.IObject /* cross-framework: NSString */, containerAnchor ICollectionLayoutAnchor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("supplementaryItemWithLayoutSize:elementKind:containerAnchor:"), layoutSize, elementKind, containerAnchor)
	return rv
}


// Creates a supplementary item of the specified size and element kind, an anchor relative to a container, and an anchor relative to an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/init(layoutSize:elementKind:containerAnchor:itemAnchor:)
func (cc _CollectionLayoutSupplementaryItemClass) SupplementaryItemWithLayoutSizeElementKindContainerAnchorItemAnchor(layoutSize ICollectionLayoutSize, elementKind objc.IObject /* cross-framework: NSString */, containerAnchor ICollectionLayoutAnchor, itemAnchor ICollectionLayoutAnchor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("supplementaryItemWithLayoutSize:elementKind:containerAnchor:itemAnchor:"), layoutSize, elementKind, containerAnchor, itemAnchor)
	return rv
}


// The anchor between the supplementary item and the container it’s attached to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/containerAnchor
func (c_ CollectionLayoutSupplementaryItem) ContainerAnchor() ICollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](c_.ID, objc.Sel("containerAnchor"))
	return rv
}


// A string that identifies the type of supplementary item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/elementKind
func (c_ CollectionLayoutSupplementaryItem) ElementKind() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("elementKind"))
	return rv
}


// The anchor between the supplementary item and the item it’s attached to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/itemAnchor
func (c_ CollectionLayoutSupplementaryItem) ItemAnchor() ICollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](c_.ID, objc.Sel("itemAnchor"))
	return rv
}


// The vertical stacking order of the supplementary item in relation to other items in the section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/zIndex
func (c_ CollectionLayoutSupplementaryItem) ZIndex() int {
	rv := objc.Send[int](c_.ID, objc.Sel("zIndex"))
	return rv
}


// The vertical stacking order of the supplementary item in relation to other items in the section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/zIndex
func (c_ CollectionLayoutSupplementaryItem) SetZIndex(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZIndex:"), value)
}



// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CollectionLayoutBoundarySupplementaryItem] class.
var (
	CollectionLayoutBoundarySupplementaryItemClass     _CollectionLayoutBoundarySupplementaryItemClass
	CollectionLayoutBoundarySupplementaryItemClassOnce sync.Once
)

func getCollectionLayoutBoundarySupplementaryItemClass() _CollectionLayoutBoundarySupplementaryItemClass {
	CollectionLayoutBoundarySupplementaryItemClassOnce.Do(func() {
		CollectionLayoutBoundarySupplementaryItemClass = _CollectionLayoutBoundarySupplementaryItemClass{objc.GetClass("NSCollectionLayoutBoundarySupplementaryItem")}
	})
	return CollectionLayoutBoundarySupplementaryItemClass
}

type _CollectionLayoutBoundarySupplementaryItemClass struct {
	class objc.Class
}





// An interface definition for the [CollectionLayoutBoundarySupplementaryItem] class.
type ICollectionLayoutBoundarySupplementaryItem interface {
	ICollectionLayoutSupplementaryItem
	

	// properties:
	Alignment() RectAlignment
	ExtendsBoundary() bool
	SetExtendsBoundary(value bool)
	Offset() corefoundation.CGPoint
	PinToVisibleBounds() bool
	SetPinToVisibleBounds(value bool)
	BoundarySupplementaryItems() ICollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value ICollectionLayoutBoundarySupplementaryItem)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutBoundarySupplementaryItemClass) Alloc() CollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[CollectionLayoutBoundarySupplementaryItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionLayoutBoundarySupplementaryItemClass) New() CollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[CollectionLayoutBoundarySupplementaryItem](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionLayoutBoundarySupplementaryItem) Init() CollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[CollectionLayoutBoundarySupplementaryItem](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionLayoutBoundarySupplementaryItem) Autorelease() CollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[CollectionLayoutBoundarySupplementaryItem](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionLayoutBoundarySupplementaryItem creates a new CollectionLayoutBoundarySupplementaryItem instance.
func NewCollectionLayoutBoundarySupplementaryItem() CollectionLayoutBoundarySupplementaryItem {
	return getCollectionLayoutBoundarySupplementaryItemClass().New()
}





// An object used to add headers or footers to a collection view.
//
// A boundary supplementary item is a specialized type of supplementary item ( ). You use boundary supplementary items to add headers or footers to a section of a collection view or the entire collection view. Each type of supplementary item must have a unique element kind. Consider tracking these strings together in a way that makes it straightforward to identify each element, for example: Add boundary supplementary items to a section by setting that section’s property:


// An object used to add headers or footers to a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem
type CollectionLayoutBoundarySupplementaryItem struct {
	CollectionLayoutSupplementaryItem
}

// CollectionLayoutBoundarySupplementaryItemFrom constructs a [CollectionLayoutBoundarySupplementaryItem] from an unsafe.Pointer.
//
// An object used to add headers or footers to a collection view.
func CollectionLayoutBoundarySupplementaryItemFrom(ptr unsafe.Pointer) CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItem{
		CollectionLayoutSupplementaryItem: CollectionLayoutSupplementaryItemFrom(ptr),
	}
}






// Creates a boundary supplementary item of the specified size and element kind, with an alignment relative to a section or layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem/init(layoutSize:elementKind:alignment:)
func NewCollectionLayoutBoundarySupplementaryItemWithLayoutSizeElementKindAlignment(layoutSize ICollectionLayoutSize, elementKind foundation.foundation.INSString, alignment RectAlignment) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[CollectionLayoutBoundarySupplementaryItem](objc.ID(getCollectionLayoutBoundarySupplementaryItemClass().class), objc.Sel("boundarySupplementaryItemWithLayoutSize:elementKind:alignment:"), layoutSize, elementKind, alignment)
	return rv
}


// Creates a boundary supplementary item of the specified size and element kind, with an alignment relative to a section or layout at an absolute offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem/init(layoutSize:elementKind:alignment:absoluteOffset:)
func NewCollectionLayoutBoundarySupplementaryItemWithLayoutSizeElementKindAlignmentAbsoluteOffset(layoutSize ICollectionLayoutSize, elementKind foundation.foundation.INSString, alignment RectAlignment, absoluteOffset corefoundation.CGPoint) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[CollectionLayoutBoundarySupplementaryItem](objc.ID(getCollectionLayoutBoundarySupplementaryItemClass().class), objc.Sel("boundarySupplementaryItemWithLayoutSize:elementKind:alignment:absoluteOffset:"), layoutSize, elementKind, alignment, absoluteOffset)
	return rv
}







// Creates a boundary supplementary item of the specified size and element kind, with an alignment relative to a section or layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem/init(layoutSize:elementKind:alignment:)
func (cc _CollectionLayoutBoundarySupplementaryItemClass) BoundarySupplementaryItemWithLayoutSizeElementKindAlignment(layoutSize ICollectionLayoutSize, elementKind foundation.foundation.INSString, alignment RectAlignment) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("boundarySupplementaryItemWithLayoutSize:elementKind:alignment:"), layoutSize, elementKind, alignment)
	return rv
}


// Creates a boundary supplementary item of the specified size and element kind, with an alignment relative to a section or layout at an absolute offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem/init(layoutSize:elementKind:alignment:absoluteOffset:)
func (cc _CollectionLayoutBoundarySupplementaryItemClass) BoundarySupplementaryItemWithLayoutSizeElementKindAlignmentAbsoluteOffset(layoutSize ICollectionLayoutSize, elementKind foundation.foundation.INSString, alignment RectAlignment, absoluteOffset corefoundation.CGPoint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("boundarySupplementaryItemWithLayoutSize:elementKind:alignment:absoluteOffset:"), layoutSize, elementKind, alignment, absoluteOffset)
	return rv
}

















// The alignment of the boundary supplementary item relative to the section or layout it’s attached to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem/alignment
func (c_ CollectionLayoutBoundarySupplementaryItem) Alignment() RectAlignment {
	rv := objc.Send[RectAlignment](c_.ID, objc.Sel("alignment"))
	return rv
}


// A Boolean value that indicates whether a boundary supplementary item extends the content area of the section or layout it’s attached to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem/extendsBoundary
func (c_ CollectionLayoutBoundarySupplementaryItem) ExtendsBoundary() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("extendsBoundary"))
	return rv
}


// A Boolean value that indicates whether a boundary supplementary item extends the content area of the section or layout it’s attached to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem/extendsBoundary
func (c_ CollectionLayoutBoundarySupplementaryItem) SetExtendsBoundary(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExtendsBoundary:"), value)
}


// The floating-point value of the boundary supplementary item’s offset from the section or layout it’s attached to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem/offset
func (c_ CollectionLayoutBoundarySupplementaryItem) Offset() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c_.ID, objc.Sel("offset"))
	return rv
}


// A Boolean value that indicates whether a header or footer is pinned to the top or bottom visible boundary of the section or layout it’s attached to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem/pinToVisibleBounds
func (c_ CollectionLayoutBoundarySupplementaryItem) PinToVisibleBounds() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("pinToVisibleBounds"))
	return rv
}


// A Boolean value that indicates whether a header or footer is pinned to the top or bottom visible boundary of the section or layout it’s attached to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem/pinToVisibleBounds
func (c_ CollectionLayoutBoundarySupplementaryItem) SetPinToVisibleBounds(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPinToVisibleBounds:"), value)
}


// An array of the supplementary items that are associated with the boundary edges of the entire layout, such as global headers and footers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayoutconfiguration/boundarysupplementaryitems
func (c_ CollectionLayoutBoundarySupplementaryItem) BoundarySupplementaryItems() ICollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[CollectionLayoutBoundarySupplementaryItem](c_.ID, objc.Sel("boundarySupplementaryItems"))
	return rv
}


// An array of the supplementary items that are associated with the boundary edges of the entire layout, such as global headers and footers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayoutconfiguration/boundarysupplementaryitems
func (c_ CollectionLayoutBoundarySupplementaryItem) SetBoundarySupplementaryItems(value ICollectionLayoutBoundarySupplementaryItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBoundarySupplementaryItems:"), value)
}








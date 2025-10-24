// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CollectionViewLayoutAttributes] class.
var (
	CollectionViewLayoutAttributesClass     _CollectionViewLayoutAttributesClass
	CollectionViewLayoutAttributesClassOnce sync.Once
)

func getCollectionViewLayoutAttributesClass() _CollectionViewLayoutAttributesClass {
	CollectionViewLayoutAttributesClassOnce.Do(func() {
		CollectionViewLayoutAttributesClass = _CollectionViewLayoutAttributesClass{objc.GetClass("NSCollectionViewLayoutAttributes")}
	})
	return CollectionViewLayoutAttributesClass
}

type _CollectionViewLayoutAttributesClass struct {
	class objc.Class
}

// An interface definition for the [CollectionViewLayoutAttributes] class.
type ICollectionViewLayoutAttributes interface {
	objectivec.IObject
	// properties:
	Alpha() float64
	SetAlpha(value float64)
	Frame() objc.IObject /* cross-framework: Rect */
	SetFrame(value objc.IObject /* cross-framework: Rect */)
	IndexPath() foundation.IndexPath
	SetIndexPath(value foundation.IndexPath)
	Hidden() bool
	SetHidden(value bool)
	RepresentedElementCategory() CollectionElementCategory
	RepresentedElementKind() objc.IObject /* cross-framework: NSString */
	Size() objc.IObject /* cross-framework: Size */
	SetSize(value objc.IObject /* cross-framework: Size */)
	ZIndex() int
	SetZIndex(value int)
	IsHidden() bool
	SetIsHidden(value bool)
	// methods:
}

// An object that contains layout-related attributes for an element in a collection view.
//
// During the layout, the layout object creates instances of for each element displayed in the collection view. The layout attributes describe the position of an element and other information such as its alpha and position on the z axis. The collection view later applies the layout attributes to the onscreen elements. The only time you interact with layout attribute objects is when you implement a custom layout, and the interactions are straightforward. When asked for layout attributes for a specific element, your layout object uses the methods of this class to create an appropriate instance of the class based on the type of the requested element. It then configures the properties of the object and returns it to the requester.


// An object that contains layout-related attributes for an element in a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes
type CollectionViewLayoutAttributes struct {
	objectivec.Object
}

// CollectionViewLayoutAttributesFrom constructs a [CollectionViewLayoutAttributes] from an unsafe.Pointer.
//
// An object that contains layout-related attributes for an element in a collection view.
func CollectionViewLayoutAttributesFrom(ptr unsafe.Pointer) CollectionViewLayoutAttributes {
	return CollectionViewLayoutAttributes{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CollectionViewLayoutAttributesClass) Alloc() CollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CollectionViewLayoutAttributesClass) New() CollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionViewLayoutAttributes) Init() CollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionViewLayoutAttributes) Autorelease() CollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionViewLayoutAttributes creates a new CollectionViewLayoutAttributes instance.
func NewCollectionViewLayoutAttributes() CollectionViewLayoutAttributes {
	return getCollectionViewLayoutAttributesClass().New()
}



// Creates and returns a layout attributes object for a decoration view based on the specified information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forDecorationViewOfKind:with:)
func NewCollectionViewLayoutAttributesForDecorationViewOfKindWithIndexPath(decorationViewKind objc.IObject /* cross-framework: CollectionViewDecorationElementKind */, indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](objc.ID(getCollectionViewLayoutAttributesClass().class), objc.Sel("layoutAttributesForDecorationViewOfKind:withIndexPath:"), decorationViewKind, indexPath)
	return rv
}


// Creates and returns a layout attributes object for an inter-item gap view at the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forInterItemGapBefore:)
func NewCollectionViewLayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](objc.ID(getCollectionViewLayoutAttributesClass().class), objc.Sel("layoutAttributesForInterItemGapBeforeIndexPath:"), indexPath)
	return rv
}


// Creates and returns a layout attributes object for the item at the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forItemWith:)
func NewCollectionViewLayoutAttributesForItemWithIndexPath(indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](objc.ID(getCollectionViewLayoutAttributesClass().class), objc.Sel("layoutAttributesForItemWithIndexPath:"), indexPath)
	return rv
}


// Creates and returns a layout attributes object for a supplementary view based on the specified information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forSupplementaryViewOfKind:with:)
func NewCollectionViewLayoutAttributesForSupplementaryViewOfKindWithIndexPath(elementKind objc.IObject /* cross-framework: CollectionViewSupplementaryElementKind */, indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](objc.ID(getCollectionViewLayoutAttributesClass().class), objc.Sel("layoutAttributesForSupplementaryViewOfKind:withIndexPath:"), elementKind, indexPath)
	return rv
}



// Creates and returns a layout attributes object for a decoration view based on the specified information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forDecorationViewOfKind:with:)
func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForDecorationViewOfKindWithIndexPath(decorationViewKind objc.IObject /* cross-framework: CollectionViewDecorationElementKind */, indexPath foundation.IndexPath) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layoutAttributesForDecorationViewOfKind:withIndexPath:"), decorationViewKind, indexPath)
	return rv
}


// Creates and returns a layout attributes object for an inter-item gap view at the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forInterItemGapBefore:)
func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IndexPath) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layoutAttributesForInterItemGapBeforeIndexPath:"), indexPath)
	return rv
}


// Creates and returns a layout attributes object for the item at the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forItemWith:)
func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForItemWithIndexPath(indexPath foundation.IndexPath) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layoutAttributesForItemWithIndexPath:"), indexPath)
	return rv
}


// Creates and returns a layout attributes object for a supplementary view based on the specified information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forSupplementaryViewOfKind:with:)
func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForSupplementaryViewOfKindWithIndexPath(elementKind objc.IObject /* cross-framework: CollectionViewSupplementaryElementKind */, indexPath foundation.IndexPath) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layoutAttributesForSupplementaryViewOfKind:withIndexPath:"), elementKind, indexPath)
	return rv
}


// The transparency of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/alpha
func (c_ CollectionViewLayoutAttributes) Alpha() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("alpha"))
	return rv
}


// The transparency of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/alpha
func (c_ CollectionViewLayoutAttributes) SetAlpha(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlpha:"), value)
}


// The frame rectangle of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/frame
func (c_ CollectionViewLayoutAttributes) Frame() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](c_.ID, objc.Sel("frame"))
	return rv
}


// The frame rectangle of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/frame
func (c_ CollectionViewLayoutAttributes) SetFrame(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFrame:"), value)
}


// The index path of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/indexPath
func (c_ CollectionViewLayoutAttributes) IndexPath() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](c_.ID, objc.Sel("indexPath"))
	return rv
}


// The index path of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/indexPath
func (c_ CollectionViewLayoutAttributes) SetIndexPath(value foundation.IndexPath) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIndexPath:"), value)
}


// A Boolean value indicating whether the element is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/isHidden
func (c_ CollectionViewLayoutAttributes) Hidden() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hidden"))
	return rv
}


// A Boolean value indicating whether the element is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/isHidden
func (c_ CollectionViewLayoutAttributes) SetHidden(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHidden:"), value)
}


// The type of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/representedElementCategory
func (c_ CollectionViewLayoutAttributes) RepresentedElementCategory() CollectionElementCategory {
	rv := objc.Send[CollectionElementCategory](c_.ID, objc.Sel("representedElementCategory"))
	return rv
}


// The identifier for specific elements of your collection view interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/representedElementKind
func (c_ CollectionViewLayoutAttributes) RepresentedElementKind() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("representedElementKind"))
	return rv
}


// The size of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/size
func (c_ CollectionViewLayoutAttributes) Size() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](c_.ID, objc.Sel("size"))
	return rv
}


// The size of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/size
func (c_ CollectionViewLayoutAttributes) SetSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSize:"), value)
}


// The element’s position on the z axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/zIndex
func (c_ CollectionViewLayoutAttributes) ZIndex() int {
	rv := objc.Send[int](c_.ID, objc.Sel("zIndex"))
	return rv
}


// The element’s position on the z axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/zIndex
func (c_ CollectionViewLayoutAttributes) SetZIndex(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZIndex:"), value)
}


// A Boolean value indicating whether the element is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/ishidden
func (c_ CollectionViewLayoutAttributes) IsHidden() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHidden"))
	return rv
}


// A Boolean value indicating whether the element is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/ishidden
func (c_ CollectionViewLayoutAttributes) SetIsHidden(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHidden:"), value)
}



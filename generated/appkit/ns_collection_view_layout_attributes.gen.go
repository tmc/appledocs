// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCollectionViewLayoutAttributes */


/* debug [class_header]: Header for NSCollectionViewLayoutAttributes */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionViewLayoutAttributes */
// An interface definition for the [CollectionViewLayoutAttributes] class.
type ICollectionViewLayoutAttributes interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CollectionViewLayoutAttributes */
	// properties:
	Alpha() float64
	SetAlpha(value float64)
	Frame() Rect /* not a class type */
	SetFrame(value Rect /* not a class type */)
	IndexPath() foundation.IndexPath
	SetIndexPath(value foundation.IndexPath)
	Hidden() bool
	SetHidden(value bool)
	RepresentedElementCategory() CollectionElementCategory
	RepresentedElementKind() objc.IObject /* cross-framework: NSString */
	Size() Size /* not a class type */
	SetSize(value Size /* not a class type */)
	ZIndex() int
	SetZIndex(value int)
	IsHidden() bool
	SetIsHidden(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionViewLayoutAttributes */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionViewLayoutAttributes */
// Alloc allocates a new instance without initialization.
func (cc _CollectionViewLayoutAttributesClass) Alloc() CollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionViewLayoutAttributes */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionViewLayoutAttributes */

// Creates and returns a layout attributes object for a decoration view based on the specified information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forDecorationViewOfKind:with:)
func NewCollectionViewLayoutAttributesForDecorationViewOfKindWithIndexPath(decorationViewKind CollectionViewDecorationElementKind /* typedef */, indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](objc.ID(getCollectionViewLayoutAttributesClass().class), objc.Sel("layoutAttributesForDecorationViewOfKind:withIndexPath:"), decorationViewKind, indexPath)
	return rv
}/* debug [class_init_methods/constructor]: NewCollectionViewLayoutAttributesForDecorationViewOfKindWithIndexPath */


// Creates and returns a layout attributes object for an inter-item gap view at the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forInterItemGapBefore:)
func NewCollectionViewLayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](objc.ID(getCollectionViewLayoutAttributesClass().class), objc.Sel("layoutAttributesForInterItemGapBeforeIndexPath:"), indexPath)
	return rv
}/* debug [class_init_methods/constructor]: NewCollectionViewLayoutAttributesForInterItemGapBeforeIndexPath */


// Creates and returns a layout attributes object for the item at the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forItemWith:)
func NewCollectionViewLayoutAttributesForItemWithIndexPath(indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](objc.ID(getCollectionViewLayoutAttributesClass().class), objc.Sel("layoutAttributesForItemWithIndexPath:"), indexPath)
	return rv
}/* debug [class_init_methods/constructor]: NewCollectionViewLayoutAttributesForItemWithIndexPath */


// Creates and returns a layout attributes object for a supplementary view based on the specified information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forSupplementaryViewOfKind:with:)
func NewCollectionViewLayoutAttributesForSupplementaryViewOfKindWithIndexPath(elementKind CollectionViewSupplementaryElementKind /* typedef */, indexPath foundation.IndexPath) CollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](objc.ID(getCollectionViewLayoutAttributesClass().class), objc.Sel("layoutAttributesForSupplementaryViewOfKind:withIndexPath:"), elementKind, indexPath)
	return rv
}/* debug [class_init_methods/constructor]: NewCollectionViewLayoutAttributesForSupplementaryViewOfKindWithIndexPath */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionViewLayoutAttributes */

// Creates and returns a layout attributes object for a decoration view based on the specified information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forDecorationViewOfKind:with:)
func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForDecorationViewOfKindWithIndexPath(decorationViewKind CollectionViewDecorationElementKind /* typedef */, indexPath foundation.IndexPath) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("layoutAttributesForDecorationViewOfKind:withIndexPath:"), decorationViewKind, indexPath)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayoutAttributesForDecorationViewOfKindWithIndexPath) */


// Creates and returns a layout attributes object for an inter-item gap view at the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forInterItemGapBefore:)
func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IndexPath) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("layoutAttributesForInterItemGapBeforeIndexPath:"), indexPath)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayoutAttributesForInterItemGapBeforeIndexPath) */


// Creates and returns a layout attributes object for the item at the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forItemWith:)
func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForItemWithIndexPath(indexPath foundation.IndexPath) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("layoutAttributesForItemWithIndexPath:"), indexPath)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayoutAttributesForItemWithIndexPath) */


// Creates and returns a layout attributes object for a supplementary view based on the specified information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forSupplementaryViewOfKind:with:)
func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForSupplementaryViewOfKindWithIndexPath(elementKind CollectionViewSupplementaryElementKind /* typedef */, indexPath foundation.IndexPath) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("layoutAttributesForSupplementaryViewOfKind:withIndexPath:"), elementKind, indexPath)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayoutAttributesForSupplementaryViewOfKindWithIndexPath) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionViewLayoutAttributes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionViewLayoutAttributes */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionViewLayoutAttributes */

// The transparency of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/alpha
func (c_ CollectionViewLayoutAttributes) Alpha() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// The transparency of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/alpha
func (c_ CollectionViewLayoutAttributes) SetAlpha(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlpha:"), value)
}/* debug [instance_properties/setter]: alpha */


// The frame rectangle of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/frame
func (c_ CollectionViewLayoutAttributes) Frame() Rect /* not a class type */ {
	rv := objc.Send[Rect](c_.ID, objc.Sel("frame"))
	return rv
}/* debug [instance_properties/getter]: frame */


// The frame rectangle of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/frame
func (c_ CollectionViewLayoutAttributes) SetFrame(value Rect /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFrame:"), value)
}/* debug [instance_properties/setter]: frame */


// The index path of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/indexPath
func (c_ CollectionViewLayoutAttributes) IndexPath() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](c_.ID, objc.Sel("indexPath"))
	return rv
}/* debug [instance_properties/getter]: indexPath */


// The index path of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/indexPath
func (c_ CollectionViewLayoutAttributes) SetIndexPath(value foundation.IndexPath) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIndexPath:"), value)
}/* debug [instance_properties/setter]: indexPath */


// A Boolean value indicating whether the element is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/isHidden
func (c_ CollectionViewLayoutAttributes) Hidden() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hidden"))
	return rv
}/* debug [instance_properties/getter]: hidden */


// A Boolean value indicating whether the element is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/isHidden
func (c_ CollectionViewLayoutAttributes) SetHidden(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHidden:"), value)
}/* debug [instance_properties/setter]: hidden */


// The type of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/representedElementCategory
func (c_ CollectionViewLayoutAttributes) RepresentedElementCategory() CollectionElementCategory {
	rv := objc.Send[CollectionElementCategory](c_.ID, objc.Sel("representedElementCategory"))
	return rv
}/* debug [instance_properties/getter]: representedElementCategory */


// The identifier for specific elements of your collection view interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/representedElementKind
func (c_ CollectionViewLayoutAttributes) RepresentedElementKind() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("representedElementKind"))
	return rv
}/* debug [instance_properties/getter]: representedElementKind */


// The size of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/size
func (c_ CollectionViewLayoutAttributes) Size() Size /* not a class type */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */


// The size of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/size
func (c_ CollectionViewLayoutAttributes) SetSize(value Size /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSize:"), value)
}/* debug [instance_properties/setter]: size */


// The element’s position on the z axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/zIndex
func (c_ CollectionViewLayoutAttributes) ZIndex() int {
	rv := objc.Send[int](c_.ID, objc.Sel("zIndex"))
	return rv
}/* debug [instance_properties/getter]: zIndex */


// The element’s position on the z axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/zIndex
func (c_ CollectionViewLayoutAttributes) SetZIndex(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZIndex:"), value)
}/* debug [instance_properties/setter]: zIndex */


// A Boolean value indicating whether the element is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/ishidden
func (c_ CollectionViewLayoutAttributes) IsHidden() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHidden"))
	return rv
}/* debug [instance_properties/getter]: isHidden */


// A Boolean value indicating whether the element is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/ishidden
func (c_ CollectionViewLayoutAttributes) SetIsHidden(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHidden:"), value)
}/* debug [instance_properties/setter]: isHidden */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionViewLayoutAttributes */



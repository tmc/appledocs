// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCollectionViewLayoutInvalidationContext */


/* debug [class_header]: Header for NSCollectionViewLayoutInvalidationContext */
// The class instance for the [CollectionViewLayoutInvalidationContext] class.
var (
	CollectionViewLayoutInvalidationContextClass     _CollectionViewLayoutInvalidationContextClass
	CollectionViewLayoutInvalidationContextClassOnce sync.Once
)

func getCollectionViewLayoutInvalidationContextClass() _CollectionViewLayoutInvalidationContextClass {
	CollectionViewLayoutInvalidationContextClassOnce.Do(func() {
		CollectionViewLayoutInvalidationContextClass = _CollectionViewLayoutInvalidationContextClass{objc.GetClass("NSCollectionViewLayoutInvalidationContext")}
	})
	return CollectionViewLayoutInvalidationContextClass
}

type _CollectionViewLayoutInvalidationContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionViewLayoutInvalidationContext */
// An interface definition for the [CollectionViewLayoutInvalidationContext] class.
type ICollectionViewLayoutInvalidationContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CollectionViewLayoutInvalidationContext */
	// properties:
	ContentOffsetAdjustment() vision.Point
	SetContentOffsetAdjustment(value vision.Point)
	ContentSizeAdjustment() Size /* not a class type */
	SetContentSizeAdjustment(value Size /* not a class type */)
	InvalidateDataSourceCounts() bool
	InvalidateEverything() bool
	InvalidatedDecorationIndexPaths() foundation.IDictionary
	InvalidatedItemIndexPaths() unsafe.Pointer
	InvalidatedSupplementaryIndexPaths() foundation.IDictionary
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionViewLayoutInvalidationContext */
	// methods:
	InvalidateDecorationElementsOfKindAtIndexPaths(elementKind CollectionViewDecorationElementKind /* typedef */, indexPaths unsafe.Pointer)
	InvalidateItemsAtIndexPaths(indexPaths unsafe.Pointer)
	InvalidateSupplementaryElementsOfKindAtIndexPaths(elementKind CollectionViewSupplementaryElementKind /* typedef */, indexPaths unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionViewLayoutInvalidationContext */
// Alloc allocates a new instance without initialization.
func (cc _CollectionViewLayoutInvalidationContextClass) Alloc() CollectionViewLayoutInvalidationContext {
	rv := objc.Send[CollectionViewLayoutInvalidationContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionViewLayoutInvalidationContextClass) New() CollectionViewLayoutInvalidationContext {
	rv := objc.Send[CollectionViewLayoutInvalidationContext](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionViewLayoutInvalidationContext) Init() CollectionViewLayoutInvalidationContext {
	rv := objc.Send[CollectionViewLayoutInvalidationContext](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionViewLayoutInvalidationContext) Autorelease() CollectionViewLayoutInvalidationContext {
	rv := objc.Send[CollectionViewLayoutInvalidationContext](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionViewLayoutInvalidationContext creates a new CollectionViewLayoutInvalidationContext instance.
func NewCollectionViewLayoutInvalidationContext() CollectionViewLayoutInvalidationContext {
	return getCollectionViewLayoutInvalidationContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionViewLayoutInvalidationContext */
// An object that identifies the portions of your layout that need to be updated.
//
// Invalidation contexts are a way to improve the efficiency of layout operations and must be supported explicitly by the layout object. Instead of invalidating the entire layout, you can create an invalidation layout object that specifies only the portions of the layout that changed. You then pass that invalidation context to the method of the layout object. Typically, you ask the layout object to create an invalidation context for you. The class defines methods for creating a supported invalidation context. If you define a custom layout, you can define additional methods for creating invalidation contexts with custom information. Layout objects may also create invalidation contexts in response to specific changes. For example, layout objects automatically create invalidation contexts when you change the collection view’s data source, when you insert or delete items, and when you reload the collection view’s data.


// An object that identifies the portions of your layout that need to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext
type CollectionViewLayoutInvalidationContext struct {
	objectivec.Object
}

// CollectionViewLayoutInvalidationContextFrom constructs a [CollectionViewLayoutInvalidationContext] from an unsafe.Pointer.
//
// An object that identifies the portions of your layout that need to be updated.
func CollectionViewLayoutInvalidationContextFrom(ptr unsafe.Pointer) CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContext{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionViewLayoutInvalidationContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionViewLayoutInvalidationContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionViewLayoutInvalidationContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionViewLayoutInvalidationContext */

// Marks the specified decoration views as invalid so that their layout information can be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/invalidateDecorationElements(ofKind:at:)
func (c_ CollectionViewLayoutInvalidationContext) InvalidateDecorationElementsOfKindAtIndexPaths(elementKind CollectionViewDecorationElementKind /* typedef */, indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("invalidateDecorationElementsOfKind:atIndexPaths:"), elementKind, indexPaths)
}/* debug [instance_methods/method]: InvalidateDecorationElementsOfKindAtIndexPaths */


// Marks the specified items as invalid so that their layout information can be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/invalidateItems(at:)
func (c_ CollectionViewLayoutInvalidationContext) InvalidateItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("invalidateItemsAtIndexPaths:"), indexPaths)
}/* debug [instance_methods/method]: InvalidateItemsAtIndexPaths */


// Marks the specified supplementary views as invalid so that their layout information can be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/invalidateSupplementaryElements(ofKind:at:)
func (c_ CollectionViewLayoutInvalidationContext) InvalidateSupplementaryElementsOfKindAtIndexPaths(elementKind CollectionViewSupplementaryElementKind /* typedef */, indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("invalidateSupplementaryElementsOfKind:atIndexPaths:"), elementKind, indexPaths)
}/* debug [instance_methods/method]: InvalidateSupplementaryElementsOfKindAtIndexPaths */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionViewLayoutInvalidationContext */

// The delta value to add to the collection view’s content offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/contentOffsetAdjustment
func (c_ CollectionViewLayoutInvalidationContext) ContentOffsetAdjustment() vision.Point {
	rv := objc.Send[vision.Point](c_.ID, objc.Sel("contentOffsetAdjustment"))
	return rv
}/* debug [instance_properties/getter]: contentOffsetAdjustment */


// The delta value to add to the collection view’s content offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/contentOffsetAdjustment
func (c_ CollectionViewLayoutInvalidationContext) SetContentOffsetAdjustment(value vision.Point) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentOffsetAdjustment:"), value)
}/* debug [instance_properties/setter]: contentOffsetAdjustment */


// The delta value to add to the collection view’s content size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/contentSizeAdjustment
func (c_ CollectionViewLayoutInvalidationContext) ContentSizeAdjustment() Size /* not a class type */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("contentSizeAdjustment"))
	return rv
}/* debug [instance_properties/getter]: contentSizeAdjustment */


// The delta value to add to the collection view’s content size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/contentSizeAdjustment
func (c_ CollectionViewLayoutInvalidationContext) SetContentSizeAdjustment(value Size /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentSizeAdjustment:"), value)
}/* debug [instance_properties/setter]: contentSizeAdjustment */


// A Boolean that indicates whether the layout object should ask for new section and item counts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/invalidateDataSourceCounts
func (c_ CollectionViewLayoutInvalidationContext) InvalidateDataSourceCounts() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("invalidateDataSourceCounts"))
	return rv
}/* debug [instance_properties/getter]: invalidateDataSourceCounts */


// A Boolean that indicates whether all layout data should be marked as invalid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/invalidateEverything
func (c_ CollectionViewLayoutInvalidationContext) InvalidateEverything() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("invalidateEverything"))
	return rv
}/* debug [instance_properties/getter]: invalidateEverything */


// A dictionary containing the decoration views whose layout attributes are invalid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/invalidatedDecorationIndexPaths
func (c_ CollectionViewLayoutInvalidationContext) InvalidatedDecorationIndexPaths() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("invalidatedDecorationIndexPaths"))
	return rv
}/* debug [instance_properties/getter]: invalidatedDecorationIndexPaths */


// The set of items whose layout attributes are invalid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/invalidatedItemIndexPaths
func (c_ CollectionViewLayoutInvalidationContext) InvalidatedItemIndexPaths() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("invalidatedItemIndexPaths"))
	return rv
}/* debug [instance_properties/getter]: invalidatedItemIndexPaths */


// A dictionary containing the supplementary views whose layout attributes are invalid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/invalidatedSupplementaryIndexPaths
func (c_ CollectionViewLayoutInvalidationContext) InvalidatedSupplementaryIndexPaths() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("invalidatedSupplementaryIndexPaths"))
	return rv
}/* debug [instance_properties/getter]: invalidatedSupplementaryIndexPaths */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionViewLayoutInvalidationContext */




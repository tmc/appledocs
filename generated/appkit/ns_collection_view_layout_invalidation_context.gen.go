// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CollectionViewLayoutInvalidationContext] class.
type ICollectionViewLayoutInvalidationContext interface {
	objectivec.IObject
	ContentOffsetAdjustment() coregraphics.CGPoint
	SetContentOffsetAdjustment(value coregraphics.CGPoint)
	ContentSizeAdjustment() coregraphics.CGSize
	SetContentSizeAdjustment(value coregraphics.CGSize)
	InvalidateDataSourceCounts() bool
	SetInvalidateDataSourceCounts(value bool)
	InvalidateEverything() bool
	SetInvalidateEverything(value bool)
	InvalidatedDecorationIndexPaths() foundation.IndexPath
	SetInvalidatedDecorationIndexPaths(value foundation.IIndexPath)
	InvalidatedItemIndexPaths() foundation.IndexPath
	SetInvalidatedItemIndexPaths(value foundation.IIndexPath)
	InvalidatedSupplementaryIndexPaths() foundation.IndexPath
	SetInvalidatedSupplementaryIndexPaths(value foundation.IIndexPath)
}

// An object that identifies the portions of your layout that need to be updated.
//
// Invalidation contexts are a way to improve the efficiency of layout operations and must be supported explicitly by the layout object. Instead of invalidating the entire layout, you can create an invalidation layout object that specifies only the portions of the layout that changed. You then pass that invalidation context to the method of the layout object. Typically, you ask the layout object to create an invalidation context for you. The class defines methods for creating a supported invalidation context. If you define a custom layout, you can define additional methods for creating invalidation contexts with custom information. Layout objects may also create invalidation contexts in response to specific changes. For example, layout objects automatically create invalidation contexts when you change the collection view’s data source, when you insert or delete items, and when you reload the collection view’s data.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CollectionViewLayoutInvalidationContextClass) Alloc() CollectionViewLayoutInvalidationContext {
	rv := objc.Send[CollectionViewLayoutInvalidationContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The delta value to add to the collection view’s content offset.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/contentoffsetadjustment
func (c_ CollectionViewLayoutInvalidationContext) ContentOffsetAdjustment() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](c_.ID, objc.Sel("contentOffsetAdjustment"))
	return rv
}


// SetContentOffsetAdjustment sets the value of the contentOffsetAdjustment property.
// The delta value to add to the collection view’s content offset.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/contentoffsetadjustment
func (c_ CollectionViewLayoutInvalidationContext) SetContentOffsetAdjustment(value coregraphics.CGPoint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentOffsetAdjustment:"), value)
}

// The delta value to add to the collection view’s content size.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/contentsizeadjustment
func (c_ CollectionViewLayoutInvalidationContext) ContentSizeAdjustment() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("contentSizeAdjustment"))
	return rv
}


// SetContentSizeAdjustment sets the value of the contentSizeAdjustment property.
// The delta value to add to the collection view’s content size.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/contentsizeadjustment
func (c_ CollectionViewLayoutInvalidationContext) SetContentSizeAdjustment(value coregraphics.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentSizeAdjustment:"), value)
}

// A Boolean that indicates whether the layout object should ask for new section and item counts.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/invalidatedatasourcecounts
func (c_ CollectionViewLayoutInvalidationContext) InvalidateDataSourceCounts() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("invalidateDataSourceCounts"))
	return rv
}


// SetInvalidateDataSourceCounts sets the value of the invalidateDataSourceCounts property.
// A Boolean that indicates whether the layout object should ask for new section and item counts.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/invalidatedatasourcecounts
func (c_ CollectionViewLayoutInvalidationContext) SetInvalidateDataSourceCounts(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInvalidateDataSourceCounts:"), value)
}

// A Boolean that indicates whether all layout data should be marked as invalid.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/invalidateeverything
func (c_ CollectionViewLayoutInvalidationContext) InvalidateEverything() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("invalidateEverything"))
	return rv
}


// SetInvalidateEverything sets the value of the invalidateEverything property.
// A Boolean that indicates whether all layout data should be marked as invalid.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/invalidateeverything
func (c_ CollectionViewLayoutInvalidationContext) SetInvalidateEverything(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInvalidateEverything:"), value)
}

// A dictionary containing the decoration views whose layout attributes are invalid.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/invalidateddecorationindexpaths
func (c_ CollectionViewLayoutInvalidationContext) InvalidatedDecorationIndexPaths() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](c_.ID, objc.Sel("invalidatedDecorationIndexPaths"))
	return rv
}


// SetInvalidatedDecorationIndexPaths sets the value of the invalidatedDecorationIndexPaths property.
// A dictionary containing the decoration views whose layout attributes are invalid.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/invalidateddecorationindexpaths
func (c_ CollectionViewLayoutInvalidationContext) SetInvalidatedDecorationIndexPaths(value foundation.IIndexPath) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInvalidatedDecorationIndexPaths:"), value)
}

// The set of items whose layout attributes are invalid.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/invalidateditemindexpaths
func (c_ CollectionViewLayoutInvalidationContext) InvalidatedItemIndexPaths() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](c_.ID, objc.Sel("invalidatedItemIndexPaths"))
	return rv
}


// SetInvalidatedItemIndexPaths sets the value of the invalidatedItemIndexPaths property.
// The set of items whose layout attributes are invalid.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/invalidateditemindexpaths
func (c_ CollectionViewLayoutInvalidationContext) SetInvalidatedItemIndexPaths(value foundation.IIndexPath) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInvalidatedItemIndexPaths:"), value)
}

// A dictionary containing the supplementary views whose layout attributes are invalid.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/invalidatedsupplementaryindexpaths
func (c_ CollectionViewLayoutInvalidationContext) InvalidatedSupplementaryIndexPaths() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](c_.ID, objc.Sel("invalidatedSupplementaryIndexPaths"))
	return rv
}


// SetInvalidatedSupplementaryIndexPaths sets the value of the invalidatedSupplementaryIndexPaths property.
// A dictionary containing the supplementary views whose layout attributes are invalid.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutinvalidationcontext/invalidatedsupplementaryindexpaths
func (c_ CollectionViewLayoutInvalidationContext) SetInvalidatedSupplementaryIndexPaths(value foundation.IIndexPath) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInvalidatedSupplementaryIndexPaths:"), value)
}




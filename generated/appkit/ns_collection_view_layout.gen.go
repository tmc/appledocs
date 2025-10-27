// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CollectionViewLayout] class.
var (
	CollectionViewLayoutClass     _CollectionViewLayoutClass
	CollectionViewLayoutClassOnce sync.Once
)

func getCollectionViewLayoutClass() _CollectionViewLayoutClass {
	CollectionViewLayoutClassOnce.Do(func() {
		CollectionViewLayoutClass = _CollectionViewLayoutClass{objc.GetClass("NSCollectionViewLayout")}
	})
	return CollectionViewLayoutClass
}

type _CollectionViewLayoutClass struct {
	class objc.Class
}





// An interface definition for the [CollectionViewLayout] class.
type ICollectionViewLayout interface {
	objectivec.IObject
	

	// properties:
	CollectionView() CollectionView /* not a class type */
	CollectionViewContentSize() corefoundation.CGSize
	CollectionViewLayout() ICollectionViewLayout
	SetCollectionViewLayout(value ICollectionViewLayout)


	

	// methods:
	FinalLayoutAttributesForDisappearingDecorationElementOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.foundation.INSIndexPath) ICollectionViewLayoutAttributes
	FinalLayoutAttributesForDisappearingItemAtIndexPath(itemIndexPath foundation.foundation.INSIndexPath) ICollectionViewLayoutAttributes
	FinalLayoutAttributesForDisappearingSupplementaryElementOfKindAtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.foundation.INSIndexPath) ICollectionViewLayoutAttributes
	FinalizeAnimatedBoundsChange()
	FinalizeCollectionViewUpdates()
	FinalizeLayoutTransition()
	IndexPathsToDeleteForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) unsafe.Pointer
	IndexPathsToDeleteForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) unsafe.Pointer
	IndexPathsToInsertForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) unsafe.Pointer
	IndexPathsToInsertForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) unsafe.Pointer
	InvalidateLayout()
	InvalidateLayoutWithContext(context ICollectionViewLayoutInvalidationContext)
	InvalidationContextForBoundsChange(newBounds corefoundation.CGRect) ICollectionViewLayoutInvalidationContext
	InvalidationContextForPreferredLayoutAttributesWithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) ICollectionViewLayoutInvalidationContext
	LayoutAttributesForDecorationViewOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, indexPath foundation.foundation.INSIndexPath) ICollectionViewLayoutAttributes
	LayoutAttributesForDropTargetAtPoint(pointInCollectionView corefoundation.CGPoint) ICollectionViewLayoutAttributes
	LayoutAttributesForElementsInRect(rect corefoundation.CGRect) []CollectionViewLayoutAttributes
	LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.foundation.INSIndexPath) ICollectionViewLayoutAttributes
	LayoutAttributesForItemAtIndexPath(indexPath foundation.foundation.INSIndexPath) ICollectionViewLayoutAttributes
	LayoutAttributesForSupplementaryViewOfKindAtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.foundation.INSIndexPath) ICollectionViewLayoutAttributes
	PrepareLayout()
	PrepareForAnimatedBoundsChange(oldBounds corefoundation.CGRect)
	PrepareForCollectionViewUpdates(updateItems []CollectionViewUpdateItem)
	PrepareForTransitionFromLayout(oldLayout ICollectionViewLayout)
	PrepareForTransitionToLayout(newLayout ICollectionViewLayout)
	RegisterClassForDecorationViewOfKind(viewClass objc.Class, elementKind CollectionViewDecorationElementKind)
	RegisterNibForDecorationViewOfKind(nib INib, elementKind CollectionViewDecorationElementKind)
	ShouldInvalidateLayoutForBoundsChange(newBounds corefoundation.CGRect) bool
	ShouldInvalidateLayoutForPreferredLayoutAttributesWithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) bool
	TargetContentOffsetForProposedContentOffset(proposedContentOffset corefoundation.CGPoint) corefoundation.CGPoint
	TargetContentOffsetForProposedContentOffsetWithScrollingVelocity(proposedContentOffset corefoundation.CGPoint, velocity corefoundation.CGPoint) corefoundation.CGPoint


}





// Alloc allocates a new instance without initialization.
func (cc _CollectionViewLayoutClass) Alloc() CollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionViewLayoutClass) New() CollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionViewLayout) Init() CollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionViewLayout) Autorelease() CollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionViewLayout creates a new CollectionViewLayout instance.
func NewCollectionViewLayout() CollectionViewLayout {
	return getCollectionViewLayoutClass().New()
}





// An abstract base class that you subclass and use to generate layout information for a collection view.
//
// The job of a layout object is to perform the calculations needed to determine the placement and appearance of items, supplementary views, and other content in the collection view. The layout object does not apply the layout attributes it generates to the views in your interface. Instead, it passes those layout attributes to the collection view, which then creates the needed views and applies the layout attributes to them. You do not create instances of this class directly. Instead, you create instances of one of its subclasses and associate that object with your collection view either programmatically (using the property) or at design time in Interface Builder. Changing the layout object of a collection view forces an immediate update of the layout information. Collection views support many different types of elements, most of which are visual and all of which require layout attributes: are the main elements managed by the layout. Each item represents a single piece of data in the collection view. A collection view can have a single group of items or it can divide the items into multiple sections. are optional views associated with a specific section. The layout object defines the placement and use of supplementary views. For example, grid and flow layouts use supplementary views to implement headers and footers for each section. Supplementary views cannot be selected by the user. are visual adornments used to implement themes or to present visual content that is unrelated to the data being managed by the collection view. Decoration views are optional and the layout object defines their use and placement. supply a drop target for dragged content. Gaps do not have a direct visual representation, but they do have layout attributes, which the collection view uses for hit testing. The layout object provides attributes for inter-item gaps only when asked to do so. Each concrete layout object defines a specific organization for the contained elements and provides the appropriate layout attributes. The placement and appearance of items is determined entirely by the layout object. The and subclasses define variants of a grid-based layout, but you can create custom layouts that arrange elements in different ways. For example, you might define a layout class that arranges items in a circle or define a class that groups items into stacks that resemble a pile of photos on a table.


// An abstract base class that you subclass and use to generate layout information for a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout
type CollectionViewLayout struct {
	objectivec.Object
}

// CollectionViewLayoutFrom constructs a [CollectionViewLayout] from an unsafe.Pointer.
//
// An abstract base class that you subclass and use to generate layout information for a collection view.
func CollectionViewLayoutFrom(ptr unsafe.Pointer) CollectionViewLayout {
	return CollectionViewLayout{objectivec.Object{objc.ID(ptr)}}
}






// Returns the starting layout information for a decoration view being added to the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/initialLayoutAttributesForAppearingDecorationElement(ofKind:at:)
func NewCollectionViewLayoutialLayoutAttributesForAppearingDecorationElementOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.foundation.INSIndexPath) CollectionViewLayout {
	instance := getCollectionViewLayoutClass().Alloc()
	rv := objc.Send[CollectionViewLayout](instance.ID, objc.Sel("initialLayoutAttributesForAppearingDecorationElementOfKind:atIndexPath:"), elementKind, decorationIndexPath)
	rv.Autorelease()
	return rv
}


// Returns the starting layout information for an item being inserted into the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/initialLayoutAttributesForAppearingItem(at:)
func NewCollectionViewLayoutialLayoutAttributesForAppearingItemAtIndexPath(itemIndexPath foundation.foundation.INSIndexPath) CollectionViewLayout {
	instance := getCollectionViewLayoutClass().Alloc()
	rv := objc.Send[CollectionViewLayout](instance.ID, objc.Sel("initialLayoutAttributesForAppearingItemAtIndexPath:"), itemIndexPath)
	rv.Autorelease()
	return rv
}


// Returns the starting layout information for a supplementary view being added to the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/initialLayoutAttributesForAppearingSupplementaryElement(ofKind:at:)
func NewCollectionViewLayoutialLayoutAttributesForAppearingSupplementaryElementOfKindAtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.foundation.INSIndexPath) CollectionViewLayout {
	instance := getCollectionViewLayoutClass().Alloc()
	rv := objc.Send[CollectionViewLayout](instance.ID, objc.Sel("initialLayoutAttributesForAppearingSupplementaryElementOfKind:atIndexPath:"), elementKind, elementIndexPath)
	rv.Autorelease()
	return rv
}












// Returns the class to use when creating an invalidation context object for the layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/invalidationContextClass
func (cc _CollectionViewLayoutClass) InvalidationContextClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(cc.class), objc.Sel("invalidationContextClass"))
	return rv
}

// Returns the class to use for layout attribute objects
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/layoutAttributesClass
func (cc _CollectionViewLayoutClass) LayoutAttributesClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(cc.class), objc.Sel("layoutAttributesClass"))
	return rv
}






// Returns the ending layout information for a decoration view being removed from the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/finalLayoutAttributesForDisappearingDecorationElement(ofKind:at:)
func (c_ CollectionViewLayout) FinalLayoutAttributesForDisappearingDecorationElementOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.foundation.INSIndexPath) ICollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](c_.ID, objc.Sel("finalLayoutAttributesForDisappearingDecorationElementOfKind:atIndexPath:"), elementKind, decorationIndexPath)
	return rv
}


// Returns the ending layout information for an item being removed from the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/finalLayoutAttributesForDisappearingItem(at:)
func (c_ CollectionViewLayout) FinalLayoutAttributesForDisappearingItemAtIndexPath(itemIndexPath foundation.foundation.INSIndexPath) ICollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](c_.ID, objc.Sel("finalLayoutAttributesForDisappearingItemAtIndexPath:"), itemIndexPath)
	return rv
}


// Returns the ending layout information for a supplementary view being removed from the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/finalLayoutAttributesForDisappearingSupplementaryElement(ofKind:at:)
func (c_ CollectionViewLayout) FinalLayoutAttributesForDisappearingSupplementaryElementOfKindAtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.foundation.INSIndexPath) ICollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](c_.ID, objc.Sel("finalLayoutAttributesForDisappearingSupplementaryElementOfKind:atIndexPath:"), elementKind, elementIndexPath)
	return rv
}


// Cleans up after any animated changes to the collection view’s bounds or after the insertion or deletion of items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/finalizeAnimatedBoundsChange()
func (c_ CollectionViewLayout) FinalizeAnimatedBoundsChange() {
	objc.Send[objc.ID](c_.ID, objc.Sel("finalizeAnimatedBoundsChange"))
}


// Performs needed steps after items are inserted, deleted, or moved within a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/finalizeCollectionViewUpdates()
func (c_ CollectionViewLayout) FinalizeCollectionViewUpdates() {
	objc.Send[objc.ID](c_.ID, objc.Sel("finalizeCollectionViewUpdates"))
}


// Performs any final steps related to a layout transition before the transition animations actually occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/finalizeLayoutTransition()
func (c_ CollectionViewLayout) FinalizeLayoutTransition() {
	objc.Send[objc.ID](c_.ID, objc.Sel("finalizeLayoutTransition"))
}


// Returns index paths for any decoration views that the layout object wants to remove from the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/indexPathsToDeleteForDecorationView(ofKind:)
func (c_ CollectionViewLayout) IndexPathsToDeleteForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("indexPathsToDeleteForDecorationViewOfKind:"), elementKind)
	return rv
}


// Returns the index paths for any supplementary views that the layout object wants to remove from the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/indexPathsToDeleteForSupplementaryView(ofKind:)
func (c_ CollectionViewLayout) IndexPathsToDeleteForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("indexPathsToDeleteForSupplementaryViewOfKind:"), elementKind)
	return rv
}


// Returns the index paths for any decoration views that the layout object wants to add to the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/indexPathsToInsertForDecorationView(ofKind:)
func (c_ CollectionViewLayout) IndexPathsToInsertForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("indexPathsToInsertForDecorationViewOfKind:"), elementKind)
	return rv
}


// Returns the index paths for any supplementary views that the layout object wants to add to the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/indexPathsToInsertForSupplementaryView(ofKind:)
func (c_ CollectionViewLayout) IndexPathsToInsertForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("indexPathsToInsertForSupplementaryViewOfKind:"), elementKind)
	return rv
}


// Invalidates all layout information and triggers a layout update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/invalidateLayout()
func (c_ CollectionViewLayout) InvalidateLayout() {
	objc.Send[objc.ID](c_.ID, objc.Sel("invalidateLayout"))
}


// Invalidates specific parts of the layout using the specified context object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/invalidateLayout(with:)
func (c_ CollectionViewLayout) InvalidateLayoutWithContext(context ICollectionViewLayoutInvalidationContext) {
	objc.Send[objc.ID](c_.ID, objc.Sel("invalidateLayoutWithContext:"), context)
}


// Returns an invalidation context object that defines the portions of the layout that need to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/invalidationContext(forBoundsChange:)
func (c_ CollectionViewLayout) InvalidationContextForBoundsChange(newBounds corefoundation.CGRect) ICollectionViewLayoutInvalidationContext {
	rv := objc.Send[CollectionViewLayoutInvalidationContext](c_.ID, objc.Sel("invalidationContextForBoundsChange:"), newBounds)
	return rv
}


// Returns an invalidation context object that defines the portions of the layout that need to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/invalidationContext(forPreferredLayoutAttributes:withOriginalAttributes:)
func (c_ CollectionViewLayout) InvalidationContextForPreferredLayoutAttributesWithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) ICollectionViewLayoutInvalidationContext {
	rv := objc.Send[CollectionViewLayoutInvalidationContext](c_.ID, objc.Sel("invalidationContextForPreferredLayoutAttributes:withOriginalAttributes:"), preferredAttributes, originalAttributes)
	return rv
}


// Returns the layout attributes of the decoration view at the specified location in your layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/layoutAttributesForDecorationView(ofKind:at:)
func (c_ CollectionViewLayout) LayoutAttributesForDecorationViewOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, indexPath foundation.foundation.INSIndexPath) ICollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](c_.ID, objc.Sel("layoutAttributesForDecorationViewOfKind:atIndexPath:"), elementKind, indexPath)
	return rv
}


// Returns layout attributes for the drop target at the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/layoutAttributesForDropTarget(at:)
func (c_ CollectionViewLayout) LayoutAttributesForDropTargetAtPoint(pointInCollectionView corefoundation.CGPoint) ICollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](c_.ID, objc.Sel("layoutAttributesForDropTargetAtPoint:"), pointInCollectionView)
	return rv
}


// Returns the layout attribute objects for all items and views in the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/layoutAttributesForElements(in:)
func (c_ CollectionViewLayout) LayoutAttributesForElementsInRect(rect corefoundation.CGRect) []CollectionViewLayoutAttributes {
	rv := objc.Send[[]CollectionViewLayoutAttributes](c_.ID, objc.Sel("layoutAttributesForElementsInRect:"), rect)
	return rv
}


// Returns layout attributes for the inter-item gap at the specified location in your layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/layoutAttributesForInterItemGap(before:)
func (c_ CollectionViewLayout) LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.foundation.INSIndexPath) ICollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](c_.ID, objc.Sel("layoutAttributesForInterItemGapBeforeIndexPath:"), indexPath)
	return rv
}


// Returns the layout attributes for the item at the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/layoutAttributesForItem(at:)
func (c_ CollectionViewLayout) LayoutAttributesForItemAtIndexPath(indexPath foundation.foundation.INSIndexPath) ICollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](c_.ID, objc.Sel("layoutAttributesForItemAtIndexPath:"), indexPath)
	return rv
}


// Returns the layout attributes of the supplementary view at the specified location in your layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/layoutAttributesForSupplementaryView(ofKind:at:)
func (c_ CollectionViewLayout) LayoutAttributesForSupplementaryViewOfKindAtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.foundation.INSIndexPath) ICollectionViewLayoutAttributes {
	rv := objc.Send[CollectionViewLayoutAttributes](c_.ID, objc.Sel("layoutAttributesForSupplementaryViewOfKind:atIndexPath:"), elementKind, indexPath)
	return rv
}


// Prepares the layout object to begin laying out content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/prepare()
func (c_ CollectionViewLayout) PrepareLayout() {
	objc.Send[objc.ID](c_.ID, objc.Sel("prepareLayout"))
}


// Prepares the layout object for animated changes to the collection view’s bounds or for the insertion or deletion of items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/prepare(forAnimatedBoundsChange:)
func (c_ CollectionViewLayout) PrepareForAnimatedBoundsChange(oldBounds corefoundation.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("prepareForAnimatedBoundsChange:"), oldBounds)
}


// Performs needed tasks before items are inserted, deleted, or moved within the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/prepare(forCollectionViewUpdates:)
func (c_ CollectionViewLayout) PrepareForCollectionViewUpdates(updateItems []CollectionViewUpdateItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("prepareForCollectionViewUpdates:"), updateItems)
}


// Prepares the layout object to be installed in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/prepareForTransition(from:)
func (c_ CollectionViewLayout) PrepareForTransitionFromLayout(oldLayout ICollectionViewLayout) {
	objc.Send[objc.ID](c_.ID, objc.Sel("prepareForTransitionFromLayout:"), oldLayout)
}


// Prepares the layout object to be uninstalled from the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/prepareForTransition(to:)
func (c_ CollectionViewLayout) PrepareForTransitionToLayout(newLayout ICollectionViewLayout) {
	objc.Send[objc.ID](c_.ID, objc.Sel("prepareForTransitionToLayout:"), newLayout)
}


// Registers a class to use when creating the layout’s decoration views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/register(_:forDecorationViewOfKind:)-44qmc
func (c_ CollectionViewLayout) RegisterClassForDecorationViewOfKind(viewClass objc.Class, elementKind CollectionViewDecorationElementKind) {
	objc.Send[objc.ID](c_.ID, objc.Sel("registerClass:forDecorationViewOfKind:"), viewClass, elementKind)
}


// Registers a nib file to use when creating the layout’s decoration views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/register(_:forDecorationViewOfKind:)-7z7uf
func (c_ CollectionViewLayout) RegisterNibForDecorationViewOfKind(nib INib, elementKind CollectionViewDecorationElementKind) {
	objc.Send[objc.ID](c_.ID, objc.Sel("registerNib:forDecorationViewOfKind:"), nib, elementKind)
}


// Returns a Boolean indicating whether a bounds change triggers a layout update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/shouldInvalidateLayout(forBoundsChange:)
func (c_ CollectionViewLayout) ShouldInvalidateLayoutForBoundsChange(newBounds corefoundation.CGRect) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldInvalidateLayoutForBoundsChange:"), newBounds)
	return rv
}


// Returns a Boolean indicating whether changes to a cell’s layout attributes trigger a larger layout update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/shouldInvalidateLayout(forPreferredLayoutAttributes:withOriginalAttributes:)
func (c_ CollectionViewLayout) ShouldInvalidateLayoutForPreferredLayoutAttributesWithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldInvalidateLayoutForPreferredLayoutAttributes:withOriginalAttributes:"), preferredAttributes, originalAttributes)
	return rv
}


// Returns the offset value to use after an animated layout update or change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/targetContentOffset(forProposedContentOffset:)
func (c_ CollectionViewLayout) TargetContentOffsetForProposedContentOffset(proposedContentOffset corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c_.ID, objc.Sel("targetContentOffsetForProposedContentOffset:"), proposedContentOffset)
	return rv
}


// Returns the offset value to use for the collection view’s content at the end of scrolling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/targetContentOffset(forProposedContentOffset:withScrollingVelocity:)
func (c_ CollectionViewLayout) TargetContentOffsetForProposedContentOffsetWithScrollingVelocity(proposedContentOffset corefoundation.CGPoint, velocity corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c_.ID, objc.Sel("targetContentOffsetForProposedContentOffset:withScrollingVelocity:"), proposedContentOffset, velocity)
	return rv
}







// The collection view object currently using this layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/collectionView
func (c_ CollectionViewLayout) CollectionView() CollectionView /* not a class type */ {
	rv := objc.Send[CollectionView](c_.ID, objc.Sel("collectionView"))
	return rv
}


// The width and height of the collection view’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/collectionViewContentSize
func (c_ CollectionViewLayout) CollectionViewContentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("collectionViewContentSize"))
	return rv
}


// Returns the class to use when creating an invalidation context object for the layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/invalidationContextClass
func (c_ CollectionViewLayout) InvalidationContextClass() objc.Class {
	rv := objc.Send[objc.Class](c_.ID, objc.Sel("invalidationContextClass"))
	return rv
}


// Returns the class to use for layout attribute objects
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/layoutAttributesClass
func (c_ CollectionViewLayout) LayoutAttributesClass() objc.Class {
	rv := objc.Send[objc.Class](c_.ID, objc.Sel("layoutAttributesClass"))
	return rv
}


// The layout object used to organize the collection view’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/collectionviewlayout
func (c_ CollectionViewLayout) CollectionViewLayout() ICollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](c_.ID, objc.Sel("collectionViewLayout"))
	return rv
}


// The layout object used to organize the collection view’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/collectionviewlayout
func (c_ CollectionViewLayout) SetCollectionViewLayout(value ICollectionViewLayout) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCollectionViewLayout:"), value)
}








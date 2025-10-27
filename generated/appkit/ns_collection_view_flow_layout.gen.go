// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)





// The class instance for the [CollectionViewFlowLayout] class.
var (
	CollectionViewFlowLayoutClass     _CollectionViewFlowLayoutClass
	CollectionViewFlowLayoutClassOnce sync.Once
)

func getCollectionViewFlowLayoutClass() _CollectionViewFlowLayoutClass {
	CollectionViewFlowLayoutClassOnce.Do(func() {
		CollectionViewFlowLayoutClass = _CollectionViewFlowLayoutClass{objc.GetClass("NSCollectionViewFlowLayout")}
	})
	return CollectionViewFlowLayoutClass
}

type _CollectionViewFlowLayoutClass struct {
	class objc.Class
}





// An interface definition for the [CollectionViewFlowLayout] class.
type ICollectionViewFlowLayout interface {
	ICollectionViewLayout
	

	// properties:
	EstimatedItemSize() corefoundation.CGSize
	SetEstimatedItemSize(value corefoundation.CGSize)
	FooterReferenceSize() corefoundation.CGSize
	SetFooterReferenceSize(value corefoundation.CGSize)
	HeaderReferenceSize() corefoundation.CGSize
	SetHeaderReferenceSize(value corefoundation.CGSize)
	ItemSize() corefoundation.CGSize
	SetItemSize(value corefoundation.CGSize)
	MinimumInteritemSpacing() float64
	SetMinimumInteritemSpacing(value float64)
	MinimumLineSpacing() float64
	SetMinimumLineSpacing(value float64)
	ScrollDirection() CollectionViewScrollDirection
	SetScrollDirection(value CollectionViewScrollDirection)
	SectionFootersPinToVisibleBounds() bool
	SetSectionFootersPinToVisibleBounds(value bool)
	SectionHeadersPinToVisibleBounds() bool
	SetSectionHeadersPinToVisibleBounds(value bool)
	SectionInset() foundation.EdgeInsets
	SetSectionInset(value foundation.EdgeInsets)
	NSZeroSize() corefoundation.CGSize


	

	// methods:
	CollapseSectionAtIndex(sectionIndex uint)
	ExpandSectionAtIndex(sectionIndex uint)
	SectionAtIndexIsCollapsed(sectionIndex uint) bool


}





// Alloc allocates a new instance without initialization.
func (cc _CollectionViewFlowLayoutClass) Alloc() CollectionViewFlowLayout {
	rv := objc.Send[CollectionViewFlowLayout](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionViewFlowLayoutClass) New() CollectionViewFlowLayout {
	rv := objc.Send[CollectionViewFlowLayout](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionViewFlowLayout) Init() CollectionViewFlowLayout {
	rv := objc.Send[CollectionViewFlowLayout](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionViewFlowLayout) Autorelease() CollectionViewFlowLayout {
	rv := objc.Send[CollectionViewFlowLayout](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionViewFlowLayout creates a new CollectionViewFlowLayout instance.
func NewCollectionViewFlowLayout() CollectionViewFlowLayout {
	return getCollectionViewFlowLayoutClass().New()
}





// A layout that organizes items into a flexible and configurable arrangement.
//
// In a flow layout, the first item is positioned in the top-left corner and other items are laid out either horizontally or vertically based on the scroll direction, which is configurable. Items may be the same size or different sizes, and you may use the flow layout object or the collection view’s delegate object to specify the size of items and the spacing around them. The flow layout also lets you specify custom header and footer views for each section. You can use an object as-is or subclass it to modify more aspects of the layout behavior. There are several ways to customize the basic layout behavior that do not require subclassing. For example, you can use a delegate object to change the size and spacing of items dynamically. Subclassing is appropriate for more advanced layout changes, such as adding supplementary views or decoration views, supporting custom layout attributes, or customizing the layout animations when inserting or deleting items.


// A layout that organizes items into a flexible and configurable arrangement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout
type CollectionViewFlowLayout struct {
	CollectionViewLayout
}

// CollectionViewFlowLayoutFrom constructs a [CollectionViewFlowLayout] from an unsafe.Pointer.
//
// A layout that organizes items into a flexible and configurable arrangement.
func CollectionViewFlowLayoutFrom(ptr unsafe.Pointer) CollectionViewFlowLayout {
	return CollectionViewFlowLayout{
		CollectionViewLayout: CollectionViewLayoutFrom(ptr),
	}
}




















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/collapseSection(at:)
func (c_ CollectionViewFlowLayout) CollapseSectionAtIndex(sectionIndex uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("collapseSectionAtIndex:"), sectionIndex)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/expandSection(at:)
func (c_ CollectionViewFlowLayout) ExpandSectionAtIndex(sectionIndex uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("expandSectionAtIndex:"), sectionIndex)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/section(atIndexIsCollapsed:)
func (c_ CollectionViewFlowLayout) SectionAtIndexIsCollapsed(sectionIndex uint) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sectionAtIndexIsCollapsed:"), sectionIndex)
	return rv
}







// The estimated size of items in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/estimatedItemSize
func (c_ CollectionViewFlowLayout) EstimatedItemSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("estimatedItemSize"))
	return rv
}


// The estimated size of items in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/estimatedItemSize
func (c_ CollectionViewFlowLayout) SetEstimatedItemSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEstimatedItemSize:"), value)
}


// The default size to use for section footers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/footerReferenceSize
func (c_ CollectionViewFlowLayout) FooterReferenceSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("footerReferenceSize"))
	return rv
}


// The default size to use for section footers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/footerReferenceSize
func (c_ CollectionViewFlowLayout) SetFooterReferenceSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFooterReferenceSize:"), value)
}


// The default size to use for section headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/headerReferenceSize
func (c_ CollectionViewFlowLayout) HeaderReferenceSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("headerReferenceSize"))
	return rv
}


// The default size to use for section headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/headerReferenceSize
func (c_ CollectionViewFlowLayout) SetHeaderReferenceSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHeaderReferenceSize:"), value)
}


// The default size to use for items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/itemSize
func (c_ CollectionViewFlowLayout) ItemSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("itemSize"))
	return rv
}


// The default size to use for items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/itemSize
func (c_ CollectionViewFlowLayout) SetItemSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setItemSize:"), value)
}


// The minimum spacing (in points) to use between items in the same row or column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/minimumInteritemSpacing
func (c_ CollectionViewFlowLayout) MinimumInteritemSpacing() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("minimumInteritemSpacing"))
	return rv
}


// The minimum spacing (in points) to use between items in the same row or column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/minimumInteritemSpacing
func (c_ CollectionViewFlowLayout) SetMinimumInteritemSpacing(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumInteritemSpacing:"), value)
}


// The minimum spacing (in points) to use between rows or columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/minimumLineSpacing
func (c_ CollectionViewFlowLayout) MinimumLineSpacing() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("minimumLineSpacing"))
	return rv
}


// The minimum spacing (in points) to use between rows or columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/minimumLineSpacing
func (c_ CollectionViewFlowLayout) SetMinimumLineSpacing(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumLineSpacing:"), value)
}


// The scroll direction of the layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/scrollDirection
func (c_ CollectionViewFlowLayout) ScrollDirection() CollectionViewScrollDirection {
	rv := objc.Send[CollectionViewScrollDirection](c_.ID, objc.Sel("scrollDirection"))
	return rv
}


// The scroll direction of the layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/scrollDirection
func (c_ CollectionViewFlowLayout) SetScrollDirection(value CollectionViewScrollDirection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScrollDirection:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/sectionFootersPinToVisibleBounds
func (c_ CollectionViewFlowLayout) SectionFootersPinToVisibleBounds() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sectionFootersPinToVisibleBounds"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/sectionFootersPinToVisibleBounds
func (c_ CollectionViewFlowLayout) SetSectionFootersPinToVisibleBounds(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSectionFootersPinToVisibleBounds:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/sectionHeadersPinToVisibleBounds
func (c_ CollectionViewFlowLayout) SectionHeadersPinToVisibleBounds() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sectionHeadersPinToVisibleBounds"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/sectionHeadersPinToVisibleBounds
func (c_ CollectionViewFlowLayout) SetSectionHeadersPinToVisibleBounds(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSectionHeadersPinToVisibleBounds:"), value)
}


// The margins used to lay out content in a section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/sectionInset
func (c_ CollectionViewFlowLayout) SectionInset() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](c_.ID, objc.Sel("sectionInset"))
	return rv
}


// The margins used to lay out content in a section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/sectionInset
func (c_ CollectionViewFlowLayout) SetSectionInset(value foundation.EdgeInsets) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSectionInset:"), value)
}


// An `NSSize` structure set to `0` in both dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSZeroSize
func (c_ CollectionViewFlowLayout) NSZeroSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("NSZeroSize"))
	return rv
}









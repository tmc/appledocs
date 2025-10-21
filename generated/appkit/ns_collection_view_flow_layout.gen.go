// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
}

// A layout that organizes items into a flexible and configurable arrangement.
//
// In a flow layout, the first item is positioned in the top-left corner and other items are laid out either horizontally or vertically based on the scroll direction, which is configurable. Items may be the same size or different sizes, and you may use the flow layout object or the collection view’s delegate object to specify the size of items and the spacing around them. The flow layout also lets you specify custom header and footer views for each section. You can use an object as-is or subclass it to modify more aspects of the layout behavior. There are several ways to customize the basic layout behavior that do not require subclassing. For example, you can use a delegate object to change the size and spacing of items dynamically. Subclassing is appropriate for more advanced layout changes, such as adding supplementary views or decoration views, supporting custom layout attributes, or customizing the layout animations when inserting or deleting items.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CollectionViewFlowLayoutClass) Alloc() CollectionViewFlowLayout {
	rv := objc.Send[CollectionViewFlowLayout](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The layout object used to organize the collection view’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/collectionviewlayout
func (c_ CollectionViewFlowLayout) CollectionViewLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("collectionViewLayout"))
	return rv
}


// SetCollectionViewLayout sets the value of the collectionViewLayout property.
// The layout object used to organize the collection view’s content.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/collectionviewlayout
func (c_ CollectionViewFlowLayout) SetCollectionViewLayout(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCollectionViewLayout:"), value)
}

// The collection view’s delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/delegate
func (c_ CollectionViewFlowLayout) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The collection view’s delegate object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/delegate
func (c_ CollectionViewFlowLayout) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}

// The estimated size of items in the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/estimateditemsize
func (c_ CollectionViewFlowLayout) EstimatedItemSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("estimatedItemSize"))
	return rv
}


// SetEstimatedItemSize sets the value of the estimatedItemSize property.
// The estimated size of items in the collection view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/estimateditemsize
func (c_ CollectionViewFlowLayout) SetEstimatedItemSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEstimatedItemSize:"), value)
}

// The default size to use for section footers.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/footerreferencesize
func (c_ CollectionViewFlowLayout) FooterReferenceSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("footerReferenceSize"))
	return rv
}


// SetFooterReferenceSize sets the value of the footerReferenceSize property.
// The default size to use for section footers.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/footerreferencesize
func (c_ CollectionViewFlowLayout) SetFooterReferenceSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFooterReferenceSize:"), value)
}

// The default size to use for section headers.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/headerreferencesize
func (c_ CollectionViewFlowLayout) HeaderReferenceSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("headerReferenceSize"))
	return rv
}


// SetHeaderReferenceSize sets the value of the headerReferenceSize property.
// The default size to use for section headers.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/headerreferencesize
func (c_ CollectionViewFlowLayout) SetHeaderReferenceSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHeaderReferenceSize:"), value)
}

// The default size to use for items.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/itemsize
func (c_ CollectionViewFlowLayout) ItemSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("itemSize"))
	return rv
}


// SetItemSize sets the value of the itemSize property.
// The default size to use for items.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/itemsize
func (c_ CollectionViewFlowLayout) SetItemSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setItemSize:"), value)
}

// The minimum spacing (in points) to use between items in the same row or column.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/minimuminteritemspacing
func (c_ CollectionViewFlowLayout) MinimumInteritemSpacing() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("minimumInteritemSpacing"))
	return rv
}


// SetMinimumInteritemSpacing sets the value of the minimumInteritemSpacing property.
// The minimum spacing (in points) to use between items in the same row or column.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/minimuminteritemspacing
func (c_ CollectionViewFlowLayout) SetMinimumInteritemSpacing(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumInteritemSpacing:"), value)
}

// The minimum spacing (in points) to use between rows or columns.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/minimumlinespacing
func (c_ CollectionViewFlowLayout) MinimumLineSpacing() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("minimumLineSpacing"))
	return rv
}


// SetMinimumLineSpacing sets the value of the minimumLineSpacing property.
// The minimum spacing (in points) to use between rows or columns.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/minimumlinespacing
func (c_ CollectionViewFlowLayout) SetMinimumLineSpacing(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumLineSpacing:"), value)
}

// The scroll direction of the layout.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/scrolldirection
func (c_ CollectionViewFlowLayout) ScrollDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("scrollDirection"))
	return rv
}


// SetScrollDirection sets the value of the scrollDirection property.
// The scroll direction of the layout.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/scrolldirection
func (c_ CollectionViewFlowLayout) SetScrollDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScrollDirection:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/sectionfooterspintovisiblebounds
func (c_ CollectionViewFlowLayout) SectionFootersPinToVisibleBounds() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sectionFootersPinToVisibleBounds"))
	return rv
}


// SetSectionFootersPinToVisibleBounds sets the value of the sectionFootersPinToVisibleBounds property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/sectionfooterspintovisiblebounds
func (c_ CollectionViewFlowLayout) SetSectionFootersPinToVisibleBounds(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSectionFootersPinToVisibleBounds:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/sectionheaderspintovisiblebounds
func (c_ CollectionViewFlowLayout) SectionHeadersPinToVisibleBounds() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sectionHeadersPinToVisibleBounds"))
	return rv
}


// SetSectionHeadersPinToVisibleBounds sets the value of the sectionHeadersPinToVisibleBounds property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/sectionheaderspintovisiblebounds
func (c_ CollectionViewFlowLayout) SetSectionHeadersPinToVisibleBounds(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSectionHeadersPinToVisibleBounds:"), value)
}

// The margins used to lay out content in a section.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/sectioninset
func (c_ CollectionViewFlowLayout) SectionInset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sectionInset"))
	return rv
}


// SetSectionInset sets the value of the sectionInset property.
// The margins used to lay out content in a section.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewflowlayout/sectioninset
func (c_ CollectionViewFlowLayout) SetSectionInset(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSectionInset:"), value)
}

// An `NSSize` structure set to `0` in both dimensions.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSZeroSize
func (c_ CollectionViewFlowLayout) NSZeroSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("NSZeroSize"))
	return rv
}




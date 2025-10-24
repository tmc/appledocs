// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NSCollectionViewFlowLayout */


/* debug [class_header]: Header for NSCollectionViewFlowLayout */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionViewFlowLayout */
// An interface definition for the [CollectionViewFlowLayout] class.
type ICollectionViewFlowLayout interface {
	ICollectionViewLayout
	
/* debug [class_interface_properties]: Properties for CollectionViewFlowLayout */
	// properties:
	EstimatedItemSize() Size /* not a class type */
	SetEstimatedItemSize(value Size /* not a class type */)
	FooterReferenceSize() Size /* not a class type */
	SetFooterReferenceSize(value Size /* not a class type */)
	HeaderReferenceSize() Size /* not a class type */
	SetHeaderReferenceSize(value Size /* not a class type */)
	ItemSize() Size /* not a class type */
	SetItemSize(value Size /* not a class type */)
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
	Delegate() objc.IObject /* cross-framework: CollectionViewDelegate */
	SetDelegate(value objc.IObject /* cross-framework: CollectionViewDelegate */)
	NSZeroSize() Size /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionViewFlowLayout */
	// methods:
	CollapseSectionAtIndex(sectionIndex uint)
	ExpandSectionAtIndex(sectionIndex uint)
	SectionAtIndexIsCollapsed(sectionIndex uint) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionViewFlowLayout */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionViewFlowLayout */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionViewFlowLayout *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionViewFlowLayout */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionViewFlowLayout */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionViewFlowLayout */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/collapseSection(at:)
func (c_ CollectionViewFlowLayout) CollapseSectionAtIndex(sectionIndex uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("collapseSectionAtIndex:"), sectionIndex)
}/* debug [instance_methods/method]: CollapseSectionAtIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/expandSection(at:)
func (c_ CollectionViewFlowLayout) ExpandSectionAtIndex(sectionIndex uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("expandSectionAtIndex:"), sectionIndex)
}/* debug [instance_methods/method]: ExpandSectionAtIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/section(atIndexIsCollapsed:)
func (c_ CollectionViewFlowLayout) SectionAtIndexIsCollapsed(sectionIndex uint) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sectionAtIndexIsCollapsed:"), sectionIndex)
	return rv
}/* debug [instance_methods/method]: SectionAtIndexIsCollapsed */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionViewFlowLayout */

// The estimated size of items in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/estimatedItemSize
func (c_ CollectionViewFlowLayout) EstimatedItemSize() Size /* not a class type */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("estimatedItemSize"))
	return rv
}/* debug [instance_properties/getter]: estimatedItemSize */


// The estimated size of items in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/estimatedItemSize
func (c_ CollectionViewFlowLayout) SetEstimatedItemSize(value Size /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEstimatedItemSize:"), value)
}/* debug [instance_properties/setter]: estimatedItemSize */


// The default size to use for section footers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/footerReferenceSize
func (c_ CollectionViewFlowLayout) FooterReferenceSize() Size /* not a class type */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("footerReferenceSize"))
	return rv
}/* debug [instance_properties/getter]: footerReferenceSize */


// The default size to use for section footers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/footerReferenceSize
func (c_ CollectionViewFlowLayout) SetFooterReferenceSize(value Size /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFooterReferenceSize:"), value)
}/* debug [instance_properties/setter]: footerReferenceSize */


// The default size to use for section headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/headerReferenceSize
func (c_ CollectionViewFlowLayout) HeaderReferenceSize() Size /* not a class type */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("headerReferenceSize"))
	return rv
}/* debug [instance_properties/getter]: headerReferenceSize */


// The default size to use for section headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/headerReferenceSize
func (c_ CollectionViewFlowLayout) SetHeaderReferenceSize(value Size /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHeaderReferenceSize:"), value)
}/* debug [instance_properties/setter]: headerReferenceSize */


// The default size to use for items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/itemSize
func (c_ CollectionViewFlowLayout) ItemSize() Size /* not a class type */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("itemSize"))
	return rv
}/* debug [instance_properties/getter]: itemSize */


// The default size to use for items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/itemSize
func (c_ CollectionViewFlowLayout) SetItemSize(value Size /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setItemSize:"), value)
}/* debug [instance_properties/setter]: itemSize */


// The minimum spacing (in points) to use between items in the same row or column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/minimumInteritemSpacing
func (c_ CollectionViewFlowLayout) MinimumInteritemSpacing() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("minimumInteritemSpacing"))
	return rv
}/* debug [instance_properties/getter]: minimumInteritemSpacing */


// The minimum spacing (in points) to use between items in the same row or column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/minimumInteritemSpacing
func (c_ CollectionViewFlowLayout) SetMinimumInteritemSpacing(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumInteritemSpacing:"), value)
}/* debug [instance_properties/setter]: minimumInteritemSpacing */


// The minimum spacing (in points) to use between rows or columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/minimumLineSpacing
func (c_ CollectionViewFlowLayout) MinimumLineSpacing() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("minimumLineSpacing"))
	return rv
}/* debug [instance_properties/getter]: minimumLineSpacing */


// The minimum spacing (in points) to use between rows or columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/minimumLineSpacing
func (c_ CollectionViewFlowLayout) SetMinimumLineSpacing(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumLineSpacing:"), value)
}/* debug [instance_properties/setter]: minimumLineSpacing */


// The scroll direction of the layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/scrollDirection
func (c_ CollectionViewFlowLayout) ScrollDirection() CollectionViewScrollDirection {
	rv := objc.Send[CollectionViewScrollDirection](c_.ID, objc.Sel("scrollDirection"))
	return rv
}/* debug [instance_properties/getter]: scrollDirection */


// The scroll direction of the layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/scrollDirection
func (c_ CollectionViewFlowLayout) SetScrollDirection(value CollectionViewScrollDirection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScrollDirection:"), value)
}/* debug [instance_properties/setter]: scrollDirection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/sectionFootersPinToVisibleBounds
func (c_ CollectionViewFlowLayout) SectionFootersPinToVisibleBounds() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sectionFootersPinToVisibleBounds"))
	return rv
}/* debug [instance_properties/getter]: sectionFootersPinToVisibleBounds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/sectionFootersPinToVisibleBounds
func (c_ CollectionViewFlowLayout) SetSectionFootersPinToVisibleBounds(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSectionFootersPinToVisibleBounds:"), value)
}/* debug [instance_properties/setter]: sectionFootersPinToVisibleBounds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/sectionHeadersPinToVisibleBounds
func (c_ CollectionViewFlowLayout) SectionHeadersPinToVisibleBounds() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sectionHeadersPinToVisibleBounds"))
	return rv
}/* debug [instance_properties/getter]: sectionHeadersPinToVisibleBounds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/sectionHeadersPinToVisibleBounds
func (c_ CollectionViewFlowLayout) SetSectionHeadersPinToVisibleBounds(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSectionHeadersPinToVisibleBounds:"), value)
}/* debug [instance_properties/setter]: sectionHeadersPinToVisibleBounds */


// The margins used to lay out content in a section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/sectionInset
func (c_ CollectionViewFlowLayout) SectionInset() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](c_.ID, objc.Sel("sectionInset"))
	return rv
}/* debug [instance_properties/getter]: sectionInset */


// The margins used to lay out content in a section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/sectionInset
func (c_ CollectionViewFlowLayout) SetSectionInset(value foundation.EdgeInsets) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSectionInset:"), value)
}/* debug [instance_properties/setter]: sectionInset */


// The collection view’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/delegate
func (c_ CollectionViewFlowLayout) Delegate() objc.IObject /* cross-framework: CollectionViewDelegate */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The collection view’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/delegate
func (c_ CollectionViewFlowLayout) SetDelegate(value objc.IObject /* cross-framework: CollectionViewDelegate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// An `NSSize` structure set to `0` in both dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSZeroSize
func (c_ CollectionViewFlowLayout) NSZeroSize() Size /* not a class type */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("NSZeroSize"))
	return rv
}/* debug [instance_properties/getter]: NSZeroSize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionViewFlowLayout */




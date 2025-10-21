// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
}

// An abstract base class that you subclass and use to generate layout information for a collection view.
//
// The job of a layout object is to perform the calculations needed to determine the placement and appearance of items, supplementary views, and other content in the collection view. The layout object does not apply the layout attributes it generates to the views in your interface. Instead, it passes those layout attributes to the collection view, which then creates the needed views and applies the layout attributes to them. You do not create instances of this class directly. Instead, you create instances of one of its subclasses and associate that object with your collection view either programmatically (using the property) or at design time in Interface Builder. Changing the layout object of a collection view forces an immediate update of the layout information. Collection views support many different types of elements, most of which are visual and all of which require layout attributes: are the main elements managed by the layout. Each item represents a single piece of data in the collection view. A collection view can have a single group of items or it can divide the items into multiple sections. are optional views associated with a specific section. The layout object defines the placement and use of supplementary views. For example, grid and flow layouts use supplementary views to implement headers and footers for each section. Supplementary views cannot be selected by the user. are visual adornments used to implement themes or to present visual content that is unrelated to the data being managed by the collection view. Decoration views are optional and the layout object defines their use and placement. supply a drop target for dragged content. Gaps do not have a direct visual representation, but they do have layout attributes, which the collection view uses for hit testing. The layout object provides attributes for inter-item gaps only when asked to do so. Each concrete layout object defines a specific organization for the contained elements and provides the appropriate layout attributes. The placement and appearance of items is determined entirely by the layout object. The and subclasses define variants of a grid-based layout, but you can create custom layouts that arrange elements in different ways. For example, you might define a layout class that arranges items in a circle or define a class that groups items into stacks that resemble a pile of photos on a table.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CollectionViewLayoutClass) Alloc() CollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The layout object used to organize the collection view’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/collectionviewlayout
func (c_ CollectionViewLayout) CollectionViewLayout() NSCollectionViewLayout {
	rv := objc.Send[NSCollectionViewLayout](c_.ID, objc.Sel("collectionViewLayout"))
	return rv
}


// SetCollectionViewLayout sets the value of the collectionViewLayout property.
// The layout object used to organize the collection view’s content.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/collectionviewlayout
func (c_ CollectionViewLayout) SetCollectionViewLayout(value ICollectionViewLayout) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCollectionViewLayout:"), value)
}

// The collection view object currently using this layout.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/collectionview
func (c_ CollectionViewLayout) CollectionView() NSCollectionView {
	rv := objc.Send[NSCollectionView](c_.ID, objc.Sel("collectionView"))
	return rv
}


// SetCollectionView sets the value of the collectionView property.
// The collection view object currently using this layout.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/collectionview
func (c_ CollectionViewLayout) SetCollectionView(value ICollectionView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCollectionView:"), value)
}

// The width and height of the collection view’s contents.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/collectionviewcontentsize
func (c_ CollectionViewLayout) CollectionViewContentSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("collectionViewContentSize"))
	return rv
}


// SetCollectionViewContentSize sets the value of the collectionViewContentSize property.
// The width and height of the collection view’s contents.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/collectionviewcontentsize
func (c_ CollectionViewLayout) SetCollectionViewContentSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCollectionViewContentSize:"), value)
}




// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCollectionLayoutSize */


/* debug [class_header]: Header for NSCollectionLayoutSize */
// The class instance for the [CollectionLayoutSize] class.
var (
	CollectionLayoutSizeClass     _CollectionLayoutSizeClass
	CollectionLayoutSizeClassOnce sync.Once
)

func getCollectionLayoutSizeClass() _CollectionLayoutSizeClass {
	CollectionLayoutSizeClassOnce.Do(func() {
		CollectionLayoutSizeClass = _CollectionLayoutSizeClass{objc.GetClass("NSCollectionLayoutSize")}
	})
	return CollectionLayoutSizeClass
}

type _CollectionLayoutSizeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionLayoutSize */
// An interface definition for the [CollectionLayoutSize] class.
type ICollectionLayoutSize interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CollectionLayoutSize */
	// properties:
	HeightDimension() ICollectionLayoutDimension
	WidthDimension() ICollectionLayoutDimension
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionLayoutSize */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionLayoutSize */
// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutSizeClass) Alloc() CollectionLayoutSize {
	rv := objc.Send[CollectionLayoutSize](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionLayoutSizeClass) New() CollectionLayoutSize {
	rv := objc.Send[CollectionLayoutSize](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionLayoutSize) Init() CollectionLayoutSize {
	rv := objc.Send[CollectionLayoutSize](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionLayoutSize) Autorelease() CollectionLayoutSize {
	rv := objc.Send[CollectionLayoutSize](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionLayoutSize creates a new CollectionLayoutSize instance.
func NewCollectionLayoutSize() CollectionLayoutSize {
	return getCollectionLayoutSizeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionLayoutSize */
// The width and the height of an item in a collection view.
//
// A size is a pair of dimensions ( ): a width dimension and a height dimension. Every component of a collection view layout has an explicit size.


// The width and the height of an item in a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSize
type CollectionLayoutSize struct {
	objectivec.Object
}

// CollectionLayoutSizeFrom constructs a [CollectionLayoutSize] from an unsafe.Pointer.
//
// The width and the height of an item in a collection view.
func CollectionLayoutSizeFrom(ptr unsafe.Pointer) CollectionLayoutSize {
	return CollectionLayoutSize{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionLayoutSize */

// Creates a size with the specified width and height dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSize/init(widthDimension:heightDimension:)
func NewCollectionLayoutSizeWithWidthDimensionHeightDimension(width ICollectionLayoutDimension, height ICollectionLayoutDimension) CollectionLayoutSize {
	rv := objc.Send[CollectionLayoutSize](objc.ID(getCollectionLayoutSizeClass().class), objc.Sel("sizeWithWidthDimension:heightDimension:"), width, height)
	return rv
}/* debug [class_init_methods/constructor]: NewCollectionLayoutSizeWithWidthDimensionHeightDimension */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionLayoutSize */

// Creates a size with the specified width and height dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSize/init(widthDimension:heightDimension:)
func (cc _CollectionLayoutSizeClass) SizeWithWidthDimensionHeightDimension(width ICollectionLayoutDimension, height ICollectionLayoutDimension) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("sizeWithWidthDimension:heightDimension:"), width, height)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SizeWithWidthDimensionHeightDimension) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionLayoutSize */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionLayoutSize */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionLayoutSize */

// The height dimension of an item in a collection view layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSize/heightDimension
func (c_ CollectionLayoutSize) HeightDimension() ICollectionLayoutDimension {
	rv := objc.Send[CollectionLayoutDimension](c_.ID, objc.Sel("heightDimension"))
	return rv
}/* debug [instance_properties/getter]: heightDimension */


// The width dimension of an item in a collection view layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSize/widthDimension
func (c_ CollectionLayoutSize) WidthDimension() ICollectionLayoutDimension {
	rv := objc.Send[CollectionLayoutDimension](c_.ID, objc.Sel("widthDimension"))
	return rv
}/* debug [instance_properties/getter]: widthDimension */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionLayoutSize */



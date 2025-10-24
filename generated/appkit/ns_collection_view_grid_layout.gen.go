// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSCollectionViewGridLayout */


/* debug [class_header]: Header for NSCollectionViewGridLayout */
// The class instance for the [CollectionViewGridLayout] class.
var (
	CollectionViewGridLayoutClass     _CollectionViewGridLayoutClass
	CollectionViewGridLayoutClassOnce sync.Once
)

func getCollectionViewGridLayoutClass() _CollectionViewGridLayoutClass {
	CollectionViewGridLayoutClassOnce.Do(func() {
		CollectionViewGridLayoutClass = _CollectionViewGridLayoutClass{objc.GetClass("NSCollectionViewGridLayout")}
	})
	return CollectionViewGridLayoutClass
}

type _CollectionViewGridLayoutClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionViewGridLayout */
// An interface definition for the [CollectionViewGridLayout] class.
type ICollectionViewGridLayout interface {
	ICollectionViewLayout
	
/* debug [class_interface_properties]: Properties for CollectionViewGridLayout */
	// properties:
	BackgroundColors() []Color
	SetBackgroundColors(value []Color)
	Margins() foundation.EdgeInsets
	SetMargins(value foundation.EdgeInsets)
	MaximumItemSize() Size /* not a class type */
	SetMaximumItemSize(value Size /* not a class type */)
	MaximumNumberOfColumns() uint
	SetMaximumNumberOfColumns(value uint)
	MaximumNumberOfRows() uint
	SetMaximumNumberOfRows(value uint)
	MinimumInteritemSpacing() float64
	SetMinimumInteritemSpacing(value float64)
	MinimumItemSize() Size /* not a class type */
	SetMinimumItemSize(value Size /* not a class type */)
	MinimumLineSpacing() float64
	SetMinimumLineSpacing(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionViewGridLayout */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionViewGridLayout */
// Alloc allocates a new instance without initialization.
func (cc _CollectionViewGridLayoutClass) Alloc() CollectionViewGridLayout {
	rv := objc.Send[CollectionViewGridLayout](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionViewGridLayoutClass) New() CollectionViewGridLayout {
	rv := objc.Send[CollectionViewGridLayout](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionViewGridLayout) Init() CollectionViewGridLayout {
	rv := objc.Send[CollectionViewGridLayout](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionViewGridLayout) Autorelease() CollectionViewGridLayout {
	rv := objc.Send[CollectionViewGridLayout](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionViewGridLayout creates a new CollectionViewGridLayout instance.
func NewCollectionViewGridLayout() CollectionViewGridLayout {
	return getCollectionViewGridLayoutClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionViewGridLayout */
// A layout that displays a single section of items in a row and column grid.
//
// The object provides the same layout behavior offered by the class prior to macOS 10.11, and you can use it in cases where you want to maintain the old appearance while still taking advantage of newer collection view features.


// A layout that displays a single section of items in a row and column grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout
type CollectionViewGridLayout struct {
	CollectionViewLayout
}

// CollectionViewGridLayoutFrom constructs a [CollectionViewGridLayout] from an unsafe.Pointer.
//
// A layout that displays a single section of items in a row and column grid.
func CollectionViewGridLayoutFrom(ptr unsafe.Pointer) CollectionViewGridLayout {
	return CollectionViewGridLayout{
		CollectionViewLayout: CollectionViewLayoutFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionViewGridLayout *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionViewGridLayout */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionViewGridLayout */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionViewGridLayout */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionViewGridLayout */

// The array of background colors to use when drawing the grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/backgroundColors
func (c_ CollectionViewGridLayout) BackgroundColors() []Color {
	rv := objc.Send[[]Color](c_.ID, objc.Sel("backgroundColors"))
	return rv
}/* debug [instance_properties/getter]: backgroundColors */


// The array of background colors to use when drawing the grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/backgroundColors
func (c_ CollectionViewGridLayout) SetBackgroundColors(value []Color) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setBackgroundColors:"), nsArray)
}/* debug [instance_properties/setter]: backgroundColors */


// The amount of empty space (in points) around the grid’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/margins
func (c_ CollectionViewGridLayout) Margins() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](c_.ID, objc.Sel("margins"))
	return rv
}/* debug [instance_properties/getter]: margins */


// The amount of empty space (in points) around the grid’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/margins
func (c_ CollectionViewGridLayout) SetMargins(value foundation.EdgeInsets) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMargins:"), value)
}/* debug [instance_properties/setter]: margins */


// The largest allowable size for an item’s view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/maximumItemSize
func (c_ CollectionViewGridLayout) MaximumItemSize() Size /* not a class type */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("maximumItemSize"))
	return rv
}/* debug [instance_properties/getter]: maximumItemSize */


// The largest allowable size for an item’s view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/maximumItemSize
func (c_ CollectionViewGridLayout) SetMaximumItemSize(value Size /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaximumItemSize:"), value)
}/* debug [instance_properties/setter]: maximumItemSize */


// The maximum number of columns to display in the collection view’s visible area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/maximumNumberOfColumns
func (c_ CollectionViewGridLayout) MaximumNumberOfColumns() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maximumNumberOfColumns"))
	return rv
}/* debug [instance_properties/getter]: maximumNumberOfColumns */


// The maximum number of columns to display in the collection view’s visible area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/maximumNumberOfColumns
func (c_ CollectionViewGridLayout) SetMaximumNumberOfColumns(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaximumNumberOfColumns:"), value)
}/* debug [instance_properties/setter]: maximumNumberOfColumns */


// The maximum number of rows to display in the collection view’s visible area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/maximumNumberOfRows
func (c_ CollectionViewGridLayout) MaximumNumberOfRows() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maximumNumberOfRows"))
	return rv
}/* debug [instance_properties/getter]: maximumNumberOfRows */


// The maximum number of rows to display in the collection view’s visible area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/maximumNumberOfRows
func (c_ CollectionViewGridLayout) SetMaximumNumberOfRows(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaximumNumberOfRows:"), value)
}/* debug [instance_properties/setter]: maximumNumberOfRows */


// The minimum spacing (in points) to use between items in the same row or column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/minimumInteritemSpacing
func (c_ CollectionViewGridLayout) MinimumInteritemSpacing() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("minimumInteritemSpacing"))
	return rv
}/* debug [instance_properties/getter]: minimumInteritemSpacing */


// The minimum spacing (in points) to use between items in the same row or column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/minimumInteritemSpacing
func (c_ CollectionViewGridLayout) SetMinimumInteritemSpacing(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumInteritemSpacing:"), value)
}/* debug [instance_properties/setter]: minimumInteritemSpacing */


// The smallest allowable size for an item’s view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/minimumItemSize
func (c_ CollectionViewGridLayout) MinimumItemSize() Size /* not a class type */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("minimumItemSize"))
	return rv
}/* debug [instance_properties/getter]: minimumItemSize */


// The smallest allowable size for an item’s view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/minimumItemSize
func (c_ CollectionViewGridLayout) SetMinimumItemSize(value Size /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumItemSize:"), value)
}/* debug [instance_properties/setter]: minimumItemSize */


// The minimum spacing (in points) to use between rows or columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/minimumLineSpacing
func (c_ CollectionViewGridLayout) MinimumLineSpacing() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("minimumLineSpacing"))
	return rv
}/* debug [instance_properties/getter]: minimumLineSpacing */


// The minimum spacing (in points) to use between rows or columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/minimumLineSpacing
func (c_ CollectionViewGridLayout) SetMinimumLineSpacing(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumLineSpacing:"), value)
}/* debug [instance_properties/setter]: minimumLineSpacing */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionViewGridLayout */




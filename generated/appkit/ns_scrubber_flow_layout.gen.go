// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NSScrubberFlowLayout */


/* debug [class_header]: Header for NSScrubberFlowLayout */
// The class instance for the [ScrubberFlowLayout] class.
var (
	ScrubberFlowLayoutClass     _ScrubberFlowLayoutClass
	ScrubberFlowLayoutClassOnce sync.Once
)

func getScrubberFlowLayoutClass() _ScrubberFlowLayoutClass {
	ScrubberFlowLayoutClassOnce.Do(func() {
		ScrubberFlowLayoutClass = _ScrubberFlowLayoutClass{objc.GetClass("NSScrubberFlowLayout")}
	})
	return ScrubberFlowLayoutClass
}

type _ScrubberFlowLayoutClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScrubberFlowLayout */
// An interface definition for the [ScrubberFlowLayout] class.
type IScrubberFlowLayout interface {
	IScrubberLayout
	
/* debug [class_interface_properties]: Properties for ScrubberFlowLayout */
	// properties:
	ItemSize() Size /* not a class type */
	SetItemSize(value Size /* not a class type */)
	ItemSpacing() float64
	SetItemSpacing(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScrubberFlowLayout */
	// methods:
	InvalidateLayoutForItemsAtIndexes(invalidItemIndexes foundation.IndexSet)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScrubberFlowLayout */
// Alloc allocates a new instance without initialization.
func (sc _ScrubberFlowLayoutClass) Alloc() ScrubberFlowLayout {
	rv := objc.Send[ScrubberFlowLayout](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrubberFlowLayoutClass) New() ScrubberFlowLayout {
	rv := objc.Send[ScrubberFlowLayout](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberFlowLayout) Init() ScrubberFlowLayout {
	rv := objc.Send[ScrubberFlowLayout](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberFlowLayout) Autorelease() ScrubberFlowLayout {
	rv := objc.Send[ScrubberFlowLayout](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberFlowLayout creates a new ScrubberFlowLayout instance.
func NewScrubberFlowLayout() ScrubberFlowLayout {
	return getScrubberFlowLayoutClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScrubberFlowLayout */
// A concrete layout object that arranges items end-to-end in a linear strip.
//
// To set the size of items on a per-item basis, ensure that your scrubber delegate conforms to the protocol, and provides an implementation of the method.


// A concrete layout object that arranges items end-to-end in a linear strip.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout
type ScrubberFlowLayout struct {
	ScrubberLayout
}

// ScrubberFlowLayoutFrom constructs a [ScrubberFlowLayout] from an unsafe.Pointer.
//
// A concrete layout object that arranges items end-to-end in a linear strip.
func ScrubberFlowLayoutFrom(ptr unsafe.Pointer) ScrubberFlowLayout {
	return ScrubberFlowLayout{
		ScrubberLayout: ScrubberLayoutFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScrubberFlowLayout *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScrubberFlowLayout */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScrubberFlowLayout */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScrubberFlowLayout */

// Informs the scrubber that it should perform a new layout pass for the items at the specified indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout/invalidateLayoutForItems(at:)
func (s_ ScrubberFlowLayout) InvalidateLayoutForItemsAtIndexes(invalidItemIndexes foundation.IndexSet) {
	objc.Send[objc.ID](s_.ID, objc.Sel("invalidateLayoutForItemsAtIndexes:"), invalidItemIndexes)
}/* debug [instance_methods/method]: InvalidateLayoutForItemsAtIndexes */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScrubberFlowLayout */

// The frame size for each item in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout/itemSize
func (s_ ScrubberFlowLayout) ItemSize() Size /* not a class type */ {
	rv := objc.Send[Size](s_.ID, objc.Sel("itemSize"))
	return rv
}/* debug [instance_properties/getter]: itemSize */


// The frame size for each item in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout/itemSize
func (s_ ScrubberFlowLayout) SetItemSize(value Size /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setItemSize:"), value)
}/* debug [instance_properties/setter]: itemSize */


// The horizontal spacing between items, specified in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout/itemSpacing
func (s_ ScrubberFlowLayout) ItemSpacing() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("itemSpacing"))
	return rv
}/* debug [instance_properties/getter]: itemSpacing */


// The horizontal spacing between items, specified in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout/itemSpacing
func (s_ ScrubberFlowLayout) SetItemSpacing(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setItemSpacing:"), value)
}/* debug [instance_properties/setter]: itemSpacing */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScrubberFlowLayout */




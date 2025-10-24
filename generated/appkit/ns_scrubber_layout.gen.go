// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSScrubberLayout */


/* debug [class_header]: Header for NSScrubberLayout */
// The class instance for the [ScrubberLayout] class.
var (
	ScrubberLayoutClass     _ScrubberLayoutClass
	ScrubberLayoutClassOnce sync.Once
)

func getScrubberLayoutClass() _ScrubberLayoutClass {
	ScrubberLayoutClassOnce.Do(func() {
		ScrubberLayoutClass = _ScrubberLayoutClass{objc.GetClass("NSScrubberLayout")}
	})
	return ScrubberLayoutClass
}

type _ScrubberLayoutClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScrubberLayout */
// An interface definition for the [ScrubberLayout] class.
type IScrubberLayout interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ScrubberLayout */
	// properties:
	AutomaticallyMirrorsInRightToLeftLayout() bool
	Scrubber() IScrubber
	ScrubberContentSize() Size /* not a class type */
	ShouldInvalidateLayoutForHighlightChange() bool
	ShouldInvalidateLayoutForSelectionChange() bool
	VisibleRect() Rect /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScrubberLayout */
	// methods:
	InvalidateLayout()
	LayoutAttributesForItemAtIndex(index int) IScrubberLayoutAttributes
	LayoutAttributesForItemsInRect(rect Rect /* not a class type */) unsafe.Pointer
	PrepareLayout()
	ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect(fromVisibleRect Rect /* not a class type */, toVisibleRect Rect /* not a class type */) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScrubberLayout */
// Alloc allocates a new instance without initialization.
func (sc _ScrubberLayoutClass) Alloc() ScrubberLayout {
	rv := objc.Send[ScrubberLayout](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrubberLayoutClass) New() ScrubberLayout {
	rv := objc.Send[ScrubberLayout](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberLayout) Init() ScrubberLayout {
	rv := objc.Send[ScrubberLayout](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberLayout) Autorelease() ScrubberLayout {
	rv := objc.Send[ScrubberLayout](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberLayout creates a new ScrubberLayout instance.
func NewScrubberLayout() ScrubberLayout {
	return getScrubberLayoutClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScrubberLayout */
// An abstract class that describes the layout of items within a scrubber control.
//
// To determine the layout of items in a scrubber, use one of the built-in subclasses ( or ), or create a custom subclass to implement your own layout.


// An abstract class that describes the layout of items within a scrubber control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout
type ScrubberLayout struct {
	objectivec.Object
}

// ScrubberLayoutFrom constructs a [ScrubberLayout] from an unsafe.Pointer.
//
// An abstract class that describes the layout of items within a scrubber control.
func ScrubberLayoutFrom(ptr unsafe.Pointer) ScrubberLayout {
	return ScrubberLayout{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScrubberLayout */

// Initializes and returns a newly allocated scrubber layout object from a storyboard or nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/init(coder:)
func NewScrubberLayoutWithCoder(coder foundation.Coder) ScrubberLayout {
	instance := getScrubberLayoutClass().Alloc()
	rv := objc.Send[ScrubberLayout](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewScrubberLayoutWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScrubberLayout */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScrubberLayout */

// A property containing a class that describes layout attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/layoutAttributesClass
func (sc _ScrubberLayoutClass) LayoutAttributesClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(sc.class), objc.Sel("layoutAttributesClass"))
	return rv
}/* debug [class_properties_class/property]: layoutAttributesClass */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScrubberLayout */

// Signals that the layout has been invalidated, and that the scrubber control should perform a new layout pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/invalidateLayout()
func (s_ ScrubberLayout) InvalidateLayout() {
	objc.Send[objc.ID](s_.ID, objc.Sel("invalidateLayout"))
}/* debug [instance_methods/method]: InvalidateLayout */


// The layout attributes for the item with the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/layoutAttributesForItem(at:)
func (s_ ScrubberLayout) LayoutAttributesForItemAtIndex(index int) IScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](s_.ID, objc.Sel("layoutAttributesForItemAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: LayoutAttributesForItemAtIndex */


// The set of layout attributes for all items within the provided rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/layoutAttributesForItems(in:)
func (s_ ScrubberLayout) LayoutAttributesForItemsInRect(rect Rect /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("layoutAttributesForItemsInRect:"), rect)
	return rv
}/* debug [instance_methods/method]: LayoutAttributesForItemsInRect */


// Gives you an opportunity to perform layout calculations when the scrubber’s layout is invalidated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/prepare()
func (s_ ScrubberLayout) PrepareLayout() {
	objc.Send[objc.ID](s_.ID, objc.Sel("prepareLayout"))
}/* debug [instance_methods/method]: PrepareLayout */


// Determines whether the scrubber should refresh its layout in response to a change of its visible region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/shouldInvalidateLayoutForChange(fromVisibleRect:toVisibleRect:)
func (s_ ScrubberLayout) ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect(fromVisibleRect Rect /* not a class type */, toVisibleRect Rect /* not a class type */) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldInvalidateLayoutForChangeFromVisibleRect:toVisibleRect:"), fromVisibleRect, toVisibleRect)
	return rv
}/* debug [instance_methods/method]: ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScrubberLayout */

// Determines whether the scrubber mirrors its layout for right-to-left layouts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/automaticallyMirrorsInRightToLeftLayout
func (s_ ScrubberLayout) AutomaticallyMirrorsInRightToLeftLayout() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticallyMirrorsInRightToLeftLayout"))
	return rv
}/* debug [instance_properties/getter]: automaticallyMirrorsInRightToLeftLayout */


// A property containing a class that describes layout attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/layoutAttributesClass
func (s_ ScrubberLayout) LayoutAttributesClass() objc.Class {
	rv := objc.Send[objc.Class](s_.ID, objc.Sel("layoutAttributesClass"))
	return rv
}/* debug [instance_properties/getter]: layoutAttributesClass */


// The scrubber control that this layout is assigned to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/scrubber
func (s_ ScrubberLayout) Scrubber() IScrubber {
	rv := objc.Send[Scrubber](s_.ID, objc.Sel("scrubber"))
	return rv
}/* debug [instance_properties/getter]: scrubber */


// The size required to contain all elements within the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/scrubberContentSize
func (s_ ScrubberLayout) ScrubberContentSize() Size /* not a class type */ {
	rv := objc.Send[Size](s_.ID, objc.Sel("scrubberContentSize"))
	return rv
}/* debug [instance_properties/getter]: scrubberContentSize */


// Determines whether the scrubber should refresh its layout when an item is highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/shouldInvalidateLayoutForHighlightChange
func (s_ ScrubberLayout) ShouldInvalidateLayoutForHighlightChange() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldInvalidateLayoutForHighlightChange"))
	return rv
}/* debug [instance_properties/getter]: shouldInvalidateLayoutForHighlightChange */


// Determines whether the scrubber should refresh its layout when the selection changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/shouldInvalidateLayoutForSelectionChange
func (s_ ScrubberLayout) ShouldInvalidateLayoutForSelectionChange() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldInvalidateLayoutForSelectionChange"))
	return rv
}/* debug [instance_properties/getter]: shouldInvalidateLayoutForSelectionChange */


// The currently visible rectangle, in the coordinate space of the scrubber content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/visibleRect
func (s_ ScrubberLayout) VisibleRect() Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("visibleRect"))
	return rv
}/* debug [instance_properties/getter]: visibleRect */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScrubberLayout */



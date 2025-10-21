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

// An interface definition for the [ScrubberLayout] class.
type IScrubberLayout interface {
	objectivec.IObject
	InvalidateLayout()
	LayoutAttributesForItemAtIndex(index int) ScrubberLayoutAttributes
	LayoutAttributesForItemsInRect(rect coregraphics.CGRect) unsafe.Pointer
	PrepareLayout()
	ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect(fromVisibleRect coregraphics.CGRect, toVisibleRect coregraphics.CGRect) bool
}

// An abstract class that describes the layout of items within a scrubber control.
//
// To determine the layout of items in a scrubber, use one of the built-in subclasses ( or ), or create a custom subclass to implement your own layout.
//
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

// Alloc allocates a new instance without initialization.
func (sc _ScrubberLayoutClass) Alloc() ScrubberLayout {
	rv := objc.Send[ScrubberLayout](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes and returns a newly allocated scrubber layout object from a storyboard or nib file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/init(coder:)
func NewScrubberLayoutWithCoder(coder foundation.ICoder) ScrubberLayout {
	instance := getScrubberLayoutClass().Alloc()
	rv := objc.Send[ScrubberLayout](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// A property containing a class that describes layout attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/layoutAttributesClass
func (sc _ScrubberLayoutClass) LayoutAttributesClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(sc.class), objc.Sel("layoutAttributesClass"))
	return rv
}
// Signals that the layout has been invalidated, and that the scrubber control should perform a new layout pass.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/invalidateLayout()
func (s_ ScrubberLayout) InvalidateLayout() {
	objc.Send[objc.ID](s_.ID, objc.Sel("invalidateLayout"))
}

// The layout attributes for the item with the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/layoutAttributesForItem(at:)
func (s_ ScrubberLayout) LayoutAttributesForItemAtIndex(index int) ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](s_.ID, objc.Sel("layoutAttributesForItemAtIndex:"), index)
	return rv
}

// The set of layout attributes for all items within the provided rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/layoutAttributesForItems(in:)
func (s_ ScrubberLayout) LayoutAttributesForItemsInRect(rect coregraphics.CGRect) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("layoutAttributesForItemsInRect:"), rect)
	return rv
}

// Gives you an opportunity to perform layout calculations when the scrubber’s layout is invalidated.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/prepare()
func (s_ ScrubberLayout) PrepareLayout() {
	objc.Send[objc.ID](s_.ID, objc.Sel("prepareLayout"))
}

// Determines whether the scrubber should refresh its layout in response to a change of its visible region.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/shouldInvalidateLayoutForChange(fromVisibleRect:toVisibleRect:)
func (s_ ScrubberLayout) ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect(fromVisibleRect coregraphics.CGRect, toVisibleRect coregraphics.CGRect) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldInvalidateLayoutForChangeFromVisibleRect:toVisibleRect:"), fromVisibleRect, toVisibleRect)
	return rv
}

// Determines whether the scrubber mirrors its layout for right-to-left layouts.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/automaticallyMirrorsInRightToLeftLayout
func (s_ ScrubberLayout) AutomaticallyMirrorsInRightToLeftLayout() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticallyMirrorsInRightToLeftLayout"))
	return rv
}

// A property containing a class that describes layout attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/layoutAttributesClass
func (s_ ScrubberLayout) LayoutAttributesClass() objc.Class {
	rv := objc.Send[objc.Class](s_.ID, objc.Sel("layoutAttributesClass"))
	return rv
}

// The scrubber control that this layout is assigned to.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/scrubber
func (s_ ScrubberLayout) Scrubber() NSScrubber {
	rv := objc.Send[NSScrubber](s_.ID, objc.Sel("scrubber"))
	return rv
}

// The size required to contain all elements within the scrubber.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/scrubberContentSize
func (s_ ScrubberLayout) ScrubberContentSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](s_.ID, objc.Sel("scrubberContentSize"))
	return rv
}

// Determines whether the scrubber should refresh its layout when an item is highlighted.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/shouldInvalidateLayoutForHighlightChange
func (s_ ScrubberLayout) ShouldInvalidateLayoutForHighlightChange() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldInvalidateLayoutForHighlightChange"))
	return rv
}

// Determines whether the scrubber should refresh its layout when the selection changes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/shouldInvalidateLayoutForSelectionChange
func (s_ ScrubberLayout) ShouldInvalidateLayoutForSelectionChange() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldInvalidateLayoutForSelectionChange"))
	return rv
}

// The currently visible rectangle, in the coordinate space of the scrubber content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/visibleRect
func (s_ ScrubberLayout) VisibleRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("visibleRect"))
	return rv
}



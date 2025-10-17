// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScrubberLayout] class.
var scrubberLayoutClass = _ScrubberLayoutClass{objc.GetClass("NSScrubberLayout")}

type _ScrubberLayoutClass struct {
	class objc.Class
}

// An abstract class that describes the layout of items within a scrubber control. [Full Topic]
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

// Signals that the layout has been invalidated, and that the scrubber control should perform a new layout pass. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/invalidateLayout()
func (s_ ScrubberLayout) InvalidateLayout() {
	objc.Send[objc.ID](s_.ID, objc.Sel("invalidateLayout"))
}
// The layout attributes for the item with the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/layoutAttributesForItem(at:)
func (s_ ScrubberLayout) LayoutAttributesForItemAtIndex(index int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("layoutAttributesForItemAtIndex:"), index)
	return rv
}
// The set of layout attributes for all items within the provided rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/layoutAttributesForItems(in:)
func (s_ ScrubberLayout) LayoutAttributesForItemsInRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("layoutAttributesForItemsInRect:"), rect)
	return rv
}
// Gives you an opportunity to perform layout calculations when the scrubber’s layout is invalidated. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/prepare()
func (s_ ScrubberLayout) PrepareLayout() {
	objc.Send[objc.ID](s_.ID, objc.Sel("prepareLayout"))
}
// Determines whether the scrubber should refresh its layout in response to a change of its visible region. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/shouldInvalidateLayoutForChange(fromVisibleRect:toVisibleRect:)
func (s_ ScrubberLayout) ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect(fromVisibleRect unsafe.Pointer, toVisibleRect unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldInvalidateLayoutForChangeFromVisibleRect:toVisibleRect:"), fromVisibleRect, toVisibleRect)
	return rv
}



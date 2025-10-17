// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberLayout] class.
var ScrubberLayoutClass objc.Class

func init() {
	ScrubberLayoutClass = objc.GetClass("NSScrubberLayout")
}

type ScrubberLayout struct {
	objc.ID
}

func ScrubberLayoutFrom(ptr unsafe.Pointer) ScrubberLayout {
	return ScrubberLayout{
		ID: objc.ID(ptr),
	}
}


// Signals that the layout has been invalidated, and that the scrubber control should perform a new layout pass. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayout/invalidateLayout()
func (s_ ScrubberLayout) InvalidateLayout() {
	sel := objc.RegisterName("invalidateLayout")
	s_.ID.Send(sel)
}
// The layout attributes for the item with the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayout/layoutAttributesForItem(at:)
func (s_ ScrubberLayout) LayoutAttributesForItemAtIndex(index int) unsafe.Pointer {
	sel := objc.RegisterName("layoutAttributesForItemAtIndex:")
	ret := s_.ID.Send(sel, index)
	return unsafe.Pointer(ret)
}
// The set of layout attributes for all items within the provided rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayout/layoutAttributesForItems(in:)
func (s_ ScrubberLayout) LayoutAttributesForItemsInRect(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("layoutAttributesForItemsInRect:")
	ret := s_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Gives you an opportunity to perform layout calculations when the scrubber’s layout is invalidated. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayout/prepare()
func (s_ ScrubberLayout) PrepareLayout() {
	sel := objc.RegisterName("prepareLayout")
	s_.ID.Send(sel)
}
// Determines whether the scrubber should refresh its layout in response to a change of its visible region. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayout/shouldInvalidateLayoutForChange(fromVisibleRect:toVisibleRect:)
func (s_ ScrubberLayout) ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect(fromVisibleRect unsafe.Pointer, toVisibleRect unsafe.Pointer) bool {
	sel := objc.RegisterName("shouldInvalidateLayoutForChangeFromVisibleRect:toVisibleRect:")
	ret := s_.ID.Send(sel, fromVisibleRect, toVisibleRect)
	return ret != 0
}



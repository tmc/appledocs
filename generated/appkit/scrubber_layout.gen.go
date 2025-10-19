// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScrubberLayout] class.
var (
	scrubberLayoutClass     _ScrubberLayoutClass
	scrubberLayoutClassOnce sync.Once
)

func getScrubberLayoutClass() _ScrubberLayoutClass {
	scrubberLayoutClassOnce.Do(func() {
		scrubberLayoutClass = _ScrubberLayoutClass{objc.GetClass("NSScrubberLayout")}
	})
	return scrubberLayoutClass
}

type _ScrubberLayoutClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberLayout] class.
type IScrubberLayout interface {
	objectivec.IObject
	InvalidateLayout()
	LayoutAttributesForItemAtIndex(index int) unsafe.Pointer
	LayoutAttributesForItemsInRect(rect unsafe.Pointer) unsafe.Pointer
	PrepareLayout()
	ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect(fromVisibleRect unsafe.Pointer, toVisibleRect unsafe.Pointer) bool
}

// An abstract class that describes the layout of items within a scrubber control.
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


// Signals that the layout has been invalidated, and that the scrubber control should perform a new layout pass.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/invalidateLayout()
func (s_ ScrubberLayout) InvalidateLayout() {
	objc.Send[objc.ID](s_.ID, objc.Sel("invalidateLayout"))
}

// The layout attributes for the item with the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/layoutAttributesForItem(at:)
func (s_ ScrubberLayout) LayoutAttributesForItemAtIndex(index int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("layoutAttributesForItemAtIndex:"), index)
	return rv
}

// The set of layout attributes for all items within the provided rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/layoutAttributesForItems(in:)
func (s_ ScrubberLayout) LayoutAttributesForItemsInRect(rect unsafe.Pointer) unsafe.Pointer {
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
func (s_ ScrubberLayout) ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect(fromVisibleRect unsafe.Pointer, toVisibleRect unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldInvalidateLayoutForChangeFromVisibleRect:toVisibleRect:"), fromVisibleRect, toVisibleRect)
	return rv
}




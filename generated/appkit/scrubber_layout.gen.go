
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// The class instance for the [ScrubberLayout] class.
var ScrubberLayoutClass _ScrubberLayoutClass

func init() {
	ScrubberLayoutClass = _ScrubberLayoutClass{objc.GetClass("NSScrubberLayout")}
}

type _ScrubberLayoutClass struct {
	objc.Class
}

// An interface definition for the [ScrubberLayout] class.
type IScrubberLayout interface {
	ID() objc.ID
	InvalidateLayout()
	LayoutAttributesForItemAtIndex(index int) unsafe.Pointer
	LayoutAttributesForItemsInRect(rect foundation.Rect) unsafe.Pointer
	PrepareLayout()
	ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect(fromVisibleRect foundation.Rect, toVisibleRect foundation.Rect) bool
}

type ScrubberLayout struct {
	id objc.ID
}

func ScrubberLayoutFrom(ptr unsafe.Pointer) ScrubberLayout {
	return ScrubberLayout{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ ScrubberLayout) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberLayoutClass) Alloc() ScrubberLayout {
	rv := objc.Send[ScrubberLayout](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _ScrubberLayoutClass) New() ScrubberLayout {
	rv := objc.Send[ScrubberLayout](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewScrubberLayout creates and returns a new initialized instance.
func NewScrubberLayout() ScrubberLayout {
	return ScrubberLayoutClass.New()
}

// Init initializes the instance.
func (s_ ScrubberLayout) Init() ScrubberLayout {
	rv := objc.Send[ScrubberLayout](s_.ID(), selInit)
	return rv
}
// Signals that the layout has been invalidated, and that the scrubber control should perform a new layout pass. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayout/invalidateLayout()
func (s_ ScrubberLayout) InvalidateLayout() {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("invalidateLayout"))
}
// The layout attributes for the item with the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayout/layoutAttributesForItem(at:)
func (s_ ScrubberLayout) LayoutAttributesForItemAtIndex(index int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("layoutAttributesForItemAtIndex:"), index)
	return rv
}
// The set of layout attributes for all items within the provided rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayout/layoutAttributesForItems(in:)
func (s_ ScrubberLayout) LayoutAttributesForItemsInRect(rect foundation.Rect) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("layoutAttributesForItemsInRect:"), rect)
	return rv
}
// Gives you an opportunity to perform layout calculations when the scrubber’s layout is invalidated. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayout/prepare()
func (s_ ScrubberLayout) PrepareLayout() {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("prepareLayout"))
}
// Determines whether the scrubber should refresh its layout in response to a change of its visible region. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayout/shouldInvalidateLayoutForChange(fromVisibleRect:toVisibleRect:)
func (s_ ScrubberLayout) ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect(fromVisibleRect foundation.Rect, toVisibleRect foundation.Rect) bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("shouldInvalidateLayoutForChangeFromVisibleRect:toVisibleRect:"), fromVisibleRect, toVisibleRect)
	return rv
}
// The size required to contain all elements within the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayout/scrubberContentSize
func (s_ ScrubberLayout) ScrubberContentSize() foundation.Size {
	rv := objc.Send[foundation.Size](s_.ID(), objc.RegisterName("scrubberContentSize"))
	return rv
}
// Determines whether the scrubber should refresh its layout when an item is highlighted. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayout/shouldInvalidateLayoutForHighlightChange
func (s_ ScrubberLayout) ShouldInvalidateLayoutForHighlightChange() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("shouldInvalidateLayoutForHighlightChange"))
	return rv
}
// Determines whether the scrubber should refresh its layout when the selection changes. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayout/shouldInvalidateLayoutForSelectionChange
func (s_ ScrubberLayout) ShouldInvalidateLayoutForSelectionChange() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("shouldInvalidateLayoutForSelectionChange"))
	return rv
}

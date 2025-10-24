// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	// properties:
	AutomaticallyMirrorsInRightToLeftLayout() bool
	SetAutomaticallyMirrorsInRightToLeftLayout(value bool)
	Scrubber() IScrubber
	SetScrubber(value IScrubber)
	ScrubberContentSize() objc.IObject /* cross-framework: Size */
	SetScrubberContentSize(value objc.IObject /* cross-framework: Size */)
	ShouldInvalidateLayoutForHighlightChange() bool
	SetShouldInvalidateLayoutForHighlightChange(value bool)
	ShouldInvalidateLayoutForSelectionChange() bool
	SetShouldInvalidateLayoutForSelectionChange(value bool)
	VisibleRect() objc.IObject /* cross-framework: Rect */
	SetVisibleRect(value objc.IObject /* cross-framework: Rect */)
	// methods:
}

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



// Determines whether the scrubber mirrors its layout for right-to-left layouts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/automaticallymirrorsinrighttoleftlayout
func (s_ ScrubberLayout) AutomaticallyMirrorsInRightToLeftLayout() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticallyMirrorsInRightToLeftLayout"))
	return rv
}


// Determines whether the scrubber mirrors its layout for right-to-left layouts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/automaticallymirrorsinrighttoleftlayout
func (s_ ScrubberLayout) SetAutomaticallyMirrorsInRightToLeftLayout(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutomaticallyMirrorsInRightToLeftLayout:"), value)
}


// The scrubber control that this layout is assigned to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/scrubber
func (s_ ScrubberLayout) Scrubber() IScrubber {
	rv := objc.Send[Scrubber](s_.ID, objc.Sel("scrubber"))
	return rv
}


// The scrubber control that this layout is assigned to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/scrubber
func (s_ ScrubberLayout) SetScrubber(value IScrubber) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrubber:"), value)
}


// The size required to contain all elements within the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/scrubbercontentsize
func (s_ ScrubberLayout) ScrubberContentSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](s_.ID, objc.Sel("scrubberContentSize"))
	return rv
}


// The size required to contain all elements within the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/scrubbercontentsize
func (s_ ScrubberLayout) SetScrubberContentSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrubberContentSize:"), value)
}


// Determines whether the scrubber should refresh its layout when an item is highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/shouldinvalidatelayoutforhighlightchange
func (s_ ScrubberLayout) ShouldInvalidateLayoutForHighlightChange() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldInvalidateLayoutForHighlightChange"))
	return rv
}


// Determines whether the scrubber should refresh its layout when an item is highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/shouldinvalidatelayoutforhighlightchange
func (s_ ScrubberLayout) SetShouldInvalidateLayoutForHighlightChange(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShouldInvalidateLayoutForHighlightChange:"), value)
}


// Determines whether the scrubber should refresh its layout when the selection changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/shouldinvalidatelayoutforselectionchange
func (s_ ScrubberLayout) ShouldInvalidateLayoutForSelectionChange() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldInvalidateLayoutForSelectionChange"))
	return rv
}


// Determines whether the scrubber should refresh its layout when the selection changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/shouldinvalidatelayoutforselectionchange
func (s_ ScrubberLayout) SetShouldInvalidateLayoutForSelectionChange(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShouldInvalidateLayoutForSelectionChange:"), value)
}


// The currently visible rectangle, in the coordinate space of the scrubber content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/visiblerect
func (s_ ScrubberLayout) VisibleRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](s_.ID, objc.Sel("visibleRect"))
	return rv
}


// The currently visible rectangle, in the coordinate space of the scrubber content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/visiblerect
func (s_ ScrubberLayout) SetVisibleRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVisibleRect:"), value)
}




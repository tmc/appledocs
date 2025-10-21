// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScrubberArrangedView] class.
var (
	ScrubberArrangedViewClass     _ScrubberArrangedViewClass
	ScrubberArrangedViewClassOnce sync.Once
)

func getScrubberArrangedViewClass() _ScrubberArrangedViewClass {
	ScrubberArrangedViewClassOnce.Do(func() {
		ScrubberArrangedViewClass = _ScrubberArrangedViewClass{objc.GetClass("NSScrubberArrangedView")}
	})
	return ScrubberArrangedViewClass
}

type _ScrubberArrangedViewClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberArrangedView] class.
type IScrubberArrangedView interface {
	IView
	ApplyLayoutAttributes(layoutAttributes unsafe.Pointer)
}

// An abstract base class for the views whose layout is managed by a scrubber.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView
type ScrubberArrangedView struct {
	View
}

// ScrubberArrangedViewFrom constructs a [ScrubberArrangedView] from an unsafe.Pointer.
//
// An abstract base class for the views whose layout is managed by a scrubber.
func ScrubberArrangedViewFrom(ptr unsafe.Pointer) ScrubberArrangedView {
	return ScrubberArrangedView{
		View: ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberArrangedViewClass) Alloc() ScrubberArrangedView {
	rv := objc.Send[ScrubberArrangedView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScrubberArrangedViewClass) New() ScrubberArrangedView {
	rv := objc.Send[ScrubberArrangedView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberArrangedView) Init() ScrubberArrangedView {
	rv := objc.Send[ScrubberArrangedView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberArrangedView) Autorelease() ScrubberArrangedView {
	rv := objc.Send[ScrubberArrangedView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberArrangedView creates a new ScrubberArrangedView instance.
func NewScrubberArrangedView() ScrubberArrangedView {
	return getScrubberArrangedViewClass().New()
}


// Updates the layout of the arranged view to respect the provided layout attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/apply(_:)
func (s_ ScrubberArrangedView) ApplyLayoutAttributes(layoutAttributes unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("applyLayoutAttributes:"), layoutAttributes)
}

// A Boolean value that specifies whether the view is currently highlighted.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/isHighlighted
func (s_ ScrubberArrangedView) Highlighted() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("highlighted"))
	return rv
}


// SetHighlighted sets the value of the highlighted property.
// A Boolean value that specifies whether the view is currently highlighted.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/isHighlighted
func (s_ ScrubberArrangedView) SetHighlighted(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHighlighted:"), value)
}

// A Boolean value that specifies whether the current view is selected.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/isSelected
func (s_ ScrubberArrangedView) Selected() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("selected"))
	return rv
}


// SetSelected sets the value of the selected property.
// A Boolean value that specifies whether the current view is selected.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/isSelected
func (s_ ScrubberArrangedView) SetSelected(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelected:"), value)
}




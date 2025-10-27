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
	

	// properties:
	Highlighted() bool
	SetHighlighted(value bool)
	Selected() bool
	SetSelected(value bool)
	IsHighlighted() bool
	SetIsHighlighted(value bool)
	IsSelected() bool
	SetIsSelected(value bool)


	

	// methods:
	ApplyLayoutAttributes(layoutAttributes IScrubberLayoutAttributes)


}





// Alloc allocates a new instance without initialization.
func (sc _ScrubberArrangedViewClass) Alloc() ScrubberArrangedView {
	rv := objc.Send[ScrubberArrangedView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An abstract base class for the views whose layout is managed by a scrubber.


// An abstract base class for the views whose layout is managed by a scrubber.
//
// [Full Topic]
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




















// Updates the layout of the arranged view to respect the provided layout attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/apply(_:)
func (s_ ScrubberArrangedView) ApplyLayoutAttributes(layoutAttributes IScrubberLayoutAttributes) {
	objc.Send[objc.ID](s_.ID, objc.Sel("applyLayoutAttributes:"), layoutAttributes)
}







// A Boolean value that specifies whether the view is currently highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/isHighlighted
func (s_ ScrubberArrangedView) Highlighted() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("highlighted"))
	return rv
}


// A Boolean value that specifies whether the view is currently highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/isHighlighted
func (s_ ScrubberArrangedView) SetHighlighted(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHighlighted:"), value)
}


// A Boolean value that specifies whether the current view is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/isSelected
func (s_ ScrubberArrangedView) Selected() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("selected"))
	return rv
}


// A Boolean value that specifies whether the current view is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/isSelected
func (s_ ScrubberArrangedView) SetSelected(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelected:"), value)
}


// A Boolean value that specifies whether the view is currently highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberarrangedview/ishighlighted
func (s_ ScrubberArrangedView) IsHighlighted() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isHighlighted"))
	return rv
}


// A Boolean value that specifies whether the view is currently highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberarrangedview/ishighlighted
func (s_ ScrubberArrangedView) SetIsHighlighted(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsHighlighted:"), value)
}


// A Boolean value that specifies whether the current view is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberarrangedview/isselected
func (s_ ScrubberArrangedView) IsSelected() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSelected"))
	return rv
}


// A Boolean value that specifies whether the current view is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberarrangedview/isselected
func (s_ ScrubberArrangedView) SetIsSelected(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsSelected:"), value)
}









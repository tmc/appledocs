// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CandidateListTouchBarItem] class.
var (
	CandidateListTouchBarItemClass     _CandidateListTouchBarItemClass
	CandidateListTouchBarItemClassOnce sync.Once
)

func getCandidateListTouchBarItemClass() _CandidateListTouchBarItemClass {
	CandidateListTouchBarItemClassOnce.Do(func() {
		CandidateListTouchBarItemClass = _CandidateListTouchBarItemClass{objc.GetClass("NSCandidateListTouchBarItem")}
	})
	return CandidateListTouchBarItemClass
}

type _CandidateListTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [CandidateListTouchBarItem] class.
type ICandidateListTouchBarItem interface {
	ITouchBarItem
}

// A bar item that, along with its delegate, provides a list of textual suggestions for the current text view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem
type CandidateListTouchBarItem struct {
	TouchBarItem
}

// CandidateListTouchBarItemFrom constructs a [CandidateListTouchBarItem] from an unsafe.Pointer.
//
// A bar item that, along with its delegate, provides a list of textual suggestions for the current text view.
func CandidateListTouchBarItemFrom(ptr unsafe.Pointer) CandidateListTouchBarItem {
	return CandidateListTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CandidateListTouchBarItemClass) Alloc() CandidateListTouchBarItem {
	rv := objc.Send[CandidateListTouchBarItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CandidateListTouchBarItemClass) New() CandidateListTouchBarItem {
	rv := objc.Send[CandidateListTouchBarItem](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CandidateListTouchBarItem) Init() CandidateListTouchBarItem {
	rv := objc.Send[CandidateListTouchBarItem](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CandidateListTouchBarItem) Autorelease() CandidateListTouchBarItem {
	rv := objc.Send[CandidateListTouchBarItem](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCandidateListTouchBarItem creates a new CandidateListTouchBarItem instance.
func NewCandidateListTouchBarItem() CandidateListTouchBarItem {
	return getCandidateListTouchBarItemClass().New()
}

// A block that converts a candidate object into an attributed string for display in the candidate list item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/attributedStringForCandidate
func (c_ CandidateListTouchBarItem) AttributedStringForCandidate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("attributedStringForCandidate"))
	return rv
}

// SetAttributedStringForCandidate sets the value of the attributedStringForCandidate property.
// A block that converts a candidate object into an attributed string for display in the candidate list item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/attributedStringForCandidate
func (c_ CandidateListTouchBarItem) SetAttributedStringForCandidate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAttributedStringForCandidate:"), value)
}

// A Boolean value that controls the visibility of the candidate list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/isCollapsed
func (c_ CandidateListTouchBarItem) Collapsed() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("collapsed"))
	return rv
}

// SetCollapsed sets the value of the collapsed property.
// A Boolean value that controls the visibility of the candidate list.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/isCollapsed
func (c_ CandidateListTouchBarItem) SetCollapsed(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCollapsed:"), value)
}

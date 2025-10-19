// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CandidateListTouchBarItem] class.
var candidateListTouchBarItemClass = _CandidateListTouchBarItemClass{objc.GetClass("NSCandidateListTouchBarItem")}

type _CandidateListTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [CandidateListTouchBarItem] class.
type ICandidateListTouchBarItem interface {
	ITouchBarItem
}

// A bar item that, along with its delegate, provides a list of textual suggestions for the current text view. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return candidateListTouchBarItemClass.New()
}





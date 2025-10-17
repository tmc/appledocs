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




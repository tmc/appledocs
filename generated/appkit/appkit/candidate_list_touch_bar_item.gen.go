// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CandidateListTouchBarItem] class.
var CandidateListTouchBarItemClass objc.Class

func init() {
	CandidateListTouchBarItemClass = objc.GetClass("NSCandidateListTouchBarItem")
}

type CandidateListTouchBarItem struct {
	objc.ID
}

func CandidateListTouchBarItemFrom(ptr unsafe.Pointer) CandidateListTouchBarItem {
	return CandidateListTouchBarItem{
		ID: objc.ID(ptr),
	}
}




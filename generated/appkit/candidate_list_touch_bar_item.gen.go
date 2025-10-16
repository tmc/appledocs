
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CandidateListTouchBarItem] class.
var CandidateListTouchBarItemClass _CandidateListTouchBarItemClass

func init() {
	CandidateListTouchBarItemClass = _CandidateListTouchBarItemClass{objc.GetClass("NSCandidateListTouchBarItem")}
}

type _CandidateListTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [CandidateListTouchBarItem] class.
type ICandidateListTouchBarItem interface {
	ID() objc.ID
}

type CandidateListTouchBarItem struct {
	id objc.ID
}

func CandidateListTouchBarItemFrom(ptr unsafe.Pointer) CandidateListTouchBarItem {
	return CandidateListTouchBarItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ CandidateListTouchBarItem) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _CandidateListTouchBarItemClass) Alloc() CandidateListTouchBarItem {
	rv := objc.Send[CandidateListTouchBarItem](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _CandidateListTouchBarItemClass) New() CandidateListTouchBarItem {
	rv := objc.Send[CandidateListTouchBarItem](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewCandidateListTouchBarItem creates and returns a new initialized instance.
func NewCandidateListTouchBarItem() CandidateListTouchBarItem {
	return CandidateListTouchBarItemClass.New()
}

// Init initializes the instance.
func (c_ CandidateListTouchBarItem) Init() CandidateListTouchBarItem {
	rv := objc.Send[CandidateListTouchBarItem](c_.ID(), selInit)
	return rv
}

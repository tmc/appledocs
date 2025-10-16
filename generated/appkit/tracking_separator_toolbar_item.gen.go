
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TrackingSeparatorToolbarItem] class.
var TrackingSeparatorToolbarItemClass _TrackingSeparatorToolbarItemClass

func init() {
	TrackingSeparatorToolbarItemClass = _TrackingSeparatorToolbarItemClass{objc.GetClass("NSTrackingSeparatorToolbarItem")}
}

type _TrackingSeparatorToolbarItemClass struct {
	objc.Class
}

// An interface definition for the [TrackingSeparatorToolbarItem] class.
type ITrackingSeparatorToolbarItem interface {
	ID() objc.ID
}

type TrackingSeparatorToolbarItem struct {
	id objc.ID
}

func TrackingSeparatorToolbarItemFrom(ptr unsafe.Pointer) TrackingSeparatorToolbarItem {
	return TrackingSeparatorToolbarItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TrackingSeparatorToolbarItem) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TrackingSeparatorToolbarItemClass) Alloc() TrackingSeparatorToolbarItem {
	rv := objc.Send[TrackingSeparatorToolbarItem](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TrackingSeparatorToolbarItemClass) New() TrackingSeparatorToolbarItem {
	rv := objc.Send[TrackingSeparatorToolbarItem](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTrackingSeparatorToolbarItem creates and returns a new initialized instance.
func NewTrackingSeparatorToolbarItem() TrackingSeparatorToolbarItem {
	return TrackingSeparatorToolbarItemClass.New()
}

// Init initializes the instance.
func (t_ TrackingSeparatorToolbarItem) Init() TrackingSeparatorToolbarItem {
	rv := objc.Send[TrackingSeparatorToolbarItem](t_.ID(), selInit)
	return rv
}

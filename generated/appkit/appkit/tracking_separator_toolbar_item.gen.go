// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TrackingSeparatorToolbarItem] class.
var TrackingSeparatorToolbarItemClass objc.Class

func init() {
	TrackingSeparatorToolbarItemClass = objc.GetClass("NSTrackingSeparatorToolbarItem")
}

type TrackingSeparatorToolbarItem struct {
	objc.ID
}

func TrackingSeparatorToolbarItemFrom(ptr unsafe.Pointer) TrackingSeparatorToolbarItem {
	return TrackingSeparatorToolbarItem{
		ID: objc.ID(ptr),
	}
}





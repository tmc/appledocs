// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Scroller] class.
var ScrollerClass objc.Class

func init() {
	ScrollerClass = objc.GetClass("NSScroller")
}

type Scroller struct {
	objc.ID
}

func ScrollerFrom(ptr unsafe.Pointer) Scroller {
	return Scroller{
		ID: objc.ID(ptr),
	}
}





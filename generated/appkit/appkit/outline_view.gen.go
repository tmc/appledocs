// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OutlineView] class.
var OutlineViewClass objc.Class

func init() {
	OutlineViewClass = objc.GetClass("NSOutlineView")
}

type OutlineView struct {
	objc.ID
}

func OutlineViewFrom(ptr unsafe.Pointer) OutlineView {
	return OutlineView{
		ID: objc.ID(ptr),
	}
}




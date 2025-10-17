// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LayoutManager] class.
var LayoutManagerClass objc.Class

func init() {
	LayoutManagerClass = objc.GetClass("NSLayoutManager")
}

type LayoutManager struct {
	objc.ID
}

func LayoutManagerFrom(ptr unsafe.Pointer) LayoutManager {
	return LayoutManager{
		ID: objc.ID(ptr),
	}
}




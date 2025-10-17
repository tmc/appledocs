// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [StatusItem] class.
var StatusItemClass objc.Class

func init() {
	StatusItemClass = objc.GetClass("NSStatusItem")
}

type StatusItem struct {
	objc.ID
}

func StatusItemFrom(ptr unsafe.Pointer) StatusItem {
	return StatusItem{
		ID: objc.ID(ptr),
	}
}





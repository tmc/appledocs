// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TouchBarItem] class.
var TouchBarItemClass objc.Class

func init() {
	TouchBarItemClass = objc.GetClass("NSTouchBarItem")
}

type TouchBarItem struct {
	objc.ID
}

func TouchBarItemFrom(ptr unsafe.Pointer) TouchBarItem {
	return TouchBarItem{
		ID: objc.ID(ptr),
	}
}




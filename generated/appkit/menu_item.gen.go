// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MenuItem] class.
var MenuItemClass objc.Class

func init() {
	MenuItemClass = objc.GetClass("NSMenuItem")
}

type MenuItem struct {
	objc.ID
}

func MenuItemFrom(ptr unsafe.Pointer) MenuItem {
	return MenuItem{
		ID: objc.ID(ptr),
	}
}




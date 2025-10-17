// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Drawer] class.
var DrawerClass objc.Class

func init() {
	DrawerClass = objc.GetClass("NSDrawer")
}

type Drawer struct {
	objc.ID
}

func DrawerFrom(ptr unsafe.Pointer) Drawer {
	return Drawer{
		ID: objc.ID(ptr),
	}
}





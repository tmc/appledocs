// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DockTile] class.
var DockTileClass objc.Class

func init() {
	DockTileClass = objc.GetClass("NSDockTile")
}

type DockTile struct {
	objc.ID
}

func DockTileFrom(ptr unsafe.Pointer) DockTile {
	return DockTile{
		ID: objc.ID(ptr),
	}
}





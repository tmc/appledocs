// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DockTile] class.
var dockTileClass = _DockTileClass{objc.GetClass("NSDockTile")}

type _DockTileClass struct {
	class objc.Class
}

// The visual representation of your app’s miniaturized windows and app icon as they appear in the Dock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDockTile

type DockTile struct {
	objectivec.Object
}

// DockTileFrom constructs a [DockTile] from an unsafe.Pointer.
//
// The visual representation of your app’s miniaturized windows and app icon as they appear in the Dock.
func DockTileFrom(ptr unsafe.Pointer) DockTile {
	return DockTile{objectivec.Object{objc.ID(ptr)}}
}




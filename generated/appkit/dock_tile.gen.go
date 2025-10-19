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

// An interface definition for the [DockTile] class.
type IDockTile interface {
	objectivec.IObject
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
// Alloc allocates a new instance without initialization.
func (dc _DockTileClass) Alloc() DockTile {
	rv := objc.Send[DockTile](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (dc _DockTileClass) New() DockTile {
	rv := objc.Send[DockTile](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DockTile) Init() DockTile {
	rv := objc.Send[DockTile](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DockTile) Autorelease() DockTile {
	rv := objc.Send[DockTile](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDockTile creates a new DockTile instance.
func NewDockTile() DockTile {
	return dockTileClass.New()
}





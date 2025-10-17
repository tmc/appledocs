
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DockTile] class.
var DockTileClass _DockTileClass

func init() {
	DockTileClass = _DockTileClass{objc.GetClass("NSDockTile")}
}

type _DockTileClass struct {
	objc.Class
}

// An interface definition for the [DockTile] class.
type IDockTile interface {
	ID() objc.ID
}

type DockTile struct {
	id objc.ID
}

func DockTileFrom(ptr unsafe.Pointer) DockTile {
	return DockTile{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ DockTile) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _DockTileClass) Alloc() DockTile {
	rv := objc.Send[DockTile](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _DockTileClass) New() DockTile {
	rv := objc.Send[DockTile](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewDockTile creates and returns a new initialized instance.
func NewDockTile() DockTile {
	return DockTileClass.New()
}

// Init initializes the instance.
func (d_ DockTile) Init() DockTile {
	rv := objc.Send[DockTile](d_.ID(), selInit)
	return rv
}

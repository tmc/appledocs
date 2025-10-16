
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GridCell] class.
var GridCellClass _GridCellClass

func init() {
	GridCellClass = _GridCellClass{objc.GetClass("NSGridCell")}
}

type _GridCellClass struct {
	objc.Class
}

// An interface definition for the [GridCell] class.
type IGridCell interface {
	ID() objc.ID
}

type GridCell struct {
	id objc.ID
}

func GridCellFrom(ptr unsafe.Pointer) GridCell {
	return GridCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (g_ GridCell) ID() objc.ID {
	return g_.id
}

// Alloc allocates a new instance without initialization.
func (gc _GridCellClass) Alloc() GridCell {
	rv := objc.Send[GridCell](objc.ID(gc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (gc _GridCellClass) New() GridCell {
	rv := objc.Send[GridCell](objc.ID(gc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewGridCell creates and returns a new initialized instance.
func NewGridCell() GridCell {
	return GridCellClass.New()
}

// Init initializes the instance.
func (g_ GridCell) Init() GridCell {
	rv := objc.Send[GridCell](g_.ID(), selInit)
	return rv
}

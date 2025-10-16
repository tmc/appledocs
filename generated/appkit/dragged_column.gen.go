
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [draggedColumn] class.
var draggedColumnClass _draggedColumnClass

func init() {
	draggedColumnClass = _draggedColumnClass{objc.GetClass("draggedColumn")}
}

type _draggedColumnClass struct {
	objc.Class
}

// An interface definition for the [draggedColumn] class.
type IdraggedColumn interface {
	ID() objc.ID
}

type draggedColumn struct {
	id objc.ID
}

func draggedColumnFrom(ptr unsafe.Pointer) draggedColumn {
	return draggedColumn{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ draggedColumn) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _draggedColumnClass) Alloc() draggedColumn {
	rv := objc.Send[draggedColumn](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _draggedColumnClass) New() draggedColumn {
	rv := objc.Send[draggedColumn](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdraggedColumn creates and returns a new initialized instance.
func NewdraggedColumn() draggedColumn {
	return draggedColumnClass.New()
}

// Init initializes the instance.
func (d_ draggedColumn) Init() draggedColumn {
	rv := objc.Send[draggedColumn](d_.ID(), selInit)
	return rv
}

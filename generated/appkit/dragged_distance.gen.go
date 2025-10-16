
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [draggedDistance] class.
var draggedDistanceClass _draggedDistanceClass

func init() {
	draggedDistanceClass = _draggedDistanceClass{objc.GetClass("draggedDistance")}
}

type _draggedDistanceClass struct {
	objc.Class
}

// An interface definition for the [draggedDistance] class.
type IdraggedDistance interface {
	ID() objc.ID
}

type draggedDistance struct {
	id objc.ID
}

func draggedDistanceFrom(ptr unsafe.Pointer) draggedDistance {
	return draggedDistance{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ draggedDistance) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _draggedDistanceClass) Alloc() draggedDistance {
	rv := objc.Send[draggedDistance](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _draggedDistanceClass) New() draggedDistance {
	rv := objc.Send[draggedDistance](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdraggedDistance creates and returns a new initialized instance.
func NewdraggedDistance() draggedDistance {
	return draggedDistanceClass.New()
}

// Init initializes the instance.
func (d_ draggedDistance) Init() draggedDistance {
	rv := objc.Send[draggedDistance](d_.ID(), selInit)
	return rv
}

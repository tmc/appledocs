
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [draggingImageComponents] class.
var draggingImageComponentsClass _draggingImageComponentsClass

func init() {
	draggingImageComponentsClass = _draggingImageComponentsClass{objc.GetClass("draggingImageComponents")}
}

type _draggingImageComponentsClass struct {
	objc.Class
}

// An interface definition for the [draggingImageComponents] class.
type IdraggingImageComponents interface {
	ID() objc.ID
}

type draggingImageComponents struct {
	id objc.ID
}

func draggingImageComponentsFrom(ptr unsafe.Pointer) draggingImageComponents {
	return draggingImageComponents{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ draggingImageComponents) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _draggingImageComponentsClass) Alloc() draggingImageComponents {
	rv := objc.Send[draggingImageComponents](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _draggingImageComponentsClass) New() draggingImageComponents {
	rv := objc.Send[draggingImageComponents](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdraggingImageComponents creates and returns a new initialized instance.
func NewdraggingImageComponents() draggingImageComponents {
	return draggingImageComponentsClass.New()
}

// Init initializes the instance.
func (d_ draggingImageComponents) Init() draggingImageComponents {
	rv := objc.Send[draggingImageComponents](d_.ID(), selInit)
	return rv
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DraggingImageComponent] class.
var DraggingImageComponentClass _DraggingImageComponentClass

func init() {
	DraggingImageComponentClass = _DraggingImageComponentClass{objc.GetClass("NSDraggingImageComponent")}
}

type _DraggingImageComponentClass struct {
	objc.Class
}

// An interface definition for the [DraggingImageComponent] class.
type IDraggingImageComponent interface {
	ID() objc.ID
}

type DraggingImageComponent struct {
	id objc.ID
}

func DraggingImageComponentFrom(ptr unsafe.Pointer) DraggingImageComponent {
	return DraggingImageComponent{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ DraggingImageComponent) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _DraggingImageComponentClass) Alloc() DraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _DraggingImageComponentClass) New() DraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewDraggingImageComponent creates and returns a new initialized instance.
func NewDraggingImageComponent() DraggingImageComponent {
	return DraggingImageComponentClass.New()
}

// Init initializes the instance.
func (d_ DraggingImageComponent) Init() DraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](d_.ID(), selInit)
	return rv
}

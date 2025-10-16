
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [selectionOverlayStyle] class.
var selectionOverlayStyleClass _selectionOverlayStyleClass

func init() {
	selectionOverlayStyleClass = _selectionOverlayStyleClass{objc.GetClass("selectionOverlayStyle")}
}

type _selectionOverlayStyleClass struct {
	objc.Class
}

// An interface definition for the [selectionOverlayStyle] class.
type IselectionOverlayStyle interface {
	ID() objc.ID
}

type selectionOverlayStyle struct {
	id objc.ID
}

func selectionOverlayStyleFrom(ptr unsafe.Pointer) selectionOverlayStyle {
	return selectionOverlayStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ selectionOverlayStyle) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _selectionOverlayStyleClass) Alloc() selectionOverlayStyle {
	rv := objc.Send[selectionOverlayStyle](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _selectionOverlayStyleClass) New() selectionOverlayStyle {
	rv := objc.Send[selectionOverlayStyle](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewselectionOverlayStyle creates and returns a new initialized instance.
func NewselectionOverlayStyle() selectionOverlayStyle {
	return selectionOverlayStyleClass.New()
}

// Init initializes the instance.
func (s_ selectionOverlayStyle) Init() selectionOverlayStyle {
	rv := objc.Send[selectionOverlayStyle](s_.ID(), selInit)
	return rv
}

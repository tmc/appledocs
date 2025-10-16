
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [selectionHighlightStyle] class.
var selectionHighlightStyleClass _selectionHighlightStyleClass

func init() {
	selectionHighlightStyleClass = _selectionHighlightStyleClass{objc.GetClass("selectionHighlightStyle")}
}

type _selectionHighlightStyleClass struct {
	objc.Class
}

// An interface definition for the [selectionHighlightStyle] class.
type IselectionHighlightStyle interface {
	ID() objc.ID
}

type selectionHighlightStyle struct {
	id objc.ID
}

func selectionHighlightStyleFrom(ptr unsafe.Pointer) selectionHighlightStyle {
	return selectionHighlightStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ selectionHighlightStyle) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _selectionHighlightStyleClass) Alloc() selectionHighlightStyle {
	rv := objc.Send[selectionHighlightStyle](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _selectionHighlightStyleClass) New() selectionHighlightStyle {
	rv := objc.Send[selectionHighlightStyle](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewselectionHighlightStyle creates and returns a new initialized instance.
func NewselectionHighlightStyle() selectionHighlightStyle {
	return selectionHighlightStyleClass.New()
}

// Init initializes the instance.
func (s_ selectionHighlightStyle) Init() selectionHighlightStyle {
	rv := objc.Send[selectionHighlightStyle](s_.ID(), selInit)
	return rv
}

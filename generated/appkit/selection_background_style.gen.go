
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [selectionBackgroundStyle] class.
var selectionBackgroundStyleClass _selectionBackgroundStyleClass

func init() {
	selectionBackgroundStyleClass = _selectionBackgroundStyleClass{objc.GetClass("selectionBackgroundStyle")}
}

type _selectionBackgroundStyleClass struct {
	objc.Class
}

// An interface definition for the [selectionBackgroundStyle] class.
type IselectionBackgroundStyle interface {
	ID() objc.ID
}

type selectionBackgroundStyle struct {
	id objc.ID
}

func selectionBackgroundStyleFrom(ptr unsafe.Pointer) selectionBackgroundStyle {
	return selectionBackgroundStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ selectionBackgroundStyle) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _selectionBackgroundStyleClass) Alloc() selectionBackgroundStyle {
	rv := objc.Send[selectionBackgroundStyle](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _selectionBackgroundStyleClass) New() selectionBackgroundStyle {
	rv := objc.Send[selectionBackgroundStyle](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewselectionBackgroundStyle creates and returns a new initialized instance.
func NewselectionBackgroundStyle() selectionBackgroundStyle {
	return selectionBackgroundStyleClass.New()
}

// Init initializes the instance.
func (s_ selectionBackgroundStyle) Init() selectionBackgroundStyle {
	rv := objc.Send[selectionBackgroundStyle](s_.ID(), selInit)
	return rv
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [rowSizeStyle] class.
var rowSizeStyleClass _rowSizeStyleClass

func init() {
	rowSizeStyleClass = _rowSizeStyleClass{objc.GetClass("rowSizeStyle")}
}

type _rowSizeStyleClass struct {
	objc.Class
}

// An interface definition for the [rowSizeStyle] class.
type IrowSizeStyle interface {
	ID() objc.ID
}

type rowSizeStyle struct {
	id objc.ID
}

func rowSizeStyleFrom(ptr unsafe.Pointer) rowSizeStyle {
	return rowSizeStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ rowSizeStyle) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _rowSizeStyleClass) Alloc() rowSizeStyle {
	rv := objc.Send[rowSizeStyle](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _rowSizeStyleClass) New() rowSizeStyle {
	rv := objc.Send[rowSizeStyle](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrowSizeStyle creates and returns a new initialized instance.
func NewrowSizeStyle() rowSizeStyle {
	return rowSizeStyleClass.New()
}

// Init initializes the instance.
func (r_ rowSizeStyle) Init() rowSizeStyle {
	rv := objc.Send[rowSizeStyle](r_.ID(), selInit)
	return rv
}

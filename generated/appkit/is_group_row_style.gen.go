
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isGroupRowStyle] class.
var isGroupRowStyleClass _isGroupRowStyleClass

func init() {
	isGroupRowStyleClass = _isGroupRowStyleClass{objc.GetClass("isGroupRowStyle")}
}

type _isGroupRowStyleClass struct {
	objc.Class
}

// An interface definition for the [isGroupRowStyle] class.
type IisGroupRowStyle interface {
	ID() objc.ID
}

type isGroupRowStyle struct {
	id objc.ID
}

func isGroupRowStyleFrom(ptr unsafe.Pointer) isGroupRowStyle {
	return isGroupRowStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isGroupRowStyle) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isGroupRowStyleClass) Alloc() isGroupRowStyle {
	rv := objc.Send[isGroupRowStyle](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isGroupRowStyleClass) New() isGroupRowStyle {
	rv := objc.Send[isGroupRowStyle](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisGroupRowStyle creates and returns a new initialized instance.
func NewisGroupRowStyle() isGroupRowStyle {
	return isGroupRowStyleClass.New()
}

// Init initializes the instance.
func (i_ isGroupRowStyle) Init() isGroupRowStyle {
	rv := objc.Send[isGroupRowStyle](i_.ID(), selInit)
	return rv
}

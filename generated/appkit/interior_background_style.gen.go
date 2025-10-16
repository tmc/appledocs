
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [interiorBackgroundStyle] class.
var interiorBackgroundStyleClass _interiorBackgroundStyleClass

func init() {
	interiorBackgroundStyleClass = _interiorBackgroundStyleClass{objc.GetClass("interiorBackgroundStyle")}
}

type _interiorBackgroundStyleClass struct {
	objc.Class
}

// An interface definition for the [interiorBackgroundStyle] class.
type IinteriorBackgroundStyle interface {
	ID() objc.ID
}

type interiorBackgroundStyle struct {
	id objc.ID
}

func interiorBackgroundStyleFrom(ptr unsafe.Pointer) interiorBackgroundStyle {
	return interiorBackgroundStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ interiorBackgroundStyle) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _interiorBackgroundStyleClass) Alloc() interiorBackgroundStyle {
	rv := objc.Send[interiorBackgroundStyle](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _interiorBackgroundStyleClass) New() interiorBackgroundStyle {
	rv := objc.Send[interiorBackgroundStyle](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewinteriorBackgroundStyle creates and returns a new initialized instance.
func NewinteriorBackgroundStyle() interiorBackgroundStyle {
	return interiorBackgroundStyleClass.New()
}

// Init initializes the instance.
func (i_ interiorBackgroundStyle) Init() interiorBackgroundStyle {
	rv := objc.Send[interiorBackgroundStyle](i_.ID(), selInit)
	return rv
}

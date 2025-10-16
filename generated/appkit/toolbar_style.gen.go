
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [toolbarStyle] class.
var toolbarStyleClass _toolbarStyleClass

func init() {
	toolbarStyleClass = _toolbarStyleClass{objc.GetClass("toolbarStyle")}
}

type _toolbarStyleClass struct {
	objc.Class
}

// An interface definition for the [toolbarStyle] class.
type ItoolbarStyle interface {
	ID() objc.ID
}

type toolbarStyle struct {
	id objc.ID
}

func toolbarStyleFrom(ptr unsafe.Pointer) toolbarStyle {
	return toolbarStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ toolbarStyle) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _toolbarStyleClass) Alloc() toolbarStyle {
	rv := objc.Send[toolbarStyle](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _toolbarStyleClass) New() toolbarStyle {
	rv := objc.Send[toolbarStyle](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtoolbarStyle creates and returns a new initialized instance.
func NewtoolbarStyle() toolbarStyle {
	return toolbarStyleClass.New()
}

// Init initializes the instance.
func (t_ toolbarStyle) Init() toolbarStyle {
	rv := objc.Send[toolbarStyle](t_.ID(), selInit)
	return rv
}

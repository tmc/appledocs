
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [viewDidMoveToWindow] class.
var viewDidMoveToWindowClass _viewDidMoveToWindowClass

func init() {
	viewDidMoveToWindowClass = _viewDidMoveToWindowClass{objc.GetClass("viewDidMoveToWindow")}
}

type _viewDidMoveToWindowClass struct {
	objc.Class
}

// An interface definition for the [viewDidMoveToWindow] class.
type IviewDidMoveToWindow interface {
	ID() objc.ID
}

type viewDidMoveToWindow struct {
	id objc.ID
}

func viewDidMoveToWindowFrom(ptr unsafe.Pointer) viewDidMoveToWindow {
	return viewDidMoveToWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ viewDidMoveToWindow) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _viewDidMoveToWindowClass) Alloc() viewDidMoveToWindow {
	rv := objc.Send[viewDidMoveToWindow](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _viewDidMoveToWindowClass) New() viewDidMoveToWindow {
	rv := objc.Send[viewDidMoveToWindow](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewviewDidMoveToWindow creates and returns a new initialized instance.
func NewviewDidMoveToWindow() viewDidMoveToWindow {
	return viewDidMoveToWindowClass.New()
}

// Init initializes the instance.
func (v_ viewDidMoveToWindow) Init() viewDidMoveToWindow {
	rv := objc.Send[viewDidMoveToWindow](v_.ID(), selInit)
	return rv
}

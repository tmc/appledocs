
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [viewDidUnhide] class.
var viewDidUnhideClass _viewDidUnhideClass

func init() {
	viewDidUnhideClass = _viewDidUnhideClass{objc.GetClass("viewDidUnhide")}
}

type _viewDidUnhideClass struct {
	objc.Class
}

// An interface definition for the [viewDidUnhide] class.
type IviewDidUnhide interface {
	ID() objc.ID
}

type viewDidUnhide struct {
	id objc.ID
}

func viewDidUnhideFrom(ptr unsafe.Pointer) viewDidUnhide {
	return viewDidUnhide{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ viewDidUnhide) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _viewDidUnhideClass) Alloc() viewDidUnhide {
	rv := objc.Send[viewDidUnhide](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _viewDidUnhideClass) New() viewDidUnhide {
	rv := objc.Send[viewDidUnhide](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewviewDidUnhide creates and returns a new initialized instance.
func NewviewDidUnhide() viewDidUnhide {
	return viewDidUnhideClass.New()
}

// Init initializes the instance.
func (v_ viewDidUnhide) Init() viewDidUnhide {
	rv := objc.Send[viewDidUnhide](v_.ID(), selInit)
	return rv
}

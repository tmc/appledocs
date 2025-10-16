
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [viewDidHide] class.
var viewDidHideClass _viewDidHideClass

func init() {
	viewDidHideClass = _viewDidHideClass{objc.GetClass("viewDidHide")}
}

type _viewDidHideClass struct {
	objc.Class
}

// An interface definition for the [viewDidHide] class.
type IviewDidHide interface {
	ID() objc.ID
}

type viewDidHide struct {
	id objc.ID
}

func viewDidHideFrom(ptr unsafe.Pointer) viewDidHide {
	return viewDidHide{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ viewDidHide) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _viewDidHideClass) Alloc() viewDidHide {
	rv := objc.Send[viewDidHide](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _viewDidHideClass) New() viewDidHide {
	rv := objc.Send[viewDidHide](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewviewDidHide creates and returns a new initialized instance.
func NewviewDidHide() viewDidHide {
	return viewDidHideClass.New()
}

// Init initializes the instance.
func (v_ viewDidHide) Init() viewDidHide {
	rv := objc.Send[viewDidHide](v_.ID(), selInit)
	return rv
}

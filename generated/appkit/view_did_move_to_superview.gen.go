
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [viewDidMoveToSuperview] class.
var viewDidMoveToSuperviewClass _viewDidMoveToSuperviewClass

func init() {
	viewDidMoveToSuperviewClass = _viewDidMoveToSuperviewClass{objc.GetClass("viewDidMoveToSuperview")}
}

type _viewDidMoveToSuperviewClass struct {
	objc.Class
}

// An interface definition for the [viewDidMoveToSuperview] class.
type IviewDidMoveToSuperview interface {
	ID() objc.ID
}

type viewDidMoveToSuperview struct {
	id objc.ID
}

func viewDidMoveToSuperviewFrom(ptr unsafe.Pointer) viewDidMoveToSuperview {
	return viewDidMoveToSuperview{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ viewDidMoveToSuperview) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _viewDidMoveToSuperviewClass) Alloc() viewDidMoveToSuperview {
	rv := objc.Send[viewDidMoveToSuperview](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _viewDidMoveToSuperviewClass) New() viewDidMoveToSuperview {
	rv := objc.Send[viewDidMoveToSuperview](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewviewDidMoveToSuperview creates and returns a new initialized instance.
func NewviewDidMoveToSuperview() viewDidMoveToSuperview {
	return viewDidMoveToSuperviewClass.New()
}

// Init initializes the instance.
func (v_ viewDidMoveToSuperview) Init() viewDidMoveToSuperview {
	rv := objc.Send[viewDidMoveToSuperview](v_.ID(), selInit)
	return rv
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [viewWillDraw] class.
var viewWillDrawClass _viewWillDrawClass

func init() {
	viewWillDrawClass = _viewWillDrawClass{objc.GetClass("viewWillDraw")}
}

type _viewWillDrawClass struct {
	objc.Class
}

// An interface definition for the [viewWillDraw] class.
type IviewWillDraw interface {
	ID() objc.ID
}

type viewWillDraw struct {
	id objc.ID
}

func viewWillDrawFrom(ptr unsafe.Pointer) viewWillDraw {
	return viewWillDraw{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ viewWillDraw) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _viewWillDrawClass) Alloc() viewWillDraw {
	rv := objc.Send[viewWillDraw](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _viewWillDrawClass) New() viewWillDraw {
	rv := objc.Send[viewWillDraw](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewviewWillDraw creates and returns a new initialized instance.
func NewviewWillDraw() viewWillDraw {
	return viewWillDrawClass.New()
}

// Init initializes the instance.
func (v_ viewWillDraw) Init() viewWillDraw {
	rv := objc.Send[viewWillDraw](v_.ID(), selInit)
	return rv
}

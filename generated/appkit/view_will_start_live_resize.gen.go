
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [viewWillStartLiveResize] class.
var viewWillStartLiveResizeClass _viewWillStartLiveResizeClass

func init() {
	viewWillStartLiveResizeClass = _viewWillStartLiveResizeClass{objc.GetClass("viewWillStartLiveResize")}
}

type _viewWillStartLiveResizeClass struct {
	objc.Class
}

// An interface definition for the [viewWillStartLiveResize] class.
type IviewWillStartLiveResize interface {
	ID() objc.ID
}

type viewWillStartLiveResize struct {
	id objc.ID
}

func viewWillStartLiveResizeFrom(ptr unsafe.Pointer) viewWillStartLiveResize {
	return viewWillStartLiveResize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ viewWillStartLiveResize) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _viewWillStartLiveResizeClass) Alloc() viewWillStartLiveResize {
	rv := objc.Send[viewWillStartLiveResize](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _viewWillStartLiveResizeClass) New() viewWillStartLiveResize {
	rv := objc.Send[viewWillStartLiveResize](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewviewWillStartLiveResize creates and returns a new initialized instance.
func NewviewWillStartLiveResize() viewWillStartLiveResize {
	return viewWillStartLiveResizeClass.New()
}

// Init initializes the instance.
func (v_ viewWillStartLiveResize) Init() viewWillStartLiveResize {
	rv := objc.Send[viewWillStartLiveResize](v_.ID(), selInit)
	return rv
}

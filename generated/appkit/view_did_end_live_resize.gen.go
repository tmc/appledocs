
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [viewDidEndLiveResize] class.
var viewDidEndLiveResizeClass _viewDidEndLiveResizeClass

func init() {
	viewDidEndLiveResizeClass = _viewDidEndLiveResizeClass{objc.GetClass("viewDidEndLiveResize")}
}

type _viewDidEndLiveResizeClass struct {
	objc.Class
}

// An interface definition for the [viewDidEndLiveResize] class.
type IviewDidEndLiveResize interface {
	ID() objc.ID
}

type viewDidEndLiveResize struct {
	id objc.ID
}

func viewDidEndLiveResizeFrom(ptr unsafe.Pointer) viewDidEndLiveResize {
	return viewDidEndLiveResize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ viewDidEndLiveResize) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _viewDidEndLiveResizeClass) Alloc() viewDidEndLiveResize {
	rv := objc.Send[viewDidEndLiveResize](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _viewDidEndLiveResizeClass) New() viewDidEndLiveResize {
	rv := objc.Send[viewDidEndLiveResize](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewviewDidEndLiveResize creates and returns a new initialized instance.
func NewviewDidEndLiveResize() viewDidEndLiveResize {
	return viewDidEndLiveResizeClass.New()
}

// Init initializes the instance.
func (v_ viewDidEndLiveResize) Init() viewDidEndLiveResize {
	rv := objc.Send[viewDidEndLiveResize](v_.ID(), selInit)
	return rv
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [inLiveResize] class.
var inLiveResizeClass _inLiveResizeClass

func init() {
	inLiveResizeClass = _inLiveResizeClass{objc.GetClass("inLiveResize")}
}

type _inLiveResizeClass struct {
	objc.Class
}

// An interface definition for the [inLiveResize] class.
type IinLiveResize interface {
	ID() objc.ID
}

type inLiveResize struct {
	id objc.ID
}

func inLiveResizeFrom(ptr unsafe.Pointer) inLiveResize {
	return inLiveResize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ inLiveResize) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _inLiveResizeClass) Alloc() inLiveResize {
	rv := objc.Send[inLiveResize](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _inLiveResizeClass) New() inLiveResize {
	rv := objc.Send[inLiveResize](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewinLiveResize creates and returns a new initialized instance.
func NewinLiveResize() inLiveResize {
	return inLiveResizeClass.New()
}

// Init initializes the instance.
func (i_ inLiveResize) Init() inLiveResize {
	rv := objc.Send[inLiveResize](i_.ID(), selInit)
	return rv
}

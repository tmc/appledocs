
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [rectPreservedDuringLiveResize] class.
var rectPreservedDuringLiveResizeClass _rectPreservedDuringLiveResizeClass

func init() {
	rectPreservedDuringLiveResizeClass = _rectPreservedDuringLiveResizeClass{objc.GetClass("rectPreservedDuringLiveResize")}
}

type _rectPreservedDuringLiveResizeClass struct {
	objc.Class
}

// An interface definition for the [rectPreservedDuringLiveResize] class.
type IrectPreservedDuringLiveResize interface {
	ID() objc.ID
}

type rectPreservedDuringLiveResize struct {
	id objc.ID
}

func rectPreservedDuringLiveResizeFrom(ptr unsafe.Pointer) rectPreservedDuringLiveResize {
	return rectPreservedDuringLiveResize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ rectPreservedDuringLiveResize) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _rectPreservedDuringLiveResizeClass) Alloc() rectPreservedDuringLiveResize {
	rv := objc.Send[rectPreservedDuringLiveResize](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _rectPreservedDuringLiveResizeClass) New() rectPreservedDuringLiveResize {
	rv := objc.Send[rectPreservedDuringLiveResize](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrectPreservedDuringLiveResize creates and returns a new initialized instance.
func NewrectPreservedDuringLiveResize() rectPreservedDuringLiveResize {
	return rectPreservedDuringLiveResizeClass.New()
}

// Init initializes the instance.
func (r_ rectPreservedDuringLiveResize) Init() rectPreservedDuringLiveResize {
	rv := objc.Send[rectPreservedDuringLiveResize](r_.ID(), selInit)
	return rv
}

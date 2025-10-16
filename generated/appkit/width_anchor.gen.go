
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [widthAnchor] class.
var widthAnchorClass _widthAnchorClass

func init() {
	widthAnchorClass = _widthAnchorClass{objc.GetClass("widthAnchor")}
}

type _widthAnchorClass struct {
	objc.Class
}

// An interface definition for the [widthAnchor] class.
type IwidthAnchor interface {
	ID() objc.ID
}

type widthAnchor struct {
	id objc.ID
}

func widthAnchorFrom(ptr unsafe.Pointer) widthAnchor {
	return widthAnchor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ widthAnchor) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _widthAnchorClass) Alloc() widthAnchor {
	rv := objc.Send[widthAnchor](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _widthAnchorClass) New() widthAnchor {
	rv := objc.Send[widthAnchor](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewwidthAnchor creates and returns a new initialized instance.
func NewwidthAnchor() widthAnchor {
	return widthAnchorClass.New()
}

// Init initializes the instance.
func (w_ widthAnchor) Init() widthAnchor {
	rv := objc.Send[widthAnchor](w_.ID(), selInit)
	return rv
}

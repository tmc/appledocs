
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [leftAnchor] class.
var leftAnchorClass _leftAnchorClass

func init() {
	leftAnchorClass = _leftAnchorClass{objc.GetClass("leftAnchor")}
}

type _leftAnchorClass struct {
	objc.Class
}

// An interface definition for the [leftAnchor] class.
type IleftAnchor interface {
	ID() objc.ID
}

type leftAnchor struct {
	id objc.ID
}

func leftAnchorFrom(ptr unsafe.Pointer) leftAnchor {
	return leftAnchor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ leftAnchor) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _leftAnchorClass) Alloc() leftAnchor {
	rv := objc.Send[leftAnchor](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _leftAnchorClass) New() leftAnchor {
	rv := objc.Send[leftAnchor](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewleftAnchor creates and returns a new initialized instance.
func NewleftAnchor() leftAnchor {
	return leftAnchorClass.New()
}

// Init initializes the instance.
func (l_ leftAnchor) Init() leftAnchor {
	rv := objc.Send[leftAnchor](l_.ID(), selInit)
	return rv
}

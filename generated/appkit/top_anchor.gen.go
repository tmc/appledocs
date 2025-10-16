
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [topAnchor] class.
var topAnchorClass _topAnchorClass

func init() {
	topAnchorClass = _topAnchorClass{objc.GetClass("topAnchor")}
}

type _topAnchorClass struct {
	objc.Class
}

// An interface definition for the [topAnchor] class.
type ItopAnchor interface {
	ID() objc.ID
}

type topAnchor struct {
	id objc.ID
}

func topAnchorFrom(ptr unsafe.Pointer) topAnchor {
	return topAnchor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ topAnchor) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _topAnchorClass) Alloc() topAnchor {
	rv := objc.Send[topAnchor](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _topAnchorClass) New() topAnchor {
	rv := objc.Send[topAnchor](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtopAnchor creates and returns a new initialized instance.
func NewtopAnchor() topAnchor {
	return topAnchorClass.New()
}

// Init initializes the instance.
func (t_ topAnchor) Init() topAnchor {
	rv := objc.Send[topAnchor](t_.ID(), selInit)
	return rv
}

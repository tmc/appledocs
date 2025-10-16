
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [trailingAnchor] class.
var trailingAnchorClass _trailingAnchorClass

func init() {
	trailingAnchorClass = _trailingAnchorClass{objc.GetClass("trailingAnchor")}
}

type _trailingAnchorClass struct {
	objc.Class
}

// An interface definition for the [trailingAnchor] class.
type ItrailingAnchor interface {
	ID() objc.ID
}

type trailingAnchor struct {
	id objc.ID
}

func trailingAnchorFrom(ptr unsafe.Pointer) trailingAnchor {
	return trailingAnchor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ trailingAnchor) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _trailingAnchorClass) Alloc() trailingAnchor {
	rv := objc.Send[trailingAnchor](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _trailingAnchorClass) New() trailingAnchor {
	rv := objc.Send[trailingAnchor](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtrailingAnchor creates and returns a new initialized instance.
func NewtrailingAnchor() trailingAnchor {
	return trailingAnchorClass.New()
}

// Init initializes the instance.
func (t_ trailingAnchor) Init() trailingAnchor {
	rv := objc.Send[trailingAnchor](t_.ID(), selInit)
	return rv
}

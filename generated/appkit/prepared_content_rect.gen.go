
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [preparedContentRect] class.
var preparedContentRectClass _preparedContentRectClass

func init() {
	preparedContentRectClass = _preparedContentRectClass{objc.GetClass("preparedContentRect")}
}

type _preparedContentRectClass struct {
	objc.Class
}

// An interface definition for the [preparedContentRect] class.
type IpreparedContentRect interface {
	ID() objc.ID
}

type preparedContentRect struct {
	id objc.ID
}

func preparedContentRectFrom(ptr unsafe.Pointer) preparedContentRect {
	return preparedContentRect{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ preparedContentRect) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _preparedContentRectClass) Alloc() preparedContentRect {
	rv := objc.Send[preparedContentRect](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _preparedContentRectClass) New() preparedContentRect {
	rv := objc.Send[preparedContentRect](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpreparedContentRect creates and returns a new initialized instance.
func NewpreparedContentRect() preparedContentRect {
	return preparedContentRectClass.New()
}

// Init initializes the instance.
func (p_ preparedContentRect) Init() preparedContentRect {
	rv := objc.Send[preparedContentRect](p_.ID(), selInit)
	return rv
}

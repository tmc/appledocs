
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [contentRect] class.
var contentRectClass _contentRectClass

func init() {
	contentRectClass = _contentRectClass{objc.GetClass("contentRect")}
}

type _contentRectClass struct {
	objc.Class
}

// An interface definition for the [contentRect] class.
type IcontentRect interface {
	ID() objc.ID
}

type contentRect struct {
	id objc.ID
}

func contentRectFrom(ptr unsafe.Pointer) contentRect {
	return contentRect{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ contentRect) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _contentRectClass) Alloc() contentRect {
	rv := objc.Send[contentRect](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _contentRectClass) New() contentRect {
	rv := objc.Send[contentRect](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontentRect creates and returns a new initialized instance.
func NewcontentRect() contentRect {
	return contentRectClass.New()
}

// Init initializes the instance.
func (c_ contentRect) Init() contentRect {
	rv := objc.Send[contentRect](c_.ID(), selInit)
	return rv
}

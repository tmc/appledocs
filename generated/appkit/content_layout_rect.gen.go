
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [contentLayoutRect] class.
var contentLayoutRectClass _contentLayoutRectClass

func init() {
	contentLayoutRectClass = _contentLayoutRectClass{objc.GetClass("contentLayoutRect")}
}

type _contentLayoutRectClass struct {
	objc.Class
}

// An interface definition for the [contentLayoutRect] class.
type IcontentLayoutRect interface {
	ID() objc.ID
}

type contentLayoutRect struct {
	id objc.ID
}

func contentLayoutRectFrom(ptr unsafe.Pointer) contentLayoutRect {
	return contentLayoutRect{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ contentLayoutRect) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _contentLayoutRectClass) Alloc() contentLayoutRect {
	rv := objc.Send[contentLayoutRect](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _contentLayoutRectClass) New() contentLayoutRect {
	rv := objc.Send[contentLayoutRect](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontentLayoutRect creates and returns a new initialized instance.
func NewcontentLayoutRect() contentLayoutRect {
	return contentLayoutRectClass.New()
}

// Init initializes the instance.
func (c_ contentLayoutRect) Init() contentLayoutRect {
	rv := objc.Send[contentLayoutRect](c_.ID(), selInit)
	return rv
}

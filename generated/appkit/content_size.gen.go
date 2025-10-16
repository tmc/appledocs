
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [contentSize] class.
var contentSizeClass _contentSizeClass

func init() {
	contentSizeClass = _contentSizeClass{objc.GetClass("contentSize")}
}

type _contentSizeClass struct {
	objc.Class
}

// An interface definition for the [contentSize] class.
type IcontentSize interface {
	ID() objc.ID
}

type contentSize struct {
	id objc.ID
}

func contentSizeFrom(ptr unsafe.Pointer) contentSize {
	return contentSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ contentSize) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _contentSizeClass) Alloc() contentSize {
	rv := objc.Send[contentSize](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _contentSizeClass) New() contentSize {
	rv := objc.Send[contentSize](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontentSize creates and returns a new initialized instance.
func NewcontentSize() contentSize {
	return contentSizeClass.New()
}

// Init initializes the instance.
func (c_ contentSize) Init() contentSize {
	rv := objc.Send[contentSize](c_.ID(), selInit)
	return rv
}

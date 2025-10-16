
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [contentMaxSize] class.
var contentMaxSizeClass _contentMaxSizeClass

func init() {
	contentMaxSizeClass = _contentMaxSizeClass{objc.GetClass("contentMaxSize")}
}

type _contentMaxSizeClass struct {
	objc.Class
}

// An interface definition for the [contentMaxSize] class.
type IcontentMaxSize interface {
	ID() objc.ID
}

type contentMaxSize struct {
	id objc.ID
}

func contentMaxSizeFrom(ptr unsafe.Pointer) contentMaxSize {
	return contentMaxSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ contentMaxSize) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _contentMaxSizeClass) Alloc() contentMaxSize {
	rv := objc.Send[contentMaxSize](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _contentMaxSizeClass) New() contentMaxSize {
	rv := objc.Send[contentMaxSize](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontentMaxSize creates and returns a new initialized instance.
func NewcontentMaxSize() contentMaxSize {
	return contentMaxSizeClass.New()
}

// Init initializes the instance.
func (c_ contentMaxSize) Init() contentMaxSize {
	rv := objc.Send[contentMaxSize](c_.ID(), selInit)
	return rv
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [contentAspectRatio] class.
var contentAspectRatioClass _contentAspectRatioClass

func init() {
	contentAspectRatioClass = _contentAspectRatioClass{objc.GetClass("contentAspectRatio")}
}

type _contentAspectRatioClass struct {
	objc.Class
}

// An interface definition for the [contentAspectRatio] class.
type IcontentAspectRatio interface {
	ID() objc.ID
}

type contentAspectRatio struct {
	id objc.ID
}

func contentAspectRatioFrom(ptr unsafe.Pointer) contentAspectRatio {
	return contentAspectRatio{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ contentAspectRatio) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _contentAspectRatioClass) Alloc() contentAspectRatio {
	rv := objc.Send[contentAspectRatio](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _contentAspectRatioClass) New() contentAspectRatio {
	rv := objc.Send[contentAspectRatio](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontentAspectRatio creates and returns a new initialized instance.
func NewcontentAspectRatio() contentAspectRatio {
	return contentAspectRatioClass.New()
}

// Init initializes the instance.
func (c_ contentAspectRatio) Init() contentAspectRatio {
	rv := objc.Send[contentAspectRatio](c_.ID(), selInit)
	return rv
}

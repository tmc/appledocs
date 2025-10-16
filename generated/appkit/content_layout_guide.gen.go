
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [contentLayoutGuide] class.
var contentLayoutGuideClass _contentLayoutGuideClass

func init() {
	contentLayoutGuideClass = _contentLayoutGuideClass{objc.GetClass("contentLayoutGuide")}
}

type _contentLayoutGuideClass struct {
	objc.Class
}

// An interface definition for the [contentLayoutGuide] class.
type IcontentLayoutGuide interface {
	ID() objc.ID
}

type contentLayoutGuide struct {
	id objc.ID
}

func contentLayoutGuideFrom(ptr unsafe.Pointer) contentLayoutGuide {
	return contentLayoutGuide{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ contentLayoutGuide) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _contentLayoutGuideClass) Alloc() contentLayoutGuide {
	rv := objc.Send[contentLayoutGuide](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _contentLayoutGuideClass) New() contentLayoutGuide {
	rv := objc.Send[contentLayoutGuide](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontentLayoutGuide creates and returns a new initialized instance.
func NewcontentLayoutGuide() contentLayoutGuide {
	return contentLayoutGuideClass.New()
}

// Init initializes the instance.
func (c_ contentLayoutGuide) Init() contentLayoutGuide {
	rv := objc.Send[contentLayoutGuide](c_.ID(), selInit)
	return rv
}

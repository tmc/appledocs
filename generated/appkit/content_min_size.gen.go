
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [contentMinSize] class.
var contentMinSizeClass _contentMinSizeClass

func init() {
	contentMinSizeClass = _contentMinSizeClass{objc.GetClass("contentMinSize")}
}

type _contentMinSizeClass struct {
	objc.Class
}

// An interface definition for the [contentMinSize] class.
type IcontentMinSize interface {
	ID() objc.ID
}

type contentMinSize struct {
	id objc.ID
}

func contentMinSizeFrom(ptr unsafe.Pointer) contentMinSize {
	return contentMinSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ contentMinSize) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _contentMinSizeClass) Alloc() contentMinSize {
	rv := objc.Send[contentMinSize](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _contentMinSizeClass) New() contentMinSize {
	rv := objc.Send[contentMinSize](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontentMinSize creates and returns a new initialized instance.
func NewcontentMinSize() contentMinSize {
	return contentMinSizeClass.New()
}

// Init initializes the instance.
func (c_ contentMinSize) Init() contentMinSize {
	rv := objc.Send[contentMinSize](c_.ID(), selInit)
	return rv
}

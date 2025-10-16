
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [contentInsets] class.
var contentInsetsClass _contentInsetsClass

func init() {
	contentInsetsClass = _contentInsetsClass{objc.GetClass("contentInsets")}
}

type _contentInsetsClass struct {
	objc.Class
}

// An interface definition for the [contentInsets] class.
type IcontentInsets interface {
	ID() objc.ID
}

type contentInsets struct {
	id objc.ID
}

func contentInsetsFrom(ptr unsafe.Pointer) contentInsets {
	return contentInsets{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ contentInsets) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _contentInsetsClass) Alloc() contentInsets {
	rv := objc.Send[contentInsets](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _contentInsetsClass) New() contentInsets {
	rv := objc.Send[contentInsets](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontentInsets creates and returns a new initialized instance.
func NewcontentInsets() contentInsets {
	return contentInsetsClass.New()
}

// Init initializes the instance.
func (c_ contentInsets) Init() contentInsets {
	rv := objc.Send[contentInsets](c_.ID(), selInit)
	return rv
}

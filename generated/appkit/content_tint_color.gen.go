
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [contentTintColor] class.
var contentTintColorClass _contentTintColorClass

func init() {
	contentTintColorClass = _contentTintColorClass{objc.GetClass("contentTintColor")}
}

type _contentTintColorClass struct {
	objc.Class
}

// An interface definition for the [contentTintColor] class.
type IcontentTintColor interface {
	ID() objc.ID
}

type contentTintColor struct {
	id objc.ID
}

func contentTintColorFrom(ptr unsafe.Pointer) contentTintColor {
	return contentTintColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ contentTintColor) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _contentTintColorClass) Alloc() contentTintColor {
	rv := objc.Send[contentTintColor](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _contentTintColorClass) New() contentTintColor {
	rv := objc.Send[contentTintColor](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontentTintColor creates and returns a new initialized instance.
func NewcontentTintColor() contentTintColor {
	return contentTintColorClass.New()
}

// Init initializes the instance.
func (c_ contentTintColor) Init() contentTintColor {
	rv := objc.Send[contentTintColor](c_.ID(), selInit)
	return rv
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [clipsToBounds] class.
var clipsToBoundsClass _clipsToBoundsClass

func init() {
	clipsToBoundsClass = _clipsToBoundsClass{objc.GetClass("clipsToBounds")}
}

type _clipsToBoundsClass struct {
	objc.Class
}

// An interface definition for the [clipsToBounds] class.
type IclipsToBounds interface {
	ID() objc.ID
}

type clipsToBounds struct {
	id objc.ID
}

func clipsToBoundsFrom(ptr unsafe.Pointer) clipsToBounds {
	return clipsToBounds{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ clipsToBounds) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _clipsToBoundsClass) Alloc() clipsToBounds {
	rv := objc.Send[clipsToBounds](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _clipsToBoundsClass) New() clipsToBounds {
	rv := objc.Send[clipsToBounds](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewclipsToBounds creates and returns a new initialized instance.
func NewclipsToBounds() clipsToBounds {
	return clipsToBoundsClass.New()
}

// Init initializes the instance.
func (c_ clipsToBounds) Init() clipsToBounds {
	rv := objc.Send[clipsToBounds](c_.ID(), selInit)
	return rv
}

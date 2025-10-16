
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [screensHaveSeparateSpaces] class.
var screensHaveSeparateSpacesClass _screensHaveSeparateSpacesClass

func init() {
	screensHaveSeparateSpacesClass = _screensHaveSeparateSpacesClass{objc.GetClass("screensHaveSeparateSpaces")}
}

type _screensHaveSeparateSpacesClass struct {
	objc.Class
}

// An interface definition for the [screensHaveSeparateSpaces] class.
type IscreensHaveSeparateSpaces interface {
	ID() objc.ID
}

type screensHaveSeparateSpaces struct {
	id objc.ID
}

func screensHaveSeparateSpacesFrom(ptr unsafe.Pointer) screensHaveSeparateSpaces {
	return screensHaveSeparateSpaces{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ screensHaveSeparateSpaces) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _screensHaveSeparateSpacesClass) Alloc() screensHaveSeparateSpaces {
	rv := objc.Send[screensHaveSeparateSpaces](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _screensHaveSeparateSpacesClass) New() screensHaveSeparateSpaces {
	rv := objc.Send[screensHaveSeparateSpaces](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewscreensHaveSeparateSpaces creates and returns a new initialized instance.
func NewscreensHaveSeparateSpaces() screensHaveSeparateSpaces {
	return screensHaveSeparateSpacesClass.New()
}

// Init initializes the instance.
func (s_ screensHaveSeparateSpaces) Init() screensHaveSeparateSpaces {
	rv := objc.Send[screensHaveSeparateSpaces](s_.ID(), selInit)
	return rv
}

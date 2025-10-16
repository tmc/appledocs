
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [styleMask] class.
var styleMaskClass _styleMaskClass

func init() {
	styleMaskClass = _styleMaskClass{objc.GetClass("styleMask")}
}

type _styleMaskClass struct {
	objc.Class
}

// An interface definition for the [styleMask] class.
type IstyleMask interface {
	ID() objc.ID
}

type styleMask struct {
	id objc.ID
}

func styleMaskFrom(ptr unsafe.Pointer) styleMask {
	return styleMask{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ styleMask) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _styleMaskClass) Alloc() styleMask {
	rv := objc.Send[styleMask](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _styleMaskClass) New() styleMask {
	rv := objc.Send[styleMask](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewstyleMask creates and returns a new initialized instance.
func NewstyleMask() styleMask {
	return styleMaskClass.New()
}

// Init initializes the instance.
func (s_ styleMask) Init() styleMask {
	rv := objc.Send[styleMask](s_.ID(), selInit)
	return rv
}

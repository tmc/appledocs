
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [mode] class.
var modeClass _modeClass

func init() {
	modeClass = _modeClass{objc.GetClass("mode")}
}

type _modeClass struct {
	objc.Class
}

// An interface definition for the [mode] class.
type Imode interface {
	ID() objc.ID
}

type mode struct {
	id objc.ID
}

func modeFrom(ptr unsafe.Pointer) mode {
	return mode{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ mode) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _modeClass) Alloc() mode {
	rv := objc.Send[mode](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _modeClass) New() mode {
	rv := objc.Send[mode](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newmode creates and returns a new initialized instance.
func Newmode() mode {
	return modeClass.New()
}

// Init initializes the instance.
func (m_ mode) Init() mode {
	rv := objc.Send[mode](m_.ID(), selInit)
	return rv
}

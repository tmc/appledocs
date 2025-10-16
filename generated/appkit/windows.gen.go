
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [windows] class.
var windowsClass _windowsClass

func init() {
	windowsClass = _windowsClass{objc.GetClass("windows")}
}

type _windowsClass struct {
	objc.Class
}

// An interface definition for the [windows] class.
type Iwindows interface {
	ID() objc.ID
}

type windows struct {
	id objc.ID
}

func windowsFrom(ptr unsafe.Pointer) windows {
	return windows{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ windows) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _windowsClass) Alloc() windows {
	rv := objc.Send[windows](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _windowsClass) New() windows {
	rv := objc.Send[windows](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newwindows creates and returns a new initialized instance.
func Newwindows() windows {
	return windowsClass.New()
}

// Init initializes the instance.
func (w_ windows) Init() windows {
	rv := objc.Send[windows](w_.ID(), selInit)
	return rv
}

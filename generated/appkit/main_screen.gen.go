
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [mainScreen] class.
var mainScreenClass _mainScreenClass

func init() {
	mainScreenClass = _mainScreenClass{objc.GetClass("mainScreen")}
}

type _mainScreenClass struct {
	objc.Class
}

// An interface definition for the [mainScreen] class.
type ImainScreen interface {
	ID() objc.ID
}

type mainScreen struct {
	id objc.ID
}

func mainScreenFrom(ptr unsafe.Pointer) mainScreen {
	return mainScreen{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ mainScreen) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _mainScreenClass) Alloc() mainScreen {
	rv := objc.Send[mainScreen](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _mainScreenClass) New() mainScreen {
	rv := objc.Send[mainScreen](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmainScreen creates and returns a new initialized instance.
func NewmainScreen() mainScreen {
	return mainScreenClass.New()
}

// Init initializes the instance.
func (m_ mainScreen) Init() mainScreen {
	rv := objc.Send[mainScreen](m_.ID(), selInit)
	return rv
}

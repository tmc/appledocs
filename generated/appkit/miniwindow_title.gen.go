
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [miniwindowTitle] class.
var miniwindowTitleClass _miniwindowTitleClass

func init() {
	miniwindowTitleClass = _miniwindowTitleClass{objc.GetClass("miniwindowTitle")}
}

type _miniwindowTitleClass struct {
	objc.Class
}

// An interface definition for the [miniwindowTitle] class.
type IminiwindowTitle interface {
	ID() objc.ID
}

type miniwindowTitle struct {
	id objc.ID
}

func miniwindowTitleFrom(ptr unsafe.Pointer) miniwindowTitle {
	return miniwindowTitle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ miniwindowTitle) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _miniwindowTitleClass) Alloc() miniwindowTitle {
	rv := objc.Send[miniwindowTitle](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _miniwindowTitleClass) New() miniwindowTitle {
	rv := objc.Send[miniwindowTitle](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewminiwindowTitle creates and returns a new initialized instance.
func NewminiwindowTitle() miniwindowTitle {
	return miniwindowTitleClass.New()
}

// Init initializes the instance.
func (m_ miniwindowTitle) Init() miniwindowTitle {
	rv := objc.Send[miniwindowTitle](m_.ID(), selInit)
	return rv
}

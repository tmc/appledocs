
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isInFullScreenMode] class.
var isInFullScreenModeClass _isInFullScreenModeClass

func init() {
	isInFullScreenModeClass = _isInFullScreenModeClass{objc.GetClass("isInFullScreenMode")}
}

type _isInFullScreenModeClass struct {
	objc.Class
}

// An interface definition for the [isInFullScreenMode] class.
type IisInFullScreenMode interface {
	ID() objc.ID
}

type isInFullScreenMode struct {
	id objc.ID
}

func isInFullScreenModeFrom(ptr unsafe.Pointer) isInFullScreenMode {
	return isInFullScreenMode{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isInFullScreenMode) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isInFullScreenModeClass) Alloc() isInFullScreenMode {
	rv := objc.Send[isInFullScreenMode](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isInFullScreenModeClass) New() isInFullScreenMode {
	rv := objc.Send[isInFullScreenMode](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisInFullScreenMode creates and returns a new initialized instance.
func NewisInFullScreenMode() isInFullScreenMode {
	return isInFullScreenModeClass.New()
}

// Init initializes the instance.
func (i_ isInFullScreenMode) Init() isInFullScreenMode {
	rv := objc.Send[isInFullScreenMode](i_.ID(), selInit)
	return rv
}

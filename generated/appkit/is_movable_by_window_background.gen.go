
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isMovableByWindowBackground] class.
var isMovableByWindowBackgroundClass _isMovableByWindowBackgroundClass

func init() {
	isMovableByWindowBackgroundClass = _isMovableByWindowBackgroundClass{objc.GetClass("isMovableByWindowBackground")}
}

type _isMovableByWindowBackgroundClass struct {
	objc.Class
}

// An interface definition for the [isMovableByWindowBackground] class.
type IisMovableByWindowBackground interface {
	ID() objc.ID
}

type isMovableByWindowBackground struct {
	id objc.ID
}

func isMovableByWindowBackgroundFrom(ptr unsafe.Pointer) isMovableByWindowBackground {
	return isMovableByWindowBackground{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isMovableByWindowBackground) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isMovableByWindowBackgroundClass) Alloc() isMovableByWindowBackground {
	rv := objc.Send[isMovableByWindowBackground](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isMovableByWindowBackgroundClass) New() isMovableByWindowBackground {
	rv := objc.Send[isMovableByWindowBackground](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisMovableByWindowBackground creates and returns a new initialized instance.
func NewisMovableByWindowBackground() isMovableByWindowBackground {
	return isMovableByWindowBackgroundClass.New()
}

// Init initializes the instance.
func (i_ isMovableByWindowBackground) Init() isMovableByWindowBackground {
	rv := objc.Send[isMovableByWindowBackground](i_.ID(), selInit)
	return rv
}

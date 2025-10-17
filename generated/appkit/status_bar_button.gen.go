
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [StatusBarButton] class.
var StatusBarButtonClass _StatusBarButtonClass

func init() {
	StatusBarButtonClass = _StatusBarButtonClass{objc.GetClass("NSStatusBarButton")}
}

type _StatusBarButtonClass struct {
	objc.Class
}

// An interface definition for the [StatusBarButton] class.
type IStatusBarButton interface {
	ID() objc.ID
}

type StatusBarButton struct {
	id objc.ID
}

func StatusBarButtonFrom(ptr unsafe.Pointer) StatusBarButton {
	return StatusBarButton{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ StatusBarButton) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _StatusBarButtonClass) Alloc() StatusBarButton {
	rv := objc.Send[StatusBarButton](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _StatusBarButtonClass) New() StatusBarButton {
	rv := objc.Send[StatusBarButton](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewStatusBarButton creates and returns a new initialized instance.
func NewStatusBarButton() StatusBarButton {
	return StatusBarButtonClass.New()
}

// Init initializes the instance.
func (s_ StatusBarButton) Init() StatusBarButton {
	rv := objc.Send[StatusBarButton](s_.ID(), selInit)
	return rv
}

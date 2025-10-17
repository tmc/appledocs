
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [StatusBar] class.
var StatusBarClass _StatusBarClass

func init() {
	StatusBarClass = _StatusBarClass{objc.GetClass("NSStatusBar")}
}

type _StatusBarClass struct {
	objc.Class
}

// An interface definition for the [StatusBar] class.
type IStatusBar interface {
	ID() objc.ID
}

type StatusBar struct {
	id objc.ID
}

func StatusBarFrom(ptr unsafe.Pointer) StatusBar {
	return StatusBar{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ StatusBar) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _StatusBarClass) Alloc() StatusBar {
	rv := objc.Send[StatusBar](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _StatusBarClass) New() StatusBar {
	rv := objc.Send[StatusBar](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewStatusBar creates and returns a new initialized instance.
func NewStatusBar() StatusBar {
	return StatusBarClass.New()
}

// Init initializes the instance.
func (s_ StatusBar) Init() StatusBar {
	rv := objc.Send[StatusBar](s_.ID(), selInit)
	return rv
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isMainWindow] class.
var isMainWindowClass _isMainWindowClass

func init() {
	isMainWindowClass = _isMainWindowClass{objc.GetClass("isMainWindow")}
}

type _isMainWindowClass struct {
	objc.Class
}

// An interface definition for the [isMainWindow] class.
type IisMainWindow interface {
	ID() objc.ID
}

type isMainWindow struct {
	id objc.ID
}

func isMainWindowFrom(ptr unsafe.Pointer) isMainWindow {
	return isMainWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isMainWindow) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isMainWindowClass) Alloc() isMainWindow {
	rv := objc.Send[isMainWindow](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isMainWindowClass) New() isMainWindow {
	rv := objc.Send[isMainWindow](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisMainWindow creates and returns a new initialized instance.
func NewisMainWindow() isMainWindow {
	return isMainWindowClass.New()
}

// Init initializes the instance.
func (i_ isMainWindow) Init() isMainWindow {
	rv := objc.Send[isMainWindow](i_.ID(), selInit)
	return rv
}

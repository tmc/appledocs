
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isKeyWindow] class.
var isKeyWindowClass _isKeyWindowClass

func init() {
	isKeyWindowClass = _isKeyWindowClass{objc.GetClass("isKeyWindow")}
}

type _isKeyWindowClass struct {
	objc.Class
}

// An interface definition for the [isKeyWindow] class.
type IisKeyWindow interface {
	ID() objc.ID
}

type isKeyWindow struct {
	id objc.ID
}

func isKeyWindowFrom(ptr unsafe.Pointer) isKeyWindow {
	return isKeyWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isKeyWindow) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isKeyWindowClass) Alloc() isKeyWindow {
	rv := objc.Send[isKeyWindow](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isKeyWindowClass) New() isKeyWindow {
	rv := objc.Send[isKeyWindow](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisKeyWindow creates and returns a new initialized instance.
func NewisKeyWindow() isKeyWindow {
	return isKeyWindowClass.New()
}

// Init initializes the instance.
func (i_ isKeyWindow) Init() isKeyWindow {
	rv := objc.Send[isKeyWindow](i_.ID(), selInit)
	return rv
}

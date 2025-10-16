
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [resignMainWindow] class.
var resignMainWindowClass _resignMainWindowClass

func init() {
	resignMainWindowClass = _resignMainWindowClass{objc.GetClass("resignMainWindow")}
}

type _resignMainWindowClass struct {
	objc.Class
}

// An interface definition for the [resignMainWindow] class.
type IresignMainWindow interface {
	ID() objc.ID
}

type resignMainWindow struct {
	id objc.ID
}

func resignMainWindowFrom(ptr unsafe.Pointer) resignMainWindow {
	return resignMainWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ resignMainWindow) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _resignMainWindowClass) Alloc() resignMainWindow {
	rv := objc.Send[resignMainWindow](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _resignMainWindowClass) New() resignMainWindow {
	rv := objc.Send[resignMainWindow](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewresignMainWindow creates and returns a new initialized instance.
func NewresignMainWindow() resignMainWindow {
	return resignMainWindowClass.New()
}

// Init initializes the instance.
func (r_ resignMainWindow) Init() resignMainWindow {
	rv := objc.Send[resignMainWindow](r_.ID(), selInit)
	return rv
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [parentWindow] class.
var parentWindowClass _parentWindowClass

func init() {
	parentWindowClass = _parentWindowClass{objc.GetClass("parentWindow")}
}

type _parentWindowClass struct {
	objc.Class
}

// An interface definition for the [parentWindow] class.
type IparentWindow interface {
	ID() objc.ID
}

type parentWindow struct {
	id objc.ID
}

func parentWindowFrom(ptr unsafe.Pointer) parentWindow {
	return parentWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ parentWindow) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _parentWindowClass) Alloc() parentWindow {
	rv := objc.Send[parentWindow](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _parentWindowClass) New() parentWindow {
	rv := objc.Send[parentWindow](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewparentWindow creates and returns a new initialized instance.
func NewparentWindow() parentWindow {
	return parentWindowClass.New()
}

// Init initializes the instance.
func (p_ parentWindow) Init() parentWindow {
	rv := objc.Send[parentWindow](p_.ID(), selInit)
	return rv
}

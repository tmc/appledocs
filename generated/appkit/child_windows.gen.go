
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [childWindows] class.
var childWindowsClass _childWindowsClass

func init() {
	childWindowsClass = _childWindowsClass{objc.GetClass("childWindows")}
}

type _childWindowsClass struct {
	objc.Class
}

// An interface definition for the [childWindows] class.
type IchildWindows interface {
	ID() objc.ID
}

type childWindows struct {
	id objc.ID
}

func childWindowsFrom(ptr unsafe.Pointer) childWindows {
	return childWindows{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ childWindows) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _childWindowsClass) Alloc() childWindows {
	rv := objc.Send[childWindows](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _childWindowsClass) New() childWindows {
	rv := objc.Send[childWindows](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewchildWindows creates and returns a new initialized instance.
func NewchildWindows() childWindows {
	return childWindowsClass.New()
}

// Init initializes the instance.
func (c_ childWindows) Init() childWindows {
	rv := objc.Send[childWindows](c_.ID(), selInit)
	return rv
}

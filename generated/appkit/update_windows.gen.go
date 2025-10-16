
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [updateWindows] class.
var updateWindowsClass _updateWindowsClass

func init() {
	updateWindowsClass = _updateWindowsClass{objc.GetClass("updateWindows")}
}

type _updateWindowsClass struct {
	objc.Class
}

// An interface definition for the [updateWindows] class.
type IupdateWindows interface {
	ID() objc.ID
}

type updateWindows struct {
	id objc.ID
}

func updateWindowsFrom(ptr unsafe.Pointer) updateWindows {
	return updateWindows{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ updateWindows) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _updateWindowsClass) Alloc() updateWindows {
	rv := objc.Send[updateWindows](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _updateWindowsClass) New() updateWindows {
	rv := objc.Send[updateWindows](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewupdateWindows creates and returns a new initialized instance.
func NewupdateWindows() updateWindows {
	return updateWindowsClass.New()
}

// Init initializes the instance.
func (u_ updateWindows) Init() updateWindows {
	rv := objc.Send[updateWindows](u_.ID(), selInit)
	return rv
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [enclosingMenuItem] class.
var enclosingMenuItemClass _enclosingMenuItemClass

func init() {
	enclosingMenuItemClass = _enclosingMenuItemClass{objc.GetClass("enclosingMenuItem")}
}

type _enclosingMenuItemClass struct {
	objc.Class
}

// An interface definition for the [enclosingMenuItem] class.
type IenclosingMenuItem interface {
	ID() objc.ID
}

type enclosingMenuItem struct {
	id objc.ID
}

func enclosingMenuItemFrom(ptr unsafe.Pointer) enclosingMenuItem {
	return enclosingMenuItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ enclosingMenuItem) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _enclosingMenuItemClass) Alloc() enclosingMenuItem {
	rv := objc.Send[enclosingMenuItem](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _enclosingMenuItemClass) New() enclosingMenuItem {
	rv := objc.Send[enclosingMenuItem](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewenclosingMenuItem creates and returns a new initialized instance.
func NewenclosingMenuItem() enclosingMenuItem {
	return enclosingMenuItemClass.New()
}

// Init initializes the instance.
func (e_ enclosingMenuItem) Init() enclosingMenuItem {
	rv := objc.Send[enclosingMenuItem](e_.ID(), selInit)
	return rv
}

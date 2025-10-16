
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsUndo] class.
var allowsUndoClass _allowsUndoClass

func init() {
	allowsUndoClass = _allowsUndoClass{objc.GetClass("allowsUndo")}
}

type _allowsUndoClass struct {
	objc.Class
}

// An interface definition for the [allowsUndo] class.
type IallowsUndo interface {
	ID() objc.ID
}

type allowsUndo struct {
	id objc.ID
}

func allowsUndoFrom(ptr unsafe.Pointer) allowsUndo {
	return allowsUndo{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsUndo) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsUndoClass) Alloc() allowsUndo {
	rv := objc.Send[allowsUndo](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsUndoClass) New() allowsUndo {
	rv := objc.Send[allowsUndo](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsUndo creates and returns a new initialized instance.
func NewallowsUndo() allowsUndo {
	return allowsUndoClass.New()
}

// Init initializes the instance.
func (a_ allowsUndo) Init() allowsUndo {
	rv := objc.Send[allowsUndo](a_.ID(), selInit)
	return rv
}

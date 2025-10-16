
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [undoManager] class.
var undoManagerClass _undoManagerClass

func init() {
	undoManagerClass = _undoManagerClass{objc.GetClass("undoManager")}
}

type _undoManagerClass struct {
	objc.Class
}

// An interface definition for the [undoManager] class.
type IundoManager interface {
	ID() objc.ID
}

type undoManager struct {
	id objc.ID
}

func undoManagerFrom(ptr unsafe.Pointer) undoManager {
	return undoManager{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ undoManager) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _undoManagerClass) Alloc() undoManager {
	rv := objc.Send[undoManager](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _undoManagerClass) New() undoManager {
	rv := objc.Send[undoManager](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewundoManager creates and returns a new initialized instance.
func NewundoManager() undoManager {
	return undoManagerClass.New()
}

// Init initializes the instance.
func (u_ undoManager) Init() undoManager {
	rv := objc.Send[undoManager](u_.ID(), selInit)
	return rv
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [attachedSheet] class.
var attachedSheetClass _attachedSheetClass

func init() {
	attachedSheetClass = _attachedSheetClass{objc.GetClass("attachedSheet")}
}

type _attachedSheetClass struct {
	objc.Class
}

// An interface definition for the [attachedSheet] class.
type IattachedSheet interface {
	ID() objc.ID
}

type attachedSheet struct {
	id objc.ID
}

func attachedSheetFrom(ptr unsafe.Pointer) attachedSheet {
	return attachedSheet{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ attachedSheet) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _attachedSheetClass) Alloc() attachedSheet {
	rv := objc.Send[attachedSheet](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _attachedSheetClass) New() attachedSheet {
	rv := objc.Send[attachedSheet](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewattachedSheet creates and returns a new initialized instance.
func NewattachedSheet() attachedSheet {
	return attachedSheetClass.New()
}

// Init initializes the instance.
func (a_ attachedSheet) Init() attachedSheet {
	rv := objc.Send[attachedSheet](a_.ID(), selInit)
	return rv
}

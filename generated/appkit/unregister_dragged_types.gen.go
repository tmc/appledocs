
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [unregisterDraggedTypes] class.
var unregisterDraggedTypesClass _unregisterDraggedTypesClass

func init() {
	unregisterDraggedTypesClass = _unregisterDraggedTypesClass{objc.GetClass("unregisterDraggedTypes")}
}

type _unregisterDraggedTypesClass struct {
	objc.Class
}

// An interface definition for the [unregisterDraggedTypes] class.
type IunregisterDraggedTypes interface {
	ID() objc.ID
}

type unregisterDraggedTypes struct {
	id objc.ID
}

func unregisterDraggedTypesFrom(ptr unsafe.Pointer) unregisterDraggedTypes {
	return unregisterDraggedTypes{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ unregisterDraggedTypes) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _unregisterDraggedTypesClass) Alloc() unregisterDraggedTypes {
	rv := objc.Send[unregisterDraggedTypes](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _unregisterDraggedTypesClass) New() unregisterDraggedTypes {
	rv := objc.Send[unregisterDraggedTypes](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewunregisterDraggedTypes creates and returns a new initialized instance.
func NewunregisterDraggedTypes() unregisterDraggedTypes {
	return unregisterDraggedTypesClass.New()
}

// Init initializes the instance.
func (u_ unregisterDraggedTypes) Init() unregisterDraggedTypes {
	rv := objc.Send[unregisterDraggedTypes](u_.ID(), selInit)
	return rv
}

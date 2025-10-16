
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [registeredDraggedTypes] class.
var registeredDraggedTypesClass _registeredDraggedTypesClass

func init() {
	registeredDraggedTypesClass = _registeredDraggedTypesClass{objc.GetClass("registeredDraggedTypes")}
}

type _registeredDraggedTypesClass struct {
	objc.Class
}

// An interface definition for the [registeredDraggedTypes] class.
type IregisteredDraggedTypes interface {
	ID() objc.ID
}

type registeredDraggedTypes struct {
	id objc.ID
}

func registeredDraggedTypesFrom(ptr unsafe.Pointer) registeredDraggedTypes {
	return registeredDraggedTypes{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ registeredDraggedTypes) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _registeredDraggedTypesClass) Alloc() registeredDraggedTypes {
	rv := objc.Send[registeredDraggedTypes](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _registeredDraggedTypesClass) New() registeredDraggedTypes {
	rv := objc.Send[registeredDraggedTypes](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewregisteredDraggedTypes creates and returns a new initialized instance.
func NewregisteredDraggedTypes() registeredDraggedTypes {
	return registeredDraggedTypesClass.New()
}

// Init initializes the instance.
func (r_ registeredDraggedTypes) Init() registeredDraggedTypes {
	rv := objc.Send[registeredDraggedTypes](r_.ID(), selInit)
	return rv
}

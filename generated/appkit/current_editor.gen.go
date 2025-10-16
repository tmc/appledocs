
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [currentEditor] class.
var currentEditorClass _currentEditorClass

func init() {
	currentEditorClass = _currentEditorClass{objc.GetClass("currentEditor")}
}

type _currentEditorClass struct {
	objc.Class
}

// An interface definition for the [currentEditor] class.
type IcurrentEditor interface {
	ID() objc.ID
}

type currentEditor struct {
	id objc.ID
}

func currentEditorFrom(ptr unsafe.Pointer) currentEditor {
	return currentEditor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ currentEditor) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _currentEditorClass) Alloc() currentEditor {
	rv := objc.Send[currentEditor](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _currentEditorClass) New() currentEditor {
	rv := objc.Send[currentEditor](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcurrentEditor creates and returns a new initialized instance.
func NewcurrentEditor() currentEditor {
	return currentEditorClass.New()
}

// Init initializes the instance.
func (c_ currentEditor) Init() currentEditor {
	rv := objc.Send[currentEditor](c_.ID(), selInit)
	return rv
}

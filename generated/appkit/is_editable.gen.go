
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isEditable] class.
var isEditableClass _isEditableClass

func init() {
	isEditableClass = _isEditableClass{objc.GetClass("isEditable")}
}

type _isEditableClass struct {
	objc.Class
}

// An interface definition for the [isEditable] class.
type IisEditable interface {
	ID() objc.ID
}

type isEditable struct {
	id objc.ID
}

func isEditableFrom(ptr unsafe.Pointer) isEditable {
	return isEditable{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isEditable) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isEditableClass) Alloc() isEditable {
	rv := objc.Send[isEditable](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isEditableClass) New() isEditable {
	rv := objc.Send[isEditable](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisEditable creates and returns a new initialized instance.
func NewisEditable() isEditable {
	return isEditableClass.New()
}

// Init initializes the instance.
func (i_ isEditable) Init() isEditable {
	rv := objc.Send[isEditable](i_.ID(), selInit)
	return rv
}

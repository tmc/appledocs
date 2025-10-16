
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isFieldEditor] class.
var isFieldEditorClass _isFieldEditorClass

func init() {
	isFieldEditorClass = _isFieldEditorClass{objc.GetClass("isFieldEditor")}
}

type _isFieldEditorClass struct {
	objc.Class
}

// An interface definition for the [isFieldEditor] class.
type IisFieldEditor interface {
	ID() objc.ID
}

type isFieldEditor struct {
	id objc.ID
}

func isFieldEditorFrom(ptr unsafe.Pointer) isFieldEditor {
	return isFieldEditor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isFieldEditor) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isFieldEditorClass) Alloc() isFieldEditor {
	rv := objc.Send[isFieldEditor](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isFieldEditorClass) New() isFieldEditor {
	rv := objc.Send[isFieldEditor](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisFieldEditor creates and returns a new initialized instance.
func NewisFieldEditor() isFieldEditor {
	return isFieldEditorClass.New()
}

// Init initializes the instance.
func (i_ isFieldEditor) Init() isFieldEditor {
	rv := objc.Send[isFieldEditor](i_.ID(), selInit)
	return rv
}

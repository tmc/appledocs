
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextField] class.
var TextFieldClass _TextFieldClass

func init() {
	TextFieldClass = _TextFieldClass{objc.GetClass("NSTextField")}
}

type _TextFieldClass struct {
	objc.Class
}

// An interface definition for the [TextField] class.
type ITextField interface {
	ID() objc.ID
}

type TextField struct {
	id objc.ID
}

func TextFieldFrom(ptr unsafe.Pointer) TextField {
	return TextField{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextField) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextFieldClass) Alloc() TextField {
	rv := objc.Send[TextField](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextFieldClass) New() TextField {
	rv := objc.Send[TextField](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextField creates and returns a new initialized instance.
func NewTextField() TextField {
	return TextFieldClass.New()
}

// Init initializes the instance.
func (t_ TextField) Init() TextField {
	rv := objc.Send[TextField](t_.ID(), selInit)
	return rv
}

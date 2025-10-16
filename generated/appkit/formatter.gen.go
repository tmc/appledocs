
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [formatter] class.
var formatterClass _formatterClass

func init() {
	formatterClass = _formatterClass{objc.GetClass("formatter")}
}

type _formatterClass struct {
	objc.Class
}

// An interface definition for the [formatter] class.
type Iformatter interface {
	ID() objc.ID
}

type formatter struct {
	id objc.ID
}

func formatterFrom(ptr unsafe.Pointer) formatter {
	return formatter{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ formatter) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _formatterClass) Alloc() formatter {
	rv := objc.Send[formatter](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _formatterClass) New() formatter {
	rv := objc.Send[formatter](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newformatter creates and returns a new initialized instance.
func Newformatter() formatter {
	return formatterClass.New()
}

// Init initializes the instance.
func (f_ formatter) Init() formatter {
	rv := objc.Send[formatter](f_.ID(), selInit)
	return rv
}

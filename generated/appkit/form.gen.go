
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Form] class.
var FormClass _FormClass

func init() {
	FormClass = _FormClass{objc.GetClass("NSForm")}
}

type _FormClass struct {
	objc.Class
}

// An interface definition for the [Form] class.
type IForm interface {
	ID() objc.ID
}

type Form struct {
	id objc.ID
}

func FormFrom(ptr unsafe.Pointer) Form {
	return Form{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ Form) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _FormClass) Alloc() Form {
	rv := objc.Send[Form](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _FormClass) New() Form {
	rv := objc.Send[Form](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewForm creates and returns a new initialized instance.
func NewForm() Form {
	return FormClass.New()
}

// Init initializes the instance.
func (f_ Form) Init() Form {
	rv := objc.Send[Form](f_.ID(), selInit)
	return rv
}

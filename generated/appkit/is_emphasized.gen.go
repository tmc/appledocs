
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isEmphasized] class.
var isEmphasizedClass _isEmphasizedClass

func init() {
	isEmphasizedClass = _isEmphasizedClass{objc.GetClass("isEmphasized")}
}

type _isEmphasizedClass struct {
	objc.Class
}

// An interface definition for the [isEmphasized] class.
type IisEmphasized interface {
	ID() objc.ID
}

type isEmphasized struct {
	id objc.ID
}

func isEmphasizedFrom(ptr unsafe.Pointer) isEmphasized {
	return isEmphasized{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isEmphasized) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isEmphasizedClass) Alloc() isEmphasized {
	rv := objc.Send[isEmphasized](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isEmphasizedClass) New() isEmphasized {
	rv := objc.Send[isEmphasized](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisEmphasized creates and returns a new initialized instance.
func NewisEmphasized() isEmphasized {
	return isEmphasizedClass.New()
}

// Init initializes the instance.
func (i_ isEmphasized) Init() isEmphasized {
	rv := objc.Send[isEmphasized](i_.ID(), selInit)
	return rv
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isRestorable] class.
var isRestorableClass _isRestorableClass

func init() {
	isRestorableClass = _isRestorableClass{objc.GetClass("isRestorable")}
}

type _isRestorableClass struct {
	objc.Class
}

// An interface definition for the [isRestorable] class.
type IisRestorable interface {
	ID() objc.ID
}

type isRestorable struct {
	id objc.ID
}

func isRestorableFrom(ptr unsafe.Pointer) isRestorable {
	return isRestorable{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isRestorable) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isRestorableClass) Alloc() isRestorable {
	rv := objc.Send[isRestorable](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isRestorableClass) New() isRestorable {
	rv := objc.Send[isRestorable](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisRestorable creates and returns a new initialized instance.
func NewisRestorable() isRestorable {
	return isRestorableClass.New()
}

// Init initializes the instance.
func (i_ isRestorable) Init() isRestorable {
	rv := objc.Send[isRestorable](i_.ID(), selInit)
	return rv
}

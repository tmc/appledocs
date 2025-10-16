
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isEnabled] class.
var isEnabledClass _isEnabledClass

func init() {
	isEnabledClass = _isEnabledClass{objc.GetClass("isEnabled")}
}

type _isEnabledClass struct {
	objc.Class
}

// An interface definition for the [isEnabled] class.
type IisEnabled interface {
	ID() objc.ID
}

type isEnabled struct {
	id objc.ID
}

func isEnabledFrom(ptr unsafe.Pointer) isEnabled {
	return isEnabled{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isEnabled) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isEnabledClass) Alloc() isEnabled {
	rv := objc.Send[isEnabled](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isEnabledClass) New() isEnabled {
	rv := objc.Send[isEnabled](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisEnabled creates and returns a new initialized instance.
func NewisEnabled() isEnabled {
	return isEnabledClass.New()
}

// Init initializes the instance.
func (i_ isEnabled) Init() isEnabled {
	rv := objc.Send[isEnabled](i_.ID(), selInit)
	return rv
}

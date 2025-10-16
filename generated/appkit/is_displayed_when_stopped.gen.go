
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isDisplayedWhenStopped] class.
var isDisplayedWhenStoppedClass _isDisplayedWhenStoppedClass

func init() {
	isDisplayedWhenStoppedClass = _isDisplayedWhenStoppedClass{objc.GetClass("isDisplayedWhenStopped")}
}

type _isDisplayedWhenStoppedClass struct {
	objc.Class
}

// An interface definition for the [isDisplayedWhenStopped] class.
type IisDisplayedWhenStopped interface {
	ID() objc.ID
}

type isDisplayedWhenStopped struct {
	id objc.ID
}

func isDisplayedWhenStoppedFrom(ptr unsafe.Pointer) isDisplayedWhenStopped {
	return isDisplayedWhenStopped{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isDisplayedWhenStopped) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isDisplayedWhenStoppedClass) Alloc() isDisplayedWhenStopped {
	rv := objc.Send[isDisplayedWhenStopped](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isDisplayedWhenStoppedClass) New() isDisplayedWhenStopped {
	rv := objc.Send[isDisplayedWhenStopped](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisDisplayedWhenStopped creates and returns a new initialized instance.
func NewisDisplayedWhenStopped() isDisplayedWhenStopped {
	return isDisplayedWhenStoppedClass.New()
}

// Init initializes the instance.
func (i_ isDisplayedWhenStopped) Init() isDisplayedWhenStopped {
	rv := objc.Send[isDisplayedWhenStopped](i_.ID(), selInit)
	return rv
}

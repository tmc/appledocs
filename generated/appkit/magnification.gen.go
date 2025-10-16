
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [magnification] class.
var magnificationClass _magnificationClass

func init() {
	magnificationClass = _magnificationClass{objc.GetClass("magnification")}
}

type _magnificationClass struct {
	objc.Class
}

// An interface definition for the [magnification] class.
type Imagnification interface {
	ID() objc.ID
}

type magnification struct {
	id objc.ID
}

func magnificationFrom(ptr unsafe.Pointer) magnification {
	return magnification{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ magnification) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _magnificationClass) Alloc() magnification {
	rv := objc.Send[magnification](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _magnificationClass) New() magnification {
	rv := objc.Send[magnification](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newmagnification creates and returns a new initialized instance.
func Newmagnification() magnification {
	return magnificationClass.New()
}

// Init initializes the instance.
func (m_ magnification) Init() magnification {
	rv := objc.Send[magnification](m_.ID(), selInit)
	return rv
}

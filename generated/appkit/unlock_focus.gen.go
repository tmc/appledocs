
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [unlockFocus] class.
var unlockFocusClass _unlockFocusClass

func init() {
	unlockFocusClass = _unlockFocusClass{objc.GetClass("unlockFocus")}
}

type _unlockFocusClass struct {
	objc.Class
}

// An interface definition for the [unlockFocus] class.
type IunlockFocus interface {
	ID() objc.ID
}

type unlockFocus struct {
	id objc.ID
}

func unlockFocusFrom(ptr unsafe.Pointer) unlockFocus {
	return unlockFocus{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ unlockFocus) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _unlockFocusClass) Alloc() unlockFocus {
	rv := objc.Send[unlockFocus](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _unlockFocusClass) New() unlockFocus {
	rv := objc.Send[unlockFocus](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewunlockFocus creates and returns a new initialized instance.
func NewunlockFocus() unlockFocus {
	return unlockFocusClass.New()
}

// Init initializes the instance.
func (u_ unlockFocus) Init() unlockFocus {
	rv := objc.Send[unlockFocus](u_.ID(), selInit)
	return rv
}

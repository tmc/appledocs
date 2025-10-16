
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isAutodisplay] class.
var isAutodisplayClass _isAutodisplayClass

func init() {
	isAutodisplayClass = _isAutodisplayClass{objc.GetClass("isAutodisplay")}
}

type _isAutodisplayClass struct {
	objc.Class
}

// An interface definition for the [isAutodisplay] class.
type IisAutodisplay interface {
	ID() objc.ID
}

type isAutodisplay struct {
	id objc.ID
}

func isAutodisplayFrom(ptr unsafe.Pointer) isAutodisplay {
	return isAutodisplay{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isAutodisplay) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isAutodisplayClass) Alloc() isAutodisplay {
	rv := objc.Send[isAutodisplay](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isAutodisplayClass) New() isAutodisplay {
	rv := objc.Send[isAutodisplay](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisAutodisplay creates and returns a new initialized instance.
func NewisAutodisplay() isAutodisplay {
	return isAutodisplayClass.New()
}

// Init initializes the instance.
func (i_ isAutodisplay) Init() isAutodisplay {
	rv := objc.Send[isAutodisplay](i_.ID(), selInit)
	return rv
}

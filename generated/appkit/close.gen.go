
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [close] class.
var closeClass _closeClass

func init() {
	closeClass = _closeClass{objc.GetClass("close")}
}

type _closeClass struct {
	objc.Class
}

// An interface definition for the [close] class.
type Iclose interface {
	ID() objc.ID
}

type close struct {
	id objc.ID
}

func closeFrom(ptr unsafe.Pointer) close {
	return close{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ close) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _closeClass) Alloc() close {
	rv := objc.Send[close](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _closeClass) New() close {
	rv := objc.Send[close](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newclose creates and returns a new initialized instance.
func Newclose() close {
	return closeClass.New()
}

// Init initializes the instance.
func (c_ close) Init() close {
	rv := objc.Send[close](c_.ID(), selInit)
	return rv
}

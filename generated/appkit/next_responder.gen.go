
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [nextResponder] class.
var nextResponderClass _nextResponderClass

func init() {
	nextResponderClass = _nextResponderClass{objc.GetClass("nextResponder")}
}

type _nextResponderClass struct {
	objc.Class
}

// An interface definition for the [nextResponder] class.
type InextResponder interface {
	ID() objc.ID
}

type nextResponder struct {
	id objc.ID
}

func nextResponderFrom(ptr unsafe.Pointer) nextResponder {
	return nextResponder{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ nextResponder) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _nextResponderClass) Alloc() nextResponder {
	rv := objc.Send[nextResponder](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _nextResponderClass) New() nextResponder {
	rv := objc.Send[nextResponder](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewnextResponder creates and returns a new initialized instance.
func NewnextResponder() nextResponder {
	return nextResponderClass.New()
}

// Init initializes the instance.
func (n_ nextResponder) Init() nextResponder {
	rv := objc.Send[nextResponder](n_.ID(), selInit)
	return rv
}

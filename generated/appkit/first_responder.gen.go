
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [firstResponder] class.
var firstResponderClass _firstResponderClass

func init() {
	firstResponderClass = _firstResponderClass{objc.GetClass("firstResponder")}
}

type _firstResponderClass struct {
	objc.Class
}

// An interface definition for the [firstResponder] class.
type IfirstResponder interface {
	ID() objc.ID
}

type firstResponder struct {
	id objc.ID
}

func firstResponderFrom(ptr unsafe.Pointer) firstResponder {
	return firstResponder{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ firstResponder) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _firstResponderClass) Alloc() firstResponder {
	rv := objc.Send[firstResponder](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _firstResponderClass) New() firstResponder {
	rv := objc.Send[firstResponder](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewfirstResponder creates and returns a new initialized instance.
func NewfirstResponder() firstResponder {
	return firstResponderClass.New()
}

// Init initializes the instance.
func (f_ firstResponder) Init() firstResponder {
	rv := objc.Send[firstResponder](f_.ID(), selInit)
	return rv
}

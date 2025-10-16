
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [becomeFirstResponder] class.
var becomeFirstResponderClass _becomeFirstResponderClass

func init() {
	becomeFirstResponderClass = _becomeFirstResponderClass{objc.GetClass("becomeFirstResponder")}
}

type _becomeFirstResponderClass struct {
	objc.Class
}

// An interface definition for the [becomeFirstResponder] class.
type IbecomeFirstResponder interface {
	ID() objc.ID
}

type becomeFirstResponder struct {
	id objc.ID
}

func becomeFirstResponderFrom(ptr unsafe.Pointer) becomeFirstResponder {
	return becomeFirstResponder{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ becomeFirstResponder) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _becomeFirstResponderClass) Alloc() becomeFirstResponder {
	rv := objc.Send[becomeFirstResponder](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _becomeFirstResponderClass) New() becomeFirstResponder {
	rv := objc.Send[becomeFirstResponder](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbecomeFirstResponder creates and returns a new initialized instance.
func NewbecomeFirstResponder() becomeFirstResponder {
	return becomeFirstResponderClass.New()
}

// Init initializes the instance.
func (b_ becomeFirstResponder) Init() becomeFirstResponder {
	rv := objc.Send[becomeFirstResponder](b_.ID(), selInit)
	return rv
}

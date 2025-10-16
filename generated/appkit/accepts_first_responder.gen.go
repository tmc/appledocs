
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [acceptsFirstResponder] class.
var acceptsFirstResponderClass _acceptsFirstResponderClass

func init() {
	acceptsFirstResponderClass = _acceptsFirstResponderClass{objc.GetClass("acceptsFirstResponder")}
}

type _acceptsFirstResponderClass struct {
	objc.Class
}

// An interface definition for the [acceptsFirstResponder] class.
type IacceptsFirstResponder interface {
	ID() objc.ID
}

type acceptsFirstResponder struct {
	id objc.ID
}

func acceptsFirstResponderFrom(ptr unsafe.Pointer) acceptsFirstResponder {
	return acceptsFirstResponder{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ acceptsFirstResponder) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _acceptsFirstResponderClass) Alloc() acceptsFirstResponder {
	rv := objc.Send[acceptsFirstResponder](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _acceptsFirstResponderClass) New() acceptsFirstResponder {
	rv := objc.Send[acceptsFirstResponder](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewacceptsFirstResponder creates and returns a new initialized instance.
func NewacceptsFirstResponder() acceptsFirstResponder {
	return acceptsFirstResponderClass.New()
}

// Init initializes the instance.
func (a_ acceptsFirstResponder) Init() acceptsFirstResponder {
	rv := objc.Send[acceptsFirstResponder](a_.ID(), selInit)
	return rv
}

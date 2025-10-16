
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [initialFirstResponder] class.
var initialFirstResponderClass _initialFirstResponderClass

func init() {
	initialFirstResponderClass = _initialFirstResponderClass{objc.GetClass("initialFirstResponder")}
}

type _initialFirstResponderClass struct {
	objc.Class
}

// An interface definition for the [initialFirstResponder] class.
type IinitialFirstResponder interface {
	ID() objc.ID
}

type initialFirstResponder struct {
	id objc.ID
}

func initialFirstResponderFrom(ptr unsafe.Pointer) initialFirstResponder {
	return initialFirstResponder{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ initialFirstResponder) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _initialFirstResponderClass) Alloc() initialFirstResponder {
	rv := objc.Send[initialFirstResponder](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _initialFirstResponderClass) New() initialFirstResponder {
	rv := objc.Send[initialFirstResponder](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewinitialFirstResponder creates and returns a new initialized instance.
func NewinitialFirstResponder() initialFirstResponder {
	return initialFirstResponderClass.New()
}

// Init initializes the instance.
func (i_ initialFirstResponder) Init() initialFirstResponder {
	rv := objc.Send[initialFirstResponder](i_.ID(), selInit)
	return rv
}

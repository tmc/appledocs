
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [inputContext] class.
var inputContextClass _inputContextClass

func init() {
	inputContextClass = _inputContextClass{objc.GetClass("inputContext")}
}

type _inputContextClass struct {
	objc.Class
}

// An interface definition for the [inputContext] class.
type IinputContext interface {
	ID() objc.ID
}

type inputContext struct {
	id objc.ID
}

func inputContextFrom(ptr unsafe.Pointer) inputContext {
	return inputContext{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ inputContext) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _inputContextClass) Alloc() inputContext {
	rv := objc.Send[inputContext](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _inputContextClass) New() inputContext {
	rv := objc.Send[inputContext](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewinputContext creates and returns a new initialized instance.
func NewinputContext() inputContext {
	return inputContextClass.New()
}

// Init initializes the instance.
func (i_ inputContext) Init() inputContext {
	rv := objc.Send[inputContext](i_.ID(), selInit)
	return rv
}

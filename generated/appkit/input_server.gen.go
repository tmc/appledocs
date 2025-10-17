
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [InputServer] class.
var InputServerClass _InputServerClass

func init() {
	InputServerClass = _InputServerClass{objc.GetClass("NSInputServer")}
}

type _InputServerClass struct {
	objc.Class
}

// An interface definition for the [InputServer] class.
type IInputServer interface {
	ID() objc.ID
}

type InputServer struct {
	id objc.ID
}

func InputServerFrom(ptr unsafe.Pointer) InputServer {
	return InputServer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ InputServer) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _InputServerClass) Alloc() InputServer {
	rv := objc.Send[InputServer](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _InputServerClass) New() InputServer {
	rv := objc.Send[InputServer](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewInputServer creates and returns a new initialized instance.
func NewInputServer() InputServer {
	return InputServerClass.New()
}

// Init initializes the instance.
func (i_ InputServer) Init() InputServer {
	rv := objc.Send[InputServer](i_.ID(), selInit)
	return rv
}

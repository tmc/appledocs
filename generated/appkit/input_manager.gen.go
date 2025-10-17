
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [InputManager] class.
var InputManagerClass _InputManagerClass

func init() {
	InputManagerClass = _InputManagerClass{objc.GetClass("NSInputManager")}
}

type _InputManagerClass struct {
	objc.Class
}

// An interface definition for the [InputManager] class.
type IInputManager interface {
	ID() objc.ID
}

type InputManager struct {
	id objc.ID
}

func InputManagerFrom(ptr unsafe.Pointer) InputManager {
	return InputManager{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ InputManager) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _InputManagerClass) Alloc() InputManager {
	rv := objc.Send[InputManager](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _InputManagerClass) New() InputManager {
	rv := objc.Send[InputManager](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewInputManager creates and returns a new initialized instance.
func NewInputManager() InputManager {
	return InputManagerClass.New()
}

// Init initializes the instance.
func (i_ InputManager) Init() InputManager {
	rv := objc.Send[InputManager](i_.ID(), selInit)
	return rv
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [state] class.
var stateClass _stateClass

func init() {
	stateClass = _stateClass{objc.GetClass("state")}
}

type _stateClass struct {
	objc.Class
}

// An interface definition for the [state] class.
type Istate interface {
	ID() objc.ID
}

type state struct {
	id objc.ID
}

func stateFrom(ptr unsafe.Pointer) state {
	return state{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ state) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _stateClass) Alloc() state {
	rv := objc.Send[state](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _stateClass) New() state {
	rv := objc.Send[state](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newstate creates and returns a new initialized instance.
func Newstate() state {
	return stateClass.New()
}

// Init initializes the instance.
func (s_ state) Init() state {
	rv := objc.Send[state](s_.ID(), selInit)
	return rv
}

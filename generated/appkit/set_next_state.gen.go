
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [setNextState] class.
var setNextStateClass _setNextStateClass

func init() {
	setNextStateClass = _setNextStateClass{objc.GetClass("setNextState")}
}

type _setNextStateClass struct {
	objc.Class
}

// An interface definition for the [setNextState] class.
type IsetNextState interface {
	ID() objc.ID
}

type setNextState struct {
	id objc.ID
}

func setNextStateFrom(ptr unsafe.Pointer) setNextState {
	return setNextState{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ setNextState) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _setNextStateClass) Alloc() setNextState {
	rv := objc.Send[setNextState](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _setNextStateClass) New() setNextState {
	rv := objc.Send[setNextState](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewsetNextState creates and returns a new initialized instance.
func NewsetNextState() setNextState {
	return setNextStateClass.New()
}

// Init initializes the instance.
func (s_ setNextState) Init() setNextState {
	rv := objc.Send[setNextState](s_.ID(), selInit)
	return rv
}

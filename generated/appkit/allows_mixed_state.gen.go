
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsMixedState] class.
var allowsMixedStateClass _allowsMixedStateClass

func init() {
	allowsMixedStateClass = _allowsMixedStateClass{objc.GetClass("allowsMixedState")}
}

type _allowsMixedStateClass struct {
	objc.Class
}

// An interface definition for the [allowsMixedState] class.
type IallowsMixedState interface {
	ID() objc.ID
}

type allowsMixedState struct {
	id objc.ID
}

func allowsMixedStateFrom(ptr unsafe.Pointer) allowsMixedState {
	return allowsMixedState{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsMixedState) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsMixedStateClass) Alloc() allowsMixedState {
	rv := objc.Send[allowsMixedState](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsMixedStateClass) New() allowsMixedState {
	rv := objc.Send[allowsMixedState](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsMixedState creates and returns a new initialized instance.
func NewallowsMixedState() allowsMixedState {
	return allowsMixedStateClass.New()
}

// Init initializes the instance.
func (a_ allowsMixedState) Init() allowsMixedState {
	rv := objc.Send[allowsMixedState](a_.ID(), selInit)
	return rv
}

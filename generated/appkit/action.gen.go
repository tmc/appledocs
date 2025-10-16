
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [action] class.
var actionClass _actionClass

func init() {
	actionClass = _actionClass{objc.GetClass("action")}
}

type _actionClass struct {
	objc.Class
}

// An interface definition for the [action] class.
type Iaction interface {
	ID() objc.ID
}

type action struct {
	id objc.ID
}

func actionFrom(ptr unsafe.Pointer) action {
	return action{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ action) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _actionClass) Alloc() action {
	rv := objc.Send[action](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _actionClass) New() action {
	rv := objc.Send[action](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newaction creates and returns a new initialized instance.
func Newaction() action {
	return actionClass.New()
}

// Init initializes the instance.
func (a_ action) Init() action {
	rv := objc.Send[action](a_.ID(), selInit)
	return rv
}

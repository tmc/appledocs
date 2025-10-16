
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [nextState] class.
var nextStateClass _nextStateClass

func init() {
	nextStateClass = _nextStateClass{objc.GetClass("nextState")}
}

type _nextStateClass struct {
	objc.Class
}

// An interface definition for the [nextState] class.
type InextState interface {
	ID() objc.ID
}

type nextState struct {
	id objc.ID
}

func nextStateFrom(ptr unsafe.Pointer) nextState {
	return nextState{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ nextState) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _nextStateClass) Alloc() nextState {
	rv := objc.Send[nextState](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _nextStateClass) New() nextState {
	rv := objc.Send[nextState](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewnextState creates and returns a new initialized instance.
func NewnextState() nextState {
	return nextStateClass.New()
}

// Init initializes the instance.
func (n_ nextState) Init() nextState {
	rv := objc.Send[nextState](n_.ID(), selInit)
	return rv
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RunningApplication] class.
var RunningApplicationClass _RunningApplicationClass

func init() {
	RunningApplicationClass = _RunningApplicationClass{objc.GetClass("NSRunningApplication")}
}

type _RunningApplicationClass struct {
	objc.Class
}

// An interface definition for the [RunningApplication] class.
type IRunningApplication interface {
	ID() objc.ID
}

type RunningApplication struct {
	id objc.ID
}

func RunningApplicationFrom(ptr unsafe.Pointer) RunningApplication {
	return RunningApplication{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ RunningApplication) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _RunningApplicationClass) Alloc() RunningApplication {
	rv := objc.Send[RunningApplication](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _RunningApplicationClass) New() RunningApplication {
	rv := objc.Send[RunningApplication](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewRunningApplication creates and returns a new initialized instance.
func NewRunningApplication() RunningApplication {
	return RunningApplicationClass.New()
}

// Init initializes the instance.
func (r_ RunningApplication) Init() RunningApplication {
	rv := objc.Send[RunningApplication](r_.ID(), selInit)
	return rv
}

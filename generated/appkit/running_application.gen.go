// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RunningApplication] class.
var runningApplicationClass = _RunningApplicationClass{objc.GetClass("NSRunningApplication")}

type _RunningApplicationClass struct {
	class objc.Class
}

// An interface definition for the [RunningApplication] class.
type IRunningApplication interface {
	objectivec.IObject
}

// An object that can manipulate and provide information for a single instance of an app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication

type RunningApplication struct {
	objectivec.Object
}

// RunningApplicationFrom constructs a [RunningApplication] from an unsafe.Pointer.
//
// An object that can manipulate and provide information for a single instance of an app.
func RunningApplicationFrom(ptr unsafe.Pointer) RunningApplication {
	return RunningApplication{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (rc _RunningApplicationClass) Alloc() RunningApplication {
	rv := objc.Send[RunningApplication](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (rc _RunningApplicationClass) New() RunningApplication {
	rv := objc.Send[RunningApplication](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RunningApplication) Init() RunningApplication {
	rv := objc.Send[RunningApplication](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RunningApplication) Autorelease() RunningApplication {
	rv := objc.Send[RunningApplication](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRunningApplication creates a new RunningApplication instance.
func NewRunningApplication() RunningApplication {
	return runningApplicationClass.New()
}





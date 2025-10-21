// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [RunningApplication] class.
var (
	RunningApplicationClass     _RunningApplicationClass
	RunningApplicationClassOnce sync.Once
)

func getRunningApplicationClass() _RunningApplicationClass {
	RunningApplicationClassOnce.Do(func() {
		RunningApplicationClass = _RunningApplicationClass{objc.GetClass("NSRunningApplication")}
	})
	return RunningApplicationClass
}

type _RunningApplicationClass struct {
	class objc.Class
}

// An interface definition for the [RunningApplication] class.
type IRunningApplication interface {
	objectivec.IObject
}

// An object that can manipulate and provide information for a single instance of an app.
//
// Some properties of an app are fixed, such as the bundle identifier. Other properties may vary over time, such as whether the app is hidden. Properties that vary can be observed with key-value observing, in which case the description comment for the method notes this capability. Properties that vary over time are inherently race-prone. For example, a hidden app may unhide itself at any time. To ameliorate this, properties persist until the next turn of the main run loop in a common mode. For example, if you repeatedly poll an unhidden app for its hidden property without allowing the run loop to run, it will continue to return , even if the app hides, until the next turn of the run loop. is thread safe, in that its properties are returned atomically. However, it is still subject to the main run loop policy described above. If you access an instance of from a background thread, be aware that its time-varying properties may change from under you as the main run loop runs (or not). An instance remains valid after the app exits. However, most properties lose their significance, and some properties may not be available on a terminated application. To access the list of all running apps, use the method in .
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getRunningApplicationClass().New()
}





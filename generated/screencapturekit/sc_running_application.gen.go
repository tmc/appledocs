// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RunningApplication] class.
var (
	RunningApplicationClass     _RunningApplicationClass
	RunningApplicationClassOnce sync.Once
)

func getRunningApplicationClass() _RunningApplicationClass {
	RunningApplicationClassOnce.Do(func() {
		RunningApplicationClass = _RunningApplicationClass{objc.GetClass("SCRunningApplication")}
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

// An instance that represents an app running on a device.
//
// Retrieve the available apps from an instance of . Select one or more apps to capture and use them to create an instance of . Apply the filter to an instance of to limit its output to content matching your criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRunningApplication
type RunningApplication struct {
	objectivec.Object
}

// RunningApplicationFrom constructs a [RunningApplication] from an unsafe.Pointer.
//
// An instance that represents an app running on a device.
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


// The display name of the app.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRunningApplication/applicationName
func (r_ RunningApplication) ApplicationName() appkit.string {
	rv := objc.Send[appkit.string](r_.ID, objc.Sel("applicationName"))
	return rv
}

// The unique bundle identifier of the app.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRunningApplication/bundleIdentifier
func (r_ RunningApplication) BundleIdentifier() appkit.string {
	rv := objc.Send[appkit.string](r_.ID, objc.Sel("bundleIdentifier"))
	return rv
}

// The system process identifier of the app.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRunningApplication/processID
func (r_ RunningApplication) ProcessID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("processID"))
	return rv
}




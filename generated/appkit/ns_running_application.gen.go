// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	ActivationPolicy() ApplicationActivationPolicy
	BundleIdentifier() foundation.foundation.INSString
	BundleURL() foundation.foundation.INSURL
	ExecutableArchitecture() int
	ExecutableURL() foundation.foundation.INSURL
	Icon() IImage
	Active() bool
	FinishedLaunching() bool
	Hidden() bool
	Terminated() bool
	LaunchDate() foundation.foundation.INSDate
	LocalizedName() foundation.foundation.INSString
	OwnsMenuBar() bool
	ProcessIdentifier() objectivec.IObject
	IsActive() bool
	SetIsActive(value bool)
	IsFinishedLaunching() bool
	SetIsFinishedLaunching(value bool)
	IsHidden() bool
	SetIsHidden(value bool)
	IsTerminated() bool
	SetIsTerminated(value bool)
	RunningApplications() IRunningApplication
	SetRunningApplications(value IRunningApplication)


	

	// methods:
	ActivateFromApplicationOptions(application IRunningApplication, options ApplicationActivationOptions) bool
	ActivateWithOptions(options ApplicationActivationOptions) bool
	ForceTerminate() bool
	Hide() bool
	Terminate() bool
	Unhide() bool


}





// Alloc allocates a new instance without initialization.
func (rc _RunningApplicationClass) Alloc() RunningApplication {
	rv := objc.Send[RunningApplication](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An object that can manipulate and provide information for a single instance of an app.
//
// Some properties of an app are fixed, such as the bundle identifier. Other properties may vary over time, such as whether the app is hidden. Properties that vary can be observed with key-value observing, in which case the description comment for the method notes this capability. Properties that vary over time are inherently race-prone. For example, a hidden app may unhide itself at any time. To ameliorate this, properties persist until the next turn of the main run loop in a common mode. For example, if you repeatedly poll an unhidden app for its hidden property without allowing the run loop to run, it will continue to return , even if the app hides, until the next turn of the run loop. is thread safe, in that its properties are returned atomically. However, it is still subject to the main run loop policy described above. If you access an instance of from a background thread, be aware that its time-varying properties may change from under you as the main run loop runs (or not). An instance remains valid after the app exits. However, most properties lose their significance, and some properties may not be available on a terminated application. To access the list of all running apps, use the method in .


// An object that can manipulate and provide information for a single instance of an app.
//
// [Full Topic]
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






// Returns the running application with the given process identifier, or nil if no application has that pid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/init(processIdentifier:)
func NewRunningApplicationWithProcessIdentifier(pid objectivec.IObject) RunningApplication {
	rv := objc.Send[RunningApplication](objc.ID(getRunningApplicationClass().class), objc.Sel("runningApplicationWithProcessIdentifier:"), pid)
	return rv
}







// Returns the running application with the given process identifier, or nil if no application has that pid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/init(processIdentifier:)
func (rc _RunningApplicationClass) RunningApplicationWithProcessIdentifier(pid objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(rc.class), objc.Sel("runningApplicationWithProcessIdentifier:"), pid)
	return rv
}


// Returns an array of currently running applications with the specified bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/runningApplications(withBundleIdentifier:)
func (rc _RunningApplicationClass) RunningApplicationsWithBundleIdentifier(bundleIdentifier foundation.foundation.INSString) []RunningApplication {
	rv := objc.Send[[]RunningApplication](objc.ID(rc.class), objc.Sel("runningApplicationsWithBundleIdentifier:"), bundleIdentifier)
	return rv
}


// Terminates invisibly running applications as if triggered by system memory pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/terminateAutomaticallyTerminableApplications()
func (rc _RunningApplicationClass) TerminateAutomaticallyTerminableApplications() {
	objc.Send[objc.ID](objc.ID(rc.class), objc.Sel("terminateAutomaticallyTerminableApplications"))
}







// Returns an representing this application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/current
func (rc _RunningApplicationClass) CurrentApplication() RunningApplication {
	rv := objc.Send[RunningApplication](objc.ID(rc.class), objc.Sel("currentApplication"))
	return rv
}






// Attempts to activate the application using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/activate(from:options:)
func (r_ RunningApplication) ActivateFromApplicationOptions(application IRunningApplication, options ApplicationActivationOptions) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("activateFromApplication:options:"), application, options)
	return rv
}


// Attempts to activate the application using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/activate(options:)
func (r_ RunningApplication) ActivateWithOptions(options ApplicationActivationOptions) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("activateWithOptions:"), options)
	return rv
}


// Attempts to force the receiver to quit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/forceTerminate()
func (r_ RunningApplication) ForceTerminate() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("forceTerminate"))
	return rv
}


// Attempts to hide or the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/hide()
func (r_ RunningApplication) Hide() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("hide"))
	return rv
}


// Attempts to quit the receiver normally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/terminate()
func (r_ RunningApplication) Terminate() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("terminate"))
	return rv
}


// Attempts to unhide or the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/unhide()
func (r_ RunningApplication) Unhide() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("unhide"))
	return rv
}







// Indicates the activation policy of the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/activationPolicy
func (r_ RunningApplication) ActivationPolicy() ApplicationActivationPolicy {
	rv := objc.Send[ApplicationActivationPolicy](r_.ID, objc.Sel("activationPolicy"))
	return rv
}


// Indicates the of the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/bundleIdentifier
func (r_ RunningApplication) BundleIdentifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("bundleIdentifier"))
	return rv
}


// Indicates the URL to the application’s bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/bundleURL
func (r_ RunningApplication) BundleURL() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](r_.ID, objc.Sel("bundleURL"))
	return rv
}


// Returns an representing this application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/current
func (r_ RunningApplication) CurrentApplication() IRunningApplication {
	rv := objc.Send[RunningApplication](r_.ID, objc.Sel("currentApplication"))
	return rv
}


// Indicates the executing processor architecture for the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/executableArchitecture
func (r_ RunningApplication) ExecutableArchitecture() int {
	rv := objc.Send[int](r_.ID, objc.Sel("executableArchitecture"))
	return rv
}


// Indicates the URL to the application’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/executableURL
func (r_ RunningApplication) ExecutableURL() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](r_.ID, objc.Sel("executableURL"))
	return rv
}


// Returns the icon for the receiver’s application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/icon
func (r_ RunningApplication) Icon() IImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("icon"))
	return rv
}


// Indicates whether the application is currently frontmost.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/isActive
func (r_ RunningApplication) Active() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("active"))
	return rv
}


// A Boolean value that determines whether the receiver’s process has finished launching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/isFinishedLaunching
func (r_ RunningApplication) FinishedLaunching() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("finishedLaunching"))
	return rv
}


// Indicates whether the application is currently hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/isHidden
func (r_ RunningApplication) Hidden() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("hidden"))
	return rv
}


// Indicates that the receiver’s application has terminated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/isTerminated
func (r_ RunningApplication) Terminated() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("terminated"))
	return rv
}


// Indicates the date when the application was launched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/launchDate
func (r_ RunningApplication) LaunchDate() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](r_.ID, objc.Sel("launchDate"))
	return rv
}


// Indicates the localized name of the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/localizedName
func (r_ RunningApplication) LocalizedName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("localizedName"))
	return rv
}


// Returns whether the application owns the current menu bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/ownsMenuBar
func (r_ RunningApplication) OwnsMenuBar() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("ownsMenuBar"))
	return rv
}


// Indicates the process identifier (pid) of the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication/processIdentifier
func (r_ RunningApplication) ProcessIdentifier() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("processIdentifier"))
	return rv
}


// Indicates whether the application is currently frontmost.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/isactive
func (r_ RunningApplication) IsActive() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isActive"))
	return rv
}


// Indicates whether the application is currently frontmost.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/isactive
func (r_ RunningApplication) SetIsActive(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsActive:"), value)
}


// A Boolean value that determines whether the receiver’s process has finished launching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/isfinishedlaunching
func (r_ RunningApplication) IsFinishedLaunching() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isFinishedLaunching"))
	return rv
}


// A Boolean value that determines whether the receiver’s process has finished launching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/isfinishedlaunching
func (r_ RunningApplication) SetIsFinishedLaunching(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsFinishedLaunching:"), value)
}


// Indicates whether the application is currently hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/ishidden
func (r_ RunningApplication) IsHidden() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isHidden"))
	return rv
}


// Indicates whether the application is currently hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/ishidden
func (r_ RunningApplication) SetIsHidden(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsHidden:"), value)
}


// Indicates that the receiver’s application has terminated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/isterminated
func (r_ RunningApplication) IsTerminated() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isTerminated"))
	return rv
}


// Indicates that the receiver’s application has terminated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/isterminated
func (r_ RunningApplication) SetIsTerminated(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsTerminated:"), value)
}


// Returns an array of running apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/runningapplications
func (r_ RunningApplication) RunningApplications() IRunningApplication {
	rv := objc.Send[RunningApplication](r_.ID, objc.Sel("runningApplications"))
	return rv
}


// Returns an array of running apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/runningapplications
func (r_ RunningApplication) SetRunningApplications(value IRunningApplication) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRunningApplications:"), value)
}








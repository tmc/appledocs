// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	ActivationPolicy() unsafe.Pointer
	SetActivationPolicy(value unsafe.Pointer)
	BundleIdentifier() string
	SetBundleIdentifier(value string)
	BundleURL() foundation.URL
	SetBundleURL(value foundation.IURL)
	ExecutableArchitecture() int
	SetExecutableArchitecture(value int)
	ExecutableURL() foundation.URL
	SetExecutableURL(value foundation.IURL)
	Icon() Image
	SetIcon(value IImage)
	IsActive() bool
	SetIsActive(value bool)
	IsFinishedLaunching() bool
	SetIsFinishedLaunching(value bool)
	IsHidden() bool
	SetIsHidden(value bool)
	IsTerminated() bool
	SetIsTerminated(value bool)
	LaunchDate() foundation.Date
	SetLaunchDate(value foundation.IDate)
	LocalizedName() string
	SetLocalizedName(value string)
	OwnsMenuBar() bool
	SetOwnsMenuBar(value bool)
	ProcessIdentifier() unsafe.Pointer
	SetProcessIdentifier(value unsafe.Pointer)
	RunningApplications() NSRunningApplication
	SetRunningApplications(value IRunningApplication)
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



// Indicates the activation policy of the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/activationpolicy

func (r_ RunningApplication) ActivationPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("activationPolicy"))
	return rv
}


// Indicates the activation policy of the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/activationpolicy

func (r_ RunningApplication) SetActivationPolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setActivationPolicy:"), value)
}


// Indicates the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/bundleidentifier

func (r_ RunningApplication) BundleIdentifier() string {
	rv := objc.Send[string](r_.ID, objc.Sel("bundleIdentifier"))
	return rv
}


// Indicates the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/bundleidentifier

func (r_ RunningApplication) SetBundleIdentifier(value string) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBundleIdentifier:"), objc.String(value))
}


// Indicates the URL to the application’s bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/bundleurl

func (r_ RunningApplication) BundleURL() foundation.URL {
	rv := objc.Send[foundation.URL](r_.ID, objc.Sel("bundleURL"))
	return rv
}


// Indicates the URL to the application’s bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/bundleurl

func (r_ RunningApplication) SetBundleURL(value foundation.IURL) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBundleURL:"), value)
}


// Indicates the executing processor architecture for the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/executablearchitecture

func (r_ RunningApplication) ExecutableArchitecture() int {
	rv := objc.Send[int](r_.ID, objc.Sel("executableArchitecture"))
	return rv
}


// Indicates the executing processor architecture for the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/executablearchitecture

func (r_ RunningApplication) SetExecutableArchitecture(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setExecutableArchitecture:"), value)
}


// Indicates the URL to the application’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/executableurl

func (r_ RunningApplication) ExecutableURL() foundation.URL {
	rv := objc.Send[foundation.URL](r_.ID, objc.Sel("executableURL"))
	return rv
}


// Indicates the URL to the application’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/executableurl

func (r_ RunningApplication) SetExecutableURL(value foundation.IURL) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setExecutableURL:"), value)
}


// Returns the icon for the receiver’s application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/icon

func (r_ RunningApplication) Icon() Image {
	rv := objc.Send[Image](r_.ID, objc.Sel("icon"))
	return rv
}


// Returns the icon for the receiver’s application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/icon

func (r_ RunningApplication) SetIcon(value IImage) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIcon:"), value)
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


// Indicates the date when the application was launched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/launchdate

func (r_ RunningApplication) LaunchDate() foundation.Date {
	rv := objc.Send[foundation.Date](r_.ID, objc.Sel("launchDate"))
	return rv
}


// Indicates the date when the application was launched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/launchdate

func (r_ RunningApplication) SetLaunchDate(value foundation.IDate) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLaunchDate:"), value)
}


// Indicates the localized name of the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/localizedname

func (r_ RunningApplication) LocalizedName() string {
	rv := objc.Send[string](r_.ID, objc.Sel("localizedName"))
	return rv
}


// Indicates the localized name of the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/localizedname

func (r_ RunningApplication) SetLocalizedName(value string) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLocalizedName:"), objc.String(value))
}


// Returns whether the application owns the current menu bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/ownsmenubar

func (r_ RunningApplication) OwnsMenuBar() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("ownsMenuBar"))
	return rv
}


// Returns whether the application owns the current menu bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/ownsmenubar

func (r_ RunningApplication) SetOwnsMenuBar(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOwnsMenuBar:"), value)
}


// Indicates the process identifier (pid) of the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/processidentifier

func (r_ RunningApplication) ProcessIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("processIdentifier"))
	return rv
}


// Indicates the process identifier (pid) of the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/processidentifier

func (r_ RunningApplication) SetProcessIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setProcessIdentifier:"), value)
}


// Returns an array of running apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/runningapplications

func (r_ RunningApplication) RunningApplications() NSRunningApplication {
	rv := objc.Send[NSRunningApplication](r_.ID, objc.Sel("runningApplications"))
	return rv
}


// Returns an array of running apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/runningapplications

func (r_ RunningApplication) SetRunningApplications(value IRunningApplication) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRunningApplications:"), value)
}




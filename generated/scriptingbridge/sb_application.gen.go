// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

package scriptingbridge

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [SBApplication] class.
var (
	SBApplicationClass     _SBApplicationClass
	SBApplicationClassOnce sync.Once
)

func getSBApplicationClass() _SBApplicationClass {
	SBApplicationClassOnce.Do(func() {
		SBApplicationClass = _SBApplicationClass{objc.GetClass("SBApplication")}
	})
	return SBApplicationClass
}

type _SBApplicationClass struct {
	class objc.Class
}

// An interface definition for the [SBApplication] class.
type ISBApplication interface {
	ISBObject
	Activate()
	ClassForScriptingClass(className string) objc.Class
}

// The class provides a mechanism enabling an Objective-C program to send Apple events to a scriptable application and receive Apple events in response. It thereby makes it possible for that program to control the application and exchange data with it. Scripting Bridge works by bridging data types between Apple event descriptors and Cocoa objects.
//
// Although includes methods that manually send and process Apple events, you should never have to call these methods directly. Instead, subclasses of implement application-specific methods that handle the sending of Apple events automatically. For example, if you wanted to get the current iTunes track, you can simply use the method of the dynamically defined subclass for the iTunes application—which handles the details of sending the Apple event for you—rather than figuring out the more complicated, low-level alternative: If you do need to send Apple events manually, consider using the class.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication
type SBApplication struct {
	SBObject
}

// SBApplicationFrom constructs a [SBApplication] from an unsafe.Pointer.
//
// The class provides a mechanism enabling an Objective-C program to send Apple events to a scriptable application and receive Apple events in response. It thereby makes it possible for that program to control the application and exchange data with it. Scripting Bridge works by bridging data types between Apple event descriptors and Cocoa objects.
func SBApplicationFrom(ptr unsafe.Pointer) SBApplication {
	return SBApplication{
		SBObject: SBObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SBApplicationClass) Alloc() SBApplication {
	rv := objc.Send[SBApplication](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SBApplicationClass) New() SBApplication {
	rv := objc.Send[SBApplication](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SBApplication) Init() SBApplication {
	rv := objc.Send[SBApplication](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SBApplication) Autorelease() SBApplication {
	rv := objc.Send[SBApplication](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSBApplication creates a new SBApplication instance.
func NewSBApplication() SBApplication {
	return getSBApplicationClass().New()
}




// Returns an instance of an subclass that represents the target application identified by the given bundle identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/init(bundleIdentifier:)
func NewSBApplicationWithBundleIdentifier(ident string) SBApplication {
	instance := getSBApplicationClass().Alloc()
	rv := objc.Send[SBApplication](instance.ID, objc.Sel("initWithBundleIdentifier:"), objc.String(ident))
	rv.Autorelease()
	return rv
}



// Returns an instance of an subclass that represents the target application identified by the given process identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/init(processIdentifier:)
func NewSBApplicationWithProcessIdentifier(pid unsafe.Pointer) SBApplication {
	instance := getSBApplicationClass().Alloc()
	rv := objc.Send[SBApplication](instance.ID, objc.Sel("initWithProcessIdentifier:"), pid)
	rv.Autorelease()
	return rv
}



// Returns an instance of an subclass that represents the target application identified by the given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/init(url:)
func NewSBApplicationWithURL(url foundation.URL) SBApplication {
	instance := getSBApplicationClass().Alloc()
	rv := objc.Send[SBApplication](instance.ID, objc.Sel("initWithURL:"), url)
	rv.Autorelease()
	return rv
}


// Returns the shared instance representing the target application specified by its bundle identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/applicationWithBundleIdentifier:
func (sc _SBApplicationClass) ApplicationWithBundleIdentifier(ident string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("applicationWithBundleIdentifier:"), objc.String(ident))
	return rv
}

// Returns the shared instance representing a target application specified by its process identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/applicationWithProcessIdentifier:
func (sc _SBApplicationClass) ApplicationWithProcessIdentifier(pid unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("applicationWithProcessIdentifier:"), pid)
	return rv
}

// Returns the shared instance representing a target application specified by the given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/applicationWithURL:
func (sc _SBApplicationClass) ApplicationWithURL(url foundation.URL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("applicationWithURL:"), url)
	return rv
}

// Moves the target application to the foreground immediately.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/activate()
func (s_ SBApplication) Activate() {
	objc.Send[objc.ID](s_.ID, objc.Sel("activate"))
}

// Returns a class object that represents a particular class in the target application.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/class(forScriptingClass:)
func (s_ SBApplication) ClassForScriptingClass(className string) objc.Class {
	rv := objc.Send[objc.Class](s_.ID, objc.Sel("classForScriptingClass:"), objc.String(className))
	return rv
}

// A Boolean that indicates whether the target application represented by the
//
// [Full Topic]: https://developer.apple.com/documentation/scriptingbridge/sbapplication/isrunning
func (s_ SBApplication) IsRunning() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isRunning"))
	return rv
}


// SetIsRunning sets the value of the isRunning property.
// A Boolean that indicates whether the target application represented by the

//
// [Full Topic]: https://developer.apple.com/documentation/scriptingbridge/sbapplication/isrunning
func (s_ SBApplication) SetIsRunning(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsRunning:"), value)
}

// The error-handling delegate of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/delegate
func (s_ SBApplication) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The error-handling delegate of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/delegate
func (s_ SBApplication) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean that indicates whether the target application represented by the receiver is running.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/isRunning
func (s_ SBApplication) Running() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("running"))
	return rv
}

// The launch flags for the application represented by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/launchFlags
func (s_ SBApplication) LaunchFlags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("launchFlags"))
	return rv
}


// SetLaunchFlags sets the value of the launchFlags property.
// The launch flags for the application represented by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/launchFlags
func (s_ SBApplication) SetLaunchFlags(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLaunchFlags:"), value)
}

// The mode for sending Apple events to the target application.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/sendMode
func (s_ SBApplication) SendMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sendMode"))
	return rv
}


// SetSendMode sets the value of the sendMode property.
// The mode for sending Apple events to the target application.

//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/sendMode
func (s_ SBApplication) SetSendMode(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSendMode:"), value)
}

// The period the application will wait to receive reply Apple events.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/timeout
func (s_ SBApplication) Timeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("timeout"))
	return rv
}


// SetTimeout sets the value of the timeout property.
// The period the application will wait to receive reply Apple events.

//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/timeout
func (s_ SBApplication) SetTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTimeout:"), value)
}



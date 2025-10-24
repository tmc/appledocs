// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

package scriptingbridge

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class SBApplication */


/* debug [class_header]: Header for SBApplication */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SBApplication */
// An interface definition for the [SBApplication] class.
type ISBApplication interface {
	ISBObject
	
/* debug [class_interface_properties]: Properties for SBApplication */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Running() bool
	LaunchFlags() unsafe.Pointer
	SetLaunchFlags(value unsafe.Pointer)
	SendMode() unsafe.Pointer
	SetSendMode(value unsafe.Pointer)
	Timeout() unsafe.Pointer
	SetTimeout(value unsafe.Pointer)
	IsRunning() bool
	SetIsRunning(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SBApplication */
	// methods:
	Activate()
	ClassForScriptingClass(className objc.IObject /* cross-framework: NSString */) objc.Class
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SBApplication */
// Alloc allocates a new instance without initialization.
func (sc _SBApplicationClass) Alloc() SBApplication {
	rv := objc.Send[SBApplication](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SBApplication */
// The class provides a mechanism enabling an Objective-C program to send Apple events to a scriptable application and receive Apple events in response. It thereby makes it possible for that program to control the application and exchange data with it. Scripting Bridge works by bridging data types between Apple event descriptors and Cocoa objects.
//
// Although includes methods that manually send and process Apple events, you should never have to call these methods directly. Instead, subclasses of implement application-specific methods that handle the sending of Apple events automatically. For example, if you wanted to get the current iTunes track, you can simply use the method of the dynamically defined subclass for the iTunes application—which handles the details of sending the Apple event for you—rather than figuring out the more complicated, low-level alternative: If you do need to send Apple events manually, consider using the class.


// The class provides a mechanism enabling an Objective-C program to send Apple events to a scriptable application and receive Apple events in response. It thereby makes it possible for that program to control the application and exchange data with it. Scripting Bridge works by bridging data types between Apple event descriptors and Cocoa objects.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SBApplication */

// Returns an instance of an subclass that represents the target application identified by the given bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/init(bundleIdentifier:)
func NewSBApplicationWithBundleIdentifier(ident objc.IObject /* cross-framework: NSString */) SBApplication {
	instance := getSBApplicationClass().Alloc()
	rv := objc.Send[SBApplication](instance.ID, objc.Sel("initWithBundleIdentifier:"), ident)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSBApplicationWithBundleIdentifier */


// Returns an instance of an subclass that represents the target application identified by the given process identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/init(processIdentifier:)
func NewSBApplicationWithProcessIdentifier(pid unsafe.Pointer) SBApplication {
	instance := getSBApplicationClass().Alloc()
	rv := objc.Send[SBApplication](instance.ID, objc.Sel("initWithProcessIdentifier:"), pid)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSBApplicationWithProcessIdentifier */


// Returns an instance of an subclass that represents the target application identified by the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/init(url:)
func NewSBApplicationWithURL(url objc.IObject /* cross-framework: NSURL */) SBApplication {
	instance := getSBApplicationClass().Alloc()
	rv := objc.Send[SBApplication](instance.ID, objc.Sel("initWithURL:"), url)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSBApplicationWithURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SBApplication */

// Returns the shared instance representing the target application specified by its bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/applicationWithBundleIdentifier:
func (sc _SBApplicationClass) ApplicationWithBundleIdentifier(ident objc.IObject /* cross-framework: NSString */) SBApplication {
	rv := objc.Send[SBApplication](objc.ID(sc.class), objc.Sel("applicationWithBundleIdentifier:"), ident)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ApplicationWithBundleIdentifier) */


// Returns the shared instance representing a target application specified by its process identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/applicationWithProcessIdentifier:
func (sc _SBApplicationClass) ApplicationWithProcessIdentifier(pid unsafe.Pointer) SBApplication {
	rv := objc.Send[SBApplication](objc.ID(sc.class), objc.Sel("applicationWithProcessIdentifier:"), pid)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ApplicationWithProcessIdentifier) */


// Returns the shared instance representing a target application specified by the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/applicationWithURL:
func (sc _SBApplicationClass) ApplicationWithURL(url objc.IObject /* cross-framework: NSURL */) SBApplication {
	rv := objc.Send[SBApplication](objc.ID(sc.class), objc.Sel("applicationWithURL:"), url)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ApplicationWithURL) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SBApplication */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SBApplication */

// Moves the target application to the foreground immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/activate()
func (s_ SBApplication) Activate() {
	objc.Send[objc.ID](s_.ID, objc.Sel("activate"))
}/* debug [instance_methods/method]: Activate */


// Returns a class object that represents a particular class in the target application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/class(forScriptingClass:)
func (s_ SBApplication) ClassForScriptingClass(className objc.IObject /* cross-framework: NSString */) objc.Class {
	rv := objc.Send[objc.Class](s_.ID, objc.Sel("classForScriptingClass:"), className)
	return rv
}/* debug [instance_methods/method]: ClassForScriptingClass */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SBApplication */

// The error-handling delegate of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/delegate
func (s_ SBApplication) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The error-handling delegate of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/delegate
func (s_ SBApplication) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean that indicates whether the target application represented by the receiver is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/isRunning
func (s_ SBApplication) Running() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("running"))
	return rv
}/* debug [instance_properties/getter]: running */


// The launch flags for the application represented by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/launchFlags
func (s_ SBApplication) LaunchFlags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("launchFlags"))
	return rv
}/* debug [instance_properties/getter]: launchFlags */


// The launch flags for the application represented by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/launchFlags
func (s_ SBApplication) SetLaunchFlags(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLaunchFlags:"), value)
}/* debug [instance_properties/setter]: launchFlags */


// The mode for sending Apple events to the target application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/sendMode
func (s_ SBApplication) SendMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sendMode"))
	return rv
}/* debug [instance_properties/getter]: sendMode */


// The mode for sending Apple events to the target application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/sendMode
func (s_ SBApplication) SetSendMode(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSendMode:"), value)
}/* debug [instance_properties/setter]: sendMode */


// The period the application will wait to receive reply Apple events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/timeout
func (s_ SBApplication) Timeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("timeout"))
	return rv
}/* debug [instance_properties/getter]: timeout */


// The period the application will wait to receive reply Apple events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/timeout
func (s_ SBApplication) SetTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTimeout:"), value)
}/* debug [instance_properties/setter]: timeout */


// A Boolean that indicates whether the target application represented by the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/scriptingbridge/sbapplication/isrunning
func (s_ SBApplication) IsRunning() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isRunning"))
	return rv
}/* debug [instance_properties/getter]: isRunning */


// A Boolean that indicates whether the target application represented by the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/scriptingbridge/sbapplication/isrunning
func (s_ SBApplication) SetIsRunning(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsRunning:"), value)
}/* debug [instance_properties/setter]: isRunning */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SBApplication */



// Code generated from Apple documentation for SafetyKit. DO NOT EDIT.

package safetykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [SACrashDetectionManager] class.
var (
	SACrashDetectionManagerClass     _SACrashDetectionManagerClass
	SACrashDetectionManagerClassOnce sync.Once
)

func getSACrashDetectionManagerClass() _SACrashDetectionManagerClass {
	SACrashDetectionManagerClassOnce.Do(func() {
		SACrashDetectionManagerClass = _SACrashDetectionManagerClass{objc.GetClass("SACrashDetectionManager")}
	})
	return SACrashDetectionManagerClass
}

type _SACrashDetectionManagerClass struct {
	class objc.Class
}

// An interface definition for the [SACrashDetectionManager] class.
type ISACrashDetectionManager interface {
	objectivec.IObject
	RequestAuthorizationWithCompletionHandler(handler unsafe.Pointer)
}

// Provides registration and management of Crash Detection events.
//
// Use this class to determine Crash Detection availabilty on iPhone, detect authorization status, and register for Crash Detection events. Not all iPhones support Crash Detection, so verify that returns . Check the value of to determine if the person designates this app on their iPhone to receive Crash Detection events. If the value is not , set and call to request authorization. After your app has authorization to receive Crash Detection events, adopt and implement . If a vehicular crash occurs, the system calls the method with the Crash Detection event.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionManager
type SACrashDetectionManager struct {
	objectivec.Object
}

// SACrashDetectionManagerFrom constructs a [SACrashDetectionManager] from an unsafe.Pointer.
//
// Provides registration and management of Crash Detection events.
func SACrashDetectionManagerFrom(ptr unsafe.Pointer) SACrashDetectionManager {
	return SACrashDetectionManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SACrashDetectionManagerClass) Alloc() SACrashDetectionManager {
	rv := objc.Send[SACrashDetectionManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SACrashDetectionManagerClass) New() SACrashDetectionManager {
	rv := objc.Send[SACrashDetectionManager](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SACrashDetectionManager) Init() SACrashDetectionManager {
	rv := objc.Send[SACrashDetectionManager](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SACrashDetectionManager) Autorelease() SACrashDetectionManager {
	rv := objc.Send[SACrashDetectionManager](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSACrashDetectionManager creates a new SACrashDetectionManager instance.
func NewSACrashDetectionManager() SACrashDetectionManager {
	return getSACrashDetectionManagerClass().New()
}


// A Boolean value that indicates if Crash Detection is available.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionManager/isAvailable
func (sc _SACrashDetectionManagerClass) Available() bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("available"))
	return rv
}
// Requests permission to access Crash Detection information.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionManager/requestAuthorization(completionHandler:)
func (s_ SACrashDetectionManager) RequestAuthorizationWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("requestAuthorizationWithCompletionHandler:"), handler)
}

// A value that indicates if the person authorized the app to receive Crash Detection events.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionManager/authorizationStatus
func (s_ SACrashDetectionManager) AuthorizationStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("authorizationStatus"))
	return rv
}

// The object that receives Crash Detection events.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionManager/delegate
func (s_ SACrashDetectionManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The object that receives Crash Detection events.

//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionManager/delegate
func (s_ SACrashDetectionManager) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates if Crash Detection is available.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionManager/isAvailable
func (s_ SACrashDetectionManager) Available() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("available"))
	return rv
}




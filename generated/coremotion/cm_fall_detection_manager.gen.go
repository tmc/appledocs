// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FallDetectionManager] class.
var (
	FallDetectionManagerClass     _FallDetectionManagerClass
	FallDetectionManagerClassOnce sync.Once
)

func getFallDetectionManagerClass() _FallDetectionManagerClass {
	FallDetectionManagerClassOnce.Do(func() {
		FallDetectionManagerClass = _FallDetectionManagerClass{objc.GetClass("CMFallDetectionManager")}
	})
	return FallDetectionManagerClass
}

type _FallDetectionManagerClass struct {
	class objc.Class
}

// An interface definition for the [FallDetectionManager] class.
type IFallDetectionManager interface {
	objectivec.IObject
	RequestAuthorizationWithHandler(handler unsafe.Pointer)
	AuthorizationStatus() AuthorizationStatus
	Delegate() objc.ID
	SetDelegate(value objc.ID)
}

// An object for managing fall detection events.
//
// In Series 4 and later, Apple Watch can detect when a wearer falls, and contact emergency services if necessary. Using the , your app can request the user’s authorization, and set up a delegate to receive notifications about these . For more information, see . requires an entitlement from Apple. To apply for the entitlement, see . This entitlement allows the app to run in the background without requiring any additional capabilities. However, you can add capabilities for other background modes, as needed by your app. There are two approaches to detecting falls in your app. You can either query for samples in HealthKit, or you can use Core Motion’s .


// An object for managing fall detection events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionManager

type FallDetectionManager struct {
	objectivec.Object
}

// FallDetectionManagerFrom constructs a [FallDetectionManager] from an unsafe.Pointer.
//
// An object for managing fall detection events.
func FallDetectionManagerFrom(ptr unsafe.Pointer) FallDetectionManager {
	return FallDetectionManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FallDetectionManagerClass) Alloc() FallDetectionManager {
	rv := objc.Send[FallDetectionManager](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FallDetectionManagerClass) New() FallDetectionManager {
	rv := objc.Send[FallDetectionManager](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FallDetectionManager) Init() FallDetectionManager {
	rv := objc.Send[FallDetectionManager](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FallDetectionManager) Autorelease() FallDetectionManager {
	rv := objc.Send[FallDetectionManager](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFallDetectionManager creates a new FallDetectionManager instance.
func NewFallDetectionManager() FallDetectionManager {
	return getFallDetectionManagerClass().New()
}



// A Boolean value that indicates whether the current device supports fall detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionManager/isAvailable

func (fc _FallDetectionManagerClass) Available() bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("available"))
	return rv
}


// Requests authorization to receive notifications about fall detection events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionManager/requestAuthorization(handler:)

func (f_ FallDetectionManager) RequestAuthorizationWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("requestAuthorizationWithHandler:"), handler)
}


// The authorization status for receiving fall detection event notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionManager/authorizationStatus

func (f_ FallDetectionManager) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](f_.ID, objc.Sel("authorizationStatus"))
	return rv
}


// A delegate that can receive notifications about fall detection events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionManager/delegate

func (f_ FallDetectionManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("delegate"))
	return rv
}


// A delegate that can receive notifications about fall detection events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionManager/delegate

func (f_ FallDetectionManager) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that indicates whether the current device supports fall detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionManager/isAvailable

func (f_ FallDetectionManager) Available() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("available"))
	return rv
}




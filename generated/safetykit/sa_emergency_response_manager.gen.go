// Code generated from Apple documentation for SafetyKit. DO NOT EDIT.

package safetykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SAEmergencyResponseManager] class.
var (
	SAEmergencyResponseManagerClass     _SAEmergencyResponseManagerClass
	SAEmergencyResponseManagerClassOnce sync.Once
)

func getSAEmergencyResponseManagerClass() _SAEmergencyResponseManagerClass {
	SAEmergencyResponseManagerClassOnce.Do(func() {
		SAEmergencyResponseManagerClass = _SAEmergencyResponseManagerClass{objc.GetClass("SAEmergencyResponseManager")}
	})
	return SAEmergencyResponseManagerClass
}

type _SAEmergencyResponseManagerClass struct {
	class objc.Class
}

// An interface definition for the [SAEmergencyResponseManager] class.
type ISAEmergencyResponseManager interface {
	objectivec.IObject
	// properties:
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	// methods:
	DialVoiceCallToPhoneNumberCompletionHandler(phoneNumber objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
}

// Provides actions in response to a Crash Detection event.
//
// Use the manager to place a voice call to an emergency contact upon receipt of a Crash Detection event. Provide an object that adopts in order to respond to the status of the voice call.


// Provides actions in response to a Crash Detection event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAEmergencyResponseManager
type SAEmergencyResponseManager struct {
	objectivec.Object
}

// SAEmergencyResponseManagerFrom constructs a [SAEmergencyResponseManager] from an unsafe.Pointer.
//
// Provides actions in response to a Crash Detection event.
func SAEmergencyResponseManagerFrom(ptr unsafe.Pointer) SAEmergencyResponseManager {
	return SAEmergencyResponseManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SAEmergencyResponseManagerClass) Alloc() SAEmergencyResponseManager {
	rv := objc.Send[SAEmergencyResponseManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SAEmergencyResponseManagerClass) New() SAEmergencyResponseManager {
	rv := objc.Send[SAEmergencyResponseManager](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SAEmergencyResponseManager) Init() SAEmergencyResponseManager {
	rv := objc.Send[SAEmergencyResponseManager](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SAEmergencyResponseManager) Autorelease() SAEmergencyResponseManager {
	rv := objc.Send[SAEmergencyResponseManager](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSAEmergencyResponseManager creates a new SAEmergencyResponseManager instance.
func NewSAEmergencyResponseManager() SAEmergencyResponseManager {
	return getSAEmergencyResponseManagerClass().New()
}



// Request the system to dial a voice call on behalf of someone involved in a crash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAEmergencyResponseManager/dialVoiceCall(toPhoneNumber:completionHandler:)
func (s_ SAEmergencyResponseManager) DialVoiceCallToPhoneNumberCompletionHandler(phoneNumber objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("dialVoiceCallToPhoneNumber:completionHandler:"), phoneNumber, handler)
}


// The object that receives voice call status updates and requested emergency response actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAEmergencyResponseManager/delegate
func (s_ SAEmergencyResponseManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// The object that receives voice call status updates and requested emergency response actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAEmergencyResponseManager/delegate
func (s_ SAEmergencyResponseManager) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}





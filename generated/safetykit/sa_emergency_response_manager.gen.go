// Code generated from Apple documentation for SafetyKit. DO NOT EDIT.

package safetykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SAEmergencyResponseManager */


/* debug [class_header]: Header for SAEmergencyResponseManager */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SAEmergencyResponseManager */
// An interface definition for the [SAEmergencyResponseManager] class.
type ISAEmergencyResponseManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SAEmergencyResponseManager */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SAEmergencyResponseManager */
	// methods:
	DialVoiceCallToPhoneNumberCompletionHandler(phoneNumber objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SAEmergencyResponseManager */
// Alloc allocates a new instance without initialization.
func (sc _SAEmergencyResponseManagerClass) Alloc() SAEmergencyResponseManager {
	rv := objc.Send[SAEmergencyResponseManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SAEmergencyResponseManager */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SAEmergencyResponseManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SAEmergencyResponseManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SAEmergencyResponseManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SAEmergencyResponseManager */

// Request the system to dial a voice call on behalf of someone involved in a crash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAEmergencyResponseManager/dialVoiceCall(toPhoneNumber:completionHandler:)
func (s_ SAEmergencyResponseManager) DialVoiceCallToPhoneNumberCompletionHandler(phoneNumber objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("dialVoiceCallToPhoneNumber:completionHandler:"), phoneNumber, handler)
}/* debug [instance_methods/method]: DialVoiceCallToPhoneNumberCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SAEmergencyResponseManager */

// The object that receives voice call status updates and requested emergency response actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAEmergencyResponseManager/delegate
func (s_ SAEmergencyResponseManager) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The object that receives voice call status updates and requested emergency response actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAEmergencyResponseManager/delegate
func (s_ SAEmergencyResponseManager) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SAEmergencyResponseManager */






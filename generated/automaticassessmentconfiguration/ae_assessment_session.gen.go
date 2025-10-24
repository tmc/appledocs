// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AEAssessmentSession */


/* debug [class_header]: Header for AEAssessmentSession */
// The class instance for the [AEAssessmentSession] class.
var (
	AEAssessmentSessionClass     _AEAssessmentSessionClass
	AEAssessmentSessionClassOnce sync.Once
)

func getAEAssessmentSessionClass() _AEAssessmentSessionClass {
	AEAssessmentSessionClassOnce.Do(func() {
		AEAssessmentSessionClass = _AEAssessmentSessionClass{objc.GetClass("AEAssessmentSession")}
	})
	return AEAssessmentSessionClass
}

type _AEAssessmentSessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AEAssessmentSession */
// An interface definition for the [AEAssessmentSession] class.
type IAEAssessmentSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AEAssessmentSession */
	// properties:
	Configuration() IAEAssessmentConfiguration
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Active() bool
	IsActive() bool
	SetIsActive(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AEAssessmentSession */
	// methods:
	Begin()
	End()
	UpdateToConfiguration(configuration IAEAssessmentConfiguration)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AEAssessmentSession */
// Alloc allocates a new instance without initialization.
func (ac _AEAssessmentSessionClass) Alloc() AEAssessmentSession {
	rv := objc.Send[AEAssessmentSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AEAssessmentSessionClass) New() AEAssessmentSession {
	rv := objc.Send[AEAssessmentSession](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AEAssessmentSession) Init() AEAssessmentSession {
	rv := objc.Send[AEAssessmentSession](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AEAssessmentSession) Autorelease() AEAssessmentSession {
	rv := objc.Send[AEAssessmentSession](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAEAssessmentSession creates a new AEAssessmentSession instance.
func NewAEAssessmentSession() AEAssessmentSession {
	return getAEAssessmentSessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AEAssessmentSession */
// A session that your app uses to protect an assessment.
//
// Use the class to manage an assessment session. The system allows only one active session at a time across all processes. The first session to run gets exclusive access to the system; subsequent session attempts fail until the first session ends. To create an assessment session, pass a new instance to the method. Then, provide the session with a delegate that conforms to the protocol: You can indicate exceptions to the restrictions imposed by an assessment session by setting the properties of the configuration instance, or you can use the default restrictions as shown above. The session tells its delegate about state changes during its life cycle. To start a session, call the session’s method: The method returns immediately, and the session starts disabling system features. After achieving the desired state, the session calls its delegate’s method. Only after receiving this callback is it safe to begin your assessment. Be sure to keep a strong reference to the session as long as you want it to remain active. If the system deallocates an active session, the session automatically ends. After completing an assessment and hiding all sensitive information, call the session’s method: After making the call, wait for the session to call its delegate’s method before reporting assessment completion to the user. During assessment, the session’s delegate might receive an callback to indicate a failure. If this happens, immediately stop the assessment, hide all sensitive content, and end the session. Because it might take time for your app to finalize the assessment, the session relies on your app to call the session’s method:


// A session that your app uses to protect an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession
type AEAssessmentSession struct {
	objectivec.Object
}

// AEAssessmentSessionFrom constructs a [AEAssessmentSession] from an unsafe.Pointer.
//
// A session that your app uses to protect an assessment.
func AEAssessmentSessionFrom(ptr unsafe.Pointer) AEAssessmentSession {
	return AEAssessmentSession{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AEAssessmentSession */

// Creates a new assessment session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/init(configuration:)
func NewAEAssessmentSessionWithConfiguration(configuration IAEAssessmentConfiguration) AEAssessmentSession {
	instance := getAEAssessmentSessionClass().Alloc()
	rv := objc.Send[AEAssessmentSession](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAEAssessmentSessionWithConfiguration */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AEAssessmentSession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AEAssessmentSession */

// A Boolean that indicates whether the current device or platform supports updating a session’s configuration after the session has begun.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/supportsConfigurationUpdates
func (ac _AEAssessmentSessionClass) SupportsConfigurationUpdates() bool {
	rv := objc.Send[bool](objc.ID(ac.class), objc.Sel("supportsConfigurationUpdates"))
	return rv
}/* debug [class_properties_class/property]: supportsConfigurationUpdates */

// A Boolean that indicates whether the current device or platform supports a configuration with one or more participant applications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/supportsMultipleParticipants
func (ac _AEAssessmentSessionClass) SupportsMultipleParticipants() bool {
	rv := objc.Send[bool](objc.ID(ac.class), objc.Sel("supportsMultipleParticipants"))
	return rv
}/* debug [class_properties_class/property]: supportsMultipleParticipants */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AEAssessmentSession */

// Starts an assessment session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/begin()
func (a_ AEAssessmentSession) Begin() {
	objc.Send[objc.ID](a_.ID, objc.Sel("begin"))
}/* debug [instance_methods/method]: Begin */


// Ends an assessment session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/end()
func (a_ AEAssessmentSession) End() {
	objc.Send[objc.ID](a_.ID, objc.Sel("end"))
}/* debug [instance_methods/method]: End */


// Changes the session to use the specified configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/update(to:)
func (a_ AEAssessmentSession) UpdateToConfiguration(configuration IAEAssessmentConfiguration) {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateToConfiguration:"), configuration)
}/* debug [instance_methods/method]: UpdateToConfiguration */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AEAssessmentSession */

// The current configuration of the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/configuration
func (a_ AEAssessmentSession) Configuration() IAEAssessmentConfiguration {
	rv := objc.Send[AEAssessmentConfiguration](a_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// A delegate to which the session provides state change updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/delegate
func (a_ AEAssessmentSession) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// A delegate to which the session provides state change updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/delegate
func (a_ AEAssessmentSession) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean that indicates whether an assessment session is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/isActive
func (a_ AEAssessmentSession) Active() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("active"))
	return rv
}/* debug [instance_properties/getter]: active */


// A Boolean that indicates whether the current device or platform supports updating a session’s configuration after the session has begun.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/supportsConfigurationUpdates
func (a_ AEAssessmentSession) SupportsConfigurationUpdates() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("supportsConfigurationUpdates"))
	return rv
}/* debug [instance_properties/getter]: supportsConfigurationUpdates */


// A Boolean that indicates whether the current device or platform supports a configuration with one or more participant applications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/supportsMultipleParticipants
func (a_ AEAssessmentSession) SupportsMultipleParticipants() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("supportsMultipleParticipants"))
	return rv
}/* debug [instance_properties/getter]: supportsMultipleParticipants */


// A Boolean that indicates whether an assessment session is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentsession/isactive
func (a_ AEAssessmentSession) IsActive() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isActive"))
	return rv
}/* debug [instance_properties/getter]: isActive */


// A Boolean that indicates whether an assessment session is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentsession/isactive
func (a_ AEAssessmentSession) SetIsActive(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsActive:"), value)
}/* debug [instance_properties/setter]: isActive */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AEAssessmentSession */



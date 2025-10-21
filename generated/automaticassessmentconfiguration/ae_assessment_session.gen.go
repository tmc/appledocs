// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AEAssessmentSession] class.
type IAEAssessmentSession interface {
	objectivec.IObject
	Begin()
	End()
	UpdateToConfiguration(configuration IAEAssessmentConfiguration)
}

// A session that your app uses to protect an assessment.
//
// Use the class to manage an assessment session. The system allows only one active session at a time across all processes. The first session to run gets exclusive access to the system; subsequent session attempts fail until the first session ends. To create an assessment session, pass a new instance to the method. Then, provide the session with a delegate that conforms to the protocol: You can indicate exceptions to the restrictions imposed by an assessment session by setting the properties of the configuration instance, or you can use the default restrictions as shown above. The session tells its delegate about state changes during its life cycle. To start a session, call the session’s method: The method returns immediately, and the session starts disabling system features. After achieving the desired state, the session calls its delegate’s method. Only after receiving this callback is it safe to begin your assessment. Be sure to keep a strong reference to the session as long as you want it to remain active. If the system deallocates an active session, the session automatically ends. After completing an assessment and hiding all sensitive information, call the session’s method: After making the call, wait for the session to call its delegate’s method before reporting assessment completion to the user. During assessment, the session’s delegate might receive an callback to indicate a failure. If this happens, immediately stop the assessment, hide all sensitive content, and end the session. Because it might take time for your app to finalize the assessment, the session relies on your app to call the session’s method:
//
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

// Alloc allocates a new instance without initialization.
func (ac _AEAssessmentSessionClass) Alloc() AEAssessmentSession {
	rv := objc.Send[AEAssessmentSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a new assessment session.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/init(configuration:)
func NewAEAssessmentSessionWithConfiguration(configuration IAEAssessmentConfiguration) AEAssessmentSession {
	instance := getAEAssessmentSessionClass().Alloc()
	rv := objc.Send[AEAssessmentSession](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	rv.Autorelease()
	return rv
}


// A Boolean that indicates whether the current device or platform supports updating a session’s configuration after the session has begun.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/supportsConfigurationUpdates
func (ac _AEAssessmentSessionClass) SupportsConfigurationUpdates() bool {
	rv := objc.Send[bool](objc.ID(ac.class), objc.Sel("supportsConfigurationUpdates"))
	return rv
}
// A Boolean that indicates whether the current device or platform supports a configuration with one or more participant applications.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/supportsMultipleParticipants
func (ac _AEAssessmentSessionClass) SupportsMultipleParticipants() bool {
	rv := objc.Send[bool](objc.ID(ac.class), objc.Sel("supportsMultipleParticipants"))
	return rv
}
// Starts an assessment session.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/begin()
func (a_ AEAssessmentSession) Begin() {
	objc.Send[objc.ID](a_.ID, objc.Sel("begin"))
}

// Ends an assessment session.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/end()
func (a_ AEAssessmentSession) End() {
	objc.Send[objc.ID](a_.ID, objc.Sel("end"))
}

// Changes the session to use the specified configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/update(to:)
func (a_ AEAssessmentSession) UpdateToConfiguration(configuration IAEAssessmentConfiguration) {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateToConfiguration:"), configuration)
}

// The current configuration of the session.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/configuration
func (a_ AEAssessmentSession) Configuration() AEAssessmentConfiguration {
	rv := objc.Send[AEAssessmentConfiguration](a_.ID, objc.Sel("configuration"))
	return rv
}

// A delegate to which the session provides state change updates.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/delegate
func (a_ AEAssessmentSession) Delegate() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// A delegate to which the session provides state change updates.

//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/delegate
func (a_ AEAssessmentSession) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean that indicates whether an assessment session is running.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/isActive
func (a_ AEAssessmentSession) Active() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("active"))
	return rv
}

// A Boolean that indicates whether the current device or platform supports updating a session’s configuration after the session has begun.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/supportsConfigurationUpdates
func (a_ AEAssessmentSession) SupportsConfigurationUpdates() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("supportsConfigurationUpdates"))
	return rv
}

// A Boolean that indicates whether the current device or platform supports a configuration with one or more participant applications.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/supportsMultipleParticipants
func (a_ AEAssessmentSession) SupportsMultipleParticipants() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("supportsMultipleParticipants"))
	return rv
}

// A Boolean that indicates whether an assessment session is running.
//
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentsession/isactive
func (a_ AEAssessmentSession) IsActive() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isActive"))
	return rv
}


// SetIsActive sets the value of the isActive property.
// A Boolean that indicates whether an assessment session is running.

//
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentsession/isactive
func (a_ AEAssessmentSession) SetIsActive(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsActive:"), value)
}



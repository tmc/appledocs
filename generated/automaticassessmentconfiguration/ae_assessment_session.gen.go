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
	// properties:
	Configuration() IAEAssessmentConfiguration
	SetConfiguration(value IAEAssessmentConfiguration)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	IsActive() bool
	SetIsActive(value bool)
	// methods:
	Begin()
	End()
	UpdateToConfiguration(configuration IAEAssessmentConfiguration)
}

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



// Starts an assessment session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/begin()
func (a_ AEAssessmentSession) Begin() {
	objc.Send[objc.ID](a_.ID, objc.Sel("begin"))
}


// Ends an assessment session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/end()
func (a_ AEAssessmentSession) End() {
	objc.Send[objc.ID](a_.ID, objc.Sel("end"))
}


// Changes the session to use the specified configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentSession/update(to:)
func (a_ AEAssessmentSession) UpdateToConfiguration(configuration IAEAssessmentConfiguration) {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateToConfiguration:"), configuration)
}


// The current configuration of the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentsession/configuration
func (a_ AEAssessmentSession) Configuration() IAEAssessmentConfiguration {
	rv := objc.Send[AEAssessmentConfiguration](a_.ID, objc.Sel("configuration"))
	return rv
}


// The current configuration of the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentsession/configuration
func (a_ AEAssessmentSession) SetConfiguration(value IAEAssessmentConfiguration) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setConfiguration:"), value)
}


// A delegate to which the session provides state change updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentsession/delegate
func (a_ AEAssessmentSession) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}


// A delegate to which the session provides state change updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentsession/delegate
func (a_ AEAssessmentSession) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean that indicates whether an assessment session is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentsession/isactive
func (a_ AEAssessmentSession) IsActive() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isActive"))
	return rv
}


// A Boolean that indicates whether an assessment session is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentsession/isactive
func (a_ AEAssessmentSession) SetIsActive(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsActive:"), value)
}





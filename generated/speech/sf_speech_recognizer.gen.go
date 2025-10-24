// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSpeechRecognizer */


/* debug [class_header]: Header for SFSpeechRecognizer */
// The class instance for the [SFSpeechRecognizer] class.
var (
	SFSpeechRecognizerClass     _SFSpeechRecognizerClass
	SFSpeechRecognizerClassOnce sync.Once
)

func getSFSpeechRecognizerClass() _SFSpeechRecognizerClass {
	SFSpeechRecognizerClassOnce.Do(func() {
		SFSpeechRecognizerClass = _SFSpeechRecognizerClass{objc.GetClass("SFSpeechRecognizer")}
	})
	return SFSpeechRecognizerClass
}

type _SFSpeechRecognizerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSpeechRecognizer */
// An interface definition for the [SFSpeechRecognizer] class.
type ISFSpeechRecognizer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFSpeechRecognizer */
	// properties:
	DefaultTaskHint() SFSpeechRecognitionTaskHint
	SetDefaultTaskHint(value SFSpeechRecognitionTaskHint)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Available() bool
	Locale() foundation.Locale
	Queue() foundation.OperationQueue
	SetQueue(value foundation.OperationQueue)
	SupportsOnDeviceRecognition() bool
	SetSupportsOnDeviceRecognition(value bool)
	IsAvailable() bool
	SetIsAvailable(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSpeechRecognizer */
	// methods:
	RecognitionTaskWithRequestDelegate(request ISFSpeechRecognitionRequest, delegate unsafe.Pointer) ISFSpeechRecognitionTask
	RecognitionTaskWithRequestResultHandler(request ISFSpeechRecognitionRequest, resultHandler unsafe.Pointer) ISFSpeechRecognitionTask
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSpeechRecognizer */
// Alloc allocates a new instance without initialization.
func (sc _SFSpeechRecognizerClass) Alloc() SFSpeechRecognizer {
	rv := objc.Send[SFSpeechRecognizer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSpeechRecognizerClass) New() SFSpeechRecognizer {
	rv := objc.Send[SFSpeechRecognizer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSpeechRecognizer) Init() SFSpeechRecognizer {
	rv := objc.Send[SFSpeechRecognizer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSpeechRecognizer) Autorelease() SFSpeechRecognizer {
	rv := objc.Send[SFSpeechRecognizer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechRecognizer creates a new SFSpeechRecognizer instance.
func NewSFSpeechRecognizer() SFSpeechRecognizer {
	return getSFSpeechRecognizerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSpeechRecognizer */
// An object you use to check for the availability of the speech recognition service, and to initiate the speech recognition process.
//
// An object is the central object for managing the speech recognizer process. Use this object to: Request authorization to use speech recognition services. Specify the language to use during the recognition process. Initiate new speech recognition tasks.


// An object you use to check for the availability of the speech recognition service, and to initiate the speech recognition process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer
type SFSpeechRecognizer struct {
	objectivec.Object
}

// SFSpeechRecognizerFrom constructs a [SFSpeechRecognizer] from an unsafe.Pointer.
//
// An object you use to check for the availability of the speech recognition service, and to initiate the speech recognition process.
func SFSpeechRecognizerFrom(ptr unsafe.Pointer) SFSpeechRecognizer {
	return SFSpeechRecognizer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSpeechRecognizer */

// Creates a speech recognizer associated with the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/init(locale:)
func NewSFSpeechRecognizerWithLocale(locale foundation.Locale) SFSpeechRecognizer {
	instance := getSFSpeechRecognizerClass().Alloc()
	rv := objc.Send[SFSpeechRecognizer](instance.ID, objc.Sel("initWithLocale:"), locale)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSFSpeechRecognizerWithLocale */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSpeechRecognizer */

// Returns your app’s current authorization to perform speech recognition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/authorizationStatus()
func (sc _SFSpeechRecognizerClass) AuthorizationStatus() SFSpeechRecognizerAuthorizationStatus {
	rv := objc.Send[SFSpeechRecognizerAuthorizationStatus](objc.ID(sc.class), objc.Sel("authorizationStatus"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AuthorizationStatus) */


// Asks the user to allow your app to perform speech recognition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/requestAuthorization(_:)
func (sc _SFSpeechRecognizerClass) RequestAuthorization(handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("requestAuthorization:"), handler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RequestAuthorization) */


// Returns the set of locales that are supported by the speech recognizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/supportedLocales()
func (sc _SFSpeechRecognizerClass) SupportedLocales() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("supportedLocales"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportedLocales) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSpeechRecognizer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSpeechRecognizer */

// Recognizes speech from the audio source associated with the specified request, using the specified delegate to manage the results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/recognitionTask(with:delegate:)
func (s_ SFSpeechRecognizer) RecognitionTaskWithRequestDelegate(request ISFSpeechRecognitionRequest, delegate unsafe.Pointer) ISFSpeechRecognitionTask {
	rv := objc.Send[SFSpeechRecognitionTask](s_.ID, objc.Sel("recognitionTaskWithRequest:delegate:"), request, delegate)
	return rv
}/* debug [instance_methods/method]: RecognitionTaskWithRequestDelegate */


// Executes the speech recognition request and delivers the results to the specified handler block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/recognitionTask(with:resultHandler:)
func (s_ SFSpeechRecognizer) RecognitionTaskWithRequestResultHandler(request ISFSpeechRecognitionRequest, resultHandler unsafe.Pointer) ISFSpeechRecognitionTask {
	rv := objc.Send[SFSpeechRecognitionTask](s_.ID, objc.Sel("recognitionTaskWithRequest:resultHandler:"), request, resultHandler)
	return rv
}/* debug [instance_methods/method]: RecognitionTaskWithRequestResultHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSpeechRecognizer */

// A hint that indicates the type of speech recognition being requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/defaultTaskHint
func (s_ SFSpeechRecognizer) DefaultTaskHint() SFSpeechRecognitionTaskHint {
	rv := objc.Send[SFSpeechRecognitionTaskHint](s_.ID, objc.Sel("defaultTaskHint"))
	return rv
}/* debug [instance_properties/getter]: defaultTaskHint */


// A hint that indicates the type of speech recognition being requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/defaultTaskHint
func (s_ SFSpeechRecognizer) SetDefaultTaskHint(value SFSpeechRecognitionTaskHint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDefaultTaskHint:"), value)
}/* debug [instance_properties/setter]: defaultTaskHint */


// The delegate object that handles changes to the availability of speech recognition services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/delegate
func (s_ SFSpeechRecognizer) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate object that handles changes to the availability of speech recognition services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/delegate
func (s_ SFSpeechRecognizer) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value that indicates whether the speech recognizer is currently available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/isAvailable
func (s_ SFSpeechRecognizer) Available() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("available"))
	return rv
}/* debug [instance_properties/getter]: available */


// The locale of the speech recognizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/locale
func (s_ SFSpeechRecognizer) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](s_.ID, objc.Sel("locale"))
	return rv
}/* debug [instance_properties/getter]: locale */


// The queue on which to execute recognition task handlers and delegate methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/queue
func (s_ SFSpeechRecognizer) Queue() foundation.OperationQueue {
	rv := objc.Send[foundation.OperationQueue](s_.ID, objc.Sel("queue"))
	return rv
}/* debug [instance_properties/getter]: queue */


// The queue on which to execute recognition task handlers and delegate methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/queue
func (s_ SFSpeechRecognizer) SetQueue(value foundation.OperationQueue) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setQueue:"), value)
}/* debug [instance_properties/setter]: queue */


// A Boolean value that indicates whether the speech recognizer can operate without network access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/supportsOnDeviceRecognition
func (s_ SFSpeechRecognizer) SupportsOnDeviceRecognition() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("supportsOnDeviceRecognition"))
	return rv
}/* debug [instance_properties/getter]: supportsOnDeviceRecognition */


// A Boolean value that indicates whether the speech recognizer can operate without network access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/supportsOnDeviceRecognition
func (s_ SFSpeechRecognizer) SetSupportsOnDeviceRecognition(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSupportsOnDeviceRecognition:"), value)
}/* debug [instance_properties/setter]: supportsOnDeviceRecognition */


// A Boolean value that indicates whether the speech recognizer is currently available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognizer/isavailable
func (s_ SFSpeechRecognizer) IsAvailable() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isAvailable"))
	return rv
}/* debug [instance_properties/getter]: isAvailable */


// A Boolean value that indicates whether the speech recognizer is currently available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognizer/isavailable
func (s_ SFSpeechRecognizer) SetIsAvailable(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsAvailable:"), value)
}/* debug [instance_properties/setter]: isAvailable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSpeechRecognizer */



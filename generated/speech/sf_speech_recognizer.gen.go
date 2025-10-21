// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [SFSpeechRecognizer] class.
type ISFSpeechRecognizer interface {
	objectivec.IObject
	RecognitionTaskWithRequestDelegate(request unsafe.Pointer, delegate objc.ID) unsafe.Pointer
}

// An object you use to check for the availability of the speech recognition service, and to initiate the speech recognition process.
//
// An object is the central object for managing the speech recognizer process. Use this object to: Request authorization to use speech recognition services. Specify the language to use during the recognition process. Initiate new speech recognition tasks.
//
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

// Alloc allocates a new instance without initialization.
func (sc _SFSpeechRecognizerClass) Alloc() SFSpeechRecognizer {
	rv := objc.Send[SFSpeechRecognizer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Recognizes speech from the audio source associated with the specified request, using the specified delegate to manage the results.
//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/recognitionTask(with:delegate:)
func (s_ SFSpeechRecognizer) RecognitionTaskWithRequestDelegate(request unsafe.Pointer, delegate objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("recognitionTaskWithRequest:delegate:"), request, delegate)
	return rv
}

// A Boolean value that indicates whether the speech recognizer can operate without network access.
//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/supportsOnDeviceRecognition
func (s_ SFSpeechRecognizer) SupportsOnDeviceRecognition() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("supportsOnDeviceRecognition"))
	return rv
}


// SetSupportsOnDeviceRecognition sets the value of the supportsOnDeviceRecognition property.
// A Boolean value that indicates whether the speech recognizer can operate without network access.

//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/supportsOnDeviceRecognition
func (s_ SFSpeechRecognizer) SetSupportsOnDeviceRecognition(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSupportsOnDeviceRecognition:"), value)
}

// A hint that indicates the type of speech recognition being requested.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognizer/defaulttaskhint
func (s_ SFSpeechRecognizer) DefaultTaskHint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("defaultTaskHint"))
	return rv
}


// SetDefaultTaskHint sets the value of the defaultTaskHint property.
// A hint that indicates the type of speech recognition being requested.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognizer/defaulttaskhint
func (s_ SFSpeechRecognizer) SetDefaultTaskHint(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDefaultTaskHint:"), value)
}

// The delegate object that handles changes to the availability of speech recognition services.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognizer/delegate
func (s_ SFSpeechRecognizer) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object that handles changes to the availability of speech recognition services.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognizer/delegate
func (s_ SFSpeechRecognizer) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the speech recognizer is currently available.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognizer/isavailable
func (s_ SFSpeechRecognizer) IsAvailable() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isAvailable"))
	return rv
}


// SetIsAvailable sets the value of the isAvailable property.
// A Boolean value that indicates whether the speech recognizer is currently available.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognizer/isavailable
func (s_ SFSpeechRecognizer) SetIsAvailable(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsAvailable:"), value)
}

// The locale of the speech recognizer.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognizer/locale
func (s_ SFSpeechRecognizer) Locale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("locale"))
	return rv
}


// SetLocale sets the value of the locale property.
// The locale of the speech recognizer.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognizer/locale
func (s_ SFSpeechRecognizer) SetLocale(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLocale:"), value)
}

// The queue on which to execute recognition task handlers and delegate methods.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognizer/queue
func (s_ SFSpeechRecognizer) Queue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("queue"))
	return rv
}


// SetQueue sets the value of the queue property.
// The queue on which to execute recognition task handlers and delegate methods.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognizer/queue
func (s_ SFSpeechRecognizer) SetQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setQueue:"), value)
}




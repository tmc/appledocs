// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SpeechRecognizer] class.
var speechRecognizerClass = _SpeechRecognizerClass{objc.GetClass("NSSpeechRecognizer")}

type _SpeechRecognizerClass struct {
	class objc.Class
}

// An interface definition for the [SpeechRecognizer] class.
type ISpeechRecognizer interface {
	objectivec.IObject
}

// The Cocoa interface to speech recognition in macOS. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer

type SpeechRecognizer struct {
	objectivec.Object
}

// SpeechRecognizerFrom constructs a [SpeechRecognizer] from an unsafe.Pointer.
//
// The Cocoa interface to speech recognition in macOS.
func SpeechRecognizerFrom(ptr unsafe.Pointer) SpeechRecognizer {
	return SpeechRecognizer{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (sc _SpeechRecognizerClass) Alloc() SpeechRecognizer {
	rv := objc.Send[SpeechRecognizer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SpeechRecognizerClass) New() SpeechRecognizer {
	rv := objc.Send[SpeechRecognizer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpeechRecognizer) Init() SpeechRecognizer {
	rv := objc.Send[SpeechRecognizer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpeechRecognizer) Autorelease() SpeechRecognizer {
	rv := objc.Send[SpeechRecognizer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpeechRecognizer creates a new SpeechRecognizer instance.
func NewSpeechRecognizer() SpeechRecognizer {
	return speechRecognizerClass.New()
}





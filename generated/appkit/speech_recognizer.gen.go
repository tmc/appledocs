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




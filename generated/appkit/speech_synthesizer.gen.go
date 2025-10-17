// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SpeechSynthesizer] class.
var speechSynthesizerClass = _SpeechSynthesizerClass{objc.GetClass("NSSpeechSynthesizer")}

type _SpeechSynthesizerClass struct {
	class objc.Class
}

// The Cocoa interface to speech synthesis in macOS. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer

type SpeechSynthesizer struct {
	objectivec.Object
}

// SpeechSynthesizerFrom constructs a [SpeechSynthesizer] from an unsafe.Pointer.
//
// The Cocoa interface to speech synthesis in macOS.
func SpeechSynthesizerFrom(ptr unsafe.Pointer) SpeechSynthesizer {
	return SpeechSynthesizer{objectivec.Object{objc.ID(ptr)}}
}




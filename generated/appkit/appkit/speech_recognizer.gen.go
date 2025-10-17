// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SpeechRecognizer] class.
var SpeechRecognizerClass objc.Class

func init() {
	SpeechRecognizerClass = objc.GetClass("NSSpeechRecognizer")
}

type SpeechRecognizer struct {
	objc.ID
}

func SpeechRecognizerFrom(ptr unsafe.Pointer) SpeechRecognizer {
	return SpeechRecognizer{
		ID: objc.ID(ptr),
	}
}





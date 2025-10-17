// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SpeechSynthesizer] class.
var SpeechSynthesizerClass objc.Class

func init() {
	SpeechSynthesizerClass = objc.GetClass("NSSpeechSynthesizer")
}

type SpeechSynthesizer struct {
	objc.ID
}

func SpeechSynthesizerFrom(ptr unsafe.Pointer) SpeechSynthesizer {
	return SpeechSynthesizer{
		ID: objc.ID(ptr),
	}
}





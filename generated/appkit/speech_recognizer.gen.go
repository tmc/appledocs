
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SpeechRecognizer] class.
var SpeechRecognizerClass _SpeechRecognizerClass

func init() {
	SpeechRecognizerClass = _SpeechRecognizerClass{objc.GetClass("NSSpeechRecognizer")}
}

type _SpeechRecognizerClass struct {
	objc.Class
}

// An interface definition for the [SpeechRecognizer] class.
type ISpeechRecognizer interface {
	ID() objc.ID
}

type SpeechRecognizer struct {
	id objc.ID
}

func SpeechRecognizerFrom(ptr unsafe.Pointer) SpeechRecognizer {
	return SpeechRecognizer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SpeechRecognizer) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SpeechRecognizerClass) Alloc() SpeechRecognizer {
	rv := objc.Send[SpeechRecognizer](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SpeechRecognizerClass) New() SpeechRecognizer {
	rv := objc.Send[SpeechRecognizer](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSpeechRecognizer creates and returns a new initialized instance.
func NewSpeechRecognizer() SpeechRecognizer {
	return SpeechRecognizerClass.New()
}

// Init initializes the instance.
func (s_ SpeechRecognizer) Init() SpeechRecognizer {
	rv := objc.Send[SpeechRecognizer](s_.ID(), selInit)
	return rv
}

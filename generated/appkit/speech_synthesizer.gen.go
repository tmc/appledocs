
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SpeechSynthesizer] class.
var SpeechSynthesizerClass _SpeechSynthesizerClass

func init() {
	SpeechSynthesizerClass = _SpeechSynthesizerClass{objc.GetClass("NSSpeechSynthesizer")}
}

type _SpeechSynthesizerClass struct {
	objc.Class
}

// An interface definition for the [SpeechSynthesizer] class.
type ISpeechSynthesizer interface {
	ID() objc.ID
}

type SpeechSynthesizer struct {
	id objc.ID
}

func SpeechSynthesizerFrom(ptr unsafe.Pointer) SpeechSynthesizer {
	return SpeechSynthesizer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SpeechSynthesizer) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SpeechSynthesizerClass) Alloc() SpeechSynthesizer {
	rv := objc.Send[SpeechSynthesizer](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SpeechSynthesizerClass) New() SpeechSynthesizer {
	rv := objc.Send[SpeechSynthesizer](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSpeechSynthesizer creates and returns a new initialized instance.
func NewSpeechSynthesizer() SpeechSynthesizer {
	return SpeechSynthesizerClass.New()
}

// Init initializes the instance.
func (s_ SpeechSynthesizer) Init() SpeechSynthesizer {
	rv := objc.Send[SpeechSynthesizer](s_.ID(), selInit)
	return rv
}

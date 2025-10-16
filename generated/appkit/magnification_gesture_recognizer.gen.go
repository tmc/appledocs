
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MagnificationGestureRecognizer] class.
var MagnificationGestureRecognizerClass _MagnificationGestureRecognizerClass

func init() {
	MagnificationGestureRecognizerClass = _MagnificationGestureRecognizerClass{objc.GetClass("NSMagnificationGestureRecognizer")}
}

type _MagnificationGestureRecognizerClass struct {
	objc.Class
}

// An interface definition for the [MagnificationGestureRecognizer] class.
type IMagnificationGestureRecognizer interface {
	ID() objc.ID
}

type MagnificationGestureRecognizer struct {
	id objc.ID
}

func MagnificationGestureRecognizerFrom(ptr unsafe.Pointer) MagnificationGestureRecognizer {
	return MagnificationGestureRecognizer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ MagnificationGestureRecognizer) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _MagnificationGestureRecognizerClass) Alloc() MagnificationGestureRecognizer {
	rv := objc.Send[MagnificationGestureRecognizer](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _MagnificationGestureRecognizerClass) New() MagnificationGestureRecognizer {
	rv := objc.Send[MagnificationGestureRecognizer](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewMagnificationGestureRecognizer creates and returns a new initialized instance.
func NewMagnificationGestureRecognizer() MagnificationGestureRecognizer {
	return MagnificationGestureRecognizerClass.New()
}

// Init initializes the instance.
func (m_ MagnificationGestureRecognizer) Init() MagnificationGestureRecognizer {
	rv := objc.Send[MagnificationGestureRecognizer](m_.ID(), selInit)
	return rv
}

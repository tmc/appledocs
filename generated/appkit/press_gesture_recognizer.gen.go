
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PressGestureRecognizer] class.
var PressGestureRecognizerClass _PressGestureRecognizerClass

func init() {
	PressGestureRecognizerClass = _PressGestureRecognizerClass{objc.GetClass("NSPressGestureRecognizer")}
}

type _PressGestureRecognizerClass struct {
	objc.Class
}

// An interface definition for the [PressGestureRecognizer] class.
type IPressGestureRecognizer interface {
	ID() objc.ID
}

type PressGestureRecognizer struct {
	id objc.ID
}

func PressGestureRecognizerFrom(ptr unsafe.Pointer) PressGestureRecognizer {
	return PressGestureRecognizer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PressGestureRecognizer) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PressGestureRecognizerClass) Alloc() PressGestureRecognizer {
	rv := objc.Send[PressGestureRecognizer](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PressGestureRecognizerClass) New() PressGestureRecognizer {
	rv := objc.Send[PressGestureRecognizer](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPressGestureRecognizer creates and returns a new initialized instance.
func NewPressGestureRecognizer() PressGestureRecognizer {
	return PressGestureRecognizerClass.New()
}

// Init initializes the instance.
func (p_ PressGestureRecognizer) Init() PressGestureRecognizer {
	rv := objc.Send[PressGestureRecognizer](p_.ID(), selInit)
	return rv
}

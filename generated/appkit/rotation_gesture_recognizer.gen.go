
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RotationGestureRecognizer] class.
var RotationGestureRecognizerClass _RotationGestureRecognizerClass

func init() {
	RotationGestureRecognizerClass = _RotationGestureRecognizerClass{objc.GetClass("NSRotationGestureRecognizer")}
}

type _RotationGestureRecognizerClass struct {
	objc.Class
}

// An interface definition for the [RotationGestureRecognizer] class.
type IRotationGestureRecognizer interface {
	ID() objc.ID
}

type RotationGestureRecognizer struct {
	id objc.ID
}

func RotationGestureRecognizerFrom(ptr unsafe.Pointer) RotationGestureRecognizer {
	return RotationGestureRecognizer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ RotationGestureRecognizer) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _RotationGestureRecognizerClass) Alloc() RotationGestureRecognizer {
	rv := objc.Send[RotationGestureRecognizer](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _RotationGestureRecognizerClass) New() RotationGestureRecognizer {
	rv := objc.Send[RotationGestureRecognizer](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewRotationGestureRecognizer creates and returns a new initialized instance.
func NewRotationGestureRecognizer() RotationGestureRecognizer {
	return RotationGestureRecognizerClass.New()
}

// Init initializes the instance.
func (r_ RotationGestureRecognizer) Init() RotationGestureRecognizer {
	rv := objc.Send[RotationGestureRecognizer](r_.ID(), selInit)
	return rv
}

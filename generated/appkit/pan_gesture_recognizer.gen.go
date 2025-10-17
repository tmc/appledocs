
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PanGestureRecognizer] class.
var PanGestureRecognizerClass _PanGestureRecognizerClass

func init() {
	PanGestureRecognizerClass = _PanGestureRecognizerClass{objc.GetClass("NSPanGestureRecognizer")}
}

type _PanGestureRecognizerClass struct {
	objc.Class
}

// An interface definition for the [PanGestureRecognizer] class.
type IPanGestureRecognizer interface {
	ID() objc.ID
}

type PanGestureRecognizer struct {
	id objc.ID
}

func PanGestureRecognizerFrom(ptr unsafe.Pointer) PanGestureRecognizer {
	return PanGestureRecognizer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PanGestureRecognizer) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PanGestureRecognizerClass) Alloc() PanGestureRecognizer {
	rv := objc.Send[PanGestureRecognizer](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PanGestureRecognizerClass) New() PanGestureRecognizer {
	rv := objc.Send[PanGestureRecognizer](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPanGestureRecognizer creates and returns a new initialized instance.
func NewPanGestureRecognizer() PanGestureRecognizer {
	return PanGestureRecognizerClass.New()
}

// Init initializes the instance.
func (p_ PanGestureRecognizer) Init() PanGestureRecognizer {
	rv := objc.Send[PanGestureRecognizer](p_.ID(), selInit)
	return rv
}

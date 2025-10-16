
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ClickGestureRecognizer] class.
var ClickGestureRecognizerClass _ClickGestureRecognizerClass

func init() {
	ClickGestureRecognizerClass = _ClickGestureRecognizerClass{objc.GetClass("NSClickGestureRecognizer")}
}

type _ClickGestureRecognizerClass struct {
	objc.Class
}

// An interface definition for the [ClickGestureRecognizer] class.
type IClickGestureRecognizer interface {
	ID() objc.ID
}

type ClickGestureRecognizer struct {
	id objc.ID
}

func ClickGestureRecognizerFrom(ptr unsafe.Pointer) ClickGestureRecognizer {
	return ClickGestureRecognizer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ ClickGestureRecognizer) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _ClickGestureRecognizerClass) Alloc() ClickGestureRecognizer {
	rv := objc.Send[ClickGestureRecognizer](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _ClickGestureRecognizerClass) New() ClickGestureRecognizer {
	rv := objc.Send[ClickGestureRecognizer](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewClickGestureRecognizer creates and returns a new initialized instance.
func NewClickGestureRecognizer() ClickGestureRecognizer {
	return ClickGestureRecognizerClass.New()
}

// Init initializes the instance.
func (c_ ClickGestureRecognizer) Init() ClickGestureRecognizer {
	rv := objc.Send[ClickGestureRecognizer](c_.ID(), selInit)
	return rv
}

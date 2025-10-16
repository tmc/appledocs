
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [sendsActionOnEndEditing] class.
var sendsActionOnEndEditingClass _sendsActionOnEndEditingClass

func init() {
	sendsActionOnEndEditingClass = _sendsActionOnEndEditingClass{objc.GetClass("sendsActionOnEndEditing")}
}

type _sendsActionOnEndEditingClass struct {
	objc.Class
}

// An interface definition for the [sendsActionOnEndEditing] class.
type IsendsActionOnEndEditing interface {
	ID() objc.ID
}

type sendsActionOnEndEditing struct {
	id objc.ID
}

func sendsActionOnEndEditingFrom(ptr unsafe.Pointer) sendsActionOnEndEditing {
	return sendsActionOnEndEditing{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ sendsActionOnEndEditing) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _sendsActionOnEndEditingClass) Alloc() sendsActionOnEndEditing {
	rv := objc.Send[sendsActionOnEndEditing](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _sendsActionOnEndEditingClass) New() sendsActionOnEndEditing {
	rv := objc.Send[sendsActionOnEndEditing](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewsendsActionOnEndEditing creates and returns a new initialized instance.
func NewsendsActionOnEndEditing() sendsActionOnEndEditing {
	return sendsActionOnEndEditingClass.New()
}

// Init initializes the instance.
func (s_ sendsActionOnEndEditing) Init() sendsActionOnEndEditing {
	rv := objc.Send[sendsActionOnEndEditing](s_.ID(), selInit)
	return rv
}

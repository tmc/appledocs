
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [abortEditing] class.
var abortEditingClass _abortEditingClass

func init() {
	abortEditingClass = _abortEditingClass{objc.GetClass("abortEditing")}
}

type _abortEditingClass struct {
	objc.Class
}

// An interface definition for the [abortEditing] class.
type IabortEditing interface {
	ID() objc.ID
}

type abortEditing struct {
	id objc.ID
}

func abortEditingFrom(ptr unsafe.Pointer) abortEditing {
	return abortEditing{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ abortEditing) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _abortEditingClass) Alloc() abortEditing {
	rv := objc.Send[abortEditing](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _abortEditingClass) New() abortEditing {
	rv := objc.Send[abortEditing](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewabortEditing creates and returns a new initialized instance.
func NewabortEditing() abortEditing {
	return abortEditingClass.New()
}

// Init initializes the instance.
func (a_ abortEditing) Init() abortEditing {
	rv := objc.Send[abortEditing](a_.ID(), selInit)
	return rv
}

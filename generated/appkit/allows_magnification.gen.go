
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsMagnification] class.
var allowsMagnificationClass _allowsMagnificationClass

func init() {
	allowsMagnificationClass = _allowsMagnificationClass{objc.GetClass("allowsMagnification")}
}

type _allowsMagnificationClass struct {
	objc.Class
}

// An interface definition for the [allowsMagnification] class.
type IallowsMagnification interface {
	ID() objc.ID
}

type allowsMagnification struct {
	id objc.ID
}

func allowsMagnificationFrom(ptr unsafe.Pointer) allowsMagnification {
	return allowsMagnification{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsMagnification) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsMagnificationClass) Alloc() allowsMagnification {
	rv := objc.Send[allowsMagnification](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsMagnificationClass) New() allowsMagnification {
	rv := objc.Send[allowsMagnification](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsMagnification creates and returns a new initialized instance.
func NewallowsMagnification() allowsMagnification {
	return allowsMagnificationClass.New()
}

// Init initializes the instance.
func (a_ allowsMagnification) Init() allowsMagnification {
	rv := objc.Send[allowsMagnification](a_.ID(), selInit)
	return rv
}

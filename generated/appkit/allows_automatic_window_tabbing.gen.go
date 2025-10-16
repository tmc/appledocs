
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsAutomaticWindowTabbing] class.
var allowsAutomaticWindowTabbingClass _allowsAutomaticWindowTabbingClass

func init() {
	allowsAutomaticWindowTabbingClass = _allowsAutomaticWindowTabbingClass{objc.GetClass("allowsAutomaticWindowTabbing")}
}

type _allowsAutomaticWindowTabbingClass struct {
	objc.Class
}

// An interface definition for the [allowsAutomaticWindowTabbing] class.
type IallowsAutomaticWindowTabbing interface {
	ID() objc.ID
}

type allowsAutomaticWindowTabbing struct {
	id objc.ID
}

func allowsAutomaticWindowTabbingFrom(ptr unsafe.Pointer) allowsAutomaticWindowTabbing {
	return allowsAutomaticWindowTabbing{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsAutomaticWindowTabbing) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsAutomaticWindowTabbingClass) Alloc() allowsAutomaticWindowTabbing {
	rv := objc.Send[allowsAutomaticWindowTabbing](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsAutomaticWindowTabbingClass) New() allowsAutomaticWindowTabbing {
	rv := objc.Send[allowsAutomaticWindowTabbing](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsAutomaticWindowTabbing creates and returns a new initialized instance.
func NewallowsAutomaticWindowTabbing() allowsAutomaticWindowTabbing {
	return allowsAutomaticWindowTabbingClass.New()
}

// Init initializes the instance.
func (a_ allowsAutomaticWindowTabbing) Init() allowsAutomaticWindowTabbing {
	rv := objc.Send[allowsAutomaticWindowTabbing](a_.ID(), selInit)
	return rv
}

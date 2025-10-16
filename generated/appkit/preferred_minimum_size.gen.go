
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [preferredMinimumSize] class.
var preferredMinimumSizeClass _preferredMinimumSizeClass

func init() {
	preferredMinimumSizeClass = _preferredMinimumSizeClass{objc.GetClass("preferredMinimumSize")}
}

type _preferredMinimumSizeClass struct {
	objc.Class
}

// An interface definition for the [preferredMinimumSize] class.
type IpreferredMinimumSize interface {
	ID() objc.ID
}

type preferredMinimumSize struct {
	id objc.ID
}

func preferredMinimumSizeFrom(ptr unsafe.Pointer) preferredMinimumSize {
	return preferredMinimumSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ preferredMinimumSize) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _preferredMinimumSizeClass) Alloc() preferredMinimumSize {
	rv := objc.Send[preferredMinimumSize](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _preferredMinimumSizeClass) New() preferredMinimumSize {
	rv := objc.Send[preferredMinimumSize](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpreferredMinimumSize creates and returns a new initialized instance.
func NewpreferredMinimumSize() preferredMinimumSize {
	return preferredMinimumSizeClass.New()
}

// Init initializes the instance.
func (p_ preferredMinimumSize) Init() preferredMinimumSize {
	rv := objc.Send[preferredMinimumSize](p_.ID(), selInit)
	return rv
}

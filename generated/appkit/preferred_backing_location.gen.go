
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [preferredBackingLocation] class.
var preferredBackingLocationClass _preferredBackingLocationClass

func init() {
	preferredBackingLocationClass = _preferredBackingLocationClass{objc.GetClass("preferredBackingLocation")}
}

type _preferredBackingLocationClass struct {
	objc.Class
}

// An interface definition for the [preferredBackingLocation] class.
type IpreferredBackingLocation interface {
	ID() objc.ID
}

type preferredBackingLocation struct {
	id objc.ID
}

func preferredBackingLocationFrom(ptr unsafe.Pointer) preferredBackingLocation {
	return preferredBackingLocation{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ preferredBackingLocation) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _preferredBackingLocationClass) Alloc() preferredBackingLocation {
	rv := objc.Send[preferredBackingLocation](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _preferredBackingLocationClass) New() preferredBackingLocation {
	rv := objc.Send[preferredBackingLocation](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpreferredBackingLocation creates and returns a new initialized instance.
func NewpreferredBackingLocation() preferredBackingLocation {
	return preferredBackingLocationClass.New()
}

// Init initializes the instance.
func (p_ preferredBackingLocation) Init() preferredBackingLocation {
	rv := objc.Send[preferredBackingLocation](p_.ID(), selInit)
	return rv
}

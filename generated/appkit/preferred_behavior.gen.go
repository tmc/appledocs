
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [preferredBehavior] class.
var preferredBehaviorClass _preferredBehaviorClass

func init() {
	preferredBehaviorClass = _preferredBehaviorClass{objc.GetClass("preferredBehavior")}
}

type _preferredBehaviorClass struct {
	objc.Class
}

// An interface definition for the [preferredBehavior] class.
type IpreferredBehavior interface {
	ID() objc.ID
}

type preferredBehavior struct {
	id objc.ID
}

func preferredBehaviorFrom(ptr unsafe.Pointer) preferredBehavior {
	return preferredBehavior{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ preferredBehavior) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _preferredBehaviorClass) Alloc() preferredBehavior {
	rv := objc.Send[preferredBehavior](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _preferredBehaviorClass) New() preferredBehavior {
	rv := objc.Send[preferredBehavior](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpreferredBehavior creates and returns a new initialized instance.
func NewpreferredBehavior() preferredBehavior {
	return preferredBehaviorClass.New()
}

// Init initializes the instance.
func (p_ preferredBehavior) Init() preferredBehavior {
	rv := objc.Send[preferredBehavior](p_.ID(), selInit)
	return rv
}

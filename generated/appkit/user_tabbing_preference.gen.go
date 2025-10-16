
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [userTabbingPreference] class.
var userTabbingPreferenceClass _userTabbingPreferenceClass

func init() {
	userTabbingPreferenceClass = _userTabbingPreferenceClass{objc.GetClass("userTabbingPreference")}
}

type _userTabbingPreferenceClass struct {
	objc.Class
}

// An interface definition for the [userTabbingPreference] class.
type IuserTabbingPreference interface {
	ID() objc.ID
}

type userTabbingPreference struct {
	id objc.ID
}

func userTabbingPreferenceFrom(ptr unsafe.Pointer) userTabbingPreference {
	return userTabbingPreference{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ userTabbingPreference) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _userTabbingPreferenceClass) Alloc() userTabbingPreference {
	rv := objc.Send[userTabbingPreference](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _userTabbingPreferenceClass) New() userTabbingPreference {
	rv := objc.Send[userTabbingPreference](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewuserTabbingPreference creates and returns a new initialized instance.
func NewuserTabbingPreference() userTabbingPreference {
	return userTabbingPreferenceClass.New()
}

// Init initializes the instance.
func (u_ userTabbingPreference) Init() userTabbingPreference {
	rv := objc.Send[userTabbingPreference](u_.ID(), selInit)
	return rv
}

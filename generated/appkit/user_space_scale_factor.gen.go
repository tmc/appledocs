
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [userSpaceScaleFactor] class.
var userSpaceScaleFactorClass _userSpaceScaleFactorClass

func init() {
	userSpaceScaleFactorClass = _userSpaceScaleFactorClass{objc.GetClass("userSpaceScaleFactor")}
}

type _userSpaceScaleFactorClass struct {
	objc.Class
}

// An interface definition for the [userSpaceScaleFactor] class.
type IuserSpaceScaleFactor interface {
	ID() objc.ID
}

type userSpaceScaleFactor struct {
	id objc.ID
}

func userSpaceScaleFactorFrom(ptr unsafe.Pointer) userSpaceScaleFactor {
	return userSpaceScaleFactor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ userSpaceScaleFactor) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _userSpaceScaleFactorClass) Alloc() userSpaceScaleFactor {
	rv := objc.Send[userSpaceScaleFactor](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _userSpaceScaleFactorClass) New() userSpaceScaleFactor {
	rv := objc.Send[userSpaceScaleFactor](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewuserSpaceScaleFactor creates and returns a new initialized instance.
func NewuserSpaceScaleFactor() userSpaceScaleFactor {
	return userSpaceScaleFactorClass.New()
}

// Init initializes the instance.
func (u_ userSpaceScaleFactor) Init() userSpaceScaleFactor {
	rv := objc.Send[userSpaceScaleFactor](u_.ID(), selInit)
	return rv
}

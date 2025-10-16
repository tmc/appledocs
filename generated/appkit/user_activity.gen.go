
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [userActivity] class.
var userActivityClass _userActivityClass

func init() {
	userActivityClass = _userActivityClass{objc.GetClass("userActivity")}
}

type _userActivityClass struct {
	objc.Class
}

// An interface definition for the [userActivity] class.
type IuserActivity interface {
	ID() objc.ID
}

type userActivity struct {
	id objc.ID
}

func userActivityFrom(ptr unsafe.Pointer) userActivity {
	return userActivity{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ userActivity) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _userActivityClass) Alloc() userActivity {
	rv := objc.Send[userActivity](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _userActivityClass) New() userActivity {
	rv := objc.Send[userActivity](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewuserActivity creates and returns a new initialized instance.
func NewuserActivity() userActivity {
	return userActivityClass.New()
}

// Init initializes the instance.
func (u_ userActivity) Init() userActivity {
	rv := objc.Send[userActivity](u_.ID(), selInit)
	return rv
}

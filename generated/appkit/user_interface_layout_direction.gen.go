
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [userInterfaceLayoutDirection] class.
var userInterfaceLayoutDirectionClass _userInterfaceLayoutDirectionClass

func init() {
	userInterfaceLayoutDirectionClass = _userInterfaceLayoutDirectionClass{objc.GetClass("userInterfaceLayoutDirection")}
}

type _userInterfaceLayoutDirectionClass struct {
	objc.Class
}

// An interface definition for the [userInterfaceLayoutDirection] class.
type IuserInterfaceLayoutDirection interface {
	ID() objc.ID
}

type userInterfaceLayoutDirection struct {
	id objc.ID
}

func userInterfaceLayoutDirectionFrom(ptr unsafe.Pointer) userInterfaceLayoutDirection {
	return userInterfaceLayoutDirection{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ userInterfaceLayoutDirection) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _userInterfaceLayoutDirectionClass) Alloc() userInterfaceLayoutDirection {
	rv := objc.Send[userInterfaceLayoutDirection](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _userInterfaceLayoutDirectionClass) New() userInterfaceLayoutDirection {
	rv := objc.Send[userInterfaceLayoutDirection](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewuserInterfaceLayoutDirection creates and returns a new initialized instance.
func NewuserInterfaceLayoutDirection() userInterfaceLayoutDirection {
	return userInterfaceLayoutDirectionClass.New()
}

// Init initializes the instance.
func (u_ userInterfaceLayoutDirection) Init() userInterfaceLayoutDirection {
	rv := objc.Send[userInterfaceLayoutDirection](u_.ID(), selInit)
	return rv
}

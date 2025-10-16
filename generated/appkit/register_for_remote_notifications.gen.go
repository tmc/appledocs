
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [registerForRemoteNotifications] class.
var registerForRemoteNotificationsClass _registerForRemoteNotificationsClass

func init() {
	registerForRemoteNotificationsClass = _registerForRemoteNotificationsClass{objc.GetClass("registerForRemoteNotifications")}
}

type _registerForRemoteNotificationsClass struct {
	objc.Class
}

// An interface definition for the [registerForRemoteNotifications] class.
type IregisterForRemoteNotifications interface {
	ID() objc.ID
}

type registerForRemoteNotifications struct {
	id objc.ID
}

func registerForRemoteNotificationsFrom(ptr unsafe.Pointer) registerForRemoteNotifications {
	return registerForRemoteNotifications{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ registerForRemoteNotifications) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _registerForRemoteNotificationsClass) Alloc() registerForRemoteNotifications {
	rv := objc.Send[registerForRemoteNotifications](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _registerForRemoteNotificationsClass) New() registerForRemoteNotifications {
	rv := objc.Send[registerForRemoteNotifications](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewregisterForRemoteNotifications creates and returns a new initialized instance.
func NewregisterForRemoteNotifications() registerForRemoteNotifications {
	return registerForRemoteNotificationsClass.New()
}

// Init initializes the instance.
func (r_ registerForRemoteNotifications) Init() registerForRemoteNotifications {
	rv := objc.Send[registerForRemoteNotifications](r_.ID(), selInit)
	return rv
}

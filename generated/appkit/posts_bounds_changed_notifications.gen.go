
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [postsBoundsChangedNotifications] class.
var postsBoundsChangedNotificationsClass _postsBoundsChangedNotificationsClass

func init() {
	postsBoundsChangedNotificationsClass = _postsBoundsChangedNotificationsClass{objc.GetClass("postsBoundsChangedNotifications")}
}

type _postsBoundsChangedNotificationsClass struct {
	objc.Class
}

// An interface definition for the [postsBoundsChangedNotifications] class.
type IpostsBoundsChangedNotifications interface {
	ID() objc.ID
}

type postsBoundsChangedNotifications struct {
	id objc.ID
}

func postsBoundsChangedNotificationsFrom(ptr unsafe.Pointer) postsBoundsChangedNotifications {
	return postsBoundsChangedNotifications{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ postsBoundsChangedNotifications) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _postsBoundsChangedNotificationsClass) Alloc() postsBoundsChangedNotifications {
	rv := objc.Send[postsBoundsChangedNotifications](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _postsBoundsChangedNotificationsClass) New() postsBoundsChangedNotifications {
	rv := objc.Send[postsBoundsChangedNotifications](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpostsBoundsChangedNotifications creates and returns a new initialized instance.
func NewpostsBoundsChangedNotifications() postsBoundsChangedNotifications {
	return postsBoundsChangedNotificationsClass.New()
}

// Init initializes the instance.
func (p_ postsBoundsChangedNotifications) Init() postsBoundsChangedNotifications {
	rv := objc.Send[postsBoundsChangedNotifications](p_.ID(), selInit)
	return rv
}

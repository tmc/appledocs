
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [postsFrameChangedNotifications] class.
var postsFrameChangedNotificationsClass _postsFrameChangedNotificationsClass

func init() {
	postsFrameChangedNotificationsClass = _postsFrameChangedNotificationsClass{objc.GetClass("postsFrameChangedNotifications")}
}

type _postsFrameChangedNotificationsClass struct {
	objc.Class
}

// An interface definition for the [postsFrameChangedNotifications] class.
type IpostsFrameChangedNotifications interface {
	ID() objc.ID
}

type postsFrameChangedNotifications struct {
	id objc.ID
}

func postsFrameChangedNotificationsFrom(ptr unsafe.Pointer) postsFrameChangedNotifications {
	return postsFrameChangedNotifications{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ postsFrameChangedNotifications) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _postsFrameChangedNotificationsClass) Alloc() postsFrameChangedNotifications {
	rv := objc.Send[postsFrameChangedNotifications](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _postsFrameChangedNotificationsClass) New() postsFrameChangedNotifications {
	rv := objc.Send[postsFrameChangedNotifications](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpostsFrameChangedNotifications creates and returns a new initialized instance.
func NewpostsFrameChangedNotifications() postsFrameChangedNotifications {
	return postsFrameChangedNotificationsClass.New()
}

// Init initializes the instance.
func (p_ postsFrameChangedNotifications) Init() postsFrameChangedNotifications {
	rv := objc.Send[postsFrameChangedNotifications](p_.ID(), selInit)
	return rv
}

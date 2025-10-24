// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NotificationCenter] class.
var (
	NotificationCenterClass     _NotificationCenterClass
	NotificationCenterClassOnce sync.Once
)

func getNotificationCenterClass() _NotificationCenterClass {
	NotificationCenterClassOnce.Do(func() {
		NotificationCenterClass = _NotificationCenterClass{objc.GetClass("NSNotificationCenter")}
	})
	return NotificationCenterClass
}

type _NotificationCenterClass struct {
	class objc.Class
}

// An interface definition for the [NotificationCenter] class.
type INotificationCenter interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A notification dispatch mechanism that enables the broadcast of information to registered observers.
//
// Callers register with a notification center to receive one or both of the following: objects, when working in Objective-C or with frameworks that only support . Objects register with a notification center to receive notifications ( objects) using the or methods, specifying a notification name and optionally a source object. When a caller adds itself as an observer, it specifies which notifications it should receive. and instances for use with Swift code, providing strong typing, appropriate actor isolation, and a more idiomatic Swift experience. Callers register with the notification center using the various flavors of the method, specifying either a message type or a convenience to identify the notification messages to receive. See for more information about this API. Callers may add observers for many different notifications, or even the same notification name or message type as produced by different source objects. Each running app has a notification center, and you can create new notification centers to organize communications in particular contexts. A notification center can deliver notifications only within a single program. On macOS, if you want to post a notification to other processes or receive notifications from other processes, use instead.


// A notification dispatch mechanism that enables the broadcast of information to registered observers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter
type NotificationCenter struct {
	objectivec.Object
}

// NotificationCenterFrom constructs a [NotificationCenter] from an unsafe.Pointer.
//
// A notification dispatch mechanism that enables the broadcast of information to registered observers.
func NotificationCenterFrom(ptr unsafe.Pointer) NotificationCenter {
	return NotificationCenter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NotificationCenterClass) Alloc() NotificationCenter {
	rv := objc.Send[NotificationCenter](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NotificationCenterClass) New() NotificationCenter {
	rv := objc.Send[NotificationCenter](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NotificationCenter) Init() NotificationCenter {
	rv := objc.Send[NotificationCenter](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NotificationCenter) Autorelease() NotificationCenter {
	rv := objc.Send[NotificationCenter](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNotificationCenter creates a new NotificationCenter instance.
func NewNotificationCenter() NotificationCenter {
	return getNotificationCenterClass().New()
}





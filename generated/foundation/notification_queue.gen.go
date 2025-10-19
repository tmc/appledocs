// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NotificationQueue] class.
var notificationQueueClass = _NotificationQueueClass{objc.GetClass("NSNotificationQueue")}

type _NotificationQueueClass struct {
	class objc.Class
}

// An interface definition for the [NotificationQueue] class.
type INotificationQueue interface {
	objectivec.IObject
}

// A notification center buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue

type NotificationQueue struct {
	objectivec.Object
}

// NotificationQueueFrom constructs a [NotificationQueue] from an unsafe.Pointer.
//
// A notification center buffer.
func NotificationQueueFrom(ptr unsafe.Pointer) NotificationQueue {
	return NotificationQueue{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _NotificationQueueClass) Alloc() NotificationQueue {
	rv := objc.Send[NotificationQueue](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NotificationQueueClass) New() NotificationQueue {
	rv := objc.Send[NotificationQueue](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NotificationQueue) Init() NotificationQueue {
	rv := objc.Send[NotificationQueue](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NotificationQueue) Autorelease() NotificationQueue {
	rv := objc.Send[NotificationQueue](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNotificationQueue creates a new NotificationQueue instance.
func NewNotificationQueue() NotificationQueue {
	return notificationQueueClass.New()
}





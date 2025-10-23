// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NotificationQueue] class.
var (
	NotificationQueueClass     _NotificationQueueClass
	NotificationQueueClassOnce sync.Once
)

func getNotificationQueueClass() _NotificationQueueClass {
	NotificationQueueClassOnce.Do(func() {
		NotificationQueueClass = _NotificationQueueClass{objc.GetClass("NSNotificationQueue")}
	})
	return NotificationQueueClass
}

type _NotificationQueueClass struct {
	class objc.Class
}

// An interface definition for the [NotificationQueue] class.
type INotificationQueue interface {
	objectivec.IObject
	DequeueNotificationsMatchingCoalesceMask(notification INotification, coalesceMask uint)
	EnqueueNotificationPostingStyle(notification INotification, postingStyle NSPostingStyle)
	EnqueueNotificationPostingStyleCoalesceMaskForModes(notification INotification, postingStyle NSPostingStyle, coalesceMask NSNotificationCoalescing, modes []string)
}

// A notification center buffer.
//
// Whereas a notification center distributes notifications when posted, notifications placed into the queue can be delayed until the end of the current pass through the run loop or until the run loop is idle. Duplicate notifications can be coalesced so that only one notification is sent although multiple notifications are posted. A notification queue maintains notifications in first in, first out (FIFO) order. When a notification moves to the front of the queue, the queue posts it to the notification center, which in turn dispatches the notification to all objects registered as observers. Every thread has a default notification queue, which is associated with the default notification center for the process. You can create your own notification queues and have multiple queues per center and thread.


// A notification center buffer.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getNotificationQueueClass().New()
}



// Initializes and returns a notification queue for the specified notification center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/init(notificationCenter:)
func NewNotificationQueueWithNotificationCenter(notificationCenter INotificationCenter) NotificationQueue {
	instance := getNotificationQueueClass().Alloc()
	rv := objc.Send[NotificationQueue](instance.ID, objc.Sel("initWithNotificationCenter:"), notificationCenter)
	rv.Autorelease()
	return rv
}



// Returns the default notification queue for the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/default
func (nc _NotificationQueueClass) DefaultQueue() NotificationQueue {
	rv := objc.Send[NotificationQueue](objc.ID(nc.class), objc.Sel("defaultQueue"))
	return rv
}

// Removes all notifications from the queue that match a provided notification using provided matching criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/dequeueNotifications(matching:coalesceMask:)
func (n_ NotificationQueue) DequeueNotificationsMatchingCoalesceMask(notification INotification, coalesceMask uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("dequeueNotificationsMatching:coalesceMask:"), notification, coalesceMask)
}


// Adds a notification to the notification queue with a specified posting style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/enqueue(_:postingStyle:)
func (n_ NotificationQueue) EnqueueNotificationPostingStyle(notification INotification, postingStyle NSPostingStyle) {
	objc.Send[objc.ID](n_.ID, objc.Sel("enqueueNotification:postingStyle:"), notification, postingStyle)
}


// Adds a notification to the notification queue with a specified posting style, criteria for coalescing, and run loop mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/enqueue(_:postingStyle:coalesceMask:forModes:)
func (n_ NotificationQueue) EnqueueNotificationPostingStyleCoalesceMaskForModes(notification INotification, postingStyle NSPostingStyle, coalesceMask NSNotificationCoalescing, modes []string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("enqueueNotification:postingStyle:coalesceMask:forModes:"), notification, postingStyle, coalesceMask, modes)
}


// Returns the default notification queue for the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationQueue/default
func (n_ NotificationQueue) DefaultQueue() INotificationQueue {
	rv := objc.Send[NotificationQueue](n_.ID, objc.Sel("defaultQueue"))
	return rv
}



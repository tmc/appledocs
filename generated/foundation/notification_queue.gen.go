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




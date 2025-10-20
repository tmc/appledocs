// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var notificationQueueClass _NotificationQueueClass

func init() {
	notificationQueueClass = _NotificationQueueClass{objc.GetClass("NSNotificationQueue")}
}

type _NotificationQueueClass struct {
	class objc.Class
}

type NotificationQueue struct {
	objc.ID
}

func NotificationQueueFrom(ptr unsafe.Pointer) NotificationQueue {
	return NotificationQueue{
		ID: objc.ID(ptr),
	}
}





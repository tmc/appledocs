// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [NotificationQueue] class.
var NotificationQueueClass objc.Class

func init() {
	NotificationQueueClass = objc.GetClass("NSNotificationQueue")
}

type NotificationQueue struct {
	objc.ID
}

func NotificationQueueFrom(ptr unsafe.Pointer) NotificationQueue {
	return NotificationQueue{
		ID: objc.ID(ptr),
	}
}





// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DistributedNotificationCenter] class.
var distributedNotificationCenterClass = _DistributedNotificationCenterClass{objc.GetClass("NSDistributedNotificationCenter")}

type _DistributedNotificationCenterClass struct {
	class objc.Class
}

// An interface definition for the [DistributedNotificationCenter] class.
type IDistributedNotificationCenter interface {
	INotificationCenter
}

// A notification dispatch mechanism that enables the broadcast of notifications across task boundaries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter

type DistributedNotificationCenter struct {
	NotificationCenter
}

// DistributedNotificationCenterFrom constructs a [DistributedNotificationCenter] from an unsafe.Pointer.
//
// A notification dispatch mechanism that enables the broadcast of notifications across task boundaries.
func DistributedNotificationCenterFrom(ptr unsafe.Pointer) DistributedNotificationCenter {
	return DistributedNotificationCenter{
		NotificationCenter: NotificationCenterFrom(ptr),
	}
}




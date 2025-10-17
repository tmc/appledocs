// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DistributedNotificationCenter] class.
var DistributedNotificationCenterClass = _DistributedNotificationCenterClass{objc.GetClass("NSDistributedNotificationCenter")}

type _DistributedNotificationCenterClass struct {
	class objc.Class
}

type DistributedNotificationCenter struct {
	objc.ID
}

func DistributedNotificationCenterFrom(ptr unsafe.Pointer) DistributedNotificationCenter {
	return DistributedNotificationCenter{
		ID: objc.ID(ptr),
	}
}





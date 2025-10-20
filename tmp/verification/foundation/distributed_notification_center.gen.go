// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var distributedNotificationCenterClass _DistributedNotificationCenterClass

func init() {
	distributedNotificationCenterClass = _DistributedNotificationCenterClass{objc.GetClass("NSDistributedNotificationCenter")}
}

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





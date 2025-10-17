// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DistributedNotificationCenter] class.
var DistributedNotificationCenterClass objc.Class

func init() {
	DistributedNotificationCenterClass = objc.GetClass("NSDistributedNotificationCenter")
}

type DistributedNotificationCenter struct {
	objc.ID
}

func DistributedNotificationCenterFrom(ptr unsafe.Pointer) DistributedNotificationCenter {
	return DistributedNotificationCenter{
		ID: objc.ID(ptr),
	}
}




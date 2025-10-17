// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Notification] class.
var NotificationClass objc.Class

func init() {
	NotificationClass = objc.GetClass("NSNotification")
}

type Notification struct {
	objc.ID
}

func NotificationFrom(ptr unsafe.Pointer) Notification {
	return Notification{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc Notification) Alloc() Notification {
	ret := objc.ID(NotificationClass).Send(objc.RegisterName("alloc"))
	return Notification{ret}
}

// Init initializes the instance.
func (n_ Notification) Init() Notification {
	ret := n_.ID.Send(objc.RegisterName("init"))
	return Notification{ret}
}
// Initializes a notification with the data from an unarchiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSNotification/init(coder:)
func NewNotificationWithCoder(coder unsafe.Pointer) Notification {
	instance := Notification{}.Alloc()
	sel := objc.RegisterName("initWithCoder:")
	ret := instance.ID.Send(sel, coder)
	instance = Notification{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}




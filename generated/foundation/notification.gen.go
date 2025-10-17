// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Notification] class.
var notificationClass = _NotificationClass{objc.GetClass("NSNotification")}

type _NotificationClass struct {
	class objc.Class
}

// An interface definition for the [Notification] class.
type INotification interface {
	objectivec.IObject
}

// A container for information broadcast through a notification center to all registered observers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification

type Notification struct {
	objectivec.Object
}

// NotificationFrom constructs a [Notification] from an unsafe.Pointer.
//
// A container for information broadcast through a notification center to all registered observers.
func NotificationFrom(ptr unsafe.Pointer) Notification {
	return Notification{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _NotificationClass) Alloc() Notification {
	rv := objc.Send[Notification](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NotificationClass) New() Notification {
	rv := objc.Send[Notification](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ Notification) Init() Notification {
	rv := objc.Send[Notification](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ Notification) Autorelease() Notification {
	rv := objc.Send[Notification](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNotification creates a new Notification instance.
func NewNotification() Notification {
	return notificationClass.New()
}
// Initializes a notification with the data from an unarchiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/init(coder:)
func NewNotificationWithCoder(coder unsafe.Pointer) Notification {
	instance := notificationClass.Alloc()
	rv := objc.Send[Notification](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}




// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Notification] class.
var (
	NotificationClass     _NotificationClass
	NotificationClassOnce sync.Once
)

func getNotificationClass() _NotificationClass {
	NotificationClassOnce.Do(func() {
		NotificationClass = _NotificationClass{objc.GetClass("NSNotification")}
	})
	return NotificationClass
}

type _NotificationClass struct {
	class objc.Class
}

// An interface definition for the [Notification] class.
type INotification interface {
	objectivec.IObject
}

// A container for information broadcast through a notification center to all registered observers.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. A notification contains a name, an object, and an optional dictionary, and is broadcast to by instances of or . The name is a tag identifying the notification. The object is any object that the poster of the notification wants to send to observers of that notification (typically, the object posting the notification). The dictionary stores other related objects, if any. objects are immutable. You don’t usually create your own notifications directly, but instead call the methods and .
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getNotificationClass().New()
}




// Initializes a notification with the data from an unarchiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/init(coder:)
func NewNotificationWithCoder(coder unsafe.Pointer) Notification {
	instance := getNotificationClass().Alloc()
	rv := objc.Send[Notification](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Returns a new notification object with a specified name and object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/init(name:object:)
func NewNotificationWithNameObject(aName unsafe.Pointer, anObject objc.ID) Notification {
	rv := objc.Send[Notification](objc.ID(getNotificationClass().class), objc.Sel("notificationWithName:object:"), aName, anObject)
	return rv
}



// Initializes a notification with a specified name, object, and user information.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/init(name:object:userInfo:)
func NewNotificationWithNameObjectUserInfo(name unsafe.Pointer, object objc.ID, userInfo objc.ID) Notification {
	instance := getNotificationClass().Alloc()
	rv := objc.Send[Notification](instance.ID, objc.Sel("initWithName:object:userInfo:"), name, object, userInfo)
	rv.Autorelease()
	return rv
}


// Returns a new notification object with a specified name and object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/init(name:object:)
func (nc _NotificationClass) NotificationWithNameObject(aName unsafe.Pointer, anObject objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("notificationWithName:object:"), aName, anObject)
	return rv
}

// Returns a notification object with a specified name, object, and user information.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/notificationWithName:object:userInfo:
func (nc _NotificationClass) NotificationWithNameObjectUserInfo(aName unsafe.Pointer, anObject objc.ID, aUserInfo objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("notificationWithName:object:userInfo:"), aName, anObject, aUserInfo)
	return rv
}

// The name of the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/name-swift.property
func (n_ Notification) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("name"))
	return rv
}

// The object associated with the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/object
func (n_ Notification) GetObject() objc.ID {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("object"))
	return rv
}

// The user information dictionary associated with the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/userInfo
func (n_ Notification) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("userInfo"))
	return rv
}



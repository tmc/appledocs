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
	// properties:
	Name() unsafe.Pointer
	SetName(value unsafe.Pointer)
	GetObject() unsafe.Pointer
	SetGetObject(value unsafe.Pointer)
	UserInfo() unsafe.Pointer
	SetUserInfo(value unsafe.Pointer)
	// methods:
}

// A container for information broadcast through a notification center to all registered observers.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. A notification contains a name, an object, and an optional dictionary, and is broadcast to by instances of or . The name is a tag identifying the notification. The object is any object that the poster of the notification wants to send to observers of that notification (typically, the object posting the notification). The dictionary stores other related objects, if any. objects are immutable. You don’t usually create your own notifications directly, but instead call the methods and .


// A container for information broadcast through a notification center to all registered observers.
//
// [Full Topic]
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



// The name of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.property
func (n_ Notification) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("name"))
	return rv
}


// The name of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.property
func (n_ Notification) SetName(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setName:"), value)
}


// The object associated with the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotification/object
func (n_ Notification) GetObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("object"))
	return rv
}


// The object associated with the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotification/object
func (n_ Notification) SetGetObject(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGetObject:"), value)
}


// The user information dictionary associated with the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotification/userinfo
func (n_ Notification) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("userInfo"))
	return rv
}


// The user information dictionary associated with the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotification/userinfo
func (n_ Notification) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUserInfo:"), value)
}




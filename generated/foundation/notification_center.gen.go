// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NotificationCenter] class.
var (
	notificationCenterClass     _NotificationCenterClass
	notificationCenterClassOnce sync.Once
)

func getNotificationCenterClass() _NotificationCenterClass {
	notificationCenterClassOnce.Do(func() {
		notificationCenterClass = _NotificationCenterClass{objc.GetClass("NSNotificationCenter")}
	})
	return notificationCenterClass
}

type _NotificationCenterClass struct {
	class objc.Class
}

// An interface definition for the [NotificationCenter] class.
type INotificationCenter interface {
	objectivec.IObject
	AddObserverSelectorNameObject(observer objc.ID, aSelector objc.SEL, aName unsafe.Pointer, anObject objc.ID)
	AddObserverForNameObjectQueueUsingBlock(name unsafe.Pointer, obj objc.ID, queue unsafe.Pointer, block unsafe.Pointer) unsafe.Pointer
	RemoveObserver(observer objc.ID)
	RemoveObserverNameObject(observer objc.ID, aName unsafe.Pointer, anObject objc.ID)
}

// A notification dispatch mechanism that enables the broadcast of information to registered observers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter
type NotificationCenter struct {
	objectivec.Object
}

// NotificationCenterFrom constructs a [NotificationCenter] from an unsafe.Pointer.
//
// A notification dispatch mechanism that enables the broadcast of information to registered observers.
func NotificationCenterFrom(ptr unsafe.Pointer) NotificationCenter {
	return NotificationCenter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NotificationCenterClass) Alloc() NotificationCenter {
	rv := objc.Send[NotificationCenter](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NotificationCenterClass) New() NotificationCenter {
	rv := objc.Send[NotificationCenter](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NotificationCenter) Init() NotificationCenter {
	rv := objc.Send[NotificationCenter](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NotificationCenter) Autorelease() NotificationCenter {
	rv := objc.Send[NotificationCenter](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNotificationCenter creates a new NotificationCenter instance.
func NewNotificationCenter() NotificationCenter {
	return getNotificationCenterClass().New()
}


// Adds an entry to the notification center to call the provided selector with the notification. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter/addObserver(_:selector:name:object:)
func (n_ NotificationCenter) AddObserverSelectorNameObject(observer objc.ID, aSelector objc.SEL, aName unsafe.Pointer, anObject objc.ID) {
	objc.Send[objc.ID](n_.ID, objc.Sel("addObserver:selector:name:object:"), observer, aSelector, aName, anObject)
}
// Adds an entry to the notification center to receive notifications that passed to the provided block. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter/addObserver(forName:object:queue:using:)
func (n_ NotificationCenter) AddObserverForNameObjectQueueUsingBlock(name unsafe.Pointer, obj objc.ID, queue unsafe.Pointer, block unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("addObserverForName:object:queue:usingBlock:"), name, obj, queue, block)
	return rv
}
// Posts a given notification to the notification center. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter/post(_:)-3x2st
func (n_ NotificationCenter) PostNotification(notification unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("postNotification:"), notification)
}
// Creates a notification with a given name and sender and posts it to the notification center. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter/post(name:object:)
func (n_ NotificationCenter) PostNotificationNameObject(aName unsafe.Pointer, anObject objc.ID) {
	objc.Send[objc.ID](n_.ID, objc.Sel("postNotificationName:object:"), aName, anObject)
}
// Creates a notification with a given name, sender, and information and posts it to the notification center. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter/post(name:object:userInfo:)
func (n_ NotificationCenter) PostNotificationNameObjectUserInfo(aName unsafe.Pointer, anObject objc.ID, aUserInfo unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("postNotificationName:object:userInfo:"), aName, anObject, aUserInfo)
}
// Removes all entries specifying an observer from the notification center’s dispatch table. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter/removeObserver(_:)-2yciv
func (n_ NotificationCenter) RemoveObserver(observer objc.ID) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeObserver:"), observer)
}
// Removes matching entries from the notification center’s dispatch table. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter/removeObserver(_:name:object:)
func (n_ NotificationCenter) RemoveObserverNameObject(observer objc.ID, aName unsafe.Pointer, anObject objc.ID) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeObserver:name:object:"), observer, aName, anObject)
}



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
	NotificationCenterClass     _NotificationCenterClass
	NotificationCenterClassOnce sync.Once
)

func getNotificationCenterClass() _NotificationCenterClass {
	NotificationCenterClassOnce.Do(func() {
		NotificationCenterClass = _NotificationCenterClass{objc.GetClass("NSNotificationCenter")}
	})
	return NotificationCenterClass
}

type _NotificationCenterClass struct {
	class objc.Class
}

// An interface definition for the [NotificationCenter] class.
type INotificationCenter interface {
	objectivec.IObject
	AddObserverSelectorNameObject(observer objectivec.IObject, aSelector objc.SEL, aName NotificationName, anObject objectivec.IObject)
	AddObserverForNameObjectQueueUsingBlock(name NotificationName, obj objectivec.IObject, queue IOperationQueue, block unsafe.Pointer) objc.ID
	PostNotification(notification INotification)
	PostNotificationNameObject(aName NotificationName, anObject objectivec.IObject)
	PostNotificationNameObjectUserInfo(aName NotificationName, anObject objectivec.IObject, aUserInfo objectivec.IObject)
	RemoveObserver(observer objectivec.IObject)
	RemoveObserverNameObject(observer objectivec.IObject, aName NotificationName, anObject objectivec.IObject)
}

// A notification dispatch mechanism that enables the broadcast of information to registered observers.
//
// Callers register with a notification center to receive one or both of the following: objects, when working in Objective-C or with frameworks that only support . Objects register with a notification center to receive notifications ( objects) using the or methods, specifying a notification name and optionally a source object. When a caller adds itself as an observer, it specifies which notifications it should receive. and instances for use with Swift code, providing strong typing, appropriate actor isolation, and a more idiomatic Swift experience. Callers register with the notification center using the various flavors of the method, specifying either a message type or a convenience to identify the notification messages to receive. See for more information about this API. Callers may add observers for many different notifications, or even the same notification name or message type as produced by different source objects. Each running app has a notification center, and you can create new notification centers to organize communications in particular contexts. A notification center can deliver notifications only within a single program. On macOS, if you want to post a notification to other processes or receive notifications from other processes, use instead.


// A notification dispatch mechanism that enables the broadcast of information to registered observers.
//
// [Full Topic]
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



// The app’s default notification center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter/default
func (nc _NotificationCenterClass) DefaultCenter() NotificationCenter {
	rv := objc.Send[NotificationCenter](objc.ID(nc.class), objc.Sel("defaultCenter"))
	return rv
}

// Adds an entry to the notification center to call the provided selector with the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter/addObserver(_:selector:name:object:)
func (n_ NotificationCenter) AddObserverSelectorNameObject(observer objectivec.IObject, aSelector objc.SEL, aName NotificationName, anObject objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("addObserver:selector:name:object:"), observer, aSelector, aName, anObject)
}


// Adds an entry to the notification center to receive notifications that passed to the provided block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter/addObserver(forName:object:queue:using:)
func (n_ NotificationCenter) AddObserverForNameObjectQueueUsingBlock(name NotificationName, obj objectivec.IObject, queue IOperationQueue, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("addObserverForName:object:queue:usingBlock:"), name, obj, queue, block)
	return rv
}


// Posts a given notification to the notification center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter/post(_:)-3x2st
func (n_ NotificationCenter) PostNotification(notification INotification) {
	objc.Send[objc.ID](n_.ID, objc.Sel("postNotification:"), notification)
}


// Creates a notification with a given name and sender and posts it to the notification center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter/post(name:object:)
func (n_ NotificationCenter) PostNotificationNameObject(aName NotificationName, anObject objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("postNotificationName:object:"), aName, anObject)
}


// Creates a notification with a given name, sender, and information and posts it to the notification center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter/post(name:object:userInfo:)
func (n_ NotificationCenter) PostNotificationNameObjectUserInfo(aName NotificationName, anObject objectivec.IObject, aUserInfo objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("postNotificationName:object:userInfo:"), aName, anObject, aUserInfo)
}


// Removes all entries specifying an observer from the notification center’s dispatch table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter/removeObserver(_:)-2yciv
func (n_ NotificationCenter) RemoveObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeObserver:"), observer)
}


// Removes matching entries from the notification center’s dispatch table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter/removeObserver(_:name:object:)
func (n_ NotificationCenter) RemoveObserverNameObject(observer objectivec.IObject, aName NotificationName, anObject objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeObserver:name:object:"), observer, aName, anObject)
}


// The app’s default notification center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NotificationCenter/default
func (n_ NotificationCenter) DefaultCenter() INotificationCenter {
	rv := objc.Send[NotificationCenter](n_.ID, objc.Sel("defaultCenter"))
	return rv
}




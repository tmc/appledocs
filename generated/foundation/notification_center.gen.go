// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [NotificationCenter] class.
var NotificationCenterClass objc.Class

func init() {
	NotificationCenterClass = objc.GetClass("NSNotificationCenter")
}

type NotificationCenter struct {
	objc.ID
}

func NotificationCenterFrom(ptr unsafe.Pointer) NotificationCenter {
	return NotificationCenter{
		ID: objc.ID(ptr),
	}
}


// Adds an entry to the notification center to call the provided selector with the notification. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NotificationCenter/addObserver(_:selector:name:object:)
func (n_ NotificationCenter) AddObserverSelectorNameObject(observer objc.ID, aSelector objc.SEL, aName unsafe.Pointer, anObject objc.ID) {
	sel := objc.RegisterName("addObserver:selector:name:object:")
	n_.ID.Send(sel, observer, aSelector, aName, anObject)
}
// Adds an entry to the notification center to receive notifications that passed to the provided block. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NotificationCenter/addObserver(forName:object:queue:using:)
func (n_ NotificationCenter) AddObserverForNameObjectQueueUsingBlock(name unsafe.Pointer, obj objc.ID, queue unsafe.Pointer, block unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("addObserverForName:object:queue:usingBlock:")
	ret := n_.ID.Send(sel, name, obj, queue, block)
	return unsafe.Pointer(ret)
}
// Posts a given notification to the notification center. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NotificationCenter/post(_:)-3x2st
func (n_ NotificationCenter) PostNotification(notification unsafe.Pointer) {
	sel := objc.RegisterName("postNotification:")
	n_.ID.Send(sel, notification)
}
// Creates a notification with a given name and sender and posts it to the notification center. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NotificationCenter/post(name:object:)
func (n_ NotificationCenter) PostNotificationNameObject(aName unsafe.Pointer, anObject objc.ID) {
	sel := objc.RegisterName("postNotificationName:object:")
	n_.ID.Send(sel, aName, anObject)
}
// Creates a notification with a given name, sender, and information and posts it to the notification center. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NotificationCenter/post(name:object:userInfo:)
func (n_ NotificationCenter) PostNotificationNameObjectUserInfo(aName unsafe.Pointer, anObject objc.ID, aUserInfo unsafe.Pointer) {
	sel := objc.RegisterName("postNotificationName:object:userInfo:")
	n_.ID.Send(sel, aName, anObject, aUserInfo)
}
// Removes all entries specifying an observer from the notification center’s dispatch table. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NotificationCenter/removeObserver(_:)-2yciv
func (n_ NotificationCenter) RemoveObserver(observer objc.ID) {
	sel := objc.RegisterName("removeObserver:")
	n_.ID.Send(sel, observer)
}
// Removes matching entries from the notification center’s dispatch table. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NotificationCenter/removeObserver(_:name:object:)
func (n_ NotificationCenter) RemoveObserverNameObject(observer objc.ID, aName unsafe.Pointer, anObject objc.ID) {
	sel := objc.RegisterName("removeObserver:name:object:")
	n_.ID.Send(sel, observer, aName, anObject)
}


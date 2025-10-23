// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DistributedNotificationCenter] class.
var (
	DistributedNotificationCenterClass     _DistributedNotificationCenterClass
	DistributedNotificationCenterClassOnce sync.Once
)

func getDistributedNotificationCenterClass() _DistributedNotificationCenterClass {
	DistributedNotificationCenterClassOnce.Do(func() {
		DistributedNotificationCenterClass = _DistributedNotificationCenterClass{objc.GetClass("NSDistributedNotificationCenter")}
	})
	return DistributedNotificationCenterClass
}

type _DistributedNotificationCenterClass struct {
	class objc.Class
}

// An interface definition for the [DistributedNotificationCenter] class.
type IDistributedNotificationCenter interface {
	INotificationCenter
	AddObserverSelectorNameObject(observer objectivec.IObject, aSelector objc.SEL, aName NotificationName, anObject string)
	AddObserverSelectorNameObjectSuspensionBehavior(observer objectivec.IObject, selector objc.SEL, name NotificationName, object string, suspensionBehavior NSNotificationSuspensionBehavior)
	PostNotificationNameObject(aName NotificationName, anObject string)
	PostNotificationNameObjectUserInfo(aName NotificationName, anObject string, aUserInfo objectivec.IObject)
	PostNotificationNameObjectUserInfoDeliverImmediately(name NotificationName, object string, userInfo objectivec.IObject, deliverImmediately bool)
	PostNotificationNameObjectUserInfoOptions(name NotificationName, object string, userInfo objectivec.IObject, options NSDistributedNotificationOptions)
	RemoveObserverNameObject(observer objectivec.IObject, aName NotificationName, anObject string)
	Suspended() bool
	SetSuspended(value bool)
}

// A notification dispatch mechanism that enables the broadcast of notifications across task boundaries.
//
// A instance broadcasts objects to objects in other tasks that have registered for the notification with their task’s default distributed notification center.


// A notification dispatch mechanism that enables the broadcast of notifications across task boundaries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter
type DistributedNotificationCenter struct {
	NotificationCenter
}

// DistributedNotificationCenterFrom constructs a [DistributedNotificationCenter] from an unsafe.Pointer.
//
// A notification dispatch mechanism that enables the broadcast of notifications across task boundaries.
func DistributedNotificationCenterFrom(ptr unsafe.Pointer) DistributedNotificationCenter {
	return DistributedNotificationCenter{
		NotificationCenter: NotificationCenterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DistributedNotificationCenterClass) Alloc() DistributedNotificationCenter {
	rv := objc.Send[DistributedNotificationCenter](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DistributedNotificationCenterClass) New() DistributedNotificationCenter {
	rv := objc.Send[DistributedNotificationCenter](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DistributedNotificationCenter) Init() DistributedNotificationCenter {
	rv := objc.Send[DistributedNotificationCenter](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DistributedNotificationCenter) Autorelease() DistributedNotificationCenter {
	rv := objc.Send[DistributedNotificationCenter](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDistributedNotificationCenter creates a new DistributedNotificationCenter instance.
func NewDistributedNotificationCenter() DistributedNotificationCenter {
	return getDistributedNotificationCenterClass().New()
}



// Returns the default distributed notification center, representing the local notification center for the computer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/default()
func (dc _DistributedNotificationCenterClass) DefaultCenter() DistributedNotificationCenter {
	rv := objc.Send[DistributedNotificationCenter](objc.ID(dc.class), objc.Sel("defaultCenter"))
	return rv
}


// Returns the distributed notification center for a particular notification center type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/forType(_:)
func (dc _DistributedNotificationCenterClass) NotificationCenterForType(notificationCenterType DistributedNotificationCenterType) DistributedNotificationCenter {
	rv := objc.Send[DistributedNotificationCenter](objc.ID(dc.class), objc.Sel("notificationCenterForType:"), notificationCenterType)
	return rv
}


// Adds an entry to the notification center’s dispatch table with an observer, a selector, and an optional notification name and sender.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/addObserver(_:selector:name:object:)
func (d_ DistributedNotificationCenter) AddObserverSelectorNameObject(observer objectivec.IObject, aSelector objc.SEL, aName NotificationName, anObject string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("addObserver:selector:name:object:"), observer, aSelector, aName, objc.String(anObject))
}


// Adds an entry to the receiver’s dispatch table with a specific observer and suspended-notifications behavior, and optional notification name and sender.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/addObserver(_:selector:name:object:suspensionBehavior:)
func (d_ DistributedNotificationCenter) AddObserverSelectorNameObjectSuspensionBehavior(observer objectivec.IObject, selector objc.SEL, name NotificationName, object string, suspensionBehavior NSNotificationSuspensionBehavior) {
	objc.Send[objc.ID](d_.ID, objc.Sel("addObserver:selector:name:object:suspensionBehavior:"), observer, selector, name, objc.String(object), suspensionBehavior)
}


// Creates a notification, and posts it to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/post(name:object:)
func (d_ DistributedNotificationCenter) PostNotificationNameObject(aName NotificationName, anObject string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("postNotificationName:object:"), aName, objc.String(anObject))
}


// Creates a notification with information, and posts it to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/post(name:object:userInfo:)
func (d_ DistributedNotificationCenter) PostNotificationNameObjectUserInfo(aName NotificationName, anObject string, aUserInfo objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("postNotificationName:object:userInfo:"), aName, objc.String(anObject), aUserInfo)
}


// Creates a notification with information and an immediate-delivery specifier, and posts it to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/postNotificationName(_:object:userInfo:deliverImmediately:)
func (d_ DistributedNotificationCenter) PostNotificationNameObjectUserInfoDeliverImmediately(name NotificationName, object string, userInfo objectivec.IObject, deliverImmediately bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("postNotificationName:object:userInfo:deliverImmediately:"), name, objc.String(object), userInfo, deliverImmediately)
}


// Creates a notification with information, and posts it to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/postNotificationName(_:object:userInfo:options:)
func (d_ DistributedNotificationCenter) PostNotificationNameObjectUserInfoOptions(name NotificationName, object string, userInfo objectivec.IObject, options NSDistributedNotificationOptions) {
	objc.Send[objc.ID](d_.ID, objc.Sel("postNotificationName:object:userInfo:options:"), name, objc.String(object), userInfo, options)
}


// Removes matching entries from the receiver’s dispatch table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/removeObserver(_:name:object:)
func (d_ DistributedNotificationCenter) RemoveObserverNameObject(observer objectivec.IObject, aName NotificationName, anObject string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("removeObserver:name:object:"), observer, aName, objc.String(anObject))
}


// Suspends or resumes notification delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/suspended
func (d_ DistributedNotificationCenter) Suspended() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("suspended"))
	return rv
}


// Suspends or resumes notification delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/suspended
func (d_ DistributedNotificationCenter) SetSuspended(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSuspended:"), value)
}




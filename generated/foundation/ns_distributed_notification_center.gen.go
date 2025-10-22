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
	PostNotificationNameObjectUserInfo(aName INotificationName, anObject string, aUserInfo objectivec.IObject)
	PostNotificationNameObjectUserInfoDeliverImmediately(name INotificationName, object string, userInfo objectivec.IObject, deliverImmediately bool)
}

// A notification dispatch mechanism that enables the broadcast of notifications across task boundaries.
//
// A instance broadcasts objects to objects in other tasks that have registered for the notification with their task’s default distributed notification center.
//
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


// Creates a notification with information, and posts it to the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/post(name:object:userInfo:)
func (d_ DistributedNotificationCenter) PostNotificationNameObjectUserInfo(aName INotificationName, anObject string, aUserInfo objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("postNotificationName:object:userInfo:"), aName, objc.String(anObject), aUserInfo)
}

// Creates a notification with information and an immediate-delivery specifier, and posts it to the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/postNotificationName(_:object:userInfo:deliverImmediately:)
func (d_ DistributedNotificationCenter) PostNotificationNameObjectUserInfoDeliverImmediately(name INotificationName, object string, userInfo objectivec.IObject, deliverImmediately bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("postNotificationName:object:userInfo:deliverImmediately:"), name, objc.String(object), userInfo, deliverImmediately)
}

// Suspends or resumes notification delivery.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/distributednotificationcenter/suspended
func (d_ DistributedNotificationCenter) Suspended() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("suspended"))
	return rv
}


// SetSuspended sets the value of the suspended property.
// Suspends or resumes notification delivery.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/distributednotificationcenter/suspended
func (d_ DistributedNotificationCenter) SetSuspended(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSuspended:"), value)
}




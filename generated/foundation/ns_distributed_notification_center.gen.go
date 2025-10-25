// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSDistributedNotificationCenter */


/* debug [class_header]: Header for NSDistributedNotificationCenter */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DistributedNotificationCenter */
// An interface definition for the [DistributedNotificationCenter] class.
type IDistributedNotificationCenter interface {
	INotificationCenter
	
/* debug [class_interface_properties]: Properties for DistributedNotificationCenter */
	// properties:
	Suspended() bool
	SetSuspended(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DistributedNotificationCenter */
	// methods:
	AddObserverSelectorNameObjectSuspensionBehavior(observer objc.IObject, selector objc.SEL, name NotificationName, object IString, suspensionBehavior NotificationSuspensionBehavior)
	PostNotificationNameObjectUserInfoDeliverImmediately(name NotificationName, object IString, userInfo IDictionary, deliverImmediately bool)
	PostNotificationNameObjectUserInfoOptions(name NotificationName, object IString, userInfo IDictionary, options DistributedNotificationOptions)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DistributedNotificationCenter */
// Alloc allocates a new instance without initialization.
func (dc _DistributedNotificationCenterClass) Alloc() DistributedNotificationCenter {
	rv := objc.Send[DistributedNotificationCenter](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DistributedNotificationCenter */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DistributedNotificationCenter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DistributedNotificationCenter */

// Returns the default distributed notification center, representing the local notification center for the computer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/default()
func (dc _DistributedNotificationCenterClass) DefaultCenter() IDistributedNotificationCenter {
	rv := objc.Send[DistributedNotificationCenter](objc.ID(dc.class), objc.Sel("defaultCenter"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultCenter) */


// Returns the distributed notification center for a particular notification center type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/forType(_:)
func (dc _DistributedNotificationCenterClass) NotificationCenterForType(notificationCenterType DistributedNotificationCenterType) IDistributedNotificationCenter {
	rv := objc.Send[DistributedNotificationCenter](objc.ID(dc.class), objc.Sel("notificationCenterForType:"), notificationCenterType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NotificationCenterForType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DistributedNotificationCenter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DistributedNotificationCenter */

// Adds an entry to the receiver’s dispatch table with a specific observer and suspended-notifications behavior, and optional notification name and sender.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/addObserver(_:selector:name:object:suspensionBehavior:)
func (d_ DistributedNotificationCenter) AddObserverSelectorNameObjectSuspensionBehavior(observer objc.IObject, selector objc.SEL, name NotificationName, object IString, suspensionBehavior NotificationSuspensionBehavior) {
	objc.Send[objc.ID](d_.ID, objc.Sel("addObserver:selector:name:object:suspensionBehavior:"), observer, selector, name, object, suspensionBehavior)
}/* debug [instance_methods/method]: AddObserverSelectorNameObjectSuspensionBehavior */


// Creates a notification with information and an immediate-delivery specifier, and posts it to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/postNotificationName(_:object:userInfo:deliverImmediately:)
func (d_ DistributedNotificationCenter) PostNotificationNameObjectUserInfoDeliverImmediately(name NotificationName, object IString, userInfo IDictionary, deliverImmediately bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("postNotificationName:object:userInfo:deliverImmediately:"), name, object, userInfo, deliverImmediately)
}/* debug [instance_methods/method]: PostNotificationNameObjectUserInfoDeliverImmediately */


// Creates a notification with information, and posts it to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/postNotificationName(_:object:userInfo:options:)
func (d_ DistributedNotificationCenter) PostNotificationNameObjectUserInfoOptions(name NotificationName, object IString, userInfo IDictionary, options DistributedNotificationOptions) {
	objc.Send[objc.ID](d_.ID, objc.Sel("postNotificationName:object:userInfo:options:"), name, object, userInfo, options)
}/* debug [instance_methods/method]: PostNotificationNameObjectUserInfoOptions */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DistributedNotificationCenter */

// Suspends or resumes notification delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/suspended
func (d_ DistributedNotificationCenter) Suspended() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("suspended"))
	return rv
}/* debug [instance_properties/getter]: suspended */


// Suspends or resumes notification delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/suspended
func (d_ DistributedNotificationCenter) SetSuspended(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSuspended:"), value)
}/* debug [instance_properties/setter]: suspended */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDistributedNotificationCenter */




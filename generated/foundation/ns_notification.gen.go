// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSNotification */


/* debug [class_header]: Header for NSNotification */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Notification */
// An interface definition for the [Notification] class.
type INotification interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Notification */
	// properties:
	Name() NotificationName /* typedef */
	GetObject() objc.ID
	UserInfo() IDictionary
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Notification */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Notification */
// Alloc allocates a new instance without initialization.
func (nc _NotificationClass) Alloc() Notification {
	rv := objc.Send[Notification](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Notification */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Notification */

// Initializes a notification with the data from an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/init(coder:)
func NewNotificationWithCoder(coder ICoder) Notification {
	instance := getNotificationClass().Alloc()
	rv := objc.Send[Notification](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNotificationWithCoder */


// Returns a new notification object with a specified name and object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/init(name:object:)
func NewNotificationWithNameObject(aName NotificationName /* typedef */, anObject objc.IObject) Notification {
	rv := objc.Send[Notification](objc.ID(getNotificationClass().class), objc.Sel("notificationWithName:object:"), aName, anObject)
	return rv
}/* debug [class_init_methods/constructor]: NewNotificationWithNameObject */


// Initializes a notification with a specified name, object, and user information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/init(name:object:userInfo:)
func NewNotificationWithNameObjectUserInfo(name NotificationName /* typedef */, object objc.IObject, userInfo IDictionary) Notification {
	instance := getNotificationClass().Alloc()
	rv := objc.Send[Notification](instance.ID, objc.Sel("initWithName:object:userInfo:"), name, object, userInfo)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNotificationWithNameObjectUserInfo */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Notification */

// Returns a new notification object with a specified name and object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/init(name:object:)
func (nc _NotificationClass) NotificationWithNameObject(aName NotificationName /* typedef */, anObject objc.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("notificationWithName:object:"), aName, anObject)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NotificationWithNameObject) */


// Returns a notification object with a specified name, object, and user information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/notificationWithName:object:userInfo:
func (nc _NotificationClass) NotificationWithNameObjectUserInfo(aName NotificationName /* typedef */, anObject objc.IObject, aUserInfo IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("notificationWithName:object:userInfo:"), aName, anObject, aUserInfo)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NotificationWithNameObjectUserInfo) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Notification */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Notification */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Notification */

// The name of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/name-swift.property
func (n_ Notification) Name() NotificationName /* typedef */ {
	rv := objc.Send[String](n_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The object associated with the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/object
func (n_ Notification) GetObject() objc.ID {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("object"))
	return rv
}/* debug [instance_properties/getter]: object */


// The user information dictionary associated with the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/userInfo
func (n_ Notification) UserInfo() IDictionary {
	rv := objc.Send[Dictionary](n_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSNotification */



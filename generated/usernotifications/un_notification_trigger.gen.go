// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class UNNotificationTrigger */


/* debug [class_header]: Header for UNNotificationTrigger */
// The class instance for the [UNNotificationTrigger] class.
var (
	UNNotificationTriggerClass     _UNNotificationTriggerClass
	UNNotificationTriggerClassOnce sync.Once
)

func getUNNotificationTriggerClass() _UNNotificationTriggerClass {
	UNNotificationTriggerClassOnce.Do(func() {
		UNNotificationTriggerClass = _UNNotificationTriggerClass{objc.GetClass("UNNotificationTrigger")}
	})
	return UNNotificationTriggerClass
}

type _UNNotificationTriggerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UNNotificationTrigger */
// An interface definition for the [UNNotificationTrigger] class.
type IUNNotificationTrigger interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for UNNotificationTrigger */
	// properties:
	Repeats() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UNNotificationTrigger */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UNNotificationTrigger */
// Alloc allocates a new instance without initialization.
func (uc _UNNotificationTriggerClass) Alloc() UNNotificationTrigger {
	rv := objc.Send[UNNotificationTrigger](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UNNotificationTriggerClass) New() UNNotificationTrigger {
	rv := objc.Send[UNNotificationTrigger](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNNotificationTrigger) Init() UNNotificationTrigger {
	rv := objc.Send[UNNotificationTrigger](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNNotificationTrigger) Autorelease() UNNotificationTrigger {
	rv := objc.Send[UNNotificationTrigger](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationTrigger creates a new UNNotificationTrigger instance.
func NewUNNotificationTrigger() UNNotificationTrigger {
	return getUNNotificationTriggerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UNNotificationTrigger */
// The common behavior for subclasses that trigger the delivery of a local or remote notification.
//
// The class is an abstract class for representing an event that triggers the delivery of a notification. You don’t create instances of this class directly. Instead, you instantiate the concrete subclass that defines the trigger condition you want for your notification. You then assign the resulting object to the object that you use to schedule your notification. Concrete trigger classes include the following:


// The common behavior for subclasses that trigger the delivery of a local or remote notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationTrigger
type UNNotificationTrigger struct {
	objectivec.Object
}

// UNNotificationTriggerFrom constructs a [UNNotificationTrigger] from an unsafe.Pointer.
//
// The common behavior for subclasses that trigger the delivery of a local or remote notification.
func UNNotificationTriggerFrom(ptr unsafe.Pointer) UNNotificationTrigger {
	return UNNotificationTrigger{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UNNotificationTrigger *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UNNotificationTrigger */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UNNotificationTrigger */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UNNotificationTrigger */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UNNotificationTrigger */

// A Boolean value indicating whether the system reschedules the notification after it’s delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationTrigger/repeats
func (u_ UNNotificationTrigger) Repeats() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("repeats"))
	return rv
}/* debug [instance_properties/getter]: repeats */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class UNNotificationTrigger */




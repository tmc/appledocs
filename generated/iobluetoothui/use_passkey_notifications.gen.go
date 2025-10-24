// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class usePasskeyNotifications */


/* debug [class_header]: Header for usePasskeyNotifications */
// The class instance for the [usePasskeyNotifications] class.
var (
	UsePasskeyNotificationsClass     _usePasskeyNotificationsClass
	UsePasskeyNotificationsClassOnce sync.Once
)

func getusePasskeyNotificationsClass() _usePasskeyNotificationsClass {
	UsePasskeyNotificationsClassOnce.Do(func() {
		UsePasskeyNotificationsClass = _usePasskeyNotificationsClass{objc.GetClass("usePasskeyNotifications")}
	})
	return UsePasskeyNotificationsClass
}

type _usePasskeyNotificationsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for usePasskeyNotifications */
// An interface definition for the [usePasskeyNotifications] class.
type IusePasskeyNotifications interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for usePasskeyNotifications */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for usePasskeyNotifications */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for usePasskeyNotifications */
// Alloc allocates a new instance without initialization.
func (uc _usePasskeyNotificationsClass) Alloc() usePasskeyNotifications {
	rv := objc.Send[usePasskeyNotifications](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _usePasskeyNotificationsClass) New() usePasskeyNotifications {
	rv := objc.Send[usePasskeyNotifications](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ usePasskeyNotifications) Init() usePasskeyNotifications {
	rv := objc.Send[usePasskeyNotifications](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ usePasskeyNotifications) Autorelease() usePasskeyNotifications {
	rv := objc.Send[usePasskeyNotifications](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewusePasskeyNotifications creates a new usePasskeyNotifications instance.
func NewusePasskeyNotifications() usePasskeyNotifications {
	return getusePasskeyNotificationsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for usePasskeyNotifications */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/usePasskeyNotifications
type usePasskeyNotifications struct {
	objectivec.Object
}

// usePasskeyNotificationsFrom constructs a [usePasskeyNotifications] from an unsafe.Pointer.
func usePasskeyNotificationsFrom(ptr unsafe.Pointer) usePasskeyNotifications {
	return usePasskeyNotifications{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for usePasskeyNotifications *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for usePasskeyNotifications */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for usePasskeyNotifications */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for usePasskeyNotifications */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for usePasskeyNotifications */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class usePasskeyNotifications */




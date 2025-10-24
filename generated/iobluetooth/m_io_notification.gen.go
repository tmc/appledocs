// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mIONotification */


/* debug [class_header]: Header for mIONotification */
// The class instance for the [mIONotification] class.
var (
	MIONotificationClass     _mIONotificationClass
	MIONotificationClassOnce sync.Once
)

func getmIONotificationClass() _mIONotificationClass {
	MIONotificationClassOnce.Do(func() {
		MIONotificationClass = _mIONotificationClass{objc.GetClass("mIONotification")}
	})
	return MIONotificationClass
}

type _mIONotificationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mIONotification */
// An interface definition for the [mIONotification] class.
type ImIONotification interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mIONotification */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mIONotification */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mIONotification */
// Alloc allocates a new instance without initialization.
func (mc _mIONotificationClass) Alloc() mIONotification {
	rv := objc.Send[mIONotification](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mIONotificationClass) New() mIONotification {
	rv := objc.Send[mIONotification](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIONotification) Init() mIONotification {
	rv := objc.Send[mIONotification](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIONotification) Autorelease() mIONotification {
	rv := objc.Send[mIONotification](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIONotification creates a new mIONotification instance.
func NewmIONotification() mIONotification {
	return getmIONotificationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mIONotification */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothObject/mIONotification
type mIONotification struct {
	objectivec.Object
}

// mIONotificationFrom constructs a [mIONotification] from an unsafe.Pointer.
func mIONotificationFrom(ptr unsafe.Pointer) mIONotification {
	return mIONotification{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mIONotification *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mIONotification */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mIONotification */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mIONotification */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mIONotification */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mIONotification */




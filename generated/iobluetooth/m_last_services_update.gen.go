// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mLastServicesUpdate */


/* debug [class_header]: Header for mLastServicesUpdate */
// The class instance for the [mLastServicesUpdate] class.
var (
	MLastServicesUpdateClass     _mLastServicesUpdateClass
	MLastServicesUpdateClassOnce sync.Once
)

func getmLastServicesUpdateClass() _mLastServicesUpdateClass {
	MLastServicesUpdateClassOnce.Do(func() {
		MLastServicesUpdateClass = _mLastServicesUpdateClass{objc.GetClass("mLastServicesUpdate")}
	})
	return MLastServicesUpdateClass
}

type _mLastServicesUpdateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mLastServicesUpdate */
// An interface definition for the [mLastServicesUpdate] class.
type ImLastServicesUpdate interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mLastServicesUpdate */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mLastServicesUpdate */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mLastServicesUpdate */
// Alloc allocates a new instance without initialization.
func (mc _mLastServicesUpdateClass) Alloc() mLastServicesUpdate {
	rv := objc.Send[mLastServicesUpdate](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mLastServicesUpdateClass) New() mLastServicesUpdate {
	rv := objc.Send[mLastServicesUpdate](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mLastServicesUpdate) Init() mLastServicesUpdate {
	rv := objc.Send[mLastServicesUpdate](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mLastServicesUpdate) Autorelease() mLastServicesUpdate {
	rv := objc.Send[mLastServicesUpdate](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmLastServicesUpdate creates a new mLastServicesUpdate instance.
func NewmLastServicesUpdate() mLastServicesUpdate {
	return getmLastServicesUpdateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mLastServicesUpdate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mLastServicesUpdate
type mLastServicesUpdate struct {
	objectivec.Object
}

// mLastServicesUpdateFrom constructs a [mLastServicesUpdate] from an unsafe.Pointer.
func mLastServicesUpdateFrom(ptr unsafe.Pointer) mLastServicesUpdate {
	return mLastServicesUpdate{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mLastServicesUpdate *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mLastServicesUpdate */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mLastServicesUpdate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mLastServicesUpdate */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mLastServicesUpdate */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mLastServicesUpdate */




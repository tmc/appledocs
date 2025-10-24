// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mOpenConnectionTimeout */


/* debug [class_header]: Header for mOpenConnectionTimeout */
// The class instance for the [mOpenConnectionTimeout] class.
var (
	MOpenConnectionTimeoutClass     _mOpenConnectionTimeoutClass
	MOpenConnectionTimeoutClassOnce sync.Once
)

func getmOpenConnectionTimeoutClass() _mOpenConnectionTimeoutClass {
	MOpenConnectionTimeoutClassOnce.Do(func() {
		MOpenConnectionTimeoutClass = _mOpenConnectionTimeoutClass{objc.GetClass("mOpenConnectionTimeout")}
	})
	return MOpenConnectionTimeoutClass
}

type _mOpenConnectionTimeoutClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mOpenConnectionTimeout */
// An interface definition for the [mOpenConnectionTimeout] class.
type ImOpenConnectionTimeout interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mOpenConnectionTimeout */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mOpenConnectionTimeout */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mOpenConnectionTimeout */
// Alloc allocates a new instance without initialization.
func (mc _mOpenConnectionTimeoutClass) Alloc() mOpenConnectionTimeout {
	rv := objc.Send[mOpenConnectionTimeout](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mOpenConnectionTimeoutClass) New() mOpenConnectionTimeout {
	rv := objc.Send[mOpenConnectionTimeout](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOpenConnectionTimeout) Init() mOpenConnectionTimeout {
	rv := objc.Send[mOpenConnectionTimeout](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOpenConnectionTimeout) Autorelease() mOpenConnectionTimeout {
	rv := objc.Send[mOpenConnectionTimeout](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOpenConnectionTimeout creates a new mOpenConnectionTimeout instance.
func NewmOpenConnectionTimeout() mOpenConnectionTimeout {
	return getmOpenConnectionTimeoutClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mOpenConnectionTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/mOpenConnectionTimeout
type mOpenConnectionTimeout struct {
	objectivec.Object
}

// mOpenConnectionTimeoutFrom constructs a [mOpenConnectionTimeout] from an unsafe.Pointer.
func mOpenConnectionTimeoutFrom(ptr unsafe.Pointer) mOpenConnectionTimeout {
	return mOpenConnectionTimeout{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mOpenConnectionTimeout *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mOpenConnectionTimeout */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mOpenConnectionTimeout */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mOpenConnectionTimeout */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mOpenConnectionTimeout */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mOpenConnectionTimeout */




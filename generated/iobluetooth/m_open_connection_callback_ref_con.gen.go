// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mOpenConnectionCallbackRefCon */


/* debug [class_header]: Header for mOpenConnectionCallbackRefCon */
// The class instance for the [mOpenConnectionCallbackRefCon] class.
var (
	MOpenConnectionCallbackRefConClass     _mOpenConnectionCallbackRefConClass
	MOpenConnectionCallbackRefConClassOnce sync.Once
)

func getmOpenConnectionCallbackRefConClass() _mOpenConnectionCallbackRefConClass {
	MOpenConnectionCallbackRefConClassOnce.Do(func() {
		MOpenConnectionCallbackRefConClass = _mOpenConnectionCallbackRefConClass{objc.GetClass("mOpenConnectionCallbackRefCon")}
	})
	return MOpenConnectionCallbackRefConClass
}

type _mOpenConnectionCallbackRefConClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mOpenConnectionCallbackRefCon */
// An interface definition for the [mOpenConnectionCallbackRefCon] class.
type ImOpenConnectionCallbackRefCon interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mOpenConnectionCallbackRefCon */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mOpenConnectionCallbackRefCon */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mOpenConnectionCallbackRefCon */
// Alloc allocates a new instance without initialization.
func (mc _mOpenConnectionCallbackRefConClass) Alloc() mOpenConnectionCallbackRefCon {
	rv := objc.Send[mOpenConnectionCallbackRefCon](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mOpenConnectionCallbackRefConClass) New() mOpenConnectionCallbackRefCon {
	rv := objc.Send[mOpenConnectionCallbackRefCon](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOpenConnectionCallbackRefCon) Init() mOpenConnectionCallbackRefCon {
	rv := objc.Send[mOpenConnectionCallbackRefCon](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOpenConnectionCallbackRefCon) Autorelease() mOpenConnectionCallbackRefCon {
	rv := objc.Send[mOpenConnectionCallbackRefCon](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOpenConnectionCallbackRefCon creates a new mOpenConnectionCallbackRefCon instance.
func NewmOpenConnectionCallbackRefCon() mOpenConnectionCallbackRefCon {
	return getmOpenConnectionCallbackRefConClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mOpenConnectionCallbackRefCon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/mOpenConnectionCallbackRefCon
type mOpenConnectionCallbackRefCon struct {
	objectivec.Object
}

// mOpenConnectionCallbackRefConFrom constructs a [mOpenConnectionCallbackRefCon] from an unsafe.Pointer.
func mOpenConnectionCallbackRefConFrom(ptr unsafe.Pointer) mOpenConnectionCallbackRefCon {
	return mOpenConnectionCallbackRefCon{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mOpenConnectionCallbackRefCon *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mOpenConnectionCallbackRefCon */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mOpenConnectionCallbackRefCon */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mOpenConnectionCallbackRefCon */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mOpenConnectionCallbackRefCon */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mOpenConnectionCallbackRefCon */




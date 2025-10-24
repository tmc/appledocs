// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mOpenConnectionRefCon */


/* debug [class_header]: Header for mOpenConnectionRefCon */
// The class instance for the [mOpenConnectionRefCon] class.
var (
	MOpenConnectionRefConClass     _mOpenConnectionRefConClass
	MOpenConnectionRefConClassOnce sync.Once
)

func getmOpenConnectionRefConClass() _mOpenConnectionRefConClass {
	MOpenConnectionRefConClassOnce.Do(func() {
		MOpenConnectionRefConClass = _mOpenConnectionRefConClass{objc.GetClass("mOpenConnectionRefCon")}
	})
	return MOpenConnectionRefConClass
}

type _mOpenConnectionRefConClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mOpenConnectionRefCon */
// An interface definition for the [mOpenConnectionRefCon] class.
type ImOpenConnectionRefCon interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mOpenConnectionRefCon */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mOpenConnectionRefCon */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mOpenConnectionRefCon */
// Alloc allocates a new instance without initialization.
func (mc _mOpenConnectionRefConClass) Alloc() mOpenConnectionRefCon {
	rv := objc.Send[mOpenConnectionRefCon](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mOpenConnectionRefConClass) New() mOpenConnectionRefCon {
	rv := objc.Send[mOpenConnectionRefCon](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOpenConnectionRefCon) Init() mOpenConnectionRefCon {
	rv := objc.Send[mOpenConnectionRefCon](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOpenConnectionRefCon) Autorelease() mOpenConnectionRefCon {
	rv := objc.Send[mOpenConnectionRefCon](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOpenConnectionRefCon creates a new mOpenConnectionRefCon instance.
func NewmOpenConnectionRefCon() mOpenConnectionRefCon {
	return getmOpenConnectionRefConClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mOpenConnectionRefCon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/mOpenConnectionRefCon
type mOpenConnectionRefCon struct {
	objectivec.Object
}

// mOpenConnectionRefConFrom constructs a [mOpenConnectionRefCon] from an unsafe.Pointer.
func mOpenConnectionRefConFrom(ptr unsafe.Pointer) mOpenConnectionRefCon {
	return mOpenConnectionRefCon{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mOpenConnectionRefCon *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mOpenConnectionRefCon */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mOpenConnectionRefCon */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mOpenConnectionRefCon */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mOpenConnectionRefCon */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mOpenConnectionRefCon */




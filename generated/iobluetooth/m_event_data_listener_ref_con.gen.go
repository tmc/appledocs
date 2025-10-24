// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mEventDataListenerRefCon */


/* debug [class_header]: Header for mEventDataListenerRefCon */
// The class instance for the [mEventDataListenerRefCon] class.
var (
	MEventDataListenerRefConClass     _mEventDataListenerRefConClass
	MEventDataListenerRefConClassOnce sync.Once
)

func getmEventDataListenerRefConClass() _mEventDataListenerRefConClass {
	MEventDataListenerRefConClassOnce.Do(func() {
		MEventDataListenerRefConClass = _mEventDataListenerRefConClass{objc.GetClass("mEventDataListenerRefCon")}
	})
	return MEventDataListenerRefConClass
}

type _mEventDataListenerRefConClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mEventDataListenerRefCon */
// An interface definition for the [mEventDataListenerRefCon] class.
type ImEventDataListenerRefCon interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mEventDataListenerRefCon */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mEventDataListenerRefCon */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mEventDataListenerRefCon */
// Alloc allocates a new instance without initialization.
func (mc _mEventDataListenerRefConClass) Alloc() mEventDataListenerRefCon {
	rv := objc.Send[mEventDataListenerRefCon](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mEventDataListenerRefConClass) New() mEventDataListenerRefCon {
	rv := objc.Send[mEventDataListenerRefCon](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mEventDataListenerRefCon) Init() mEventDataListenerRefCon {
	rv := objc.Send[mEventDataListenerRefCon](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mEventDataListenerRefCon) Autorelease() mEventDataListenerRefCon {
	rv := objc.Send[mEventDataListenerRefCon](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmEventDataListenerRefCon creates a new mEventDataListenerRefCon instance.
func NewmEventDataListenerRefCon() mEventDataListenerRefCon {
	return getmEventDataListenerRefConClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mEventDataListenerRefCon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mEventDataListenerRefCon
type mEventDataListenerRefCon struct {
	objectivec.Object
}

// mEventDataListenerRefConFrom constructs a [mEventDataListenerRefCon] from an unsafe.Pointer.
func mEventDataListenerRefConFrom(ptr unsafe.Pointer) mEventDataListenerRefCon {
	return mEventDataListenerRefCon{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mEventDataListenerRefCon *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mEventDataListenerRefCon */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mEventDataListenerRefCon */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mEventDataListenerRefCon */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mEventDataListenerRefCon */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mEventDataListenerRefCon */




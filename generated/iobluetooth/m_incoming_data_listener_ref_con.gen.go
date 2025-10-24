// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mIncomingDataListenerRefCon */


/* debug [class_header]: Header for mIncomingDataListenerRefCon */
// The class instance for the [mIncomingDataListenerRefCon] class.
var (
	MIncomingDataListenerRefConClass     _mIncomingDataListenerRefConClass
	MIncomingDataListenerRefConClassOnce sync.Once
)

func getmIncomingDataListenerRefConClass() _mIncomingDataListenerRefConClass {
	MIncomingDataListenerRefConClassOnce.Do(func() {
		MIncomingDataListenerRefConClass = _mIncomingDataListenerRefConClass{objc.GetClass("mIncomingDataListenerRefCon")}
	})
	return MIncomingDataListenerRefConClass
}

type _mIncomingDataListenerRefConClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mIncomingDataListenerRefCon */
// An interface definition for the [mIncomingDataListenerRefCon] class.
type ImIncomingDataListenerRefCon interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mIncomingDataListenerRefCon */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mIncomingDataListenerRefCon */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mIncomingDataListenerRefCon */
// Alloc allocates a new instance without initialization.
func (mc _mIncomingDataListenerRefConClass) Alloc() mIncomingDataListenerRefCon {
	rv := objc.Send[mIncomingDataListenerRefCon](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mIncomingDataListenerRefConClass) New() mIncomingDataListenerRefCon {
	rv := objc.Send[mIncomingDataListenerRefCon](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIncomingDataListenerRefCon) Init() mIncomingDataListenerRefCon {
	rv := objc.Send[mIncomingDataListenerRefCon](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIncomingDataListenerRefCon) Autorelease() mIncomingDataListenerRefCon {
	rv := objc.Send[mIncomingDataListenerRefCon](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIncomingDataListenerRefCon creates a new mIncomingDataListenerRefCon instance.
func NewmIncomingDataListenerRefCon() mIncomingDataListenerRefCon {
	return getmIncomingDataListenerRefConClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mIncomingDataListenerRefCon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mIncomingDataListenerRefCon
type mIncomingDataListenerRefCon struct {
	objectivec.Object
}

// mIncomingDataListenerRefConFrom constructs a [mIncomingDataListenerRefCon] from an unsafe.Pointer.
func mIncomingDataListenerRefConFrom(ptr unsafe.Pointer) mIncomingDataListenerRefCon {
	return mIncomingDataListenerRefCon{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mIncomingDataListenerRefCon *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mIncomingDataListenerRefCon */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mIncomingDataListenerRefCon */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mIncomingDataListenerRefCon */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mIncomingDataListenerRefCon */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mIncomingDataListenerRefCon */




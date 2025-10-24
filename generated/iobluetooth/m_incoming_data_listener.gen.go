// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mIncomingDataListener */


/* debug [class_header]: Header for mIncomingDataListener */
// The class instance for the [mIncomingDataListener] class.
var (
	MIncomingDataListenerClass     _mIncomingDataListenerClass
	MIncomingDataListenerClassOnce sync.Once
)

func getmIncomingDataListenerClass() _mIncomingDataListenerClass {
	MIncomingDataListenerClassOnce.Do(func() {
		MIncomingDataListenerClass = _mIncomingDataListenerClass{objc.GetClass("mIncomingDataListener")}
	})
	return MIncomingDataListenerClass
}

type _mIncomingDataListenerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mIncomingDataListener */
// An interface definition for the [mIncomingDataListener] class.
type ImIncomingDataListener interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mIncomingDataListener */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mIncomingDataListener */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mIncomingDataListener */
// Alloc allocates a new instance without initialization.
func (mc _mIncomingDataListenerClass) Alloc() mIncomingDataListener {
	rv := objc.Send[mIncomingDataListener](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mIncomingDataListenerClass) New() mIncomingDataListener {
	rv := objc.Send[mIncomingDataListener](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIncomingDataListener) Init() mIncomingDataListener {
	rv := objc.Send[mIncomingDataListener](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIncomingDataListener) Autorelease() mIncomingDataListener {
	rv := objc.Send[mIncomingDataListener](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIncomingDataListener creates a new mIncomingDataListener instance.
func NewmIncomingDataListener() mIncomingDataListener {
	return getmIncomingDataListenerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mIncomingDataListener */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mIncomingDataListener
type mIncomingDataListener struct {
	objectivec.Object
}

// mIncomingDataListenerFrom constructs a [mIncomingDataListener] from an unsafe.Pointer.
func mIncomingDataListenerFrom(ptr unsafe.Pointer) mIncomingDataListener {
	return mIncomingDataListener{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mIncomingDataListener *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mIncomingDataListener */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mIncomingDataListener */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mIncomingDataListener */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mIncomingDataListener */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mIncomingDataListener */




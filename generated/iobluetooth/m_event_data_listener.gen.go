// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mEventDataListener */


/* debug [class_header]: Header for mEventDataListener */
// The class instance for the [mEventDataListener] class.
var (
	MEventDataListenerClass     _mEventDataListenerClass
	MEventDataListenerClassOnce sync.Once
)

func getmEventDataListenerClass() _mEventDataListenerClass {
	MEventDataListenerClassOnce.Do(func() {
		MEventDataListenerClass = _mEventDataListenerClass{objc.GetClass("mEventDataListener")}
	})
	return MEventDataListenerClass
}

type _mEventDataListenerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mEventDataListener */
// An interface definition for the [mEventDataListener] class.
type ImEventDataListener interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mEventDataListener */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mEventDataListener */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mEventDataListener */
// Alloc allocates a new instance without initialization.
func (mc _mEventDataListenerClass) Alloc() mEventDataListener {
	rv := objc.Send[mEventDataListener](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mEventDataListenerClass) New() mEventDataListener {
	rv := objc.Send[mEventDataListener](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mEventDataListener) Init() mEventDataListener {
	rv := objc.Send[mEventDataListener](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mEventDataListener) Autorelease() mEventDataListener {
	rv := objc.Send[mEventDataListener](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmEventDataListener creates a new mEventDataListener instance.
func NewmEventDataListener() mEventDataListener {
	return getmEventDataListenerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mEventDataListener */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mEventDataListener
type mEventDataListener struct {
	objectivec.Object
}

// mEventDataListenerFrom constructs a [mEventDataListener] from an unsafe.Pointer.
func mEventDataListenerFrom(ptr unsafe.Pointer) mEventDataListener {
	return mEventDataListener{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mEventDataListener *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mEventDataListener */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mEventDataListener */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mEventDataListener */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mEventDataListener */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mEventDataListener */




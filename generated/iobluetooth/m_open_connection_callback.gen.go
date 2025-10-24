// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mOpenConnectionCallback */


/* debug [class_header]: Header for mOpenConnectionCallback */
// The class instance for the [mOpenConnectionCallback] class.
var (
	MOpenConnectionCallbackClass     _mOpenConnectionCallbackClass
	MOpenConnectionCallbackClassOnce sync.Once
)

func getmOpenConnectionCallbackClass() _mOpenConnectionCallbackClass {
	MOpenConnectionCallbackClassOnce.Do(func() {
		MOpenConnectionCallbackClass = _mOpenConnectionCallbackClass{objc.GetClass("mOpenConnectionCallback")}
	})
	return MOpenConnectionCallbackClass
}

type _mOpenConnectionCallbackClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mOpenConnectionCallback */
// An interface definition for the [mOpenConnectionCallback] class.
type ImOpenConnectionCallback interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mOpenConnectionCallback */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mOpenConnectionCallback */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mOpenConnectionCallback */
// Alloc allocates a new instance without initialization.
func (mc _mOpenConnectionCallbackClass) Alloc() mOpenConnectionCallback {
	rv := objc.Send[mOpenConnectionCallback](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mOpenConnectionCallbackClass) New() mOpenConnectionCallback {
	rv := objc.Send[mOpenConnectionCallback](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOpenConnectionCallback) Init() mOpenConnectionCallback {
	rv := objc.Send[mOpenConnectionCallback](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOpenConnectionCallback) Autorelease() mOpenConnectionCallback {
	rv := objc.Send[mOpenConnectionCallback](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOpenConnectionCallback creates a new mOpenConnectionCallback instance.
func NewmOpenConnectionCallback() mOpenConnectionCallback {
	return getmOpenConnectionCallbackClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mOpenConnectionCallback */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/mOpenConnectionCallback
type mOpenConnectionCallback struct {
	objectivec.Object
}

// mOpenConnectionCallbackFrom constructs a [mOpenConnectionCallback] from an unsafe.Pointer.
func mOpenConnectionCallbackFrom(ptr unsafe.Pointer) mOpenConnectionCallback {
	return mOpenConnectionCallback{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mOpenConnectionCallback *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mOpenConnectionCallback */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mOpenConnectionCallback */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mOpenConnectionCallback */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mOpenConnectionCallback */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mOpenConnectionCallback */




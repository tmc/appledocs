// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mRFCOMMConnection */


/* debug [class_header]: Header for mRFCOMMConnection */
// The class instance for the [mRFCOMMConnection] class.
var (
	MRFCOMMConnectionClass     _mRFCOMMConnectionClass
	MRFCOMMConnectionClassOnce sync.Once
)

func getmRFCOMMConnectionClass() _mRFCOMMConnectionClass {
	MRFCOMMConnectionClassOnce.Do(func() {
		MRFCOMMConnectionClass = _mRFCOMMConnectionClass{objc.GetClass("mRFCOMMConnection")}
	})
	return MRFCOMMConnectionClass
}

type _mRFCOMMConnectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mRFCOMMConnection */
// An interface definition for the [mRFCOMMConnection] class.
type ImRFCOMMConnection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mRFCOMMConnection */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mRFCOMMConnection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mRFCOMMConnection */
// Alloc allocates a new instance without initialization.
func (mc _mRFCOMMConnectionClass) Alloc() mRFCOMMConnection {
	rv := objc.Send[mRFCOMMConnection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mRFCOMMConnectionClass) New() mRFCOMMConnection {
	rv := objc.Send[mRFCOMMConnection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mRFCOMMConnection) Init() mRFCOMMConnection {
	rv := objc.Send[mRFCOMMConnection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mRFCOMMConnection) Autorelease() mRFCOMMConnection {
	rv := objc.Send[mRFCOMMConnection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmRFCOMMConnection creates a new mRFCOMMConnection instance.
func NewmRFCOMMConnection() mRFCOMMConnection {
	return getmRFCOMMConnectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mRFCOMMConnection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mRFCOMMConnection
type mRFCOMMConnection struct {
	objectivec.Object
}

// mRFCOMMConnectionFrom constructs a [mRFCOMMConnection] from an unsafe.Pointer.
func mRFCOMMConnectionFrom(ptr unsafe.Pointer) mRFCOMMConnection {
	return mRFCOMMConnection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mRFCOMMConnection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mRFCOMMConnection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mRFCOMMConnection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mRFCOMMConnection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mRFCOMMConnection */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mRFCOMMConnection */




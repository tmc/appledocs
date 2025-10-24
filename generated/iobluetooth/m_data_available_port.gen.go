// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mDataAvailablePort */


/* debug [class_header]: Header for mDataAvailablePort */
// The class instance for the [mDataAvailablePort] class.
var (
	MDataAvailablePortClass     _mDataAvailablePortClass
	MDataAvailablePortClassOnce sync.Once
)

func getmDataAvailablePortClass() _mDataAvailablePortClass {
	MDataAvailablePortClassOnce.Do(func() {
		MDataAvailablePortClass = _mDataAvailablePortClass{objc.GetClass("mDataAvailablePort")}
	})
	return MDataAvailablePortClass
}

type _mDataAvailablePortClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mDataAvailablePort */
// An interface definition for the [mDataAvailablePort] class.
type ImDataAvailablePort interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mDataAvailablePort */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mDataAvailablePort */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mDataAvailablePort */
// Alloc allocates a new instance without initialization.
func (mc _mDataAvailablePortClass) Alloc() mDataAvailablePort {
	rv := objc.Send[mDataAvailablePort](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mDataAvailablePortClass) New() mDataAvailablePort {
	rv := objc.Send[mDataAvailablePort](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mDataAvailablePort) Init() mDataAvailablePort {
	rv := objc.Send[mDataAvailablePort](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mDataAvailablePort) Autorelease() mDataAvailablePort {
	rv := objc.Send[mDataAvailablePort](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmDataAvailablePort creates a new mDataAvailablePort instance.
func NewmDataAvailablePort() mDataAvailablePort {
	return getmDataAvailablePortClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mDataAvailablePort */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mDataAvailablePort
type mDataAvailablePort struct {
	objectivec.Object
}

// mDataAvailablePortFrom constructs a [mDataAvailablePort] from an unsafe.Pointer.
func mDataAvailablePortFrom(ptr unsafe.Pointer) mDataAvailablePort {
	return mDataAvailablePort{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mDataAvailablePort *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mDataAvailablePort */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mDataAvailablePort */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mDataAvailablePort */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mDataAvailablePort */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mDataAvailablePort */




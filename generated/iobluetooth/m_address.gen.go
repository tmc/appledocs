// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mAddress */


/* debug [class_header]: Header for mAddress */
// The class instance for the [mAddress] class.
var (
	MAddressClass     _mAddressClass
	MAddressClassOnce sync.Once
)

func getmAddressClass() _mAddressClass {
	MAddressClassOnce.Do(func() {
		MAddressClass = _mAddressClass{objc.GetClass("mAddress")}
	})
	return MAddressClass
}

type _mAddressClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mAddress */
// An interface definition for the [mAddress] class.
type ImAddress interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mAddress */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mAddress */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mAddress */
// Alloc allocates a new instance without initialization.
func (mc _mAddressClass) Alloc() mAddress {
	rv := objc.Send[mAddress](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mAddressClass) New() mAddress {
	rv := objc.Send[mAddress](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mAddress) Init() mAddress {
	rv := objc.Send[mAddress](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mAddress) Autorelease() mAddress {
	rv := objc.Send[mAddress](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmAddress creates a new mAddress instance.
func NewmAddress() mAddress {
	return getmAddressClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mAddress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mAddress
type mAddress struct {
	objectivec.Object
}

// mAddressFrom constructs a [mAddress] from an unsafe.Pointer.
func mAddressFrom(ptr unsafe.Pointer) mAddress {
	return mAddress{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mAddress *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mAddress */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mAddress */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mAddress */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mAddress */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mAddress */




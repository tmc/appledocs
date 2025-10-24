// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mMTU */


/* debug [class_header]: Header for mMTU */
// The class instance for the [mMTU] class.
var (
	MMTUClass     _mMTUClass
	MMTUClassOnce sync.Once
)

func getmMTUClass() _mMTUClass {
	MMTUClassOnce.Do(func() {
		MMTUClass = _mMTUClass{objc.GetClass("mMTU")}
	})
	return MMTUClass
}

type _mMTUClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mMTU */
// An interface definition for the [mMTU] class.
type ImMTU interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mMTU */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mMTU */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mMTU */
// Alloc allocates a new instance without initialization.
func (mc _mMTUClass) Alloc() mMTU {
	rv := objc.Send[mMTU](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mMTUClass) New() mMTU {
	rv := objc.Send[mMTU](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mMTU) Init() mMTU {
	rv := objc.Send[mMTU](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mMTU) Autorelease() mMTU {
	rv := objc.Send[mMTU](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmMTU creates a new mMTU instance.
func NewmMTU() mMTU {
	return getmMTUClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mMTU */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/mMTU
type mMTU struct {
	objectivec.Object
}

// mMTUFrom constructs a [mMTU] from an unsafe.Pointer.
func mMTUFrom(ptr unsafe.Pointer) mMTU {
	return mMTU{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mMTU *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mMTU */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mMTU */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mMTU */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mMTU */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mMTU */




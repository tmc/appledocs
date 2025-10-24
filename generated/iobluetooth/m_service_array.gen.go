// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mServiceArray */


/* debug [class_header]: Header for mServiceArray */
// The class instance for the [mServiceArray] class.
var (
	MServiceArrayClass     _mServiceArrayClass
	MServiceArrayClassOnce sync.Once
)

func getmServiceArrayClass() _mServiceArrayClass {
	MServiceArrayClassOnce.Do(func() {
		MServiceArrayClass = _mServiceArrayClass{objc.GetClass("mServiceArray")}
	})
	return MServiceArrayClass
}

type _mServiceArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mServiceArray */
// An interface definition for the [mServiceArray] class.
type ImServiceArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mServiceArray */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mServiceArray */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mServiceArray */
// Alloc allocates a new instance without initialization.
func (mc _mServiceArrayClass) Alloc() mServiceArray {
	rv := objc.Send[mServiceArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mServiceArrayClass) New() mServiceArray {
	rv := objc.Send[mServiceArray](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mServiceArray) Init() mServiceArray {
	rv := objc.Send[mServiceArray](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mServiceArray) Autorelease() mServiceArray {
	rv := objc.Send[mServiceArray](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmServiceArray creates a new mServiceArray instance.
func NewmServiceArray() mServiceArray {
	return getmServiceArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mServiceArray */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mServiceArray
type mServiceArray struct {
	objectivec.Object
}

// mServiceArrayFrom constructs a [mServiceArray] from an unsafe.Pointer.
func mServiceArrayFrom(ptr unsafe.Pointer) mServiceArray {
	return mServiceArray{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mServiceArray *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mServiceArray */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mServiceArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mServiceArray */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mServiceArray */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mServiceArray */




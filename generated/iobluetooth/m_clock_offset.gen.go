// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mClockOffset */


/* debug [class_header]: Header for mClockOffset */
// The class instance for the [mClockOffset] class.
var (
	MClockOffsetClass     _mClockOffsetClass
	MClockOffsetClassOnce sync.Once
)

func getmClockOffsetClass() _mClockOffsetClass {
	MClockOffsetClassOnce.Do(func() {
		MClockOffsetClass = _mClockOffsetClass{objc.GetClass("mClockOffset")}
	})
	return MClockOffsetClass
}

type _mClockOffsetClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mClockOffset */
// An interface definition for the [mClockOffset] class.
type ImClockOffset interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mClockOffset */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mClockOffset */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mClockOffset */
// Alloc allocates a new instance without initialization.
func (mc _mClockOffsetClass) Alloc() mClockOffset {
	rv := objc.Send[mClockOffset](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mClockOffsetClass) New() mClockOffset {
	rv := objc.Send[mClockOffset](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mClockOffset) Init() mClockOffset {
	rv := objc.Send[mClockOffset](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mClockOffset) Autorelease() mClockOffset {
	rv := objc.Send[mClockOffset](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmClockOffset creates a new mClockOffset instance.
func NewmClockOffset() mClockOffset {
	return getmClockOffsetClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mClockOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mClockOffset
type mClockOffset struct {
	objectivec.Object
}

// mClockOffsetFrom constructs a [mClockOffset] from an unsafe.Pointer.
func mClockOffsetFrom(ptr unsafe.Pointer) mClockOffset {
	return mClockOffset{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mClockOffset *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mClockOffset */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mClockOffset */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mClockOffset */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mClockOffset */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mClockOffset */




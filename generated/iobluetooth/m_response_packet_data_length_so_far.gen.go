// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mResponsePacketDataLengthSoFar */


/* debug [class_header]: Header for mResponsePacketDataLengthSoFar */
// The class instance for the [mResponsePacketDataLengthSoFar] class.
var (
	MResponsePacketDataLengthSoFarClass     _mResponsePacketDataLengthSoFarClass
	MResponsePacketDataLengthSoFarClassOnce sync.Once
)

func getmResponsePacketDataLengthSoFarClass() _mResponsePacketDataLengthSoFarClass {
	MResponsePacketDataLengthSoFarClassOnce.Do(func() {
		MResponsePacketDataLengthSoFarClass = _mResponsePacketDataLengthSoFarClass{objc.GetClass("mResponsePacketDataLengthSoFar")}
	})
	return MResponsePacketDataLengthSoFarClass
}

type _mResponsePacketDataLengthSoFarClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mResponsePacketDataLengthSoFar */
// An interface definition for the [mResponsePacketDataLengthSoFar] class.
type ImResponsePacketDataLengthSoFar interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mResponsePacketDataLengthSoFar */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mResponsePacketDataLengthSoFar */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mResponsePacketDataLengthSoFar */
// Alloc allocates a new instance without initialization.
func (mc _mResponsePacketDataLengthSoFarClass) Alloc() mResponsePacketDataLengthSoFar {
	rv := objc.Send[mResponsePacketDataLengthSoFar](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mResponsePacketDataLengthSoFarClass) New() mResponsePacketDataLengthSoFar {
	rv := objc.Send[mResponsePacketDataLengthSoFar](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mResponsePacketDataLengthSoFar) Init() mResponsePacketDataLengthSoFar {
	rv := objc.Send[mResponsePacketDataLengthSoFar](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mResponsePacketDataLengthSoFar) Autorelease() mResponsePacketDataLengthSoFar {
	rv := objc.Send[mResponsePacketDataLengthSoFar](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmResponsePacketDataLengthSoFar creates a new mResponsePacketDataLengthSoFar instance.
func NewmResponsePacketDataLengthSoFar() mResponsePacketDataLengthSoFar {
	return getmResponsePacketDataLengthSoFarClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mResponsePacketDataLengthSoFar */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mResponsePacketDataLengthSoFar
type mResponsePacketDataLengthSoFar struct {
	objectivec.Object
}

// mResponsePacketDataLengthSoFarFrom constructs a [mResponsePacketDataLengthSoFar] from an unsafe.Pointer.
func mResponsePacketDataLengthSoFarFrom(ptr unsafe.Pointer) mResponsePacketDataLengthSoFar {
	return mResponsePacketDataLengthSoFar{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mResponsePacketDataLengthSoFar *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mResponsePacketDataLengthSoFar */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mResponsePacketDataLengthSoFar */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mResponsePacketDataLengthSoFar */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mResponsePacketDataLengthSoFar */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mResponsePacketDataLengthSoFar */




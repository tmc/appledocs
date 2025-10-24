// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mResponsePacketDataLengthExpected */


/* debug [class_header]: Header for mResponsePacketDataLengthExpected */
// The class instance for the [mResponsePacketDataLengthExpected] class.
var (
	MResponsePacketDataLengthExpectedClass     _mResponsePacketDataLengthExpectedClass
	MResponsePacketDataLengthExpectedClassOnce sync.Once
)

func getmResponsePacketDataLengthExpectedClass() _mResponsePacketDataLengthExpectedClass {
	MResponsePacketDataLengthExpectedClassOnce.Do(func() {
		MResponsePacketDataLengthExpectedClass = _mResponsePacketDataLengthExpectedClass{objc.GetClass("mResponsePacketDataLengthExpected")}
	})
	return MResponsePacketDataLengthExpectedClass
}

type _mResponsePacketDataLengthExpectedClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mResponsePacketDataLengthExpected */
// An interface definition for the [mResponsePacketDataLengthExpected] class.
type ImResponsePacketDataLengthExpected interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mResponsePacketDataLengthExpected */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mResponsePacketDataLengthExpected */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mResponsePacketDataLengthExpected */
// Alloc allocates a new instance without initialization.
func (mc _mResponsePacketDataLengthExpectedClass) Alloc() mResponsePacketDataLengthExpected {
	rv := objc.Send[mResponsePacketDataLengthExpected](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mResponsePacketDataLengthExpectedClass) New() mResponsePacketDataLengthExpected {
	rv := objc.Send[mResponsePacketDataLengthExpected](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mResponsePacketDataLengthExpected) Init() mResponsePacketDataLengthExpected {
	rv := objc.Send[mResponsePacketDataLengthExpected](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mResponsePacketDataLengthExpected) Autorelease() mResponsePacketDataLengthExpected {
	rv := objc.Send[mResponsePacketDataLengthExpected](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmResponsePacketDataLengthExpected creates a new mResponsePacketDataLengthExpected instance.
func NewmResponsePacketDataLengthExpected() mResponsePacketDataLengthExpected {
	return getmResponsePacketDataLengthExpectedClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mResponsePacketDataLengthExpected */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mResponsePacketDataLengthExpected
type mResponsePacketDataLengthExpected struct {
	objectivec.Object
}

// mResponsePacketDataLengthExpectedFrom constructs a [mResponsePacketDataLengthExpected] from an unsafe.Pointer.
func mResponsePacketDataLengthExpectedFrom(ptr unsafe.Pointer) mResponsePacketDataLengthExpected {
	return mResponsePacketDataLengthExpected{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mResponsePacketDataLengthExpected *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mResponsePacketDataLengthExpected */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mResponsePacketDataLengthExpected */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mResponsePacketDataLengthExpected */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mResponsePacketDataLengthExpected */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mResponsePacketDataLengthExpected */




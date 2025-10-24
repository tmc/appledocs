// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mIOConnection */


/* debug [class_header]: Header for mIOConnection */
// The class instance for the [mIOConnection] class.
var (
	MIOConnectionClass     _mIOConnectionClass
	MIOConnectionClassOnce sync.Once
)

func getmIOConnectionClass() _mIOConnectionClass {
	MIOConnectionClassOnce.Do(func() {
		MIOConnectionClass = _mIOConnectionClass{objc.GetClass("mIOConnection")}
	})
	return MIOConnectionClass
}

type _mIOConnectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mIOConnection */
// An interface definition for the [mIOConnection] class.
type ImIOConnection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mIOConnection */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mIOConnection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mIOConnection */
// Alloc allocates a new instance without initialization.
func (mc _mIOConnectionClass) Alloc() mIOConnection {
	rv := objc.Send[mIOConnection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mIOConnectionClass) New() mIOConnection {
	rv := objc.Send[mIOConnection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIOConnection) Init() mIOConnection {
	rv := objc.Send[mIOConnection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIOConnection) Autorelease() mIOConnection {
	rv := objc.Send[mIOConnection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIOConnection creates a new mIOConnection instance.
func NewmIOConnection() mIOConnection {
	return getmIOConnectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mIOConnection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothObject/mIOConnection
type mIOConnection struct {
	objectivec.Object
}

// mIOConnectionFrom constructs a [mIOConnection] from an unsafe.Pointer.
func mIOConnectionFrom(ptr unsafe.Pointer) mIOConnection {
	return mIOConnection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mIOConnection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mIOConnection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mIOConnection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mIOConnection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mIOConnection */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mIOConnection */




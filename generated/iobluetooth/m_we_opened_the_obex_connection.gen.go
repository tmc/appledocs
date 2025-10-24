// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mWeOpenedTheOBEXConnection */


/* debug [class_header]: Header for mWeOpenedTheOBEXConnection */
// The class instance for the [mWeOpenedTheOBEXConnection] class.
var (
	MWeOpenedTheOBEXConnectionClass     _mWeOpenedTheOBEXConnectionClass
	MWeOpenedTheOBEXConnectionClassOnce sync.Once
)

func getmWeOpenedTheOBEXConnectionClass() _mWeOpenedTheOBEXConnectionClass {
	MWeOpenedTheOBEXConnectionClassOnce.Do(func() {
		MWeOpenedTheOBEXConnectionClass = _mWeOpenedTheOBEXConnectionClass{objc.GetClass("mWeOpenedTheOBEXConnection")}
	})
	return MWeOpenedTheOBEXConnectionClass
}

type _mWeOpenedTheOBEXConnectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mWeOpenedTheOBEXConnection */
// An interface definition for the [mWeOpenedTheOBEXConnection] class.
type ImWeOpenedTheOBEXConnection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mWeOpenedTheOBEXConnection */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mWeOpenedTheOBEXConnection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mWeOpenedTheOBEXConnection */
// Alloc allocates a new instance without initialization.
func (mc _mWeOpenedTheOBEXConnectionClass) Alloc() mWeOpenedTheOBEXConnection {
	rv := objc.Send[mWeOpenedTheOBEXConnection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mWeOpenedTheOBEXConnectionClass) New() mWeOpenedTheOBEXConnection {
	rv := objc.Send[mWeOpenedTheOBEXConnection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mWeOpenedTheOBEXConnection) Init() mWeOpenedTheOBEXConnection {
	rv := objc.Send[mWeOpenedTheOBEXConnection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mWeOpenedTheOBEXConnection) Autorelease() mWeOpenedTheOBEXConnection {
	rv := objc.Send[mWeOpenedTheOBEXConnection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmWeOpenedTheOBEXConnection creates a new mWeOpenedTheOBEXConnection instance.
func NewmWeOpenedTheOBEXConnection() mWeOpenedTheOBEXConnection {
	return getmWeOpenedTheOBEXConnectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mWeOpenedTheOBEXConnection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mWeOpenedTheOBEXConnection
type mWeOpenedTheOBEXConnection struct {
	objectivec.Object
}

// mWeOpenedTheOBEXConnectionFrom constructs a [mWeOpenedTheOBEXConnection] from an unsafe.Pointer.
func mWeOpenedTheOBEXConnectionFrom(ptr unsafe.Pointer) mWeOpenedTheOBEXConnection {
	return mWeOpenedTheOBEXConnection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mWeOpenedTheOBEXConnection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mWeOpenedTheOBEXConnection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mWeOpenedTheOBEXConnection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mWeOpenedTheOBEXConnection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mWeOpenedTheOBEXConnection */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mWeOpenedTheOBEXConnection */




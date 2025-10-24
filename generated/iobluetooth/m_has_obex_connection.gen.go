// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mHasOBEXConnection */


/* debug [class_header]: Header for mHasOBEXConnection */
// The class instance for the [mHasOBEXConnection] class.
var (
	MHasOBEXConnectionClass     _mHasOBEXConnectionClass
	MHasOBEXConnectionClassOnce sync.Once
)

func getmHasOBEXConnectionClass() _mHasOBEXConnectionClass {
	MHasOBEXConnectionClassOnce.Do(func() {
		MHasOBEXConnectionClass = _mHasOBEXConnectionClass{objc.GetClass("mHasOBEXConnection")}
	})
	return MHasOBEXConnectionClass
}

type _mHasOBEXConnectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mHasOBEXConnection */
// An interface definition for the [mHasOBEXConnection] class.
type ImHasOBEXConnection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mHasOBEXConnection */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mHasOBEXConnection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mHasOBEXConnection */
// Alloc allocates a new instance without initialization.
func (mc _mHasOBEXConnectionClass) Alloc() mHasOBEXConnection {
	rv := objc.Send[mHasOBEXConnection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mHasOBEXConnectionClass) New() mHasOBEXConnection {
	rv := objc.Send[mHasOBEXConnection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mHasOBEXConnection) Init() mHasOBEXConnection {
	rv := objc.Send[mHasOBEXConnection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mHasOBEXConnection) Autorelease() mHasOBEXConnection {
	rv := objc.Send[mHasOBEXConnection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmHasOBEXConnection creates a new mHasOBEXConnection instance.
func NewmHasOBEXConnection() mHasOBEXConnection {
	return getmHasOBEXConnectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mHasOBEXConnection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mHasOBEXConnection
type mHasOBEXConnection struct {
	objectivec.Object
}

// mHasOBEXConnectionFrom constructs a [mHasOBEXConnection] from an unsafe.Pointer.
func mHasOBEXConnectionFrom(ptr unsafe.Pointer) mHasOBEXConnection {
	return mHasOBEXConnection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mHasOBEXConnection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mHasOBEXConnection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mHasOBEXConnection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mHasOBEXConnection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mHasOBEXConnection */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mHasOBEXConnection */




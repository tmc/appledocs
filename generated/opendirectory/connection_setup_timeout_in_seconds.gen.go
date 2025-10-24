// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class connectionSetupTimeoutInSeconds */


/* debug [class_header]: Header for connectionSetupTimeoutInSeconds */
// The class instance for the [connectionSetupTimeoutInSeconds] class.
var (
	ConnectionSetupTimeoutInSecondsClass     _connectionSetupTimeoutInSecondsClass
	ConnectionSetupTimeoutInSecondsClassOnce sync.Once
)

func getconnectionSetupTimeoutInSecondsClass() _connectionSetupTimeoutInSecondsClass {
	ConnectionSetupTimeoutInSecondsClassOnce.Do(func() {
		ConnectionSetupTimeoutInSecondsClass = _connectionSetupTimeoutInSecondsClass{objc.GetClass("connectionSetupTimeoutInSeconds")}
	})
	return ConnectionSetupTimeoutInSecondsClass
}

type _connectionSetupTimeoutInSecondsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for connectionSetupTimeoutInSeconds */
// An interface definition for the [connectionSetupTimeoutInSeconds] class.
type IconnectionSetupTimeoutInSeconds interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for connectionSetupTimeoutInSeconds */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for connectionSetupTimeoutInSeconds */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for connectionSetupTimeoutInSeconds */
// Alloc allocates a new instance without initialization.
func (cc _connectionSetupTimeoutInSecondsClass) Alloc() connectionSetupTimeoutInSeconds {
	rv := objc.Send[connectionSetupTimeoutInSeconds](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _connectionSetupTimeoutInSecondsClass) New() connectionSetupTimeoutInSeconds {
	rv := objc.Send[connectionSetupTimeoutInSeconds](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ connectionSetupTimeoutInSeconds) Init() connectionSetupTimeoutInSeconds {
	rv := objc.Send[connectionSetupTimeoutInSeconds](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ connectionSetupTimeoutInSeconds) Autorelease() connectionSetupTimeoutInSeconds {
	rv := objc.Send[connectionSetupTimeoutInSeconds](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewconnectionSetupTimeoutInSeconds creates a new connectionSetupTimeoutInSeconds instance.
func NewconnectionSetupTimeoutInSeconds() connectionSetupTimeoutInSeconds {
	return getconnectionSetupTimeoutInSecondsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for connectionSetupTimeoutInSeconds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionSetupTimeoutInSeconds-c.ivar
type connectionSetupTimeoutInSeconds struct {
	objectivec.Object
}

// connectionSetupTimeoutInSecondsFrom constructs a [connectionSetupTimeoutInSeconds] from an unsafe.Pointer.
func connectionSetupTimeoutInSecondsFrom(ptr unsafe.Pointer) connectionSetupTimeoutInSeconds {
	return connectionSetupTimeoutInSeconds{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for connectionSetupTimeoutInSeconds *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for connectionSetupTimeoutInSeconds */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for connectionSetupTimeoutInSeconds */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for connectionSetupTimeoutInSeconds */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for connectionSetupTimeoutInSeconds */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class connectionSetupTimeoutInSeconds */




// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class connectionIdleTimeoutInSeconds */


/* debug [class_header]: Header for connectionIdleTimeoutInSeconds */
// The class instance for the [connectionIdleTimeoutInSeconds] class.
var (
	ConnectionIdleTimeoutInSecondsClass     _connectionIdleTimeoutInSecondsClass
	ConnectionIdleTimeoutInSecondsClassOnce sync.Once
)

func getconnectionIdleTimeoutInSecondsClass() _connectionIdleTimeoutInSecondsClass {
	ConnectionIdleTimeoutInSecondsClassOnce.Do(func() {
		ConnectionIdleTimeoutInSecondsClass = _connectionIdleTimeoutInSecondsClass{objc.GetClass("connectionIdleTimeoutInSeconds")}
	})
	return ConnectionIdleTimeoutInSecondsClass
}

type _connectionIdleTimeoutInSecondsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for connectionIdleTimeoutInSeconds */
// An interface definition for the [connectionIdleTimeoutInSeconds] class.
type IconnectionIdleTimeoutInSeconds interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for connectionIdleTimeoutInSeconds */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for connectionIdleTimeoutInSeconds */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for connectionIdleTimeoutInSeconds */
// Alloc allocates a new instance without initialization.
func (cc _connectionIdleTimeoutInSecondsClass) Alloc() connectionIdleTimeoutInSeconds {
	rv := objc.Send[connectionIdleTimeoutInSeconds](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _connectionIdleTimeoutInSecondsClass) New() connectionIdleTimeoutInSeconds {
	rv := objc.Send[connectionIdleTimeoutInSeconds](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ connectionIdleTimeoutInSeconds) Init() connectionIdleTimeoutInSeconds {
	rv := objc.Send[connectionIdleTimeoutInSeconds](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ connectionIdleTimeoutInSeconds) Autorelease() connectionIdleTimeoutInSeconds {
	rv := objc.Send[connectionIdleTimeoutInSeconds](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewconnectionIdleTimeoutInSeconds creates a new connectionIdleTimeoutInSeconds instance.
func NewconnectionIdleTimeoutInSeconds() connectionIdleTimeoutInSeconds {
	return getconnectionIdleTimeoutInSecondsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for connectionIdleTimeoutInSeconds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionIdleTimeoutInSeconds-c.ivar
type connectionIdleTimeoutInSeconds struct {
	objectivec.Object
}

// connectionIdleTimeoutInSecondsFrom constructs a [connectionIdleTimeoutInSeconds] from an unsafe.Pointer.
func connectionIdleTimeoutInSecondsFrom(ptr unsafe.Pointer) connectionIdleTimeoutInSeconds {
	return connectionIdleTimeoutInSeconds{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for connectionIdleTimeoutInSeconds *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for connectionIdleTimeoutInSeconds */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for connectionIdleTimeoutInSeconds */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for connectionIdleTimeoutInSeconds */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for connectionIdleTimeoutInSeconds */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class connectionIdleTimeoutInSeconds */




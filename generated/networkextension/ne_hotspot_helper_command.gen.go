// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEHotspotHelperCommand */


/* debug [class_header]: Header for NEHotspotHelperCommand */
// The class instance for the [NEHotspotHelperCommand] class.
var (
	NEHotspotHelperCommandClass     _NEHotspotHelperCommandClass
	NEHotspotHelperCommandClassOnce sync.Once
)

func getNEHotspotHelperCommandClass() _NEHotspotHelperCommandClass {
	NEHotspotHelperCommandClassOnce.Do(func() {
		NEHotspotHelperCommandClass = _NEHotspotHelperCommandClass{objc.GetClass("NEHotspotHelperCommand")}
	})
	return NEHotspotHelperCommandClass
}

type _NEHotspotHelperCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEHotspotHelperCommand */
// An interface definition for the [NEHotspotHelperCommand] class.
type INEHotspotHelperCommand interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEHotspotHelperCommand */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEHotspotHelperCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEHotspotHelperCommand */
// Alloc allocates a new instance without initialization.
func (nc _NEHotspotHelperCommandClass) Alloc() NEHotspotHelperCommand {
	rv := objc.Send[NEHotspotHelperCommand](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEHotspotHelperCommandClass) New() NEHotspotHelperCommand {
	rv := objc.Send[NEHotspotHelperCommand](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEHotspotHelperCommand) Init() NEHotspotHelperCommand {
	rv := objc.Send[NEHotspotHelperCommand](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEHotspotHelperCommand) Autorelease() NEHotspotHelperCommand {
	rv := objc.Send[NEHotspotHelperCommand](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEHotspotHelperCommand creates a new NEHotspotHelperCommand instance.
func NewNEHotspotHelperCommand() NEHotspotHelperCommand {
	return getNEHotspotHelperCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEHotspotHelperCommand */
// A command for the hotspot helper to handle.
//
// NEHotspostHelperCommand objects are passed to the the Hotspot Helper app’s command handler block. The Hotspot Helper app processes the command, instantiates an object, sets the annotated or ( or commands only), and then delivers the response to the system.


// A command for the hotspot helper to handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommand
type NEHotspotHelperCommand struct {
	objectivec.Object
}

// NEHotspotHelperCommandFrom constructs a [NEHotspotHelperCommand] from an unsafe.Pointer.
//
// A command for the hotspot helper to handle.
func NEHotspotHelperCommandFrom(ptr unsafe.Pointer) NEHotspotHelperCommand {
	return NEHotspotHelperCommand{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEHotspotHelperCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEHotspotHelperCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEHotspotHelperCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEHotspotHelperCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEHotspotHelperCommand */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEHotspotHelperCommand */



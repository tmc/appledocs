// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterMessages */


/* debug [class_header]: Header for MTRClusterMessages */
// The class instance for the [MTRClusterMessages] class.
var (
	MTRClusterMessagesClass     _MTRClusterMessagesClass
	MTRClusterMessagesClassOnce sync.Once
)

func getMTRClusterMessagesClass() _MTRClusterMessagesClass {
	MTRClusterMessagesClassOnce.Do(func() {
		MTRClusterMessagesClass = _MTRClusterMessagesClass{objc.GetClass("MTRClusterMessages")}
	})
	return MTRClusterMessagesClass
}

type _MTRClusterMessagesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterMessages */
// An interface definition for the [MTRClusterMessages] class.
type IMTRClusterMessages interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterMessages */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterMessages */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterMessages */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterMessagesClass) Alloc() MTRClusterMessages {
	rv := objc.Send[MTRClusterMessages](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterMessagesClass) New() MTRClusterMessages {
	rv := objc.Send[MTRClusterMessages](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterMessages) Init() MTRClusterMessages {
	rv := objc.Send[MTRClusterMessages](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterMessages) Autorelease() MTRClusterMessages {
	rv := objc.Send[MTRClusterMessages](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterMessages creates a new MTRClusterMessages instance.
func NewMTRClusterMessages() MTRClusterMessages {
	return getMTRClusterMessagesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterMessages */
// Cluster Messages This cluster provides an interface for passing messages to be presented by a device.


// Cluster Messages This cluster provides an interface for passing messages to be presented by a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMessages
type MTRClusterMessages struct {
	MTRGenericCluster
}

// MTRClusterMessagesFrom constructs a [MTRClusterMessages] from an unsafe.Pointer.
//
// Cluster Messages This cluster provides an interface for passing messages to be presented by a device.
func MTRClusterMessagesFrom(ptr unsafe.Pointer) MTRClusterMessages {
	return MTRClusterMessages{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterMessages */

// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMessages/init(device:endpointID:queue:)
func NewMTRClusterMessagesWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterMessages {
	instance := getMTRClusterMessagesClass().Alloc()
	rv := objc.Send[MTRClusterMessages](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRClusterMessagesWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterMessages */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterMessages */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterMessages */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterMessages */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterMessages */



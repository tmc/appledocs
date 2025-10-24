// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterContentAppObserver */


/* debug [class_header]: Header for MTRBaseClusterContentAppObserver */
// The class instance for the [MTRBaseClusterContentAppObserver] class.
var (
	MTRBaseClusterContentAppObserverClass     _MTRBaseClusterContentAppObserverClass
	MTRBaseClusterContentAppObserverClassOnce sync.Once
)

func getMTRBaseClusterContentAppObserverClass() _MTRBaseClusterContentAppObserverClass {
	MTRBaseClusterContentAppObserverClassOnce.Do(func() {
		MTRBaseClusterContentAppObserverClass = _MTRBaseClusterContentAppObserverClass{objc.GetClass("MTRBaseClusterContentAppObserver")}
	})
	return MTRBaseClusterContentAppObserverClass
}

type _MTRBaseClusterContentAppObserverClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterContentAppObserver */
// An interface definition for the [MTRBaseClusterContentAppObserver] class.
type IMTRBaseClusterContentAppObserver interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterContentAppObserver */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterContentAppObserver */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterContentAppObserver */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterContentAppObserverClass) Alloc() MTRBaseClusterContentAppObserver {
	rv := objc.Send[MTRBaseClusterContentAppObserver](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterContentAppObserverClass) New() MTRBaseClusterContentAppObserver {
	rv := objc.Send[MTRBaseClusterContentAppObserver](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterContentAppObserver) Init() MTRBaseClusterContentAppObserver {
	rv := objc.Send[MTRBaseClusterContentAppObserver](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterContentAppObserver) Autorelease() MTRBaseClusterContentAppObserver {
	rv := objc.Send[MTRBaseClusterContentAppObserver](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterContentAppObserver creates a new MTRBaseClusterContentAppObserver instance.
func NewMTRBaseClusterContentAppObserver() MTRBaseClusterContentAppObserver {
	return getMTRBaseClusterContentAppObserverClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterContentAppObserver */
// Cluster Content App Observer
//
// This cluster provides an interface for sending targeted commands to an Observer of a Content App on a Video Player device such as a Streaming Media Player, Smart TV or Smart Screen. The cluster server for Content App Observer is implemented by an endpoint that communicates with a Content App, such as a Casting Video Client. The cluster client for Content App Observer is implemented by a Content App endpoint. A Content App is informed of the NodeId of an Observer when a binding is set on the Content App. The Content App can then send the ContentAppMessage to the Observer (server cluster), and the Observer responds with a ContentAppMessageResponse.


// Cluster Content App Observer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver
type MTRBaseClusterContentAppObserver struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterContentAppObserverFrom constructs a [MTRBaseClusterContentAppObserver] from an unsafe.Pointer.
//
// Cluster Content App Observer
func MTRBaseClusterContentAppObserverFrom(ptr unsafe.Pointer) MTRBaseClusterContentAppObserver {
	return MTRBaseClusterContentAppObserver{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterContentAppObserver */

// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/init(device:endpointID:queue:)
func NewMTRBaseClusterContentAppObserverWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterContentAppObserver {
	instance := getMTRBaseClusterContentAppObserverClass().Alloc()
	rv := objc.Send[MTRBaseClusterContentAppObserver](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterContentAppObserverWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterContentAppObserver */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterContentAppObserver */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterContentAppObserver */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterContentAppObserver */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterContentAppObserver */



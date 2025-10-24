// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterTimeSynchronization */


/* debug [class_header]: Header for MTRClusterTimeSynchronization */
// The class instance for the [MTRClusterTimeSynchronization] class.
var (
	MTRClusterTimeSynchronizationClass     _MTRClusterTimeSynchronizationClass
	MTRClusterTimeSynchronizationClassOnce sync.Once
)

func getMTRClusterTimeSynchronizationClass() _MTRClusterTimeSynchronizationClass {
	MTRClusterTimeSynchronizationClassOnce.Do(func() {
		MTRClusterTimeSynchronizationClass = _MTRClusterTimeSynchronizationClass{objc.GetClass("MTRClusterTimeSynchronization")}
	})
	return MTRClusterTimeSynchronizationClass
}

type _MTRClusterTimeSynchronizationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterTimeSynchronization */
// An interface definition for the [MTRClusterTimeSynchronization] class.
type IMTRClusterTimeSynchronization interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterTimeSynchronization */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterTimeSynchronization */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterTimeSynchronization */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterTimeSynchronizationClass) Alloc() MTRClusterTimeSynchronization {
	rv := objc.Send[MTRClusterTimeSynchronization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterTimeSynchronizationClass) New() MTRClusterTimeSynchronization {
	rv := objc.Send[MTRClusterTimeSynchronization](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterTimeSynchronization) Init() MTRClusterTimeSynchronization {
	rv := objc.Send[MTRClusterTimeSynchronization](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterTimeSynchronization) Autorelease() MTRClusterTimeSynchronization {
	rv := objc.Send[MTRClusterTimeSynchronization](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterTimeSynchronization creates a new MTRClusterTimeSynchronization instance.
func NewMTRClusterTimeSynchronization() MTRClusterTimeSynchronization {
	return getMTRClusterTimeSynchronizationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterTimeSynchronization */
// Cluster Time Synchronization Accurate time is required for a number of reasons, including scheduling, display and validating security materials.


// Cluster Time Synchronization Accurate time is required for a number of reasons, including scheduling, display and validating security materials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization
type MTRClusterTimeSynchronization struct {
	MTRGenericCluster
}

// MTRClusterTimeSynchronizationFrom constructs a [MTRClusterTimeSynchronization] from an unsafe.Pointer.
//
// Cluster Time Synchronization Accurate time is required for a number of reasons, including scheduling, display and validating security materials.
func MTRClusterTimeSynchronizationFrom(ptr unsafe.Pointer) MTRClusterTimeSynchronization {
	return MTRClusterTimeSynchronization{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterTimeSynchronization */

// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/init(device:endpointID:queue:)
func NewMTRClusterTimeSynchronizationWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterTimeSynchronization {
	instance := getMTRClusterTimeSynchronizationClass().Alloc()
	rv := objc.Send[MTRClusterTimeSynchronization](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRClusterTimeSynchronizationWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterTimeSynchronization */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterTimeSynchronization */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterTimeSynchronization */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterTimeSynchronization */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterTimeSynchronization */



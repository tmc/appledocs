// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterOvenMode */


/* debug [class_header]: Header for MTRBaseClusterOvenMode */
// The class instance for the [MTRBaseClusterOvenMode] class.
var (
	MTRBaseClusterOvenModeClass     _MTRBaseClusterOvenModeClass
	MTRBaseClusterOvenModeClassOnce sync.Once
)

func getMTRBaseClusterOvenModeClass() _MTRBaseClusterOvenModeClass {
	MTRBaseClusterOvenModeClassOnce.Do(func() {
		MTRBaseClusterOvenModeClass = _MTRBaseClusterOvenModeClass{objc.GetClass("MTRBaseClusterOvenMode")}
	})
	return MTRBaseClusterOvenModeClass
}

type _MTRBaseClusterOvenModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterOvenMode */
// An interface definition for the [MTRBaseClusterOvenMode] class.
type IMTRBaseClusterOvenMode interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterOvenMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterOvenMode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterOvenMode */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterOvenModeClass) Alloc() MTRBaseClusterOvenMode {
	rv := objc.Send[MTRBaseClusterOvenMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterOvenModeClass) New() MTRBaseClusterOvenMode {
	rv := objc.Send[MTRBaseClusterOvenMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterOvenMode) Init() MTRBaseClusterOvenMode {
	rv := objc.Send[MTRBaseClusterOvenMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterOvenMode) Autorelease() MTRBaseClusterOvenMode {
	rv := objc.Send[MTRBaseClusterOvenMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterOvenMode creates a new MTRBaseClusterOvenMode instance.
func NewMTRBaseClusterOvenMode() MTRBaseClusterOvenMode {
	return getMTRBaseClusterOvenModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterOvenMode */
// Cluster Oven Mode
//
// Attributes and commands for selecting a mode from a list of supported options.


// Cluster Oven Mode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenMode
type MTRBaseClusterOvenMode struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterOvenModeFrom constructs a [MTRBaseClusterOvenMode] from an unsafe.Pointer.
//
// Cluster Oven Mode
func MTRBaseClusterOvenModeFrom(ptr unsafe.Pointer) MTRBaseClusterOvenMode {
	return MTRBaseClusterOvenMode{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterOvenMode */

// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenMode/init(device:endpointID:queue:)
func NewMTRBaseClusterOvenModeWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterOvenMode {
	instance := getMTRBaseClusterOvenModeClass().Alloc()
	rv := objc.Send[MTRBaseClusterOvenMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterOvenModeWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterOvenMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterOvenMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterOvenMode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterOvenMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterOvenMode */



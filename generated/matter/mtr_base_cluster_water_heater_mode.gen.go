// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterWaterHeaterMode */


/* debug [class_header]: Header for MTRBaseClusterWaterHeaterMode */
// The class instance for the [MTRBaseClusterWaterHeaterMode] class.
var (
	MTRBaseClusterWaterHeaterModeClass     _MTRBaseClusterWaterHeaterModeClass
	MTRBaseClusterWaterHeaterModeClassOnce sync.Once
)

func getMTRBaseClusterWaterHeaterModeClass() _MTRBaseClusterWaterHeaterModeClass {
	MTRBaseClusterWaterHeaterModeClassOnce.Do(func() {
		MTRBaseClusterWaterHeaterModeClass = _MTRBaseClusterWaterHeaterModeClass{objc.GetClass("MTRBaseClusterWaterHeaterMode")}
	})
	return MTRBaseClusterWaterHeaterModeClass
}

type _MTRBaseClusterWaterHeaterModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterWaterHeaterMode */
// An interface definition for the [MTRBaseClusterWaterHeaterMode] class.
type IMTRBaseClusterWaterHeaterMode interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterWaterHeaterMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterWaterHeaterMode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterWaterHeaterMode */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterWaterHeaterModeClass) Alloc() MTRBaseClusterWaterHeaterMode {
	rv := objc.Send[MTRBaseClusterWaterHeaterMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterWaterHeaterModeClass) New() MTRBaseClusterWaterHeaterMode {
	rv := objc.Send[MTRBaseClusterWaterHeaterMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterWaterHeaterMode) Init() MTRBaseClusterWaterHeaterMode {
	rv := objc.Send[MTRBaseClusterWaterHeaterMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterWaterHeaterMode) Autorelease() MTRBaseClusterWaterHeaterMode {
	rv := objc.Send[MTRBaseClusterWaterHeaterMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterWaterHeaterMode creates a new MTRBaseClusterWaterHeaterMode instance.
func NewMTRBaseClusterWaterHeaterMode() MTRBaseClusterWaterHeaterMode {
	return getMTRBaseClusterWaterHeaterModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterWaterHeaterMode */
// Cluster Water Heater Mode
//
// Attributes and commands for selecting a mode from a list of supported options.


// Cluster Water Heater Mode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode
type MTRBaseClusterWaterHeaterMode struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterWaterHeaterModeFrom constructs a [MTRBaseClusterWaterHeaterMode] from an unsafe.Pointer.
//
// Cluster Water Heater Mode
func MTRBaseClusterWaterHeaterModeFrom(ptr unsafe.Pointer) MTRBaseClusterWaterHeaterMode {
	return MTRBaseClusterWaterHeaterMode{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterWaterHeaterMode */

// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/init(device:endpointID:queue:)
func NewMTRBaseClusterWaterHeaterModeWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterWaterHeaterMode {
	instance := getMTRBaseClusterWaterHeaterModeClass().Alloc()
	rv := objc.Send[MTRBaseClusterWaterHeaterMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterWaterHeaterModeWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterWaterHeaterMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterWaterHeaterMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterWaterHeaterMode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterWaterHeaterMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterWaterHeaterMode */



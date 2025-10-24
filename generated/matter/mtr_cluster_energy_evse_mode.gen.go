// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterEnergyEVSEMode */


/* debug [class_header]: Header for MTRClusterEnergyEVSEMode */
// The class instance for the [MTRClusterEnergyEVSEMode] class.
var (
	MTRClusterEnergyEVSEModeClass     _MTRClusterEnergyEVSEModeClass
	MTRClusterEnergyEVSEModeClassOnce sync.Once
)

func getMTRClusterEnergyEVSEModeClass() _MTRClusterEnergyEVSEModeClass {
	MTRClusterEnergyEVSEModeClassOnce.Do(func() {
		MTRClusterEnergyEVSEModeClass = _MTRClusterEnergyEVSEModeClass{objc.GetClass("MTRClusterEnergyEVSEMode")}
	})
	return MTRClusterEnergyEVSEModeClass
}

type _MTRClusterEnergyEVSEModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterEnergyEVSEMode */
// An interface definition for the [MTRClusterEnergyEVSEMode] class.
type IMTRClusterEnergyEVSEMode interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterEnergyEVSEMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterEnergyEVSEMode */
	// methods:
	ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEModeClusterChangeToModeParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterEnergyEVSEMode */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterEnergyEVSEModeClass) Alloc() MTRClusterEnergyEVSEMode {
	rv := objc.Send[MTRClusterEnergyEVSEMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterEnergyEVSEModeClass) New() MTRClusterEnergyEVSEMode {
	rv := objc.Send[MTRClusterEnergyEVSEMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterEnergyEVSEMode) Init() MTRClusterEnergyEVSEMode {
	rv := objc.Send[MTRClusterEnergyEVSEMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterEnergyEVSEMode) Autorelease() MTRClusterEnergyEVSEMode {
	rv := objc.Send[MTRClusterEnergyEVSEMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterEnergyEVSEMode creates a new MTRClusterEnergyEVSEMode instance.
func NewMTRClusterEnergyEVSEMode() MTRClusterEnergyEVSEMode {
	return getMTRClusterEnergyEVSEModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterEnergyEVSEMode */
// Cluster Energy EVSE Mode Attributes and commands for selecting a mode from a list of supported options.


// Cluster Energy EVSE Mode Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSEMode
type MTRClusterEnergyEVSEMode struct {
	MTRGenericCluster
}

// MTRClusterEnergyEVSEModeFrom constructs a [MTRClusterEnergyEVSEMode] from an unsafe.Pointer.
//
// Cluster Energy EVSE Mode Attributes and commands for selecting a mode from a list of supported options.
func MTRClusterEnergyEVSEModeFrom(ptr unsafe.Pointer) MTRClusterEnergyEVSEMode {
	return MTRClusterEnergyEVSEMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterEnergyEVSEMode */

// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSEMode/init(device:endpointID:queue:)
func NewMTRClusterEnergyEVSEModeWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterEnergyEVSEMode {
	instance := getMTRClusterEnergyEVSEModeClass().Alloc()
	rv := objc.Send[MTRClusterEnergyEVSEMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRClusterEnergyEVSEModeWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterEnergyEVSEMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterEnergyEVSEMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterEnergyEVSEMode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSEMode/changeToMode(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSEMode) ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEModeClusterChangeToModeParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("changeToModeWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}/* debug [instance_methods/method]: ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSEMode/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterEnergyEVSEMode) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}/* debug [instance_methods/method]: ReadAttributeAcceptedCommandListWithParams */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterEnergyEVSEMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterEnergyEVSEMode */



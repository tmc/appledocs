// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams */


/* debug [class_header]: Header for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams */
// The class instance for the [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams] class.
var (
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass     _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass() _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass {
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass = _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass{objc.GetClass("MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams")}
	})
	return MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass
}

type _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams */
// An interface definition for the [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams] class.
type IMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams */
	// properties:
	NewMode() objc.IObject /* cross-framework: NSNumber */
	SetNewMode(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass) Alloc() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass) New() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams) Init() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams) Autorelease() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams creates a new MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams instance.
func NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams {
	return getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams
type MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsFrom constructs a [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams {
	return MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams/newMode
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams) NewMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newMode"))
	return rv
}/* debug [instance_properties/getter]: newMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams/newMode
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams) SetNewMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}/* debug [instance_properties/setter]: newMode */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrefrigeratorandtemperaturecontrolledcabinetmodeclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrefrigeratorandtemperaturecontrolledcabinetmodeclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrefrigeratorandtemperaturecontrolledcabinetmodeclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrefrigeratorandtemperaturecontrolledcabinetmodeclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams */




// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROvenModeClusterChangeToModeParams */


/* debug [class_header]: Header for MTROvenModeClusterChangeToModeParams */
// The class instance for the [MTROvenModeClusterChangeToModeParams] class.
var (
	MTROvenModeClusterChangeToModeParamsClass     _MTROvenModeClusterChangeToModeParamsClass
	MTROvenModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTROvenModeClusterChangeToModeParamsClass() _MTROvenModeClusterChangeToModeParamsClass {
	MTROvenModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTROvenModeClusterChangeToModeParamsClass = _MTROvenModeClusterChangeToModeParamsClass{objc.GetClass("MTROvenModeClusterChangeToModeParams")}
	})
	return MTROvenModeClusterChangeToModeParamsClass
}

type _MTROvenModeClusterChangeToModeParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROvenModeClusterChangeToModeParams */
// An interface definition for the [MTROvenModeClusterChangeToModeParams] class.
type IMTROvenModeClusterChangeToModeParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROvenModeClusterChangeToModeParams */
	// properties:
	NewMode() objc.IObject /* cross-framework: NSNumber */
	SetNewMode(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROvenModeClusterChangeToModeParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROvenModeClusterChangeToModeParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROvenModeClusterChangeToModeParamsClass) Alloc() MTROvenModeClusterChangeToModeParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROvenModeClusterChangeToModeParamsClass) New() MTROvenModeClusterChangeToModeParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenModeClusterChangeToModeParams) Init() MTROvenModeClusterChangeToModeParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenModeClusterChangeToModeParams) Autorelease() MTROvenModeClusterChangeToModeParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenModeClusterChangeToModeParams creates a new MTROvenModeClusterChangeToModeParams instance.
func NewMTROvenModeClusterChangeToModeParams() MTROvenModeClusterChangeToModeParams {
	return getMTROvenModeClusterChangeToModeParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROvenModeClusterChangeToModeParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams
type MTROvenModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTROvenModeClusterChangeToModeParamsFrom constructs a [MTROvenModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTROvenModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTROvenModeClusterChangeToModeParams {
	return MTROvenModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROvenModeClusterChangeToModeParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROvenModeClusterChangeToModeParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROvenModeClusterChangeToModeParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROvenModeClusterChangeToModeParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROvenModeClusterChangeToModeParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams/newMode
func (m_ MTROvenModeClusterChangeToModeParams) NewMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newMode"))
	return rv
}/* debug [instance_properties/getter]: newMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams/newMode
func (m_ MTROvenModeClusterChangeToModeParams) SetNewMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}/* debug [instance_properties/setter]: newMode */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovenmodeclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTROvenModeClusterChangeToModeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovenmodeclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTROvenModeClusterChangeToModeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovenmodeclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTROvenModeClusterChangeToModeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovenmodeclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTROvenModeClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROvenModeClusterChangeToModeParams */




// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDishwasherModeClusterChangeToModeParams */


/* debug [class_header]: Header for MTRDishwasherModeClusterChangeToModeParams */
// The class instance for the [MTRDishwasherModeClusterChangeToModeParams] class.
var (
	MTRDishwasherModeClusterChangeToModeParamsClass     _MTRDishwasherModeClusterChangeToModeParamsClass
	MTRDishwasherModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTRDishwasherModeClusterChangeToModeParamsClass() _MTRDishwasherModeClusterChangeToModeParamsClass {
	MTRDishwasherModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTRDishwasherModeClusterChangeToModeParamsClass = _MTRDishwasherModeClusterChangeToModeParamsClass{objc.GetClass("MTRDishwasherModeClusterChangeToModeParams")}
	})
	return MTRDishwasherModeClusterChangeToModeParamsClass
}

type _MTRDishwasherModeClusterChangeToModeParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDishwasherModeClusterChangeToModeParams */
// An interface definition for the [MTRDishwasherModeClusterChangeToModeParams] class.
type IMTRDishwasherModeClusterChangeToModeParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDishwasherModeClusterChangeToModeParams */
	// properties:
	NewMode() objc.IObject /* cross-framework: NSNumber */
	SetNewMode(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDishwasherModeClusterChangeToModeParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDishwasherModeClusterChangeToModeParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDishwasherModeClusterChangeToModeParamsClass) Alloc() MTRDishwasherModeClusterChangeToModeParams {
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDishwasherModeClusterChangeToModeParamsClass) New() MTRDishwasherModeClusterChangeToModeParams {
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDishwasherModeClusterChangeToModeParams) Init() MTRDishwasherModeClusterChangeToModeParams {
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDishwasherModeClusterChangeToModeParams) Autorelease() MTRDishwasherModeClusterChangeToModeParams {
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDishwasherModeClusterChangeToModeParams creates a new MTRDishwasherModeClusterChangeToModeParams instance.
func NewMTRDishwasherModeClusterChangeToModeParams() MTRDishwasherModeClusterChangeToModeParams {
	return getMTRDishwasherModeClusterChangeToModeParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDishwasherModeClusterChangeToModeParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeParams
type MTRDishwasherModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTRDishwasherModeClusterChangeToModeParamsFrom constructs a [MTRDishwasherModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTRDishwasherModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTRDishwasherModeClusterChangeToModeParams {
	return MTRDishwasherModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDishwasherModeClusterChangeToModeParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDishwasherModeClusterChangeToModeParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDishwasherModeClusterChangeToModeParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDishwasherModeClusterChangeToModeParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDishwasherModeClusterChangeToModeParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeParams/newMode
func (m_ MTRDishwasherModeClusterChangeToModeParams) NewMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newMode"))
	return rv
}/* debug [instance_properties/getter]: newMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeParams/newMode
func (m_ MTRDishwasherModeClusterChangeToModeParams) SetNewMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}/* debug [instance_properties/setter]: newMode */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwashermodeclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTRDishwasherModeClusterChangeToModeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwashermodeclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTRDishwasherModeClusterChangeToModeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwashermodeclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTRDishwasherModeClusterChangeToModeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwashermodeclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTRDishwasherModeClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDishwasherModeClusterChangeToModeParams */




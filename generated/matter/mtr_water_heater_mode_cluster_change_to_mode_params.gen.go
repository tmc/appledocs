// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWaterHeaterModeClusterChangeToModeParams */


/* debug [class_header]: Header for MTRWaterHeaterModeClusterChangeToModeParams */
// The class instance for the [MTRWaterHeaterModeClusterChangeToModeParams] class.
var (
	MTRWaterHeaterModeClusterChangeToModeParamsClass     _MTRWaterHeaterModeClusterChangeToModeParamsClass
	MTRWaterHeaterModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTRWaterHeaterModeClusterChangeToModeParamsClass() _MTRWaterHeaterModeClusterChangeToModeParamsClass {
	MTRWaterHeaterModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTRWaterHeaterModeClusterChangeToModeParamsClass = _MTRWaterHeaterModeClusterChangeToModeParamsClass{objc.GetClass("MTRWaterHeaterModeClusterChangeToModeParams")}
	})
	return MTRWaterHeaterModeClusterChangeToModeParamsClass
}

type _MTRWaterHeaterModeClusterChangeToModeParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWaterHeaterModeClusterChangeToModeParams */
// An interface definition for the [MTRWaterHeaterModeClusterChangeToModeParams] class.
type IMTRWaterHeaterModeClusterChangeToModeParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWaterHeaterModeClusterChangeToModeParams */
	// properties:
	NewMode() objc.IObject /* cross-framework: NSNumber */
	SetNewMode(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWaterHeaterModeClusterChangeToModeParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWaterHeaterModeClusterChangeToModeParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterModeClusterChangeToModeParamsClass) Alloc() MTRWaterHeaterModeClusterChangeToModeParams {
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWaterHeaterModeClusterChangeToModeParamsClass) New() MTRWaterHeaterModeClusterChangeToModeParams {
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterModeClusterChangeToModeParams) Init() MTRWaterHeaterModeClusterChangeToModeParams {
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterModeClusterChangeToModeParams) Autorelease() MTRWaterHeaterModeClusterChangeToModeParams {
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterModeClusterChangeToModeParams creates a new MTRWaterHeaterModeClusterChangeToModeParams instance.
func NewMTRWaterHeaterModeClusterChangeToModeParams() MTRWaterHeaterModeClusterChangeToModeParams {
	return getMTRWaterHeaterModeClusterChangeToModeParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWaterHeaterModeClusterChangeToModeParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeParams
type MTRWaterHeaterModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTRWaterHeaterModeClusterChangeToModeParamsFrom constructs a [MTRWaterHeaterModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTRWaterHeaterModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTRWaterHeaterModeClusterChangeToModeParams {
	return MTRWaterHeaterModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWaterHeaterModeClusterChangeToModeParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWaterHeaterModeClusterChangeToModeParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWaterHeaterModeClusterChangeToModeParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWaterHeaterModeClusterChangeToModeParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWaterHeaterModeClusterChangeToModeParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeParams/newMode
func (m_ MTRWaterHeaterModeClusterChangeToModeParams) NewMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newMode"))
	return rv
}/* debug [instance_properties/getter]: newMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeParams/newMode
func (m_ MTRWaterHeaterModeClusterChangeToModeParams) SetNewMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}/* debug [instance_properties/setter]: newMode */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermodeclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTRWaterHeaterModeClusterChangeToModeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermodeclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTRWaterHeaterModeClusterChangeToModeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermodeclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTRWaterHeaterModeClusterChangeToModeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermodeclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTRWaterHeaterModeClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWaterHeaterModeClusterChangeToModeParams */




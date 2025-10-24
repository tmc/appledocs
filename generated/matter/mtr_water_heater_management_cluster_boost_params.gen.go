// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWaterHeaterManagementClusterBoostParams */


/* debug [class_header]: Header for MTRWaterHeaterManagementClusterBoostParams */
// The class instance for the [MTRWaterHeaterManagementClusterBoostParams] class.
var (
	MTRWaterHeaterManagementClusterBoostParamsClass     _MTRWaterHeaterManagementClusterBoostParamsClass
	MTRWaterHeaterManagementClusterBoostParamsClassOnce sync.Once
)

func getMTRWaterHeaterManagementClusterBoostParamsClass() _MTRWaterHeaterManagementClusterBoostParamsClass {
	MTRWaterHeaterManagementClusterBoostParamsClassOnce.Do(func() {
		MTRWaterHeaterManagementClusterBoostParamsClass = _MTRWaterHeaterManagementClusterBoostParamsClass{objc.GetClass("MTRWaterHeaterManagementClusterBoostParams")}
	})
	return MTRWaterHeaterManagementClusterBoostParamsClass
}

type _MTRWaterHeaterManagementClusterBoostParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWaterHeaterManagementClusterBoostParams */
// An interface definition for the [MTRWaterHeaterManagementClusterBoostParams] class.
type IMTRWaterHeaterManagementClusterBoostParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWaterHeaterManagementClusterBoostParams */
	// properties:
	BoostInfo() IMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct
	SetBoostInfo(value IMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWaterHeaterManagementClusterBoostParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWaterHeaterManagementClusterBoostParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterManagementClusterBoostParamsClass) Alloc() MTRWaterHeaterManagementClusterBoostParams {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWaterHeaterManagementClusterBoostParamsClass) New() MTRWaterHeaterManagementClusterBoostParams {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterManagementClusterBoostParams) Init() MTRWaterHeaterManagementClusterBoostParams {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterManagementClusterBoostParams) Autorelease() MTRWaterHeaterManagementClusterBoostParams {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterManagementClusterBoostParams creates a new MTRWaterHeaterManagementClusterBoostParams instance.
func NewMTRWaterHeaterManagementClusterBoostParams() MTRWaterHeaterManagementClusterBoostParams {
	return getMTRWaterHeaterManagementClusterBoostParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWaterHeaterManagementClusterBoostParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostParams
type MTRWaterHeaterManagementClusterBoostParams struct {
	objectivec.Object
}

// MTRWaterHeaterManagementClusterBoostParamsFrom constructs a [MTRWaterHeaterManagementClusterBoostParams] from an unsafe.Pointer.
func MTRWaterHeaterManagementClusterBoostParamsFrom(ptr unsafe.Pointer) MTRWaterHeaterManagementClusterBoostParams {
	return MTRWaterHeaterManagementClusterBoostParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWaterHeaterManagementClusterBoostParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWaterHeaterManagementClusterBoostParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWaterHeaterManagementClusterBoostParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWaterHeaterManagementClusterBoostParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWaterHeaterManagementClusterBoostParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostParams/boostInfo
func (m_ MTRWaterHeaterManagementClusterBoostParams) BoostInfo() IMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct {
	rv := objc.Send[MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct](m_.ID, objc.Sel("boostInfo"))
	return rv
}/* debug [instance_properties/getter]: boostInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostParams/boostInfo
func (m_ MTRWaterHeaterManagementClusterBoostParams) SetBoostInfo(value IMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBoostInfo:"), value)
}/* debug [instance_properties/setter]: boostInfo */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermanagementclusterboostparams/serversideprocessingtimeout
func (m_ MTRWaterHeaterManagementClusterBoostParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermanagementclusterboostparams/serversideprocessingtimeout
func (m_ MTRWaterHeaterManagementClusterBoostParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermanagementclusterboostparams/timedinvoketimeoutms
func (m_ MTRWaterHeaterManagementClusterBoostParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermanagementclusterboostparams/timedinvoketimeoutms
func (m_ MTRWaterHeaterManagementClusterBoostParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWaterHeaterManagementClusterBoostParams */




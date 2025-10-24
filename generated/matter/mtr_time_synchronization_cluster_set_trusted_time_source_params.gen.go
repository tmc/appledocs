// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRTimeSynchronizationClusterSetTrustedTimeSourceParams */


/* debug [class_header]: Header for MTRTimeSynchronizationClusterSetTrustedTimeSourceParams */
// The class instance for the [MTRTimeSynchronizationClusterSetTrustedTimeSourceParams] class.
var (
	MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass     _MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass
	MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass() _MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass {
	MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClassOnce.Do(func() {
		MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass = _MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass{objc.GetClass("MTRTimeSynchronizationClusterSetTrustedTimeSourceParams")}
	})
	return MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass
}

type _MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTimeSynchronizationClusterSetTrustedTimeSourceParams */
// An interface definition for the [MTRTimeSynchronizationClusterSetTrustedTimeSourceParams] class.
type IMTRTimeSynchronizationClusterSetTrustedTimeSourceParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRTimeSynchronizationClusterSetTrustedTimeSourceParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TrustedTimeSource() IMTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct
	SetTrustedTimeSource(value IMTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTimeSynchronizationClusterSetTrustedTimeSourceParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTimeSynchronizationClusterSetTrustedTimeSourceParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass) Alloc() MTRTimeSynchronizationClusterSetTrustedTimeSourceParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTrustedTimeSourceParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass) New() MTRTimeSynchronizationClusterSetTrustedTimeSourceParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTrustedTimeSourceParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterSetTrustedTimeSourceParams) Init() MTRTimeSynchronizationClusterSetTrustedTimeSourceParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTrustedTimeSourceParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterSetTrustedTimeSourceParams) Autorelease() MTRTimeSynchronizationClusterSetTrustedTimeSourceParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTrustedTimeSourceParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterSetTrustedTimeSourceParams creates a new MTRTimeSynchronizationClusterSetTrustedTimeSourceParams instance.
func NewMTRTimeSynchronizationClusterSetTrustedTimeSourceParams() MTRTimeSynchronizationClusterSetTrustedTimeSourceParams {
	return getMTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTimeSynchronizationClusterSetTrustedTimeSourceParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTrustedTimeSourceParams
type MTRTimeSynchronizationClusterSetTrustedTimeSourceParams struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsFrom constructs a [MTRTimeSynchronizationClusterSetTrustedTimeSourceParams] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterSetTrustedTimeSourceParams {
	return MTRTimeSynchronizationClusterSetTrustedTimeSourceParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTimeSynchronizationClusterSetTrustedTimeSourceParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTimeSynchronizationClusterSetTrustedTimeSourceParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTimeSynchronizationClusterSetTrustedTimeSourceParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTimeSynchronizationClusterSetTrustedTimeSourceParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTimeSynchronizationClusterSetTrustedTimeSourceParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTrustedTimeSourceParams/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetTrustedTimeSourceParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTrustedTimeSourceParams/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetTrustedTimeSourceParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersettrustedtimesourceparams/timedinvoketimeoutms
func (m_ MTRTimeSynchronizationClusterSetTrustedTimeSourceParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersettrustedtimesourceparams/timedinvoketimeoutms
func (m_ MTRTimeSynchronizationClusterSetTrustedTimeSourceParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersettrustedtimesourceparams/trustedtimesource
func (m_ MTRTimeSynchronizationClusterSetTrustedTimeSourceParams) TrustedTimeSource() IMTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct](m_.ID, objc.Sel("trustedTimeSource"))
	return rv
}/* debug [instance_properties/getter]: trustedTimeSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersettrustedtimesourceparams/trustedtimesource
func (m_ MTRTimeSynchronizationClusterSetTrustedTimeSourceParams) SetTrustedTimeSource(value IMTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrustedTimeSource:"), value)
}/* debug [instance_properties/setter]: trustedTimeSource */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTimeSynchronizationClusterSetTrustedTimeSourceParams */




// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams */


/* debug [class_header]: Header for MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams */
// The class instance for the [MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams] class.
var (
	MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass     _MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass
	MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClassOnce sync.Once
)

func getMTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass() _MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass {
	MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClassOnce.Do(func() {
		MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass = _MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass{objc.GetClass("MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams")}
	})
	return MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass
}

type _MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams */
// An interface definition for the [MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams] class.
type IMTRThreadNetworkDirectoryClusterGetOperationalDatasetParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams */
	// properties:
	ExtendedPanID() objc.IObject /* cross-framework: NSData */
	SetExtendedPanID(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass) Alloc() MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass) New() MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams) Init() MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams) Autorelease() MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDirectoryClusterGetOperationalDatasetParams creates a new MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams instance.
func NewMTRThreadNetworkDirectoryClusterGetOperationalDatasetParams() MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams {
	return getMTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams
type MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams struct {
	objectivec.Object
}

// MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsFrom constructs a [MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams] from an unsafe.Pointer.
func MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsFrom(ptr unsafe.Pointer) MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams {
	return MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams/extendedPanID
func (m_ MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams) ExtendedPanID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("extendedPanID"))
	return rv
}/* debug [instance_properties/getter]: extendedPanID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams/extendedPanID
func (m_ MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams) SetExtendedPanID(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedPanID:"), value)
}/* debug [instance_properties/setter]: extendedPanID */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclustergetoperationaldatasetparams/serversideprocessingtimeout
func (m_ MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclustergetoperationaldatasetparams/serversideprocessingtimeout
func (m_ MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclustergetoperationaldatasetparams/timedinvoketimeoutms
func (m_ MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclustergetoperationaldatasetparams/timedinvoketimeoutms
func (m_ MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams */




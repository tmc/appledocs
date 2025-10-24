// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadNetworkDirectoryClusterAddNetworkParams */


/* debug [class_header]: Header for MTRThreadNetworkDirectoryClusterAddNetworkParams */
// The class instance for the [MTRThreadNetworkDirectoryClusterAddNetworkParams] class.
var (
	MTRThreadNetworkDirectoryClusterAddNetworkParamsClass     _MTRThreadNetworkDirectoryClusterAddNetworkParamsClass
	MTRThreadNetworkDirectoryClusterAddNetworkParamsClassOnce sync.Once
)

func getMTRThreadNetworkDirectoryClusterAddNetworkParamsClass() _MTRThreadNetworkDirectoryClusterAddNetworkParamsClass {
	MTRThreadNetworkDirectoryClusterAddNetworkParamsClassOnce.Do(func() {
		MTRThreadNetworkDirectoryClusterAddNetworkParamsClass = _MTRThreadNetworkDirectoryClusterAddNetworkParamsClass{objc.GetClass("MTRThreadNetworkDirectoryClusterAddNetworkParams")}
	})
	return MTRThreadNetworkDirectoryClusterAddNetworkParamsClass
}

type _MTRThreadNetworkDirectoryClusterAddNetworkParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadNetworkDirectoryClusterAddNetworkParams */
// An interface definition for the [MTRThreadNetworkDirectoryClusterAddNetworkParams] class.
type IMTRThreadNetworkDirectoryClusterAddNetworkParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadNetworkDirectoryClusterAddNetworkParams */
	// properties:
	OperationalDataset() objc.IObject /* cross-framework: NSData */
	SetOperationalDataset(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadNetworkDirectoryClusterAddNetworkParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadNetworkDirectoryClusterAddNetworkParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDirectoryClusterAddNetworkParamsClass) Alloc() MTRThreadNetworkDirectoryClusterAddNetworkParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterAddNetworkParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThreadNetworkDirectoryClusterAddNetworkParamsClass) New() MTRThreadNetworkDirectoryClusterAddNetworkParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterAddNetworkParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDirectoryClusterAddNetworkParams) Init() MTRThreadNetworkDirectoryClusterAddNetworkParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterAddNetworkParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDirectoryClusterAddNetworkParams) Autorelease() MTRThreadNetworkDirectoryClusterAddNetworkParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterAddNetworkParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDirectoryClusterAddNetworkParams creates a new MTRThreadNetworkDirectoryClusterAddNetworkParams instance.
func NewMTRThreadNetworkDirectoryClusterAddNetworkParams() MTRThreadNetworkDirectoryClusterAddNetworkParams {
	return getMTRThreadNetworkDirectoryClusterAddNetworkParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadNetworkDirectoryClusterAddNetworkParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterAddNetworkParams
type MTRThreadNetworkDirectoryClusterAddNetworkParams struct {
	objectivec.Object
}

// MTRThreadNetworkDirectoryClusterAddNetworkParamsFrom constructs a [MTRThreadNetworkDirectoryClusterAddNetworkParams] from an unsafe.Pointer.
func MTRThreadNetworkDirectoryClusterAddNetworkParamsFrom(ptr unsafe.Pointer) MTRThreadNetworkDirectoryClusterAddNetworkParams {
	return MTRThreadNetworkDirectoryClusterAddNetworkParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadNetworkDirectoryClusterAddNetworkParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadNetworkDirectoryClusterAddNetworkParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadNetworkDirectoryClusterAddNetworkParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadNetworkDirectoryClusterAddNetworkParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadNetworkDirectoryClusterAddNetworkParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterAddNetworkParams/operationalDataset
func (m_ MTRThreadNetworkDirectoryClusterAddNetworkParams) OperationalDataset() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("operationalDataset"))
	return rv
}/* debug [instance_properties/getter]: operationalDataset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterAddNetworkParams/operationalDataset
func (m_ MTRThreadNetworkDirectoryClusterAddNetworkParams) SetOperationalDataset(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalDataset:"), value)
}/* debug [instance_properties/setter]: operationalDataset */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclusteraddnetworkparams/serversideprocessingtimeout
func (m_ MTRThreadNetworkDirectoryClusterAddNetworkParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclusteraddnetworkparams/serversideprocessingtimeout
func (m_ MTRThreadNetworkDirectoryClusterAddNetworkParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclusteraddnetworkparams/timedinvoketimeoutms
func (m_ MTRThreadNetworkDirectoryClusterAddNetworkParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclusteraddnetworkparams/timedinvoketimeoutms
func (m_ MTRThreadNetworkDirectoryClusterAddNetworkParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadNetworkDirectoryClusterAddNetworkParams */




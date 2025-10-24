// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadNetworkDirectoryClusterRemoveNetworkParams */


/* debug [class_header]: Header for MTRThreadNetworkDirectoryClusterRemoveNetworkParams */
// The class instance for the [MTRThreadNetworkDirectoryClusterRemoveNetworkParams] class.
var (
	MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass     _MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass
	MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClassOnce sync.Once
)

func getMTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass() _MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass {
	MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClassOnce.Do(func() {
		MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass = _MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass{objc.GetClass("MTRThreadNetworkDirectoryClusterRemoveNetworkParams")}
	})
	return MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass
}

type _MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadNetworkDirectoryClusterRemoveNetworkParams */
// An interface definition for the [MTRThreadNetworkDirectoryClusterRemoveNetworkParams] class.
type IMTRThreadNetworkDirectoryClusterRemoveNetworkParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadNetworkDirectoryClusterRemoveNetworkParams */
	// properties:
	ExtendedPanID() objc.IObject /* cross-framework: NSData */
	SetExtendedPanID(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadNetworkDirectoryClusterRemoveNetworkParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadNetworkDirectoryClusterRemoveNetworkParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass) Alloc() MTRThreadNetworkDirectoryClusterRemoveNetworkParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterRemoveNetworkParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass) New() MTRThreadNetworkDirectoryClusterRemoveNetworkParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterRemoveNetworkParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDirectoryClusterRemoveNetworkParams) Init() MTRThreadNetworkDirectoryClusterRemoveNetworkParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterRemoveNetworkParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDirectoryClusterRemoveNetworkParams) Autorelease() MTRThreadNetworkDirectoryClusterRemoveNetworkParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterRemoveNetworkParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDirectoryClusterRemoveNetworkParams creates a new MTRThreadNetworkDirectoryClusterRemoveNetworkParams instance.
func NewMTRThreadNetworkDirectoryClusterRemoveNetworkParams() MTRThreadNetworkDirectoryClusterRemoveNetworkParams {
	return getMTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadNetworkDirectoryClusterRemoveNetworkParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterRemoveNetworkParams
type MTRThreadNetworkDirectoryClusterRemoveNetworkParams struct {
	objectivec.Object
}

// MTRThreadNetworkDirectoryClusterRemoveNetworkParamsFrom constructs a [MTRThreadNetworkDirectoryClusterRemoveNetworkParams] from an unsafe.Pointer.
func MTRThreadNetworkDirectoryClusterRemoveNetworkParamsFrom(ptr unsafe.Pointer) MTRThreadNetworkDirectoryClusterRemoveNetworkParams {
	return MTRThreadNetworkDirectoryClusterRemoveNetworkParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadNetworkDirectoryClusterRemoveNetworkParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadNetworkDirectoryClusterRemoveNetworkParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadNetworkDirectoryClusterRemoveNetworkParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadNetworkDirectoryClusterRemoveNetworkParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadNetworkDirectoryClusterRemoveNetworkParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterRemoveNetworkParams/extendedPanID
func (m_ MTRThreadNetworkDirectoryClusterRemoveNetworkParams) ExtendedPanID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("extendedPanID"))
	return rv
}/* debug [instance_properties/getter]: extendedPanID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterRemoveNetworkParams/extendedPanID
func (m_ MTRThreadNetworkDirectoryClusterRemoveNetworkParams) SetExtendedPanID(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedPanID:"), value)
}/* debug [instance_properties/setter]: extendedPanID */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclusterremovenetworkparams/serversideprocessingtimeout
func (m_ MTRThreadNetworkDirectoryClusterRemoveNetworkParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclusterremovenetworkparams/serversideprocessingtimeout
func (m_ MTRThreadNetworkDirectoryClusterRemoveNetworkParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclusterremovenetworkparams/timedinvoketimeoutms
func (m_ MTRThreadNetworkDirectoryClusterRemoveNetworkParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclusterremovenetworkparams/timedinvoketimeoutms
func (m_ MTRThreadNetworkDirectoryClusterRemoveNetworkParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadNetworkDirectoryClusterRemoveNetworkParams */




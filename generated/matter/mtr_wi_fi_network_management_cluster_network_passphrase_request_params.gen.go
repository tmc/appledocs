// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams */


/* debug [class_header]: Header for MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams */
// The class instance for the [MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams] class.
var (
	MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass     _MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass
	MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClassOnce sync.Once
)

func getMTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass() _MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass {
	MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClassOnce.Do(func() {
		MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass = _MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass{objc.GetClass("MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams")}
	})
	return MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass
}

type _MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams */
// An interface definition for the [MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams] class.
type IMTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass) Alloc() MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams {
	rv := objc.Send[MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass) New() MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams {
	rv := objc.Send[MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams) Init() MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams {
	rv := objc.Send[MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams) Autorelease() MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams {
	rv := objc.Send[MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams creates a new MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams instance.
func NewMTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams() MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams {
	return getMTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams
type MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams struct {
	objectivec.Object
}

// MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsFrom constructs a [MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams] from an unsafe.Pointer.
func MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsFrom(ptr unsafe.Pointer) MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams {
	return MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams/serverSideProcessingTimeout
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams/serverSideProcessingTimeout
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams/timedInvokeTimeoutMs
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams/timedInvokeTimeoutMs
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams */




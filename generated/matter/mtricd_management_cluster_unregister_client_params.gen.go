// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRICDManagementClusterUnregisterClientParams */


/* debug [class_header]: Header for MTRICDManagementClusterUnregisterClientParams */
// The class instance for the [MTRICDManagementClusterUnregisterClientParams] class.
var (
	MTRICDManagementClusterUnregisterClientParamsClass     _MTRICDManagementClusterUnregisterClientParamsClass
	MTRICDManagementClusterUnregisterClientParamsClassOnce sync.Once
)

func getMTRICDManagementClusterUnregisterClientParamsClass() _MTRICDManagementClusterUnregisterClientParamsClass {
	MTRICDManagementClusterUnregisterClientParamsClassOnce.Do(func() {
		MTRICDManagementClusterUnregisterClientParamsClass = _MTRICDManagementClusterUnregisterClientParamsClass{objc.GetClass("MTRICDManagementClusterUnregisterClientParams")}
	})
	return MTRICDManagementClusterUnregisterClientParamsClass
}

type _MTRICDManagementClusterUnregisterClientParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRICDManagementClusterUnregisterClientParams */
// An interface definition for the [MTRICDManagementClusterUnregisterClientParams] class.
type IMTRICDManagementClusterUnregisterClientParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRICDManagementClusterUnregisterClientParams */
	// properties:
	CheckInNodeID() objc.IObject /* cross-framework: NSNumber */
	SetCheckInNodeID(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	VerificationKey() foundation.Data
	SetVerificationKey(value foundation.Data)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRICDManagementClusterUnregisterClientParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRICDManagementClusterUnregisterClientParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRICDManagementClusterUnregisterClientParamsClass) Alloc() MTRICDManagementClusterUnregisterClientParams {
	rv := objc.Send[MTRICDManagementClusterUnregisterClientParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRICDManagementClusterUnregisterClientParamsClass) New() MTRICDManagementClusterUnregisterClientParams {
	rv := objc.Send[MTRICDManagementClusterUnregisterClientParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRICDManagementClusterUnregisterClientParams) Init() MTRICDManagementClusterUnregisterClientParams {
	rv := objc.Send[MTRICDManagementClusterUnregisterClientParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRICDManagementClusterUnregisterClientParams) Autorelease() MTRICDManagementClusterUnregisterClientParams {
	rv := objc.Send[MTRICDManagementClusterUnregisterClientParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRICDManagementClusterUnregisterClientParams creates a new MTRICDManagementClusterUnregisterClientParams instance.
func NewMTRICDManagementClusterUnregisterClientParams() MTRICDManagementClusterUnregisterClientParams {
	return getMTRICDManagementClusterUnregisterClientParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRICDManagementClusterUnregisterClientParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterUnregisterClientParams
type MTRICDManagementClusterUnregisterClientParams struct {
	objectivec.Object
}

// MTRICDManagementClusterUnregisterClientParamsFrom constructs a [MTRICDManagementClusterUnregisterClientParams] from an unsafe.Pointer.
func MTRICDManagementClusterUnregisterClientParamsFrom(ptr unsafe.Pointer) MTRICDManagementClusterUnregisterClientParams {
	return MTRICDManagementClusterUnregisterClientParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRICDManagementClusterUnregisterClientParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRICDManagementClusterUnregisterClientParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRICDManagementClusterUnregisterClientParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRICDManagementClusterUnregisterClientParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRICDManagementClusterUnregisterClientParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterUnregisterClientParams/checkInNodeID
func (m_ MTRICDManagementClusterUnregisterClientParams) CheckInNodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("checkInNodeID"))
	return rv
}/* debug [instance_properties/getter]: checkInNodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterUnregisterClientParams/checkInNodeID
func (m_ MTRICDManagementClusterUnregisterClientParams) SetCheckInNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCheckInNodeID:"), value)
}/* debug [instance_properties/setter]: checkInNodeID */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterunregisterclientparams/serversideprocessingtimeout
func (m_ MTRICDManagementClusterUnregisterClientParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterunregisterclientparams/serversideprocessingtimeout
func (m_ MTRICDManagementClusterUnregisterClientParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterunregisterclientparams/timedinvoketimeoutms
func (m_ MTRICDManagementClusterUnregisterClientParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterunregisterclientparams/timedinvoketimeoutms
func (m_ MTRICDManagementClusterUnregisterClientParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterunregisterclientparams/verificationkey
func (m_ MTRICDManagementClusterUnregisterClientParams) VerificationKey() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("verificationKey"))
	return rv
}/* debug [instance_properties/getter]: verificationKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterunregisterclientparams/verificationkey
func (m_ MTRICDManagementClusterUnregisterClientParams) SetVerificationKey(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVerificationKey:"), value)
}/* debug [instance_properties/setter]: verificationKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRICDManagementClusterUnregisterClientParams */




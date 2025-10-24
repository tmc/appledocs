// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRICDManagementClusterRegisterClientParams */


/* debug [class_header]: Header for MTRICDManagementClusterRegisterClientParams */
// The class instance for the [MTRICDManagementClusterRegisterClientParams] class.
var (
	MTRICDManagementClusterRegisterClientParamsClass     _MTRICDManagementClusterRegisterClientParamsClass
	MTRICDManagementClusterRegisterClientParamsClassOnce sync.Once
)

func getMTRICDManagementClusterRegisterClientParamsClass() _MTRICDManagementClusterRegisterClientParamsClass {
	MTRICDManagementClusterRegisterClientParamsClassOnce.Do(func() {
		MTRICDManagementClusterRegisterClientParamsClass = _MTRICDManagementClusterRegisterClientParamsClass{objc.GetClass("MTRICDManagementClusterRegisterClientParams")}
	})
	return MTRICDManagementClusterRegisterClientParamsClass
}

type _MTRICDManagementClusterRegisterClientParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRICDManagementClusterRegisterClientParams */
// An interface definition for the [MTRICDManagementClusterRegisterClientParams] class.
type IMTRICDManagementClusterRegisterClientParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRICDManagementClusterRegisterClientParams */
	// properties:
	CheckInNodeID() objc.IObject /* cross-framework: NSNumber */
	SetCheckInNodeID(value objc.IObject /* cross-framework: NSNumber */)
	ClientType() objc.IObject /* cross-framework: NSNumber */
	SetClientType(value objc.IObject /* cross-framework: NSNumber */)
	Key() foundation.Data
	SetKey(value foundation.Data)
	MonitoredSubject() objc.IObject /* cross-framework: NSNumber */
	SetMonitoredSubject(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	VerificationKey() foundation.Data
	SetVerificationKey(value foundation.Data)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRICDManagementClusterRegisterClientParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRICDManagementClusterRegisterClientParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRICDManagementClusterRegisterClientParamsClass) Alloc() MTRICDManagementClusterRegisterClientParams {
	rv := objc.Send[MTRICDManagementClusterRegisterClientParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRICDManagementClusterRegisterClientParamsClass) New() MTRICDManagementClusterRegisterClientParams {
	rv := objc.Send[MTRICDManagementClusterRegisterClientParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRICDManagementClusterRegisterClientParams) Init() MTRICDManagementClusterRegisterClientParams {
	rv := objc.Send[MTRICDManagementClusterRegisterClientParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRICDManagementClusterRegisterClientParams) Autorelease() MTRICDManagementClusterRegisterClientParams {
	rv := objc.Send[MTRICDManagementClusterRegisterClientParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRICDManagementClusterRegisterClientParams creates a new MTRICDManagementClusterRegisterClientParams instance.
func NewMTRICDManagementClusterRegisterClientParams() MTRICDManagementClusterRegisterClientParams {
	return getMTRICDManagementClusterRegisterClientParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRICDManagementClusterRegisterClientParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams
type MTRICDManagementClusterRegisterClientParams struct {
	objectivec.Object
}

// MTRICDManagementClusterRegisterClientParamsFrom constructs a [MTRICDManagementClusterRegisterClientParams] from an unsafe.Pointer.
func MTRICDManagementClusterRegisterClientParamsFrom(ptr unsafe.Pointer) MTRICDManagementClusterRegisterClientParams {
	return MTRICDManagementClusterRegisterClientParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRICDManagementClusterRegisterClientParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRICDManagementClusterRegisterClientParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRICDManagementClusterRegisterClientParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRICDManagementClusterRegisterClientParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRICDManagementClusterRegisterClientParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/checkInNodeID
func (m_ MTRICDManagementClusterRegisterClientParams) CheckInNodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("checkInNodeID"))
	return rv
}/* debug [instance_properties/getter]: checkInNodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/checkInNodeID
func (m_ MTRICDManagementClusterRegisterClientParams) SetCheckInNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCheckInNodeID:"), value)
}/* debug [instance_properties/setter]: checkInNodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterregisterclientparams/clienttype
func (m_ MTRICDManagementClusterRegisterClientParams) ClientType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("clientType"))
	return rv
}/* debug [instance_properties/getter]: clientType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterregisterclientparams/clienttype
func (m_ MTRICDManagementClusterRegisterClientParams) SetClientType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setClientType:"), value)
}/* debug [instance_properties/setter]: clientType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterregisterclientparams/key
func (m_ MTRICDManagementClusterRegisterClientParams) Key() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("key"))
	return rv
}/* debug [instance_properties/getter]: key */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterregisterclientparams/key
func (m_ MTRICDManagementClusterRegisterClientParams) SetKey(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKey:"), value)
}/* debug [instance_properties/setter]: key */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterregisterclientparams/monitoredsubject
func (m_ MTRICDManagementClusterRegisterClientParams) MonitoredSubject() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("monitoredSubject"))
	return rv
}/* debug [instance_properties/getter]: monitoredSubject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterregisterclientparams/monitoredsubject
func (m_ MTRICDManagementClusterRegisterClientParams) SetMonitoredSubject(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMonitoredSubject:"), value)
}/* debug [instance_properties/setter]: monitoredSubject */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterregisterclientparams/serversideprocessingtimeout
func (m_ MTRICDManagementClusterRegisterClientParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterregisterclientparams/serversideprocessingtimeout
func (m_ MTRICDManagementClusterRegisterClientParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterregisterclientparams/timedinvoketimeoutms
func (m_ MTRICDManagementClusterRegisterClientParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterregisterclientparams/timedinvoketimeoutms
func (m_ MTRICDManagementClusterRegisterClientParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterregisterclientparams/verificationkey
func (m_ MTRICDManagementClusterRegisterClientParams) VerificationKey() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("verificationKey"))
	return rv
}/* debug [instance_properties/getter]: verificationKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterregisterclientparams/verificationkey
func (m_ MTRICDManagementClusterRegisterClientParams) SetVerificationKey(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVerificationKey:"), value)
}/* debug [instance_properties/setter]: verificationKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRICDManagementClusterRegisterClientParams */




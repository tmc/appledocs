// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterSetAliroReaderConfigParams */


/* debug [class_header]: Header for MTRDoorLockClusterSetAliroReaderConfigParams */
// The class instance for the [MTRDoorLockClusterSetAliroReaderConfigParams] class.
var (
	MTRDoorLockClusterSetAliroReaderConfigParamsClass     _MTRDoorLockClusterSetAliroReaderConfigParamsClass
	MTRDoorLockClusterSetAliroReaderConfigParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetAliroReaderConfigParamsClass() _MTRDoorLockClusterSetAliroReaderConfigParamsClass {
	MTRDoorLockClusterSetAliroReaderConfigParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetAliroReaderConfigParamsClass = _MTRDoorLockClusterSetAliroReaderConfigParamsClass{objc.GetClass("MTRDoorLockClusterSetAliroReaderConfigParams")}
	})
	return MTRDoorLockClusterSetAliroReaderConfigParamsClass
}

type _MTRDoorLockClusterSetAliroReaderConfigParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterSetAliroReaderConfigParams */
// An interface definition for the [MTRDoorLockClusterSetAliroReaderConfigParams] class.
type IMTRDoorLockClusterSetAliroReaderConfigParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterSetAliroReaderConfigParams */
	// properties:
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	GroupIdentifier() foundation.Data
	SetGroupIdentifier(value foundation.Data)
	GroupResolvingKey() foundation.Data
	SetGroupResolvingKey(value foundation.Data)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	SigningKey() foundation.Data
	SetSigningKey(value foundation.Data)
	VerificationKey() foundation.Data
	SetVerificationKey(value foundation.Data)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterSetAliroReaderConfigParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterSetAliroReaderConfigParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetAliroReaderConfigParamsClass) Alloc() MTRDoorLockClusterSetAliroReaderConfigParams {
	rv := objc.Send[MTRDoorLockClusterSetAliroReaderConfigParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterSetAliroReaderConfigParamsClass) New() MTRDoorLockClusterSetAliroReaderConfigParams {
	rv := objc.Send[MTRDoorLockClusterSetAliroReaderConfigParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) Init() MTRDoorLockClusterSetAliroReaderConfigParams {
	rv := objc.Send[MTRDoorLockClusterSetAliroReaderConfigParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) Autorelease() MTRDoorLockClusterSetAliroReaderConfigParams {
	rv := objc.Send[MTRDoorLockClusterSetAliroReaderConfigParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetAliroReaderConfigParams creates a new MTRDoorLockClusterSetAliroReaderConfigParams instance.
func NewMTRDoorLockClusterSetAliroReaderConfigParams() MTRDoorLockClusterSetAliroReaderConfigParams {
	return getMTRDoorLockClusterSetAliroReaderConfigParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterSetAliroReaderConfigParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams
type MTRDoorLockClusterSetAliroReaderConfigParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetAliroReaderConfigParamsFrom constructs a [MTRDoorLockClusterSetAliroReaderConfigParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetAliroReaderConfigParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetAliroReaderConfigParams {
	return MTRDoorLockClusterSetAliroReaderConfigParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterSetAliroReaderConfigParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterSetAliroReaderConfigParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterSetAliroReaderConfigParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterSetAliroReaderConfigParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterSetAliroReaderConfigParams */

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetaliroreaderconfigparams/groupidentifier
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) GroupIdentifier() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("groupIdentifier"))
	return rv
}/* debug [instance_properties/getter]: groupIdentifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetaliroreaderconfigparams/groupidentifier
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetGroupIdentifier(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupIdentifier:"), value)
}/* debug [instance_properties/setter]: groupIdentifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetaliroreaderconfigparams/groupresolvingkey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) GroupResolvingKey() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("groupResolvingKey"))
	return rv
}/* debug [instance_properties/getter]: groupResolvingKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetaliroreaderconfigparams/groupresolvingkey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetGroupResolvingKey(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupResolvingKey:"), value)
}/* debug [instance_properties/setter]: groupResolvingKey */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetaliroreaderconfigparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetaliroreaderconfigparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetaliroreaderconfigparams/signingkey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SigningKey() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("signingKey"))
	return rv
}/* debug [instance_properties/getter]: signingKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetaliroreaderconfigparams/signingkey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetSigningKey(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSigningKey:"), value)
}/* debug [instance_properties/setter]: signingKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetaliroreaderconfigparams/verificationkey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) VerificationKey() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("verificationKey"))
	return rv
}/* debug [instance_properties/getter]: verificationKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetaliroreaderconfigparams/verificationkey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetVerificationKey(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVerificationKey:"), value)
}/* debug [instance_properties/setter]: verificationKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterSetAliroReaderConfigParams */




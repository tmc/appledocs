// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterUnboltDoorParams */


/* debug [class_header]: Header for MTRDoorLockClusterUnboltDoorParams */
// The class instance for the [MTRDoorLockClusterUnboltDoorParams] class.
var (
	MTRDoorLockClusterUnboltDoorParamsClass     _MTRDoorLockClusterUnboltDoorParamsClass
	MTRDoorLockClusterUnboltDoorParamsClassOnce sync.Once
)

func getMTRDoorLockClusterUnboltDoorParamsClass() _MTRDoorLockClusterUnboltDoorParamsClass {
	MTRDoorLockClusterUnboltDoorParamsClassOnce.Do(func() {
		MTRDoorLockClusterUnboltDoorParamsClass = _MTRDoorLockClusterUnboltDoorParamsClass{objc.GetClass("MTRDoorLockClusterUnboltDoorParams")}
	})
	return MTRDoorLockClusterUnboltDoorParamsClass
}

type _MTRDoorLockClusterUnboltDoorParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterUnboltDoorParams */
// An interface definition for the [MTRDoorLockClusterUnboltDoorParams] class.
type IMTRDoorLockClusterUnboltDoorParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterUnboltDoorParams */
	// properties:
	PinCode() objc.IObject /* cross-framework: NSData */
	SetPinCode(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterUnboltDoorParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterUnboltDoorParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterUnboltDoorParamsClass) Alloc() MTRDoorLockClusterUnboltDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnboltDoorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterUnboltDoorParamsClass) New() MTRDoorLockClusterUnboltDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnboltDoorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterUnboltDoorParams) Init() MTRDoorLockClusterUnboltDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnboltDoorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterUnboltDoorParams) Autorelease() MTRDoorLockClusterUnboltDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnboltDoorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterUnboltDoorParams creates a new MTRDoorLockClusterUnboltDoorParams instance.
func NewMTRDoorLockClusterUnboltDoorParams() MTRDoorLockClusterUnboltDoorParams {
	return getMTRDoorLockClusterUnboltDoorParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterUnboltDoorParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnboltDoorParams
type MTRDoorLockClusterUnboltDoorParams struct {
	objectivec.Object
}

// MTRDoorLockClusterUnboltDoorParamsFrom constructs a [MTRDoorLockClusterUnboltDoorParams] from an unsafe.Pointer.
func MTRDoorLockClusterUnboltDoorParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterUnboltDoorParams {
	return MTRDoorLockClusterUnboltDoorParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterUnboltDoorParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterUnboltDoorParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterUnboltDoorParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterUnboltDoorParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterUnboltDoorParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnboltDoorParams/pinCode
func (m_ MTRDoorLockClusterUnboltDoorParams) PinCode() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("pinCode"))
	return rv
}/* debug [instance_properties/getter]: pinCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnboltDoorParams/pinCode
func (m_ MTRDoorLockClusterUnboltDoorParams) SetPinCode(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPinCode:"), value)
}/* debug [instance_properties/setter]: pinCode */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnboltDoorParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterUnboltDoorParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnboltDoorParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterUnboltDoorParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterunboltdoorparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterUnboltDoorParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterunboltdoorparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterUnboltDoorParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterUnboltDoorParams */




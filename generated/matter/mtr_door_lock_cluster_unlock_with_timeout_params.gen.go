// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterUnlockWithTimeoutParams */


/* debug [class_header]: Header for MTRDoorLockClusterUnlockWithTimeoutParams */
// The class instance for the [MTRDoorLockClusterUnlockWithTimeoutParams] class.
var (
	MTRDoorLockClusterUnlockWithTimeoutParamsClass     _MTRDoorLockClusterUnlockWithTimeoutParamsClass
	MTRDoorLockClusterUnlockWithTimeoutParamsClassOnce sync.Once
)

func getMTRDoorLockClusterUnlockWithTimeoutParamsClass() _MTRDoorLockClusterUnlockWithTimeoutParamsClass {
	MTRDoorLockClusterUnlockWithTimeoutParamsClassOnce.Do(func() {
		MTRDoorLockClusterUnlockWithTimeoutParamsClass = _MTRDoorLockClusterUnlockWithTimeoutParamsClass{objc.GetClass("MTRDoorLockClusterUnlockWithTimeoutParams")}
	})
	return MTRDoorLockClusterUnlockWithTimeoutParamsClass
}

type _MTRDoorLockClusterUnlockWithTimeoutParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterUnlockWithTimeoutParams */
// An interface definition for the [MTRDoorLockClusterUnlockWithTimeoutParams] class.
type IMTRDoorLockClusterUnlockWithTimeoutParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterUnlockWithTimeoutParams */
	// properties:
	PinCode() objc.IObject /* cross-framework: NSData */
	SetPinCode(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	Timeout() objc.IObject /* cross-framework: NSNumber */
	SetTimeout(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterUnlockWithTimeoutParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterUnlockWithTimeoutParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterUnlockWithTimeoutParamsClass) Alloc() MTRDoorLockClusterUnlockWithTimeoutParams {
	rv := objc.Send[MTRDoorLockClusterUnlockWithTimeoutParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterUnlockWithTimeoutParamsClass) New() MTRDoorLockClusterUnlockWithTimeoutParams {
	rv := objc.Send[MTRDoorLockClusterUnlockWithTimeoutParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) Init() MTRDoorLockClusterUnlockWithTimeoutParams {
	rv := objc.Send[MTRDoorLockClusterUnlockWithTimeoutParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) Autorelease() MTRDoorLockClusterUnlockWithTimeoutParams {
	rv := objc.Send[MTRDoorLockClusterUnlockWithTimeoutParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterUnlockWithTimeoutParams creates a new MTRDoorLockClusterUnlockWithTimeoutParams instance.
func NewMTRDoorLockClusterUnlockWithTimeoutParams() MTRDoorLockClusterUnlockWithTimeoutParams {
	return getMTRDoorLockClusterUnlockWithTimeoutParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterUnlockWithTimeoutParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockWithTimeoutParams
type MTRDoorLockClusterUnlockWithTimeoutParams struct {
	objectivec.Object
}

// MTRDoorLockClusterUnlockWithTimeoutParamsFrom constructs a [MTRDoorLockClusterUnlockWithTimeoutParams] from an unsafe.Pointer.
func MTRDoorLockClusterUnlockWithTimeoutParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterUnlockWithTimeoutParams {
	return MTRDoorLockClusterUnlockWithTimeoutParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterUnlockWithTimeoutParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterUnlockWithTimeoutParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterUnlockWithTimeoutParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterUnlockWithTimeoutParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterUnlockWithTimeoutParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockWithTimeoutParams/pinCode
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) PinCode() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("pinCode"))
	return rv
}/* debug [instance_properties/getter]: pinCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockWithTimeoutParams/pinCode
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) SetPinCode(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPinCode:"), value)
}/* debug [instance_properties/setter]: pinCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockWithTimeoutParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockWithTimeoutParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockWithTimeoutParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockWithTimeoutParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockWithTimeoutParams/timeout
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) Timeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timeout"))
	return rv
}/* debug [instance_properties/getter]: timeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockWithTimeoutParams/timeout
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) SetTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeout:"), value)
}/* debug [instance_properties/setter]: timeout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterUnlockWithTimeoutParams */




// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterUnlockDoorParams */


/* debug [class_header]: Header for MTRDoorLockClusterUnlockDoorParams */
// The class instance for the [MTRDoorLockClusterUnlockDoorParams] class.
var (
	MTRDoorLockClusterUnlockDoorParamsClass     _MTRDoorLockClusterUnlockDoorParamsClass
	MTRDoorLockClusterUnlockDoorParamsClassOnce sync.Once
)

func getMTRDoorLockClusterUnlockDoorParamsClass() _MTRDoorLockClusterUnlockDoorParamsClass {
	MTRDoorLockClusterUnlockDoorParamsClassOnce.Do(func() {
		MTRDoorLockClusterUnlockDoorParamsClass = _MTRDoorLockClusterUnlockDoorParamsClass{objc.GetClass("MTRDoorLockClusterUnlockDoorParams")}
	})
	return MTRDoorLockClusterUnlockDoorParamsClass
}

type _MTRDoorLockClusterUnlockDoorParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterUnlockDoorParams */
// An interface definition for the [MTRDoorLockClusterUnlockDoorParams] class.
type IMTRDoorLockClusterUnlockDoorParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterUnlockDoorParams */
	// properties:
	PinCode() objc.IObject /* cross-framework: NSData */
	SetPinCode(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterUnlockDoorParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterUnlockDoorParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterUnlockDoorParamsClass) Alloc() MTRDoorLockClusterUnlockDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnlockDoorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterUnlockDoorParamsClass) New() MTRDoorLockClusterUnlockDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnlockDoorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterUnlockDoorParams) Init() MTRDoorLockClusterUnlockDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnlockDoorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterUnlockDoorParams) Autorelease() MTRDoorLockClusterUnlockDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnlockDoorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterUnlockDoorParams creates a new MTRDoorLockClusterUnlockDoorParams instance.
func NewMTRDoorLockClusterUnlockDoorParams() MTRDoorLockClusterUnlockDoorParams {
	return getMTRDoorLockClusterUnlockDoorParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterUnlockDoorParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockDoorParams
type MTRDoorLockClusterUnlockDoorParams struct {
	objectivec.Object
}

// MTRDoorLockClusterUnlockDoorParamsFrom constructs a [MTRDoorLockClusterUnlockDoorParams] from an unsafe.Pointer.
func MTRDoorLockClusterUnlockDoorParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterUnlockDoorParams {
	return MTRDoorLockClusterUnlockDoorParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterUnlockDoorParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterUnlockDoorParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterUnlockDoorParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterUnlockDoorParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterUnlockDoorParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockDoorParams/pinCode
func (m_ MTRDoorLockClusterUnlockDoorParams) PinCode() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("pinCode"))
	return rv
}/* debug [instance_properties/getter]: pinCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockDoorParams/pinCode
func (m_ MTRDoorLockClusterUnlockDoorParams) SetPinCode(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPinCode:"), value)
}/* debug [instance_properties/setter]: pinCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockDoorParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterUnlockDoorParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockDoorParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterUnlockDoorParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockDoorParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterUnlockDoorParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockDoorParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterUnlockDoorParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterUnlockDoorParams */




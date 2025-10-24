// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterLockDoorParams */


/* debug [class_header]: Header for MTRDoorLockClusterLockDoorParams */
// The class instance for the [MTRDoorLockClusterLockDoorParams] class.
var (
	MTRDoorLockClusterLockDoorParamsClass     _MTRDoorLockClusterLockDoorParamsClass
	MTRDoorLockClusterLockDoorParamsClassOnce sync.Once
)

func getMTRDoorLockClusterLockDoorParamsClass() _MTRDoorLockClusterLockDoorParamsClass {
	MTRDoorLockClusterLockDoorParamsClassOnce.Do(func() {
		MTRDoorLockClusterLockDoorParamsClass = _MTRDoorLockClusterLockDoorParamsClass{objc.GetClass("MTRDoorLockClusterLockDoorParams")}
	})
	return MTRDoorLockClusterLockDoorParamsClass
}

type _MTRDoorLockClusterLockDoorParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterLockDoorParams */
// An interface definition for the [MTRDoorLockClusterLockDoorParams] class.
type IMTRDoorLockClusterLockDoorParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterLockDoorParams */
	// properties:
	PinCode() objc.IObject /* cross-framework: NSData */
	SetPinCode(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterLockDoorParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterLockDoorParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterLockDoorParamsClass) Alloc() MTRDoorLockClusterLockDoorParams {
	rv := objc.Send[MTRDoorLockClusterLockDoorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterLockDoorParamsClass) New() MTRDoorLockClusterLockDoorParams {
	rv := objc.Send[MTRDoorLockClusterLockDoorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterLockDoorParams) Init() MTRDoorLockClusterLockDoorParams {
	rv := objc.Send[MTRDoorLockClusterLockDoorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterLockDoorParams) Autorelease() MTRDoorLockClusterLockDoorParams {
	rv := objc.Send[MTRDoorLockClusterLockDoorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterLockDoorParams creates a new MTRDoorLockClusterLockDoorParams instance.
func NewMTRDoorLockClusterLockDoorParams() MTRDoorLockClusterLockDoorParams {
	return getMTRDoorLockClusterLockDoorParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterLockDoorParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockDoorParams
type MTRDoorLockClusterLockDoorParams struct {
	objectivec.Object
}

// MTRDoorLockClusterLockDoorParamsFrom constructs a [MTRDoorLockClusterLockDoorParams] from an unsafe.Pointer.
func MTRDoorLockClusterLockDoorParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterLockDoorParams {
	return MTRDoorLockClusterLockDoorParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterLockDoorParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterLockDoorParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterLockDoorParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterLockDoorParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterLockDoorParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockDoorParams/pinCode
func (m_ MTRDoorLockClusterLockDoorParams) PinCode() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("pinCode"))
	return rv
}/* debug [instance_properties/getter]: pinCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockDoorParams/pinCode
func (m_ MTRDoorLockClusterLockDoorParams) SetPinCode(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPinCode:"), value)
}/* debug [instance_properties/setter]: pinCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockDoorParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterLockDoorParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockDoorParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterLockDoorParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockDoorParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterLockDoorParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockDoorParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterLockDoorParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterLockDoorParams */




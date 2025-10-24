// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterSetCredentialResponseParams */


/* debug [class_header]: Header for MTRDoorLockClusterSetCredentialResponseParams */
// The class instance for the [MTRDoorLockClusterSetCredentialResponseParams] class.
var (
	MTRDoorLockClusterSetCredentialResponseParamsClass     _MTRDoorLockClusterSetCredentialResponseParamsClass
	MTRDoorLockClusterSetCredentialResponseParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetCredentialResponseParamsClass() _MTRDoorLockClusterSetCredentialResponseParamsClass {
	MTRDoorLockClusterSetCredentialResponseParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetCredentialResponseParamsClass = _MTRDoorLockClusterSetCredentialResponseParamsClass{objc.GetClass("MTRDoorLockClusterSetCredentialResponseParams")}
	})
	return MTRDoorLockClusterSetCredentialResponseParamsClass
}

type _MTRDoorLockClusterSetCredentialResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterSetCredentialResponseParams */
// An interface definition for the [MTRDoorLockClusterSetCredentialResponseParams] class.
type IMTRDoorLockClusterSetCredentialResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterSetCredentialResponseParams */
	// properties:
	NextCredentialIndex() objc.IObject /* cross-framework: NSNumber */
	SetNextCredentialIndex(value objc.IObject /* cross-framework: NSNumber */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterSetCredentialResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterSetCredentialResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetCredentialResponseParamsClass) Alloc() MTRDoorLockClusterSetCredentialResponseParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterSetCredentialResponseParamsClass) New() MTRDoorLockClusterSetCredentialResponseParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetCredentialResponseParams) Init() MTRDoorLockClusterSetCredentialResponseParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetCredentialResponseParams) Autorelease() MTRDoorLockClusterSetCredentialResponseParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetCredentialResponseParams creates a new MTRDoorLockClusterSetCredentialResponseParams instance.
func NewMTRDoorLockClusterSetCredentialResponseParams() MTRDoorLockClusterSetCredentialResponseParams {
	return getMTRDoorLockClusterSetCredentialResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterSetCredentialResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialResponseParams
type MTRDoorLockClusterSetCredentialResponseParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetCredentialResponseParamsFrom constructs a [MTRDoorLockClusterSetCredentialResponseParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetCredentialResponseParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetCredentialResponseParams {
	return MTRDoorLockClusterSetCredentialResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterSetCredentialResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialResponseParams/init(responseValue:)
func NewMTRDoorLockClusterSetCredentialResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRDoorLockClusterSetCredentialResponseParams {
	instance := getMTRDoorLockClusterSetCredentialResponseParamsClass().Alloc()
	rv := objc.Send[MTRDoorLockClusterSetCredentialResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDoorLockClusterSetCredentialResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterSetCredentialResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterSetCredentialResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterSetCredentialResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterSetCredentialResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialResponseParams/nextCredentialIndex
func (m_ MTRDoorLockClusterSetCredentialResponseParams) NextCredentialIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nextCredentialIndex"))
	return rv
}/* debug [instance_properties/getter]: nextCredentialIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialResponseParams/nextCredentialIndex
func (m_ MTRDoorLockClusterSetCredentialResponseParams) SetNextCredentialIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNextCredentialIndex:"), value)
}/* debug [instance_properties/setter]: nextCredentialIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialResponseParams/status
func (m_ MTRDoorLockClusterSetCredentialResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialResponseParams/status
func (m_ MTRDoorLockClusterSetCredentialResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialResponseParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetCredentialResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialResponseParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetCredentialResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialResponseParams/userIndex
func (m_ MTRDoorLockClusterSetCredentialResponseParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialResponseParams/userIndex
func (m_ MTRDoorLockClusterSetCredentialResponseParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterSetCredentialResponseParams */



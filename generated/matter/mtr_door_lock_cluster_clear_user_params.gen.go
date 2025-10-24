// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterClearUserParams */


/* debug [class_header]: Header for MTRDoorLockClusterClearUserParams */
// The class instance for the [MTRDoorLockClusterClearUserParams] class.
var (
	MTRDoorLockClusterClearUserParamsClass     _MTRDoorLockClusterClearUserParamsClass
	MTRDoorLockClusterClearUserParamsClassOnce sync.Once
)

func getMTRDoorLockClusterClearUserParamsClass() _MTRDoorLockClusterClearUserParamsClass {
	MTRDoorLockClusterClearUserParamsClassOnce.Do(func() {
		MTRDoorLockClusterClearUserParamsClass = _MTRDoorLockClusterClearUserParamsClass{objc.GetClass("MTRDoorLockClusterClearUserParams")}
	})
	return MTRDoorLockClusterClearUserParamsClass
}

type _MTRDoorLockClusterClearUserParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterClearUserParams */
// An interface definition for the [MTRDoorLockClusterClearUserParams] class.
type IMTRDoorLockClusterClearUserParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterClearUserParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterClearUserParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterClearUserParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterClearUserParamsClass) Alloc() MTRDoorLockClusterClearUserParams {
	rv := objc.Send[MTRDoorLockClusterClearUserParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterClearUserParamsClass) New() MTRDoorLockClusterClearUserParams {
	rv := objc.Send[MTRDoorLockClusterClearUserParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterClearUserParams) Init() MTRDoorLockClusterClearUserParams {
	rv := objc.Send[MTRDoorLockClusterClearUserParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterClearUserParams) Autorelease() MTRDoorLockClusterClearUserParams {
	rv := objc.Send[MTRDoorLockClusterClearUserParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterClearUserParams creates a new MTRDoorLockClusterClearUserParams instance.
func NewMTRDoorLockClusterClearUserParams() MTRDoorLockClusterClearUserParams {
	return getMTRDoorLockClusterClearUserParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterClearUserParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearUserParams
type MTRDoorLockClusterClearUserParams struct {
	objectivec.Object
}

// MTRDoorLockClusterClearUserParamsFrom constructs a [MTRDoorLockClusterClearUserParams] from an unsafe.Pointer.
func MTRDoorLockClusterClearUserParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterClearUserParams {
	return MTRDoorLockClusterClearUserParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterClearUserParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterClearUserParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterClearUserParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterClearUserParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterClearUserParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearUserParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterClearUserParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearUserParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterClearUserParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearUserParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterClearUserParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearUserParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterClearUserParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearUserParams/userIndex
func (m_ MTRDoorLockClusterClearUserParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearUserParams/userIndex
func (m_ MTRDoorLockClusterClearUserParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterClearUserParams */




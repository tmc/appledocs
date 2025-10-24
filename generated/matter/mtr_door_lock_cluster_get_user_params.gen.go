// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterGetUserParams */


/* debug [class_header]: Header for MTRDoorLockClusterGetUserParams */
// The class instance for the [MTRDoorLockClusterGetUserParams] class.
var (
	MTRDoorLockClusterGetUserParamsClass     _MTRDoorLockClusterGetUserParamsClass
	MTRDoorLockClusterGetUserParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetUserParamsClass() _MTRDoorLockClusterGetUserParamsClass {
	MTRDoorLockClusterGetUserParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetUserParamsClass = _MTRDoorLockClusterGetUserParamsClass{objc.GetClass("MTRDoorLockClusterGetUserParams")}
	})
	return MTRDoorLockClusterGetUserParamsClass
}

type _MTRDoorLockClusterGetUserParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterGetUserParams */
// An interface definition for the [MTRDoorLockClusterGetUserParams] class.
type IMTRDoorLockClusterGetUserParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterGetUserParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterGetUserParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterGetUserParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetUserParamsClass) Alloc() MTRDoorLockClusterGetUserParams {
	rv := objc.Send[MTRDoorLockClusterGetUserParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterGetUserParamsClass) New() MTRDoorLockClusterGetUserParams {
	rv := objc.Send[MTRDoorLockClusterGetUserParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetUserParams) Init() MTRDoorLockClusterGetUserParams {
	rv := objc.Send[MTRDoorLockClusterGetUserParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetUserParams) Autorelease() MTRDoorLockClusterGetUserParams {
	rv := objc.Send[MTRDoorLockClusterGetUserParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetUserParams creates a new MTRDoorLockClusterGetUserParams instance.
func NewMTRDoorLockClusterGetUserParams() MTRDoorLockClusterGetUserParams {
	return getMTRDoorLockClusterGetUserParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterGetUserParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserParams
type MTRDoorLockClusterGetUserParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetUserParamsFrom constructs a [MTRDoorLockClusterGetUserParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetUserParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetUserParams {
	return MTRDoorLockClusterGetUserParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterGetUserParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterGetUserParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterGetUserParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterGetUserParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterGetUserParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterGetUserParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterGetUserParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetUserParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetUserParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserParams/userIndex
func (m_ MTRDoorLockClusterGetUserParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserParams/userIndex
func (m_ MTRDoorLockClusterGetUserParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterGetUserParams */




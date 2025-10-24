// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterClearCredentialParams */


/* debug [class_header]: Header for MTRDoorLockClusterClearCredentialParams */
// The class instance for the [MTRDoorLockClusterClearCredentialParams] class.
var (
	MTRDoorLockClusterClearCredentialParamsClass     _MTRDoorLockClusterClearCredentialParamsClass
	MTRDoorLockClusterClearCredentialParamsClassOnce sync.Once
)

func getMTRDoorLockClusterClearCredentialParamsClass() _MTRDoorLockClusterClearCredentialParamsClass {
	MTRDoorLockClusterClearCredentialParamsClassOnce.Do(func() {
		MTRDoorLockClusterClearCredentialParamsClass = _MTRDoorLockClusterClearCredentialParamsClass{objc.GetClass("MTRDoorLockClusterClearCredentialParams")}
	})
	return MTRDoorLockClusterClearCredentialParamsClass
}

type _MTRDoorLockClusterClearCredentialParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterClearCredentialParams */
// An interface definition for the [MTRDoorLockClusterClearCredentialParams] class.
type IMTRDoorLockClusterClearCredentialParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterClearCredentialParams */
	// properties:
	Credential() IMTRDoorLockClusterCredentialStruct
	SetCredential(value IMTRDoorLockClusterCredentialStruct)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterClearCredentialParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterClearCredentialParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterClearCredentialParamsClass) Alloc() MTRDoorLockClusterClearCredentialParams {
	rv := objc.Send[MTRDoorLockClusterClearCredentialParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterClearCredentialParamsClass) New() MTRDoorLockClusterClearCredentialParams {
	rv := objc.Send[MTRDoorLockClusterClearCredentialParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterClearCredentialParams) Init() MTRDoorLockClusterClearCredentialParams {
	rv := objc.Send[MTRDoorLockClusterClearCredentialParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterClearCredentialParams) Autorelease() MTRDoorLockClusterClearCredentialParams {
	rv := objc.Send[MTRDoorLockClusterClearCredentialParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterClearCredentialParams creates a new MTRDoorLockClusterClearCredentialParams instance.
func NewMTRDoorLockClusterClearCredentialParams() MTRDoorLockClusterClearCredentialParams {
	return getMTRDoorLockClusterClearCredentialParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterClearCredentialParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearCredentialParams
type MTRDoorLockClusterClearCredentialParams struct {
	objectivec.Object
}

// MTRDoorLockClusterClearCredentialParamsFrom constructs a [MTRDoorLockClusterClearCredentialParams] from an unsafe.Pointer.
func MTRDoorLockClusterClearCredentialParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterClearCredentialParams {
	return MTRDoorLockClusterClearCredentialParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterClearCredentialParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterClearCredentialParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterClearCredentialParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterClearCredentialParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterClearCredentialParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearCredentialParams/credential
func (m_ MTRDoorLockClusterClearCredentialParams) Credential() IMTRDoorLockClusterCredentialStruct {
	rv := objc.Send[MTRDoorLockClusterCredentialStruct](m_.ID, objc.Sel("credential"))
	return rv
}/* debug [instance_properties/getter]: credential */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearCredentialParams/credential
func (m_ MTRDoorLockClusterClearCredentialParams) SetCredential(value IMTRDoorLockClusterCredentialStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredential:"), value)
}/* debug [instance_properties/setter]: credential */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearCredentialParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterClearCredentialParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearCredentialParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterClearCredentialParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearCredentialParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterClearCredentialParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearCredentialParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterClearCredentialParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterClearCredentialParams */




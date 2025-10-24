// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterGetCredentialStatusParams */


/* debug [class_header]: Header for MTRDoorLockClusterGetCredentialStatusParams */
// The class instance for the [MTRDoorLockClusterGetCredentialStatusParams] class.
var (
	MTRDoorLockClusterGetCredentialStatusParamsClass     _MTRDoorLockClusterGetCredentialStatusParamsClass
	MTRDoorLockClusterGetCredentialStatusParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetCredentialStatusParamsClass() _MTRDoorLockClusterGetCredentialStatusParamsClass {
	MTRDoorLockClusterGetCredentialStatusParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetCredentialStatusParamsClass = _MTRDoorLockClusterGetCredentialStatusParamsClass{objc.GetClass("MTRDoorLockClusterGetCredentialStatusParams")}
	})
	return MTRDoorLockClusterGetCredentialStatusParamsClass
}

type _MTRDoorLockClusterGetCredentialStatusParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterGetCredentialStatusParams */
// An interface definition for the [MTRDoorLockClusterGetCredentialStatusParams] class.
type IMTRDoorLockClusterGetCredentialStatusParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterGetCredentialStatusParams */
	// properties:
	Credential() IMTRDoorLockClusterCredentialStruct
	SetCredential(value IMTRDoorLockClusterCredentialStruct)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterGetCredentialStatusParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterGetCredentialStatusParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetCredentialStatusParamsClass) Alloc() MTRDoorLockClusterGetCredentialStatusParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterGetCredentialStatusParamsClass) New() MTRDoorLockClusterGetCredentialStatusParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetCredentialStatusParams) Init() MTRDoorLockClusterGetCredentialStatusParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetCredentialStatusParams) Autorelease() MTRDoorLockClusterGetCredentialStatusParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetCredentialStatusParams creates a new MTRDoorLockClusterGetCredentialStatusParams instance.
func NewMTRDoorLockClusterGetCredentialStatusParams() MTRDoorLockClusterGetCredentialStatusParams {
	return getMTRDoorLockClusterGetCredentialStatusParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterGetCredentialStatusParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusParams
type MTRDoorLockClusterGetCredentialStatusParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetCredentialStatusParamsFrom constructs a [MTRDoorLockClusterGetCredentialStatusParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetCredentialStatusParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetCredentialStatusParams {
	return MTRDoorLockClusterGetCredentialStatusParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterGetCredentialStatusParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterGetCredentialStatusParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterGetCredentialStatusParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterGetCredentialStatusParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterGetCredentialStatusParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusParams/credential
func (m_ MTRDoorLockClusterGetCredentialStatusParams) Credential() IMTRDoorLockClusterCredentialStruct {
	rv := objc.Send[MTRDoorLockClusterCredentialStruct](m_.ID, objc.Sel("credential"))
	return rv
}/* debug [instance_properties/getter]: credential */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusParams/credential
func (m_ MTRDoorLockClusterGetCredentialStatusParams) SetCredential(value IMTRDoorLockClusterCredentialStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredential:"), value)
}/* debug [instance_properties/setter]: credential */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterGetCredentialStatusParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterGetCredentialStatusParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetCredentialStatusParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetCredentialStatusParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterGetCredentialStatusParams */




// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterSetCredentialParams */


/* debug [class_header]: Header for MTRDoorLockClusterSetCredentialParams */
// The class instance for the [MTRDoorLockClusterSetCredentialParams] class.
var (
	MTRDoorLockClusterSetCredentialParamsClass     _MTRDoorLockClusterSetCredentialParamsClass
	MTRDoorLockClusterSetCredentialParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetCredentialParamsClass() _MTRDoorLockClusterSetCredentialParamsClass {
	MTRDoorLockClusterSetCredentialParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetCredentialParamsClass = _MTRDoorLockClusterSetCredentialParamsClass{objc.GetClass("MTRDoorLockClusterSetCredentialParams")}
	})
	return MTRDoorLockClusterSetCredentialParamsClass
}

type _MTRDoorLockClusterSetCredentialParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterSetCredentialParams */
// An interface definition for the [MTRDoorLockClusterSetCredentialParams] class.
type IMTRDoorLockClusterSetCredentialParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterSetCredentialParams */
	// properties:
	Credential() IMTRDoorLockClusterCredentialStruct
	SetCredential(value IMTRDoorLockClusterCredentialStruct)
	CredentialData() objc.IObject /* cross-framework: NSData */
	SetCredentialData(value objc.IObject /* cross-framework: NSData */)
	OperationType() objc.IObject /* cross-framework: NSNumber */
	SetOperationType(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	UserStatus() objc.IObject /* cross-framework: NSNumber */
	SetUserStatus(value objc.IObject /* cross-framework: NSNumber */)
	UserType() objc.IObject /* cross-framework: NSNumber */
	SetUserType(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterSetCredentialParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterSetCredentialParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetCredentialParamsClass) Alloc() MTRDoorLockClusterSetCredentialParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterSetCredentialParamsClass) New() MTRDoorLockClusterSetCredentialParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetCredentialParams) Init() MTRDoorLockClusterSetCredentialParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetCredentialParams) Autorelease() MTRDoorLockClusterSetCredentialParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetCredentialParams creates a new MTRDoorLockClusterSetCredentialParams instance.
func NewMTRDoorLockClusterSetCredentialParams() MTRDoorLockClusterSetCredentialParams {
	return getMTRDoorLockClusterSetCredentialParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterSetCredentialParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams
type MTRDoorLockClusterSetCredentialParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetCredentialParamsFrom constructs a [MTRDoorLockClusterSetCredentialParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetCredentialParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetCredentialParams {
	return MTRDoorLockClusterSetCredentialParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterSetCredentialParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterSetCredentialParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterSetCredentialParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterSetCredentialParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterSetCredentialParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams/credential
func (m_ MTRDoorLockClusterSetCredentialParams) Credential() IMTRDoorLockClusterCredentialStruct {
	rv := objc.Send[MTRDoorLockClusterCredentialStruct](m_.ID, objc.Sel("credential"))
	return rv
}/* debug [instance_properties/getter]: credential */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams/credential
func (m_ MTRDoorLockClusterSetCredentialParams) SetCredential(value IMTRDoorLockClusterCredentialStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredential:"), value)
}/* debug [instance_properties/setter]: credential */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams/credentialData
func (m_ MTRDoorLockClusterSetCredentialParams) CredentialData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("credentialData"))
	return rv
}/* debug [instance_properties/getter]: credentialData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams/credentialData
func (m_ MTRDoorLockClusterSetCredentialParams) SetCredentialData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialData:"), value)
}/* debug [instance_properties/setter]: credentialData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams/operationType
func (m_ MTRDoorLockClusterSetCredentialParams) OperationType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operationType"))
	return rv
}/* debug [instance_properties/getter]: operationType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams/operationType
func (m_ MTRDoorLockClusterSetCredentialParams) SetOperationType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationType:"), value)
}/* debug [instance_properties/setter]: operationType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterSetCredentialParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterSetCredentialParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetCredentialParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetCredentialParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams/userIndex
func (m_ MTRDoorLockClusterSetCredentialParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams/userIndex
func (m_ MTRDoorLockClusterSetCredentialParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams/userStatus
func (m_ MTRDoorLockClusterSetCredentialParams) UserStatus() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userStatus"))
	return rv
}/* debug [instance_properties/getter]: userStatus */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams/userStatus
func (m_ MTRDoorLockClusterSetCredentialParams) SetUserStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserStatus:"), value)
}/* debug [instance_properties/setter]: userStatus */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams/userType
func (m_ MTRDoorLockClusterSetCredentialParams) UserType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userType"))
	return rv
}/* debug [instance_properties/getter]: userType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams/userType
func (m_ MTRDoorLockClusterSetCredentialParams) SetUserType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserType:"), value)
}/* debug [instance_properties/setter]: userType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterSetCredentialParams */




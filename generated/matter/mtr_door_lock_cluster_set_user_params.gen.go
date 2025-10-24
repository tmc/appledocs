// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterSetUserParams */


/* debug [class_header]: Header for MTRDoorLockClusterSetUserParams */
// The class instance for the [MTRDoorLockClusterSetUserParams] class.
var (
	MTRDoorLockClusterSetUserParamsClass     _MTRDoorLockClusterSetUserParamsClass
	MTRDoorLockClusterSetUserParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetUserParamsClass() _MTRDoorLockClusterSetUserParamsClass {
	MTRDoorLockClusterSetUserParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetUserParamsClass = _MTRDoorLockClusterSetUserParamsClass{objc.GetClass("MTRDoorLockClusterSetUserParams")}
	})
	return MTRDoorLockClusterSetUserParamsClass
}

type _MTRDoorLockClusterSetUserParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterSetUserParams */
// An interface definition for the [MTRDoorLockClusterSetUserParams] class.
type IMTRDoorLockClusterSetUserParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterSetUserParams */
	// properties:
	CredentialRule() objc.IObject /* cross-framework: NSNumber */
	SetCredentialRule(value objc.IObject /* cross-framework: NSNumber */)
	OperationType() objc.IObject /* cross-framework: NSNumber */
	SetOperationType(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	UserName() objc.IObject /* cross-framework: NSString */
	SetUserName(value objc.IObject /* cross-framework: NSString */)
	UserStatus() objc.IObject /* cross-framework: NSNumber */
	SetUserStatus(value objc.IObject /* cross-framework: NSNumber */)
	UserType() objc.IObject /* cross-framework: NSNumber */
	SetUserType(value objc.IObject /* cross-framework: NSNumber */)
	UserUniqueId() objc.IObject /* cross-framework: NSNumber */
	SetUserUniqueId(value objc.IObject /* cross-framework: NSNumber */)
	UserUniqueID() objc.IObject /* cross-framework: NSNumber */
	SetUserUniqueID(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterSetUserParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterSetUserParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetUserParamsClass) Alloc() MTRDoorLockClusterSetUserParams {
	rv := objc.Send[MTRDoorLockClusterSetUserParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterSetUserParamsClass) New() MTRDoorLockClusterSetUserParams {
	rv := objc.Send[MTRDoorLockClusterSetUserParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetUserParams) Init() MTRDoorLockClusterSetUserParams {
	rv := objc.Send[MTRDoorLockClusterSetUserParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetUserParams) Autorelease() MTRDoorLockClusterSetUserParams {
	rv := objc.Send[MTRDoorLockClusterSetUserParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetUserParams creates a new MTRDoorLockClusterSetUserParams instance.
func NewMTRDoorLockClusterSetUserParams() MTRDoorLockClusterSetUserParams {
	return getMTRDoorLockClusterSetUserParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterSetUserParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams
type MTRDoorLockClusterSetUserParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetUserParamsFrom constructs a [MTRDoorLockClusterSetUserParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetUserParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetUserParams {
	return MTRDoorLockClusterSetUserParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterSetUserParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterSetUserParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterSetUserParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterSetUserParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterSetUserParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/credentialRule
func (m_ MTRDoorLockClusterSetUserParams) CredentialRule() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("credentialRule"))
	return rv
}/* debug [instance_properties/getter]: credentialRule */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/credentialRule
func (m_ MTRDoorLockClusterSetUserParams) SetCredentialRule(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialRule:"), value)
}/* debug [instance_properties/setter]: credentialRule */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/operationType
func (m_ MTRDoorLockClusterSetUserParams) OperationType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operationType"))
	return rv
}/* debug [instance_properties/getter]: operationType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/operationType
func (m_ MTRDoorLockClusterSetUserParams) SetOperationType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationType:"), value)
}/* debug [instance_properties/setter]: operationType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterSetUserParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterSetUserParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetUserParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetUserParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/userIndex
func (m_ MTRDoorLockClusterSetUserParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/userIndex
func (m_ MTRDoorLockClusterSetUserParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/userName
func (m_ MTRDoorLockClusterSetUserParams) UserName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("userName"))
	return rv
}/* debug [instance_properties/getter]: userName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/userName
func (m_ MTRDoorLockClusterSetUserParams) SetUserName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserName:"), value)
}/* debug [instance_properties/setter]: userName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/userStatus
func (m_ MTRDoorLockClusterSetUserParams) UserStatus() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userStatus"))
	return rv
}/* debug [instance_properties/getter]: userStatus */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/userStatus
func (m_ MTRDoorLockClusterSetUserParams) SetUserStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserStatus:"), value)
}/* debug [instance_properties/setter]: userStatus */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/userType
func (m_ MTRDoorLockClusterSetUserParams) UserType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userType"))
	return rv
}/* debug [instance_properties/getter]: userType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/userType
func (m_ MTRDoorLockClusterSetUserParams) SetUserType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserType:"), value)
}/* debug [instance_properties/setter]: userType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/userUniqueId-22tje
func (m_ MTRDoorLockClusterSetUserParams) UserUniqueId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userUniqueId"))
	return rv
}/* debug [instance_properties/getter]: userUniqueId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/userUniqueId-22tje
func (m_ MTRDoorLockClusterSetUserParams) SetUserUniqueId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserUniqueId:"), value)
}/* debug [instance_properties/setter]: userUniqueId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/userUniqueID-22tka
func (m_ MTRDoorLockClusterSetUserParams) UserUniqueID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userUniqueID"))
	return rv
}/* debug [instance_properties/getter]: userUniqueID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams/userUniqueID-22tka
func (m_ MTRDoorLockClusterSetUserParams) SetUserUniqueID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserUniqueID:"), value)
}/* debug [instance_properties/setter]: userUniqueID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterSetUserParams */




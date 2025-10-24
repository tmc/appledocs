// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterGetUserResponseParams */


/* debug [class_header]: Header for MTRDoorLockClusterGetUserResponseParams */
// The class instance for the [MTRDoorLockClusterGetUserResponseParams] class.
var (
	MTRDoorLockClusterGetUserResponseParamsClass     _MTRDoorLockClusterGetUserResponseParamsClass
	MTRDoorLockClusterGetUserResponseParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetUserResponseParamsClass() _MTRDoorLockClusterGetUserResponseParamsClass {
	MTRDoorLockClusterGetUserResponseParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetUserResponseParamsClass = _MTRDoorLockClusterGetUserResponseParamsClass{objc.GetClass("MTRDoorLockClusterGetUserResponseParams")}
	})
	return MTRDoorLockClusterGetUserResponseParamsClass
}

type _MTRDoorLockClusterGetUserResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterGetUserResponseParams */
// An interface definition for the [MTRDoorLockClusterGetUserResponseParams] class.
type IMTRDoorLockClusterGetUserResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterGetUserResponseParams */
	// properties:
	CreatorFabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetCreatorFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	CredentialRule() objc.IObject /* cross-framework: NSNumber */
	SetCredentialRule(value objc.IObject /* cross-framework: NSNumber */)
	Credentials() objc.IObject /* cross-framework: NSArray */
	SetCredentials(value objc.IObject /* cross-framework: NSArray */)
	LastModifiedFabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetLastModifiedFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	NextUserIndex() objc.IObject /* cross-framework: NSNumber */
	SetNextUserIndex(value objc.IObject /* cross-framework: NSNumber */)
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

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterGetUserResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterGetUserResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetUserResponseParamsClass) Alloc() MTRDoorLockClusterGetUserResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetUserResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterGetUserResponseParamsClass) New() MTRDoorLockClusterGetUserResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetUserResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetUserResponseParams) Init() MTRDoorLockClusterGetUserResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetUserResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetUserResponseParams) Autorelease() MTRDoorLockClusterGetUserResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetUserResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetUserResponseParams creates a new MTRDoorLockClusterGetUserResponseParams instance.
func NewMTRDoorLockClusterGetUserResponseParams() MTRDoorLockClusterGetUserResponseParams {
	return getMTRDoorLockClusterGetUserResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterGetUserResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams
type MTRDoorLockClusterGetUserResponseParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetUserResponseParamsFrom constructs a [MTRDoorLockClusterGetUserResponseParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetUserResponseParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetUserResponseParams {
	return MTRDoorLockClusterGetUserResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterGetUserResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/init(responseValue:)
func NewMTRDoorLockClusterGetUserResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRDoorLockClusterGetUserResponseParams {
	instance := getMTRDoorLockClusterGetUserResponseParamsClass().Alloc()
	rv := objc.Send[MTRDoorLockClusterGetUserResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDoorLockClusterGetUserResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterGetUserResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterGetUserResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterGetUserResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterGetUserResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/creatorFabricIndex
func (m_ MTRDoorLockClusterGetUserResponseParams) CreatorFabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("creatorFabricIndex"))
	return rv
}/* debug [instance_properties/getter]: creatorFabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/creatorFabricIndex
func (m_ MTRDoorLockClusterGetUserResponseParams) SetCreatorFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCreatorFabricIndex:"), value)
}/* debug [instance_properties/setter]: creatorFabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/credentialRule
func (m_ MTRDoorLockClusterGetUserResponseParams) CredentialRule() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("credentialRule"))
	return rv
}/* debug [instance_properties/getter]: credentialRule */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/credentialRule
func (m_ MTRDoorLockClusterGetUserResponseParams) SetCredentialRule(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialRule:"), value)
}/* debug [instance_properties/setter]: credentialRule */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/credentials
func (m_ MTRDoorLockClusterGetUserResponseParams) Credentials() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("credentials"))
	return rv
}/* debug [instance_properties/getter]: credentials */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/credentials
func (m_ MTRDoorLockClusterGetUserResponseParams) SetCredentials(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentials:"), value)
}/* debug [instance_properties/setter]: credentials */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/lastModifiedFabricIndex
func (m_ MTRDoorLockClusterGetUserResponseParams) LastModifiedFabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lastModifiedFabricIndex"))
	return rv
}/* debug [instance_properties/getter]: lastModifiedFabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/lastModifiedFabricIndex
func (m_ MTRDoorLockClusterGetUserResponseParams) SetLastModifiedFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLastModifiedFabricIndex:"), value)
}/* debug [instance_properties/setter]: lastModifiedFabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/nextUserIndex
func (m_ MTRDoorLockClusterGetUserResponseParams) NextUserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nextUserIndex"))
	return rv
}/* debug [instance_properties/getter]: nextUserIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/nextUserIndex
func (m_ MTRDoorLockClusterGetUserResponseParams) SetNextUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNextUserIndex:"), value)
}/* debug [instance_properties/setter]: nextUserIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetUserResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetUserResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/userIndex
func (m_ MTRDoorLockClusterGetUserResponseParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/userIndex
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/userName
func (m_ MTRDoorLockClusterGetUserResponseParams) UserName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("userName"))
	return rv
}/* debug [instance_properties/getter]: userName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/userName
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserName:"), value)
}/* debug [instance_properties/setter]: userName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/userStatus
func (m_ MTRDoorLockClusterGetUserResponseParams) UserStatus() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userStatus"))
	return rv
}/* debug [instance_properties/getter]: userStatus */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/userStatus
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserStatus:"), value)
}/* debug [instance_properties/setter]: userStatus */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/userType
func (m_ MTRDoorLockClusterGetUserResponseParams) UserType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userType"))
	return rv
}/* debug [instance_properties/getter]: userType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/userType
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserType:"), value)
}/* debug [instance_properties/setter]: userType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/userUniqueId-761xi
func (m_ MTRDoorLockClusterGetUserResponseParams) UserUniqueId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userUniqueId"))
	return rv
}/* debug [instance_properties/getter]: userUniqueId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/userUniqueId-761xi
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserUniqueId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserUniqueId:"), value)
}/* debug [instance_properties/setter]: userUniqueId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/userUniqueID-761ye
func (m_ MTRDoorLockClusterGetUserResponseParams) UserUniqueID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userUniqueID"))
	return rv
}/* debug [instance_properties/getter]: userUniqueID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams/userUniqueID-761ye
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserUniqueID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserUniqueID:"), value)
}/* debug [instance_properties/setter]: userUniqueID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterGetUserResponseParams */



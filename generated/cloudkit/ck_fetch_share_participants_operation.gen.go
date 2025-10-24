// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKFetchShareParticipantsOperation */


/* debug [class_header]: Header for CKFetchShareParticipantsOperation */
// The class instance for the [CKFetchShareParticipantsOperation] class.
var (
	CKFetchShareParticipantsOperationClass     _CKFetchShareParticipantsOperationClass
	CKFetchShareParticipantsOperationClassOnce sync.Once
)

func getCKFetchShareParticipantsOperationClass() _CKFetchShareParticipantsOperationClass {
	CKFetchShareParticipantsOperationClassOnce.Do(func() {
		CKFetchShareParticipantsOperationClass = _CKFetchShareParticipantsOperationClass{objc.GetClass("CKFetchShareParticipantsOperation")}
	})
	return CKFetchShareParticipantsOperationClass
}

type _CKFetchShareParticipantsOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKFetchShareParticipantsOperation */
// An interface definition for the [CKFetchShareParticipantsOperation] class.
type ICKFetchShareParticipantsOperation interface {
	ICKOperation
	
/* debug [class_interface_properties]: Properties for CKFetchShareParticipantsOperation */
	// properties:
	FetchShareParticipantsCompletionBlock() unsafe.Pointer
	SetFetchShareParticipantsCompletionBlock(value unsafe.Pointer)
	PerShareParticipantCompletionBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	SetPerShareParticipantCompletionBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer))
	ShareParticipantFetchedBlock() unsafe.Pointer
	SetShareParticipantFetchedBlock(value unsafe.Pointer)
	UserIdentityLookupInfos() []CKUserIdentityLookupInfo
	SetUserIdentityLookupInfos(value []CKUserIdentityLookupInfo)
	FetchShareParticipantsResultBlock() objectivec.IObject
	SetFetchShareParticipantsResultBlock(value objectivec.IObject)
	PerShareParticipantResultBlock() objectivec.IObject
	SetPerShareParticipantResultBlock(value objectivec.IObject)
	CKPartialErrorsByItemIDKey() objc.IObject /* cross-framework: NSString */
	UserIdentity() ICKUserIdentity
	SetUserIdentity(value ICKUserIdentity)
	HasiCloudAccount() bool
	SetHasiCloudAccount(value bool)
	UserInfo() objc.IObject /* cross-framework: NSString */
	SetUserInfo(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKFetchShareParticipantsOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKFetchShareParticipantsOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKFetchShareParticipantsOperationClass) Alloc() CKFetchShareParticipantsOperation {
	rv := objc.Send[CKFetchShareParticipantsOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKFetchShareParticipantsOperationClass) New() CKFetchShareParticipantsOperation {
	rv := objc.Send[CKFetchShareParticipantsOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKFetchShareParticipantsOperation) Init() CKFetchShareParticipantsOperation {
	rv := objc.Send[CKFetchShareParticipantsOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKFetchShareParticipantsOperation) Autorelease() CKFetchShareParticipantsOperation {
	rv := objc.Send[CKFetchShareParticipantsOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchShareParticipantsOperation creates a new CKFetchShareParticipantsOperation instance.
func NewCKFetchShareParticipantsOperation() CKFetchShareParticipantsOperation {
	return getCKFetchShareParticipantsOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKFetchShareParticipantsOperation */
// An operation that converts user identities into share participants.
//
// Participants are a fundamental part of sharing in CloudKit. A participant provides information about a user and their participation in a share, which includes their identity, acceptance status, role, and permissions. The acceptance status manages the user’s visibilty of the shared records. The role and permissions control what actions the user can perform on those records. You don’t create participants. Instead, create an instance of for each user. Provide the user’s email address or phone number, and then use this operation to convert them into participants that you can add to a share. CloudKit limits the number of participants in a share to 100, and each participant must have an active iCloud account. CloudKit queries iCloud for corresponding accounts as part of the operation. If it doesn’t find an account, the server updates the participant’s to reflect that by setting the property to . CloudKit associates a participant with their iCloud account when they accept the share. Anyone with the URL of a public share can become a participant in that share. For a private share, the owner manages its participants. A participant can’t accept a private share unless the owner adds them first. To run the operation, add it to the container’s operation queue. The operation executes its callbacks on a private serial queue. The following example demonstrates how to create the operation, configure it, and then execute it using the default container’s operation queue: The operation calls once for each item you provide, and CloudKit returns the participant, or an error if it can’t generate a particpant. CloudKit also batches per-participant errors. If the operation completes with errors, it returns a error. The error stores the individual errors in its dictionary. Use the key to extract them.


// An operation that converts user identities into share participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation
type CKFetchShareParticipantsOperation struct {
	CKOperation
}

// CKFetchShareParticipantsOperationFrom constructs a [CKFetchShareParticipantsOperation] from an unsafe.Pointer.
//
// An operation that converts user identities into share participants.
func CKFetchShareParticipantsOperationFrom(ptr unsafe.Pointer) CKFetchShareParticipantsOperation {
	return CKFetchShareParticipantsOperation{
		CKOperation: CKOperationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKFetchShareParticipantsOperation */

// Creates an operation for generating share participants from the specified user data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/init(userIdentityLookupInfos:)
func NewCKFetchShareParticipantsOperationWithUserIdentityLookupInfos(userIdentityLookupInfos []CKUserIdentityLookupInfo) CKFetchShareParticipantsOperation {
	instance := getCKFetchShareParticipantsOperationClass().Alloc()
	rv := objc.Send[CKFetchShareParticipantsOperation](instance.ID, objc.Sel("initWithUserIdentityLookupInfos:"), userIdentityLookupInfos)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKFetchShareParticipantsOperationWithUserIdentityLookupInfos */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKFetchShareParticipantsOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKFetchShareParticipantsOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKFetchShareParticipantsOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKFetchShareParticipantsOperation */

// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/fetchShareParticipantsCompletionBlock
func (c_ CKFetchShareParticipantsOperation) FetchShareParticipantsCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchShareParticipantsCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: fetchShareParticipantsCompletionBlock */


// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/fetchShareParticipantsCompletionBlock
func (c_ CKFetchShareParticipantsOperation) SetFetchShareParticipantsCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchShareParticipantsCompletionBlock:"), value)
}/* debug [instance_properties/setter]: fetchShareParticipantsCompletionBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/perShareParticipantCompletionBlock
func (c_ CKFetchShareParticipantsOperation) PerShareParticipantCompletionBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)](c_.ID, objc.Sel("perShareParticipantCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: perShareParticipantCompletionBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/perShareParticipantCompletionBlock
func (c_ CKFetchShareParticipantsOperation) SetPerShareParticipantCompletionBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareParticipantCompletionBlock:"), value)
}/* debug [instance_properties/setter]: perShareParticipantCompletionBlock */


// The closure to execute as the operation generates individual participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/shareParticipantFetchedBlock
func (c_ CKFetchShareParticipantsOperation) ShareParticipantFetchedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("shareParticipantFetchedBlock"))
	return rv
}/* debug [instance_properties/getter]: shareParticipantFetchedBlock */


// The closure to execute as the operation generates individual participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/shareParticipantFetchedBlock
func (c_ CKFetchShareParticipantsOperation) SetShareParticipantFetchedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShareParticipantFetchedBlock:"), value)
}/* debug [instance_properties/setter]: shareParticipantFetchedBlock */


// The user data for the participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/userIdentityLookupInfos
func (c_ CKFetchShareParticipantsOperation) UserIdentityLookupInfos() []CKUserIdentityLookupInfo {
	rv := objc.Send[[]CKUserIdentityLookupInfo](c_.ID, objc.Sel("userIdentityLookupInfos"))
	return rv
}/* debug [instance_properties/getter]: userIdentityLookupInfos */


// The user data for the participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/userIdentityLookupInfos
func (c_ CKFetchShareParticipantsOperation) SetUserIdentityLookupInfos(value []CKUserIdentityLookupInfo) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserIdentityLookupInfos:"), nsArray)
}/* debug [instance_properties/setter]: userIdentityLookupInfos */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/fetchshareparticipantsresultblock
func (c_ CKFetchShareParticipantsOperation) FetchShareParticipantsResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("fetchShareParticipantsResultBlock"))
	return rv
}/* debug [instance_properties/getter]: fetchShareParticipantsResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/fetchshareparticipantsresultblock
func (c_ CKFetchShareParticipantsOperation) SetFetchShareParticipantsResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchShareParticipantsResultBlock:"), value)
}/* debug [instance_properties/setter]: fetchShareParticipantsResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/pershareparticipantresultblock
func (c_ CKFetchShareParticipantsOperation) PerShareParticipantResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("perShareParticipantResultBlock"))
	return rv
}/* debug [instance_properties/getter]: perShareParticipantResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/pershareparticipantresultblock
func (c_ CKFetchShareParticipantsOperation) SetPerShareParticipantResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareParticipantResultBlock:"), value)
}/* debug [instance_properties/setter]: perShareParticipantResultBlock */


// The key to retrieve partial errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckpartialerrorsbyitemidkey
func (c_ CKFetchShareParticipantsOperation) CKPartialErrorsByItemIDKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CKPartialErrorsByItemIDKey"))
	return rv
}/* debug [instance_properties/getter]: CKPartialErrorsByItemIDKey */


// The identity of the participant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/participant/useridentity
func (c_ CKFetchShareParticipantsOperation) UserIdentity() ICKUserIdentity {
	rv := objc.Send[CKUserIdentity](c_.ID, objc.Sel("userIdentity"))
	return rv
}/* debug [instance_properties/getter]: userIdentity */


// The identity of the participant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/participant/useridentity
func (c_ CKFetchShareParticipantsOperation) SetUserIdentity(value ICKUserIdentity) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserIdentity:"), value)
}/* debug [instance_properties/setter]: userIdentity */


// A Boolean value that indicates whether the user has an iCloud account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/hasicloudaccount
func (c_ CKFetchShareParticipantsOperation) HasiCloudAccount() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasiCloudAccount"))
	return rv
}/* debug [instance_properties/getter]: hasiCloudAccount */


// A Boolean value that indicates whether the user has an iCloud account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/hasicloudaccount
func (c_ CKFetchShareParticipantsOperation) SetHasiCloudAccount(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasiCloudAccount:"), value)
}/* debug [instance_properties/setter]: hasiCloudAccount */


// The user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (c_ CKFetchShareParticipantsOperation) UserInfo() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */


// The user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (c_ CKFetchShareParticipantsOperation) SetUserInfo(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserInfo:"), value)
}/* debug [instance_properties/setter]: userInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKFetchShareParticipantsOperation */



// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CKFetchShareParticipantsOperation] class.
type ICKFetchShareParticipantsOperation interface {
	ICKOperation
	FetchShareParticipantsCompletionBlock() unsafe.Pointer
	SetFetchShareParticipantsCompletionBlock(value unsafe.Pointer)
	PerShareParticipantCompletionBlock() unsafe.Pointer
	SetPerShareParticipantCompletionBlock(value unsafe.Pointer)
	ShareParticipantFetchedBlock() unsafe.Pointer
	SetShareParticipantFetchedBlock(value unsafe.Pointer)
	UserIdentityLookupInfos() []CKUserIdentityLookupInfo
	SetUserIdentityLookupInfos(value []CKUserIdentityLookupInfo)
	FetchShareParticipantsResultBlock() unsafe.Pointer
	SetFetchShareParticipantsResultBlock(value unsafe.Pointer)
	PerShareParticipantResultBlock() unsafe.Pointer
	SetPerShareParticipantResultBlock(value unsafe.Pointer)
	CKPartialErrorsByItemIDKey() string
	UserIdentity() CKUserIdentity
	SetUserIdentity(value ICKUserIdentity)
	HasiCloudAccount() bool
	SetHasiCloudAccount(value bool)
	UserInfo() string
	SetUserInfo(value string)
}

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

// Alloc allocates a new instance without initialization.
func (cc _CKFetchShareParticipantsOperationClass) Alloc() CKFetchShareParticipantsOperation {
	rv := objc.Send[CKFetchShareParticipantsOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates an operation for generating share participants from the specified user data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/init(userIdentityLookupInfos:)
func NewCKFetchShareParticipantsOperationWithUserIdentityLookupInfos(userIdentityLookupInfos []CKUserIdentityLookupInfo) CKFetchShareParticipantsOperation {
	instance := getCKFetchShareParticipantsOperationClass().Alloc()
	rv := objc.Send[CKFetchShareParticipantsOperation](instance.ID, objc.Sel("initWithUserIdentityLookupInfos:"), userIdentityLookupInfos)
	rv.Autorelease()
	return rv
}



// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/fetchShareParticipantsCompletionBlock
func (c_ CKFetchShareParticipantsOperation) FetchShareParticipantsCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchShareParticipantsCompletionBlock"))
	return rv
}


// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/fetchShareParticipantsCompletionBlock
func (c_ CKFetchShareParticipantsOperation) SetFetchShareParticipantsCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchShareParticipantsCompletionBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/perShareParticipantCompletionBlock
func (c_ CKFetchShareParticipantsOperation) PerShareParticipantCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perShareParticipantCompletionBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/perShareParticipantCompletionBlock
func (c_ CKFetchShareParticipantsOperation) SetPerShareParticipantCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareParticipantCompletionBlock:"), value)
}


// The closure to execute as the operation generates individual participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/shareParticipantFetchedBlock
func (c_ CKFetchShareParticipantsOperation) ShareParticipantFetchedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("shareParticipantFetchedBlock"))
	return rv
}


// The closure to execute as the operation generates individual participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/shareParticipantFetchedBlock
func (c_ CKFetchShareParticipantsOperation) SetShareParticipantFetchedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShareParticipantFetchedBlock:"), value)
}


// The user data for the participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/userIdentityLookupInfos
func (c_ CKFetchShareParticipantsOperation) UserIdentityLookupInfos() []CKUserIdentityLookupInfo {
	rv := objc.Send[[]CKUserIdentityLookupInfo](c_.ID, objc.Sel("userIdentityLookupInfos"))
	return rv
}


// The user data for the participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/userIdentityLookupInfos
func (c_ CKFetchShareParticipantsOperation) SetUserIdentityLookupInfos(value []CKUserIdentityLookupInfo) {
	// Convert Go slice to NSArray
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
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/fetchshareparticipantsresultblock
func (c_ CKFetchShareParticipantsOperation) FetchShareParticipantsResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchShareParticipantsResultBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/fetchshareparticipantsresultblock
func (c_ CKFetchShareParticipantsOperation) SetFetchShareParticipantsResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchShareParticipantsResultBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/pershareparticipantresultblock
func (c_ CKFetchShareParticipantsOperation) PerShareParticipantResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perShareParticipantResultBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchshareparticipantsoperation/pershareparticipantresultblock
func (c_ CKFetchShareParticipantsOperation) SetPerShareParticipantResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareParticipantResultBlock:"), value)
}


// The key to retrieve partial errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckpartialerrorsbyitemidkey
func (c_ CKFetchShareParticipantsOperation) CKPartialErrorsByItemIDKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CKPartialErrorsByItemIDKey"))
	return rv
}


// The identity of the participant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/participant/useridentity
func (c_ CKFetchShareParticipantsOperation) UserIdentity() CKUserIdentity {
	rv := objc.Send[CKUserIdentity](c_.ID, objc.Sel("userIdentity"))
	return rv
}


// The identity of the participant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/participant/useridentity
func (c_ CKFetchShareParticipantsOperation) SetUserIdentity(value ICKUserIdentity) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserIdentity:"), value)
}


// A Boolean value that indicates whether the user has an iCloud account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/hasicloudaccount
func (c_ CKFetchShareParticipantsOperation) HasiCloudAccount() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasiCloudAccount"))
	return rv
}


// A Boolean value that indicates whether the user has an iCloud account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/hasicloudaccount
func (c_ CKFetchShareParticipantsOperation) SetHasiCloudAccount(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasiCloudAccount:"), value)
}


// The user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (c_ CKFetchShareParticipantsOperation) UserInfo() string {
	rv := objc.Send[string](c_.ID, objc.Sel("userInfo"))
	return rv
}


// The user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (c_ CKFetchShareParticipantsOperation) SetUserInfo(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserInfo:"), objc.String(value))
}



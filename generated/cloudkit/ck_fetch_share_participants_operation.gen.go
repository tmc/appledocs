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
}

// An operation that converts user identities into share participants.
//
// Participants are a fundamental part of sharing in CloudKit. A participant provides information about a user and their participation in a share, which includes their identity, acceptance status, role, and permissions. The acceptance status manages the user’s visibilty of the shared records. The role and permissions control what actions the user can perform on those records. You don’t create participants. Instead, create an instance of for each user. Provide the user’s email address or phone number, and then use this operation to convert them into participants that you can add to a share. CloudKit limits the number of participants in a share to 100, and each participant must have an active iCloud account. CloudKit queries iCloud for corresponding accounts as part of the operation. If it doesn’t find an account, the server updates the participant’s to reflect that by setting the property to . CloudKit associates a participant with their iCloud account when they accept the share. Anyone with the URL of a public share can become a participant in that share. For a private share, the owner manages its participants. A participant can’t accept a private share unless the owner adds them first. To run the operation, add it to the container’s operation queue. The operation executes its callbacks on a private serial queue. The following example demonstrates how to create the operation, configure it, and then execute it using the default container’s operation queue: The operation calls once for each item you provide, and CloudKit returns the participant, or an error if it can’t generate a particpant. CloudKit also batches per-participant errors. If the operation completes with errors, it returns a error. The error stores the individual errors in its dictionary. Use the key to extract them.
//
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
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/init(userIdentityLookupInfos:)
func NewCKFetchShareParticipantsOperationWithUserIdentityLookupInfos(userIdentityLookupInfos unsafe.Pointer) CKFetchShareParticipantsOperation {
	instance := getCKFetchShareParticipantsOperationClass().Alloc()
	rv := objc.Send[CKFetchShareParticipantsOperation](instance.ID, objc.Sel("initWithUserIdentityLookupInfos:"), userIdentityLookupInfos)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/perShareParticipantCompletionBlock
func (c_ CKFetchShareParticipantsOperation) PerShareParticipantCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perShareParticipantCompletionBlock"))
	return rv
}


// SetPerShareParticipantCompletionBlock sets the value of the perShareParticipantCompletionBlock property.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/perShareParticipantCompletionBlock
func (c_ CKFetchShareParticipantsOperation) SetPerShareParticipantCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareParticipantCompletionBlock:"), value)
}
// The user data for the participants.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchShareParticipantsOperation/userIdentityLookupInfos
func (c_ CKFetchShareParticipantsOperation) UserIdentityLookupInfos() []CKUserIdentityLookupInfo {
	rv := objc.Send[[]CKUserIdentityLookupInfo](c_.ID, objc.Sel("userIdentityLookupInfos"))
	return rv
}


// SetUserIdentityLookupInfos sets the value of the userIdentityLookupInfos property.
// The user data for the participants.

//
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


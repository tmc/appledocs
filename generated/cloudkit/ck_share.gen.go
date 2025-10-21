// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CKShare] class.
var (
	CKShareClass     _CKShareClass
	CKShareClassOnce sync.Once
)

func getCKShareClass() _CKShareClass {
	CKShareClassOnce.Do(func() {
		CKShareClass = _CKShareClass{objc.GetClass("CKShare")}
	})
	return CKShareClass
}

type _CKShareClass struct {
	class objc.Class
}

// An interface definition for the [CKShare] class.
type ICKShare interface {
	ICKRecord
	AddParticipant(participant unsafe.Pointer)
	RemoveParticipant(participant unsafe.Pointer)
}

// A specialized record type that manages a collection of shared records.
//
// A share is a specialized record type that facilitates the sharing of one or more records with many participants. You store shareable records in a custom record zone in the user’s private database. As you create records in that zone, they become eligible for record zone sharing. If you want to share a specific hierarchy of related records, rather than the entire record zone, set each record’s property to define the relationship with its parent. CloudKit infers the shared hierarchy using only the property, and ignores any custom reference fields. You create a share with either the ID of the record zone to share, or the root record, which defines the point in a record hierarchy where you want to start sharing. CloudKit shares all the records in the record zone, or every record in the hierarchy below the root. If you set the root record’s property, CloudKit ignores it. A record can take part in only a single share. This applies to every record in the shared record zone or hierarchy. If a record is participating in another share, any attempt to save the share fails, and CloudKit returns an error. Use to save the share to the server. The initial set of records the share includes must exist on the server or be part of the same save operation to succeed. CloudKit then updates the share’s property. Use to present options to the user for sharing the URL. Otherwise, distribute the URL to any participants you add to the share. You can allow anyone with the URL to take part in the share by setting to a value more permissive than . After CloudKit saves the share, a participant can fetch its corresponding metadata, which includes a reference to the share, information about the user’s participation, and, for shared hierarchies, the root record’s record ID. Create an instance of using the share’s URL and add it to the container’s queue to execute it. The operation returns an instance of for each URL you provide. This is only applicable if you manually process share acceptance. If a user receives the share URL and taps or clicks it, CloudKit automatically processes their participation. To determine the configuration of a fetched share, inspect the property of its . If the value is , the share is managing a shared record zone; otherwise, it’s managing a shared record hierarchy. CloudKit limits the number of participants in a share to 100, and each participant must have an active iCloud account. You don’t create participants. Instead, use to manage a share’s participants and their permissions. Alternatively, create an instance of for each user. Provide the user’s email address or phone number, and use to fetch the corresponding participants. CloudKit queries iCloud for corresponding accounts as part of the operation. If it doesn’t find an account, the server updates the participant’s to reflect that by setting the property to . CloudKit associates the participant with their iCloud account when they accept the share if they launch the process by tapping or clicking the share URL. Participants with write permissions can modify or delete any record that you include in the share. However, only the owner can delete a shared hierarchy’s root record. If a participant attempts to delete the share, CloudKit removes the participant. The share remains active for all other participants. If the owner deletes a share that manages a record hierarchy, CloudKit sets the root record’s property to . CloudKit deletes the share if the owner of the shared heirarchy deletes its root record. You can customize the title and image the system displays when initiating a share or accepting an invitation to participate. You can also provide a custom UTI to indicate the content of the shared records. Use the keys that defines, as the following example shows:
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare
type CKShare struct {
	CKRecord
}

// CKShareFrom constructs a [CKShare] from an unsafe.Pointer.
//
// A specialized record type that manages a collection of shared records.
func CKShareFrom(ptr unsafe.Pointer) CKShare {
	return CKShare{
		CKRecord: CKRecordFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKShareClass) Alloc() CKShare {
	rv := objc.Send[CKShare](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKShareClass) New() CKShare {
	rv := objc.Send[CKShare](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKShare) Init() CKShare {
	rv := objc.Send[CKShare](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKShare) Autorelease() CKShare {
	rv := objc.Send[CKShare](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKShare creates a new CKShare instance.
func NewCKShare() CKShare {
	return getCKShareClass().New()
}




// Creates a share from a serialized instance.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/init(coder:)
func NewCKShareWithCoder(aDecoder unsafe.Pointer) CKShare {
	instance := getCKShareClass().Alloc()
	rv := objc.Send[CKShare](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}



// Creates a new share for the specified record zone.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/init(recordZoneID:)
func NewCKShareWithRecordZoneID(recordZoneID unsafe.Pointer) CKShare {
	instance := getCKShareClass().Alloc()
	rv := objc.Send[CKShare](instance.ID, objc.Sel("initWithRecordZoneID:"), recordZoneID)
	rv.Autorelease()
	return rv
}



// Creates a new share for the specified record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/init(rootRecord:)
func NewCKShareWithRootRecord(rootRecord unsafe.Pointer) CKShare {
	instance := getCKShareClass().Alloc()
	rv := objc.Send[CKShare](instance.ID, objc.Sel("initWithRootRecord:"), rootRecord)
	rv.Autorelease()
	return rv
}



// Creates a new share for the specified record and record ID.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/init(rootRecord:shareID:)
func NewCKShareWithRootRecordShareID(rootRecord unsafe.Pointer, shareID unsafe.Pointer) CKShare {
	instance := getCKShareClass().Alloc()
	rv := objc.Send[CKShare](instance.ID, objc.Sel("initWithRootRecord:shareID:"), rootRecord, shareID)
	rv.Autorelease()
	return rv
}


// Adds a participant to the share.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/addParticipant(_:)
func (c_ CKShare) AddParticipant(participant unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addParticipant:"), participant)
}

// Removes a participant from the share.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/removeParticipant(_:)
func (c_ CKShare) RemoveParticipant(participant unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeParticipant:"), participant)
}

// The permission for anyone with access to the share’s URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/publicPermission
func (c_ CKShare) PublicPermission() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("publicPermission"))
	return rv
}


// SetPublicPermission sets the value of the publicPermission property.
// The permission for anyone with access to the share’s URL.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/publicPermission
func (c_ CKShare) SetPublicPermission(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPublicPermission:"), value)
}

// The URL for inviting participants to the share.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/url
func (c_ CKShare) URL() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("URL"))
	return rv
}

// The unique name of the record.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/id/recordname
func (c_ CKShare) RecordName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("recordName"))
	return rv
}


// SetRecordName sets the value of the recordName property.
// The unique name of the record.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/id/recordname
func (c_ CKShare) SetRecordName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordName:"), objc.String(value))
}

// A reference to the record’s parent record.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/parent
func (c_ CKShare) Parent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("parent"))
	return rv
}


// SetParent sets the value of the parent property.
// A reference to the record’s parent record.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/parent
func (c_ CKShare) SetParent(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setParent:"), value)
}

// The unique ID of the record.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordid
func (c_ CKShare) RecordID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordID"))
	return rv
}


// SetRecordID sets the value of the recordID property.
// The unique ID of the record.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordid
func (c_ CKShare) SetRecordID(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordID:"), value)
}

// A reference to the share object that determines the share status of the record.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/share
func (c_ CKShare) Share() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("share"))
	return rv
}


// SetShare sets the value of the share property.
// A reference to the share object that determines the share status of the record.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/share
func (c_ CKShare) SetShare(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShare:"), value)
}

// The name of a share record that manages a shared record zone.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordnamezonewideshare
func (c_ CKShare) CKRecordNameZoneWideShare() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CKRecordNameZoneWideShare"))
	return rv
}

// The identity of the participant.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/participant/useridentity
func (c_ CKShare) UserIdentity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("userIdentity"))
	return rv
}


// SetUserIdentity sets the value of the userIdentity property.
// The identity of the participant.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/participant/useridentity
func (c_ CKShare) SetUserIdentity(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserIdentity:"), value)
}

// Indicates whether uninvited users can request access to this share.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/allowsaccessrequests
func (c_ CKShare) AllowsAccessRequests() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsAccessRequests"))
	return rv
}


// SetAllowsAccessRequests sets the value of the allowsAccessRequests property.
// Indicates whether uninvited users can request access to this share.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/allowsaccessrequests
func (c_ CKShare) SetAllowsAccessRequests(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsAccessRequests:"), value)
}

// A list of users blocked from requesting access to this share.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/blockedidentities
func (c_ CKShare) BlockedIdentities() cloudkit.CKShare {
	rv := objc.Send[cloudkit.CKShare](c_.ID, objc.Sel("blockedIdentities"))
	return rv
}


// SetBlockedIdentities sets the value of the blockedIdentities property.
// A list of users blocked from requesting access to this share.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/blockedidentities
func (c_ CKShare) SetBlockedIdentities(value cloudkit.CKShare) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBlockedIdentities:"), value)
}

// The participant that represents the current user.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/currentuserparticipant
func (c_ CKShare) CurrentUserParticipant() cloudkit.CKShare {
	rv := objc.Send[cloudkit.CKShare](c_.ID, objc.Sel("currentUserParticipant"))
	return rv
}


// SetCurrentUserParticipant sets the value of the currentUserParticipant property.
// The participant that represents the current user.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/currentuserparticipant
func (c_ CKShare) SetCurrentUserParticipant(value cloudkit.CKShare) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCurrentUserParticipant:"), value)
}

// The participant that represents the share’s owner.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/owner
func (c_ CKShare) Owner() cloudkit.CKShare {
	rv := objc.Send[cloudkit.CKShare](c_.ID, objc.Sel("owner"))
	return rv
}


// SetOwner sets the value of the owner property.
// The participant that represents the share’s owner.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/owner
func (c_ CKShare) SetOwner(value cloudkit.CKShare) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOwner:"), value)
}

// An array that contains the share’s participants.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/participants
func (c_ CKShare) Participants() cloudkit.CKShare {
	rv := objc.Send[cloudkit.CKShare](c_.ID, objc.Sel("participants"))
	return rv
}


// SetParticipants sets the value of the participants property.
// An array that contains the share’s participants.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/participants
func (c_ CKShare) SetParticipants(value cloudkit.CKShare) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setParticipants:"), value)
}

// A list of all uninvited users who have requested access to this share.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/requesters
func (c_ CKShare) Requesters() cloudkit.CKShare {
	rv := objc.Send[cloudkit.CKShare](c_.ID, objc.Sel("requesters"))
	return rv
}


// SetRequesters sets the value of the requesters property.
// A list of all uninvited users who have requested access to this share.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/requesters
func (c_ CKShare) SetRequesters(value cloudkit.CKShare) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequesters:"), value)
}

// A Boolean value that indicates whether the user has an iCloud account.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/hasicloudaccount
func (c_ CKShare) HasiCloudAccount() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasiCloudAccount"))
	return rv
}


// SetHasiCloudAccount sets the value of the hasiCloudAccount property.
// A Boolean value that indicates whether the user has an iCloud account.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/hasicloudaccount
func (c_ CKShare) SetHasiCloudAccount(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasiCloudAccount:"), value)
}



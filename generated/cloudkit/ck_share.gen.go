// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKShare */


/* debug [class_header]: Header for CKShare */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKShare */
// An interface definition for the [CKShare] class.
type ICKShare interface {
	ICKRecord
	
/* debug [class_interface_properties]: Properties for CKShare */
	// properties:
	AllowsAccessRequests() bool
	SetAllowsAccessRequests(value bool)
	BlockedIdentities() []CKShareBlockedIdentity
	CurrentUserParticipant() ICKShareParticipant
	Owner() ICKShareParticipant
	Participants() []CKShareParticipant
	PublicPermission() CKShareParticipantPermission
	SetPublicPermission(value CKShareParticipantPermission)
	Requesters() []CKShareAccessRequester
	URL() objc.IObject /* cross-framework: NSURL */
	RecordName() objc.IObject /* cross-framework: NSString */
	SetRecordName(value objc.IObject /* cross-framework: NSString */)
	Parent() ICKReference
	SetParent(value ICKReference)
	RecordID() ICKRecordID
	SetRecordID(value ICKRecordID)
	Share() ICKReference
	SetShare(value ICKReference)
	CKRecordNameZoneWideShare() objc.IObject /* cross-framework: NSString */
	UserIdentity() ICKUserIdentity
	SetUserIdentity(value ICKUserIdentity)
	HasiCloudAccount() bool
	SetHasiCloudAccount(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKShare */
	// methods:
	AddParticipant(participant ICKShareParticipant)
	BlockRequesters(requesters []CKShareAccessRequester)
	DenyRequesters(requesters []CKShareAccessRequester)
	OneTimeURLForParticipantID(participantID objc.IObject /* cross-framework: NSString */) foundation.URL
	RemoveParticipant(participant ICKShareParticipant)
	UnblockIdentities(blockedIdentities []CKShareBlockedIdentity)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKShare */
// Alloc allocates a new instance without initialization.
func (cc _CKShareClass) Alloc() CKShare {
	rv := objc.Send[CKShare](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKShare */
// A specialized record type that manages a collection of shared records.
//
// A share is a specialized record type that facilitates the sharing of one or more records with many participants. You store shareable records in a custom record zone in the user’s private database. As you create records in that zone, they become eligible for record zone sharing. If you want to share a specific hierarchy of related records, rather than the entire record zone, set each record’s property to define the relationship with its parent. CloudKit infers the shared hierarchy using only the property, and ignores any custom reference fields. You create a share with either the ID of the record zone to share, or the root record, which defines the point in a record hierarchy where you want to start sharing. CloudKit shares all the records in the record zone, or every record in the hierarchy below the root. If you set the root record’s property, CloudKit ignores it. A record can take part in only a single share. This applies to every record in the shared record zone or hierarchy. If a record is participating in another share, any attempt to save the share fails, and CloudKit returns an error. Use to save the share to the server. The initial set of records the share includes must exist on the server or be part of the same save operation to succeed. CloudKit then updates the share’s property. Use to present options to the user for sharing the URL. Otherwise, distribute the URL to any participants you add to the share. You can allow anyone with the URL to take part in the share by setting to a value more permissive than . After CloudKit saves the share, a participant can fetch its corresponding metadata, which includes a reference to the share, information about the user’s participation, and, for shared hierarchies, the root record’s record ID. Create an instance of using the share’s URL and add it to the container’s queue to execute it. The operation returns an instance of for each URL you provide. This is only applicable if you manually process share acceptance. If a user receives the share URL and taps or clicks it, CloudKit automatically processes their participation. To determine the configuration of a fetched share, inspect the property of its . If the value is , the share is managing a shared record zone; otherwise, it’s managing a shared record hierarchy. CloudKit limits the number of participants in a share to 100, and each participant must have an active iCloud account. You don’t create participants. Instead, use to manage a share’s participants and their permissions. Alternatively, create an instance of for each user. Provide the user’s email address or phone number, and use to fetch the corresponding participants. CloudKit queries iCloud for corresponding accounts as part of the operation. If it doesn’t find an account, the server updates the participant’s to reflect that by setting the property to . CloudKit associates the participant with their iCloud account when they accept the share if they launch the process by tapping or clicking the share URL. Participants with write permissions can modify or delete any record that you include in the share. However, only the owner can delete a shared hierarchy’s root record. If a participant attempts to delete the share, CloudKit removes the participant. The share remains active for all other participants. If the owner deletes a share that manages a record hierarchy, CloudKit sets the root record’s property to . CloudKit deletes the share if the owner of the shared heirarchy deletes its root record. You can customize the title and image the system displays when initiating a share or accepting an invitation to participate. You can also provide a custom UTI to indicate the content of the shared records. Use the keys that defines, as the following example shows:


// A specialized record type that manages a collection of shared records.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKShare */

// Creates a share from a serialized instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/init(coder:)
func NewCKShareWithCoder(aDecoder foundation.Coder) CKShare {
	instance := getCKShareClass().Alloc()
	rv := objc.Send[CKShare](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKShareWithCoder */


// Creates a new share for the specified record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/init(recordZoneID:)
func NewCKShareWithRecordZoneID(recordZoneID ICKRecordZoneID) CKShare {
	instance := getCKShareClass().Alloc()
	rv := objc.Send[CKShare](instance.ID, objc.Sel("initWithRecordZoneID:"), recordZoneID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKShareWithRecordZoneID */


// Creates a new share for the specified record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/init(rootRecord:)
func NewCKShareWithRootRecord(rootRecord objc.IObject /* cross-framework: CKRecord */) CKShare {
	instance := getCKShareClass().Alloc()
	rv := objc.Send[CKShare](instance.ID, objc.Sel("initWithRootRecord:"), rootRecord)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKShareWithRootRecord */


// Creates a new share for the specified record and record ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/init(rootRecord:shareID:)
func NewCKShareWithRootRecordShareID(rootRecord objc.IObject /* cross-framework: CKRecord */, shareID ICKRecordID) CKShare {
	instance := getCKShareClass().Alloc()
	rv := objc.Send[CKShare](instance.ID, objc.Sel("initWithRootRecord:shareID:"), rootRecord, shareID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKShareWithRootRecordShareID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKShare */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKShare */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKShare */

// Adds a participant to the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/addParticipant(_:)
func (c_ CKShare) AddParticipant(participant ICKShareParticipant) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addParticipant:"), participant)
}/* debug [instance_methods/method]: AddParticipant */


// Blocks specified users from requesting access to this share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/blockRequesters(_:)
func (c_ CKShare) BlockRequesters(requesters []CKShareAccessRequester) {
	objc.Send[objc.ID](c_.ID, objc.Sel("blockRequesters:"), requesters)
}/* debug [instance_methods/method]: BlockRequesters */


// Denies access requests from specified users.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/denyRequesters(_:)
func (c_ CKShare) DenyRequesters(requesters []CKShareAccessRequester) {
	objc.Send[objc.ID](c_.ID, objc.Sel("denyRequesters:"), requesters)
}/* debug [instance_methods/method]: DenyRequesters */


// Invitation URLs that can be used by any receiver to claim the associated participantID and join the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/oneTimeURLForParticipantID:
func (c_ CKShare) OneTimeURLForParticipantID(participantID objc.IObject /* cross-framework: NSString */) foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("oneTimeURLForParticipantID:"), participantID)
	return rv
}/* debug [instance_methods/method]: OneTimeURLForParticipantID */


// Removes a participant from the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/removeParticipant(_:)
func (c_ CKShare) RemoveParticipant(participant ICKShareParticipant) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeParticipant:"), participant)
}/* debug [instance_methods/method]: RemoveParticipant */


// Unblocks previously blocked users, allowing them to request access again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/unblockIdentities(_:)
func (c_ CKShare) UnblockIdentities(blockedIdentities []CKShareBlockedIdentity) {
	objc.Send[objc.ID](c_.ID, objc.Sel("unblockIdentities:"), blockedIdentities)
}/* debug [instance_methods/method]: UnblockIdentities */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKShare */

// Indicates whether uninvited users can request access to this share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/allowsAccessRequests
func (c_ CKShare) AllowsAccessRequests() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsAccessRequests"))
	return rv
}/* debug [instance_properties/getter]: allowsAccessRequests */


// Indicates whether uninvited users can request access to this share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/allowsAccessRequests
func (c_ CKShare) SetAllowsAccessRequests(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsAccessRequests:"), value)
}/* debug [instance_properties/setter]: allowsAccessRequests */


// A list of users blocked from requesting access to this share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/blockedIdentities
func (c_ CKShare) BlockedIdentities() []CKShareBlockedIdentity {
	rv := objc.Send[[]CKShareBlockedIdentity](c_.ID, objc.Sel("blockedIdentities"))
	return rv
}/* debug [instance_properties/getter]: blockedIdentities */


// The participant that represents the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/currentUserParticipant
func (c_ CKShare) CurrentUserParticipant() ICKShareParticipant {
	rv := objc.Send[CKShareParticipant](c_.ID, objc.Sel("currentUserParticipant"))
	return rv
}/* debug [instance_properties/getter]: currentUserParticipant */


// The participant that represents the share’s owner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/owner
func (c_ CKShare) Owner() ICKShareParticipant {
	rv := objc.Send[CKShareParticipant](c_.ID, objc.Sel("owner"))
	return rv
}/* debug [instance_properties/getter]: owner */


// An array that contains the share’s participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/participants
func (c_ CKShare) Participants() []CKShareParticipant {
	rv := objc.Send[[]CKShareParticipant](c_.ID, objc.Sel("participants"))
	return rv
}/* debug [instance_properties/getter]: participants */


// The permission for anyone with access to the share’s URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/publicPermission
func (c_ CKShare) PublicPermission() CKShareParticipantPermission {
	rv := objc.Send[CKShareParticipantPermission](c_.ID, objc.Sel("publicPermission"))
	return rv
}/* debug [instance_properties/getter]: publicPermission */


// The permission for anyone with access to the share’s URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/publicPermission
func (c_ CKShare) SetPublicPermission(value CKShareParticipantPermission) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPublicPermission:"), value)
}/* debug [instance_properties/setter]: publicPermission */


// A list of all uninvited users who have requested access to this share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/requesters
func (c_ CKShare) Requesters() []CKShareAccessRequester {
	rv := objc.Send[[]CKShareAccessRequester](c_.ID, objc.Sel("requesters"))
	return rv
}/* debug [instance_properties/getter]: requesters */


// The URL for inviting participants to the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/url
func (c_ CKShare) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](c_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// The unique name of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/id/recordname
func (c_ CKShare) RecordName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("recordName"))
	return rv
}/* debug [instance_properties/getter]: recordName */


// The unique name of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/id/recordname
func (c_ CKShare) SetRecordName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordName:"), value)
}/* debug [instance_properties/setter]: recordName */


// A reference to the record’s parent record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/parent
func (c_ CKShare) Parent() ICKReference {
	rv := objc.Send[CKReference](c_.ID, objc.Sel("parent"))
	return rv
}/* debug [instance_properties/getter]: parent */


// A reference to the record’s parent record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/parent
func (c_ CKShare) SetParent(value ICKReference) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setParent:"), value)
}/* debug [instance_properties/setter]: parent */


// The unique ID of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordid
func (c_ CKShare) RecordID() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("recordID"))
	return rv
}/* debug [instance_properties/getter]: recordID */


// The unique ID of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/recordid
func (c_ CKShare) SetRecordID(value ICKRecordID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordID:"), value)
}/* debug [instance_properties/setter]: recordID */


// A reference to the share object that determines the share status of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/share
func (c_ CKShare) Share() ICKReference {
	rv := objc.Send[CKReference](c_.ID, objc.Sel("share"))
	return rv
}/* debug [instance_properties/getter]: share */


// A reference to the share object that determines the share status of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/share
func (c_ CKShare) SetShare(value ICKReference) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShare:"), value)
}/* debug [instance_properties/setter]: share */


// The name of a share record that manages a shared record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordnamezonewideshare
func (c_ CKShare) CKRecordNameZoneWideShare() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CKRecordNameZoneWideShare"))
	return rv
}/* debug [instance_properties/getter]: CKRecordNameZoneWideShare */


// The identity of the participant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/participant/useridentity
func (c_ CKShare) UserIdentity() ICKUserIdentity {
	rv := objc.Send[CKUserIdentity](c_.ID, objc.Sel("userIdentity"))
	return rv
}/* debug [instance_properties/getter]: userIdentity */


// The identity of the participant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/participant/useridentity
func (c_ CKShare) SetUserIdentity(value ICKUserIdentity) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserIdentity:"), value)
}/* debug [instance_properties/setter]: userIdentity */


// A Boolean value that indicates whether the user has an iCloud account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/hasicloudaccount
func (c_ CKShare) HasiCloudAccount() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasiCloudAccount"))
	return rv
}/* debug [instance_properties/getter]: hasiCloudAccount */


// A Boolean value that indicates whether the user has an iCloud account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/hasicloudaccount
func (c_ CKShare) SetHasiCloudAccount(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasiCloudAccount:"), value)
}/* debug [instance_properties/setter]: hasiCloudAccount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKShare */



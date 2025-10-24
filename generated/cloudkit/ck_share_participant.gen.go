// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKShareParticipant */


/* debug [class_header]: Header for CKShareParticipant */
// The class instance for the [CKShareParticipant] class.
var (
	CKShareParticipantClass     _CKShareParticipantClass
	CKShareParticipantClassOnce sync.Once
)

func getCKShareParticipantClass() _CKShareParticipantClass {
	CKShareParticipantClassOnce.Do(func() {
		CKShareParticipantClass = _CKShareParticipantClass{objc.GetClass("CKShareParticipant")}
	})
	return CKShareParticipantClass
}

type _CKShareParticipantClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKShareParticipant */
// An interface definition for the [CKShareParticipant] class.
type ICKShareParticipant interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKShareParticipant */
	// properties:
	AcceptanceStatus() CKShareParticipantAcceptanceStatus
	DateAddedToShare() objc.IObject /* cross-framework: NSDate */
	IsApprovedRequester() bool
	Permission() CKShareParticipantPermission
	SetPermission(value CKShareParticipantPermission)
	Role() CKShareParticipantRole
	SetRole(value CKShareParticipantRole)
	Type() CKShareParticipantType
	SetType(value CKShareParticipantType)
	UserIdentity() ICKUserIdentity
	ParticipantID() objc.IObject /* cross-framework: NSString */
	Participants() ICKShareParticipant
	SetParticipants(value ICKShareParticipant)
	PublicPermission() objectivec.IObject
	SetPublicPermission(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKShareParticipant */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKShareParticipant */
// Alloc allocates a new instance without initialization.
func (cc _CKShareParticipantClass) Alloc() CKShareParticipant {
	rv := objc.Send[CKShareParticipant](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKShareParticipantClass) New() CKShareParticipant {
	rv := objc.Send[CKShareParticipant](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKShareParticipant) Init() CKShareParticipant {
	rv := objc.Send[CKShareParticipant](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKShareParticipant) Autorelease() CKShareParticipant {
	rv := objc.Send[CKShareParticipant](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKShareParticipant creates a new CKShareParticipant instance.
func NewCKShareParticipant() CKShareParticipant {
	return getCKShareParticipantClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKShareParticipant */
// An object that describes a user’s participation in a share.
//
// Participants are a key element of sharing in CloudKit. A participant provides information about an iCloud user and their participation in a share, including their identity, acceptance status, permissions, and role. The acceptance status determines the participant’s visibilty of the shared records. Statuses are: , , , and . If the status is , use to accept the share. Upon acceptance, CloudKit makes the shared records available in the participant’s shared database. The records remain accessible for as long as the participant’s status is . You don’t create participants. Use the share’s property to access its existing participants. Use to manage the share’s participants and their permissions. Alternatively, you can generate participants using . Participants must have an active iCloud account. Anyone with the URL of a public share can become a participant in that share. Participants of a public share assume the role. For private shares, the owner manages the participants. An owner is any participant with the role. A participant of a private share can’t accept the share unless the owner adds them first. Private share participants assume the role. CloudKit removes any pending participants if the owner changes the share’s . CloudKit removes all participants if the new permission is . Participants with write permissions can modify or delete any record that you include in the share. However, only the owner can delete a shared hierarchy’s root record. If a participant attempts to delete the share, CloudKit removes the participant. The share remains active for all other participants.


// An object that describes a user’s participation in a share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant
type CKShareParticipant struct {
	objectivec.Object
}

// CKShareParticipantFrom constructs a [CKShareParticipant] from an unsafe.Pointer.
//
// An object that describes a user’s participation in a share.
func CKShareParticipantFrom(ptr unsafe.Pointer) CKShareParticipant {
	return CKShareParticipant{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKShareParticipant *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKShareParticipant */

// Generate a unique URL for inviting a participant without knowing their handle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/oneTimeURLParticipant()
func (cc _CKShareParticipantClass) OneTimeURLParticipant() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("oneTimeURLParticipant"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OneTimeURLParticipant) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKShareParticipant */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKShareParticipant */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKShareParticipant */

// The current state of the user’s acceptance of the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/acceptanceStatus-swift.property
func (c_ CKShareParticipant) AcceptanceStatus() CKShareParticipantAcceptanceStatus {
	rv := objc.Send[CKShareParticipantAcceptanceStatus](c_.ID, objc.Sel("acceptanceStatus"))
	return rv
}/* debug [instance_properties/getter]: acceptanceStatus */


// The date and time when the participant was added to the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/dateAddedToShare
func (c_ CKShareParticipant) DateAddedToShare() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("dateAddedToShare"))
	return rv
}/* debug [instance_properties/getter]: dateAddedToShare */


// Indicates whether the participant was originally a requester who was approved to join the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/isApprovedRequester
func (c_ CKShareParticipant) IsApprovedRequester() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isApprovedRequester"))
	return rv
}/* debug [instance_properties/getter]: isApprovedRequester */


// The participant’s permission level for the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/permission-swift.property
func (c_ CKShareParticipant) Permission() CKShareParticipantPermission {
	rv := objc.Send[CKShareParticipantPermission](c_.ID, objc.Sel("permission"))
	return rv
}/* debug [instance_properties/getter]: permission */


// The participant’s permission level for the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/permission-swift.property
func (c_ CKShareParticipant) SetPermission(value CKShareParticipantPermission) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPermission:"), value)
}/* debug [instance_properties/setter]: permission */


// The participant’s role for the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/role-swift.property
func (c_ CKShareParticipant) Role() CKShareParticipantRole {
	rv := objc.Send[CKShareParticipantRole](c_.ID, objc.Sel("role"))
	return rv
}/* debug [instance_properties/getter]: role */


// The participant’s role for the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/role-swift.property
func (c_ CKShareParticipant) SetRole(value CKShareParticipantRole) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRole:"), value)
}/* debug [instance_properties/setter]: role */


// The participant type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/type
func (c_ CKShareParticipant) Type() CKShareParticipantType {
	rv := objc.Send[CKShareParticipantType](c_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The participant type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/type
func (c_ CKShareParticipant) SetType(value CKShareParticipantType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */


// The identity of the participant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/userIdentity
func (c_ CKShareParticipant) UserIdentity() ICKUserIdentity {
	rv := objc.Send[CKUserIdentity](c_.ID, objc.Sel("userIdentity"))
	return rv
}/* debug [instance_properties/getter]: userIdentity */


// A unique identifier for this participant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareParticipant/participantID
func (c_ CKShareParticipant) ParticipantID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("participantID"))
	return rv
}/* debug [instance_properties/getter]: participantID */


// An array that contains the share’s participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/participants
func (c_ CKShareParticipant) Participants() ICKShareParticipant {
	rv := objc.Send[CKShareParticipant](c_.ID, objc.Sel("participants"))
	return rv
}/* debug [instance_properties/getter]: participants */


// An array that contains the share’s participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/participants
func (c_ CKShareParticipant) SetParticipants(value ICKShareParticipant) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setParticipants:"), value)
}/* debug [instance_properties/setter]: participants */


// The permission for anyone with access to the share’s URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/publicpermission
func (c_ CKShareParticipant) PublicPermission() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("publicPermission"))
	return rv
}/* debug [instance_properties/getter]: publicPermission */


// The permission for anyone with access to the share’s URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/publicpermission
func (c_ CKShareParticipant) SetPublicPermission(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPublicPermission:"), value)
}/* debug [instance_properties/setter]: publicPermission */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKShareParticipant */




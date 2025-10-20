// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CKShareParticipant] class.
type ICKShareParticipant interface {
	objectivec.IObject
}

// An object that describes a user’s participation in a share.
//
// Participants are a key element of sharing in CloudKit. A participant provides information about an iCloud user and their participation in a share, including their identity, acceptance status, permissions, and role. The acceptance status determines the participant’s visibilty of the shared records. Statuses are: , , , and . If the status is , use to accept the share. Upon acceptance, CloudKit makes the shared records available in the participant’s shared database. The records remain accessible for as long as the participant’s status is . You don’t create participants. Use the share’s property to access its existing participants. Use to manage the share’s participants and their permissions. Alternatively, you can generate participants using . Participants must have an active iCloud account. Anyone with the URL of a public share can become a participant in that share. Participants of a public share assume the role. For private shares, the owner manages the participants. An owner is any participant with the role. A participant of a private share can’t accept the share unless the owner adds them first. Private share participants assume the role. CloudKit removes any pending participants if the owner changes the share’s . CloudKit removes all participants if the new permission is . Participants with write permissions can modify or delete any record that you include in the share. However, only the owner can delete a shared hierarchy’s root record. If a participant attempts to delete the share, CloudKit removes the participant. The share remains active for all other participants.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CKShareParticipantClass) Alloc() CKShareParticipant {
	rv := objc.Send[CKShareParticipant](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The current state of the user’s acceptance of the share.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/acceptanceStatus-swift.property
func (c_ CKShareParticipant) AcceptanceStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("acceptanceStatus"))
	return rv
}

// The participant’s permission level for the share.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/permission-swift.property
func (c_ CKShareParticipant) Permission() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("permission"))
	return rv
}


// SetPermission sets the value of the permission property.
// The participant’s permission level for the share.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/permission-swift.property
func (c_ CKShareParticipant) SetPermission(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPermission:"), value)
}
// The participant’s role for the share.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/role-swift.property
func (c_ CKShareParticipant) Role() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("role"))
	return rv
}


// SetRole sets the value of the role property.
// The participant’s role for the share.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/role-swift.property
func (c_ CKShareParticipant) SetRole(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRole:"), value)
}
// The participant type.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/type
func (c_ CKShareParticipant) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// The participant type.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/type
func (c_ CKShareParticipant) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setType:"), value)
}
// The identity of the participant.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/userIdentity
func (c_ CKShareParticipant) UserIdentity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("userIdentity"))
	return rv
}




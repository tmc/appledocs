// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKAllowedSharingOptions] class.
var (
	CKAllowedSharingOptionsClass     _CKAllowedSharingOptionsClass
	CKAllowedSharingOptionsClassOnce sync.Once
)

func getCKAllowedSharingOptionsClass() _CKAllowedSharingOptionsClass {
	CKAllowedSharingOptionsClassOnce.Do(func() {
		CKAllowedSharingOptionsClass = _CKAllowedSharingOptionsClass{objc.GetClass("CKAllowedSharingOptions")}
	})
	return CKAllowedSharingOptionsClass
}

type _CKAllowedSharingOptionsClass struct {
	class objc.Class
}

// An interface definition for the [CKAllowedSharingOptions] class.
type ICKAllowedSharingOptions interface {
	objectivec.IObject
	// properties:
	AllowedParticipantAccessOptions() unsafe.Pointer
	SetAllowedParticipantAccessOptions(value unsafe.Pointer)
	AllowedParticipantPermissionOptions() CKSharingParticipantPermissionOption
	SetAllowedParticipantPermissionOptions(value CKSharingParticipantPermissionOption)
	AllowsAccessRequests() bool
	SetAllowsAccessRequests(value bool)
	AllowsParticipantsToInviteOthers() bool
	SetAllowsParticipantsToInviteOthers(value bool)
	// methods:
}

// An object that controls participant access and permission options.
//
// Register an instance of this class with an or when preparing a before your app invokes the share sheet. The share sheet uses the registered   object to let the user choose between the allowed options when sharing.


// An object that controls participant access and permission options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions
type CKAllowedSharingOptions struct {
	objectivec.Object
}

// CKAllowedSharingOptionsFrom constructs a [CKAllowedSharingOptions] from an unsafe.Pointer.
//
// An object that controls participant access and permission options.
func CKAllowedSharingOptionsFrom(ptr unsafe.Pointer) CKAllowedSharingOptions {
	return CKAllowedSharingOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKAllowedSharingOptionsClass) Alloc() CKAllowedSharingOptions {
	rv := objc.Send[CKAllowedSharingOptions](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKAllowedSharingOptionsClass) New() CKAllowedSharingOptions {
	rv := objc.Send[CKAllowedSharingOptions](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKAllowedSharingOptions) Init() CKAllowedSharingOptions {
	rv := objc.Send[CKAllowedSharingOptions](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKAllowedSharingOptions) Autorelease() CKAllowedSharingOptions {
	rv := objc.Send[CKAllowedSharingOptions](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKAllowedSharingOptions creates a new CKAllowedSharingOptions instance.
func NewCKAllowedSharingOptions() CKAllowedSharingOptions {
	return getCKAllowedSharingOptionsClass().New()
}



// The permission option the system uses to control whether a user can share publicly or privately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckallowedsharingoptions/allowedparticipantaccessoptions
func (c_ CKAllowedSharingOptions) AllowedParticipantAccessOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("allowedParticipantAccessOptions"))
	return rv
}


// The permission option the system uses to control whether a user can share publicly or privately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckallowedsharingoptions/allowedparticipantaccessoptions
func (c_ CKAllowedSharingOptions) SetAllowedParticipantAccessOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowedParticipantAccessOptions:"), value)
}


// The permission option the system uses to control whether a user can grant read-only or write access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckallowedsharingoptions/allowedparticipantpermissionoptions
func (c_ CKAllowedSharingOptions) AllowedParticipantPermissionOptions() CKSharingParticipantPermissionOption {
	rv := objc.Send[CKSharingParticipantPermissionOption](c_.ID, objc.Sel("allowedParticipantPermissionOptions"))
	return rv
}


// The permission option the system uses to control whether a user can grant read-only or write access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckallowedsharingoptions/allowedparticipantpermissionoptions
func (c_ CKAllowedSharingOptions) SetAllowedParticipantPermissionOptions(value CKSharingParticipantPermissionOption) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowedParticipantPermissionOptions:"), value)
}


// Default value is
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckallowedsharingoptions/allowsaccessrequests
func (c_ CKAllowedSharingOptions) AllowsAccessRequests() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsAccessRequests"))
	return rv
}


// Default value is
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckallowedsharingoptions/allowsaccessrequests
func (c_ CKAllowedSharingOptions) SetAllowsAccessRequests(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsAccessRequests:"), value)
}


// Default value is
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckallowedsharingoptions/allowsparticipantstoinviteothers
func (c_ CKAllowedSharingOptions) AllowsParticipantsToInviteOthers() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsParticipantsToInviteOthers"))
	return rv
}


// Default value is
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckallowedsharingoptions/allowsparticipantstoinviteothers
func (c_ CKAllowedSharingOptions) SetAllowsParticipantsToInviteOthers(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsParticipantsToInviteOthers:"), value)
}




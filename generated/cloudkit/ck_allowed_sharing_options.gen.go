// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKAllowedSharingOptions */


/* debug [class_header]: Header for CKAllowedSharingOptions */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKAllowedSharingOptions */
// An interface definition for the [CKAllowedSharingOptions] class.
type ICKAllowedSharingOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKAllowedSharingOptions */
	// properties:
	AllowedParticipantAccessOptions() CKSharingParticipantAccessOption
	SetAllowedParticipantAccessOptions(value CKSharingParticipantAccessOption)
	AllowedParticipantPermissionOptions() CKSharingParticipantPermissionOption
	SetAllowedParticipantPermissionOptions(value CKSharingParticipantPermissionOption)
	AllowsAccessRequests() bool
	SetAllowsAccessRequests(value bool)
	AllowsParticipantsToInviteOthers() bool
	SetAllowsParticipantsToInviteOthers(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKAllowedSharingOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKAllowedSharingOptions */
// Alloc allocates a new instance without initialization.
func (cc _CKAllowedSharingOptionsClass) Alloc() CKAllowedSharingOptions {
	rv := objc.Send[CKAllowedSharingOptions](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKAllowedSharingOptions */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKAllowedSharingOptions */

// Creates and initializes an allowed sharing options object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/init(allowedParticipantPermissionOptions:allowedParticipantAccessOptions:)
func NewCKAllowedSharingOptionsWithAllowedParticipantPermissionOptionsAllowedParticipantAccessOptions(allowedParticipantPermissionOptions CKSharingParticipantPermissionOption, allowedParticipantAccessOptions CKSharingParticipantAccessOption) CKAllowedSharingOptions {
	instance := getCKAllowedSharingOptionsClass().Alloc()
	rv := objc.Send[CKAllowedSharingOptions](instance.ID, objc.Sel("initWithAllowedParticipantPermissionOptions:allowedParticipantAccessOptions:"), allowedParticipantPermissionOptions, allowedParticipantAccessOptions)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKAllowedSharingOptionsWithAllowedParticipantPermissionOptionsAllowedParticipantAccessOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKAllowedSharingOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKAllowedSharingOptions */

// An object set to the most permissive sharing options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/standard
func (cc _CKAllowedSharingOptionsClass) StandardOptions() CKAllowedSharingOptions {
	rv := objc.Send[CKAllowedSharingOptions](objc.ID(cc.class), objc.Sel("standardOptions"))
	return rv
}/* debug [class_properties_class/property]: standardOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKAllowedSharingOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKAllowedSharingOptions */

// The permission option the system uses to control whether a user can share publicly or privately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/allowedParticipantAccessOptions
func (c_ CKAllowedSharingOptions) AllowedParticipantAccessOptions() CKSharingParticipantAccessOption {
	rv := objc.Send[CKSharingParticipantAccessOption](c_.ID, objc.Sel("allowedParticipantAccessOptions"))
	return rv
}/* debug [instance_properties/getter]: allowedParticipantAccessOptions */


// The permission option the system uses to control whether a user can share publicly or privately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/allowedParticipantAccessOptions
func (c_ CKAllowedSharingOptions) SetAllowedParticipantAccessOptions(value CKSharingParticipantAccessOption) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowedParticipantAccessOptions:"), value)
}/* debug [instance_properties/setter]: allowedParticipantAccessOptions */


// The permission option the system uses to control whether a user can grant read-only or write access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/allowedParticipantPermissionOptions
func (c_ CKAllowedSharingOptions) AllowedParticipantPermissionOptions() CKSharingParticipantPermissionOption {
	rv := objc.Send[CKSharingParticipantPermissionOption](c_.ID, objc.Sel("allowedParticipantPermissionOptions"))
	return rv
}/* debug [instance_properties/getter]: allowedParticipantPermissionOptions */


// The permission option the system uses to control whether a user can grant read-only or write access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/allowedParticipantPermissionOptions
func (c_ CKAllowedSharingOptions) SetAllowedParticipantPermissionOptions(value CKSharingParticipantPermissionOption) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowedParticipantPermissionOptions:"), value)
}/* debug [instance_properties/setter]: allowedParticipantPermissionOptions */


// Default value is . If set, the system sharing UI will allow the user to configure whether access requests are enabled on the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/allowsAccessRequests
func (c_ CKAllowedSharingOptions) AllowsAccessRequests() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsAccessRequests"))
	return rv
}/* debug [instance_properties/getter]: allowsAccessRequests */


// Default value is . If set, the system sharing UI will allow the user to configure whether access requests are enabled on the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/allowsAccessRequests
func (c_ CKAllowedSharingOptions) SetAllowsAccessRequests(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsAccessRequests:"), value)
}/* debug [instance_properties/setter]: allowsAccessRequests */


// Default value is . If set, the system sharing UI will allow the user to choose whether added participants can invite others to the share. Shares with participants will be returned as read-only to devices running OS versions prior to this role being introduced. Administrator participants on these read-only shares will be returned as .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/allowsParticipantsToInviteOthers
func (c_ CKAllowedSharingOptions) AllowsParticipantsToInviteOthers() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsParticipantsToInviteOthers"))
	return rv
}/* debug [instance_properties/getter]: allowsParticipantsToInviteOthers */


// Default value is . If set, the system sharing UI will allow the user to choose whether added participants can invite others to the share. Shares with participants will be returned as read-only to devices running OS versions prior to this role being introduced. Administrator participants on these read-only shares will be returned as .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/allowsParticipantsToInviteOthers
func (c_ CKAllowedSharingOptions) SetAllowsParticipantsToInviteOthers(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsParticipantsToInviteOthers:"), value)
}/* debug [instance_properties/setter]: allowsParticipantsToInviteOthers */


// An object set to the most permissive sharing options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/standard
func (c_ CKAllowedSharingOptions) StandardOptions() ICKAllowedSharingOptions {
	rv := objc.Send[CKAllowedSharingOptions](c_.ID, objc.Sel("standardOptions"))
	return rv
}/* debug [instance_properties/getter]: standardOptions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKAllowedSharingOptions */



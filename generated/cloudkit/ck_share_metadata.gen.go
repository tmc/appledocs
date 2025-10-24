// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKShareMetadata */


/* debug [class_header]: Header for CKShareMetadata */
// The class instance for the [CKShareMetadata] class.
var (
	CKShareMetadataClass     _CKShareMetadataClass
	CKShareMetadataClassOnce sync.Once
)

func getCKShareMetadataClass() _CKShareMetadataClass {
	CKShareMetadataClassOnce.Do(func() {
		CKShareMetadataClass = _CKShareMetadataClass{objc.GetClass("CKShareMetadata")}
	})
	return CKShareMetadataClass
}

type _CKShareMetadataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKShareMetadata */
// An interface definition for the [CKShareMetadata] class.
type ICKShareMetadata interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKShareMetadata */
	// properties:
	ContainerIdentifier() objc.IObject /* cross-framework: NSString */
	HierarchicalRootRecordID() ICKRecordID
	OwnerIdentity() ICKUserIdentity
	ParticipantPermission() CKShareParticipantPermission
	ParticipantRole() CKShareParticipantRole
	ParticipantStatus() CKShareParticipantAcceptanceStatus
	ParticipantType() CKShareParticipantType
	RootRecord() objc.IObject /* cross-framework: CKRecord */
	RootRecordID() ICKRecordID
	Share() ICKShare
	RootRecordDesiredKeys() objectivec.IObject
	SetRootRecordDesiredKeys(value objectivec.IObject)
	ShouldFetchRootRecord() bool
	SetShouldFetchRootRecord(value bool)
	Url() foundation.URL
	SetUrl(value foundation.URL)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKShareMetadata */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKShareMetadata */
// Alloc allocates a new instance without initialization.
func (cc _CKShareMetadataClass) Alloc() CKShareMetadata {
	rv := objc.Send[CKShareMetadata](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKShareMetadataClass) New() CKShareMetadata {
	rv := objc.Send[CKShareMetadata](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKShareMetadata) Init() CKShareMetadata {
	rv := objc.Send[CKShareMetadata](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKShareMetadata) Autorelease() CKShareMetadata {
	rv := objc.Send[CKShareMetadata](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKShareMetadata creates a new CKShareMetadata instance.
func NewCKShareMetadata() CKShareMetadata {
	return getCKShareMetadataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKShareMetadata */
// An object that describes a shared record’s metadata.
//
// A share’s metadata is an intermediary object that provides access to the share, its owner, and, for a shared record hierarchy, its root record. Metadata also includes details about the current user’s participation in the share. You don’t create metadata. CloudKit provides it to your app when the user taps or clicks a share’s , such as in an email or a message. The method CloudKit calls varies by platform and app configuration, and includes the following: For a scene-based iOS app in a running or suspended state, CloudKit calls the method on your window scene delegate. For a scene-based iOS app that’s not running, the system launches your app in response to the tap or click, and calls the method on your scene delegate. The parameter contains the metadata. Use its property to access it. For an iOS app that doesn’t use scenes, CloudKit calls your app delegate’s method. For a macOS app, CloudKit calls your app delegate’s method. For a watchOS app, CloudKit calls the method on your watch extension delegate. Respond by checking the of the provided metadata. If the status is , use to accept participation in the share. You can also fetch metadata independent of this flow using . For a shared record hierarchy, the property contains the ID of the share’s root record. When using to fetch metadata, you can include the entire root record by setting the operation’s property to . CloudKit then populates the property before it returns the metadata. You can further customize this behavior using the operation’s property to specify which fields to return. This functionality isn’t applicable for a shared record zone because, unlike a shared record hierarchy, it doesn’t have a nominated root record. The participant properties provide the current user’s acceptance status, permissions, and role. Use these values to determine what functionality to provide to the user. For example, only display editing controls for accepted participants with permissions.


// An object that describes a shared record’s metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata
type CKShareMetadata struct {
	objectivec.Object
}

// CKShareMetadataFrom constructs a [CKShareMetadata] from an unsafe.Pointer.
//
// An object that describes a shared record’s metadata.
func CKShareMetadataFrom(ptr unsafe.Pointer) CKShareMetadata {
	return CKShareMetadata{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKShareMetadata *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKShareMetadata */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKShareMetadata */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKShareMetadata */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKShareMetadata */

// The ID of the share’s container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/containerIdentifier
func (c_ CKShareMetadata) ContainerIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}/* debug [instance_properties/getter]: containerIdentifier */


// The record ID of the shared hierarchy’s root record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/hierarchicalRootRecordID
func (c_ CKShareMetadata) HierarchicalRootRecordID() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("hierarchicalRootRecordID"))
	return rv
}/* debug [instance_properties/getter]: hierarchicalRootRecordID */


// The identity of the share’s owner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/ownerIdentity
func (c_ CKShareMetadata) OwnerIdentity() ICKUserIdentity {
	rv := objc.Send[CKUserIdentity](c_.ID, objc.Sel("ownerIdentity"))
	return rv
}/* debug [instance_properties/getter]: ownerIdentity */


// The share’s permissions for the user who retrieves the metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/participantPermission
func (c_ CKShareMetadata) ParticipantPermission() CKShareParticipantPermission {
	rv := objc.Send[CKShareParticipantPermission](c_.ID, objc.Sel("participantPermission"))
	return rv
}/* debug [instance_properties/getter]: participantPermission */


// The share’s participant role for the user who retrieves the metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/participantRole
func (c_ CKShareMetadata) ParticipantRole() CKShareParticipantRole {
	rv := objc.Send[CKShareParticipantRole](c_.ID, objc.Sel("participantRole"))
	return rv
}/* debug [instance_properties/getter]: participantRole */


// The share’s participation status for the user who retrieves the metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/participantStatus
func (c_ CKShareMetadata) ParticipantStatus() CKShareParticipantAcceptanceStatus {
	rv := objc.Send[CKShareParticipantAcceptanceStatus](c_.ID, objc.Sel("participantStatus"))
	return rv
}/* debug [instance_properties/getter]: participantStatus */


// The share’s participation type for the user who retrieves the metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/participantType
func (c_ CKShareMetadata) ParticipantType() CKShareParticipantType {
	rv := objc.Send[CKShareParticipantType](c_.ID, objc.Sel("participantType"))
	return rv
}/* debug [instance_properties/getter]: participantType */


// The share’s root record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/rootRecord
func (c_ CKShareMetadata) RootRecord() objc.IObject /* cross-framework: CKRecord */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("rootRecord"))
	return rv
}/* debug [instance_properties/getter]: rootRecord */


// The record ID of the share’s root record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/rootRecordID
func (c_ CKShareMetadata) RootRecordID() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("rootRecordID"))
	return rv
}/* debug [instance_properties/getter]: rootRecordID */


// The share that owns the metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/share
func (c_ CKShareMetadata) Share() ICKShare {
	rv := objc.Send[CKShare](c_.ID, objc.Sel("share"))
	return rv
}/* debug [instance_properties/getter]: share */


// The fields to return when fetching the root record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/rootrecorddesiredkeys-3xrex
func (c_ CKShareMetadata) RootRecordDesiredKeys() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("rootRecordDesiredKeys"))
	return rv
}/* debug [instance_properties/getter]: rootRecordDesiredKeys */


// The fields to return when fetching the root record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/rootrecorddesiredkeys-3xrex
func (c_ CKShareMetadata) SetRootRecordDesiredKeys(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRootRecordDesiredKeys:"), value)
}/* debug [instance_properties/setter]: rootRecordDesiredKeys */


// A Boolean value that indicates whether to retrieve the root record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/shouldfetchrootrecord
func (c_ CKShareMetadata) ShouldFetchRootRecord() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldFetchRootRecord"))
	return rv
}/* debug [instance_properties/getter]: shouldFetchRootRecord */


// A Boolean value that indicates whether to retrieve the root record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchsharemetadataoperation/shouldfetchrootrecord
func (c_ CKShareMetadata) SetShouldFetchRootRecord(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldFetchRootRecord:"), value)
}/* debug [instance_properties/setter]: shouldFetchRootRecord */


// The URL for inviting participants to the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/url
func (c_ CKShareMetadata) Url() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// The URL for inviting participants to the share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshare/url
func (c_ CKShareMetadata) SetUrl(value foundation.URL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUrl:"), value)
}/* debug [instance_properties/setter]: url */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKShareMetadata */



